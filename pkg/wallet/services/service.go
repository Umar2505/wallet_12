package service

import (
	"errors"

	"github.com/Umar2505/wallet_12/pkg/wallet/types"
	"github.com/google/uuid"
)

var ErrPaymentNotFound = errors.New("payment not found")

type Service struct {
	nextAccountID int64
	nextPaymentID string
	accounts []*types.Account
	payments []*types.Payment
	favorites []*types.Favorite
}

type testService struct {
	*Service
}

func NewTestService() *testService {
	return &testService{Service: &Service{}}
}

func (s *testService) Register(phone types.Phone) (*types.Account,error) {
	for _, v := range s.accounts {
		if v.Phone==phone {
			return v,errors.New("phone is existed")
		}
	}
	s.nextAccountID++
	acc:= types.Account{
		ID: s.nextAccountID,
		Phone: phone,
		Balance: 0,
	}
	s.accounts=append(s.accounts, &acc)
	return &acc,nil
}

func (s *testService,) Deposit(accountID int64,amount types.Money) (*types.Payment,error) {
	var payment types.Payment
	if amount<=0 {
		return nil, errors.New("amount should be greater than 0")
	}
	
	var targetAccount *types.Account
	for _, v := range s.accounts {
		if v.ID==accountID {
			targetAccount=v
		}
	}
	if targetAccount==nil {
		return nil, errors.New("account not found")
	}

	targetAccount.Balance+=amount
	s.nextPaymentID=uuid.NewString()
	payment.ID=s.nextPaymentID
	payment.AccountID=targetAccount.ID
	payment.Amount=amount
	s.payments=append(s.payments, &payment)
	return &payment,nil
}

func (s *testService) Reject(paymentID string) error  {
	for _, v := range s.payments {
		if v.ID==paymentID {
			for _, a := range s.accounts {
				if v.AccountID==a.ID {
					a.Balance+=v.Amount
					return nil
				}
			}
		}
	}
	return  ErrPaymentNotFound
}

func (s *testService) Repeat(paymentID string) (*types.Payment,error) {
	targetP:= types.Payment{}
	targetPayment:=&targetP
	for _, v := range s.payments {
		if v.ID==paymentID {
			targetPayment.AccountID=v.AccountID
			targetPayment.Amount=v.Amount
			targetPayment.Category=v.Category
			targetPayment.Status=v.Status
			s.nextPaymentID=uuid.NewString()
			targetPayment.ID=s.nextPaymentID
		}
	}

	if targetPayment==&targetP {
		return nil, errors.New("payment not found")
	}

	return targetPayment,nil
}

func (s *testService) FavoritePayment(paymentID string, name string) (*types.Favorite, error) {
	var tPayment *types.Payment
	for _, payment := range s.payments {
		if payment.ID==paymentID {
			tPayment=payment
		}
	}
	if tPayment==nil {
		return nil, errors.New("couldn't find payment")
	}

	var favorite *types.Favorite
	favorite.AccountID= tPayment.AccountID
	favorite.Amount=tPayment.Amount
	favorite.Category=tPayment.Category
	favorite.ID=tPayment.ID
	favorite.Name=name
	s.favorites=append(s.favorites, favorite)
	return favorite,nil
}

func (s *testService) PayFromFavorite(favoriteID string) (*types.Payment, error) {
	var tFavorite *types.Favorite
	for _, favorite := range s.favorites {
		if favorite.ID==favoriteID {
			tFavorite=favorite
		}
	}

	if tFavorite==nil {
		return nil,errors.New("couldn't find favorite")
	}

	payment,err:= s.Deposit(tFavorite.AccountID,tFavorite.Amount)
	if err!=nil {
		return nil,err
	}
	return payment,nil
}
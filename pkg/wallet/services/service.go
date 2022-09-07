package service

import (
	"errors"

	"github.com/Umar2505/wallet_12/pkg/wallet/types"
)

var ErrPaymentNotFound = errors.New("payment not found")

type Service struct {
	nextAccountID int64
	accounts []*types.Account
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

func (s *testService,) Deposit(accountID int64,amount types.Money) (error) {
	if amount<=0 {
		return errors.New("amount should be greater than 0")
	}
	
	var targetAccount *types.Account
	for _, v := range s.accounts {
		if v.ID==accountID {
			targetAccount=v
		}
	}
	if targetAccount==nil {
		return errors.New("account not found")
	}

	targetAccount.Balance+=amount
	return nil
}
package service

import (
	"errors"

	"github.com/Umar2505/wallet_12/pkg/wallet/types"
)

var ErrPaymentNotFound = errors.New("payment not found")

type Service struct {
	accounts []*types.Account
	payments []*types.Payment
}

func (s *Service) Reject(paymentID string) error  {
	for _, v := range s.payments {
		if v.ID==paymentID {
			for _, a := range s.accounts {
				if v.AccountID==a.ID {
					a.Balance+=v.Amount
					v.Status=types.PaymentStatusFail
					return nil
				}
			}
		}
	}
	return  ErrPaymentNotFound
}
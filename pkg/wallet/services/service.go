package service

import (
	"errors"

	"github.com/Umar2505/wallet_12/pkg/wallet/types"
)

type Service struct {
	accounts []*types.Account
}

func (s *Service) FindAccountByID(accountID int64) (*types.Account, error) {
	for _, v := range s.accounts {
		if v.ID==accountID {
			return v,nil
		}
	}
	return nil,errors.New("account not found")
}
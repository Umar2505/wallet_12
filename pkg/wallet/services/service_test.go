package service

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Umar2505/wallet_12/pkg/wallet/types"
)

var s *Service = new(Service)

var accounts = []*types.Account{
	{
		ID: 1,
		Balance: 195_000,
	},
	{
		ID: 2,
		Balance: 50_000,
	},
	{
		ID: 3,
		Balance: 0,
	},
	{
		ID: 4,
		Balance: 100_000,
	},
}

var payments= []*types.Payment{
	{
		ID: "10",
		Amount: 1000,
		AccountID: 3,
	},
	{
		ID: "11",
		Amount: 5_000,
		AccountID: 1,
	},
	{
		ID: "12",
		Amount: 2_000,
		AccountID: 2,
	},
}

func TestService_Reject_positive(t *testing.T) {
	
	s.accounts=accounts
	s.payments=payments

	ex:= types.Account{
		ID: 1,
		Balance: 20checkout 0_000,
	}
	expected := &ex

	err:=s.Reject("11")
	if err!=nil {
		fmt.Println(err)
		return
	}

	if !reflect.DeepEqual(s.accounts[0],expected) {
		fmt.Printf("Expected: %v, recieved: %v",expected,s.accounts[0])
		return
	}
}

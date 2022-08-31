package service

import (
	"fmt"
	"testing"

	"github.com/Umar2505/wallet_12/pkg/wallet/types"
)

var accounts = []*types.Account{
	{
		ID: 1,
		Balance: 200_000,
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

func TestService_FindAccountByID_positive(t *testing.T) {
	
	s:= &Service{
		accounts: accounts,
	}
	

	expected := types.Account{
		ID: 2,
		Balance: 50_000,
	}

	result,err:=s.FindAccountByID(2)
	if err!=nil {
		t.Error(err.Error())
	}

	if result.ID!=expected.ID {
		fmt.Printf("Expected: %v, Recieved: %v",expected,result)
	}
}

func TestService_FindAccountByID_notFound(t *testing.T) {
	
	s:= &Service{
		accounts: accounts,
	}
	

	expected := types.Account{
		ID: 0,
		Balance: 10_000,
	}

	result,err:=s.FindAccountByID(0)
	if err!=nil {
		t.Error(err.Error())
	}

	if result.ID!=expected.ID && result.Balance!=expected.Balance {
		fmt.Printf("Expected: %v, Recieved: %v",expected,result)
	}
}
package service

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRepeat_positive(t *testing.T) {
	s:= NewTestService()
	acc,err:=s.Register("+123123123")
	if err!=nil {
		fmt.Println(err.Error())
		return
	}
	pay,err:=s.Deposit(acc.ID,1000)
	if err!=nil {
		fmt.Println(err.Error())
		return
	}

	expect:= *pay
	add,_:=strconv.Atoi(expect.ID)
	expect.ID=strconv.Itoa(add+1)

	payment,err:=s.Repeat(pay.ID)
	if err!=nil {
		fmt.Println(err.Error())
		return
	}

	if expect.AccountID!=payment.AccountID && expect.Amount!=payment.Amount {
		fmt.Println("your func() isn't work")
		return
	}
}


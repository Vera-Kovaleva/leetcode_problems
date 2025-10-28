package main

/*
type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	return Bank{
		balance: balance,
	}
}

func (bank Bank) Transfer(accNumFrom int, accNumTo int, money int64) bool {
	n := len(bank.balance)
	if accNumFrom > n || accNumTo > n || bank.balance[accNumFrom-1] < money {
		return false
	}
	bank.balance[accNumFrom-1] -= money
	bank.balance[accNumTo-1] += money
	return true
}

func (bank Bank) Deposit(accNum int, money int64) bool {
	if accNum > len(bank.balance) {
		return false
	}
	bank.balance[accNum-1] += money
	return true
}

func (bank Bank) Withdraw(accNum int, money int64) bool {
	if accNum > len(bank.balance) || bank.balance[accNum-1] < money {
		return false
	}
	bank.balance[accNum-1] -= money
	return true
}*/

type Bank []int64

func Constructor(balance []int64) Bank {
	return balance
}

func (bank Bank) Transfer(accNumFrom int, accNumTo int, money int64) bool {
	if accNumFrom > len(bank) || accNumTo > len(bank) || bank[accNumFrom-1] < money {
		return false
	}
	bank[accNumFrom-1] -= money
	bank[accNumTo-1] += money
	return true
}

func (bank Bank) Deposit(accNum int, money int64) bool {
	if accNum > len(bank) {
		return false
	}
	bank[accNum-1] += money
	return true
}

func (bank Bank) Withdraw(accNum int, money int64) bool {
	if accNum > len(bank) || bank[accNum-1] < money {
		return false
	}
	bank[accNum-1] -= money
	return true
}

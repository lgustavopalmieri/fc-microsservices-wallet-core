package entity

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Client    *Client
	ClientID  string
	Balance   float64
	CreatedAt string
	UpdatedAt string
}

func NewAccount(client *Client) *Account {
	currentTime := time.Now()
	if client == nil {
		return nil
	}
	account := &Account{
		ID:        uuid.New().String(),
		Client:    client,
		Balance:   0,
		CreatedAt: currentTime.Format("2006-01-02T15:04:05"),
		UpdatedAt: currentTime.Format("2006-01-02T15:04:05"),
	}

	return account
}

func (a *Account) Credit(amount float64) {
	currentTime := time.Now()
	a.Balance += amount
	a.UpdatedAt = currentTime.Format("2006-01-02T15:04:05")
}

func (a *Account) Debit(amount float64) {
	currentTime := time.Now()
	a.Balance -= amount
	a.UpdatedAt = currentTime.Format("2006-01-02T15:04:05")
}

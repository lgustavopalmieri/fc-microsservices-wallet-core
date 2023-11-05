package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "jondoe@mail.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithNilClietn(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "jondoe@mail.com")
	account := NewAccount(client)
	assert.Equal(t, float64(0), account.Balance)
	account.Credit(100)
	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "jondoe@mail.com")
	account := NewAccount(client)
	account.Credit(100)
	assert.Equal(t, float64(100), account.Balance)
	account.Debit(50)
	assert.Equal(t, float64(50), account.Balance)
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("John Doe", "jondoe@mail.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}

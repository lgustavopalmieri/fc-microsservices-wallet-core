package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string
	Name      string
	Email     string
	Accounts  []*Account
	CreatedAt string
	UpdatedAt string
}

func NewClient(name, email string) (*Client, error) {
	currentTime := time.Now()
	client := &Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: currentTime.Format(time.RFC3339),
		UpdatedAt: currentTime.Format(time.RFC3339),
	}
	err := client.Validate()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.Email == "" {
		return errors.New("email is required")
	}
	return nil
}
func (c *Client) Update(name, email string) error {
	currentTime := time.Now()
	c.Name = name
	c.Email = email
	c.UpdatedAt = currentTime.Format(time.RFC3339)
	err := c.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AddAccount(account *Account) error {
	if account.Client.ID != c.ID {
		return errors.New("account does not belong to client")
	}
	c.Accounts = append(c.Accounts, account)
	return nil
}


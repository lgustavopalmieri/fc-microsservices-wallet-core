package gateway

import "github.com/lgustavopalmieri/fc-microsservice-wallet-core/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}


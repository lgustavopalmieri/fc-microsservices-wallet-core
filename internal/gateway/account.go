package gateway

import "github.com/lgustavopalmieri/fc-microsservice-wallet-core/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}

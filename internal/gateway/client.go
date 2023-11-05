package gateway

import "github.com/lgustavopalmieri/fc-microsservice-wallet-core/internal/entity"

type ClientGateway interface {
	GetById(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}

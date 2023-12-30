package create_account

import (
	"testing"

	"github.com/lgustavopalmieri/fc-microsservice-wallet-core/internal/entity"
	"github.com/lgustavopalmieri/fc-microsservice-wallet-core/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "j@j")
	clientMock := &mocks.ClientGatewayMock{}
	clientMock.On("GetById", client.ID).Return(client, nil)

	accountMock := &mocks.AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountMock, clientMock)
	inputDto := CreateAccountInputDTO{
		ClientID: client.ID,
	}
	outputDto, err := uc.Execute(inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, outputDto.ID)
	clientMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)
	clientMock.AssertNumberOfCalls(t, "GetById", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
}

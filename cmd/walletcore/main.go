package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lgustavopalmieri/fc-microsservice-wallet-core/internal/database"
	"github.com/lgustavopalmieri/fc-microsservice-wallet-core/internal/event"
	createaccount "github.com/lgustavopalmieri/fc-microsservice-wallet-core/internal/usecase/create_account"
	createclient "github.com/lgustavopalmieri/fc-microsservice-wallet-core/internal/usecase/create_client"
	createtransaction "github.com/lgustavopalmieri/fc-microsservice-wallet-core/internal/usecase/create_transaction"
	"github.com/lgustavopalmieri/fc-microsservice-wallet-core/internal/web"
	"github.com/lgustavopalmieri/fc-microsservice-wallet-core/internal/web/webserver"
	"github.com/lgustavopalmieri/fc-microsservice-wallet-core/pkg/events"
	"github.com/lgustavopalmieri/fc-microsservice-wallet-core/pkg/uow"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/wallet")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	// eventDispatcher.Register("TransactionCreated", handler)
	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	//transactionDb := database.NewTransactionDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent)

	webserver := webserver.NewWebServer(":8080")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running")
	webserver.Start()
}

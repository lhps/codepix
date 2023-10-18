package model_test

import (
	uuid "github.com/satori/go.uuid"
	"testing"

	"github.com/lhps/codepix-go/domain/model"
	"github.com/stretchr/testify/require"
)

func TestNewTransaction(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Lucas"
	account, err := model.NewAccount(bank, accountNumber, ownerName)

	accountNumberDestination := "abcdestination"
	ownerName = "Camilla"
	accountDestination, _ := model.NewAccount(bank, accountNumberDestination, ownerName)

	kind := "email"
	key := "l@pinho.com"
	pixKey, err := model.NewPixKey(kind, key, accountDestination)

	require.NotEqual(t, account.ID, accountDestination.ID)

	amount := 3.10
	statusTransaction := "pending"
	transaction, err := model.NewTransaction(account, amount, pixKey, "Descrição")

	require.Nil(t, err)
	require.NotNil(t, uuid.FromStringOrNil(transaction.ID))
	require.Equal(t, transaction.Amount, amount)
	require.Equal(t, transaction.Status, statusTransaction)
	require.Equal(t, transaction.Description, "Descrição")
	require.Empty(t, transaction.CancelDescription)

	pixKeySameAccount, err := model.NewPixKey(kind, key, account)

	_, err = model.NewTransaction(account, amount, pixKeySameAccount, "Descrição")
	require.NotNil(t, err)

	_, err = model.NewTransaction(account, 0, pixKey, "Descrição")
	require.NotNil(t, err)
}

func TestModel_ChangeStatusOfATransaction(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Lucas"
	account, _ := model.NewAccount(bank, accountNumber, ownerName)

	accountNumberDestination := "abcdestination"
	ownerName = "Camilla"
	accountDestination, _ := model.NewAccount(bank, accountNumberDestination, ownerName)

	kind := "email"
	key := "l@pinho.com"
	pixKey, _ := model.NewPixKey(kind, key, accountDestination)

	amount := 3.10
	transaction, _ := model.NewTransaction(account, amount, pixKey, "Descrição")

	transaction.Complete()
	require.Equal(t, transaction.Status, model.TransactionCompleted)

	transaction.Cancel("Error")
	require.Equal(t, transaction.Status, model.TransactionError)
	require.Equal(t, transaction.CancelDescription, "Error")
}

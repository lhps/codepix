package model_test

import (
	uuid "github.com/satori/go.uuid"
	"testing"

	"github.com/lhps/codepix-go/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewAccount(t *testing.T) {

	code := "001"
	name := "Banco do Brasil"
	bank, err := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Lucas"
	account, err := model.NewAccount(bank, accountNumber, ownerName)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(account.ID))
	require.Equal(t, account.Number, accountNumber)
	require.Equal(t, account.Bank.ID, bank.ID)

	_, err = model.NewAccount(bank, "", ownerName)
	require.NotNil(t, err)
}

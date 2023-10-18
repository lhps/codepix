package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/lhps/codepix-go/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewPixKey(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Lucas"
	account, err := model.NewAccount(bank, accountNumber, ownerName)

	kind := "email"
	key := "l@pinho.com"
	pixKey, err := model.NewPixKey(kind, key, account)

	require.NotEmpty(t, uuid.FromStringOrNil(pixKey.ID))
	require.Equal(t, pixKey.Kind, kind)
	require.Equal(t, pixKey.Status, "active")

	kind = "cpf"
	_, err = model.NewPixKey(kind, key, account)
	require.Nil(t, err)
	
	kind = "invalidkind"
	_, err = model.NewPixKey(kind, key, account)
	require.NotNil(t, err)

	_, err = model.NewPixKey("nome", key, account)
	require.NotNil(t, err)
}

package repository

import (
	"Bank-INA/database"
	"Bank-INA/domain"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Join(filepath.Dir(b), "../../../../")
	_          = godotenv.Load(basepath + "/.env")
	db         = database.NewConnMysql(1)
	repo       = NewUserRepository(db)

	user = domain.User{
		Name:     "alfisar",
		Email:    "testing@gmail.com",
		Password: "Password12345",
	}
)

func TestCreateUser(t *testing.T) {
	err := repo.CreateUser(user)
	require.Equal(t, domain.ErrorData{}, err)
}

func TestCreateUserFailed(t *testing.T) {
	err := repo.CreateUser(user)
	require.NotEqual(t, domain.ErrorData{}, err)
}

func TestGetAllUser(t *testing.T) {
	_, err := repo.GetAll()
	require.Equal(t, domain.ErrorData{}, err)
}

func TestGetOneUser(t *testing.T) {
	_, err := repo.GetByID(18)
	require.Equal(t, domain.ErrorData{}, err)
}

func TestFailedGetOneUser(t *testing.T) {
	_, err := repo.GetByID(0)
	require.NotEqual(t, domain.ErrorData{}, err)
}

func TestUpdate(t *testing.T) {
	updates := map[string]interface{}{}
	updates["name"] = "alfi"
	updates["email"] = "yuhu@gmail.com"
	err := repo.UpdateByID(18, updates)
	require.Equal(t, domain.ErrorData{}, err)
}

func TestDelete(t *testing.T) {
	err := repo.DeleteByID(18)
	require.Equal(t, domain.ErrorData{}, err)
}

func TestFailedDelete(t *testing.T) {
	err := repo.DeleteByID(18)
	require.Equal(t, domain.ErrorData{}, err)
}

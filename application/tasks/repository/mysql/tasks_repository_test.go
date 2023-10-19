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
	repo       = NewTaskRepository(db)

	tasks = domain.Tasks{
		UserID: 20,
		Title:  "Approval Deposito",
		Desc:   "Dokumen approval untuk pembuatan deposito nasabah",
	}
)

func TestCreateTasks(t *testing.T) {
	err := repo.CreateTasks(tasks)
	require.Equal(t, domain.ErrorData{}, err)
}

func TestCreateTasksFailed(t *testing.T) {
	tasks.UserID = 100
	err := repo.CreateTasks(tasks)
	require.NotEqual(t, domain.ErrorData{}, err)
}

func TestGetAllTasks(t *testing.T) {
	_, err := repo.GetAll()
	require.Equal(t, domain.ErrorData{}, err)
}

func TestGetOneTasks(t *testing.T) {
	_, err := repo.GetByID(7)
	require.Equal(t, domain.ErrorData{}, err)
}

func TestFailedGetOneTasks(t *testing.T) {
	_, err := repo.GetByID(0)
	require.NotEqual(t, domain.ErrorData{}, err)
}

func TestUpdate(t *testing.T) {
	updates := map[string]interface{}{}
	updates["title"] = "Perubahan data deposito"
	updates["description"] = "Terdapat perubahan pembukaan dapositor nasabah"
	updates["status"] = "Acc Checker 1"
	err := repo.UpdateByID(7, updates)
	require.Equal(t, domain.ErrorData{}, err)
}

func TestDelete(t *testing.T) {
	err := repo.DeleteByID(7)
	require.Equal(t, domain.ErrorData{}, err)
}

func TestFailedDelete(t *testing.T) {
	err := repo.DeleteByID(7)
	require.Equal(t, domain.ErrorData{}, err)
}

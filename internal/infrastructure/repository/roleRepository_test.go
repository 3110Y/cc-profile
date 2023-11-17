package repository

import (
	"context"
	"github.com/3110Y/profile/internal/infrastructure/database"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var roleRepository RoleRepository

func init() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	connect, err = database.NewConnectTest()
	if err != nil {
		log.Fatal(err)
	}
	roleRepository = RoleRepository{db: connect}
	ctxRepository = context.Background()
}

func TestRoleRepository_Get(t *testing.T) {
	id := "7647316e-22e2-4f94-93bc-c0459dcb54de"
	roleItem, err := roleRepository.Get(ctxRepository, id)
	assert.Nil(t, err)
	assert.Equal(t, id, roleItem.Id)
}

func TestRoleRepository_List(t *testing.T) {
	roleList, err := roleRepository.List(ctxRepository, 2, 1)
	assert.Nil(t, err)
	assert.Len(t, *roleList, 2)
	roleList, err = roleRepository.List(ctxRepository, 1, 1)
	assert.Nil(t, err)
	assert.Len(t, *roleList, 1)
	roleList, err = roleRepository.List(ctxRepository, 1, 2)
	assert.Nil(t, err)
	assert.Len(t, *roleList, 1)
	roleList, err = roleRepository.List(ctxRepository, 10, 1)
	assert.Nil(t, err)
	assert.Len(t, *roleList, 2)
	roleList, err = roleRepository.List(ctxRepository, 10, 2)
	assert.Nil(t, err)
	assert.Len(t, *roleList, 0)
}

func TestRoleRepository_Count(t *testing.T) {
	count, err := roleRepository.Count(ctxRepository)
	assert.Nil(t, err)
	assert.Equal(t, uint64(2), *count)
}

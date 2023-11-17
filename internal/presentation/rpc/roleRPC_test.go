package rpc

import (
	"context"
	"github.com/3110Y/profile/internal/application/service"
	"github.com/3110Y/profile/internal/infrastructure/database"
	"github.com/3110Y/profile/internal/infrastructure/repository"
	"github.com/3110Y/profile/pkg/profileGRPC"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var roleRPC *RoleRPC
var repositoryRole *repository.RoleRepository

func init() {
	var err error
	ctx = context.Background()
	err = godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	connect, err = database.NewConnectTest()
	if err != nil {
		log.Fatal(err)
	}
	repositoryRole = repository.NewRoleRepository(connect)
	roleRPC = NewRoleRPC(
		service.NewRoleService(
			repositoryRole,
		),
	)
}

func TestRoleRPC_Item(t *testing.T) {
	id := "7647316e-22e2-4f94-93bc-c0459dcb54de"
	in := profileGRPC.RoleId{
		Id: id,
	}
	roleItem, err := roleRPC.Item(ctx, &in)
	assert.Nil(t, err)
	assert.Equal(t, id, roleItem.Id)
}

func TestRoleRPC_List(t *testing.T) {
	count := uint64(2)
	onPage := uint64(2)
	roleList, err := roleRPC.List(ctx, &profileGRPC.RoleDTOPaginator{OnPage: onPage, Page: 1})
	assert.Nil(t, err)
	assert.Equal(t, onPage, uint64(len(roleList.Data)))
	assert.Equal(t, count, roleList.AllCount)
}

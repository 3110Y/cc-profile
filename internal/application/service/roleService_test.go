package service

import (
	"context"
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/internal/application/mapping"
	"github.com/3110Y/profile/internal/domain/entity"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	ctx = context.Background()
}

func getEntityRole() entity.Role {
	return entity.Role{
		Id:   "12345678",
		Name: "Name",
	}
}

func prepareRole(t *testing.T) (
	func(),
	*MockRoleRepositoryInterface,
	*RoleService,
	entity.Role,
) {
	ctrl := gomock.NewController(t)
	repository := NewMockRoleRepositoryInterface(ctrl)
	role := NewRoleService(
		repository,
	)
	entityRole := getEntityRole()
	return ctrl.Finish, repository, role, entityRole
}

func TestRoleService_Item(t *testing.T) {
	t.Parallel()
	finish, repository, role, entityProfile := prepareRole(t)
	defer finish()
	dtoProfile := utlits.Pointer(mapping.RoleEntityMapping{Entity: entityProfile}).ToRoleDTO()
	repository.EXPECT().Get(ctx, entityProfile.Id).Return(&entityProfile, nil)
	dtoProfileFilled, err := role.Item(ctx, "12345678")
	assert.Nil(t, err)
	assert.Equal(t, dtoProfile, *dtoProfileFilled)
}

func TestRoleService_List(t *testing.T) {
	t.Parallel()
	finish, repository, role, entityProfile := prepareRole(t)
	defer finish()
	entityList := []entity.Role{entityProfile, entityProfile}
	profileListDTO := dto.RoleListDTO{
		Data: utlits.Map(entityList, func(f entity.Role) dto.RoleDTO {
			return utlits.Pointer(mapping.RoleEntityMapping{Entity: f}).ToRoleDTO()
		}),
		AllCount: 3,
	}
	repository.EXPECT().List(ctx, uint64(1), uint64(2)).Return(&entityList, nil)
	repository.EXPECT().Count(ctx).Return(&profileListDTO.AllCount, nil)
	profileListDTOFilled, err := role.List(ctx, 1, 2)
	assert.Nil(t, err)
	assert.Equal(t, profileListDTO, *profileListDTOFilled)
}

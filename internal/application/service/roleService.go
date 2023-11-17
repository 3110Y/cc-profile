package service

//go:generate mockgen -destination mock_roleService_test.go -package service . RoleRepositoryInterface

import (
	"context"
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/internal/application/mapping"
	"github.com/3110Y/profile/internal/domain/entity"
)

type RoleRepositoryInterface interface {
	Get(ctx context.Context, id string) (*entity.Role, error)
	List(ctx context.Context, onPage uint64, page uint64) (*[]entity.Role, error)
	Count(ctx context.Context) (*uint64, error)
}

type RoleService struct {
	repositoryRole RoleRepositoryInterface
}

func NewRoleService(repositoryRole RoleRepositoryInterface) *RoleService {
	return &RoleService{repositoryRole: repositoryRole}
}

func (r *RoleService) Item(ctx context.Context, id string) (dto *dto.RoleDTO, err error) {
	entityRole, err := r.repositoryRole.Get(ctx, id)
	dto = utlits.Pointer(utlits.Pointer(mapping.RoleEntityMapping{Entity: *entityRole}).ToRoleDTO())
	return
}

func (r *RoleService) List(ctx context.Context, onPage uint64, page uint64) (*dto.RoleListDTO, error) {
	roleEntityList, err := r.repositoryRole.List(ctx, onPage, page)
	if err != nil {
		return nil, err
	}
	count, err := r.repositoryRole.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &dto.RoleListDTO{
			Data: utlits.Map(*roleEntityList, func(f entity.Role) dto.RoleDTO {
				return utlits.Pointer(mapping.RoleEntityMapping{Entity: f}).ToRoleDTO()
			}),
			AllCount: *count,
		},
		nil
}

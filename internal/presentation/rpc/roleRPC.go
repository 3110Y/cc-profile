package rpc

import (
	"context"
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/pkg/profileGRPC"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceRoleInterface interface {
	Item(ctx context.Context, id string) (dto *dto.RoleDTO, err error)
	List(ctx context.Context, onPage uint64, page uint64) (*dto.RoleListDTO, error)
}

type RoleRPC struct {
	profileGRPC.UnimplementedRoleServer
	serviceRole ServiceRoleInterface
}

func NewRoleRPC(serviceRole ServiceRoleInterface) *RoleRPC {
	return &RoleRPC{serviceRole: serviceRole}
}

func (r *RoleRPC) Item(ctx context.Context, in *profileGRPC.RoleId) (*profileGRPC.RoleItem, error) {
	roleDTO, err := r.serviceRole.Item(ctx, in.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	out := &profileGRPC.RoleItem{
		Id:   roleDTO.Id,
		Name: roleDTO.Name,
	}
	return out, nil
}
func (r *RoleRPC) List(ctx context.Context, in *profileGRPC.RoleDTOPaginator) (*profileGRPC.RoleDTOListItem, error) {
	roleListDTO, err := r.serviceRole.List(ctx, in.OnPage, in.Page)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	out := &profileGRPC.RoleDTOListItem{
		Data: utlits.Map(roleListDTO.Data, func(roleDTO dto.RoleDTO) *profileGRPC.RoleItem {
			return &profileGRPC.RoleItem{
				Id:   roleDTO.Id,
				Name: roleDTO.Name,
			}
		}),
		AllCount: roleListDTO.AllCount,
	}
	return out, nil
}

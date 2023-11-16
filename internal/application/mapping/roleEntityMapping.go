package mapping

import (
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/internal/domain/entity"
)

type RoleEntityMapping struct {
	Entity entity.Role
}

func (r *RoleEntityMapping) ToRoleDTO() dto.RoleDTO {
	return dto.RoleDTO{
		Id:   r.Entity.Id,
		Name: r.Entity.Name,
	}
}

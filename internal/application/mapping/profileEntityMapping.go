package mapping

import (
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/internal/domain/entity"
)

type ProfileEntityMapping struct {
	Entity entity.Profile
}

func (p *ProfileEntityMapping) ToProfileDTO() dto.ProfileDTO {
	return dto.ProfileDTO{
		Id:         utlits.Pointer(p.Entity.Id),
		Email:      p.Entity.Email,
		Phone:      p.Entity.Phone,
		Surname:    p.Entity.Surname,
		Name:       p.Entity.Name,
		Patronymic: p.Entity.Patronymic,
		Password:   utlits.Pointer(p.Entity.Password),
		CreateAt:   utlits.Pointer(p.Entity.CreateAt),
		UpdateAt:   utlits.Pointer(p.Entity.UpdateAt),
	}
}

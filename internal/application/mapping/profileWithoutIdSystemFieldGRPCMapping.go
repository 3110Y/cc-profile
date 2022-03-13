package mapping

import (
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/pkg/profileGRPC"
)

type ProfileWithoutIdSystemFieldGRPCMapping struct {
	ProfileWithoutIdSystemField profileGRPC.ProfileWithoutIdSystemField
}

func (p *ProfileWithoutIdSystemFieldGRPCMapping) ToProfileDTO() dto.ProfileDTO {
	return dto.ProfileDTO{
		Email:      p.ProfileWithoutIdSystemField.Email,
		Phone:      p.ProfileWithoutIdSystemField.Phone,
		Surname:    p.ProfileWithoutIdSystemField.Surname,
		Name:       p.ProfileWithoutIdSystemField.Name,
		Patronymic: p.ProfileWithoutIdSystemField.Patronymic,
		Password:   &p.ProfileWithoutIdSystemField.Password,
	}
}

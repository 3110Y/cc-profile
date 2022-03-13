package mapping

import (
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/pkg/profileGRPC"
)

type ProfileWithoutSystemFieldGRPCMapping struct {
	ProfileWithoutSystemField profileGRPC.ProfileWithoutSystemField
}

func (p *ProfileWithoutSystemFieldGRPCMapping) ToProfileDTO() dto.ProfileDTO {
	return dto.ProfileDTO{
		Id:         &p.ProfileWithoutSystemField.Id,
		Email:      p.ProfileWithoutSystemField.Email,
		Phone:      p.ProfileWithoutSystemField.Phone,
		Surname:    p.ProfileWithoutSystemField.Surname,
		Name:       p.ProfileWithoutSystemField.Name,
		Patronymic: p.ProfileWithoutSystemField.Patronymic,
		Password:   &p.ProfileWithoutSystemField.Password,
	}
}

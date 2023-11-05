package mapping

import (
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/pkg/profileGRPC"
)

type ProfileWithoutIdSystemFieldPasswordMapping struct {
	ProfileWithoutIdSystemFieldPassword profileGRPC.ProfileWithoutIdSystemFieldPassword
}

func (p *ProfileWithoutIdSystemFieldPasswordMapping) ToProfileDTO() dto.ProfileDTO {
	return dto.ProfileDTO{
		Id:         &p.ProfileWithoutIdSystemFieldPassword.Id,
		Email:      p.ProfileWithoutIdSystemFieldPassword.Email,
		Phone:      p.ProfileWithoutIdSystemFieldPassword.Phone,
		Surname:    p.ProfileWithoutIdSystemFieldPassword.Surname,
		Name:       p.ProfileWithoutIdSystemFieldPassword.Name,
		Patronymic: p.ProfileWithoutIdSystemFieldPassword.Patronymic,
	}
}

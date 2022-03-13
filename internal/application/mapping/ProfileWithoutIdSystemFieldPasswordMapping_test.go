package mapping

import (
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/pkg/profileGRPC"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProfileWithoutIdSystemFieldPasswordMapping_ToProfileDTO(t *testing.T) {
	profileGRPC := profileGRPC.ProfileWithoutIdSystemFieldPassword{
		Id:         "12345678",
		Email:      "test@test.test",
		Phone:      79062579331,
		Surname:    "Surname",
		Name:       "Name",
		Patronymic: "Patronymic",
	}
	profileDTO := dto.ProfileDTO{
		Id:         &profileGRPC.Id,
		Email:      profileGRPC.Email,
		Phone:      profileGRPC.Phone,
		Surname:    profileGRPC.Surname,
		Name:       profileGRPC.Name,
		Patronymic: profileGRPC.Patronymic,
	}
	profileDTOFilled := utlits.Pointer(
		ProfileWithoutIdSystemFieldPasswordMapping{
			ProfileWithoutIdSystemFieldPassword: profileGRPC,
		},
	).ToProfileDTO()
	assert.Equal(t, profileDTO, profileDTOFilled)
}

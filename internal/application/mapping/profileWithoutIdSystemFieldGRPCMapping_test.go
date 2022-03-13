package mapping

import (
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/pkg/profileGRPC"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProfileWithoutIdSystemFieldGRPCMapping_ToProfileDTO(t *testing.T) {
	profileGRPC := profileGRPC.ProfileWithoutIdSystemField{
		Email:      "test@test.test",
		Phone:      79062579331,
		Password:   "Password8",
		Surname:    "Surname",
		Name:       "Name",
		Patronymic: "Patronymic",
	}
	profileDTO := dto.ProfileDTO{
		Email:      profileGRPC.Email,
		Phone:      profileGRPC.Phone,
		Password:   &profileGRPC.Password,
		Surname:    profileGRPC.Surname,
		Name:       profileGRPC.Name,
		Patronymic: profileGRPC.Patronymic,
	}
	profileDTOFilled := utlits.Pointer(
		ProfileWithoutIdSystemFieldGRPCMapping{
			ProfileWithoutIdSystemField: profileGRPC,
		},
	).ToProfileDTO()
	assert.Equal(t, profileDTO, profileDTOFilled)
}

package mapping

import (
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestProfileEntityMapping_ToProfileDTO(t *testing.T) {
	entityProfile := entity.Profile{
		Id:         "123456789",
		Email:      "test@test.test",
		Phone:      791101234567,
		Surname:    "Surname",
		Name:       "Name",
		Patronymic: "Patronymic",
		Password:   "Password8",
		CreateAt:   time.Now().Add(-24 * time.Hour),
		UpdateAt:   time.Now(),
	}
	profileDTO := dto.ProfileDTO{
		Id:         utlits.Pointer(entityProfile.Id),
		Email:      entityProfile.Email,
		Phone:      entityProfile.Phone,
		Surname:    entityProfile.Surname,
		Name:       entityProfile.Name,
		Patronymic: entityProfile.Patronymic,
		Password:   utlits.Pointer(entityProfile.Password),
		CreateAt:   utlits.Pointer(entityProfile.CreateAt),
		UpdateAt:   utlits.Pointer(entityProfile.UpdateAt),
	}
	profileDTOFilled := utlits.Pointer(ProfileEntityMapping{Entity: entityProfile}).ToProfileDTO()
	assert.Equal(t, profileDTO, profileDTOFilled)
}

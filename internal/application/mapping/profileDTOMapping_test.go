package mapping

import (
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/internal/domain/entity"
	"github.com/3110Y/profile/pkg/profileGRPC"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

var profileDTO dto.ProfileDTO

func init() {
	profileDTO = dto.ProfileDTO{
		Id:         utlits.Pointer("12345678"),
		Email:      "test@test.test",
		Phone:      79062579331,
		Surname:    "Surname",
		Name:       "Name",
		Patronymic: "Patronymic",
		Password:   utlits.Pointer("Password8"),
		CreateAt:   utlits.Pointer(time.Now().Add(-72 * time.Hour)),
		UpdateAt:   utlits.Pointer(time.Now()),
	}
}

func TestProfileDTOMapping_ToEntity(t *testing.T) {
	profileEntity := entity.Profile{
		Id:         *profileDTO.Id,
		Email:      profileDTO.Email,
		Phone:      profileDTO.Phone,
		Surname:    profileDTO.Surname,
		Name:       profileDTO.Name,
		Patronymic: profileDTO.Patronymic,
		Password:   *profileDTO.Password,
		CreateAt:   *profileDTO.CreateAt,
		UpdateAt:   *profileDTO.UpdateAt,
	}
	profileEntityFilled := utlits.Pointer(ProfileDTOMapping{ProfileDTO: profileDTO}).ToEntity()
	assert.Equal(t, profileEntity, profileEntityFilled)
}

func TestProfileDTOMapping_ToProfileIdGRPC(t *testing.T) {
	profileId := profileGRPC.ProfileId{
		Id: *profileDTO.Id,
	}
	profileIdFilled := utlits.Pointer(ProfileDTOMapping{ProfileDTO: profileDTO}).ToProfileIdGRPC()
	assert.Equal(t, profileId, profileIdFilled)
}

func TestProfileDTOMapping_ToProfileWithoutPasswordGRPC(t *testing.T) {
	profileWithoutPassword := profileGRPC.ProfileWithoutPassword{
		Id:         *profileDTO.Id,
		Email:      profileDTO.Email,
		Phone:      profileDTO.Phone,
		Surname:    profileDTO.Surname,
		Name:       profileDTO.Name,
		Patronymic: profileDTO.Patronymic,
		CreateAt: &timestamppb.Timestamp{
			Seconds: profileDTO.CreateAt.Unix(),
			Nanos:   int32(profileDTO.CreateAt.Nanosecond()),
		},
		UpdateAt: &timestamppb.Timestamp{
			Seconds: profileDTO.UpdateAt.Unix(),
			Nanos:   int32(profileDTO.UpdateAt.Nanosecond()),
		},
	}
	profileWithoutPasswordFilled := utlits.Pointer(
		ProfileDTOMapping{
			ProfileDTO: profileDTO,
		},
	).ToProfileWithoutPasswordGRPC()
	assert.Equal(t, profileWithoutPassword, profileWithoutPasswordFilled)
}

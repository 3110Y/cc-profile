package mapping

import (
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/pkg/profileGRPC"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestProfileListDTOMapping_ToProfileListGRPC(t *testing.T) {

	profileListDTO := dto.ProfileListDTO{
		Data: []dto.ProfileDTO{
			{
				Id:         utlits.Pointer("12345678"),
				Email:      "test@test.test",
				Phone:      79062579331,
				Surname:    "Surname",
				Name:       "Name",
				Patronymic: "Patronymic",
				Password:   utlits.Pointer("Password8"),
				CreateAt:   utlits.Pointer(time.Now().Add(-72 * time.Hour)),
				UpdateAt:   utlits.Pointer(time.Now()),
			},
			{
				Id:         utlits.Pointer("123456782"),
				Email:      "test@test.test2",
				Phone:      79062579332,
				Surname:    "Surname2",
				Name:       "Name2",
				Patronymic: "Patronymic2",
				Password:   utlits.Pointer("Password82"),
				CreateAt:   utlits.Pointer(time.Now().Add(-72 * time.Hour)),
				UpdateAt:   utlits.Pointer(time.Now()),
			},
		},
		AllCount: 123,
	}
	profileListGRPC := profileGRPC.ProfileList{
		Data: utlits.Map(profileListDTO.Data, func(f dto.ProfileDTO) *profileGRPC.ProfileWithoutPassword {
			return &profileGRPC.ProfileWithoutPassword{
				Id:         *f.Id,
				Email:      f.Email,
				Phone:      f.Phone,
				Surname:    f.Surname,
				Name:       f.Name,
				Patronymic: f.Patronymic,
				CreateAt: &timestamp.Timestamp{
					Seconds: f.CreateAt.Unix(),
					Nanos:   int32(f.CreateAt.Nanosecond()),
				},
				UpdateAt: &timestamp.Timestamp{
					Seconds: f.UpdateAt.Unix(),
					Nanos:   int32(f.UpdateAt.Nanosecond()),
				},
			}
		}),
		AllCount: profileListDTO.AllCount,
	}
	profileListGRPCFilled := utlits.Pointer(ProfileListDTOMapping{profileListDTO}).ToProfileListGRPC()
	assert.True(t, reflect.DeepEqual(profileListGRPC, profileListGRPCFilled))
}

package mapping

import (
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/internal/domain/entity"
	"github.com/3110Y/profile/pkg/profileGRPC"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type ProfileDTOMapping struct {
	ProfileDTO dto.ProfileDTO
}

func (p *ProfileDTOMapping) ToEntity() entity.Profile {
	return entity.Profile{
		Id:         utlits.ValueOrDefault(p.ProfileDTO.Id),
		Email:      p.ProfileDTO.Email,
		Phone:      p.ProfileDTO.Phone,
		Surname:    p.ProfileDTO.Surname,
		Name:       p.ProfileDTO.Name,
		Patronymic: p.ProfileDTO.Patronymic,
		Password:   utlits.ValueOrDefault(p.ProfileDTO.Password),
		CreateAt:   utlits.ValueOrDefault(p.ProfileDTO.CreateAt),
		UpdateAt:   utlits.ValueOrDefault(p.ProfileDTO.UpdateAt),
	}
}

func (p *ProfileDTOMapping) ToProfileIdGRPC() profileGRPC.ProfileId {
	return profileGRPC.ProfileId{
		Id: *p.ProfileDTO.Id,
	}
}

func (p *ProfileDTOMapping) ToProfileWithoutPasswordGRPC() profileGRPC.ProfileWithoutPassword {
	return profileGRPC.ProfileWithoutPassword{
		Id:         utlits.ValueOrDefault(p.ProfileDTO.Id),
		Email:      p.ProfileDTO.Email,
		Phone:      p.ProfileDTO.Phone,
		Surname:    p.ProfileDTO.Surname,
		Name:       p.ProfileDTO.Name,
		Patronymic: p.ProfileDTO.Patronymic,
		CreateAt: utlits.NullSafeFunction(p.ProfileDTO.CreateAt, func(t *time.Time) *timestamppb.Timestamp {
			return &timestamppb.Timestamp{
				Seconds: p.ProfileDTO.CreateAt.Unix(),
				Nanos:   int32(p.ProfileDTO.CreateAt.Nanosecond()),
			}
		}),
		UpdateAt: utlits.NullSafeFunction(p.ProfileDTO.UpdateAt, func(t *time.Time) *timestamppb.Timestamp {
			return &timestamppb.Timestamp{
				Seconds: p.ProfileDTO.UpdateAt.Unix(),
				Nanos:   int32(p.ProfileDTO.UpdateAt.Nanosecond()),
			}
		}),
	}
}

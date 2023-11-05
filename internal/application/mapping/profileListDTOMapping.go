package mapping

import (
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/pkg/profileGRPC"
)

type ProfileListDTOMapping struct {
	ProfileListDTO dto.ProfileListDTO
}

func (p *ProfileListDTOMapping) ToProfileListGRPC() profileGRPC.ProfileList {
	return profileGRPC.ProfileList{
		Data: utlits.Map(p.ProfileListDTO.Data, func(profileDTO dto.ProfileDTO) *profileGRPC.ProfileWithoutPassword {
			return utlits.Pointer(utlits.Pointer(ProfileDTOMapping{profileDTO}).ToProfileWithoutPasswordGRPC())
		}),
		AllCount: p.ProfileListDTO.AllCount,
	}
}

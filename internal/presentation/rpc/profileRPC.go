package rpc

import (
	"context"
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/internal/application/mapping"
	"github.com/3110Y/profile/pkg/profileGRPC"
)

type ServiceProfileInterface interface {
	Add(ctx context.Context, profileDTO dto.ProfileDTO) (id string, err error)
	Edit(ctx context.Context, profileDTO dto.ProfileDTO) (uint64, error)
	Item(ctx context.Context, id string) (profileDto dto.ProfileDTO, err error)
	Delete(ctx context.Context, id string) (uint64, error)
	List(ctx context.Context, onPage uint64, page uint64) (dto.ProfileListDTO, error)
}

type ValidatorProfileInterface interface {
	ValidEmail(email string) (err error)
	ValidPassword(passwordHash string) (err error)
	ValidPhone(phone uint64) (err error)
}

type ProfileRPC struct {
	profileGRPC.UnimplementedProfileServiceServer
	serviceProfile   ServiceProfileInterface
	validatorProfile ValidatorProfileInterface
}

func NewProfileRPC(serviceProfile ServiceProfileInterface, validatorProfile ValidatorProfileInterface) *ProfileRPC {
	return &ProfileRPC{serviceProfile: serviceProfile, validatorProfile: validatorProfile}
}

func (p *ProfileRPC) validateProfileDTO(profileDto dto.ProfileDTO) error {
	err := p.validatorProfile.ValidEmail(profileDto.Email)
	if err != nil {
		return err
	}
	err = p.validatorProfile.ValidPassword(*profileDto.Password)
	if err != nil {
		return err
	}
	err = p.validatorProfile.ValidPhone(profileDto.Phone)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProfileRPC) Add(
	ctx context.Context,
	in *profileGRPC.ProfileWithoutIdSystemField,
) (*profileGRPC.ProfileId, error) {
	profileDto := utlits.Pointer(
		mapping.ProfileWithoutIdSystemFieldGRPCMapping{
			ProfileWithoutIdSystemField: *in,
		},
	).ToProfileDTO()
	err := p.validateProfileDTO(profileDto)
	if err != nil {
		return nil, err
	}
	id, err := p.serviceProfile.Add(ctx, profileDto)
	return &profileGRPC.ProfileId{Id: id}, err
}

func (p *ProfileRPC) Item(ctx context.Context, in *profileGRPC.ProfileId) (*profileGRPC.ProfileWithoutPassword, error) {
	profileWithoutPassword, err := p.serviceProfile.Item(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	out := utlits.Pointer(mapping.ProfileDTOMapping{ProfileDTO: profileWithoutPassword}).ToProfileWithoutPasswordGRPC()
	return &out, nil
}

func (p *ProfileRPC) Delete(ctx context.Context, in *profileGRPC.ProfileId) (*profileGRPC.EmptyResponse, error) {
	_, err := p.serviceProfile.Delete(ctx, in.Id)
	return &profileGRPC.EmptyResponse{}, err
}

func (p *ProfileRPC) Edit(
	ctx context.Context,
	in *profileGRPC.ProfileWithoutSystemField,
) (*profileGRPC.EmptyResponse, error) {
	profileDto := utlits.Pointer(
		mapping.ProfileWithoutSystemFieldGRPCMapping{
			ProfileWithoutSystemField: *in,
		},
	).ToProfileDTO()
	err := p.validateProfileDTO(profileDto)
	if err != nil {
		return nil, err
	}
	_, err = p.serviceProfile.Edit(ctx, profileDto)
	return &profileGRPC.EmptyResponse{}, err
}

func (p *ProfileRPC) List(ctx context.Context, in *profileGRPC.ProfilePaginator) (*profileGRPC.ProfileList, error) {
	profileListDTO, err := p.serviceProfile.List(ctx, in.OnPage, in.Page)
	if err != nil {
		return nil, err
	}
	profileList := utlits.Pointer(mapping.ProfileListDTOMapping{ProfileListDTO: profileListDTO}).ToProfileListGRPC()
	return &profileList, nil
}

package rpc

import (
	"context"
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/internal/application/mapping"
	"github.com/3110Y/profile/pkg/profileGRPC"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceProfileInterface interface {
	Add(ctx context.Context, profileDTO dto.ProfileDTO) (id *string, err error)
	Edit(ctx context.Context, profileDTO dto.ProfileDTO) (*uint64, error)
	Item(ctx context.Context, id string) (profileDto *dto.ProfileDTO, err error)
	Delete(ctx context.Context, id string) (*uint64, error)
	List(ctx context.Context, onPage uint64, page uint64) (*dto.ProfileListDTO, error)
	EditWithoutPassword(ctx context.Context, profileDTO dto.ProfileDTO) (*uint64, error)
	ChangePassword(ctx context.Context, profileDTO dto.ProfileDTO) (*uint64, error)
	GetByEmailOrPhone(
		ctx context.Context,
		email string,
		phone uint64,
		password string,
	) (profileDto *dto.ProfileDTO, err error)
}

type ValidatorProfileInterface interface {
	ValidEmail(email string) (err error)
	ValidPassword(passwordHash string) (err error)
	ValidPhone(phone uint64) (err error)
}

type ProfileRPC struct {
	profileGRPC.UnimplementedProfileServer
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
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}
	id, err := p.serviceProfile.Add(ctx, profileDto)
	return &profileGRPC.ProfileId{Id: *id}, err
}

func (p *ProfileRPC) Item(ctx context.Context, in *profileGRPC.ProfileId) (*profileGRPC.ProfileWithoutPassword, error) {
	profileWithoutPassword, err := p.serviceProfile.Item(ctx, in.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	out := utlits.Pointer(mapping.ProfileDTOMapping{ProfileDTO: *profileWithoutPassword}).ToProfileWithoutPasswordGRPC()
	return &out, nil
}

func (p *ProfileRPC) Delete(ctx context.Context, in *profileGRPC.ProfileId) (*profileGRPC.EmptyResponse, error) {
	_, err := p.serviceProfile.Delete(ctx, in.Id)
	if err != nil {
		return &profileGRPC.EmptyResponse{}, status.Errorf(codes.Internal, err.Error())
	}
	return &profileGRPC.EmptyResponse{}, nil
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
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}
	_, err = p.serviceProfile.Edit(ctx, profileDto)
	return &profileGRPC.EmptyResponse{}, err
}

func (p *ProfileRPC) List(ctx context.Context, in *profileGRPC.ProfilePaginator) (*profileGRPC.ProfileList, error) {
	profileListDTO, err := p.serviceProfile.List(ctx, in.OnPage, in.Page)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	profileList := utlits.Pointer(mapping.ProfileListDTOMapping{ProfileListDTO: *profileListDTO}).ToProfileListGRPC()
	return &profileList, nil
}

func (p *ProfileRPC) EditWithoutPassword(
	ctx context.Context,
	in *profileGRPC.ProfileWithoutIdSystemFieldPassword,
) (*profileGRPC.EmptyResponse, error) {
	profileDto := utlits.Pointer(
		mapping.ProfileWithoutIdSystemFieldPasswordMapping{
			ProfileWithoutIdSystemFieldPassword: *in,
		},
	).ToProfileDTO()
	err := p.validatorProfile.ValidEmail(profileDto.Email)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}
	err = p.validatorProfile.ValidPhone(profileDto.Phone)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}
	_, err = p.serviceProfile.EditWithoutPassword(ctx, profileDto)
	return &profileGRPC.EmptyResponse{}, err
}

func (p *ProfileRPC) ChangePassword(
	ctx context.Context,
	in *profileGRPC.ProfilePassword,
) (*profileGRPC.EmptyResponse, error) {
	profileDto := dto.ProfileDTO{
		Id:       &in.Id,
		Password: &in.Password,
	}
	err := p.validatorProfile.ValidPassword(*profileDto.Password)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}
	_, err = p.serviceProfile.EditWithoutPassword(ctx, profileDto)
	return &profileGRPC.EmptyResponse{}, err
}

func (p *ProfileRPC) GetByEmailOrPhone(
	ctx context.Context,
	in *profileGRPC.ProfileEmailPhonePassword,
) (*profileGRPC.ProfileWithoutPassword, error) {
	profileWithoutPassword, err := p.serviceProfile.GetByEmailOrPhone(ctx, in.Email, in.Phone, in.Password)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	out := utlits.Pointer(mapping.ProfileDTOMapping{ProfileDTO: *profileWithoutPassword}).ToProfileWithoutPasswordGRPC()
	return &out, nil
}

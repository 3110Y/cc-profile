package service

//go:generate mockgen -destination mock_profileService_test.go -package service . ProfileRepositoryInterface,PasswordServiceInterface

import (
	"context"
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/internal/application/mapping"
	"github.com/3110Y/profile/internal/domain/entity"
	"github.com/google/uuid"
)

type ProfileRepositoryInterface interface {
	Add(ctx context.Context, profile entity.Profile) (uint64, error)
	Edit(ctx context.Context, profile entity.Profile) (uint64, error)
	Get(ctx context.Context, id string) (profile entity.Profile, err error)
	List(ctx context.Context, onPage uint64, page uint64) (profileList []entity.Profile, err error)
	Delete(ctx context.Context, id string) (uint64, error)
	Count(ctx context.Context) (uint64, error)
}

type PasswordServiceInterface interface {
	Encode(password string) (passwordHash string, err error)
}

type ProfileService struct {
	repositoryProfile ProfileRepositoryInterface
	servicePassword   PasswordServiceInterface
}

func NewProfileService(
	repositoryProfile ProfileRepositoryInterface,
	servicePassword PasswordServiceInterface,
) *ProfileService {
	return &ProfileService{
		repositoryProfile: repositoryProfile,
		servicePassword:   servicePassword,
	}
}

func (p *ProfileService) Add(ctx context.Context, profileDTO dto.ProfileDTO) (id string, err error) {
	entityProfile := utlits.Pointer(mapping.ProfileDTOMapping{ProfileDTO: profileDTO}).ToEntity()
	passwordHash, err := p.servicePassword.Encode(entityProfile.Password)
	if err != nil {
		return id, err
	}
	entityProfile.Password = passwordHash
	id = uuid.New().String()
	entityProfile.Id = id
	_, err = p.repositoryProfile.Add(ctx, entityProfile)
	return
}

func (p *ProfileService) Edit(ctx context.Context, profileDTO dto.ProfileDTO) (uint64, error) {
	entityProfile := utlits.Pointer(mapping.ProfileDTOMapping{ProfileDTO: profileDTO}).ToEntity()
	return p.repositoryProfile.Edit(ctx, entityProfile)
}

func (p *ProfileService) Item(ctx context.Context, id string) (profileDto dto.ProfileDTO, err error) {
	entityProfile, err := p.repositoryProfile.Get(ctx, id)
	profileDto = utlits.Pointer(mapping.ProfileEntityMapping{Entity: entityProfile}).ToProfileDTO()
	return
}

func (p *ProfileService) Delete(ctx context.Context, id string) (uint64, error) {
	return p.repositoryProfile.Delete(ctx, id)
}

func (p *ProfileService) List(ctx context.Context, onPage uint64, page uint64) (dto.ProfileListDTO, error) {
	profileEntityList, err := p.repositoryProfile.List(ctx, onPage, page)
	if err != nil {
		return dto.ProfileListDTO{}, err
	}
	count, err := p.repositoryProfile.Count(ctx)
	if err != nil {
		return dto.ProfileListDTO{}, err
	}
	return dto.ProfileListDTO{
			Data: utlits.Map(profileEntityList, func(f entity.Profile) dto.ProfileDTO {
				return utlits.Pointer(mapping.ProfileEntityMapping{Entity: f}).ToProfileDTO()
			}),
			AllCount: count,
		},
		nil
}

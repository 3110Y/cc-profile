package service

import (
	"context"
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/internal/application/mapping"
	"github.com/3110Y/profile/internal/domain/entity"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

func getEntityProfile() entity.Profile {
	return entity.Profile{
		Id:         "12345678",
		Email:      "test@test.test",
		Phone:      79062579331,
		Password:   "Password8",
		Surname:    "Surname",
		Name:       "Name",
		Patronymic: "Patronymic",
		CreateAt:   time.Now(),
		UpdateAt:   time.Now().Add(-72 * time.Hour),
	}
}

func prepare(t *testing.T) (
	func(),
	*MockProfileRepositoryInterface,
	*MockPasswordServiceInterface,
	*ProfileService,
	entity.Profile,
) {
	ctrl := gomock.NewController(t)
	repository := NewMockProfileRepositoryInterface(ctrl)
	passwordService := NewMockPasswordServiceInterface(ctrl)
	profile := NewProfileService(
		repository,
		passwordService,
	)
	entityProfile := getEntityProfile()
	return ctrl.Finish, repository, passwordService, profile, entityProfile
}

func TestProfileService_Add(t *testing.T) {
	t.Parallel()
	finish, repository, passwordService, profile, entityProfile := prepare(t)
	defer finish()
	dtoProfile := utlits.Pointer(mapping.ProfileEntityMapping{Entity: entityProfile}).ToProfileDTO()
	passwordService.EXPECT().Encode(entityProfile.Password).Return(entityProfile.Password, nil)
	repository.EXPECT().Add(ctx, gomock.Any()).Return(uint64(1), nil)
	id, err := profile.Add(ctx, dtoProfile)
	assert.Nil(t, err)
	assert.Greater(t, len(id), 8)
}

func TestProfileService_Edit(t *testing.T) {
	t.Parallel()
	finish, repository, _, profile, entityProfile := prepare(t)
	defer finish()
	dtoProfile := utlits.Pointer(mapping.ProfileEntityMapping{Entity: entityProfile}).ToProfileDTO()
	repository.EXPECT().Edit(ctx, entityProfile).Return(uint64(1), nil)
	rowsAffected, err := profile.Edit(ctx, dtoProfile)
	assert.Nil(t, err)
	assert.Equal(t, rowsAffected, uint64(1))
}

func TestProfileService_Item(t *testing.T) {
	t.Parallel()
	finish, repository, _, profile, entityProfile := prepare(t)
	defer finish()
	dtoProfile := utlits.Pointer(mapping.ProfileEntityMapping{Entity: entityProfile}).ToProfileDTO()
	repository.EXPECT().Get(ctx, entityProfile.Id).Return(entityProfile, nil)
	dtoProfileFilled, err := profile.Item(ctx, "12345678")
	assert.Nil(t, err)
	assert.Equal(t, dtoProfile, dtoProfileFilled)
}

func TestProfileService_Delete(t *testing.T) {
	t.Parallel()
	finish, repository, _, profile, entityProfile := prepare(t)
	defer finish()
	repository.EXPECT().Delete(ctx, entityProfile.Id).Return(uint64(1), nil)
	rowsAffected, err := profile.Delete(ctx, entityProfile.Id)
	assert.Nil(t, err)
	assert.Equal(t, rowsAffected, uint64(1))
}

func TestProfileService_List(t *testing.T) {
	t.Parallel()
	finish, repository, _, profile, entityProfile := prepare(t)
	defer finish()
	entityList := []entity.Profile{entityProfile, entityProfile}
	profileListDTO := dto.ProfileListDTO{
		Data: utlits.Map(entityList, func(f entity.Profile) dto.ProfileDTO {
			return utlits.Pointer(mapping.ProfileEntityMapping{Entity: f}).ToProfileDTO()
		}),
		AllCount: 3,
	}
	repository.EXPECT().List(ctx, uint64(1), uint64(2)).Return(entityList, nil)
	repository.EXPECT().Count(ctx).Return(profileListDTO.AllCount, nil)
	profileListDTOFilled, err := profile.List(ctx, 1, 2)
	assert.Nil(t, err)
	assert.Equal(t, profileListDTO, profileListDTOFilled)
}

func TestProfileService_EditWithoutPassword(t *testing.T) {
	t.Parallel()
	finish, repository, _, profile, entityProfile := prepare(t)
	defer finish()
	dtoProfile := utlits.Pointer(mapping.ProfileEntityMapping{Entity: entityProfile}).ToProfileDTO()
	repository.EXPECT().EditWithoutPassword(ctx, entityProfile).Return(uint64(1), nil)
	rowsAffected, err := profile.EditWithoutPassword(ctx, dtoProfile)
	assert.Nil(t, err)
	assert.Equal(t, rowsAffected, uint64(1))
}

func TestProfileService_ChangePassword(t *testing.T) {
	t.Parallel()
	finish, repository, _, profile, entityProfile := prepare(t)
	defer finish()
	dtoProfile := utlits.Pointer(mapping.ProfileEntityMapping{Entity: entityProfile}).ToProfileDTO()
	repository.EXPECT().ChangePassword(ctx, entityProfile).Return(uint64(1), nil)
	rowsAffected, err := profile.ChangePassword(ctx, dtoProfile)
	assert.Nil(t, err)
	assert.Equal(t, rowsAffected, uint64(1))
}

func TestProfileService_GetByEmailOrPhone(t *testing.T) {
	t.Parallel()
	finish, repository, passwordService, profile, entityProfile := prepare(t)
	defer finish()
	dtoProfile := utlits.Pointer(mapping.ProfileEntityMapping{Entity: entityProfile}).ToProfileDTO()
	repository.EXPECT().GetByEmailOrPhone(ctx, entityProfile.Email, entityProfile.Phone).Return(entityProfile, nil)
	passwordService.EXPECT().Encode(entityProfile.Password).Return(entityProfile.Password, nil)
	dtoProfileFilled, err := profile.GetByEmailOrPhone(
		ctx,
		entityProfile.Email,
		entityProfile.Phone,
		entityProfile.Password,
	)
	assert.Nil(t, err)
	assert.Equal(t, dtoProfile, dtoProfileFilled)
}

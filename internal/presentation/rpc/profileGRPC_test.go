package rpc

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/profile/internal/application/mapping"
	"github.com/3110Y/profile/internal/application/service"
	"github.com/3110Y/profile/internal/application/validator"
	"github.com/3110Y/profile/internal/domain/entity"
	"github.com/3110Y/profile/internal/infrastructure/database"
	"github.com/3110Y/profile/internal/infrastructure/repository"
	"github.com/3110Y/profile/pkg/profileGRPC"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var ctx context.Context
var connect *sqlx.DB
var profileRPC *ProfileRPC
var repositoryProfile *repository.ProfileRepository

func init() {
	var err error
	ctx = context.Background()
	err = godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	connect, err = database.NewConnectTest()
	if err != nil {
		log.Fatal(err)
	}
	repositoryProfile = repository.NewProfileRepository(connect)
	profileRPC = NewProfileRPC(
		service.NewProfileService(
			repositoryProfile,
			service.NewPasswordService(),
		),
		validator.ProfileValidator{},
	)
}

func add(t *testing.T) entity.Profile {
	profileEntity := entity.Profile{
		Id:         uuid.New().String(),
		Email:      "Email",
		Phone:      79062579331,
		Password:   "Password8",
		Surname:    "Surname",
		Name:       "Name",
		Patronymic: "Patronymic",
	}
	passwordSHA := sha512.Sum512([]byte(profileEntity.Password))
	profileEntity.Password = hex.EncodeToString(passwordSHA[:])
	_, err := repositoryProfile.Add(ctx, profileEntity)
	assert.Nil(t, err)
	return profileEntity
}

func TestProfile_Add(t *testing.T) {
	defer database.Clean(t, "profile", connect)
	in := profileGRPC.ProfileWithoutIdSystemField{
		Email:      "test@test.test",
		Phone:      79062579331,
		Password:   "Password8",
		Surname:    "Surname",
		Name:       "Name",
		Patronymic: "Patronymic",
	}
	add, err := profileRPC.Add(ctx, &in)
	assert.Nil(t, err)
	assert.Equal(t, len(add.Id), 36)
	profileItem := entity.Profile{}
	err = database.GetById(&profileItem, "profile", add.Id, connect)
	assert.Nil(t, err)
	//goland:noinspection GoVetCopyLock
	expected := utlits.Pointer(
		mapping.ProfileWithoutIdSystemFieldGRPCMapping{
			ProfileWithoutIdSystemField: in,
		},
	).ToProfileDTO()
	actual := utlits.Pointer(mapping.ProfileEntityMapping{Entity: profileItem}).ToProfileDTO()
	expected.Id = actual.Id
	expected.CreateAt = actual.CreateAt
	expected.UpdateAt = actual.UpdateAt
	expected.Password = actual.Password
	assert.Equal(t, expected, actual)
}

func TestProfileRPC_Item(t *testing.T) {
	defer database.Clean(t, "profile", connect)
	profileEntity := add(t)
	in := profileGRPC.ProfileId{
		Id: profileEntity.Id,
	}
	profileWithoutPassword, err := profileRPC.Item(ctx, &in)
	assert.Nil(t, err)
	assert.Equal(t, profileEntity.Id, profileWithoutPassword.Id)
}

func TestProfileRPC_Delete(t *testing.T) {
	defer database.Clean(t, "profile", connect)
	profileEntity := add(t)
	_, err := profileRPC.Delete(ctx, &profileGRPC.ProfileId{Id: profileEntity.Id})
	assert.Nil(t, err)
	var profileList []entity.Profile
	err = database.Select(&profileList, "profile", 2, 1, connect)
	assert.Nil(t, err)
	assert.Len(t, profileList, 0)
}

func TestProfileRPC_Edit(t *testing.T) {
	defer database.Clean(t, "profile", connect)
	profileEntity := add(t)
	profileWithoutSystemField := profileGRPC.ProfileWithoutSystemField{
		Id:         profileEntity.Id,
		Email:      "test@test.test",
		Phone:      79062579332,
		Password:   "Password82",
		Surname:    "Surname2",
		Name:       "Name2",
		Patronymic: "Patronymic2",
	}
	_, err := profileRPC.Edit(ctx, &profileWithoutSystemField)
	assert.Nil(t, err)
	err = database.GetById(&profileEntity, "profile", profileWithoutSystemField.Id, connect)
	assert.Nil(t, err)
	//goland:noinspection GoVetCopyLock
	expected := utlits.Pointer(
		mapping.ProfileWithoutSystemFieldGRPCMapping{
			ProfileWithoutSystemField: profileWithoutSystemField,
		},
	).ToProfileDTO()
	actual := utlits.Pointer(mapping.ProfileEntityMapping{Entity: profileEntity}).ToProfileDTO()
	assert.Equal(t, *expected.Id, *actual.Id)
	assert.Equal(t, *expected.Password, *actual.Password)
	expected.Id = actual.Id
	expected.CreateAt = actual.CreateAt
	expected.UpdateAt = actual.UpdateAt
	expected.Password = actual.Password
	assert.Equal(t, expected, actual)
}

func TestProfileRPC_List(t *testing.T) {
	defer database.Clean(t, "profile", connect)
	count := uint64(3)
	onPage := uint64(2)
	for i := uint64(0); i < count; i++ {
		add(t)
	}
	profileList, err := profileRPC.List(ctx, &profileGRPC.ProfilePaginator{OnPage: onPage, Page: 1})
	assert.Nil(t, err)
	assert.Equal(t, onPage, uint64(len(profileList.Data)))
	assert.Equal(t, count, profileList.AllCount)
}

func TestProfileRPC_EditWithoutPassword(t *testing.T) {
	defer database.Clean(t, "profile", connect)
	profileEntity := add(t)
	ProfileWithoutIdSystemFieldPassword := profileGRPC.ProfileWithoutIdSystemFieldPassword{
		Id:         profileEntity.Id,
		Email:      "test@test.test",
		Phone:      79062579332,
		Surname:    "Surname2",
		Name:       "Name2",
		Patronymic: "Patronymic2",
	}
	_, err := profileRPC.EditWithoutPassword(ctx, &ProfileWithoutIdSystemFieldPassword)
	assert.Nil(t, err)
	err = database.GetById(&profileEntity, "profile", ProfileWithoutIdSystemFieldPassword.Id, connect)
	assert.Nil(t, err)
	//goland:noinspection GoVetCopyLock
	expected := utlits.Pointer(
		mapping.ProfileWithoutIdSystemFieldPasswordMapping{
			ProfileWithoutIdSystemFieldPassword: ProfileWithoutIdSystemFieldPassword,
		},
	).ToProfileDTO()
	actual := utlits.Pointer(mapping.ProfileEntityMapping{Entity: profileEntity}).ToProfileDTO()
	assert.Equal(t, *expected.Id, *actual.Id)
	assert.NotEqual(t, expected.Password, actual.Password)
	expected.Id = actual.Id
	expected.CreateAt = actual.CreateAt
	expected.UpdateAt = actual.UpdateAt
	expected.Password = actual.Password
	assert.Equal(t, expected, actual)
}

func TestProfileRPC_ChangePassword(t *testing.T) {
	defer database.Clean(t, "profile", connect)
	profileEntity := add(t)
	ProfilePassword := profileGRPC.ProfilePassword{
		Id:       profileEntity.Id,
		Password: profileEntity.Password,
	}
	_, err := profileRPC.ChangePassword(ctx, &ProfilePassword)
	assert.Nil(t, err)
	err = database.GetById(&profileEntity, "profile", ProfilePassword.Id, connect)
	assert.Nil(t, err)
	expected := profileGRPC.ProfilePassword{
		Id:       ProfilePassword.Id,
		Password: ProfilePassword.Password,
	}
	actual := utlits.Pointer(mapping.ProfileEntityMapping{Entity: profileEntity}).ToProfileDTO()
	assert.Equal(t, expected.Id, *actual.Id)
	assert.Equal(t, expected.Password, *actual.Password)
}

func TestProfileRPC_GetByEmailOrPhone(t *testing.T) {
	defer database.Clean(t, "profile", connect)
	profileEntity := add(t)
	in := profileGRPC.ProfileEmailPhonePassword{
		Email:    profileEntity.Email,
		Phone:    profileEntity.Phone,
		Password: "Password8",
	}
	add, err := profileRPC.GetByEmailOrPhone(ctx, &in)
	assert.Nil(t, err)
	assert.Equal(t, profileEntity.Id, add.Id)

	in = profileGRPC.ProfileEmailPhonePassword{
		Email:    profileEntity.Email,
		Phone:    profileEntity.Phone,
		Password: profileEntity.Password,
	}
	add, err = profileRPC.GetByEmailOrPhone(ctx, &in)
	assert.NotNil(t, err)
	assert.Error(t, err, "password mismatch")
}

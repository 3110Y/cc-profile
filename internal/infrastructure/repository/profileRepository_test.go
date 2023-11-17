package repository

import (
	"context"
	"github.com/3110Y/profile/internal/domain/entity"
	"github.com/3110Y/profile/internal/infrastructure/database"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"testing"
)

var profileRepository ProfileRepository
var ctxRepository context.Context
var connect *sqlx.DB

func getEntityProfile() entity.Profile {
	return entity.Profile{
		Id:         uuid.New().String(),
		Email:      "Email",
		Phone:      79062579331,
		Password:   "Password8",
		Surname:    "Surname",
		Name:       "Name",
		Patronymic: "Patronymic",
	}
}

func cleanAtInProfile(profile, profileFilled *entity.Profile) {
	profile.CreateAt = profileFilled.CreateAt
	profile.UpdateAt = profileFilled.UpdateAt
}

func init() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	connect, err = database.NewConnectTest()
	if err != nil {
		log.Fatal(err)
	}
	profileRepository = ProfileRepository{db: connect}
	ctxRepository = context.Background()
}

func TestProfileRepository_Add(t *testing.T) {
	defer database.Clean(t, "profile", connect)
	profileEntity := getEntityProfile()
	rowsAffected, err := profileRepository.Add(ctxRepository, profileEntity)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), *rowsAffected)
	var profileList []entity.Profile
	err = database.SelectAll(&profileList, "profile", connect)
	assert.Nil(t, err)
	assert.Len(t, profileList, 1)
	cleanAtInProfile(&profileList[0], &profileEntity)
	assert.Equal(t, profileEntity, profileList[0])
}

func TestProfileRepository_Add_Fail(t *testing.T) {
	defer database.Clean(t, "profile", connect)
	profileEntity := getEntityProfile()
	rowsAffected, err := profileRepository.Add(ctxRepository, profileEntity)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), *rowsAffected)
	rowsAffected, err = profileRepository.Add(ctxRepository, profileEntity)
	assert.Nil(t, rowsAffected)
	assert.NotNil(t, err)
}

func TestProfileRepository_Get(t *testing.T) {
	defer database.Clean(t, "profile", connect)
	profileFilled := getEntityProfile()
	assert.NotNil(t, profileFilled)
	rowsAffected, err := profileRepository.Add(ctxRepository, profileFilled)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), *rowsAffected)
	profileItem, err := profileRepository.Get(ctxRepository, profileFilled.Id)
	assert.Nil(t, err)
	cleanAtInProfile(profileItem, &profileFilled)
	assert.Equal(t, profileFilled, *profileItem)
}

func TestProfileRepository_List(t *testing.T) {
	defer database.Clean(t, "profile", connect)
	profileFilled := getEntityProfile()
	rowsAffected, err := profileRepository.Add(ctxRepository, profileFilled)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), *rowsAffected)
	profileFilled = getEntityProfile()
	profileFilled.Email = profileFilled.Email + "two"
	profileFilled.Phone = 79062579332
	rowsAffected, err = profileRepository.Add(ctxRepository, profileFilled)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), *rowsAffected)
	profileList, err := profileRepository.List(ctxRepository, 2, 1)
	assert.Nil(t, err)
	assert.Len(t, *profileList, 2)
	profileList, err = profileRepository.List(ctxRepository, 1, 1)
	assert.Nil(t, err)
	assert.Len(t, *profileList, 1)
	profileList, err = profileRepository.List(ctxRepository, 1, 2)
	assert.Nil(t, err)
	assert.Len(t, *profileList, 1)
	profileList, err = profileRepository.List(ctxRepository, 10, 1)
	assert.Nil(t, err)
	assert.Len(t, *profileList, 2)
	profileList, err = profileRepository.List(ctxRepository, 10, 2)
	assert.Nil(t, err)
	assert.Len(t, *profileList, 0)
}

func TestProfileRepository_Delete(t *testing.T) {
	defer database.Clean(t, "profile", connect)
	profileFilled := getEntityProfile()
	rowsAffected, err := profileRepository.Add(ctxRepository, profileFilled)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), *rowsAffected)
	profileFilledTwo := getEntityProfile()
	profileFilledTwo.Email = profileFilledTwo.Email + "two"
	profileFilledTwo.Phone = 79062579332
	rowsAffected, err = profileRepository.Add(ctxRepository, profileFilledTwo)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), *rowsAffected)
	var profileList []entity.Profile
	err = database.Select(&profileList, "profile", 2, 1, connect)
	assert.Nil(t, err)
	assert.Len(t, profileList, 2)
	rowsAffected, err = profileRepository.Delete(ctxRepository, profileFilled.Id)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), *rowsAffected)
	profileList = []entity.Profile{}
	err = database.Select(&profileList, "profile", 2, 1, connect)
	assert.Nil(t, err)
	assert.Len(t, profileList, 1)
	rowsAffected, err = profileRepository.Delete(ctxRepository, profileFilledTwo.Id)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), *rowsAffected)
	profileList = []entity.Profile{}
	err = database.Select(&profileList, "profile", 2, 1, connect)
	assert.Nil(t, err)
	assert.Len(t, profileList, 0)
}

func TestProfileRepository_Edit(t *testing.T) {
	defer database.Clean(t, "profile", connect)
	profileFilled := getEntityProfile()
	rowsAffected, err := profileRepository.Add(ctxRepository, profileFilled)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), *rowsAffected)
	profileEntity := entity.Profile{
		Id:         profileFilled.Id,
		Email:      "Email2",
		Phone:      79062579332,
		Password:   "Password82",
		Surname:    "Surname2",
		Name:       "Name2",
		Patronymic: "Patronymic2",
	}
	rowsAffected, err = profileRepository.Edit(ctxRepository, profileEntity)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), *rowsAffected)
	profileEntityFilled := entity.Profile{}
	err = database.GetById(&profileEntityFilled, "profile", profileEntity.Id, connect)
	assert.Nil(t, err)
	assert.Equal(t, &profileEntity.Id, &profileEntityFilled.Id)
	profileEntity.Id = profileEntityFilled.Id
	profileEntity.CreateAt = profileEntityFilled.CreateAt
	profileEntity.UpdateAt = profileEntityFilled.UpdateAt
	assert.Equal(t, profileEntity, profileEntityFilled)
}

func TestProfileRepository_Count(t *testing.T) {
	defer database.Clean(t, "profile", connect)
	for i, iMax := 0, 2; i < iMax; i++ {
		profileFilled := getEntityProfile()
		iString := strconv.Itoa(i)
		profileFilled.Email = profileFilled.Email + iString
		profileFilled.Phone = profileFilled.Phone + uint64(i)
		rowsAffected, err := profileRepository.Add(ctxRepository, profileFilled)
		assert.Nil(t, err)
		assert.Equal(t, uint64(1), *rowsAffected)
	}
	count, err := profileRepository.Count(ctxRepository)
	assert.Nil(t, err)
	assert.Equal(t, uint64(2), *count)
}

func TestProfileRepository_EditWithoutPassword(t *testing.T) {
	defer database.Clean(t, "profile", connect)
	profileFilled := getEntityProfile()
	rowsAffected, err := profileRepository.Add(ctxRepository, profileFilled)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), *rowsAffected)
	profileEntity := entity.Profile{
		Id:         profileFilled.Id,
		Email:      "Email2",
		Phone:      79062579332,
		Surname:    "Surname2",
		Name:       "Name2",
		Patronymic: "Patronymic2",
	}
	rowsAffected, err = profileRepository.EditWithoutPassword(ctxRepository, profileEntity)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), *rowsAffected)
	profileEntityFilled := entity.Profile{}
	err = database.GetById(&profileEntityFilled, "profile", profileEntity.Id, connect)
	assert.Nil(t, err)
	assert.Equal(t, &profileEntity.Id, &profileEntityFilled.Id)
	assert.NotEqual(t, &profileEntity.Password, &profileEntityFilled.Password)
	profileEntity.Id = profileEntityFilled.Id
	profileEntity.CreateAt = profileEntityFilled.CreateAt
	profileEntity.UpdateAt = profileEntityFilled.UpdateAt
	profileEntity.Password = profileEntityFilled.Password
	assert.Equal(t, profileEntity, profileEntityFilled)
}

func TestProfileRepository_ChangePassword(t *testing.T) {
	defer database.Clean(t, "profile", connect)
	profileFilled := getEntityProfile()
	rowsAffected, err := profileRepository.Add(ctxRepository, profileFilled)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), *rowsAffected)
	profileEntity := entity.Profile{
		Id:       profileFilled.Id,
		Password: "Password82",
	}
	rowsAffected, err = profileRepository.ChangePassword(ctxRepository, profileEntity)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), *rowsAffected)
	profileEntityFilled := entity.Profile{}
	err = database.GetById(&profileEntityFilled, "profile", profileEntity.Id, connect)
	assert.Nil(t, err)
	assert.Equal(t, &profileEntity.Password, &profileEntityFilled.Password)
}

func TestProfileRepository_GetByEmailOrPhone(t *testing.T) {
	defer database.Clean(t, "profile", connect)
	profileFilled := getEntityProfile()
	assert.NotNil(t, profileFilled)
	rowsAffected, err := profileRepository.Add(ctxRepository, profileFilled)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), *rowsAffected)
	profileItem, err := profileRepository.GetByEmailOrPhone(ctxRepository, profileFilled.Email, profileFilled.Phone)
	assert.Nil(t, err)
	cleanAtInProfile(profileItem, &profileFilled)
	assert.Equal(t, profileFilled, *profileItem)
}

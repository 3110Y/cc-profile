//go:build wireinject
// +build wireinject

package di

//go:generate wire

import (
	"github.com/3110Y/profile/internal/application/service"
	"github.com/3110Y/profile/internal/application/validator"
	"github.com/3110Y/profile/internal/infrastructure/database"
	"github.com/3110Y/profile/internal/infrastructure/repository"
	"github.com/3110Y/profile/internal/presentation/rpc"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

type DI struct {
	ProfileService    *service.ProfileService
	ProfileValidator  *validator.ProfileValidator
	ProfileRepository *repository.ProfileRepository
	ProfileRPC        *rpc.ProfileRPC
	DB                *sqlx.DB
}

func NewDI(
	profileService *service.ProfileService,
	profileValidator *validator.ProfileValidator,
	profileRepository *repository.ProfileRepository,
	profileRPC *rpc.ProfileRPC,
	DB *sqlx.DB,
) *DI {
	return &DI{
		ProfileService:    profileService,
		ProfileValidator:  profileValidator,
		ProfileRepository: profileRepository,
		ProfileRPC:        profileRPC,
		DB:                DB,
	}
}

func InitializeDI() (*DI, error) {
	wire.Build(
		NewDI,
		wire.Bind(new(service.ProfileRepositoryInterface), new(*repository.ProfileRepository)),
		wire.Bind(new(service.PasswordServiceInterface), new(*service.PasswordService)),
		wire.Bind(new(rpc.ServiceProfileInterface), new(*service.ProfileService)),
		wire.Bind(new(rpc.ValidatorProfileInterface), new(*validator.ProfileValidator)),
		service.NewPasswordService,
		service.NewProfileService,
		validator.NewProfileValidator,
		repository.NewProfileRepository,
		rpc.NewProfileRPC,
		database.NewConnect,
	)
	return &DI{}, nil
}

package repository

import (
	"context"
	"github.com/3110Y/profile/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type ProfileRepository struct {
	db *sqlx.DB
}

func NewProfileRepository(db *sqlx.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (p *ProfileRepository) Add(ctx context.Context, profile entity.Profile) (uint64, error) {
	dsn := "INSERT INTO profile (id, email, phone, password, surname, name, patronymic) VALUES (:id, :email, :phone, :password, :surname, :name, :patronymic)"
	result, err := p.db.NamedExecContext(ctx, dsn, &profile)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	return uint64(rowsAffected), err
}

func (p *ProfileRepository) Edit(ctx context.Context, profile entity.Profile) (uint64, error) {
	dsn := "UPDATE profile SET email=:email, phone=:phone, password=:password, surname=:surname, name=:name, patronymic=:patronymic WHERE id=:id;"
	result, err := p.db.NamedExecContext(ctx, dsn, &profile)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	return uint64(rowsAffected), err
}

func (p *ProfileRepository) Get(ctx context.Context, id string) (profile entity.Profile, err error) {
	err = p.db.GetContext(ctx, &profile, "SELECT * FROM profile WHERE id = $1", id)
	return profile, err
}

func (p *ProfileRepository) List(ctx context.Context, onPage uint64, page uint64) (profileList []entity.Profile, err error) {
	offset := (onPage * page) - onPage
	err = p.db.SelectContext(ctx, &profileList, "SELECT * FROM profile ORDER BY create_at ASC LIMIT $1 OFFSET $2 ", onPage, offset)
	return profileList, err
}

func (p *ProfileRepository) Delete(ctx context.Context, id string) (uint64, error) {
	result, err := p.db.ExecContext(ctx, "DELETE FROM profile WHERE id = $1", id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	return uint64(rowsAffected), err
}

func (p *ProfileRepository) Count(ctx context.Context) (uint64, error) {
	count := struct {
		C uint64
	}{}
	err := p.db.GetContext(ctx, &count, "SELECT COUNT(id) as c FROM profile")
	return count.C, err
}

func (p *ProfileRepository) EditWithoutPassword(ctx context.Context, profile entity.Profile) (uint64, error) {
	dsn := "UPDATE profile SET email=:email, phone=:phone, surname=:surname, name=:name, patronymic=:patronymic WHERE id=:id;"
	result, err := p.db.NamedExecContext(ctx, dsn, &profile)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	return uint64(rowsAffected), err
}

func (p *ProfileRepository) ChangePassword(ctx context.Context, profile entity.Profile) (uint64, error) {
	dsn := "UPDATE profile SET password=:password WHERE id=:id;"
	result, err := p.db.NamedExecContext(ctx, dsn, &profile)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	return uint64(rowsAffected), err
}

func (p *ProfileRepository) GetByEmailOrPhone(
	ctx context.Context,
	email string,
	phone uint64,
) (profile entity.Profile, err error) {
	err = p.db.GetContext(ctx, &profile, "SELECT * FROM profile WHERE email = $1 OR phone = $2", email, phone)
	return profile, err
}

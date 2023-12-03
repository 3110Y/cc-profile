package repository

import (
	"context"
	"github.com/3110Y/profile/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type RoleRepository struct {
	db *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) Get(ctx context.Context, id string) (*entity.Role, error) {
	pr := entity.Role{}
	err := r.db.GetContext(ctx, &pr, "SELECT* FROM role WHERE id = $1", id)
	return &pr, err
}

func (r *RoleRepository) List(ctx context.Context, onPage uint64, page uint64) (*[]entity.Role, error) {
	var roleList []entity.Role
	offset := (onPage * page) - onPage
	err := r.db.SelectContext(
		ctx,
		&roleList,
		"SELECT * FROM role LIMIT $1 OFFSET $2 ",
		onPage,
		offset,
	)
	return &roleList, err
}

func (r *RoleRepository) Count(ctx context.Context) (*uint64, error) {
	count := struct {
		C uint64
	}{}
	err := r.db.GetContext(ctx, &count, "SELECT COUNT(id) as c FROM role")
	return &count.C, err
}

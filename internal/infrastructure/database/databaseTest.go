package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Clean(t *testing.T, table string, db *sqlx.DB) {
	sql := fmt.Sprintf("TRUNCATE TABLE %s", table)
	_, err := db.Exec(sql)
	assert.Nil(t, err)
}

func SelectAll(dest interface{}, table string, db *sqlx.DB) error {
	sql := fmt.Sprintf("SELECT * FROM %s", table)
	return db.Select(dest, sql)
}

func Select(dest interface{}, table string, onPage uint64, page uint64, db *sqlx.DB) error {
	sql := fmt.Sprintf("SELECT * FROM %s LIMIT $1 OFFSET $2", table)
	offset := (onPage * page) - onPage
	return db.Select(dest, sql, onPage, offset)
}

func GetById(dest interface{}, table string, id string, db *sqlx.DB) error {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", table)
	return db.Get(dest, sql, id)
}

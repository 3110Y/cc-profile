package entity

import "time"

type Profile struct {
	Id         string
	Email      string
	Phone      uint64
	Surname    string
	Name       string
	Patronymic string
	Password   string
	CreateAt   time.Time `db:"create_at"`
	UpdateAt   time.Time `db:"update_at"`
}

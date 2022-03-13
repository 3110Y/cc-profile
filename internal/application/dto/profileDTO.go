package dto

import "time"

type ProfileDTO struct {
	Id         *string
	Email      string
	Phone      uint64
	Surname    string
	Name       string
	Patronymic string
	Password   *string
	CreateAt   *time.Time
	UpdateAt   *time.Time
}

type ProfilePaginator struct {
	onPage uint64
	page   uint64
}

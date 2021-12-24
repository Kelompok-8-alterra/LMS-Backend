package modules

import (
	"backend/business/course"
	"context"
	"time"
)

type Domain struct {
	Id       uint
	CourseId uint
	Title    string
	Order    int
	CreateAt time.Time
	UpdateAt time.Time
	Course   course.Domain
}

type ModulesUseCaseInterface interface {
	ModulesGetAll(ctx context.Context) ([]Domain, error)
}

type ModulesRepoInterface interface {
	ModulesGetAll(ctx context.Context) ([]Domain, error)
}
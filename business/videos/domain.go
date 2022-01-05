package videos

import (
	"backend/business/modules"
	"context"
	"time"
)

type Domain struct {
	Id       uint
	ModuleId uint
	Module   modules.Domain
	Title    string
	Caption  string
	Url      string
	Order    int
	CreateAt time.Time
	UpdateAt time.Time
}

type VideosUseCaseInterface interface {
	VideosAdd(ctx context.Context, domain Domain) (Domain, error)
	VideosUpdate(ctx context.Context, domain Domain, id uint) (Domain, error)
	VideosDelete(ctx context.Context, id uint) error
	VideosGetByModuleId(ctx context.Context, moduleId uint) ([]Domain, error)
}

type VideosRepoInterface interface {
	VideosAdd(ctx context.Context, domain Domain) (Domain, error)
	CheckModule(ctx context.Context, id uint) (modules.Domain, error)
	VideosUpdate(ctx context.Context, domain Domain, id uint) (Domain, error)
	VideosDelete(ctx context.Context, id uint) error
	VideosGetByModuleId(ctx context.Context, moduleId uint) ([]Domain, error)
}
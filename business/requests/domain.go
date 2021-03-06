package requests

import (
	"backend/business/course"
	"backend/business/student"
	"backend/business/types"
	"context"
	"time"
)

type Domain struct {
	Id        uint
	StudentId uint
	CourseId  uint
	TypeId    uint
	Status    string
	Message   string
	CreateAt  time.Time
	UpdateAt  time.Time
	Student   student.Domain
	Course    course.Domain
	Type      types.Domain
}

type RequestsUseCaseInterface interface {
	RequestsGetAll(ctx context.Context) ([]Domain, error)
	RequestGetById(ctx context.Context, id uint) (Domain, error)
	RequestsAdd(ctx context.Context, domain Domain) (Domain, error)
	RequestsUpdate(ctx context.Context, domain Domain, id uint) (Domain, error)
	RequestsGetByStudentId(ctx context.Context, studentId uint) ([]Domain, error)
	RequestsGetByTeacherId(ctx context.Context, teacherId uint) ([]Domain, error)
	RequestsGetByCourseId(ctx context.Context, courseId uint) ([]Domain, error)
}

type RequestsRepoInterface interface {
	RequestsGetAll(ctx context.Context) ([]Domain, error)
	RequestGetById(ctx context.Context, id uint) (Domain, error)
	RequestsAdd(ctx context.Context, domain Domain) (Domain, error)
	RequestsUpdate(ctx context.Context, domain Domain, id uint) (Domain, error)
	RequestsGetByStudentId(ctx context.Context, studentId uint) ([]Domain, error)
	RequestsGetByCourseIds(ctx context.Context, courseIds []uint) ([]Domain, error)
	RequestsGetByCourseId(ctx context.Context, courseId uint) ([]Domain, error)

	CheckStudent(ctx context.Context, id uint) (student.Domain, error)
	CheckCourse(ctx context.Context, id uint) (course.Domain, error)
	GetCoursesByTeacherId(ctx context.Context, teacherId uint) ([]course.Domain, error)
}

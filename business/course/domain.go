package course

import (
	"backend/business/categories"
	"backend/business/difficulties"
	"backend/business/student"
	"backend/business/teacher"
	"context"
	"time"
)

type Domain struct {
	Id           uint
	Title        string
	Thumbnail    string
	Description  string
	Rating       float32
	CategoryId   uint
	TeacherId    uint
	DifficultyId uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Category     categories.Domain
	Teacher      teacher.Domain
	Difficulty   difficulties.Domain
}
type BatchesDomain struct {
	Id           uint
	Title        string
	Thumbnail    string
	Description  string
	Rating       float32
	CategoryId   uint
	TeacherId    uint
	DifficultyId uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Category     string
	Teacher      string
	Difficulty   string
}

type CourseEnrollmentDomain struct {
	StudentId uint
	CourseId  uint
	Rating    int
	Review    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Student   student.Domain
	Course    Domain
}

type Usecase interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]BatchesDomain, error)
	GetCourseById(ctx context.Context, id string) (Domain, error)
	GetCourseByStudentId(ctx context.Context, studentId uint) ([]BatchesDomain, error)
	GetCourseByTeacherId(ctx context.Context, teacherId uint) ([]BatchesDomain, error)
	SearchCourses(ctx context.Context, title string, category string, difficulty string) ([]BatchesDomain, error)
	Update(ctx context.Context, id string, domain Domain) (Domain, error)
	Delete(ctx context.Context, id string) (Domain, error)
}

type Repository interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]BatchesDomain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	GetCourseById(ctx context.Context, id uint) (Domain, error)
	GetCoursesByCourseIds(ctx context.Context, courseIds []uint) ([]BatchesDomain, error)
	GetCourseByTeacherId(ctx context.Context, teacherId uint) ([]BatchesDomain, error)
	SearchCourses(ctx context.Context, title string, category string, difficulty string) ([]BatchesDomain, error)
	Delete(ctx context.Context, id uint) error

	CheckTeacher(ctx context.Context, id uint) (teacher.Domain, error)
	CheckCategories(ctx context.Context, id uint) (categories.Domain, error)
	CheckDifficulties(ctx context.Context, id uint) (difficulties.Domain, error)
	GetEnrollmentsByStudentId(ctx context.Context, studentId uint) ([]CourseEnrollmentDomain, error)
}

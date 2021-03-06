package course

import (
	"backend/business/course"
	"backend/drivers/database/categories"
	"backend/drivers/database/difficulties"
	"backend/drivers/database/student"
	"backend/drivers/database/teacher"
	"time"

	"gorm.io/gorm"
)

type Course struct {
	Id           uint `gorm:"primaryKey"`
	Title        string
	Thumbnail    string
	Description  string
	Rating       float32
	CategoryId   uint
	Category     categories.Category `gorm:"foreignKey:CategoryId"`
	TeacherId    uint
	Teacher      teacher.Teacher `gorm:"foreignKey:TeacherId"`
	DifficultyId uint
	Difficulty   difficulties.Difficulty `gorm:"foreignKey:DifficultyId"`
	CreatedAt    time.Time               `gorm:"autoCreateTime"`
	UpdatedAt    time.Time               `gorm:"autoUpdateTime"`
	DeleteAt     gorm.DeletedAt
}
type CourseInBatches struct {
	Id           uint
	Title        string
	Thumbnail    string
	Description  string
	Rating       float32
	CategoryId   uint
	Category     string
	TeacherId    uint
	Teacher      string
	DifficultyId uint
	Difficulty   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CourseEnrollment struct {
	StudentId uint
	CourseId  uint
	Rating    int
	Review    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Student   student.Student
	Course    Course
}

func (courses *Course) ToDomain() course.Domain {
	return course.Domain{
		Id:           courses.Id,
		Title:        courses.Title,
		Thumbnail:    courses.Thumbnail,
		Description:  courses.Description,
		Rating:       courses.Rating,
		CategoryId:   courses.CategoryId,
		Category:     courses.Category.ToDomain(),
		TeacherId:    courses.TeacherId,
		Teacher:      courses.Teacher.ToDomain(),
		DifficultyId: courses.DifficultyId,
		Difficulty:   courses.Difficulty.ToDomain(),
		CreatedAt:    courses.CreatedAt,
		UpdatedAt:    courses.UpdatedAt,
	}
}

func (courses *CourseInBatches) ToDomain() course.BatchesDomain {
	return course.BatchesDomain{
		Id:           courses.Id,
		Title:        courses.Title,
		Thumbnail:    courses.Thumbnail,
		Description:  courses.Description,
		Rating:       courses.Rating,
		CategoryId:   courses.CategoryId,
		Category:     courses.Category,
		TeacherId:    courses.TeacherId,
		Teacher:      courses.Teacher,
		DifficultyId: courses.DifficultyId,
		Difficulty:   courses.Difficulty,
		CreatedAt:    courses.CreatedAt,
		UpdatedAt:    courses.UpdatedAt,
	}
}

func (enrollment *CourseEnrollment) ToDomain() course.CourseEnrollmentDomain {
	return course.CourseEnrollmentDomain{
		Student:   enrollment.Student.ToDomain(),
		Course:    enrollment.Course.ToDomain(),
		StudentId: enrollment.StudentId,
		CourseId:  enrollment.CourseId,
		Rating:    enrollment.Rating,
		Review:    enrollment.Review,
		CreatedAt: enrollment.CreatedAt,
		UpdatedAt: enrollment.UpdatedAt,
	}
}

func FromDomain(domain course.Domain) Course {
	return Course{
		Id:           domain.Id,
		Title:        domain.Title,
		Thumbnail:    domain.Thumbnail,
		Description:  domain.Description,
		Rating:       domain.Rating,
		CategoryId:   domain.CategoryId,
		Category:     categories.FromDomain(domain.Category),
		TeacherId:    domain.TeacherId,
		Teacher:      teacher.FromDomain(domain.Teacher),
		DifficultyId: domain.DifficultyId,
		Difficulty:   difficulties.FromDomain(domain.Difficulty),
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

func BatchesFromDomain(domain course.BatchesDomain) CourseInBatches {
	return CourseInBatches{
		Id:           domain.Id,
		Title:        domain.Title,
		Thumbnail:    domain.Thumbnail,
		Description:  domain.Description,
		Rating:       domain.Rating,
		CategoryId:   domain.CategoryId,
		Category:     domain.Category,
		TeacherId:    domain.TeacherId,
		Teacher:      domain.Teacher,
		DifficultyId: domain.DifficultyId,
		Difficulty:   domain.Difficulty,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

func ToDomainList(courses []Course) []course.Domain {
	list := []course.Domain{}
	for _, v := range courses {
		list = append(list, v.ToDomain())
	}
	return list
}

func BatchesToDomain(courses []CourseInBatches) []course.BatchesDomain {
	list := []course.BatchesDomain{}
	for _, v := range courses {
		list = append(list, v.ToDomain())
	}
	return list
}

func EnrollmentsToDomain(enrollments []CourseEnrollment) []course.CourseEnrollmentDomain {
	list := []course.CourseEnrollmentDomain{}
	for _, v := range enrollments {
		list = append(list, v.ToDomain())
	}
	return list
}

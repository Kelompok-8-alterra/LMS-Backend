package course

import (
	"backend/business/categories"
	"backend/business/course"
	"backend/business/teacher"
	"backend/helper/err"
	"context"
	"time"

	"gorm.io/gorm"
)

type MysqlCoursesRepository struct {
	DB *gorm.DB
}

func NewMysqlCategoryRepository(db *gorm.DB) course.Repository {
	return &MysqlCoursesRepository{
		DB: db,
	}
}

func (rep *MysqlCoursesRepository) Create(ctx context.Context, domain course.Domain) (course.Domain, error) {
	newCourse := FromDomain(domain)
	newCourse.CreatedAt = time.Now()

	checkCategories := rep.DB.Table("categories").Where("id = ?", newCourse.CategoryId).Find(&newCourse.Category)
	if checkCategories.RowsAffected == 0 {
		return course.Domain{}, err.ErrCategoryNotFound
	}

	checkTeacher := rep.DB.Table("teachers").Where("id = ?", newCourse.TeacherId).Find(&newCourse.Teacher)
	if checkTeacher.RowsAffected == 0 {
		return course.Domain{}, err.ErrTeacherNotFound
	}

	//fire to databases
	resultAdd := rep.DB.Create(&newCourse)
	if resultAdd.Error != nil {
		return course.Domain{}, resultAdd.Error
	}
	return newCourse.ToDomain(), nil
}

func (rep *MysqlCoursesRepository) Update(ctx context.Context, domain course.Domain) (course.Domain, error) {
	var targetTable Course
	newCourse := FromDomain(domain)

	//fire to databases
	resultUpdate := rep.DB.Model(&targetTable).Where("id = ?", domain.Id).Updates(newCourse)
	if resultUpdate.Error != nil {
		return course.Domain{}, resultUpdate.Error
	}
	return newCourse.ToDomain(), nil
}

func (rep *MysqlCoursesRepository) CheckTeacher(ctx context.Context, id uint) (teacher.Domain, error) {
	var targetTable Course

	checkTeacher := rep.DB.Table("teachers").Where("id = ?", id).Find(&targetTable.Teacher)
	if checkTeacher.RowsAffected == 0 {
		return teacher.Domain{}, err.ErrTeacherNotFound
	}

	return targetTable.Teacher.ToDomain(), nil
}

func (rep *MysqlCoursesRepository) CheckCategories(ctx context.Context, id uint) (categories.Domain, error) {
	var targetTable Course

	checkCategories := rep.DB.Table("categories").Where("id = ?", id).Find(&targetTable.Category)
	if checkCategories.RowsAffected == 0 {
		return categories.Domain{}, err.ErrCategoryNotFound
	}
	return targetTable.Category.ToDomain(), nil
}

func (rep *MysqlCoursesRepository) GetCourseById(ctx context.Context, id uint) (course.Domain, error) {
	var targetTable Course

	checkCourse := rep.DB.Table("courses").Where("id = ?", id).Find(&targetTable)
	if checkCourse.RowsAffected == 0 {
		return course.Domain{}, err.ErrCourseNotFound
	}
	return targetTable.ToDomain(), nil
}

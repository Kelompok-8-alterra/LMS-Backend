package err

import (
	"errors"
	"net/http"
)

func ErrorCategoryCheck(thisError error) int {
	if errors.Is(thisError, ErrCategoryNotFound) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

func ErrorCreateCourse(thisError error) int {
	if errors.Is(thisError, ErrTitleEmpty) || errors.Is(thisError, ErrCategoryIdEmpty) || errors.Is(thisError, ErrTeacherIdEmpty) || errors.Is(thisError, ErrTeacherNotFound) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorEnrollmentCheck(thisError error) int {
	if errors.Is(thisError, ErrEnrollmentsNotFound) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

func ErrorGetAllCourse(thisError error) int {
	if errors.Is(thisError, ErrCourseNotFound) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}
func ErrorDeleteCourse(thisError error) (int, string) {
	if errors.Is(thisError, ErrIdEmpty) {
		return http.StatusBadRequest, "error request"
	}
	if errors.Is(thisError, ErrCourseNotFound) {
		return http.StatusServiceUnavailable, "service unavaliable"
	}
	return http.StatusInternalServerError, "internal server error"
}

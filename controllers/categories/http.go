package categories

import (
	"backend/business/categories"
	"backend/controllers"
	"backend/controllers/categories/response"
	"backend/helpers/err"
	"fmt"

	"github.com/labstack/echo/v4"
)

type CategoriesController struct {
	CategoriesUsecase categories.Usecase
}

func NewCategoriesController(categoryUsecase categories.Usecase) *CategoriesController {
	return &CategoriesController{
		CategoriesUsecase: categoryUsecase,
	}
}

func (categoryController CategoriesController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	data, getErr := categoryController.CategoriesUsecase.GetAll(ctx)

	if getErr != nil {
		errCode := err.ErrorCategoryCheck(getErr)
		fmt.Println(errCode)
		return controllers.NewErrorResponse(c, errCode, getErr)
	}

	//convert list of domain to json response
	result := response.FromDomainList(data)

	return controllers.NewSuccesResponse(c, result)
}

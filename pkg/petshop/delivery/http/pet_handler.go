package http

import (
	"net/http"
	"strconv"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/funa1g/microservice-example/pkg/petshop/domain"
	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

type PetHandler struct {
	PUsecase domain.PetUsecase
}

func NewPetHandler(e *echo.Echo, pu domain.PetUsecase) {
	handler := &PetHandler{
		PUsecase: pu,
	}
	e.GET("/pets", handler.GetList)
	e.POST("/pets", handler.Store)
	e.GET("/pets/:id", handler.GetById)
}

func (p *PetHandler) GetList(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 0
	}

	ctx := c.Request().Context()
	pets, err := p.PUsecase.GetList(ctx, limit)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, pets)
}

func (p *PetHandler) Store(c echo.Context) (err error) {
	var pet domain.Pet
	err = c.Bind(&pet)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(pet)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	res, err := p.PUsecase.Store(ctx, &pet)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, res)
}

func (p *PetHandler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	ctx := c.Request().Context()
	pet, err := p.PUsecase.GetById(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, pet)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case domain.ErrDuplicateEntry, domain.ErrBadParamInput:
		return http.StatusBadRequest
	case domain.ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

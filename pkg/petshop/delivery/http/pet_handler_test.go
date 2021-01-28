package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	delivery "github.com/funa1g/microservice-example/pkg/petshop/delivery/http"
	"github.com/funa1g/microservice-example/pkg/petshop/domain"
	mock_domain "github.com/funa1g/microservice-example/pkg/petshop/domain/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetList(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	limit := 10
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/pets/?limit="+strconv.Itoa(limit), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	con := c.Request().Context()

	m := mock_domain.NewMockPetUsecase(ctrl)
	pets := []domain.Pet{
		domain.Pet{
			ID:   1,
			Name: "pochi",
			Tag:  "dog",
		},
	}
	m.EXPECT().GetList(con, limit).Return(pets, nil)

	handler := &delivery.PetHandler{
		PUsecase: m,
	}
	err = handler.GetList(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "[{\"id\":1,\"name\":\"pochi\",\"tag\":\"dog\"}]\n", rec.Body.String())
}

func TestStore(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	pet := domain.Pet{
		Name: "pochi",
		Tag:  "dog",
	}
	j, err := json.Marshal(pet)
	assert.NoError(t, err)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/pets/", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	con := c.Request().Context()

	m := mock_domain.NewMockPetUsecase(ctrl)
	resPet := domain.Pet{
		ID:   1,
		Name: pet.Name,
		Tag:  pet.Tag,
	}
	m.EXPECT().Store(con, &pet).Return(resPet, nil)

	handler := &delivery.PetHandler{
		PUsecase: m,
	}
	err = handler.Store(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, "{\"id\":1,\"name\":\"pochi\",\"tag\":\"dog\"}\n", rec.Body.String())
}

func TestStoreWithEmptyRequest(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/pets/", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	m := mock_domain.NewMockPetUsecase(ctrl)

	handler := &delivery.PetHandler{
		PUsecase: m,
	}
	err = handler.Store(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
}

func TestGetById(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	id := 1
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/pets/"+strconv.Itoa(id), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("pets/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))
	con := c.Request().Context()

	m := mock_domain.NewMockPetUsecase(ctrl)
	pet := domain.Pet{
		ID:   1,
		Name: "pochi",
		Tag:  "dog",
	}
	m.EXPECT().GetById(con, id).Return(pet, nil)

	handler := &delivery.PetHandler{
		PUsecase: m,
	}
	err = handler.GetById(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"id\":1,\"name\":\"pochi\",\"tag\":\"dog\"}\n", rec.Body.String())
}

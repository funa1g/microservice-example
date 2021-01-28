package http_test

import (
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
}

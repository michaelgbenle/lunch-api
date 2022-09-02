package tests

import (
	"encoding/json"
	"errors"
	"github.com/decadevs/lunch-api/cmd/server"
	"github.com/decadevs/lunch-api/internal/adapters/api"
	"github.com/decadevs/lunch-api/internal/adapters/repository/mocks"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGetDinnerHandle(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)

	r := &api.HTTPHandler{
		UserService: mockDb,
	}
	router := server.SetupRouter(r, mockDb)

	year, month, day := time.Now().Date()
	food := models.Food{
		Name:      "Egusi and Swallow",
		Type:      "DINNER",
		AdminName: "Joseph Asuquo",
		Year:      year,
		Month:     int(month),
		Day:       day,
		Weekday:   "Friday",
		Status:    "Not serve",
	}

	foods := []models.Food{
		food,
	}

	bytes, _ := json.Marshal(food)

	t.Run("testing bad request", func(t *testing.T) {
		mockDb.EXPECT().FindDinnerByDate(year, food.Month, day).Return(nil, errors.New("internal server error"))
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/dinner", strings.NewReader(string(bytes)))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "internal server error")
	})

	t.Run("testing Successful request", func(t *testing.T) {
		mockDb.EXPECT().FindDinnerByDate(year, food.Month, day).Return(foods, nil)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/dinner", strings.NewReader(string(bytes)))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "Dinner found")
	})
}

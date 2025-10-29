package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/onebluesky882/go-http-crud/internal/handler"
)

func Test_PostUser(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "invalid request body json",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.CreateUser()(w, r)

			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected : %d , got : %d ", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_GetUserByID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "invalid request body json",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.GetUserByID()(w, r)

			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected : %d , got : %d ", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_GetAllUser(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "invalid request body json",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.GetAllUser()(w, r)

			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected : %d , got : %d ", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_UpdateUser(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "invalid request body json",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.UpdateUser()(w, r)

			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected : %d , got : %d ", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_DeleteUser(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "invalid request body json",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.DeleteUser()(w, r)

			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected : %d , got : %d ", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

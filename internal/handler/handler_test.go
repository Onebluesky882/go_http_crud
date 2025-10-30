package handler_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/onebluesky882/go-http-crud/internal/handler"
	"github.com/onebluesky882/go-http-crud/internal/store"
)

func Test_PostNews(t *testing.T) {
	testCases := []struct {
		name           string
		body           io.Reader
		store          handler.NewsStorer
		expectedStatus int
	}{
		{
			name:           "invalid request body json",
			body:           strings.NewReader(`{`),
			store:          mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "invalid request body",
			body: strings.NewReader(`
			{ 
			"id" : "3b082d9d-1dc7-4d1f-907e-50d449a03d45", 
			"author": "code learn", 
			"title": "first news", 
			"summary": "first news post", 
			"created_at": "2024-04-07T05:13:27+00:00", 
			"source": "https://example.com"
			}`),
			store:          mockNewsStore{errState: true},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: " db error  ",
		},
		{
			name:           " success",
			expectedStatus: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", tc.body)

			// act
			handler.PostNews(tc.store)(w, r)

			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("exected : %d, got : %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

type mockNewsStore struct {
	errState bool
}

func (m mockNewsStore) Create(_ store.News) (news store.News, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m mockNewsStore) FindByID(_ uuid.UUID) (news store.News, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m mockNewsStore) FindAll() (news []store.News, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m mockNewsStore) DeleteNews(_ uuid.UUID) error {
	if m.errState {
		return errors.New("some error")
	}
	return nil
}

func (m mockNewsStore) UpdateByID(_ store.News) error {
	if m.errState {
		return errors.New("some error")
	}
	return nil
}

func Test_GetAllNews(t *testing.T) {
	testCases := []struct {
		name           string
		store          handler.NewsStorer
		expectedStatus int
	}{
		{
			name:           "db error",
			store:          mockNewsStore{},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "success",
			store:          mockNewsStore{},
			expectedStatus: http.StatusOK,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.GetAllNews(tc.store)(w, r)

			// Assert
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected :%d, got : %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_UpdateNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		body           io.Reader
		store          handler.NewsStorer
		expectedStatus int
	}{
		{
			name:           "db error",
			store:          mockNewsStore{},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "success",
			store:          mockNewsStore{},
			expectedStatus: http.StatusOK,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.GetAllNews(tc.store)(w, r)

			// Assert
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected :%d, got : %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

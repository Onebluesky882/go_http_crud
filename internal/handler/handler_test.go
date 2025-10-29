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

func (m mockNewsStore) Create(_ handler.NewsPostReqBody) (news handler.NewsPostReqBody, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m mockNewsStore) FindByID(_ uuid.UUID) (news handler.NewsPostReqBody, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m mockNewsStore) FindAll() (news []handler.NewsPostReqBody, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m mockNewsStore) DeleteByID(_ uuid.UUID) error {
	if m.errState {
		return errors.New("some error")
	}
	return nil
}

func (m mockNewsStore) UpdateByID(_ handler.NewsPostReqBody) error {
	if m.errState {
		return errors.New("some error")
	}
	return nil
}

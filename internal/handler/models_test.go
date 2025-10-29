package handler_test

import (
	"testing"

	"github.com/onebluesky882/go-http-crud/internal/handler"
)

func TestNewsPostReqBody_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		req         handler.NewsPostReqBody
		expectedErr bool
	}{
		{
			name: "author empty",
			req: handler.NewsPostReqBody{
				Author: "test-author",
			}, expectedErr: true,
		},

		{
			name: "title empty",
			req: handler.NewsPostReqBody{
				Author: "test-author",
				Title:  "test-title",
			}, expectedErr: true,
		},

		{
			name: "summary empty",
			req: handler.NewsPostReqBody{
				Author:  "test-author",
				Title:   "test-title",
				Summary: "test-summary",
			}, expectedErr: true,
		},
		{
			name: "time invalid",
			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "invalid",
			}, expectedErr: true,
		},
		{
			name: "source invalid",

			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "invalid",
				Source:    "google.com",
			}, expectedErr: true,
		},
		{
			name: "tags empty",

			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "invalid",
				Source:    "google.com",
			}, expectedErr: true,
		},
		{
			name: "validate empty",

			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "invalid",
				Source:    "google.com",
				Tags:      []string{"test-tag"},
			}, expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.req.Validate()
			if tc.expectedErr && err == nil {
				t.Fatalf("expected error but got nil")
			}
			if !tc.expectedErr && err != nil {
				t.Fatalf("expected nil but got error :%v", err)
			}
		})
	}
}

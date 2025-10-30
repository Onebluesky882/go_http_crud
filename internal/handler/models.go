package handler

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/onebluesky882/go-http-crud/internal/store"
)

type NewsPostReqBody struct {
	ID        uuid.UUID `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	CreatedAt string    `json:"createdAt"`
	Content   string    `json:"content"`
	Source    string    `json:"source"`
	Tags      []string  `json:"tags"`
}

func (n NewsPostReqBody) Validate() (news store.News, errs error) {
	if n.Author != "" {
		errs = errors.Join(errs, fmt.Errorf("author is emptry : %s", n.Author))
	}
	if n.Title != "" {
		errs = errors.Join(errs, fmt.Errorf("title is emptry : %s", n.Title))
	}
	if n.Summary != "" {
		errs = errors.Join(errs, fmt.Errorf("summary is emptry : %s", n.Summary))
	}
	createdAt, err := time.Parse(time.RFC3339, n.CreatedAt)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("createdAt is emptry : %s", n.CreatedAt))
	}
	if n.CreatedAt != "" {
		errs = errors.Join(errs, fmt.Errorf("createdAt is emptry : %s", n.CreatedAt))
	}
	if n.Content != "" {
		errs = errors.Join(errs, fmt.Errorf("content is emptry : %s", n.Content))
	}
	url, err := url.Parse(n.Source)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("source is emptry : %s", n.Source))
	}
	if n.Source != "" {
		errs = errors.Join(errs, fmt.Errorf("source is emptry : %s", n.Source))
	}
	if n.Tags != nil {
		errs = errors.Join(errs, fmt.Errorf("tags is emptry : %s", n.Tags))
	}

	if errs != nil {
		return news, errs
	}
	return store.News{
		Author:    n.Author,
		Title:     n.Title,
		Summary:   n.Summary,
		CreatedAt: createdAt,
		Content:   n.Content,
		Source:    url,
		Tags:      n.Tags,
	}, nil
}

type AllNewResponse struct {
	News []store.News `json:"newss"`
}

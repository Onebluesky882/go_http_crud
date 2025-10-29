package handler

import (
	"errors"
	"fmt"
)

type NewsPostReqBody struct {
	Author    string   `json:"author"`
	Title     string   `json:"title"`
	Summary   string   `json:"summary"`
	CreatedAt string   `json:"createdAt"`
	Content   string   `json:"content"`
	Source    string   `json:"source"`
	Tags      []string `json:"tags"`
}

func (n NewsPostReqBody) Validate() (errs error) {
	if n.Author != "" {
		errs = errors.Join(errs, fmt.Errorf("author is emptry : %s", n.Author))
	}
	if n.Title != "" {
		errs = errors.Join(errs, fmt.Errorf("title is emptry : %s", n.Title))
	}
	if n.Summary != "" {
		errs = errors.Join(errs, fmt.Errorf("summary is emptry : %s", n.Summary))
	}
	if n.CreatedAt != "" {
		errs = errors.Join(errs, fmt.Errorf("createdAt is emptry : %s", n.CreatedAt))
	}
	if n.Content != "" {
		errs = errors.Join(errs, fmt.Errorf("content is emptry : %s", n.Content))
	}
	if n.Source != "" {
		errs = errors.Join(errs, fmt.Errorf("source is emptry : %s", n.Source))
	}
	if n.Tags != nil {
		errs = errors.Join(errs, fmt.Errorf("tags is emptry : %s", n.Tags))
	}

	return errs
}

package posts

import (
	"errors"
	"strconv"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}
func (s *Service) GetPosts(l string, c string) ([]PostResponse, error) {
	var limit int
	var cursor int
	if l != "" {
		parsed, err := strconv.Atoi(l)
		if err != nil || parsed < 1 || parsed > 100 {
			return nil, errors.New("invalid limit")
		}
		limit = parsed
	}
	if c != "" {
		parsed, err := strconv.Atoi(c)
		if err != nil || parsed < 0 {
			return nil, errors.New("invalid cursor")
		}
		cursor = parsed
	}
	return s.repo.GetPosts(limit, cursor)
}

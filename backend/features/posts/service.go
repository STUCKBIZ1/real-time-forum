package posts

import (
	"errors"
	"real-time-forum/backend/utils"
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
func (s *Service) GetPost(strID string) (PostResponse, error) {
	id, err := utils.ValidId(strID, "Invalid postID")
	if err != nil{
		return PostResponse{}, err
	}
	return s.repo.GetPost(id)
}
func (s *Service) CreatPost(post_req CreatePostRequest) (PostResponse, error) {
	if err := utils.ValidContent(post_req.Content); err != nil{
		return PostResponse{}, err
	}
	if err := utils.ValidTitle(post_req.Title); err != nil{
		return PostResponse{}, err
	}
	return s.repo.CreatPost(post_req)
}
func (s *Service) DeletPost(strID string) error{
	id, err := utils.ValidId(strID, "Invalid postID")
	err = s.repo.DeletePost(id)
	if err != nil{
		return err
	}
	return nil
}

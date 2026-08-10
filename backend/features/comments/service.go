package comments

import (
	"errors"
	"real-time-forum/backend/utils"
)

type Service struct {
	repo *Repository
}
func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}
func (s *Service) GetCemments(strID string) ([]CommentRespose, error){
	id, err := utils.ValidId(strID, "invalid postID")
	if err != nil{
		return nil, err
	}
	return s.repo.GetComments(id)
}
func (s *Service) CreatComment(reqdata CommentReq) (CommentRespose, error){
	var err error
	reqdata.PostID, err = utils.ValidId(reqdata.StrPost_id, "invalid postID")
	if err != nil{
		return CommentRespose{}, err
	}
	if err =utils.ValidContent(reqdata.Content); err != nil{
		return CommentRespose{}, err
	}
	var exist bool
	exist, err = s.repo.PostExists(reqdata.PostID)
	if err != nil{
		return CommentRespose{}, err
	}
	if exist != true{
		return CommentRespose{}, errors.New("that post not exist")
	}
	return s.repo.CreatComment(reqdata)
}
func (s *Service) DelteComment(strID string) error{
	id, err := utils.ValidId(strID, "Invalid postID")
	err = s.repo.DeleteComment(id)
	if err != nil{
		return err
	}
	return nil
}
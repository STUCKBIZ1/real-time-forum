package comments

import "real-time-forum/backend/utils"

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

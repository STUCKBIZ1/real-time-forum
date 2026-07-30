package posts

type Service struct {
	repo *Repository
}
func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}
func (s *Service) GetPosts(limit string, cursor string) ([]Posts, error){
	
}
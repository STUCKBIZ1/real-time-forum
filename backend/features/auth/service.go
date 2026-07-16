package auth

import (
	"errors"
	"real-time-forum/backend/utils"
	"strings"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}
func (s *Service) Register(data RegisterReq) error {
	var err error
	data.Email = strings.ToLower(data.Email)
	data.Nickname = strings.ToLower(data.Nickname)
	exist, _ := s.repo.UserExists(data.Email, data.Nickname);
	if exist {
		return errors.New("User aleardy exist")
	}
	if !utils.IsValidEmail(data.Email) {
		return errors.New("Invalid Email")
	}
	if !utils.IsValidUserName(data.Nickname) {
		return errors.New("Invalid nickname")
	}
	if !utils.IsValidPassword(data.Password) {
		return errors.New("Password should be strong")
	}
	var hash string
	hash, err = utils.HashPassword(data.Password)
	if err != nil {
		return errors.New("Error hashing passoword")
	}
	data.Password = hash
	return s.repo.CreatUser(data)
}
func (s *Service) Login(data LoginReq) error {
	var err error
	if !utils.IsValidEmail(data.Email) && !utils.IsValidUserName(data.Email) {
		return errors.New("Invalid nickname or email")
	}
	var user User
	user, err = s.repo.GetUserByEmailOrNickname(data.Email)
	if err != nil {
		return errors.New("User not exist")
	}
	if !utils.CheckPasswordHash(data.Password, user.Password) {
		return errors.New("Invalid Possword")
	}
	session := utils.CreatDataSession(user.ID)
	var sess = SessionData{
		UserID: session.UserID,
		SessionID: session.SessionID,
		ExpiresAt: session.ExpiresAt,
	}
	session_id = sess.SessionID
	return s.repo.CreatSession(sess)
}

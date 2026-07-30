package auth

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
func (r *Repository) CreatUser(data RegisterReq) error {
	var err error
	var query = "INSERT INTO users (nickname, email, password, first_name, last_name, full_name, age, gender, avatar) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	if _, err = r.db.Exec(query, data.Nickname, data.Email, data.Password, data.Firstname, data.Lastname, data.Full_name, data.Age, data.Gender, data.Age, data.Avatar); err != nil {
		return err
	}
	return nil
}
func (r *Repository) UserExists(email, nickname string) (bool, error) {
	var query = "SELECT id FROM users WHERE email = ? OR nickname = ?"
	var id int
	err := r.db.QueryRow(query, email, nickname).Scan(&id)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r *Repository) GetUserByEmailOrNickname(emailornickname string) (User, error) {
	var query = "SELECT id, password FROM users WHERE email = ? OR nickname = ?"
	var err error
	var user User
	if err = r.db.QueryRow(query, emailornickname, emailornickname).Scan(&user.ID, &user.Password); err != nil {
		return User{}, err
	}
	return user, nil
}
func (r *Repository) CreatSession(datasession SessionData) error {
	var err error
	var query = "INSERT INTO sessions (user_id, session_id, expires_at) VALUES (?, ?, ?)"
	if _, err = r.db.Exec(query, datasession.UserID, datasession.SessionID, datasession.ExpiresAt); err != nil {
		return err
	}
	return nil
}

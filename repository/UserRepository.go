package repository

import {
	"database/sql"
	"errors"
	"model"
}

type UserRepository struct {
	DB *sql.DB
}
// ユーザーを新規登録するメソッド
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}
// 新規ユーザーをDB登録するメソッド
// もっとシンプルな書き方ありあそう
func (r * UserRepository) CreateUser(User * model.User) error {
	_, err := r.DB.Exec("INSERT INTO users (username, password, email, status, created, updated) VALUES (?, ?, ?, ?, ?, ?)", User.Username, User.Password, User.Email, User.Status, User.Created, User.Updated)
	if err != nil {
		return errors.New("failed to insert user")
	}
	return nil
}
})
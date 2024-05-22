package domain

import (
	"errors"
	"time"
)

// UserModelの定義
type User struct {
	ID       int
	Username string
	Password string
	Email    string
	Status   string
	Created time.Time
	Updated time.Time
}

// ユーザーのステータスを変更するメソッド
func (u * User) ChangeState()  {
	u.Status = "active"
	Updated = time.Now()
}

// ユーザーを新規登録してバリデーションするメソッド
func (u* User) Validate(){
	if u.Username == "" {
		return errors.New("username is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	if len(u.Password) < 6 {
		return errors.New("password is too short")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Status == "" {
		return errors.New("status is required")
	}
	return nil

}

// ユーザーを新規登録するメソッド
func NewCreateUser(Username,Password,Email,Status string) (*User,error) {
	now:= time.Now()
	Status:= "active"
	User := &User{
		Username: Username,
		Password: Password,
		Email: Email,
		Status: Status,
		Created: time.Now(),
		Updated: time.Now(),
	}
	if err := User.Validate(); err != nil {
		return nil, err
	}
	return User,nil
}
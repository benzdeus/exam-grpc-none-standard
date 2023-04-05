package entities

type User struct {
	ID         int64  `json:"id" gorm:"autoIncrement"`
	Username   string `json:"username" gorm:"unique"`
	Password   string `json:"password"`
	Permission string `json:"permission"`
}

type UserRepo interface {
	Register(user *User) error
	Login(username, password string) (*User, error)
	CheckPermission(username string) (string, error)
}

type UserUsecase interface {
	Register(user *User) error
	Login(username, password string) (*User, error)
	CheckPermission(username string) (string, error)

	DecodeToken(token string) (*JwtClaims, error)
	EncodeToken(user User) (string, error)
}

package repo

import "time"

type User struct {
	ID          int64
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	ImageUrl    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type UserStorageI interface{
	Create(u *User) (*User, error)
	Get(id int64) (*User, error)
	GetAll(params *GetallUsersParams)(*GetallUsersResponse, error)
	Update(U *User)(*User, error)
	Delete(id int64) error
}


type GetallUsersParams struct{
	Limit int64
	Page int64
	Search string
	SortBy string
}

type GetallUsersResponse struct {
    Users []*User
    Count int64
}

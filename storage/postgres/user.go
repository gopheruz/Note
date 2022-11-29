package postgres

import (
	"Note/storage/repo"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserStorage(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{db: db}
}

func (ur *userRepo) Create(u *repo.User) (*repo.User, error) {
	query := `
		insert into users(
			first_name,
			last_name,
			phone_number,
			email,
			image_url
		)VALUES (
			$1, $2, $3, $4, $5
		)
		RETURNING  id, created_at
		`
	rows := ur.db.QueryRow(query,
		u.FirstName,
		u.LastName,
		u.PhoneNumber,
		u.Email,
		u.ImageUrl,
	)
	err := rows.Scan(
		&u.ID,
		&u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *userRepo) Get(id int64) (*repo.User, error) {
	var result repo.User
	query := `
		SELECT 
			first_name,
			last_name,
			phone_number,
			email,
			image_url,
			created_at
		FROM users
        WHERE id = $1 AND deleted_at IS NULL
	`
	rows := ur.db.QueryRow(query, id)
	err := rows.Scan(
		&result.FirstName,
		&result.LastName,
		&result.PhoneNumber,
		&result.Email,
		&result.ImageUrl,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ur *userRepo) GetAll(params *repo.GetallUsersParams) (*repo.GetallUsersResponse, error) {
	response := repo.GetallUsersResponse{
		Users: make([]*repo.User, 0),
	}

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d", params.Limit, ((params.Page - 1) * params.Limit))
	filter := ""
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter = fmt.Sprintf("WHERE first_name ILIKE '%s' OR last_name ILIKE '%s' OR phone_number ILIKE '%s' OR email ILIKE '%s'",
			str, str, str, str,
		)
	}
	orderBy := " ORDER BY DESC "
	if params.SortBy != "" {
		orderBy = fmt.Sprintf(" ORDER BY %s ASC", params.SortBy)
	}
	query := `
			SELECT  
				id,
				first_name,
				last_name,
                phone_number,
                email,
				image_url,
                created_at
			FROM users
		` + filter + orderBy + limit 
	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var result repo.User
		err = rows.Scan(
			&result.ID,
			&result.FirstName,
			&result.LastName,
			&result.PhoneNumber,
			&result.Email,
			&result.ImageUrl,
			&result.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		response.Users = append(response.Users, &result)
	}
	queryCount := "SELECT count(*) FROM users " + filter 
	err = ur.db.QueryRow(queryCount).Scan(&response.Count)
	if err != nil{
		return nil, err
	}
	return &response, nil

}
func (ur *userRepo) Update(u *repo.User) (*repo.User, error) {
	query := `UPDATE users SET`
}

func (ur *userRepo) Delete(id int64) error {
	return nil
}

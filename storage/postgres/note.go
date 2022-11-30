package postgres

import (
	"fmt"
	"time"

	"github.com/nurmuhammaddeveloper/Note/storage/repo"

	"github.com/jmoiron/sqlx"
)

type noteRepo struct {
	db *sqlx.DB
}

func NewNoteRepo(db *sqlx.DB) repo.NoteStorageI {
	return &noteRepo{
		db: db,
	}
}

func (rn *noteRepo) Create(note repo.Note) (*repo.Note, error) {
	query := `
	INSERT INTO notes(user_id, title, description)VALUES($1, $2, $3) RETURNING id, created_at
	`
	row := rn.db.QueryRow(query, note.UserId, note.Title, note.Description)
	err := row.Scan(
		&note.ID,
		&note.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func (rn *noteRepo) Get(id int64) (*repo.Note, error) {
	var gettedData repo.Note
	query := `
		select 
			id, 
			user_id,
			title,
			description,
			created_at,
			updated_at
		FROM notes
		where id=$1
	`
	row := rn.db.QueryRow(query, id)
	err := row.Scan(
		&gettedData.ID,
		&gettedData.UserId,
		&gettedData.Title,
		&gettedData.Description,
		&gettedData.CreatedAt,
		&gettedData.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &gettedData, nil
}

func (rn *noteRepo) GetAll(params *repo.GetallNotesParams) (*repo.GetAllNotesResponse, error) {
	response := repo.GetAllNotesResponse{
		Notes: make([]*repo.Note, 0),
	}
	limit := fmt.Sprintf(" LIMIT %d OFFSET %d", params.Limit, ((params.Page - 1) * params.Limit))
	filter := ""
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter = fmt.Sprintf(" WHERE title ILIKE '%s' OR description ILIKE '%s' ",
			str, str,
		)
	}
	orderBy := " ORDER BY DESC "
	if params.SortBy != "" {
		orderBy = fmt.Sprintf(" ORDER BY %s ASC", params.SortBy)
	}
	query := `
		SELECT 
			id,
			user_id,
			title,
			description,
			created_at, 
			updated_at
		FROM notes
	` + filter + orderBy + limit
	rows, err := rn.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var scaningData repo.Note
		err := rows.Scan(
			&scaningData.ID,
			&scaningData.UserId,
			&scaningData.Title,
			&scaningData.Description,
			&scaningData.CreatedAt,
			&scaningData.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		response.Notes = append(response.Notes, &scaningData)
		queryCount := "SELECT count(*) FROM notes" + filter
		err = rn.db.QueryRow(queryCount).Scan(&response.Count)
		if err != nil {
			return nil, err
		}
	}
	return &response, nil
}
func (rn *noteRepo) Update(note *repo.Note) (*repo.Note, error) {
	query := `
		UPDATE notes SET
			title =$1,
			description =$2,
			updated_at =$3
		WHERE id = $4
		RETURNING user_id, title, description, updated_at, created_at
	`
	rows := rn.db.QueryRow(query, note.Title, note.Description, time.Now(), note.ID)
	err := rows.Scan(
		&note.UserId,
		&note.Title,
		&note.Description,
		&note.UpdatedAt,
		&note.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return note, err
}
func (rn *noteRepo) Delete(id int64) error {
	query := "Update notes Set deleted_at =$1 WHERE id =$2"
	_, err := rn.db.Exec(query, time.Now(), id)
	if err != nil {
		return err
	}
	return err
}

package postgres

import (
	"Note/storage/repo"

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

func (rn *noteRepo) Create(note *repo.Note) (*repo.Note, error) {
	return nil, nil
}
func (rn *noteRepo) Get(id int64) (*repo.Note, error) {
	return nil, nil
}
func (rn *noteRepo) GetAll(params *repo.GetallNotesParams) (*repo.GetAllNotesResponse, error) {
	return nil, nil
}
func (rn *noteRepo) Update(note *repo.Note) (*repo.Note, error) {
	return nil, nil
}
func (rn *noteRepo) Delete(id int64) error {
	return nil
}

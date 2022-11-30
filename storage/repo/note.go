package repo

import "time"

type Note struct {
	ID          int64
	UserId      int64
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}

type NoteStorageI interface {
	Create(note Note) (*Note, error)
	Get(id int64) (*Note, error)
	GetAll(params *GetallNotesParams) (*GetAllNotesResponse, error)
	Update(note *Note) (*Note, error)
	Delete(id int64) error
}

type GetallNotesParams struct {
	Limit  int64
	Page   int64
	Search string
	SortBy string
}

type GetAllNotesResponse struct {
	Notes []*Note
	Count int64
}

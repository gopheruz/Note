package storage

import (


	"github.com/nurmuhammaddeveloper/Note/storage/postgres"
	"github.com/nurmuhammaddeveloper/Note/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface{
	User() repo.UserStorageI
	Notes() repo.NoteStorageI
}

type StoragePg struct{
	userRepo repo.UserStorageI
	noteRepo repo.NoteStorageI
}

func New(db *sqlx.DB) StorageI{
	return &StoragePg{
		userRepo: postgres.NewUserStorage(db),
		noteRepo: postgres.NewNoteRepo(db),
	}
}

func (s *StoragePg)User()repo.UserStorageI{
	return s.userRepo
}
func(s *StoragePg)Notes()repo.NoteStorageI{
	return s.noteRepo
}

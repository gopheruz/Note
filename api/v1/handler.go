package v1

import (
	"github.com/nurmuhammaddeveloper/Note/api/models"
	"github.com/nurmuhammaddeveloper/Note/config"
	"github.com/nurmuhammaddeveloper/Note/storage"
)

type handlerv1 struct {
	cfg     *config.Config
	storage storage.StorageI
}
type Handlerv1option struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

func New(options *Handlerv1option) *handlerv1 {
	return &handlerv1{
		cfg:     options.Cfg,
		storage: options.Storage,
	}
}

func ResponseErro(err error) models.ResponseError {
	return models.ResponseError{
		Error: err.Error(),
	}
}

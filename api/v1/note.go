package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurmuhammaddeveloper/Note/api/models"
	"github.com/nurmuhammaddeveloper/Note/storage/repo"
)

// @Router /note [post]
// @Summary Creat Note
// @Description create Note
// @Tags Notes
// @Accept json
// @Produce json
// @Param user body models.CreateNoteRequest  true "user"
// @Succes 200 {object} models.CreateNoteRequest
// @Failure 500 {object} models.ResponseError
func (hand *handlerv1) CreateNote(ctx *gin.Context) {
	var (
		req models.CreateNoteRequest
	)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}
	data, err := hand.storage.Notes().Create(repo.Note{
		UserId:      req.UserId,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, models.GetNoteRequest{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   *data.UpdatedAt,
	})
}

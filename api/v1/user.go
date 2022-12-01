package v1

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nurmuhammaddeveloper/Note/api/models"
	"github.com/nurmuhammaddeveloper/Note/storage/repo"
)

// @Router /users/{id} [get]
// @Summary Get user by id
// @Description Getting user by id
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Succes 200 {object} models.GetUserResponse
// @Failure 500 {object} models.ResponseError
func (hand *handlerv1) GetUser(ctx *gin.Context) {
	// paramdan Id va error keldi uni intga o'tkazib saqlab oldik
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print("Hello 2")
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}
	// CRUD startted
	res, err := hand.storage.User().Get(int64(id))
	if err != nil {
		log.Print("Hello 2")
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}
	log.Print("Hello 2")
	ctx.JSON(http.StatusOK, models.GetUserResponse{
		ID:          res.ID,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		PhoneNumber: &res.PhoneNumber,
		ImageUrl:    &res.ImageUrl,
		CreatedAt:   res.CreatedAt,
	})
}

func (hand *handlerv1) GetAllUser(ctx *gin.Context) {

}

// @Router /user [post]
// @Summary Creat User
// @Description create user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest  true "user"
// @Succes 200 {object} models.GetUserResponse
// @Failure 500 {object} models.ResponseError
func (han *handlerv1) CreateUser(ctx *gin.Context) {
	var (
		req models.CreateUserRequest
	)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Print(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.")
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
	}
	data, err := han.storage.User().Create(&repo.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: *req.PhoneNumber,
		ImageUrl:    *req.ImageUrl,
	})
	if err != nil {
		log.Print(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.")

		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
	}
	ctx.JSON(http.StatusCreated, models.GetUserResponse{
		ID:          data.ID,
		FirstName:   data.FirstName,
		LastName:    data.FirstName,
		PhoneNumber: &data.PhoneNumber,
		Email:       data.Email,
		CreatedAt:   data.CreatedAt,
	})
}

// @Summary Update a user
// @Description Update a userss
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param user body models.CreateUserRequest true "user"
// @Success 200 {object} models.GetUserResponse
// @Failure 500 {object} models.ResponseError
// @Router /users/{id} [put]
func (hand *handlerv1) UpdateUser(ctx *gin.Context) {
	var (
		req models.GetUserResponse
	)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massege": err.Error(),
		})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	req.ID = int64(id)
	user, err := hand.storage.User().Update(&repo.User{
		ID:          req.ID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: *req.PhoneNumber,
		Email:       req.Email,
		ImageUrl:    *req.ImageUrl,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// @Summary Delete a User
// @Description Delete a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ResponseError
// @Router /users/{id} [delete]
func (hand *handlerv1) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseErro(err))
		return
	}
	err = hand.storage.User().Delete(int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseErro(err))
		return
	}
	ctx.JSON(http.StatusOK, models.ResponseOK{
		Message: "Succesfully deleted!",
	})
}

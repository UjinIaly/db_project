package handlers

import (
	"db_project/internal/models"
	"db_project/internal/usecases"
	"db_project/pkg/errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/mailru/easyjson"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	PostURL     string
	PostUseCase usecases.PostUseCase
}

func CreatePostHandler(router *gin.RouterGroup, postURL string, postUseCase usecases.PostUseCase) {
	handler := &PostHandler{
		PostURL:     postURL,
		PostUseCase: postUseCase,
	}

	posts := router.Group(handler.PostURL)
	{
		posts.GET("/:id/details", handler.GetPost)
		posts.POST("/:id/details", handler.UpdatePost)
	}
}

func (postHandler *PostHandler) GetPost(c *gin.Context) {
	postIDstr := c.Param("id")
	postID, err := strconv.Atoi(postIDstr)

	relatedData := c.Query("related")
	var relatedDataArr []string
	if relatedData != "" {
		relatedDataArr = strings.Split(relatedData, ",")
	}

	postFull, err := postHandler.PostUseCase.Get(int64(postID), &relatedDataArr)
	if err != nil {
		c.Data(errors.PrepareErrorResponse(err))
		return
	}

	postFullJSON, err := postFull.MarshalJSON()
	if err != nil {
		c.Data(errors.PrepareErrorResponse(err))
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", postFullJSON)
}

func (postHandler *PostHandler) UpdatePost(c *gin.Context) {
	postIDstr := c.Param("id")
	postID, err := strconv.Atoi(postIDstr)
	if err != nil {
		c.Data(errors.PrepareErrorResponse(err))
		return
	}

	postUpdate := new(models.PostUpdate)
	if err := easyjson.UnmarshalFromReader(c.Request.Body, postUpdate); err != nil {
		c.Data(errors.PrepareErrorResponse(errors.ErrBadRequest))
		return
	}

	post := &models.Post{
		ID:      int64(postID),
		Message: postUpdate.Message,
	}
	err = postHandler.PostUseCase.Update(post)
	if err != nil {
		c.Data(errors.PrepareErrorResponse(err))
		return
	}

	postJSON, err := post.MarshalJSON()
	if err != nil {
		c.Data(errors.PrepareErrorResponse(err))
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", postJSON)
}

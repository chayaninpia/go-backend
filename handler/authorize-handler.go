package handler

import (
	"apps/models"
	"apps/pkg/utils/logx"
	"apps/pkg/utils/middleware"
	"apps/pkg/utils/resx"

	"github.com/gin-gonic/gin"
)

type IAuthorizeHandler interface {
	Login(c *gin.Context)
}

type AuthorizeHandler struct {
}

func NewAuthorizeHandler() *AuthorizeHandler {
	return &AuthorizeHandler{}
}

// Login godoc
// @Summary      Login
// @Description  to get authorization token
// @Tags         Login
// @Id			 Login
// @Success      200  {object}  models.ResponseAPI[models.Token]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/login [post]
func (handler AuthorizeHandler) Login(c *gin.Context) {

	token, err := middleware.GenerateJWT()
	if err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, "", models.Token{Token: token})
}

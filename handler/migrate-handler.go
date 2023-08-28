package handler

import (
	"apps/pkg/migrate"
	"apps/pkg/utils/logx"
	"apps/pkg/utils/resx"

	_ "apps/models"

	"github.com/gin-gonic/gin"
)

type IMigrateHandler interface {
	AutoMigrate(c *gin.Context)
}

type MigrateHandler struct {
	MigrateRepository migrate.IMigrateRepository
}

func NewMigrateHandler() *MigrateHandler {
	return &MigrateHandler{
		MigrateRepository: migrate.NewMigrateRepository(),
	}
}

// AutoMigrate godoc
// @Summary      AutoMigrate
// @Description  to automigate database
// @Tags         AutoMigrate
// @Id			 AutoMigrate
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/automigrate [post]
func (handler MigrateHandler) AutoMigrate(c *gin.Context) {

	if err := handler.MigrateRepository.AutoMigrate(); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Auto Migrate Complete`, nil)
}

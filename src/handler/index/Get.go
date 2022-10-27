package index

import (
	"apps/src/util/logx"
	"apps/src/util/resx"
	"apps/src/util/validx"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	req := IndexI{}
	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	if err := validx.Struct(req); err != nil {
		logx.Error(c, err.Error())
	}

	host, _ := os.Hostname()
	log.Println("host : ", host, "\nPID :", os.Getpid())

	res := IndexO{
		Name:     req.Name,
		LastName: strings.Join(req.LastName, ``),
	}

	resx.Success(c, ``, res)
}

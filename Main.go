package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mhsyjwt/gongjus"
	"net/http"
	"os"
	"time"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

type Fanhui struct {
	Zhuangtai string
	Zhi       string
}

func main() {
	gin.DisableConsoleColor()
	f, _ := os.Create("gin" + gongjus.Time2string(time.Now(), gongjus.NYRSFMXHX) + ".log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	r.Use(Cors())
	sn := r.Group("sn")
	{
		sn.GET("/jieshou", func(c *gin.Context) {
			log.Println(gongjus.Time2stringnyrsfm(time.Now()))
			fh := Fanhui{}
			lpstr := "clztest"
			fh.Zhi = lpstr
			fh.Zhuangtai = "chenggong"
			c.JSON(http.StatusOK, fh)
			return
		})
	}
	r.Run(":8989")
}

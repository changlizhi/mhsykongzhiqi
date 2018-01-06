package main

import (
	"encoding/json"
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
	lingpai := r.Group("jwtlingpai")
	{
		lingpai.GET("/shengcheng/:xinxis", func(c *gin.Context) {
			fh := Fanhui{}
			xinxis := c.Param("xinxis")
			jx := &gongjus.Xinxis{}
			err := json.Unmarshal([]byte(xinxis), jx)
			if err != nil {
				fh.Zhi = err.Error()
				fh.Zhuangtai = "jiexishibai"
				c.JSON(http.StatusOK, fh)
				return
			}
			lpstr, err := gongjus.Shengchenglingpai(jx)
			if err != nil {
				fh.Zhi = err.Error()
				fh.Zhuangtai = "shengchengshibai"
				c.JSON(http.StatusOK, fh)
				return
			}
			fh.Zhi = lpstr
			fh.Zhuangtai = "chenggong"
			c.JSON(http.StatusOK, fh)
			return
		})
		lingpai.GET("/yanzheng/:xinxis", func(c *gin.Context) {
			xinxis := c.Param("xinxis")
			log.Println("xinxis-----", xinxis)
			lingpai := c.GetHeader("Authorization")
			fh := Fanhui{}
			clis, err := gongjus.Yanzhenglingpai(lingpai)
			if err != nil {
				fh.Zhi = err.Error()
				fh.Zhuangtai = "yanzhengshibai"
				c.JSON(http.StatusOK, fh)
				return
			}
			log.Println("clis-----", clis)
			fh.Zhi = "/zhanghao1s/yemian"
			fh.Zhuangtai = "chenggong"
			c.JSON(http.StatusOK, fh)
			return
		})
	}
	r.Run(":8989")
}

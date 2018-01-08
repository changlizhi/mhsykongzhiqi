package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"mhsykongzhiqi/kus"
	"mhsykongzhiqi/moxings"
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
	//f, _ := os.Create("gin" + gongjus.Time2string(time.Now(), gongjus.NYRSFMXHX) + ".log")
	//gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	r.Use(Cors())
	sn := r.Group("sn")
	{
		sn.POST("/jieshou", func(c *gin.Context) {
			xlh := c.PostForm("Xuliehao")
			macdizhi := c.PostForm("Macdizhi")
			sb := &moxings.Shebeis{
				Xuliehao:xlh,
				Macdizhi:macdizhi,
				Pici:1,
			}
			cg:=kus.Charushebei(sb)
			if(cg){
				c.String(http.StatusOK, "\nconfig fuwuxinxi\n\toption yijieshou '1'\n\n")
				return
			}
			c.String(http.StatusOK, "\nconfig fuwuxinxi\n\toption yijieshou '0'\n\n")
			return
		})
		sn.POST("/ceshilianwang", func(c *gin.Context) {
			ceshi := c.PostForm("Ceshilianwang")
			if(ceshi=="ceshi"){
				c.String(http.StatusOK, "\nconfig lianwang\n\toption yilianjie '1'\n\n")
				return
			}
			c.String(http.StatusOK, "\nconfig lianwang\n\toption yilianjie '0'\n\n")
			return
		})
	}
	r.Run(":8989")
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"mhsykongzhiqi/kus"
	"mhsykongzhiqi/moxings"
	"strconv"
	"github.com/lunny/log"
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
			cg := kus.Charushebei(sb)
			if (cg) {
				opt := "\nconfig fuwuxinxi\n\t" +
					"option yijieshou '1'"
				ypsjcx := moxings.Yinpinshijians{
					Xuliehao:xlh,
				}
				ypsj := kus.Chaxunyigeyinpinshijian(ypsjcx)
				if (ypsj.Dangqianshijian - ypsj.Jieyashijian > 4665600000) {
					opt = opt + "\n\toption xushanchu '1'"
				} else {
					opt = opt + "\n\toption xushanchu '0'"
				}
				opt = opt + "\n\n"
				//增加是否需要删除音频的标记
				c.String(http.StatusOK, opt)
				return
			}
			c.String(http.StatusOK, "\nconfig fuwuxinxi\n\toption yijieshou '0'\n\n")
			return
		})
		sn.POST("/ceshilianwang", func(c *gin.Context) {
			ceshi := c.PostForm("Ceshilianwang")
			if (ceshi == "ceshi") {
				c.String(http.StatusOK, "\nconfig lianwang\n\toption yilianjie '1'\n\n")
				return
			}
			c.String(http.StatusOK, "\nconfig lianwang\n\toption yilianjie '0'\n\n")
			return
		})
		sn.POST("/bofangshijian", func(c *gin.Context) {
			xlh := c.PostForm("Xuliehao")
			kssj := c.PostForm("Kaishishijian")
			jssj := c.PostForm("Jieshushijian")
			kssjint, _ := strconv.ParseInt(kssj, 10, 64)
			jssjint, _ := strconv.ParseInt(jssj, 10, 64)

			mx := &moxings.Bofangshijians{
				Xuliehao:xlh,
				Kaishishijian:kssjint,
				Jieshushijian:jssjint,
			}
			kus.Charubofangshijian(mx)

			c.String(http.StatusOK, "\nconfig lianwang\n\toption yilianjie '0'\n\n")
			return
		})
		sn.POST("/yinpinshijian", func(c *gin.Context) {
			xlh := c.PostForm("Xuliehao")
			jysj := c.PostForm("Jieyashijian")
			dqsj := c.PostForm("Dangqianshijian")
			jysjint, _ := strconv.ParseInt(jysj, 10, 64)
			dqsjint, _ := strconv.ParseInt(dqsj, 10, 64)

			mx := &moxings.Yinpinshijians{
				Xuliehao:xlh,
				Jieyashijian:jysjint,
				Dangqianshijian:dqsjint,
			}
			kus.Charuyinpinshijian(mx)

			c.String(http.StatusOK, "\nconfig lianwang\n\toption yilianjie '0'\n\n")
			return
		})
	}
	r.Run(":8989")
}

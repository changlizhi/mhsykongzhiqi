package main

import (
	"github.com/gin-gonic/gin"
	"mhsykongzhiqi/kus"
	"mhsykongzhiqi/moxings"
	"net/http"
	"strconv"
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
	r.Static("/jingtais", "./jingtais")
	r.Use(Cors())
	sn := r.Group("sn")
	{
		sn.POST("/jieshou", func(c *gin.Context) {
			xlh := c.PostForm("Xuliehao")
			macdizhi := c.PostForm("Macdizhi")
			sb := &moxings.Shebeis{
				Xuliehao: xlh,
				Macdizhi: macdizhi,
				Pici:     1,
			}
			cg := kus.Charushebei(sb)
			if cg {
				opt := "\nconfig fuwuxinxi\n\t" +
					"option yijieshou '1'"
				ypsjcx := moxings.Yinpinshijians{
					Xuliehao: xlh,
				}
				ypsj := kus.Chaxunyigeyinpinshijian(ypsjcx)
				if ypsj.Dangqianshijian-ypsj.Jieyashijian > 4665600000 {
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
			if ceshi == "ceshi" {
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
				Xuliehao:      xlh,
				Kaishishijian: kssjint,
				Jieshushijian: jssjint,
			}
			kus.Charubofangshijian(mx)

			c.String(http.StatusOK, "\nconfig quxian\n\toption zaozhongwan '111'\n\n") //在控制器中记录服务器中的判断，如播放过短，过长等信息
			return
		})
		sn.POST("/yinpinshijian", func(c *gin.Context) {
			xlh := c.PostForm("Xuliehao")
			jysj := c.PostForm("Jieyashijian")
			dqsj := c.PostForm("Dangqianshijian")
			jysjint, _ := strconv.ParseInt(jysj, 10, 64)
			dqsjint, _ := strconv.ParseInt(dqsj, 10, 64)

			mx := &moxings.Yinpinshijians{
				Xuliehao:        xlh,
				Jieyashijian:    jysjint,
				Dangqianshijian: dqsjint,
			}
			kus.Charuyinpinshijian(mx)

			c.String(http.StatusOK, "\nconfig yinpin\n\toption dangqianshijian '"+dqsj+"'\n\n") //把服务器获取到的当前时间发下去
			return
		})
		sn.POST("/yinpinlianjie", func(c *gin.Context) { //根据控制器编码下发音频链接，分配不同的音频文件
			xlh := c.PostForm("Xuliehao")
			c.String(http.StatusOK, "\nconfig lianjie\n\toption yinpin 'base64'\n\n"+xlh)
			return
		})
		sn.POST("/yinpinmima", func(c *gin.Context) { //根据不同的控制器下发不同的音频密码
			xlh := c.PostForm("Xuliehao")
			c.String(http.StatusOK, "\nconfig yinpinmima\n\toption mima '0'\n\n"+xlh)
			return
		})
		sn.POST("/xitongmima", func(c *gin.Context) { //根据不同的控制器下发不同的系统密码
			xlh := c.PostForm("Xuliehao")
			c.String(http.StatusOK, "\nconfig xitongmima\n\toption mima '0'\n\n"+xlh)
			return
		})
		sn.POST("/smarthomegengxin", func(c *gin.Context) { //根据不同的控制器发送smarthome.out是否需要更新的标记，如果需要则开机时下载smarthome.out
			xlh := c.PostForm("Xuliehao")
			c.String(http.StatusOK, "\nconfig xitong\n\toption gengxin '1'\n\n"+xlh)
			return
		})
		// 哪些需要跟服务器交互的？
		// 当联网成功后，马上上传sn，马上进行下列参数的获取
		// 控制器的播放时间，第一个间距可能会非常大，所以下发删除命令时需要慎重
		// 1.是否需要更新软件，---提供询问服务和软件链接服务
		// 2.是否需要更新音频，---提供询问服务和音频地址服务
		// 3.是否需要更新音频密码，---提供询问服务和音频密码服务
		// 4.是否需要更新音频链接，---提供询问服务和音频链下发接服务
		// 5.是否需要更新系统密码，---提供询问服务和系统密码下发服务
		// 6.是否需要失效音频，---提供询问服务和失效标记服务
		// 7.是否需要更新音频失效时间，让控制器脱离网络环境时自动失效音频---提供询问服务和时间长度下发服务
		// 8.接收系统播放时长---提供接收时长的服务，用以分析客户状态与穿戴时长的关系
		// 9.接收音频解压时间和系统当前时间---提供接收时长的服务，用以判断是否需要使音频失效的判断标准
		// 10.唯一需要c里做的就是测试联网，其他的全部都放在shell里做，只要测试联网通过了就发起所有的下载，包括上传sn

		// 重点要注意控制器中执行的顺序，要把所有的该上传该下载的都下载了才进行系统下发信息命令的option下发，
		// 所有的事情应该在半小时内做完，不再干预系统供音频播放功能
		// 启动逻辑：把所有需要执行的脚本置为可执行
		// 执行逻辑：
		// 0.启动时做一个整体计时器，计时器时间内都可以进行下列操作，计时结束则关闭这些上传下载，收集需要上传的信息。
		// 1.监测是否已经联网
		// 2.上传sn序列号、音频下载解压时间、音频播放开始结束时间、
		// 当前音频链接及端口号、当前系统密码（密文）、当前音频密码（密文）、
		// 当前软件版本信息、当前音频是否失效、当前音频失效时长设置值、
		// 当前系统软件链接及端口号、当前音频是否存在、当前smarthome.out是否存在、当前所有的脚本是否都存在、单个音频时长。
		// 3.下载更新软件的标记及链接、下载新的音频的标记及链接、下载音频密码更新的标记及链接、
		// 下载音频链接更新的标记、下载系统密码更新的标记、下载当前版本信息、下载新的失效标记、
		// 下载新的失效时长设置、音频时长
		// 4.根据第三步的标记进行实际逻辑的编写，如果软件需要更新则下载新的软件zip，
		// 如果音频需要更新（服务端判断收钱后进行确定此标记）则下载新的音频，此时旧音频可能并没有失效，但还是给下载新的音频
		// 如果超过设定的N个月时间则自动删除音频
		// 5.与4同时播放音频，每个音频时长重新检测和播放一次，检测不到音频则提示。所以所有的控制器逻辑是需要更改的。
		// 分三个服务，第一个接收数据，第二个下发标记，第三个真正下载

		// 需要更新软件的条件：增加一个初始时全部上传的操作，增加一个使用过程中的提示上传操作。
		// 感觉这些检测在shell里全部都能做，包括是否联网。c只需要对硬件进行操作即可。
		// 比如增大减小暂停音量，系统初始化：先删除音频，再删除overlay，增加一个mtd的操作和jffs2reset的。
		// 1.系统下发标记，2.软件不存在
	}
	r.Run(":8989")
}

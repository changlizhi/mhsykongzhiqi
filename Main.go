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
		sn.POST("/ceshilianwang", func(c *gin.Context) {
			ceshi := c.PostForm("Ceshilianwang")
			if ceshi == "ceshi" {
				c.String(http.StatusOK, "\nconfig lianwang\n\toption yilianjie '1'\n\n")
				return
			}
			c.String(http.StatusOK, "\nconfig lianwang\n\toption yilianjie '0'\n\n")
			return
		})
		sn.POST("/jssn", func(c *gin.Context) {
			xlh := c.PostForm("Xuliehao")
			macdizhi := c.PostForm("Macdizhi")
			sb := &moxings.Shebeis{
				Xuliehao: xlh,
				Macdizhi: macdizhi,
				Pici:     1,
			}
			kusb := kus.Chaxunyigeshebei(*sb)
			if kusb == nil {
				cg := kus.Charushebei(sb)
				if cg {
					//返回标记接收并入库成功
					//增加是否需要删除音频的标记
					c.String(http.StatusOK, "\nconfig jssn\n\toption scchenggong '1'\n\n")
					return
				}
				c.String(http.StatusOK, "\nconfig jssn\n\toption scchenggong '0'\n\n")
				return
			}
			c.String(http.StatusOK, "\nconfig jssn\n\toption scchenggong '1'\n\n")
			return
		})
		sn.POST("/jsyinpinxiazai", func(c *gin.Context) {
			//接收音频下载时间
			xlh := c.PostForm("Xuliehao")
			xzsj := c.PostForm("Xiazaishijian")
			dqsj := c.PostForm("Dangqianshijian")
			xzsjint, _ := strconv.ParseInt(xzsj, 10, 64)
			dqsjint, _ := strconv.ParseInt(dqsj, 10, 64)
			sb := &moxings.Yinpinxiazais{
				Xuliehao:        xlh,
				Xiazaishijian:   xzsjint,
				Dangqianshijian: dqsjint,
			}
			// 每次有上传都要入库，用统计的方式计算是否需要更新音频，
			// 因为控制器的时钟芯片必须联网才能更新时间，而且不一定能更新为互联网时间
			cg := kus.Charuyinpinxiazai(sb)
			if cg {
				//返回标记接收并入库成功
				c.String(http.StatusOK, "\nconfig jsyinpinxiazai\n\toption scchenggong '1'\n\n")
				return
			}
			c.String(http.StatusOK, "\nconfig jsyinpinxiazai\n\toption scchenggong '0'\n\n")
			return
		})
		sn.POST("/jsyinpinbofang", func(c *gin.Context) {
			//接收音频下载时间
			xlh := c.PostForm("Xuliehao")
			kssj := c.PostForm("Kaishishijian")
			jssj := c.PostForm("Jieshushijian")
			kssjint, _ := strconv.ParseInt(kssj, 10, 64)
			jssjint, _ := strconv.ParseInt(jssj, 10, 64)
			sb := &moxings.Yinpinbofangs{
				Xuliehao:      xlh,
				Kaishishijian: kssjint,
				Jieshushijian: jssjint,
			}
			// 每次有上传都要入库，用统计的方式计算是否需要更新音频，
			// 因为控制器的时钟芯片必须联网才能更新时间，而且不一定能更新为互联网时间
			cg := kus.Charuyinpinbofangs(sb)
			if cg {
				//返回标记接收并入库成功
				c.String(http.StatusOK, "\nconfig jsyinpinbofang\n\toption scchenggong '1'\n\n")
				return
			}
			c.String(http.StatusOK, "\nconfig jsyinpinbofang\n\toption scchenggong '0'\n\n")
			return
		})
		sn.POST("/jsyinpinlianjiejiu", func(c *gin.Context) {
			//当前音频链接需要更新，其实是当作比较，当不需要更新的时候新链接旧链接是一样的，用最新的一条去比较，id最大的
			xlh := c.PostForm("Xuliehao")
			ip := c.PostForm("Ipdizhi")
			dk := c.PostForm("Duankou")
			lj := c.PostForm("Lianjie")
			sb := &moxings.Yinpinlianjiejius{
				Xuliehao: xlh,
				Ipdizhi:  ip,
				Duankou:  dk,
				Lianjie:  lj,
			}
			yplj := kus.Chaxunyigeyinpinlianjiejiu(*sb)
			if yplj == nil {
				cg := kus.Charuyinpinlianjiejiu(sb)
				if cg {
					//返回标记接收并入库成功
					c.String(http.StatusOK, "\nconfig jsyinpinlianjiejiu\n\toption scchenggong '1'\n\n")
					return
				}
				c.String(http.StatusOK, "\nconfig jsyinpinlianjiejiu\n\toption scchenggong '0'\n\n")
				return
			}
			c.String(http.StatusOK, "\nconfig jsyinpinlianjiejiu\n\toption scchenggong '1'\n\n")
			return
		})
		sn.POST("/jsyinpinmima", func(c *gin.Context) {
			//当前音频链接需要更新，其实是当作比较，当不需要更新的时候新链接旧链接是一样的，用最新的一条去比较，id最大的
			xlh := c.PostForm("Xuliehao")
			ypmima := c.PostForm("Yinpinmima")
			sb := &moxings.Yinpinmimajius{
				Xuliehao:   xlh,
				Yinpinmima: ypmima,
			}
			yplj := kus.Chaxunyigeyinpinmimajiu(*sb)
			if yplj == nil {
				cg := kus.Charuyinpinmimajiu(sb)
				if cg {
					//返回标记接收并入库成功
					c.String(http.StatusOK, "\nconfig jsyinpinmima\n\toption scchenggong '1'\n\n")
					return
				}
				c.String(http.StatusOK, "\nconfig jsyinpinmima\n\toption scchenggong '0'\n\n")
				return
			}
			c.String(http.StatusOK, "\nconfig jsyinpinmima\n\toption scchenggong '1'\n\n")
			return
		})
		sn.POST("/jsyinpinshanchu", func(c *gin.Context) {
			//
			xlh := c.PostForm("Xuliehao")
			scbj := c.PostForm("Shanchubiaoji")
			ysc := c.PostForm("Yishanchu")
			scbjint ,_:=strconv.ParseInt(scbj,10,64)
			yscint ,_:=strconv.ParseInt(ysc,10,64)

			sb := &moxings.Yinpinshanchujius{
				Xuliehao:      xlh,
				Shanchubiaoji: scbjint,
				Yishanchu:     yscint,
			}
			cg := kus.Charuyinpinshanchujiu(sb)
			if cg {
				//返回标记接收并入库成功
				c.String(http.StatusOK, "\nconfig jsyinpinshanchu\n\toption scchenggong '1'\n\n")
				return
			}
			c.String(http.StatusOK, "\nconfig jsyinpinshanchu\n\toption scchenggong '0'\n\n")
			return
		})
		sn.POST("/jsyinpinshixiaoshijian", func(c *gin.Context) {
			//
			xlh := c.PostForm("Xuliehao")
			sc := c.PostForm("Shichang") //秒数
			scint,_ := strconv.ParseInt(sc,10,64)//秒数
			sb := &moxings.Yinpinshixiaoshijianjius{
				Xuliehao: xlh,
				Shichang: scint,
			}
			cg := kus.Charuyinpinshixiaoshijianjiu(sb)
			if cg {
				//返回标记接收并入库成功
				c.String(http.StatusOK, "\nconfig jsyinpinshixiaoshijian\n\toption scchenggong '1'\n\n")
				return
			}
			c.String(http.StatusOK, "\nconfig jsyinpinshixiaoshijian\n\toption scchenggong '0'\n\n")
			return
		})
		sn.POST("/jsyinpindaxiao", func(c *gin.Context) {
			//
			xlh := c.PostForm("Xuliehao")
			dx := c.PostForm("Daxiao")        //大小
			ysdx := c.PostForm("Yasuodaxiao") //压缩大小
			dxint,_ := strconv.Atoi(dx)
			ysdxint,_ := strconv.Atoi(ysdx)

			sb := &moxings.Yinpindaxiaojius{
				Xuliehao:    xlh,
				Daxiao:      dxint,
				Yasuodaxiao: ysdxint,
			}
			cg := kus.Charuyinpindaxiaojiu(sb)
			if cg {
				//返回标记接收并入库成功
				c.String(http.StatusOK, "\nconfig jsyinpindaxiao\n\toption scchenggong '1'\n\n")
				return
			}
			c.String(http.StatusOK, "\nconfig jsyinpindaxiao\n\toption scchenggong '0'\n\n")
			return
		})
		sn.POST("/jsruanjianjiu", func(c *gin.Context) {
			//当前音频链接需要更新，其实是当作比较，当不需要更新的时候新链接旧链接是一样的，用最新的一条去比较，id最大的
			xlh := c.PostForm("Xuliehao")
			bb := c.PostForm("Banben")
			wz := c.PostForm("Weizhi")
			nr := c.PostForm("Neirong")
			lx := c.PostForm("Leixing")
			sb := &moxings.Ruanjianjius{
				Xuliehao: xlh,
				Banben:   bb,
				Weizhi:   wz,
				Neirong:  nr,
				Leixing:  lx,
			}
			cg := kus.Charuruanjianjiu(sb)
			if cg {
				//返回标记接收并入库成功bv
				c.String(http.StatusOK, "\nconfig jsruanjianjiu\n\toption scchenggong '1'\n\n")
				return
			}
			c.String(http.StatusOK, "\nconfig jsruanjianjiu\n\toption scchenggong '0'\n\n")
			return
		})
		// 哪些需要跟服务器交互的？c中要搜集控制器的执行状态，什么情况下正常执行。什么情况下要替换版本。

		// 重点要注意控制器中执行的顺序，要把所有的该上传该下载的都下载了才进行系统下发信息命令的option下发，
		// 所有的事情应该在半小时内做完，不再干预系统供音频播放功能
		// 启动逻辑：把所有需要执行的脚本置为可执行
		// 执行逻辑：
		// 0.启动时做一个整体计时器，计时器时间内都可以进行下列操作，计时结束则关闭这些上传下载，收集需要上传的信息。
		// 1.监测是否已经联网
		// 2.上传sn序列号、音频下载解压时间、音频播放开始结束时间、
		// 当前音频链接及端口号、当前系统密码（密文）不上传需要实时更新且无法被服务器感知是否更改且每次都要询问服务器此密码、当前音频密码（密文）、
		// 当前音频失效标记及是否已删除、当前音频失效时长设置值、单个音频大小及时长压缩后的大小
		// 当前软件版本信息、当前系统软件链接及端口号、当前smarthome.out是否存在、当前所有的脚本是否都存在。

		// 3.下载更新软件的标记及链接、下载新的音频的标记及链接、下载音频密码更新的标记及链接、
		// 下载音频链接更新的标记、下载系统密码更新的标记、下载当前版本信息、下载新的失效标记、
		// 下载新的失效时长设置、音频时长

		sn.POST("/xfruanjianlianjie", func(c *gin.Context) {
			//分散上传，统一下发
			//当前音频链接需要更新，其实是当作比较，当不需要更新的时候新链接旧链接是一样的，用最新的一条去比较，id最大的
			xlh := c.PostForm("Xuliehao")
			lx := "uci" //uci,out,sh三选一，现在下发的只是uci的，out和sh文件是要通过软件更新的方式去下发
			moxing := moxings.Ruanjianxins{
				Xuliehao: xlh,
				Leixing:  lx,
			}
			kusb := kus.Suoyouruanjianxin(moxing)
			if kusb != nil {
				kusbzz := *kusb
				allconf := "\nconfig xfruanjianlianjie"
				for _, sb := range kusbzz {
					allconf += "\n\toption " + sb.Mingcheng + "_neirong '" + sb.Neirong + "'"
					allconf += "\n\toption " + sb.Mingcheng + "_banben '" + sb.Banben + "'"
					allconf += "\n\toption " + sb.Mingcheng + "_weizhi '" + sb.Weizhi + "'" // 位置存在则进行位置设置，这个标记也是控制器二次上传时的数值
				}
				allconf += "\n\n"
				c.String(http.StatusOK, allconf)
				return
			}
			c.String(http.StatusOK, "\nconfig xfruanjianlianjie\n\toption scchenggong '0'\n\n")
			return
		})
		// 控制器及时更新服务器端的数据，服务端下发的依据是最新的控制器信息，所以必须在控制器下发之前进行上传。
		// 并且上传成功控制器才能进行下载，默认是上传失败的。启动时把所有的scchenggong置为0，然后进行上传，成功的服务器端会自动下发为1的值，
		// 意思就是每次下发都是在运行时做的。启动全部的线程进行工作。
		// uci set只在启动时整体做一次，其他的都从服务端下发。真正下载解压成功了之后又设置成1

		// 4.根据第三步的标记进行实际逻辑的编写，如果软件需要更新则下载新的软件zip，
		// 如果音频需要更新（服务端判断收钱后进行确定此标记）则下载新的音频，此时旧音频可能并没有失效，但还是给下载新的音频
		// 如果超过设定的N个月时间则自动删除音频
		// 5.与4同时播放音频，每个音频时长重新检测和播放一次，检测不到音频则提示。所以所有的控制器逻辑是需要更改的。
		// 分三个服务，第一个接收数据，第二个下发标记，第三个真正下载

		// 增加一个初始时全部上传的操作就在startup.sh里面，每个上传、下载增加一个使用过程中的提示上传操作的sh。
		// 感觉这些检测在shell里全部都能做，包括是否联网。c只需要对硬件进行操作即可。
		// 比如增大减小暂停音量，系统初始化：先删除音频，再删除overlay，增加一个mtd的操作和jffs2reset的。
		// 需要更新软件的条件：
		// 1.系统下发标记，2.软件不存在

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
	}
	r.Run(":8989")
}

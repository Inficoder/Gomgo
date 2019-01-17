package routers

import (
	"Gomgo/pkg/e"
	"Gomgo/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine{
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.StaticFS("/static",http.Dir("static"))
	r.GET("/", func(c *gin.Context) {
		//c.JSON(http.StatusOK,gin.H{
		//
		//})
		c.String(http.StatusOK,"Hello ~")
	})
	//加载HTML页面
	r.LoadHTMLGlob("static/views/page/*")
	//首页路由注册
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"code": 200,
			"msg":  e.GetMsg(200),
			//"chords":data,
			})
	})
	////详情页路由注册
	r.GET("/chordDetail", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chord.html", gin.H{
			"code": 200,
			"msg":  e.GetMsg(200),
			//"chords":data,
		})
	})


	appv1 := r.Group("/play")
	{
		appv1.GET("/chord/:Id",v1.GetChord)
		//appv1.GET("/chords",v1.GetChordsList)
		appv1.GET("/chords",v1.GetChordsListWithCondition)
		appv1.POST("/chord",v1.AddChord)
		//appv1.GET("/chords",v1.GetChordsList)
		//appv1.GET("/chords",v1.GetChordsList)
	}

	return r
}


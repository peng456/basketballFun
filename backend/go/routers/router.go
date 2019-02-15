package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/EDDYCJY/go-gin-example/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/EDDYCJY/go-gin-example/pkg/export"
	"github.com/EDDYCJY/go-gin-example/pkg/qrcode"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/upload"
	"github.com/EDDYCJY/go-gin-example/routers/api"
	"github.com/EDDYCJY/go-gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	apiv1.Use()
	{
		//获取球场列表
		apiv1.GET("/courts", v1.GetCourts)
		////获取收藏球场列表
		//apiv1.GET("/courts/collection", v1.GetCollectionCourtList)
		//
		////收藏球场
		//apiv1.GET("/courts/collection/:id", v1.CollectionCourt)
		////取消收藏球场
		//apiv1.DELETE("/courts/collection/:id", v1.CancleCollectionCourtList)

		//获取指定球场
		apiv1.GET("/courts/:id", v1.GetCourt)
		//上传球场
		apiv1.POST("/courts", v1.AddCourt)
		//更新指定球场
		apiv1.PUT("/courts/:id", v1.EditCourt)
		//删除指定球场
		apiv1.DELETE("/courts/:id", v1.DeleteCourt)


		////获取用户列表
		//apiv1.GET("/users", v1.GetUserslist)
		////获取指定用户信息
		//apiv1.GET("/users/:id", v1.GetUser)
		////创建用户
		//apiv1.POST("/users", v1.AddUser)
		////更新指定User
		//apiv1.PUT("/users/:id", v1.EditUser)
		////删除指定用户
		//apiv1.DELETE("/users/:id", v1.DeleteUser)
		//
		//
		////获取球赛列表
		//apiv1.GET("/matches", v1.GetMatchlist)
		////获取指定match信息
		//apiv1.GET("/matches/:id", v1.GetMatch)
		////创建match
		//apiv1.POST("/matches", v1.AddMatch)
		////更新指定match
		//apiv1.PUT("/matches/:id", v1.EditMatch)
		////取消指定match
		//apiv1.DELETE("/matches/:id", v1.DeleteMatch)
		//
		//
		////获取参赛人员列表
		//apiv1.GET("/matchmember/:id", v1.GetMatchMember)
		////报名match
		//apiv1.POST("/matchmember/:uid", v1.AddMatchMember)
		////更新报名信息
		//apiv1.PUT("/matchmember/:uid/:matchid", v1.EditMatchMember)
		////删除、取消报名
		//apiv1.DELETE("/matchmember/:uid/:matchid", v1.DeleteMatchMember)

		//生成文章海报
		//apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)

	}

	return r
}

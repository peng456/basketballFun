package v1

import (
	"net/http"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/service/court_service"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)


type AddCourtForm struct {
	Name      string `form:"name" valid:"Required;MaxSize(100)"`
}

type EditCourtForm struct {
	ID         int    `form:"id" valid:"Required;Min(1)"`
	Name       string `form:"name" valid:"Required;MaxSize(100)"`
}

// @Summary 获取单个球场
// @Produce  json
// @Param id param int true "ID"
// @Success 200 {string} json "{"code":200,"data":{"id":3,"created_on":1516937037,"modified_on":0,"tag_id":11,"tag":{"id":11,"created_on":1516851591,"modified_on":0,"name":"312321","created_by":"4555","modified_by":"","state":1},"content":"5555","created_by":"2412","modified_by":"","state":1},"msg":"ok"}"
// @Router /api/v1/courts/{id} [get]
func GetCourt(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}


	courtService := court_service.Court{ID: id}
	exists, err := courtService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_COURT_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_COURT, nil)
		return
	}

	court, err := courtService.Get()

	//logFilename  := "log.txt"
	//
	//isExits, err := existsFile(logFilename)
	//
	//if isExits != true {
	//	//panic("pppp")
	//
	//	_,err := os.Create(logFilename)
	//	if err != nil {
	//		//fmt.Printf("open file error=%s\r\n", err.Error())
	//		panic("open file erro")
	//
	//		os.Exit(-1)
	//	}
	//
	//}
	//
	//logFile, err := os.OpenFile(logFilename, os.O_RDWR | os.O_CREATE, 0777)
	//
	////panic()
	//if err != nil {
	//	//fmt.Printf("open file error=%s\r\n", err.Error())
	//	panic("open file erro")
	//
	//	os.Exit(-1)
	//}
	//
	//defer logFile.Close()
	//logger:=log.New(logFile,"\r\n", log.Ldate | log.Ltime | log.Lshortfile)
	//jsonStu,err := json.Marshal(court)
	//str := string(jsonStu)
	//logger.Println(str)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_COURT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, court)
}



// @Summary 获取多个文章
// @Produce  json
// @Param tag_id query int false "TagID"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":[{"id":3,"created_on":1516937037,"modified_on":0,"tag_id":11,"tag":{"id":11,"created_on":1516851591,"modified_on":0,"name":"312321","created_by":"4555","modified_by":"","state":1},"content":"5555","created_by":"2412","modified_by":"","state":1}],"msg":"ok"}"
// @Router /api/v1/articles [get]
func GetCourts(c *gin.Context) {

	// https://www.cnblogs.com/-beyond/p/9391892.html gin框架学习手册
	// /user/:name/:age/:addr/:sex  ==》 name := c.Param("name")
	// （GET请求）  http://localhost:8080/user/jane/20/beijing/female?id=999&height=170&wight=100
	//  height ==》 height := c.Query("height")  Query(key)，DefaultQuery(key,default)
	//  POST请求中的参数 ==》PostForm(key)，DefaultPostForm(key,default)

	appG := app.Gin{C: c}
	valid := validation.Validation{}

	longitude := 0.0
	latitude := 0.0
	if arg := c.PostForm("longitude"); arg != "" {
		longitude = com.StrTo(arg).MustFloat64()
		valid.Range(longitude, 0.0, 180.0, "longitude").Message("longitude在-180.0~180.0")
	}

	if arg := c.PostForm("latitude"); arg != "" {
		latitude = com.StrTo(arg).MustFloat64()
		valid.Range(latitude, 0.0, 80.0, "latitude").Message("latitude在 -90~90")
	}


	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors)
		return
	}

	courtService := court_service.Court{
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
		Longitude: com.StrTo(c.DefaultQuery("longitude","")).MustFloat64(),
		Latitude: com.StrTo(c.DefaultQuery("latitude","")).MustFloat64(),
		R: com.StrTo(c.DefaultQuery("r","0")).MustFloat64(),
	}



	//total, err := courtService.Count()
	//if err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_ARTICLE_FAIL, nil)
	//	return
	//}

	courts, err := courtService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_COURT_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = courts
	data["total"] = 99999

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

//// exists returns whether the given file or directory exists or not
//func existsFile(path string) (bool, error) {
//	_, err := os.Stat(path)
//	if err == nil { return true, nil }
//	if os.IsNotExist(err) { return false, nil }
//	return true, err
//}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func AddCourt(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddCourtForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	//appG.Response(http.StatusOK, e.SUCCESS, httpCode)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	courtService := court_service.Court{Name: form.Name}


	exists, err := courtService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_TAG, nil)
		return
	}

	err = courtService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func EditCourt(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = EditCourtForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	courtService := court_service.Court{
		ID:			   form.ID,
		Name:         form.Name,

	}
	exists, err := courtService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_COURT_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_COURT, nil)
		return
	}

	err = courtService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_COURT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 删除文章
// @Produce  json
// @Param id param int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 200 {string} json "{"code":400,"data":{},"msg":"请求参数错误"}"
// @Router /api/v1/courts/{id} [delete]
func DeleteCourt(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, valid.Errors)
		return
	}

	courtService := court_service.Court{ID: id}
	exists, err := courtService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_COURT_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_COURT, nil)
		return
	}

	err = courtService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_COURT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

















const (
	QRCODE_URL = "https://github.com/EDDYCJY/blog#gin%E7%B3%BB%E5%88%97%E7%9B%AE%E5%BD%95"
)

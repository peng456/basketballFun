package court_service

import (
	//"encoding/json"

	"github.com/EDDYCJY/go-gin-example/models"

	"time"
	//"github.com/EDDYCJY/go-gin-example/service/cache_service"
	//"github.com/EDDYCJY/go-gin-example/pkg/gredis"
	//"github.com/EDDYCJY/go-gin-example/pkg/logging"
	//"encoding/json"
)

type Court struct {

	ID            int
	Name          string
	Longitude     float64
	Latitude      float64
	R		      float64
	courtype      int
	isverifyed    int
	numhoop       int
	groundtype    int
	description   string
	createor      int
	updatetime    int
	createtime    int

	PageNum  int
	PageSize int
}

func (a *Court) Get() (*models.Court, error) {

	// 先把基本的打通，在把redis 上

	//var cacheCourt *models.Court
	//
	//cache := cache_service.Court{ID: a.ID}
	//key := cache.GetCourtKey()
	//if gredis.Exists(key) {
	//	data, err := gredis.Get(key)
	//	if err != nil {
	//		logging.Info(err)
	//	} else {
	//		json.Unmarshal(data, &cacheArticle)
	//		return cacheArticle, nil
	//	}
	//}

	court, err := models.GetCourt(a.ID)
	if err != nil {
		return nil, err
	}

	//gredis.Set(key, article, 3600)
	return court, nil
}

func (t *Court) Add() error {
	return models.AddCourt(t.Name, t.createtime, t.updatetime)
}

func (a *Court) ExistByID() (bool, error) {
	return models.ExistCourtByID(a.ID)
}

func (t *Court) ExistByName() (bool, error) {
	return models.ExistCourtByName(t.Name)
}



func (t *Court) Edit() error {
	data := make(map[string]interface{})
	data["name"] = t.Name
	data["updatetime"] = time.Now().Unix()


	return models.EditCourt(t.ID, data)
}

func (a *Court) Delete() error {
	return models.DeleteCourt(a.ID)
}

func (a *Court) GetAll() ([]*models.Court, error) {
	var (
		courts []*models.Court
	)

	//cache := cache_service.Article{
	//	TagID: a.TagID,
	//	State: a.State,
	//
	//	PageNum:  a.PageNum,
	//	PageSize: a.PageSize,
	//}
	//key := cache.GetArticlesKey()
	//if gredis.Exists(key) {
	//	data, err := gredis.Get(key)
	//	if err != nil {
	//		logging.Info(err)
	//	} else {
	//		json.Unmarshal(data, &cacheArticles)
	//		return cacheArticles, nil
	//	}
	//}

	courts, err := models.GetCourts(a.PageNum, a.PageSize, a.Longitude,a.Latitude,a.R)
	if err != nil {
		return nil, err
	}

	return courts, nil
}

func (a *Court) Count() (int, error) {
	return models.GetCourtTotal(a.getMaps())
}

func (a *Court) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	//if a.State != -1 {
	//	maps["state"] = a.State
	//}
	//if a.TagID != -1 {
	//	maps["tag_id"] = a.TagID
	//}

	if a.Latitude != -1 {
		maps["latitude"] = a.Latitude
	}

	if a.Longitude != -1 {
		maps["longitude"] = a.Longitude
	}

	return maps
}
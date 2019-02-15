package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

type Court struct {
	Model

	Createtime  int `json:"createtime"`
	Updatetime int `json:"updatetime"`

	Name         string `json:"name"`
	Description          string `json:"description"`
	Courttype     int    `json:"courttype"`
	Groundtype    int    `json:"groundtype"`
	Numhoop       int    `json:"numhoop"`
	Isverifyed    int    `json:"isverifyed"`
	Longitude     float32  `json:"longitude"`
	Latitude      float32  `json:"latitude"`
}

func ExistCourtByID(id int) (bool, error) {

	var court Court
	err := db.Select("id").Where("id = ? ", id).First(&court).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if court.ID > 0 {
		return true, nil
	}

	return false, nil
}

func ExistCourtByName(name string) (bool, error) {
	var court Court
	err := db.Select("id").Where("name = ?  ", name).First(&court).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if court.ID > 0 {
		return true, nil
	}

	return false, nil
}


func GetCourt(id int) (*Court, error) {
	var court Court
	err := db.Where("id = ? ", id ).First(&court).Error
	//str,err := json.Marshal(court)
	//panic(str)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &court, nil
}


func AddCourt(name string, Createtime int, Updatetime int) error {
	court := Court{
		Name:      name,
		Createtime:     Createtime,
		Updatetime: Updatetime,
	}
	if err := db.Create(&court).Error; err != nil {
		return err
	}

	return nil
}


func GetCourts(pageNum int, pageSize int, longitude float64 ,latitude float64 ,r float64) ([]*Court, error) {
	var courts []*Court

	// GORM 指南  http://gorm.io/zh_CN/docs/index.html
	// 计算经纬距离  https://www.jianshu.com/p/ec220111d34b   https://blog.csdn.net/fanblog/article/details/79726306

	var str string

	if longitude != 0 &&  latitude != 0 {

		longitudestr := strconv.FormatFloat(longitude, 'f', 6, 64)
		latitudestr := strconv.FormatFloat(latitude, 'f', 6, 64)
		rstr := strconv.FormatFloat(r, 'f', 6, 64)
		str = " sqrt(((( " + longitudestr + " - longitude) * PI() * 12656 * cos(((" + latitudestr + " + latitude) / 2) * PI() / 180) / 180) *(( " + longitudestr + " - longitude) * PI() * 12656 * cos((("+ latitudestr + " + latitude) / 2) * PI() / 180) / 180)) +(((" + latitudestr + " - latitude) * PI() * 12656 / 180) *(("+ latitudestr + " - latitude) * PI() * 12656 / 180))) < " + rstr
		//str = " sqrt(((("
		//str += string(longitude)
		//str +=  " - longitude) * PI() * 12656 * cos(((" + latitude + " + latitude) / 2) * PI() / 180) / 180) *((" + longitude + " - longitude) * PI() * 12656 * cos(((" + latitude + " + latitude) / 2) * PI() / 180) / 180)) +((("+ latitude + " - latitude) * PI() * 12656 / 180) *((" + latitude + " - latitude) * PI() * 12656 / 180))) < ?"
	} else {
		str = ""

	}

	//result,err := db.Exec(str).Rows()
	db.LogMode(true)
	err := db.Debug().Where(str).Offset(pageNum).Limit(pageSize).Find(&courts).Error
	//err := db.Where(str).Offset(pageNum).Limit(pageSize).Find(&courts).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return courts, nil
}

func GetCourtTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Court{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

//
//func GetArticle(id int) (*Article, error) {
//	var article Article
//	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Related(&article.Tag).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return nil, err
//	}
//
//	return &article, nil
//}

func EditCourt(id int, data interface{}) error {
	if err := db.Model(&Court{}).Where("id = ? ", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
//
//func AddArticle(data map[string]interface{}) error {
//	article := Article{
//		TagID:         data["tag_id"].(int),
//		Title:         data["title"].(string),
//		Desc:          data["desc"].(string),
//		Content:       data["content"].(string),
//		CreatedBy:     data["created_by"].(string),
//		State:         data["state"].(int),
//		CoverImageUrl: data["cover_image_url"].(string),
//	}
//	if err := db.Create(&article).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
func DeleteCourt(id int) error {
	if err := db.Where("id = ?", id).Delete(Court{}).Error; err != nil {
		return err
	}

	return nil
}
//
//func CleanAllArticle() error {
//	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{}).Error; err != nil {
//		return err
//	}
//
//	return nil
//}

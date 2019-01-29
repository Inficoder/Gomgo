package v1

import (
	"Gomgo/middleware/logging"
	"Gomgo/models"
	"Gomgo/pkg/e"
	"Gomgo/pkg/setting"
	"Gomgo/pkg/util"
	"fmt"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetChord(c *gin.Context) {
	Id := com.StrTo(c.Param("Id")).String()

	data := make(map[string]interface{})
	valid := validation.Validation{}
	valid.Required(Id, "Id").Message("Id不为空")

	code := e.SUCCESS
	data["exist"] = "true"
	if !valid.HasErrors() {
		if !models.IsExistChordById(Id){
			data["exist"] = "false"
		} else {
			model := models.GetChord(Id)
			data["chord"] = model
			//fmt.Println(model.Url)
			data["imgUrls"] = util.GetChordImages(model.Url)
		}

	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetChordsList(c *gin.Context) {
	pageNum := com.StrTo(c.Query("page")).MustInt()

	data := make(map[string]interface{})

	valid := validation.Validation{}

	code := e.SUCCESS
	data["exist"] = "true"

	if !valid.HasErrors() {
		list := models.GetChordsList(pageNum,setting.PageSize)

		fmt.Println(list)
		if len(list)==0{
			data["exist"] = "false"
		} else {
			data["chords"] = list
			data["total"] = models.GetChordsCount()
		}

	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetChordsListWithCondition(c *gin.Context){
	pageNum := com.StrTo(c.Query("page")).MustInt()
	kind := c.Query("kind")
	key := c.Query("key")
	data := make(map[string]interface{})
	code := e.SUCCESS
	data["exist"] = "true"
	list := models.GetChordsListWithCondition(pageNum,setting.PageSize,kind,key)
	data["pageCalculate"]= models.GetChordsCountByCondition(kind,key)
	fmt.Println(list)
	if len(list)==0{
		data["exist"] = "false"
	} else {
		data["chords"] = list
		data["total"] = models.GetChordsCount()
	}
	data["currentPageCount"] = len(list)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//func InsertChord(name string,singer string,url string, kind string,created_by string,modified_by string,modified_on int,cover_url string) (err error) {
//	con := GetDataBase().C("chord")
//	err = con.Insert(&Chord{ID: bson.NewObjectId(), Name: name, Url: url,Kind:kind,CreatedBy:modified_by,ModifiedOn:modified_on,CoverUrl:cover_url})
//	return
//}

func AddChord(c *gin.Context){
	name := c.Query("name")
	singer := c.Query("singer")
	//url := c.Query("guitar_name")
	kind := c.Query("kind")
	key := c.Query("key")
	count := com.StrTo(c.Query("count")).MustInt()
	locateWay := com.StrTo(c.Query("locate_way")).MustInt()
	//createdBy:= c.Query("guitar_name")

	valid := validation.Validation{}
	valid.Required(name,"name").Message("名称不能为空")
	valid.MaxSize(name,20,"name").Message("名称最长20")
	valid.Required(singer,"singer").Message("singer不能为空")
	valid.Required(kind,"kind").Message("kind不能为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors(){
		if !models.IsExistChordByName(name){
			code = e.SUCCESS
			models.InsertChord(name,singer,locateWay,kind,key,count)
		}else{
			code = e.ERROR_EXIST_CHORD
		}
	}else{
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}





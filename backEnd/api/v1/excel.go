package v1
//增加 删除 修改 查询 crud
import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"oyccx/model"
)
var map1 = make(map[int]string)
var map2 = make(map[string]int)
var map3 = make(map[string]int)
var map4 = make(map[string]int)

//上传文件接口
func ToLead1(c *gin.Context) {
	//获取上传文件
	_, headers, err := c.Request.FormFile("file")
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
	}
	c.SaveUploadedFile(headers, "cache/"+headers.Filename)
	f, err := excelize.OpenFile("cache/" + headers.Filename)

	model.DealWithExcel(f)

	if err != nil {
		fmt.Println(err)
		return
	}



	map1, map2, map3 = model.ExcelToMap1(f)
	map4 = model.ExcelToMap2(f, map2 ,map3)
	model.CreateStruct(map1, map4)

	f.Save()
}
//获取多层嵌套json
func GetJSON(c *gin.Context) {
	companyName:= c.Query("companyName")
	datas := model.MultipleTreeToJSON(companyName)
	method := c.Request.Method
	//跨域问题
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")

	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
	c.JSON(http.StatusOK, gin.H{
		"data": datas,
	})
}
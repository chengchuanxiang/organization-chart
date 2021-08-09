package model

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Data struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255) " json:"name"`
	ParentID int   `gorm:"type:int" json:"parent_id"`
	CurrID int   `gorm:"type:int" json:"curr_id"`
	CompanyName string `gorm:"type:varchar(255) " json:"company_name"`
}








//返回指定公司名的数据
func GetCompany1(companyName string) []Data{
	var data []Data
	_ = Db.Where("name = ?", companyName).Find(&data)
	return data
}


func DealWithExcel1(f *excelize.File) {
	rows, _ := f.GetRows("Sheet1")
	var k int = 1
	for _, _ = range rows {

		cellVal1, _ := f.GetCellValue("Sheet1", "A"+strconv.Itoa(k))
		cellVal2, _ := f.GetCellValue("Sheet1", "B"+strconv.Itoa(k))
		cellVal3, _ := f.GetCellValue("Sheet1", "C"+strconv.Itoa(k))
		cellVal4, _ := f.GetCellValue("Sheet1", "D"+strconv.Itoa(k))
		cellVal5, _ := f.GetCellValue("Sheet1", "E"+strconv.Itoa(k))
		cellVal6, _ := f.GetCellValue("Sheet1", "F"+strconv.Itoa(k))

		if cellVal1 == "" {
			f.SetCellValue("Sheet1", "A"+strconv.Itoa(k), "1")
		}
		if cellVal2 == "" {
			f.SetCellValue("Sheet1", "B"+strconv.Itoa(k), "1")
		}
		if cellVal3 == "" {
			f.SetCellValue("Sheet1", "C"+strconv.Itoa(k), "1")
		}
		if cellVal4 == "" {
			f.SetCellValue("Sheet1", "D"+strconv.Itoa(k), "1")
		}
		if cellVal5 == "" {
			f.SetCellValue("Sheet1", "E"+strconv.Itoa(k), "1")
		}
		if cellVal6 == "" {
			f.SetCellValue("Sheet1", "F"+strconv.Itoa(k), "1")
		}
		k++

	}
}

//将excel数据存入哈希表
func ExcelToMap1(f *excelize.File) (map[int]string,map[string]int,map[string]int){
	rows, _ := f.GetRows("Sheet" + "1")
	// 循环刚刚获取到的表中的值
	map1 := make(map[int]string)
	map2 := make(map[string]int)
	map3 := make(map[string]int)
	var k int = 0
	for key,row := range rows {
		//去掉标题行
		if key > 0 {
			if row[1] != "1" && row[2] != "1" && row[4] != "1" && row[5] != "1" {
				if row[2] == row[4] {
					map1[k] = row[1] + row[2] + row[5]
				} else {
					map1[k] = row[4] + row[5]
				}
			} else if row[0] == "1"{
				map1[k] = row[1]
			}
			if row[4] != "1"{
				_, ok := map3[row[4]]
				if !ok {
					map3[row[4]] = k
				}
			}
		}
		k++
	}
	for key, v := range map1 {
		map2[v] = key
	}
	return map1,map2,map3
}

func ExcelToMap2 (f *excelize.File,map2 map[string]int,map3 map[string]int) map[string]int {
	map4 := make(map[string]int)
	rows, _ := f.GetRows("Sheet" + "1")
	var temp string = ""
	// 循环刚刚获取到的表中的值
	for key,row := range rows {
		//去掉标题行
		if key > 0 {
			if row[0] == "1" && row[2] == "1"&& row[3] == "1" && row[3] == "1"{
				map4[row[1]] = 0
			}
			if  row[0] != "1" && row[2] == row[4] &&row[3] == "1"{
				map4[row[1]+row[2]+row[5]] = map2[row[0]]
				temp = row[2]
			}
			if row[2] != row[4]{
				map4[row[4]+row[5]] = map3[temp]
			}
			if row[3] != "1" {
				map4[row[1]+row[2]+row[5]] = map3[row[3]]
			}
		}
	}
	return map4
}

//构造结构体，写入数据库
func CreateStruct(map1 map[int]string ,map2 map[string]int){
	for i := 1 ;i <= len(map1); i++ {
		 data := Data{
			CurrID:       i,
			Name:     map1[i],
			ParentID: map2[map1[i]],
			CompanyName: map1[1],
		}
		Db.Create(&data)
	}
}

//对上传excel文档进行预处理
//数据对齐（为空数据赋值，便于后续操作）

func DealWithExcel(f *excelize.File) {
	rows, _ := f.GetRows("Sheet1")
	var k int = 1
	for _, _ = range rows {

		cellVal1, _ := f.GetCellValue("Sheet1", "A"+ strconv.Itoa(k))
		cellVal2, _ := f.GetCellValue("Sheet1", "B"+ strconv.Itoa(k))
		cellVal3, _ := f.GetCellValue("Sheet1", "C"+ strconv.Itoa(k))
		cellVal4, _ := f.GetCellValue("Sheet1", "D"+ strconv.Itoa(k))
		cellVal5, _ := f.GetCellValue("Sheet1", "E"+ strconv.Itoa(k))
		cellVal6, _ := f.GetCellValue("Sheet1", "F"+ strconv.Itoa(k))

		if cellVal1 == "" {
			f.SetCellValue("Sheet1", "A"+ strconv.Itoa(k), "1")
		}
		if cellVal2 == "" {
			f.SetCellValue("Sheet1", "B"+ strconv.Itoa(k), "1")
		}
		if cellVal3 == "" {
			f.SetCellValue("Sheet1", "C"+ strconv.Itoa(k), "1")
		}
		if cellVal4 == "" {
			f.SetCellValue("Sheet1", "D"+ strconv.Itoa(k), "1")
		}
		if cellVal5 == "" {
			f.SetCellValue("Sheet1", "E"+ strconv.Itoa(k), "1")
		}
		if cellVal6 == "" {
			f.SetCellValue("Sheet1", "F"+ strconv.Itoa(k), "1")
		}
		k++

	}

}
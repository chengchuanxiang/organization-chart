package model

import (
	"container/list"
	"sort"
	"strconv"
)

type Node struct {
	Name     string
	ParentID     string
	CurrID     string
	datas *list.List
}


func getSize(node *Node)int{
	return (*node).datas.Len()

}
func addChild(currNode *Node, newNode *Node){
	(*currNode).datas.PushBack(newNode)
}

func toString(currNode *Node) string{

	result := "["
	for e := (*currNode).datas.Front(); e != nil; e = e.Next() {
		tmp,_ := (e.Value).(*Node)
		result += toString1(tmp)
		result += ","
	}
	result = result[:len(result)-1]
	result += "]"
	return result
}
func toString1(currNode *Node) string {
	result := "{" + "\"id\" : " + (*currNode).CurrID + ", \"label\" : \"" + (*currNode).Name + "\""

	if (*currNode).datas.Len() != 0 {
		result += ", \"children\": " + toString(currNode)

	}
	return result + "}"
}


//根据公司名字从数据库中获取数据list
func GetDataFromDb(companyName string)  *list.List{
	Datas := list.New()
	var count  int
	Db.Model(&Data{}).Where("company_name = ?",companyName).Count(&count)


	for i := 1; i <= count; i++ {
		maps := make(map[string]string)
		var data Data
		Db.Where(" company_name = ? AND curr_id = ?", companyName, i).First(&data)

		maps["CurrID"] = strconv.Itoa(data.CurrID)
		maps["ParentID"] = strconv.Itoa(data.ParentID)
		maps["Name"] = data.Name

		Datas.PushBack(maps)
	}

	return Datas
}




func MultipleTreeToJSON(companyName string) string{
	datas := GetDataFromDb(companyName)
	maps := make(map[string]*Node)

	var root *Node = &Node{}
	root.datas = list.New()

	for e := datas.Front(); e != nil; e = e.Next(){
		map1,_ := (e.Value).(map[string]string)
		var node *Node = &Node{}
		node.datas = list.New()
		(*node).CurrID = map1["CurrID"]
		(*node).Name = map1["Name"]
		(*node).ParentID = map1["ParentID"]
		maps[node.CurrID] = node
	}

	var  numbers []int
	for index := range maps {
		intVal,_ := strconv.Atoi(index)
		numbers = append(numbers, intVal)
	}
	sort.Ints(numbers)

	for _, val := range numbers {
		val := strconv.Itoa(val)

		if (maps[val]).ParentID == "0" {
			root = maps[val]
		} else {

			addChild(maps[(maps[val]).ParentID], maps[val])

		}

	}


	return toString1(root)

}
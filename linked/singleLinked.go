package linked

import(
	"fmt"
	"github.com/tealeg/xlsx"
	"sync"
)
var lock sync.Mutex
var aa int = 1
//结构体
type PageInfoNode struct{

	returnNum string
	title string
	desc string
	Next *PageInfoNode
}

//初始化
func InitLinked() PageInfoNode{
	node := &PageInfoNode{
	}
	return *node
}

//链表判空
func IsEmpty(pageInfoNode *PageInfoNode) bool{
	return pageInfoNode.Next == nil
}

//插入节点
func Insert(returnNum string,title string,desc string,position *PageInfoNode){
	//defer lock.Unlock()
    lock.Lock()
	tempCell := new(PageInfoNode)
	if tempCell == nil{
		fmt.Println("err:out of space")
	}
	tempCell.returnNum = returnNum
	tempCell.title = title
	tempCell.desc = desc
	tempCell.Next = nil
	for{
		if position.Next == nil {
			break
		}
		position = position.Next
	}
	tempCell.Next = position.Next
	position.Next = tempCell
	lock.Unlock()
}

//显示节点
func ShowNode(pageInfoNode *PageInfoNode){

	if pageInfoNode.Next == nil{
		fmt.Println("the linked id empty")
	}else{
		for{
			pageInfoNode = pageInfoNode.Next
			fmt.Println("[%s,%s,%s]->",pageInfoNode.returnNum,pageInfoNode.title,pageInfoNode.desc)
			
			if pageInfoNode.Next == nil{
				break
			}
			
		}
	}
}

//删除节点
func DelNode(pageInfoNode *PageInfoNode) {
	if pageInfoNode.Next == nil{
		fmt.Println("the linked id empty")
	}else{
		pageInfoNode.Next = pageInfoNode.Next.Next
	}
}


func StartWriting(pageInfoNode *PageInfoNode){
	fmt.Println("into Writing")
	aa= 1
	//wrtingCheck(*pageInfoNode)
}

func Endwriting() bool{
	aa = 0
	fmt.Println("the reading is over,the writing can be stop")
	fmt.Println(aa)
	return true
}

func WrtingCheck(list *PageInfoNode){
	fmt.Println("into Writing  00000")

	var file *xlsx.File
    var sheet *xlsx.Sheet
	var row *xlsx.Row
    var cell *xlsx.Cell
	var err error
	file = xlsx.NewFile()
    sheet, err = file.AddSheet("Sheet1")
    if err != nil {
        fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "页面号"
	cell = row.AddCell()
	cell.Value = "标题"
	cell = row.AddCell()
	cell.Value = "简介"
	for{
		//fmt.Println("in for")
		if (list.Next==nil&&aa==0){
			break
		}else if list.Next!=nil{
			//defer lock.Unlock()
    		lock.Lock()
			//getListeningResult(list)
			list = list.Next
			row = sheet.AddRow()
			cell = row.AddCell()
			cell.Value = list.returnNum
			cell = row.AddCell()
			cell.Value = list.title
			cell = row.AddCell()
			cell.Value = list.desc
			lock.Unlock()
		}
	}
	err = file.Save("../MyXLSXFile.xlsx")
    if err != nil {
        fmt.Printf(err.Error())
	}
	fmt.Println("the excel file is already")
}
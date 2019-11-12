package linked

import(
	"fmt"
	"github.com/tealeg/xlsx"
)

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
func DelNode(pageInfoNode *PageInfoNode)  {
	if pageInfoNode.Next == nil{
		fmt.Println("the linked id empty")
	}else{
		pageInfoNode.Next = pageInfoNode.Next.Next
	}
}




func GetListeningResult(list *PageInfoNode){
	//无结果，无关闭 --- 自旋锁等待
	//有结果        ---- 写入
	//无结果，有关闭 ---- 关闭线程

	//linked.IsEmpty()
	fmt.Println("The start writing")

	var file *xlsx.File
    var sheet *xlsx.Sheet
	var row *xlsx.Row
	//var row1 *xlsx.Row
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

	fmt.Println("get in loop")

	if !IsEmpty(list){
		for{
			list = list.Next
			row = sheet.AddRow()
			cell = row.AddCell()
			cell.Value = list.returnNum
			cell = row.AddCell()
			cell.Value = list.title
			cell = row.AddCell()
			cell.Value = list.desc
			if list.Next == nil{
				break
			}
		} 
	}

	err = file.Save("../MyXLSXFile.xlsx")
    if err != nil {
        fmt.Printf(err.Error())
    }
	fmt.Println("The excel end writing")
}

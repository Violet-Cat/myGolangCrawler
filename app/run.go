package main

import(
	"fmt"
	"myGolangCrawler/linked"
	"net/http"
	"log"
	"strconv"
	"regexp"
	"github.com/PuerkitoBio/goquery"
	"time"
)

func main()  {
	fmt.Println("************************")
	ch := make(chan string)
	list := linked.InitLinked()
	go linked.WrtingCheck(&list,ch)
	client := &http.Client{
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            return http.ErrUseLastResponse
        },
	}
	fmt.Println("The pageInfo is reading,please waiting")
	for i:=10001;i<10015;i++ {
		// 20191012  55575

		time.Sleep(1 * time.Second)

		url := "https://volmoe.com/comic/"+strconv.Itoa(i)+".htm"
		resp, err := client.Get(url)
		//fmt.Println(url)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		//fmt.Println(resp.StatusCode)
		defer resp.Body.Close()
		//只有查出页面的才会写入到结构体中
		if resp.StatusCode == 200 {
			var returnNum,title,desc string
			html, err := goquery.NewDocumentFromReader(resp.Body)
			if err!=nil{
				log.Fatalln(err)
			}
		
			html.Find("#author b").Each(func(i int, selection *goquery.Selection) {
				
				title = selection.Text()
			})
			html.Find("#desc_text").Each(func(i int, selection *goquery.Selection) {
				//正则表达式去除空格
				reg := regexp.MustCompile("\\s+")
    			desc = reg.ReplaceAllString(selection.Text(), "")
				
			})
			returnNum = strconv.Itoa(i)+""
			//fmt.Println(returnNum,"\n",title,"\n",desc)

			
			linked.Insert(returnNum,title,desc,&list)
			fmt.Println("the pageInfo set in struct")
		}
	}
	
	linked.Endwriting()
	
	<-ch

	fmt.Println("************************")
	fmt.Println("the main is over")
	//fmt.Println(linked.IsEmpty(&list))
	/*for !linked.IsEmpty(&list){
		fmt.Println("the linked is not empty")
		//此处睡的时间可以长长一些，不然的话文件可能写不完
		time.Sleep(10 * time.Second)
	}*/
}
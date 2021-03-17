package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

//创建文件夹
func CreateDateDir(basePath string) {
	_, err := os.Stat("./images/" + basePath)
	if err != nil {
		fmt.Println("不存在文件")
		if os.IsNotExist(err) {
			err := os.Mkdir("./images/"+basePath, os.ModePerm)
			if err != nil {
				fmt.Printf("创建失败![%v]\n", err)
				os.Exit(1)
			}
			return
		}
		return
	}
}

//下载图片
func download(herf string, path string) {
	CreateDateDir(path) //创建文件夹
	var a = time.Now().UnixNano()
	fmt.Println(herf)
	resp, err := http.Get(herf)
	if err != nil {
		fmt.Println("访问图片出错")
	}
	_data, _err2 := ioutil.ReadAll(resp.Body)
	if _err2 != nil {
		panic(_err2)
	}
	//保存到本地
	ioutil.WriteFile(fmt.Sprintf("./images/"+path+"/%d.jpg", a), _data, 0644)
	fmt.Println("图片下载成功")
}

//爬取第二岑
func getwinimage(urls string) {
	resp, err := http.Get(urls)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	pathname := doc.Find(".arc_location a").Eq(2).Text() //直接提取title的内容
	fmt.Println(pathname)
	doc.Find("div.fleft.arc_pic .swiper-wrapper a").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Attr("src"))
		_url, _err := selection.Attr("src")
		if _err {
			download(_url, pathname)
			time.Sleep(1 * time.Second)
		}
	})

}

//获取第一层
func getimages(urls string) {
	resp, err := http.Get(urls)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	doc.Find(".egeli_pic_dl dd a").Each(func(i int, selection *goquery.Selection) {
		_href, _err := selection.Attr("href")
		//fmt.Println(selection.Attr("href"))
		if _err { //如果存在
			getwinimage(_href)
		}
	})
}

func Main() {
	fmt.Println("开始爬取")
	//爬取页数
	for i := 1; i <= 1; i++ {
		getimages("https://mm.enterdesk.com/dongmanmeinv/" + strconv.Itoa(i) + ".html")
	}
	fmt.Println("结束爬取")
}

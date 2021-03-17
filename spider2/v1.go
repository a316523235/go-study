package spider2


import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

var cnt = 0

//创建文件夹
func CreateDateDir(basePath string) {
	_, err := os.Stat("./images")
	if err != nil && os.IsNotExist(err)  {
		err := os.Mkdir("./images", os.ModePerm)
		if err != nil {
			fmt.Printf("创建images失败![%v]\n", err)
			os.Exit(1)
		}
	}

	_, err = os.Stat("./images/" + basePath)
	if err != nil && os.IsNotExist(err) {
		err := os.Mkdir("./images/"+basePath, os.ModePerm)
		if err != nil {
			fmt.Printf("创建失败![%v]\n", err)
			os.Exit(1)
		}
		return
	}
}

//下载图片
func download(herf string, path string) {
	CreateDateDir(path) //创建文件夹
	//var a = time.Now().UnixNano()

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
	cnt++
	ioutil.WriteFile(fmt.Sprintf("./images/"+path+"/%d.jpg", cnt), _data, 0644)
	fmt.Println("图片下载成功")
}

//爬取第二岑
func getwinimage(urls string, id int) {
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
	pathname := strconv.Itoa(id) //直接使用id
	fmt.Println(pathname)
	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Attr("src"))
		_url, _err := selection.Attr("src")
		if _err {
			download(_url, pathname)
			time.Sleep(1 * time.Second)
		}
	})

}

func Main() {
	// 从标准输入流中接收输入数据
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("请输入id\n")
	input.Scan()
	idStr := input.Text()
	fmt.Println("id :"  + idStr + "\n")

	fmt.Println("开始爬取")
	idInt, _ := strconv.Atoi(idStr)
	arr := []int{idInt}
	//爬取页数
	for i := 0; i < len(arr); i++ {
		cnt = 0
		id := arr[i]
		//getwinimage("https://mm.enterdesk.com/dongmanmeinv/" + strconv.Itoa(id) + ".html", id)
		getwinimage("http://m.900866.com/activity/gf_" + strconv.Itoa(id), id)
	}
	fmt.Println("结束爬取")
}
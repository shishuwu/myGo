package main

import (
    "fmt"
    "strconv"
    "net/http"
    "os"
	"io"
	"testing"
)

//百度贴吧的地址规律
//第一页:https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8(&pn=0)
//第二页:https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=50
//第三页:https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=100
//所以它的最后的数字每加50,代表着下一页

//整体提取的思路：
//1、先拿地址
//2、爬
//3、取
//4、存


/***/
func HttpGet(url string) (result string, err error) {
    resp, err1 := http.Get(url)
    if err != nil {
        err = err1
        return
    }
    defer resp.Body.Close()
    //读取网页的body内容
    buf := make([]byte, 4*1024)
    for true {
        n, err := resp.Body.Read(buf)
        if err != nil {
            if err == io.EOF{
                fmt.Println("文件读取完毕")
                break
            }else {
                fmt.Println("resp.Body.Read err = ", err)
                break
            }
        }
        result += string(buf[:n])
    }
    return
}

//爬取一个网页
func SpiderPage(i int, page chan <- int)  {
    url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
    fmt.Printf("正在爬取第%d个网页\n", i)
    //爬,将所有的网页内容爬取下来
    result, err := HttpGet(url)
    if err != nil {
        fmt.Println("http.Get err = ", err)
        return
    }
    //把内容写入到文件
    filename := strconv.Itoa((i-1)*50) + ".html"
    f, err1 := os.Create(filename)
    if err1 != nil{
        fmt.Println("os.Create err = ", err1)
        return
    }
    //写内容
    f.WriteString(result)
    //关闭文件
    f.Close()
    //每爬完一个，就给个值
    page<-i
}

func DoWork(start, end int)  {
    fmt.Printf("正在爬取第%d页到%d页\n", start, end)
    //因为很有可能爬虫还没有结束下面的循环就已经结束了，所以这里就需要且到通道
    page := make(chan int)
    for i:=start; i<=end; i++ {
        //将page阻塞
        go SpiderPage(i, page)
    }
    for i:=start; i<=end; i++ {
        fmt.Printf("第%d个页面爬取完成\n",<-page)//这里直接将面码传给点位符，值直接从管道里取出
    }
}

func TestEntrance(t *testing.T){
    var start, end int
    start = 1
    end = 100

    fmt.Printf("请输入起始页>=1：> ")
    fmt.Scan(&start)
    fmt.Printf("请输入结束页：> ")
    fmt.Scan(&end)
    DoWork(start, end)
}

// func main() {
//     var start, end int
//     fmt.Printf("请输入起始页>=1：> ")
//     fmt.Scan(&start)
//     fmt.Printf("请输入结束页：> ")
//     fmt.Scan(&end)
//     DoWork(start, end)
// }
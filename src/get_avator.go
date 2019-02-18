package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
)

func httpGet(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(body)
}
func httpRegex(str string) []string {

	//正则表达式，有点菜，只会(.*?)
	// regex := "<span class=" + "\"" + "img-hash" + "\"" + ">(.*?)</span>"
	regex := "<li><a[^>]*>([^<]*)</a></li>"

	reg := regexp.MustCompile(regex)

	dataS := reg.FindAllSubmatch([]byte(str), -1)

	results := make([]string, 0)

	for _, v := range dataS {

		results = append(results, string(v[1]))
	}
	fmt.Println(results)
	var newSlice = results[:len(results)-2]
	return newSlice
}

func getImg(url string) (n int64, err error) {

	path := strings.Split(url, "/")

	var name string

	if len(path) > 1 {

		name = path[len(path)-1]
	}

	out, err := os.Create(name) //创建文件

	defer out.Close()

	pix := httpGet(url) //获取图片

	pic := []byte(pix)

	n, err = io.Copy(out, bytes.NewReader(pic))

	return
}

func main() {

	err := os.MkdirAll("image", os.ModePerm) //创建image目录

	if err != nil {

		fmt.Println(err)

	}

	os.Chdir("./image") //修改工作目录

	url := "https://github.com/krman009/Octodex-Images-Links"

	saveImg(url)

	fmt.Println(runtime.NumGoroutine())

}
func saveImg(url string) {
	str := httpGet(url)
	reg := httpRegex(str)
	// return
	results := make([]string, 0)

	for _, v := range reg {

		results = append(results, v)

	}

	//遍历url

	for _, url := range results {

		getImg(url)
	}
	return
}

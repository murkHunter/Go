package main

import (
	"fmt"
	"time"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func httpGet(url string,  i int) string {
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

	// fmt.Println(i)
	return string(body)
}
func httpPost(url string,  i int, count int) string {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader("message=233333&name=23333&id=&touxiang=8"))

	if err != nil {
		fmt.Println(err)
		return ""
	}
	count++
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(string(body))
	return string(body)
}

func main() {
	// url := "http://120.78.173.0:3000/issue/getList"
	url := "http://120.78.173.0:3000/issue/publish"
	// var r string
	var count int = 1;
		for i:=0; i<1000;i++ {
			go func(i int) {
				httpPost(url, i, count)
			}(i)
	
		}
	runtime.GOMAXPROCS(runtime.NumCPU())
	time.Sleep(time.Second)
	fmt.Println("over")
	fmt.Println(count)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(0))

}


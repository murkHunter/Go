package engine

import (
	"fmt"
	"log"

	"../fetcher"
)

func Run(seeds ...Request) { // 语法糖， ...用于接受不确定个数的参数 ，写在右边用于打散元素，方便append
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher:error"+"fetching url %s: %v", r.Url, err)
			continue
		}
		// fmt.Printf("%s", body)
		parseResult := r.ParserFunc(body)
		fmt.Println(parseResult)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}

}

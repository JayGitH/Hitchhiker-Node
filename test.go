package main

import (
	"fmt"
)

func main() {

	tCase := &testCase{
		requestBodys: []requestBody{
			requestBody{
				id:      "1",
				method:  "GET",
				url:     "https://httpbin.org/get",
				headers: make(map[string]string),
				tests:   "",
			},
			requestBody{
				id:      "2",
				method:  "POST",
				url:     "https://httpbin.org/post",
				headers: map[string]string{"11": "22"},
				body:    "test",
				tests:   "",
			},
		},
		totalCount:       10,
		concurrencyCount: 1,
		qps:              1,
		timeout:          600,
		trace: func(result runResult) {
			fmt.Printf("result: %s \n", result.body)
		},
	}

	tCase.Run()
	for rst := range tCase.results {
		fmt.Printf("result length: %d \n", len(rst))
		// for _, c := range rst {
		// 	fmt.Printf("result: %s \n", c.body)
		// }
	}
}
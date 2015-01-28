package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	t66y  string = "http://t66y.com/thread0806.php?fid=16"
	zhihu string = "http://www.zhihu.com"
	mobo  string = "http://www.moboplayer.com"
)

func main() {
	fmt.Printf("hello world\n")
	fmt.Printf("%v", t66y)

	resp, err := http.Get(zhihu)
	if err != nil {
		fmt.Printf("error: %v", err.Error())
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("error: %v", err.Error())
	}

	fmt.Printf("%v", string(body))

}

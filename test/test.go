package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/newham/hamtask"
)

func main() {
	start := time.Now().Unix()
	loop := 1
	n := 200
	c := make(chan int)
	for i := 0; i < loop; i++ {
		go func() {
			TestTask(c, n)
		}()
	}
	//waite
	for i := 0; i < loop; i++ {
		<-c
	}
	end := time.Now().Unix()
	println("cost time:", end-start, "second")
}

func TestTask(c chan int, n int) {
	hamtask.NewSimpleWorker(n, func(i int, d hamtask.Data) {
		// println(i, d.String())
		url := d.String()
		// 1. test dc.AddURL()
		url = fmt.Sprintf("%s?cc_name=%s&f_name=%s&args=[D%d,%s]", url, "dc", "AddURL", 1000+i, "https://test.com/test.res")
		// println(i, url)
		// return
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Error", err.Error())
			return
		}
		println(resp.Status)
		// b, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	// log.Println("Error", err.Error())
		// 	return
		// }
		// println(i, string(b))

	}, func() hamtask.Data {
		return hamtask.String("http://localhost:8001/fabric-iot/test")
	}, n).Start()
	c <- 1
}

// func GetURL(f_name string,)

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/newham/hamtask"
)

func main() {
	n := GetCount()
	if n == -1 {
		return
	}
	start := time.Now().Unix()
	// start := time.Now().UnixNano()
	Loop(1, n, TestDCTask)
	end := time.Now().Unix()
	println("cost time:", end-start, "second") //(end-start)/int64(time.Millisecond)
}

func GetCount() int {
	args := os.Args
	n := -1
	if len(args) > 1 {
		m, err := strconv.Atoi(args[1])
		if err != nil {
			println("bad count")
			return -1
		}
		n = m
	}
	return n
}

func Loop(loop int, n int, f func(c chan int, n int)) {
	c := make(chan int)
	for i := 0; i < loop; i++ {
		go func() {
			f(c, n)
		}()
	}
	//waite
	for i := 0; i < loop; i++ {
		<-c
	}
}

func TestDCTask(c chan int, n int) {
	i := 0
	hamtask.NewSimpleWorker(n, func(i int, d hamtask.Data) {
		url := d.String()
		_, err := http.Get(url)
		if err != nil {
			log.Println("Error", err.Error())
			return
		}
		// println(resp.Status)

	}, func() hamtask.Data {
		i++
		return hamtask.String(GetURL("http://localhost:8001/fabric-iot/test", "dc", "AddURL", strconv.Itoa(i), "https://test.res"))
	}, n).Start()
	c <- 1
}

func GetURL(url, cc_name, f_name string, args ...string) string {
	return fmt.Sprintf("%s?cc_name=%s&f_name=%s&args=[%s]", url, cc_name, f_name, strings.Join(args, ","))
}

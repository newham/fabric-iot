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
	ccName, fName, n := GetArgs()
	if n == -1 {
		println("please input client count")
		return
	}
	start := time.Now().Unix()
	// start := time.Now().UnixNano()
	if ccName == "dc" {
		if fName == "AddURL" {
			Loop(1, n, TestDCAddURL)
		} else if fName == "GetURL" {
			Loop(1, n, TestDCGetURL)
		}
	} else {
		println("please input [cc_name] [f_name] [n]")
	}

	end := time.Now().Unix()
	println("cost time:", end-start, "second") //(end-start)/int64(time.Millisecond)
}

func GetArgs() (string, string, int) {
	args := os.Args
	if len(args) > 3 {
		cc_name := args[1]
		f_name := args[2]
		n, err := strconv.Atoi(args[3])
		if err != nil {
			println("bad count")
			return "", "", -1
		}
		return cc_name, f_name, n
	}
	return "", "", -1
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

func TestDCGetURL(c chan int, n int) {
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
		return hamtask.String(GetURL("http://localhost:8001/fabric-iot/test", "dc", "GetURL", strconv.Itoa(i)))
	}, n).Start()
	c <- 1
}

func TestDCAddURL(c chan int, n int) {
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

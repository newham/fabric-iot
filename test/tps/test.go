package main

import (
	"fmt"
	"log"
	"math/rand"
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
	} else if ccName == "pc" {
		if fName == "AddPolicy" {
			Loop(1, n, TestPCAddPolicy)
		} else if fName == "GetPolicy" {
			Loop(1, n, TestPCGetPolicy)
		} else if fName == "DeletePolicy" {
			Loop(1, n, TestPCDeletePolicy)
		} else if fName == "UpdatePolicy" {
			Loop(1, n, TestPCUpdatePolicy)
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

func TestPCAddPolicy(c chan int, n int) {
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
		now := time.Now().Unix()
		policyStr := fmt.Sprintf(`{"AS":{"userId":"%s","role":"u1","group":"g1"},"AO":{"deviceId":"%s","MAC":"%s"},"AP":1,"AE":{"createdTime":%d,"endTime":%d,"allowedIP":"*.*.*.*"}}`, GetUserID(i), GetDeviceID(i), RandomMac(), now, now+100000)
		return hamtask.String(GetURL("http://localhost:8001/fabric-iot/test", "pc", "AddPolicy", policyStr))
	}, n).Start()
	c <- 1
}

func TestPCGetPolicy(c chan int, n int) {
}

func TestPCDeletePolicy(c chan int, n int) {
}

func TestPCUpdatePolicy(c chan int, n int) {
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
		return hamtask.String(GetURL("http://localhost:8001/fabric-iot/test", "dc", "GetURL", GetDeviceID(i)))
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
		return hamtask.String(GetURL("http://localhost:8001/fabric-iot/test", "dc", "AddURL", GetDeviceID(i), "https://test"+GetDeviceID(i)+".res"))
	}, n).Start()
	c <- 1
}

func GetID(i int) int {
	return i + 10000
}

func GetUserID(i int) string {
	return strconv.Itoa(GetID(i))
}

func GetDeviceID(i int) string {
	return fmt.Sprintf("D%d", GetID(i))
}

func GetURL(url, cc_name, f_name string, args ...string) string {
	return fmt.Sprintf("%s?cc_name=%s&f_name=%s&args=[%s]", url, cc_name, f_name, strings.Join(args, ","))
}

func RandomX(N, n int) string {
	ip := []string{}
	for i := 0; i < N; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		ip = append(ip, strconv.FormatInt(int64(r.Intn(n)), 16))
	}
	return strings.Join(ip, ":")
}

func RandomMac() string {
	return RandomX(6, 64)
}

func RandomIPv6() string {
	return RandomX(8, 65536)
}

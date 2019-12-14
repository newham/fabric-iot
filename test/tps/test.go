package main

import (
	"crypto/sha256"
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

var BASE_URL = "http://localhost:8001/fabric-iot/test"

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
		} else if fName == "QueryPolicy" {
			Loop(1, n, TestPCQueryPolicy)
		} else if fName == "DeletePolicy" {
			Loop(1, n, TestPCDeletePolicy)
		} else if fName == "UpdatePolicy" {
			Loop(1, n, TestPCUpdatePolicy)
		}
	} else if ccName == "ac" {
		if fName == "CheckAccess" {
			Loop(1, n, TestACCheckAccess)
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
	TestCC(c, n, func(i int) string {
		return GetURL("pc", "AddPolicy", GetPolicyReq(i, 1, 1, 1))
	})
}

func GetPolicyReq(i, r, g, AP int) string {
	now := time.Now().Unix()
	policyStr := fmt.Sprintf(`{"AS":{"userId":"%s","role":"u%d","group":"g%d"},"AO":{"deviceId":"%s","MAC":"%s"},"AP":%d,"AE":{"createdTime":%d,"endTime":%d,"allowedIP":"*.*.*.*"}}`, GetUserID(i), r, g, GetDeviceID(i), RandomMac(), AP, now, now+100000)
	return policyStr
}

func TestPCQueryPolicy(c chan int, n int) {
	TestCC(c, n, func(i int) string {
		return GetURL("pc", "QueryPolicy", GetPolicyID(i))
	})
}

func GetPolicyID(i int) string {
	return GetSHA256(GetUserID(i), GetDeviceID(i))
}

func GetSHA256(args ...string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(strings.Join(args, ""))))
}

func TestPCDeletePolicy(c chan int, n int) {
	TestCC(c, n, func(i int) string {
		return GetURL("pc", "DeletePolicy", GetPolicyID(i))
	})
}

func TestPCUpdatePolicy(c chan int, n int) {
	TestCC(c, n, func(i int) string {
		return GetURL("pc", "UpdatePolicy", GetPolicyReq(i, 1, 1, 0))
	})
}

func TestACCheckAccess(c chan int, n int) {
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
		return hamtask.String(GetURL("ac", "CheckAccess", GetACRequest(i)))
	}, n).Start()
	c <- 1
}

func TestDCGetURL(c chan int, n int) {

	TestCC(c, n, func(i int) string {
		return GetURL("dc", "GetURL", GetDeviceID(i))
	})
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
		return hamtask.String(GetURL("dc", "AddURL", GetDeviceID(i), "https://test"+GetDeviceID(i)+".res"))
	}, n).Start()
	c <- 1
}

func TestCC(c chan int, n int, f func(int) string) {
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
		return hamtask.String(f(i))
	}, n).Start()
	c <- 1
}

func GetACRequest(i int) string {
	return fmt.Sprintf(`{"AS":{"userId":"%s","role":"u1","group":"g1"},"AO":{"deviceId":"%s","MAC":"00:11:22:33:44:55"}}`, GetUserID(i), GetDeviceID(i))
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

func GetURL(cc_name, f_name string, args ...string) string {
	return fmt.Sprintf("%s?cc_name=%s&f_name=%s&args=%s", BASE_URL, cc_name, f_name, strings.Join(args, "|"))
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

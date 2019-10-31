package main

import (
	"encoding/json"
	"github.com/abusizhishen/req"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const Single = 3

var count int = 5000
func main() {
	for{
		log.Print("check")
		if vote(rand.Intn(40)+1){
			time.Sleep(time.Second*time.Duration(rand.Intn(10)))
		}else{
			for !vote(24){
				votePlus()
				count*=2

			}
			count = 5000
		}
	}
}

var userAgent = []string{
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:42.0) Gecko/20100101 Firefox/42.0",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safari/537.36 OPR/38.0.2220.41",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0)",
	"Googlebot/2.1 (+http://www.google.com/bot.html)",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
}

func vote(num int) (result bool) {
	url := "https://zhuangding.xueersi.company/vote/"
	headers := map[string]string{
		"user-agent":userAgent[rand.Intn(len(userAgent))],
	}
	if num > 0 {
		url = url+strconv.Itoa(num)
	}
	body,err := req.PostForm(url,nil,headers)
	if err != nil{
		return false
	}

	var rank []int
	err = json.Unmarshal(body,&rank)
	if err != nil{
		return false
	}

	return checkout(rank)
}

func votePlus() bool {
	log.Println("加票")
	log.Print("总数:",count)
	var t = time.Now()
	for i:=0;i<count;i++{
		for _,num := range three{
			vote(num)
		}
		time.Sleep(time.Millisecond*time.Duration(rand.Intn(10)))
	}
	log.Print("加票完毕,耗时",time.Since(t))
	return vote(24)
}

var three = map[int]int{
	3:3,18:18,20:20,24:24,
}


func checkout(i []int) bool {
	i = i[0:4]
	for _,num := range i{
		if _,ok := three[num];!ok{
			return false
		}
	}

	return true
}
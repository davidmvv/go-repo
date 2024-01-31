package main

import (
	"fmt"
	"sync"
	"time"
	"net/http"
)

func main(){
	WaitGroupsMain();
}

func WaitGroupsMain(){
	var wg sync.WaitGroup
	
	start := time.Now()
	
	/*
	wg.Add(3)

	go loopNumbers(&wg,0,100)
	go loopNumbers(&wg,100,200)
	go loopNumbers(&wg,200,300)

	wg.Wait()
	*/
	urls := []string{
		"http://google.com",
		"http://onlyfans.com",
		"http://youtube.com",
		"http://twitch.com",
		"http://twitter.com",
	}

	getStatus(&wg,urls)

	elapsed := time.Since(start)

	fmt.Printf("Representinh another process\n")
	
	fmt.Printf("Time elapsed %f \n", float32(elapsed/1000000000))
}

func loopNumbers(wg *sync.WaitGroup, init int, end int){
	time.Sleep(time.Second * 3)
	for i:=init;i<end;i++ {
		fmt.Println(i)
	}
	wg.Done()

}

func getStatus(wg *sync.WaitGroup, urls []string){
	for _, url := range urls {
		wg.Add(1)
		go func(url string){
			res,err := http.Get(url)
			if(err != nil){
				fmt.Printf("[Error: %s x -> %s]\n",err,url)
			}else{
				fmt.Printf("[%d %s] \n", res.StatusCode, url)
			}
			wg.Done()
		}(url)
	}
	wg.Wait()
	fmt.Println("All goroutines are DONE\n")
}

/*
func WaitGroupsMain(){
	var wg sync.WaitGroup
	
	start := time.Now()
	
	wg.Add(3)

	 go func() {		
		time.Sleep(time.Second*3)
		for i:=0;i<100;i++ {
			fmt.Println(i)
		}
		wg.Done()
	 }()

	go func() {
		time.Sleep(time.Second*3)
		for i:=99;i<200;i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()
	
	go func() {
		time.Sleep(time.Second*3)
		for i:=200;i<300;i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()
	
	wg.Wait()
	elapsed := time.Since(start)
	
	fmt.Printf("Time elapsed %f \n", float32(elapsed/1000000000))
}
*/
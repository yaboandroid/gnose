package main

import (
	"fmt"
	. "github.com/yaboandroid/gnose"
	"time"
	"flag"
)

func TestCloseChannel() {
	var ch chan int
	ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
	// fentch data from ch channel

	for {
		var b int
		b, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(b)
	}

}

func TestRangeChannel() {
	var c1 chan interface{}
	c1 = make(chan interface{}, 100)
	go func() {
		for i := 0; i < 100; i++ {
			if i % 2 == 0 {
				c1 <- i
			}
		}
		close(c1)
	}()

	for v := range c1 {
		fmt.Println(v)
	}

}

func IsShusu(a int) (flag bool) {
	flag = true
	for i := 2; i < a; i++ {
		if a % i == 0 {
			flag = false
			break
		}
	}
	return
}

func cacl(intChan, resultChan chan int, exitChan chan bool) {
	//judge shusu
	for v := range intChan {
		if IsShusu(v) {
			resultChan <- v
		}
	}
	exitChan <- true
}

func TestCountShusu() {
	intChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)
	go func() {
		for i := 0; i < 1000; i++ {
			intChan <- i
		}
		close(intChan)
	}()
	for i := 0; i < 8; i++ {
		go cacl(intChan, resultChan, exitChan)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(resultChan)
	}()

	for v := range resultChan {
		fmt.Println(v)
	}
}

func TestBasicFunc() {
	const (
		a = iota
		b
		c
		d
	)
	fmt.Println(a)
	second := time.Now()
	fmt.Println(second)
	type Month int
	const (
		January Month = 1 + iota
		February
		March
		April
		May
		June
		July
		August
		September
		October
		November
		December
	)
	fmt.Println(January, February, March, April, May, June, July, August, September, October, November, December)
	year, month, day := time.Now().Date()
	fmt.Printf("Year:%d\nMonth:%s\nDay:%d\n", year, month, day)
	fmt.Println(time.Now().String())
	now := time.Now().Format("2006-01-02 15:04:05.999999")
	fmt.Println(now)

	type student struct {
		Name string
		Age  int
	}

	stu := student{
		Name:"lsui",
		Age:30,
	}

	logger := NewLogger(GetCurrentDir(), "log.txt")
	logger.Info("this is info test")
	logger.Info(fmt.Sprintf("this year is %d", year))
	logger.Info(stu)
	defer func() {
		if err := recover(); err != nil {
			logger.Warning("skip this error:%v", err)
		}
	}()
	logger.Warning("this is warning msg")
	logger.Debug("this is debug message")
	logger.Error("got an error")
	logger.Info("**************")
	logger.Info()
	logger.Info("this is 1 part", stu, "this is 3 part")
	logger.Info("name is %s and age is %d", stu.Name, stu.Age)
	//logger.Exception("panic occurs")
	m := 1
	n := 2
	as := NewAssert(logger)
	as.AssertNonCriticalTrue(m == n, "m is not equal to n")
	as.AssertNonCriticalEqual(m, n, "m match to n")

}

func test(ch chan bool, i int) {
	log1 := NewLogger(GetCurrentDir(), "log1.txt")
	log1.Info("test goroutine")
	time.Sleep(time.Second * 1)
	log1.Warning("warning msg")
	log1.Debug("i is : %d", i)
	ch <- true
}

func TestLock() {
	var ch chan bool
	ch = make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go test(ch, i)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}

}

func empty() {}

func TestFlagUsage() {
	exit := flag.Bool("e", false, "exit auto after test done")
	name := flag.String("n", "ariel", "user name")
	age := flag.Int("age", 0, "age of account")
	flag.Parse()

	fmt.Printf("exit value : %v\n", *exit)
	fmt.Printf("name is : %s\n", *name)
	fmt.Printf("age is %d\n", *age)

	fmt.Println("-----------------------")
	for i, param := range flag.Args() {
		fmt.Printf("num : %d\nnon-flag : %v\n", i, param)
	}
}

func TestCreateFolder() {
	base := GetCurrentDir()
	tobecreate := "d"
	err := CreateFolder(base, tobecreate)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	fmt.Println("this is test demo")
	//TestCloseChannel()
	//TestRangeChannel()
	//TestCountShusu()
	//TestBasicFunc()
	//TestLock()
	//empty()
	//TestFlagUsage()
	//TestCreateFolder()

}

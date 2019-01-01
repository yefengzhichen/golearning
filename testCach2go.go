package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/muesli/cache2go"
)

type student struct {
	id    int
	name  string
	age   int
	score int
}

func main() {
	cache := cache2go.Cache("studentInfo")
	jack := &student{1, "jack", 12, 90}
	tom := &student{2, "tom", 13, 65}
	tony := &student{3, "tony", 12, 80}
	saveTime := 5 * time.Second
	cache.Add(1, saveTime, jack)
	cache.Add(2, saveTime, tom)
	cache.Add(3, saveTime, tony)
	fmt.Println("students num:", cache.Count())

	printFunc := func(key interface{}, stu *cache2go.CacheItem) {
		fmt.Println(key, stu.Data())
	}
	cache.Foreach(printFunc)

	loadFunc := func(key interface{}, args ...interface{}) *cache2go.CacheItem {
		ftmp := &student{key.(int), "", 0, 0}
		item := cache2go.NewCacheItem(key, 0, ftmp)
		return item
	}
	cache.SetDataLoader(loadFunc)
	tmp, ok := cache.Value(4)
	if ok == nil {
		fmt.Println(tmp.Data())
	}
	logFile, err := os.Create("./log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer logFile.Close()
	logger := log.New(logFile, "mytest_", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("this is test line")
	cache.SetLogger(logger)
	time.Sleep(7 * time.Second)
}

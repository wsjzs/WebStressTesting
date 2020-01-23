package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"webStressTesting/model"
)

var target model.Target
var isWorking = false      //switch of goroutines
var gate = make(chan bool) //switch of counter
var count uint64           //count request times

func Work(jsonCommand []byte) [] byte {
	var mama model.Mama
	var papa model.Papa
	var result []byte
	err := json.Unmarshal(jsonCommand, &mama)
	if err != nil {
		//unmarshal fail
		papa := model.BuildPapa(0, err.Error(), "")
		result, _ = json.Marshal(papa)
		return result

	}
	if mama.Purpose == "start" {
		papa = start(mama)

	}
	if mama.Purpose == "count" {
		papa = model.BuildPapa(1, "request times", strconv.FormatUint(getCount(), 10))
	}
	if mama.Purpose == "stop" {
		isWorking = false
		resetCount()
		log.Println("set isWorking false")
		papa=model.BuildPapa(1,"have stop","")
	}
	if mama.Purpose == "quit" {
		quit()
	}
	result, _ = json.Marshal(papa)
	return result
}

func quit() {
	os.Exit(0)
}
func start(mama model.Mama) model.Papa {
	if isWorking {
		//have already working,don't need to start again
		return model.BuildPapa(0, "have already working", "")
		log.Println("already run")
	}
	log.Println("start work")
	target.Url = mama.Url
	target.Interval = mama.Interval
	target.Goroutine = mama.Goroutine
	isWorking = true
	startRoutines()
	return model.BuildPapa(1, "success", "")
}

func startRoutines() {
	startCounter()
	routine := func() {
		for isWorking {
			request(target.Url)
			gate <- true
			time.Sleep(time.Duration(target.Interval) * time.Millisecond)
		}
	}
	//start routines
	for i := 0; i < target.Goroutine; i++ {
		go routine()
	}

}

func startCounter() {
	go func() {
		for i := range gate {
			if i == false {
				break
			}
			count++
		}
		return
	}()
}
func resetCount() {
	count = 0
}
func getCount() uint64 {
	return count
}
func request(url string) {
	resp, err:= http.Get(url)
	if err!=nil{
		return
	}
	_, err= ioutil.ReadAll(resp.Body)
	if err!=nil {
		return
	}
	defer resp.Body.Close()
	log.Println(url)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/0Hidder/novelupdatesscrapperv1/connection"
	"github.com/jasonlvhit/gocron"
)

func task() {
	resp, err := http.Get("http://localhost:12345/getAllNovels")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	novels := connection.ReturnNewNovelStructureArray()

	jerr := json.Unmarshal(body, &novels)

	if jerr != nil {
		fmt.Println("error:", err)
	}

	for i, s := range novels {
		fmt.Println(s.Name)
		fmt.Println(i)
	}

}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	// Do jobs without params
	gocron.Every(1).Day().At("00:00").Do(task)
	// gocron.Every(2).Seconds().Do(task)
	// gocron.Every(1).Minute().Do(task)
	// gocron.Every(2).Minutes().Do(task)
	// gocron.Every(1).Hour().Do(task)
	// gocron.Every(2).Hours().Do(task)
	// gocron.Every(1).Day().Do(task)
	// gocron.Every(2).Days().Do(task)
	// gocron.Every(1).Week().Do(task)
	// gocron.Every(2).Weeks().Do(task)

	// // Do jobs with params
	// gocron.Every(1).Second().Do(taskWithParams, 1, "hello")

	// // Do jobs on specific weekday
	// gocron.Every(1).Monday().Do(task)
	// gocron.Every(1).Thursday().Do(task)

	// // Do a job at a specific time - 'hour:min:sec' - seconds optional
	// gocron.Every(1).Day().At("10:30").Do(task)
	// gocron.Every(1).Monday().At("18:30").Do(task)
	// gocron.Every(1).Tuesday().At("18:30:59").Do(task)

	// // Begin job immediately upon start
	// gocron.Every(1).Hour().From(gocron.NextTick()).Do(task)

	// // Begin job at a specific date/time
	// t := time.Date(2019, time.November, 10, 15, 0, 0, 0, time.Local)
	// gocron.Every(1).Hour().From(&t).Do(task)

	// // NextRun gets the next running time
	// _, time := gocron.NextRun()
	// fmt.Println(time)

	// // Remove a specific job
	// gocron.Remove(task)

	// // Clear all scheduled jobs
	// gocron.Clear()

	// // Start all the pending jobs
	<-gocron.Start()

	// also, you can create a new scheduler
	// to run two schedulers concurrently
	// s := gocron.NewScheduler()
	// s.Every(20).Seconds().Do(task)
	// <-s.Start()
}

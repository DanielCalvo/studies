package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//{
//"start_time": "2020-04-23T12:03:27+00:00",
//"exit_status": "0",
//"end_time": "2020-04-23T12:03:58+00:00",
//"nomad_alloc_id": "e51d6910-4b6d-7708-0e27-c7aacf74197e",
//"nomad_job_name": "os-training-app-stage-be-cron-dcalvo_test/periodic-1587643380"
//}

type test_struct struct {
	StartTime    time.Time
	exitStatus   int
	EndTime      time.Time
	NomadAllocId string
	NomadJobName string
}

// curl -X POST -d "{\"test\": \"that\"}" http://localhost:8082/test

func HandleJson(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t test_struct
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	log.Println(t.Test)
}

func main() {
	http.HandleFunc("/test", HandleJson)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

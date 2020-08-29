package main

import (
	"encoding/json"
	"fmt"
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
//"nomad_job_name": "banana"
//}

type JobJson struct {
	StartTime    time.Time `json:"start_time"`
	ExitStatus   string    `json:"exit_status"`
	EndTime      time.Time `json:"end_time"`
	NomadAllocID string    `json:"nomad_alloc_id"`
	NomadJobName string    `json:"nomad_job_name"`
}

//JSON_STRING=$(jq -n --arg st 2020-04-23T12:03:27+00:00 --arg es 0 --arg et 2020-04-23T12:03:58+00:00 --arg na e51d6910-4b6d-7708-0e27-c7aacf74197e --arg nj banana '{start_time: $st, exit_status: $es, end_time: $et, nomad_alloc_id: $na, nomad_job_name: $nj}')
//curl -X POST -d "$JSON_STRING" http://localhost:8080/json

// curl -X POST -d "{\"test\": \"that\"}" http://localhost:8082/test

func HandleJson(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Body:")
	fmt.Println(string(body))
	var jobJson JobJson
	err = json.Unmarshal(body, &jobJson)
	if err != nil {
		panic(err)
	}
	log.Println(jobJson)
	fmt.Println(jobJson.NomadJobName)
}

func main() {
	http.HandleFunc("/json", HandleJson)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Build struct {
	Number    int    `json:"number"`
	Result    string `json:"result"`
	Timestamp int64  `json:"timestamp"`
	Duration  int64  `json:"duration"`
	UserID    string `json:"actions,omitempty"`
}

type Job struct {
	Name   string  `json:"name"`
	Builds []Build `json:"builds"`
}

type JenkinsResponse struct {
	Jobs []Job `json:"jobs"`
}

const jenkinsURL = "http://localhost:8080/job/test/api/json?tree=builds[number,result,timestamp,duration]"

func fetchJenkinsJobs() ([]Build, error) {
	jenkinsUser := os.Getenv("JENKINS_USER")
	jenkinsToken := os.Getenv("JENKINS_TOKEN")
	client := &http.Client{}
	req, err := http.NewRequest("GET", jenkinsURL, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(jenkinsUser, jenkinsToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jobData Job
	if err := json.NewDecoder(resp.Body).Decode(&jobData); err != nil {
		return nil, err
	}
	return jobData.Builds, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	jobs, err := fetchJenkinsJobs()
	if err != nil {
		http.Error(w, "Failed to fetch Jenkins jobs", http.StatusInternalServerError)
		return
	}
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Jenkins Job History</title>
	</head>
	<body>
		<h1>Jenkins Job History</h1>
		<table border="1">
			<tr><th>Build Number</th><th>Status</th><th>Start Time</th><th>Duration (ms)</th></tr>
			{{range .}}
			<tr><td>{{.Number}}</td><td>{{.Result}}</td><td>{{.Timestamp}}</td><td>{{.Duration}}</td></tr>
			{{end}}
		</table>
	</body>
	</html>
	`
	t := template.Must(template.New("webpage").Parse(tmpl))
	t.Execute(w, jobs)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started at :8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

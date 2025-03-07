package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	jenkinsURL = "http://host.docker.internal:8080"
)

type JenkinsEntry struct {
	User       string `json:"user"`
	Permission string `json:"permission"`
}

func setJenkinsEntry(entry JenkinsEntry) error {
	scriptTemplate, err := os.ReadFile("set_permission.groovy")
	if err != nil {
		return err
	}
	script := fmt.Sprintf(string(scriptTemplate), entry.Permission, entry.User)

	req, err := http.NewRequest(
		http.MethodPost,
		jenkinsURL+"/scriptText",
		bytes.NewBuffer([]byte("script="+script)))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	const (
		adminUser  = "admin"
		adminToken = "11c4e1d3cc413fa2ec4fac4f7cacdeff8e"
	)
	req.SetBasicAuth(adminUser, adminToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to set permission: %s", string(bodyBytes))
	}

	// bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// println(string(bodyBytes))

	return nil
}

func main() {
	entries := []JenkinsEntry{
		{User: "example-user", Permission: "hudson.model.Item.Read"},
		{User: "example-user", Permission: "hudson.model.Item.Build"},
		{User: "example-user", Permission: "hudson.model.Item.Configure"},
		{User: "example-user", Permission: "hudson.model.Item.Create"},
		{User: "example-user", Permission: "hudson.model.Item.Delete"},
		{User: "example-user", Permission: "hudson.model.Item.Discover"},
		{User: "example-user", Permission: "hudson.model.Item.Move"},
		{User: "example-user", Permission: "hudson.model.Item.Workspace"},
		{User: "example-user", Permission: "hudson.model.Run.Delete"},
		{User: "example-user", Permission: "hudson.model.Run.Update"},
		{User: "example-user", Permission: "hudson.model.View.Configure"},
		{User: "example-user", Permission: "hudson.model.View.Create"},
		{User: "example-user", Permission: "hudson.model.View.Delete"},
		{User: "example-user", Permission: "hudson.model.View.Read"},
	}

	for _, entry := range entries {
		err := setJenkinsEntry(entry)
		if err != nil {
			log.Fatalf("Error setting permission %s for user %s: %v", entry.Permission, entry.User, err)
		}
	}

	fmt.Println("Permission set successfully")
}

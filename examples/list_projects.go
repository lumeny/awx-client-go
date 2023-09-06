// This example shows how to list the projects.
//
// Use the following command to build and run it with all the debug output sent to the standard
// error output:
//
//	go run list_projects.go \
//		-url "https://awx.example.com/api" \
//		-username "admin" \
//		-password "..." \
//		-ca-file "ca.pem"

package main

import (
	"flag"
	"fmt"

	"github.com/CenturyLink/hca-awx-client-go/awx"
)

var (
	// url      string
	username string
	password string
	proxy    string
	// insecure bool
	caFile string
)

func init() {
	flag.StringVar(&url, "url", "https://awx.example.com/api", "API URL.")
	flag.StringVar(&username, "username", "admin", "API user name.")
	flag.StringVar(&password, "password", "password", "API user password.")
	flag.StringVar(&proxy, "proxy", "", "API proxy URL.")
	flag.BoolVar(&insecure, "insecure", false, "Don't verify server certificate.")
	flag.StringVar(&caFile, "ca-file", "", "Trusted CA certificates.")
}

func main() {
	// Parse the command line:
	flag.Parse()

	// Connect to the server, and remember to close the connection:
	connection, err := awx.NewConnectionBuilder().
		URL(url).
		Username(username).
		Password(password).
		Proxy(proxy).
		CAFile(caFile).
		Insecure(insecure).
		Build()
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	// Find the resource that manages the collection of projects:
	projectsResource := connection.Projects()

	// Send the request to get the list of projects:
	getProjectsRequest := projectsResource.Get()
	getProjectsResponse, err := getProjectsRequest.Send()
	if err != nil {
		panic(err)
	}

	// Print the results:
	projects := getProjectsResponse.Results()
	for _, project := range projects {
		fmt.Printf("%d: %s - %s\n", project.Id(), project.Name(), project.SCMURL())
	}
}

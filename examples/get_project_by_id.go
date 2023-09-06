// This example shows how to retrieve a specific project.
//
// Use the following command to build and run it with all the debug output sent to the standard
// error output:
//
//	go run get_project_by_id.go \
//		-url "https://awx.example.com/api" \
//		-username "admin" \
//		-password "..." \
//		-ca-file "ca.pem" \
//		-id 13

package main

import (
	"flag"
	"fmt"

	"github.com/CenturyLink/hca-awx-client-go/awx"
)

var (
	url      string
	username string
	password string
	proxy    string
	insecure bool
	caFile   string
	id       int
)

func init() {
	flag.StringVar(&url, "url", "https://awx.example.com/api", "API URL.")
	flag.StringVar(&username, "username", "admin", "API user name.")
	flag.StringVar(&password, "password", "password", "API user password.")
	flag.StringVar(&proxy, "proxy", "", "API proxy URL.")
	flag.BoolVar(&insecure, "insecure", false, "Don't verify server certificate.")
	flag.StringVar(&caFile, "ca-file", "", "Trusted CA certificates.")
	flag.IntVar(&id, "id", 0, "The identifier of the project.")
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

	// Find the resource that manages the project. Note that we don't need to do a search with a
	// filter because we already know the identifier of the project, maybe because we got it in a
	// previous search.
	projectResource := connection.Projects().Id(id)

	// Send the request to retrieve the project:
	getProjectResponse, err := projectResource.Get().
		Send()
	if err != nil {
		panic(err)
	}

	// Print the results:
	project := getProjectResponse.Result()
	fmt.Printf("Id is '%d'.\n", project.Id())
	fmt.Printf("Name is '%s'.\n", project.Name())
	fmt.Printf("SCM type is '%s'.\n", project.SCMType())
	fmt.Printf("SCM URL is '%s'.\n", project.SCMURL())
	fmt.Printf("SCM branch is '%s'.\n", project.SCMBranch())
}

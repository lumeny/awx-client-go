// This example shows how to list the job templates.
//
// Use the following command to build and run it with all the debug output sent to the standard
// error output:
//
//	go run list_workflow_job_templates.go \
//		-url "https://awx.example.com/api" \
//		-authtoken "aValue" \
//		-ca-file "ca.pem"

package main

import (
	"flag"
	"fmt"

	"github.com/CenturyLink/hca-awx-client-go/awx"
)

var (
	url string
	// username string
	// password string
	// proxy    string
	insecure bool
	//
	//	caFile   string
	//
	authtoken string
)

func init() {
	flag.StringVar(&url, "url", "https://ansible-awx.rke-odc-test.corp.intranet/api", "API URL.")
	flag.StringVar(&authtoken, "authtoken", "provide-a-valid-token", "Auth token")
	// flag.StringVar(&username, "username", "admin", "API user name.")
	// flag.StringVar(&password, "password", "password", "API user password.")
	// flag.StringVar(&proxy, "proxy", "", "API proxy URL.")
	flag.BoolVar(&insecure, "insecure", true, "Don't verify server certificate.")
	// flag.StringVar(&caFile, "ca-file", "", "Trusted CA certificates.")
}

func main() {
	// Parse the command line:
	flag.Parse()

	// Connect to the server, and remember to close the connection:
	connection, err := awx.NewConnectionBuilder().
		URL(url).
		Bearer(authtoken).
		// Username(username).
		// Password(password).
		// Proxy(proxy).
		// CAFile(caFile).
		Insecure(insecure).
		Build()
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	// Find the resource that manages the collection of job templates:
	templatesResource := connection.WorkflowJobTemplates()

	// Send the request to get the list of job templates:
	getTemplatesRequest := templatesResource.Get()
	getTemplatesResponse, err := getTemplatesRequest.Send()
	if err != nil {
		panic(err)
	}

	// Print the results:
	templates := getTemplatesResponse.Results()
	for _, template := range templates {
		fmt.Printf("%d: %s - Ask Limit: %v, Ask Vars: %v\n",
			template.Id(), template.Name(),
			template.AskLimitOnLaunch(), template.AskVarsOnLaunch())
	}
}

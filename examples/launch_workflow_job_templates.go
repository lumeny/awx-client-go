// This example shows how to launch a workflow job template.
//
// Use the following command to build and run it with all the debug output sent to the standard
// error output:
//
//	go run launch_workflow_job_templates.go \
//		-url "https://awx.example.com/api" \
//		-authtoken "aValue" \
//      -project "project" \
//		-template "Echo World" \
// 		-extra-vars "cli_hostname=redhawk01-dev.corp.intranet"

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"

	awx "github.com/CenturyLink/hca-awx-client-go/awx"
)

var (
	url string
	// username      string
	// password      string
	// proxy         string
	insecure bool
	// caFile        string
	project       string
	template      string
	limit         string
	extraVarsFlag string
	insecure      bool
	authtoken     string

	//	extraVar map[string]interface{}
)

func init() {
	flag.StringVar(&url, "url", "https://ansible-awx.rke-odc-test.corp.intranet/api", "API URL.")
	flag.StringVar(&authtoken, "authtoken", "provide-a-valid-token", "Auth token")
	flag.BoolVar(&insecure, "insecure", true, "Don't verify server certificate.")
	// flag.StringVar(&project, "project", "", "Project Name.")
	flag.StringVar(&template, "template", "", "Template Name.")
	// flag.StringVar(&limit, "limit", "", "Hosts limit")
	flag.StringVar(&extraVarsFlag, "extra-vars", "", "extra variables to the Job")
}

func main() {
	// Parse the command line:
	flag.Parse()

	var extraVars map[string]interface{}
	var err error
	if len(extraVarsFlag) > 0 {
		extraVars, err = parseExtraVars(extraVarsFlag)
		if err != nil {
			fmt.Printf("Failed to parse extra-vars %s: %v\n", extraVarsFlag, err)
			return
		}
	} else {
		// create default extraVars
		extraVars = map[string]interface{}{
			"node":  "example.com",
			"count": 4,
		}
	}

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

	// Get the template by name
	templatesResource := connection.WorkflowJobTemplates()
	templatesResponse, err := templatesResource.Get().
		// Filter("project__name", project).
		Filter("name", template).
		Send()

	if err != nil {
		fmt.Printf("Failed to get workflow job template resource %v\n", err)
		return
	}

	if templatesResponse.Count() == 0 {
		fmt.Printf(
			"Template '%s' not found in project '%s'\n",
			template,
			project,
		)
		return
	}

	// TODO temp
	if len(templatesResponse.Results()) > 1 {
		log.Panicf("got %d templates by name %s", len(templatesResponse.Results()), template)
	}
	// TODO temp

	// Launch all corresponding templated
	for _, t := range templatesResponse.Results() {
		log.Printf("launching workflow job template id %d", t.Id())
		launchResource := connection.WorkflowJobTemplates().Id(t.Id()).Launch()

		if limit != "" && !t.AskLimitOnLaunch() {
			log.Printf("About to launch template '%s' with limit '%s', but 'prompt-on-launch' is false. Limit will be ignored",
				template, limit)
		}

		if extraVars != nil && !t.AskVarsOnLaunch() {
			log.Printf("About to launch template '%s' with extra-vars, but 'prompt-on-launch' is false. Extra Variables will be ignored",
				template)
		}

		response, err := launchResource.Post().
			Limit(limit).
			ExtraVars(extraVars).
			// ExtraVar("my-var", "example-val").
			Send()
		if err != nil {
			fmt.Printf("Failed to get launch job %v\n", err)
			return
		}

		log.Printf(
			"Request to launch AWX job from workflow job template '%s' has been sent, job identifier is '%v'",
			template,
			response.Job,
		)
	}
}

// Parse array of strings to extra vars json
// Expected input format: "a=b x=y c={\"v\":\"w\"}"
func parseExtraVars(input string) (output map[string]interface{}, err error) {
	variables := strings.Split(input, " ")
	if len(variables) > 0 {
		output = make(map[string]interface{})
	}
	for _, currVar := range variables {
		list := strings.SplitN(currVar, "=", 2)
		if len(list) != 2 {
			err = fmt.Errorf("bad format of extra-var")
			return
		}

		key := list[0]
		val := list[1]

		if val[0] == '{' {
			// handle complex json
			var parsedJson interface{}
			err = json.Unmarshal([]byte(val), &parsedJson)
			if err != nil {
				return
			}

			output[key] = parsedJson
		} else {
			output[key] = val
		}
	}
	return
}

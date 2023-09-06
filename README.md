# hca-awx-client-go

## History
This is a fork of the defunct https://github.com/Flamacue/awx-client-go - which in turn was a fork of the defunct https://github.com/moolitayer/awx-client-go
Team HCA needs to interact with workflow job templates, which is not supported by the original authors. 
So this fork will be the go-forward client for HCA/Lumen, with no plans to submit changes back to the original defunct projects.



A golang client library for [AWX](https://github.com/ansible/awx) and [Ansible Tower](https://www.ansible.com/products/tower) REST API.


## Usage
### import
```go
import 	"github.com/CenturyLink/hca-awx-client-go/awx"
```

### Creating a connection:
```go
// Uses the builder pattern:
connection, err := awx.NewConnectionBuilder().
  URL("http://awx.example.com/api").          // URL is mandatory
  Username(username).
  Password(password).
  Token("TOKEN").                             // Preferred approach
  Bearer("BEARER").
  CAFile("/etc/pki/tls/cert.pem").
  Insecure(insecure).                        // set to true because of self-signed certs 
  Proxy("http://myproxy.example.com").
  Build()                                    // Create the client
if err != nil {
  panic(err)
}
defer connection.Close()                      // Don't forget to close the connection!
```

`URL()` points at an AWX server's root API endpoint (including the '/api' path) and is mandatory.
`Proxy()` specifies a proxy server to use for all outgoing connection to the AWX server.
#### Authentication
Use one of:
- `Username()` and `Password()` specify Basic Auth for AWX API server.
- `Token()` uses the authtoken/ endpoint and works with AWX < 1.0.5 and Ansible tower < 3.3.
- `Bearer()` uses OAuth2 and works since AWX 1.0.5 and Ansible Tower 3.3.

When Username and Password are specified the client will attempt to acquire Token Or Bearer based on what the server supports.

#### TLS
`CAFile()` specifies path of a file containing PEM encoded CA certificates used to verify the AWX server. If no CAFile is provided, the default host trust store will be used. `CAFile()` can be used multiple times to specify a list of files.  
`Insecure(true)` can be specified to disable TLS verification.

### Supported resources
- Projects
- Jobs
- Job Templates
- Workflow Job Templates (added by Team HCA after fork)

Please submit feature requests as Github [issues](https://github.com/CenturyLink/hca-awx-client-go/issues/new).

### Retrieving resources
```go
projectsResource := connection.Projects()

// Get a list of all Projects.
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
```
#### Filtering
User `Filter()` on a request to filter lists:
```go
projectsResource := connection.Projects()

// Get a list of all Projects using git SCM.
getProjectsResponse, err := projectsResource.Get().
  Filter("scm_type", "git").
  Send()
getProjectsResponse, err := getProjectsRequest.Filter("scm_type", "git").Send()

```
#### Retrieving resource by id
Use `Id(...)` on a resource list to get a single resource
```go
// Get a resource managing a project with id=4
projectResource := connection.Projects().Id(4)

// Send the request to retrieve the project:
getProjectResponse, err := projectResource.Get().Send()
```

#### Launching a Job from a Template
```go
// Launch Job Template with id=8
launchResource := connection.JobTemplates().Id(8).Launch()

response, err := launchResource.Post().
  ExtraVars(map[string]string{"awx_environment": "staging"}).
  ExtraVar("instance", "example.com").
  Limit("master.example.com").
  Send()
if err != nil {
  return err
}
```
`ExtraVars()` Specifies a map passed to AWX as extra vars.  
`ExtraVar()` Specifies a single key value pair.  
`Limit()` is an Ansible host pattern.
See [Job Template](http://docs.ansible.com/ansible-tower/latest/html/userguide/job_templates.html)

## Examples

See [examples](examples).

## Development

### Running Tests
Install development dependencies:
```
go get golang.org/x/tools/cmd/goimports

make
```

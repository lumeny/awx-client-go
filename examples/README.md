# AWX Client Examples

This directory contains a collections of examples that show how to use the AWX
client.

## Running the Examples

In order to run the examples you will need to provide the details to connect to
your AWX server, either modifiying the source code of the example or using the
command line options. For example, to run the example that lists the job
templates you can use the following command line:

```bash
$ go run list_job_templates.go \
-url "https://awx.example.com/api" \
-username "admin" \
-password "..." \
-ca-file "ca.pem"
```

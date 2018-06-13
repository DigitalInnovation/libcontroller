# libcontroller (Alpha)
A shared library for common controller functionality. It's very rough around the edges so pls raise a github issue if you run into any gotchas :innocent:

Currently this library supports loading credentials from `/etc/kubernetes/azure.json`on cluster nodes. If an environment variable is defined in your current shell, the value read from azure.json will be overriden.

## Instructions

You install the liibrary as follows:


```bash
$ go get github.com/digitalinnovation/libcontroller
```

Then you can use it in your project as follows:

```go
package main

import ( "os"
         "fmt"
         _ "github.com/digitalinnovation/libcontroller"
)

func main() {

        fmt.Printf("AZ_TENANT_ID is: %s\n", os.Getenv("AZ_TENANT_ID"))
        fmt.Printf("AZ_CLIENT_ID is: %s\n", os.Getenv("AZ_CLIENT_ID"))
        fmt.Printf("ENVIRONMENT is: %s\n", os.Getenv("ENVIRONMENT"))
}
```

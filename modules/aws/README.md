# aws

## Enable
```
liquor app enable aws
# or
go get github.com/go-liquor/liquor-sdk/modules/aws
```

in `cmd/app/main.go` add module

```go
package main

import (
	"github.com/go-liquor/framework/internal/adapters/server/http"
	"github.com/go-liquor/framework/internal/app/services"
	"github.com/go-liquor/liquor-sdk/app"
    "github.com/go-liquor/liquor-sdk/modules/aws" // add this
)

func main() {
	app.NewApp(
        aws.AwsClientModule, // add this
		http.Server,
		app.RegisterService(
			services.NewInitialService,
		),
	)
}
```

### Aws Services

- `*dynamodb.Client`
- `*kms.Client`
- `*s3.Client`
- `*sqs.Client`

### How can I use it?

In your service you can use it like this:

```go
type Service struct {
	client *s3.Client
}

func NewService(client *s3.Client) *Service {
	return &Service{
		client: client,
	}
}
```

The `*s3.Client` will be injected by the framework.
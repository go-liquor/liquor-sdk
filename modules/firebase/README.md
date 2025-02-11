# Firebase Module

Firebase integration module for Go applications.

## Installation

```bash
liquor app enable firebase
# or
go get github.com/go-liquor/liquor-sdk/modules/firebase
```

## Configuration

In your configuration file:
```yaml
firebase:
  configFile: "path/to/firebase-credentials.json"
```

Add in your main.go:

```go
package main

import (
	"github.com/go-liquor/framework/internal/adapters/server/http"
	"github.com/go-liquor/framework/internal/app/services"
	"github.com/go-liquor/liquor-sdk/app"
    "github.com/go-liquor/liquor-sdk/modules/firebase" // add this
)

func main() {
	app.NewApp(
        firebase.FirebaseModule, // add this
		http.Server,
		app.RegisterService(
			services.NewInitialService,
		),
	)
}
```

## Usage

### Authentication

```go

type Service struct {
    client *firebase.Auth
}

func NewService(client *firebase.Auth) *Service {
    return &Service{client: client}
}

func (s *Service) Run(ctx context.Context, idToken string) {
   token, err := s.client.VerifyIDToken(ctx, idToken)
    if err != nil {
        // Handle error
    }
}
```

### Firestore Database

```go

type Service struct {
    client *firebase.FirestoreClient
}

func NewService(client *firebase.FirestoreClient) *Service {
    return &Service{client: client}
}

func (s *Service) Run(ctx context.Context) {
    // Use Firestore client
    doc := s.client.Collection("users").Doc("user1")
}
```

## Features

- Firebase App initialization
- Authentication client
- Firestore database client




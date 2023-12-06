# Background

GrowthBook provides the [OpenAPI 3.1.0 specification](https://api.growthbook.io/api/v1/openapi.yaml) of their API. However, it is not trivial to immediately use it to auto-generate a Go client from this spec. There are two reasons for this. a) the spec file provided is, at the moment of creating this repository, malformed (see partially corrected file in `growthbook-openapi.yaml`) and b) from the various client generators available publicly, not all of them support 3.1.0 version of the OpenAPI spec, or do it only partially.

Some of the errors in their spec include misspelling keyword `object` as `obect`, missing `description` field in their HTTP expected responses and others.

Therefore I created this repository where I used [OpenAPI](https://openapi-generator.tech) tools to auto-generate most of the code. I however needed to make some amendments. For example, OpenAPI tools do not support `AnyOf` clause for Go clients, also the tool dumps all files in the same folder and Go package, so I split out the models in a separate folder and Go package for clarity.

# Motivation

At the moment of creating this repository there is no Golang client/SDK to interact with the complete API of GrowthBook. There is an [SDK](https://docs.growthbook.io/lib/go) provided by GrowthBook to access just one endpoint (to read features). However, the rest of the API is not accessible via the aforementioned SDK.

One use case to have a Go client to access GrowthBook API is to be able to implement a GrowthBook Terraform provider - which at the moment of writing creating this repository, is inexistent.

# How to install
This SDK requires Go >= 1.21
You can install it by issuing the classic
```
go get github.com/JoseFMP/gogrowthbook
```

As well as `go install github.com/JoseFMP/gogrowthbook` if required.

# How to use

As the code is mostly autogenerated with OpenAPI tools, it is mostly very consistent as one model and one interface per endpoint of the API.

Here a toy example:
```go
package main

import "github.com/JoseFMP/gogrowthbook"

func main() {
  cfg := gogrowthbook.NewConfiguration()
  
  // initialization of the client. Nothing really going on here.
  client := gogrowthbook.NewAPIClient(cfg)

  // secretKey can be generated in GrowthBook UI. Personal Access Token works too.
  // Depending on the operations you plan to do (read only or write) you might require 'admin' permission on the key, or just 'readonly'
  // Note: Obviously this key is sensitive.
  secretKey := "secret_key_blabla"
  clientContext := context.Background()

  // sets the secret key in to the context, this way the client can use it
  // Note: we could use different keys for different requests using the very same client
  clientContext = context.WithValue(clientContext, gogrowthbook.ContextAccessToken, secretKey)

  exampleNewFeature := models.PostFeatureRequest{
    Id:          "example-put-here-unique-feature-id",
    ValueType:   "string",
    Description: "This is an example feature, ignore it!",
  }
  newFeatureReq := client.FeaturesAPI.PostFeature(clientContext).PostFeatureRequest(exampleNewFeature)

  responsePayload, httpResponse, errDoingReq := newFeatureReq.Execute()

  fmt.Printf("Creating new feature result: %d, %s, %v\n%v", httpResponse.StatusCode, httpResponse.Status, errDoingReq, responsePayload)
}
```
You can see a fully working integration test in [client_test.go](./client_test.go)

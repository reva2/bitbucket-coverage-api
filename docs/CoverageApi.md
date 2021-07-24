# \CoverageApi

All URIs are relative to *http://localhost/rest/code-coverage/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](CoverageApi.md#Delete) | **Delete** /commits/{commitId} | 
[**Get**](CoverageApi.md#Get) | **Get** /commits/{commitId} | 
[**Post**](CoverageApi.md#Post) | **Post** /commits/{commitId} | 



## Delete

> Delete(ctx, commitId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    commitId := "commitId_example" // string | The commit ID. This must be a full 40 character commit hash

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoverageApi.Delete(context.Background(), commitId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoverageApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commitId** | **string** | The commit ID. This must be a full 40 character commit hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> CommitCoverage Get(ctx, commitId).Path(path).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    commitId := "commitId_example" // string | The commit ID. This must be a full 40 character commit hash
    path := "path_example" // string | This is the path of the file relative to the git repository

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoverageApi.Get(context.Background(), commitId).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoverageApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: CommitCoverage
    fmt.Fprintf(os.Stdout, "Response from `CoverageApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commitId** | **string** | The commit ID. This must be a full 40 character commit hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | This is the path of the file relative to the git repository | 

### Return type

[**CommitCoverage**](CommitCoverage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CommitCoverage Post(ctx, commitId).CommitCoverage(commitCoverage).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    commitId := "commitId_example" // string | The commit ID. This must be a full 40 character commit hash
    commitCoverage := *openapiclient.NewCommitCoverage([]openapiclient.FileCoverage{*openapiclient.NewFileCoverage("Path_example", "Coverage_example")}) // CommitCoverage | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoverageApi.Post(context.Background(), commitId).CommitCoverage(commitCoverage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoverageApi.Post``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Post`: CommitCoverage
    fmt.Fprintf(os.Stdout, "Response from `CoverageApi.Post`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commitId** | **string** | The commit ID. This must be a full 40 character commit hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commitCoverage** | [**CommitCoverage**](CommitCoverage.md) |  | 

### Return type

[**CommitCoverage**](CommitCoverage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


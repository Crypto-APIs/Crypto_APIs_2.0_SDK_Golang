# \DefaultApi

All URIs are relative to *https://rest.cryptoapis.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContractDetailsByAddress**](DefaultApi.md#GetContractDetailsByAddress) | **Get** /blockchain-data/{blockchain}/{network}/addresses/{contractAddress}/contract | Get Contract Details by Address



## GetContractDetailsByAddress

> GetContractDetailsByAddressResponse GetContractDetailsByAddress(ctx, blockchain, network, contractAddress).Context(context).Execute()

Get Contract Details by Address



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
    blockchain := "blockchain_example" // string | 
    network := "network_example" // string | 
    contractAddress := "contractAddress_example" // string | String identifier of the token
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetContractDetailsByAddress(context.Background(), blockchain, network, contractAddress).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetContractDetailsByAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractDetailsByAddress`: GetContractDetailsByAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetContractDetailsByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** |  | 
**network** | **string** |  | 
**contractAddress** | **string** | String identifier of the token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractDetailsByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetContractDetailsByAddressResponse**](GetContractDetailsByAddressResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


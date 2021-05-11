# \ValidatingApi

All URIs are relative to *https://rest.cryptoapis.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateAddress**](ValidatingApi.md#ValidateAddress) | **Post** /blockchain-tools/{blockchain}/{network}/addresses/validate | Validate Address



## ValidateAddress

> ValidateAddressResponse ValidateAddress(ctx, blockchain, network).Context(context).ValidateAddressRequestBody(validateAddressRequestBody).Execute()

Validate Address



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
    blockchain := "bitcoin" // string | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    validateAddressRequestBody := *openapiclient.NewValidateAddressRequestBody(*openapiclient.NewValidateAddressRequestBodyData(*openapiclient.NewValidateAddressRequestBodyDataItem("mho4jHBcrNCncKt38trJahXakuaBnS7LK5"))) // ValidateAddressRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ValidatingApi.ValidateAddress(context.Background(), blockchain, network).Context(context).ValidateAddressRequestBody(validateAddressRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ValidatingApi.ValidateAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateAddress`: ValidateAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `ValidatingApi.ValidateAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **validateAddressRequestBody** | [**ValidateAddressRequestBody**](ValidateAddressRequestBody.md) |  | 

### Return type

[**ValidateAddressResponse**](ValidateAddressResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \FeaturesApi

All URIs are relative to *https://rest.cryptoapis.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BroadcastLocallySignedTransaction**](FeaturesApi.md#BroadcastLocallySignedTransaction) | **Post** /blockchain-tools/{blockchain}/{network}/transactions/broadcast | Broadcast Locally Signed Transaction
[**GenerateAddress**](FeaturesApi.md#GenerateAddress) | **Post** /blockchain-tools/{blockchain}/{network}/addresses/generate | Generate Address
[**GetEIP1559FeeRecommendations**](FeaturesApi.md#GetEIP1559FeeRecommendations) | **Get** /blockchain-tools/{blockchain}/{network}/fees/eip1559 | Get EIP 1559 Fee Recommendations
[**ValidateAddress**](FeaturesApi.md#ValidateAddress) | **Post** /blockchain-tools/{blockchain}/{network}/addresses/validate | Validate Address



## BroadcastLocallySignedTransaction

> BroadcastLocallySignedTransactionR BroadcastLocallySignedTransaction(ctx, blockchain, network).Context(context).BroadcastLocallySignedTransactionRB(broadcastLocallySignedTransactionRB).Execute()

Broadcast Locally Signed Transaction



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
    network := "network_example" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    broadcastLocallySignedTransactionRB := *openapiclient.NewBroadcastLocallySignedTransactionRB(*openapiclient.NewBroadcastLocallySignedTransactionRBData(*openapiclient.NewBroadcastLocallySignedTransactionRBDataItem("https://example.com", "0xf86a22827d00831e8480941b85a43e2e7f52e766ddfdfa2b901c42cb1201be8801b27f33b807c0008029a084ccbf02b27e0842fb1eda7a187a5589c3759be0e969e0ca989dc469a5e5e394a02e111e1156b197f1de4c1d9ba4af26e50665ea6d617d05b3e4047da12b915e69"))) // BroadcastLocallySignedTransactionRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeaturesApi.BroadcastLocallySignedTransaction(context.Background(), blockchain, network).Context(context).BroadcastLocallySignedTransactionRB(broadcastLocallySignedTransactionRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.BroadcastLocallySignedTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BroadcastLocallySignedTransaction`: BroadcastLocallySignedTransactionR
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.BroadcastLocallySignedTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastLocallySignedTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **broadcastLocallySignedTransactionRB** | [**BroadcastLocallySignedTransactionRB**](BroadcastLocallySignedTransactionRB.md) |  | 

### Return type

[**BroadcastLocallySignedTransactionR**](BroadcastLocallySignedTransactionR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateAddress

> GenerateAddressR GenerateAddress(ctx, blockchain, network).Context(context).GenerateAddressRB(generateAddressRB).Execute()

Generate Address



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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    generateAddressRB := *openapiclient.NewGenerateAddressRB(*openapiclient.NewGenerateAddressRBData(map[string]interface{}(123))) // GenerateAddressRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeaturesApi.GenerateAddress(context.Background(), blockchain, network).Context(context).GenerateAddressRB(generateAddressRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.GenerateAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateAddress`: GenerateAddressR
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.GenerateAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **generateAddressRB** | [**GenerateAddressRB**](GenerateAddressRB.md) |  | 

### Return type

[**GenerateAddressR**](GenerateAddressR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEIP1559FeeRecommendations

> GetEIP1559FeeRecommendationsR GetEIP1559FeeRecommendations(ctx, network, blockchain).Context(context).Execute()

Get EIP 1559 Fee Recommendations



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
    network := "ropsten" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    blockchain := "ethereum" // string | Represents the specific blockchain protocol name, e.g. Ethereum.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeaturesApi.GetEIP1559FeeRecommendations(context.Background(), network, blockchain).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.GetEIP1559FeeRecommendations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEIP1559FeeRecommendations`: GetEIP1559FeeRecommendationsR
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.GetEIP1559FeeRecommendations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEIP1559FeeRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetEIP1559FeeRecommendationsR**](GetEIP1559FeeRecommendationsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAddress

> ValidateAddressR ValidateAddress(ctx, blockchain, network).Context(context).ValidateAddressRB(validateAddressRB).Execute()

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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    validateAddressRB := *openapiclient.NewValidateAddressRB(*openapiclient.NewValidateAddressRBData(*openapiclient.NewValidateAddressRBDataItem("mho4jHBcrNCncKt38trJahXakuaBnS7LK5"))) // ValidateAddressRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeaturesApi.ValidateAddress(context.Background(), blockchain, network).Context(context).ValidateAddressRB(validateAddressRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.ValidateAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateAddress`: ValidateAddressR
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.ValidateAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **validateAddressRB** | [**ValidateAddressRB**](ValidateAddressRB.md) |  | 

### Return type

[**ValidateAddressR**](ValidateAddressR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


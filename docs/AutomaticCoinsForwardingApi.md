# \AutomaticCoinsForwardingApi

All URIs are relative to *https://rest.cryptoapis.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutomaticCoinsForwarding**](AutomaticCoinsForwardingApi.md#CreateAutomaticCoinsForwarding) | **Post** /blockchain-automations/{blockchain}/{network}/coins-forwarding/automations | Create Automatic Coins Forwarding
[**DeleteAutomaticCoinsForwarding**](AutomaticCoinsForwardingApi.md#DeleteAutomaticCoinsForwarding) | **Delete** /blockchain-automations/{blockchain}/{network}/coins-forwarding/automations/{referenceId} | Delete Automatic Coins Forwarding
[**ListCoinsForwardingAutomations**](AutomaticCoinsForwardingApi.md#ListCoinsForwardingAutomations) | **Get** /blockchain-automations/{blockchain}/{network}/coins-forwarding/automations | List Coins Forwarding Automations



## CreateAutomaticCoinsForwarding

> CreateAutomaticCoinsForwardingR CreateAutomaticCoinsForwarding(ctx, blockchain, network).Context(context).CreateAutomaticCoinsForwardingRB(createAutomaticCoinsForwardingRB).Execute()

Create Automatic Coins Forwarding



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
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    createAutomaticCoinsForwardingRB := *openapiclient.NewCreateAutomaticCoinsForwardingRB(*openapiclient.NewCreateAutomaticCoinsForwardingRBData(*openapiclient.NewCreateAutomaticCoinsForwardingRBDataItem("yourSecretString", "https://example.com", int32(3), "standard", "0.02", "mzYijhgmzZrmuB7wBDazRKirnChKyow4M3"))) // CreateAutomaticCoinsForwardingRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutomaticCoinsForwardingApi.CreateAutomaticCoinsForwarding(context.Background(), blockchain, network).Context(context).CreateAutomaticCoinsForwardingRB(createAutomaticCoinsForwardingRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticCoinsForwardingApi.CreateAutomaticCoinsForwarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAutomaticCoinsForwarding`: CreateAutomaticCoinsForwardingR
    fmt.Fprintf(os.Stdout, "Response from `AutomaticCoinsForwardingApi.CreateAutomaticCoinsForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutomaticCoinsForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **createAutomaticCoinsForwardingRB** | [**CreateAutomaticCoinsForwardingRB**](CreateAutomaticCoinsForwardingRB.md) |  | 

### Return type

[**CreateAutomaticCoinsForwardingR**](CreateAutomaticCoinsForwardingR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutomaticCoinsForwarding

> DeleteAutomaticCoinsForwardingR DeleteAutomaticCoinsForwarding(ctx, blockchain, network, referenceId).Context(context).Execute()

Delete Automatic Coins Forwarding



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
    referenceId := "600955ea5e75d660e71d3c7d" // string | Represents a unique ID used to reference the specific callback subscription.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutomaticCoinsForwardingApi.DeleteAutomaticCoinsForwarding(context.Background(), blockchain, network, referenceId).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticCoinsForwardingApi.DeleteAutomaticCoinsForwarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAutomaticCoinsForwarding`: DeleteAutomaticCoinsForwardingR
    fmt.Fprintf(os.Stdout, "Response from `AutomaticCoinsForwardingApi.DeleteAutomaticCoinsForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**referenceId** | **string** | Represents a unique ID used to reference the specific callback subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutomaticCoinsForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**DeleteAutomaticCoinsForwardingR**](DeleteAutomaticCoinsForwardingR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCoinsForwardingAutomations

> ListCoinsForwardingAutomationsR ListCoinsForwardingAutomations(ctx, blockchain, network).Context(context).Limit(limit).Offset(offset).Execute()

List Coins Forwarding Automations



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
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int64(50) // int64 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int64(0) // int64 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutomaticCoinsForwardingApi.ListCoinsForwardingAutomations(context.Background(), blockchain, network).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticCoinsForwardingApi.ListCoinsForwardingAutomations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCoinsForwardingAutomations`: ListCoinsForwardingAutomationsR
    fmt.Fprintf(os.Stdout, "Response from `AutomaticCoinsForwardingApi.ListCoinsForwardingAutomations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCoinsForwardingAutomationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int64** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int64** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListCoinsForwardingAutomationsR**](ListCoinsForwardingAutomationsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


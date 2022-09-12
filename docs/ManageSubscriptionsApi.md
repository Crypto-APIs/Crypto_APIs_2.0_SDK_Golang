# \ManageSubscriptionsApi

All URIs are relative to *https://rest.cryptoapis.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateBlockchainEventSubscription**](ManageSubscriptionsApi.md#ActivateBlockchainEventSubscription) | **Post** /blockchain-events/subscriptions/{referenceId}/activate | Activate Blockchain Event Subscription
[**DeleteBlockchainEventSubscription**](ManageSubscriptionsApi.md#DeleteBlockchainEventSubscription) | **Delete** /blockchain-events/{blockchain}/{network}/subscriptions/{referenceId} | Delete Blockchain Event Subscription
[**GetBlockchainEventSubscriptionDetailsByReferenceID**](ManageSubscriptionsApi.md#GetBlockchainEventSubscriptionDetailsByReferenceID) | **Get** /blockchain-events/subscriptions/{referenceId} | Get Blockchain Event Subscription Details By Reference ID
[**ListBlockchainEventsSubscriptions**](ManageSubscriptionsApi.md#ListBlockchainEventsSubscriptions) | **Get** /blockchain-events/{blockchain}/{network}/subscriptions | List Blockchain Events Subscriptions



## ActivateBlockchainEventSubscription

> ActivateBlockchainEventSubscriptionR ActivateBlockchainEventSubscription(ctx, referenceId).Context(context).ActivateBlockchainEventSubscriptionRB(activateBlockchainEventSubscriptionRB).Execute()

Activate Blockchain Event Subscription



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
    referenceId := "bc243c86-0902-4386-b30d-e6b30fa1f2aa" // string | Represents a unique ID used to reference the specific callback subscription.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    activateBlockchainEventSubscriptionRB := *openapiclient.NewActivateBlockchainEventSubscriptionRB(*openapiclient.NewActivateBlockchainEventSubscriptionRBData(map[string]interface{}(123))) // ActivateBlockchainEventSubscriptionRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManageSubscriptionsApi.ActivateBlockchainEventSubscription(context.Background(), referenceId).Context(context).ActivateBlockchainEventSubscriptionRB(activateBlockchainEventSubscriptionRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageSubscriptionsApi.ActivateBlockchainEventSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateBlockchainEventSubscription`: ActivateBlockchainEventSubscriptionR
    fmt.Fprintf(os.Stdout, "Response from `ManageSubscriptionsApi.ActivateBlockchainEventSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**referenceId** | **string** | Represents a unique ID used to reference the specific callback subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateBlockchainEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **activateBlockchainEventSubscriptionRB** | [**ActivateBlockchainEventSubscriptionRB**](ActivateBlockchainEventSubscriptionRB.md) |  | 

### Return type

[**ActivateBlockchainEventSubscriptionR**](ActivateBlockchainEventSubscriptionR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlockchainEventSubscription

> DeleteBlockchainEventSubscriptionR DeleteBlockchainEventSubscription(ctx, blockchain, network, referenceId).Context(context).Execute()

Delete Blockchain Event Subscription



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
    referenceId := "d3fd6a0e-f2b6-4bb5-9fd3-7944bcec9e9f" // string | Represents a unique ID used to reference the specific callback subscription.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManageSubscriptionsApi.DeleteBlockchainEventSubscription(context.Background(), blockchain, network, referenceId).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageSubscriptionsApi.DeleteBlockchainEventSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBlockchainEventSubscription`: DeleteBlockchainEventSubscriptionR
    fmt.Fprintf(os.Stdout, "Response from `ManageSubscriptionsApi.DeleteBlockchainEventSubscription`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteBlockchainEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**DeleteBlockchainEventSubscriptionR**](DeleteBlockchainEventSubscriptionR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockchainEventSubscriptionDetailsByReferenceID

> GetBlockchainEventSubscriptionDetailsByReferenceIDR GetBlockchainEventSubscriptionDetailsByReferenceID(ctx, referenceId).Context(context).Execute()

Get Blockchain Event Subscription Details By Reference ID



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
    referenceId := "bc243c86-0902-4386-b30d-e6b30fa1f2aa" // string | Represents a unique ID used to reference the specific callback subscription.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManageSubscriptionsApi.GetBlockchainEventSubscriptionDetailsByReferenceID(context.Background(), referenceId).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageSubscriptionsApi.GetBlockchainEventSubscriptionDetailsByReferenceID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockchainEventSubscriptionDetailsByReferenceID`: GetBlockchainEventSubscriptionDetailsByReferenceIDR
    fmt.Fprintf(os.Stdout, "Response from `ManageSubscriptionsApi.GetBlockchainEventSubscriptionDetailsByReferenceID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**referenceId** | **string** | Represents a unique ID used to reference the specific callback subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockchainEventSubscriptionDetailsByReferenceIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetBlockchainEventSubscriptionDetailsByReferenceIDR**](GetBlockchainEventSubscriptionDetailsByReferenceIDR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlockchainEventsSubscriptions

> ListBlockchainEventsSubscriptionsR ListBlockchainEventsSubscriptions(ctx, blockchain, network).Context(context).Limit(limit).Offset(offset).Execute()

List Blockchain Events Subscriptions



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
    resp, r, err := apiClient.ManageSubscriptionsApi.ListBlockchainEventsSubscriptions(context.Background(), blockchain, network).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageSubscriptionsApi.ListBlockchainEventsSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBlockchainEventsSubscriptions`: ListBlockchainEventsSubscriptionsR
    fmt.Fprintf(os.Stdout, "Response from `ManageSubscriptionsApi.ListBlockchainEventsSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBlockchainEventsSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int64** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int64** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListBlockchainEventsSubscriptionsR**](ListBlockchainEventsSubscriptionsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


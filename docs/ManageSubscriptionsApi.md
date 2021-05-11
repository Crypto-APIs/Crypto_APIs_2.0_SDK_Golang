# \ManageSubscriptionsApi

All URIs are relative to *https://rest.cryptoapis.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBlockchainEventSubscription**](ManageSubscriptionsApi.md#DeleteBlockchainEventSubscription) | **Delete** /blockchain-events/{blockchain}/{network}/subscriptions/{referenceId} | Delete Blockchain Event Subscription
[**ListBlockchainEventsSubscriptions**](ManageSubscriptionsApi.md#ListBlockchainEventsSubscriptions) | **Get** /blockchain-events/{blockchain}/{network}/subscriptions | List Blockchain Events Subscriptions



## DeleteBlockchainEventSubscription

> DeleteBlockchainEventSubscriptionResponse DeleteBlockchainEventSubscription(ctx, blockchain, network, referenceId).Context(context).Execute()

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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
    referenceId := "d3fd6a0e-f2b6-4bb5-9fd3-7944bcec9e9f" // string | Represents a unique ID used to reference the specific callback subscription.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManageSubscriptionsApi.DeleteBlockchainEventSubscription(context.Background(), blockchain, network, referenceId).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageSubscriptionsApi.DeleteBlockchainEventSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBlockchainEventSubscription`: DeleteBlockchainEventSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageSubscriptionsApi.DeleteBlockchainEventSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**referenceId** | **string** | Represents a unique ID used to reference the specific callback subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlockchainEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**DeleteBlockchainEventSubscriptionResponse**](DeleteBlockchainEventSubscriptionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlockchainEventsSubscriptions

> ListBlockchainEventsSubscriptionsResponse ListBlockchainEventsSubscriptions(ctx, blockchain, network).Context(context).Limit(limit).Offset(offset).Execute()

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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int32(50) // int32 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int32(10) // int32 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManageSubscriptionsApi.ListBlockchainEventsSubscriptions(context.Background(), blockchain, network).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageSubscriptionsApi.ListBlockchainEventsSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBlockchainEventsSubscriptions`: ListBlockchainEventsSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageSubscriptionsApi.ListBlockchainEventsSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBlockchainEventsSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int32** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListBlockchainEventsSubscriptionsResponse**](ListBlockchainEventsSubscriptionsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


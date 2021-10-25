# \AutomaticTokensForwardingApi

All URIs are relative to *https://rest.cryptoapis.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTokensToExistingFromAddress**](AutomaticTokensForwardingApi.md#AddTokensToExistingFromAddress) | **Post** /blockchain-automations/{blockchain}/{network}/tokens-forwarding/automations/add-token | Add Tokens To Existing fromAddress
[**CreateAutomaticTokensForwarding**](AutomaticTokensForwardingApi.md#CreateAutomaticTokensForwarding) | **Post** /blockchain-automations/{blockchain}/{network}/tokens-forwarding/automations | Create Automatic Tokens Forwarding
[**DeleteAutomaticTokensForwarding**](AutomaticTokensForwardingApi.md#DeleteAutomaticTokensForwarding) | **Delete** /blockchain-automations/{blockchain}/{network}/tokens-forwarding/automations/{referenceId} | Delete Automatic Tokens Forwarding
[**GetFeeAddressDetails**](AutomaticTokensForwardingApi.md#GetFeeAddressDetails) | **Get** /blockchain-automations/{blockchain}/{network}/tokens-forwarding/fee-addresses | Get Fee Address Details
[**ListTokensForwardingAutomations**](AutomaticTokensForwardingApi.md#ListTokensForwardingAutomations) | **Get** /blockchain-automations/{blockchain}/{network}/tokens-forwarding/automations | List Tokens Forwarding Automations



## AddTokensToExistingFromAddress

> AddTokensToExistingFromAddressR AddTokensToExistingFromAddress(ctx, blockchain, network).Context(context).AddTokensToExistingFromAddressRB(addTokensToExistingFromAddressRB).Execute()

Add Tokens To Existing fromAddress



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
    addTokensToExistingFromAddressRB := *openapiclient.NewAddTokensToExistingFromAddressRB(*openapiclient.NewAddTokensToExistingFromAddressRBData(*openapiclient.NewAddTokensToExistingFromAddressRBDataItem("https://example.com", int32(3), "standard", "mizRduUBKEbJ6uzYJUegPh78gEGgM3WjAr", "0.00001", "mnumE76iEKN47bUsdni85oped5D1fRwKWi", openapiclient.AddTokensToExistingFromAddressRBTokenData{AddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken: openapiclient.NewAddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken(int32(2))}))) // AddTokensToExistingFromAddressRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomaticTokensForwardingApi.AddTokensToExistingFromAddress(context.Background(), blockchain, network).Context(context).AddTokensToExistingFromAddressRB(addTokensToExistingFromAddressRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticTokensForwardingApi.AddTokensToExistingFromAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTokensToExistingFromAddress`: AddTokensToExistingFromAddressR
    fmt.Fprintf(os.Stdout, "Response from `AutomaticTokensForwardingApi.AddTokensToExistingFromAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTokensToExistingFromAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **addTokensToExistingFromAddressRB** | [**AddTokensToExistingFromAddressRB**](AddTokensToExistingFromAddressRB.md) |  | 

### Return type

[**AddTokensToExistingFromAddressR**](AddTokensToExistingFromAddressR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAutomaticTokensForwarding

> CreateAutomaticTokensForwardingR CreateAutomaticTokensForwarding(ctx, blockchain, network).Context(context).CreateAutomaticTokensForwardingRB(createAutomaticTokensForwardingRB).Execute()

Create Automatic Tokens Forwarding



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
    createAutomaticTokensForwardingRB := *openapiclient.NewCreateAutomaticTokensForwardingRB(*openapiclient.NewCreateAutomaticTokensForwardingRBData(*openapiclient.NewCreateAutomaticTokensForwardingRBDataItem("https://example.com", "3", "standard", "0.00002", "tb1q54j7qcu7kgsrx87yn0r9zjdvsxrnvxg4qua2z6", openapiclient.CreateAutomaticTokensForwardingRBTokenData{CreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken: openapiclient.NewCreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken(int32(31))}))) // CreateAutomaticTokensForwardingRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomaticTokensForwardingApi.CreateAutomaticTokensForwarding(context.Background(), blockchain, network).Context(context).CreateAutomaticTokensForwardingRB(createAutomaticTokensForwardingRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticTokensForwardingApi.CreateAutomaticTokensForwarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAutomaticTokensForwarding`: CreateAutomaticTokensForwardingR
    fmt.Fprintf(os.Stdout, "Response from `AutomaticTokensForwardingApi.CreateAutomaticTokensForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutomaticTokensForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **createAutomaticTokensForwardingRB** | [**CreateAutomaticTokensForwardingRB**](CreateAutomaticTokensForwardingRB.md) |  | 

### Return type

[**CreateAutomaticTokensForwardingR**](CreateAutomaticTokensForwardingR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutomaticTokensForwarding

> DeleteAutomaticTokensForwardingR DeleteAutomaticTokensForwarding(ctx, blockchain, network, referenceId).Context(context).Execute()

Delete Automatic Tokens Forwarding



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
    referenceId := "6017dd02a309213863be9e55" // string | Represents a unique ID used to reference the specific callback subscription.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomaticTokensForwardingApi.DeleteAutomaticTokensForwarding(context.Background(), blockchain, network, referenceId).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticTokensForwardingApi.DeleteAutomaticTokensForwarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAutomaticTokensForwarding`: DeleteAutomaticTokensForwardingR
    fmt.Fprintf(os.Stdout, "Response from `AutomaticTokensForwardingApi.DeleteAutomaticTokensForwarding`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteAutomaticTokensForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**DeleteAutomaticTokensForwardingR**](DeleteAutomaticTokensForwardingR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeeAddressDetails

> GetFeeAddressDetailsR GetFeeAddressDetails(ctx, blockchain, network).Context(context).Execute()

Get Fee Address Details



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomaticTokensForwardingApi.GetFeeAddressDetails(context.Background(), blockchain, network).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticTokensForwardingApi.GetFeeAddressDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeeAddressDetails`: GetFeeAddressDetailsR
    fmt.Fprintf(os.Stdout, "Response from `AutomaticTokensForwardingApi.GetFeeAddressDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeeAddressDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetFeeAddressDetailsR**](GetFeeAddressDetailsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokensForwardingAutomations

> ListTokensForwardingAutomationsR ListTokensForwardingAutomations(ctx, blockchain, network).Context(context).Limit(limit).Offset(offset).Execute()

List Tokens Forwarding Automations



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
    limit := int32(50) // int32 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int32(10) // int32 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomaticTokensForwardingApi.ListTokensForwardingAutomations(context.Background(), blockchain, network).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticTokensForwardingApi.ListTokensForwardingAutomations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTokensForwardingAutomations`: ListTokensForwardingAutomationsR
    fmt.Fprintf(os.Stdout, "Response from `AutomaticTokensForwardingApi.ListTokensForwardingAutomations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTokensForwardingAutomationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int32** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListTokensForwardingAutomationsR**](ListTokensForwardingAutomationsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


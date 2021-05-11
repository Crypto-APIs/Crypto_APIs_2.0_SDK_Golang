# \CreateSubscriptionsForApi

All URIs are relative to *https://rest.cryptoapis.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MinedTransaction**](CreateSubscriptionsForApi.md#MinedTransaction) | **Post** /blockchain-events/{blockchain}/{network}/subscriptions/transaction-mined | Mined transaction
[**NewBlock**](CreateSubscriptionsForApi.md#NewBlock) | **Post** /blockchain-events/{blockchain}/{network}/subscriptions/block-mined | New Block
[**NewConfirmedCoinsTransactions**](CreateSubscriptionsForApi.md#NewConfirmedCoinsTransactions) | **Post** /blockchain-events/{blockchain}/{network}/subscriptions/address-coins-transactions-confirmed | New confirmed coins transactions
[**NewConfirmedCoinsTransactionsAndEachConfirmation**](CreateSubscriptionsForApi.md#NewConfirmedCoinsTransactionsAndEachConfirmation) | **Post** /blockchain-events/{blockchain}/{network}/subscriptions/address-coins-transactions-confirmed-each-confirmation | New confirmed coins transactions and each confirmation
[**NewConfirmedTokensTransactions**](CreateSubscriptionsForApi.md#NewConfirmedTokensTransactions) | **Post** /blockchain-events/{blockchain}/{network}/subscriptions/address-tokens-transactions-confirmed | New confirmed tokens transactions
[**NewConfirmedTokensTransactionsAndEachConfirmation**](CreateSubscriptionsForApi.md#NewConfirmedTokensTransactionsAndEachConfirmation) | **Post** /blockchain-events/{blockchain}/{network}/subscriptions/address-tokens-transactions-confirmed-each-confirmation | New confirmed tokens transactions and each confirmation
[**NewUnconfirmedCoinsTransactions**](CreateSubscriptionsForApi.md#NewUnconfirmedCoinsTransactions) | **Post** /blockchain-events/{blockchain}/{network}/subscriptions/address-coins-transactions-unconfirmed | New unconfirmed coins transactions
[**NewUnconfirmedTokensTransactions**](CreateSubscriptionsForApi.md#NewUnconfirmedTokensTransactions) | **Post** /blockchain-events/{blockchain}/{network}/subscriptions/address-tokens-transactions-unconfirmed | New unconfirmed tokens transactions



## MinedTransaction

> MinedTransactionResponse MinedTransaction(ctx, blockchain, network).Context(context).MinedTransactionRequestBody(minedTransactionRequestBody).Execute()

Mined transaction



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
    minedTransactionRequestBody := *openapiclient.NewMinedTransactionRequestBody(*openapiclient.NewMinedTransactionRequestBodyData(*openapiclient.NewMinedTransactionRequestBodyDataItem("http://example.com", "df2690ff97e72c1f8b0f2102a8cb5c1d0fa8fb8754d543c9bc0edc4d4bc34bfc"))) // MinedTransactionRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CreateSubscriptionsForApi.MinedTransaction(context.Background(), blockchain, network).Context(context).MinedTransactionRequestBody(minedTransactionRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateSubscriptionsForApi.MinedTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MinedTransaction`: MinedTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `CreateSubscriptionsForApi.MinedTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMinedTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **minedTransactionRequestBody** | [**MinedTransactionRequestBody**](MinedTransactionRequestBody.md) |  | 

### Return type

[**MinedTransactionResponse**](MinedTransactionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewBlock

> NewBlockResponse NewBlock(ctx, blockchain, network).Context(context).NewBlockRequestBody(newBlockRequestBody).Execute()

New Block



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
    newBlockRequestBody := *openapiclient.NewNewBlockRequestBody(*openapiclient.NewNewBlockRequestBodyData(*openapiclient.NewNewBlockRequestBodyDataItem("http://example.com"))) // NewBlockRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CreateSubscriptionsForApi.NewBlock(context.Background(), blockchain, network).Context(context).NewBlockRequestBody(newBlockRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateSubscriptionsForApi.NewBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewBlock`: NewBlockResponse
    fmt.Fprintf(os.Stdout, "Response from `CreateSubscriptionsForApi.NewBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **newBlockRequestBody** | [**NewBlockRequestBody**](NewBlockRequestBody.md) |  | 

### Return type

[**NewBlockResponse**](NewBlockResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewConfirmedCoinsTransactions

> NewConfirmedCoinsTransactionsResponse NewConfirmedCoinsTransactions(ctx, blockchain, network).Context(context).NewConfirmedCoinsTransactionsRequestBody(newConfirmedCoinsTransactionsRequestBody).Execute()

New confirmed coins transactions



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
    newConfirmedCoinsTransactionsRequestBody := *openapiclient.NewNewConfirmedCoinsTransactionsRequestBody(*openapiclient.NewNewConfirmedCoinsTransactionsRequestBodyData(*openapiclient.NewNewConfirmedCoinsTransactionsRequestBodyDataItem("mho4jHBcrNCncKt38trJahXakuaBnS7LK5", "http://example.com"))) // NewConfirmedCoinsTransactionsRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CreateSubscriptionsForApi.NewConfirmedCoinsTransactions(context.Background(), blockchain, network).Context(context).NewConfirmedCoinsTransactionsRequestBody(newConfirmedCoinsTransactionsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateSubscriptionsForApi.NewConfirmedCoinsTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewConfirmedCoinsTransactions`: NewConfirmedCoinsTransactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `CreateSubscriptionsForApi.NewConfirmedCoinsTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewConfirmedCoinsTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **newConfirmedCoinsTransactionsRequestBody** | [**NewConfirmedCoinsTransactionsRequestBody**](NewConfirmedCoinsTransactionsRequestBody.md) |  | 

### Return type

[**NewConfirmedCoinsTransactionsResponse**](NewConfirmedCoinsTransactionsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewConfirmedCoinsTransactionsAndEachConfirmation

> NewConfirmedCoinsTransactionsAndEachConfirmationResponse NewConfirmedCoinsTransactionsAndEachConfirmation(ctx, blockchain, network).Context(context).NewConfirmedCoinsTransactionsAndEachConfirmationRequestBody(newConfirmedCoinsTransactionsAndEachConfirmationRequestBody).Execute()

New confirmed coins transactions and each confirmation



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
    newConfirmedCoinsTransactionsAndEachConfirmationRequestBody := *openapiclient.NewNewConfirmedCoinsTransactionsAndEachConfirmationRequestBody(*openapiclient.NewNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData(*openapiclient.NewNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem("mho4jHBcrNCncKt38trJahXakuaBnS7LK5", "http://example.com", int32(3)))) // NewConfirmedCoinsTransactionsAndEachConfirmationRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CreateSubscriptionsForApi.NewConfirmedCoinsTransactionsAndEachConfirmation(context.Background(), blockchain, network).Context(context).NewConfirmedCoinsTransactionsAndEachConfirmationRequestBody(newConfirmedCoinsTransactionsAndEachConfirmationRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateSubscriptionsForApi.NewConfirmedCoinsTransactionsAndEachConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewConfirmedCoinsTransactionsAndEachConfirmation`: NewConfirmedCoinsTransactionsAndEachConfirmationResponse
    fmt.Fprintf(os.Stdout, "Response from `CreateSubscriptionsForApi.NewConfirmedCoinsTransactionsAndEachConfirmation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewConfirmedCoinsTransactionsAndEachConfirmationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **newConfirmedCoinsTransactionsAndEachConfirmationRequestBody** | [**NewConfirmedCoinsTransactionsAndEachConfirmationRequestBody**](NewConfirmedCoinsTransactionsAndEachConfirmationRequestBody.md) |  | 

### Return type

[**NewConfirmedCoinsTransactionsAndEachConfirmationResponse**](NewConfirmedCoinsTransactionsAndEachConfirmationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewConfirmedTokensTransactions

> NewConfirmedTokensTransactionsResponse NewConfirmedTokensTransactions(ctx, blockchain, network).Context(context).NewConfirmedTokensTransactionsRequestBody(newConfirmedTokensTransactionsRequestBody).Execute()

New confirmed tokens transactions



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
    newConfirmedTokensTransactionsRequestBody := *openapiclient.NewNewConfirmedTokensTransactionsRequestBody(*openapiclient.NewNewUnconfirmedTokensTransactionsRequestBodyData(*openapiclient.NewNewUnconfirmedTokensTransactionsRequestBodyDataItem("mho4jHBcrNCncKt38trJahXakuaBnS7LK5", "http://example.com"))) // NewConfirmedTokensTransactionsRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CreateSubscriptionsForApi.NewConfirmedTokensTransactions(context.Background(), blockchain, network).Context(context).NewConfirmedTokensTransactionsRequestBody(newConfirmedTokensTransactionsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateSubscriptionsForApi.NewConfirmedTokensTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewConfirmedTokensTransactions`: NewConfirmedTokensTransactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `CreateSubscriptionsForApi.NewConfirmedTokensTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewConfirmedTokensTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **newConfirmedTokensTransactionsRequestBody** | [**NewConfirmedTokensTransactionsRequestBody**](NewConfirmedTokensTransactionsRequestBody.md) |  | 

### Return type

[**NewConfirmedTokensTransactionsResponse**](NewConfirmedTokensTransactionsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewConfirmedTokensTransactionsAndEachConfirmation

> NewConfirmedTokensTransactionsAndEachConfirmationResponse NewConfirmedTokensTransactionsAndEachConfirmation(ctx, blockchain, network).Context(context).NewConfirmedTokensTransactionsAndEachConfirmationRequestBody(newConfirmedTokensTransactionsAndEachConfirmationRequestBody).Execute()

New confirmed tokens transactions and each confirmation



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
    newConfirmedTokensTransactionsAndEachConfirmationRequestBody := *openapiclient.NewNewConfirmedTokensTransactionsAndEachConfirmationRequestBody(*openapiclient.NewNewConfirmedTokensTransactionsAndEachConfirmationRequestBodyData(*openapiclient.NewNewConfirmedTokensTransactionsAndEachConfirmationRequestBodyDataItem("mho4jHBcrNCncKt38trJahXakuaBnS7LK5", "http://example.com"))) // NewConfirmedTokensTransactionsAndEachConfirmationRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CreateSubscriptionsForApi.NewConfirmedTokensTransactionsAndEachConfirmation(context.Background(), blockchain, network).Context(context).NewConfirmedTokensTransactionsAndEachConfirmationRequestBody(newConfirmedTokensTransactionsAndEachConfirmationRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateSubscriptionsForApi.NewConfirmedTokensTransactionsAndEachConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewConfirmedTokensTransactionsAndEachConfirmation`: NewConfirmedTokensTransactionsAndEachConfirmationResponse
    fmt.Fprintf(os.Stdout, "Response from `CreateSubscriptionsForApi.NewConfirmedTokensTransactionsAndEachConfirmation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewConfirmedTokensTransactionsAndEachConfirmationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **newConfirmedTokensTransactionsAndEachConfirmationRequestBody** | [**NewConfirmedTokensTransactionsAndEachConfirmationRequestBody**](NewConfirmedTokensTransactionsAndEachConfirmationRequestBody.md) |  | 

### Return type

[**NewConfirmedTokensTransactionsAndEachConfirmationResponse**](NewConfirmedTokensTransactionsAndEachConfirmationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewUnconfirmedCoinsTransactions

> NewUnconfirmedCoinsTransactionsResponse NewUnconfirmedCoinsTransactions(ctx, blockchain, network).Context(context).NewUnconfirmedCoinsTransactionsRequestBody(newUnconfirmedCoinsTransactionsRequestBody).Execute()

New unconfirmed coins transactions



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
    newUnconfirmedCoinsTransactionsRequestBody := *openapiclient.NewNewUnconfirmedCoinsTransactionsRequestBody(*openapiclient.NewNewUnconfirmedCoinsTransactionsRequestBodyData(*openapiclient.NewNewUnconfirmedCoinsTransactionsRequestBodyDataItem("mho4jHBcrNCncKt38trJahXakuaBnS7LK5", "http://example.com"))) // NewUnconfirmedCoinsTransactionsRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CreateSubscriptionsForApi.NewUnconfirmedCoinsTransactions(context.Background(), blockchain, network).Context(context).NewUnconfirmedCoinsTransactionsRequestBody(newUnconfirmedCoinsTransactionsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateSubscriptionsForApi.NewUnconfirmedCoinsTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewUnconfirmedCoinsTransactions`: NewUnconfirmedCoinsTransactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `CreateSubscriptionsForApi.NewUnconfirmedCoinsTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewUnconfirmedCoinsTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **newUnconfirmedCoinsTransactionsRequestBody** | [**NewUnconfirmedCoinsTransactionsRequestBody**](NewUnconfirmedCoinsTransactionsRequestBody.md) |  | 

### Return type

[**NewUnconfirmedCoinsTransactionsResponse**](NewUnconfirmedCoinsTransactionsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewUnconfirmedTokensTransactions

> NewUnconfirmedTokensTransactionsResponse NewUnconfirmedTokensTransactions(ctx, blockchain, network).Context(context).NewUnconfirmedTokensTransactionsRequestBody(newUnconfirmedTokensTransactionsRequestBody).Execute()

New unconfirmed tokens transactions



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
    newUnconfirmedTokensTransactionsRequestBody := *openapiclient.NewNewUnconfirmedTokensTransactionsRequestBody(*openapiclient.NewNewUnconfirmedTokensTransactionsRequestBodyData(*openapiclient.NewNewUnconfirmedTokensTransactionsRequestBodyDataItem("mho4jHBcrNCncKt38trJahXakuaBnS7LK5", "http://example.com"))) // NewUnconfirmedTokensTransactionsRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CreateSubscriptionsForApi.NewUnconfirmedTokensTransactions(context.Background(), blockchain, network).Context(context).NewUnconfirmedTokensTransactionsRequestBody(newUnconfirmedTokensTransactionsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateSubscriptionsForApi.NewUnconfirmedTokensTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewUnconfirmedTokensTransactions`: NewUnconfirmedTokensTransactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `CreateSubscriptionsForApi.NewUnconfirmedTokensTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewUnconfirmedTokensTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **newUnconfirmedTokensTransactionsRequestBody** | [**NewUnconfirmedTokensTransactionsRequestBody**](NewUnconfirmedTokensTransactionsRequestBody.md) |  | 

### Return type

[**NewUnconfirmedTokensTransactionsResponse**](NewUnconfirmedTokensTransactionsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \CreateSubscriptionsForApi

All URIs are relative to *https://rest.cryptoapis.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MinedTransaction**](CreateSubscriptionsForApi.md#MinedTransaction) | **Post** /blockchain-events/{blockchain}/{network}/subscriptions/transaction-mined | Mined transaction
[**NewBlock**](CreateSubscriptionsForApi.md#NewBlock) | **Post** /blockchain-events/{blockchain}/{network}/subscriptions/block-mined | New Block
[**NewConfirmedCoinsTransactions**](CreateSubscriptionsForApi.md#NewConfirmedCoinsTransactions) | **Post** /blockchain-events/{blockchain}/{network}/subscriptions/address-coins-transactions-confirmed | New confirmed coins transactions
[**NewConfirmedCoinsTransactionsAndEachConfirmation**](CreateSubscriptionsForApi.md#NewConfirmedCoinsTransactionsAndEachConfirmation) | **Post** /blockchain-events/{blockchain}/{network}/subscriptions/address-coins-transactions-confirmed-each-confirmation | New confirmed coins transactions and each confirmation
[**NewConfirmedInternalTransactions**](CreateSubscriptionsForApi.md#NewConfirmedInternalTransactions) | **Post** /blockchain-events/{blockchain}/{network}/subscriptions/address-internal-transactions-confirmed | New confirmed internal transactions
[**NewConfirmedInternalTransactionsAndEachConfirmation**](CreateSubscriptionsForApi.md#NewConfirmedInternalTransactionsAndEachConfirmation) | **Post** /blockchain-events/{blockchain}/{network}/subscriptions/address-internal-transactions-confirmed-each-confirmation | New confirmed internal transactions and each confirmation
[**NewConfirmedTokensTransactions**](CreateSubscriptionsForApi.md#NewConfirmedTokensTransactions) | **Post** /blockchain-events/{blockchain}/{network}/subscriptions/address-tokens-transactions-confirmed | New confirmed tokens transactions
[**NewConfirmedTokensTransactionsAndEachConfirmation**](CreateSubscriptionsForApi.md#NewConfirmedTokensTransactionsAndEachConfirmation) | **Post** /blockchain-events/{blockchain}/{network}/subscriptions/address-tokens-transactions-confirmed-each-confirmation | New confirmed tokens transactions and each confirmation
[**NewUnconfirmedCoinsTransactions**](CreateSubscriptionsForApi.md#NewUnconfirmedCoinsTransactions) | **Post** /blockchain-events/{blockchain}/{network}/subscriptions/address-coins-transactions-unconfirmed | New unconfirmed coins transactions
[**NewUnconfirmedTokensTransactions**](CreateSubscriptionsForApi.md#NewUnconfirmedTokensTransactions) | **Post** /blockchain-events/{blockchain}/{network}/subscriptions/address-tokens-transactions-unconfirmed | New unconfirmed tokens transactions



## MinedTransaction

> MinedTransactionR MinedTransaction(ctx, blockchain, network).Context(context).MinedTransactionRB(minedTransactionRB).Execute()

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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    minedTransactionRB := *openapiclient.NewMinedTransactionRB(*openapiclient.NewMinedTransactionRBData(*openapiclient.NewMinedTransactionRBDataItem("http://example.com", "df2690ff97e72c1f8b0f2102a8cb5c1d0fa8fb8754d543c9bc0edc4d4bc34bfc"))) // MinedTransactionRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CreateSubscriptionsForApi.MinedTransaction(context.Background(), blockchain, network).Context(context).MinedTransactionRB(minedTransactionRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateSubscriptionsForApi.MinedTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MinedTransaction`: MinedTransactionR
    fmt.Fprintf(os.Stdout, "Response from `CreateSubscriptionsForApi.MinedTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMinedTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **minedTransactionRB** | [**MinedTransactionRB**](MinedTransactionRB.md) |  | 

### Return type

[**MinedTransactionR**](MinedTransactionR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewBlock

> NewBlockR NewBlock(ctx, blockchain, network).Context(context).NewBlockRB(newBlockRB).Execute()

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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    newBlockRB := *openapiclient.NewNewBlockRB(*openapiclient.NewNewBlockRBData(*openapiclient.NewNewBlockRBDataItem("http://example.com"))) // NewBlockRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CreateSubscriptionsForApi.NewBlock(context.Background(), blockchain, network).Context(context).NewBlockRB(newBlockRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateSubscriptionsForApi.NewBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewBlock`: NewBlockR
    fmt.Fprintf(os.Stdout, "Response from `CreateSubscriptionsForApi.NewBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **newBlockRB** | [**NewBlockRB**](NewBlockRB.md) |  | 

### Return type

[**NewBlockR**](NewBlockR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewConfirmedCoinsTransactions

> NewConfirmedCoinsTransactionsR NewConfirmedCoinsTransactions(ctx, blockchain, network).Context(context).NewConfirmedCoinsTransactionsRB(newConfirmedCoinsTransactionsRB).Execute()

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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    newConfirmedCoinsTransactionsRB := *openapiclient.NewNewConfirmedCoinsTransactionsRB(*openapiclient.NewNewConfirmedCoinsTransactionsRBData(*openapiclient.NewNewConfirmedCoinsTransactionsRBDataItem("mho4jHBcrNCncKt38trJahXakuaBnS7LK5", "http://example.com"))) // NewConfirmedCoinsTransactionsRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CreateSubscriptionsForApi.NewConfirmedCoinsTransactions(context.Background(), blockchain, network).Context(context).NewConfirmedCoinsTransactionsRB(newConfirmedCoinsTransactionsRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateSubscriptionsForApi.NewConfirmedCoinsTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewConfirmedCoinsTransactions`: NewConfirmedCoinsTransactionsR
    fmt.Fprintf(os.Stdout, "Response from `CreateSubscriptionsForApi.NewConfirmedCoinsTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewConfirmedCoinsTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **newConfirmedCoinsTransactionsRB** | [**NewConfirmedCoinsTransactionsRB**](NewConfirmedCoinsTransactionsRB.md) |  | 

### Return type

[**NewConfirmedCoinsTransactionsR**](NewConfirmedCoinsTransactionsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewConfirmedCoinsTransactionsAndEachConfirmation

> NewConfirmedCoinsTransactionsAndEachConfirmationR NewConfirmedCoinsTransactionsAndEachConfirmation(ctx, blockchain, network).Context(context).NewConfirmedCoinsTransactionsAndEachConfirmationRB(newConfirmedCoinsTransactionsAndEachConfirmationRB).Execute()

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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    newConfirmedCoinsTransactionsAndEachConfirmationRB := *openapiclient.NewNewConfirmedCoinsTransactionsAndEachConfirmationRB(*openapiclient.NewNewConfirmedCoinsTransactionsAndEachConfirmationRBData(*openapiclient.NewNewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem("mho4jHBcrNCncKt38trJahXakuaBnS7LK5", "http://example.com", int32(3)))) // NewConfirmedCoinsTransactionsAndEachConfirmationRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CreateSubscriptionsForApi.NewConfirmedCoinsTransactionsAndEachConfirmation(context.Background(), blockchain, network).Context(context).NewConfirmedCoinsTransactionsAndEachConfirmationRB(newConfirmedCoinsTransactionsAndEachConfirmationRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateSubscriptionsForApi.NewConfirmedCoinsTransactionsAndEachConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewConfirmedCoinsTransactionsAndEachConfirmation`: NewConfirmedCoinsTransactionsAndEachConfirmationR
    fmt.Fprintf(os.Stdout, "Response from `CreateSubscriptionsForApi.NewConfirmedCoinsTransactionsAndEachConfirmation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewConfirmedCoinsTransactionsAndEachConfirmationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **newConfirmedCoinsTransactionsAndEachConfirmationRB** | [**NewConfirmedCoinsTransactionsAndEachConfirmationRB**](NewConfirmedCoinsTransactionsAndEachConfirmationRB.md) |  | 

### Return type

[**NewConfirmedCoinsTransactionsAndEachConfirmationR**](NewConfirmedCoinsTransactionsAndEachConfirmationR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewConfirmedInternalTransactions

> NewConfirmedInternalTransactionsR NewConfirmedInternalTransactions(ctx, blockchain, network).Context(context).NewConfirmedInternalTransactionsRB(newConfirmedInternalTransactionsRB).Execute()

New confirmed internal transactions



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
    blockchain := "ethereum-classic" // string | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
    network := "mordor" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    newConfirmedInternalTransactionsRB := *openapiclient.NewNewConfirmedInternalTransactionsRB(*openapiclient.NewNewConfirmedInternalTransactionsRBData(*openapiclient.NewNewConfirmedInternalTransactionsRBDataItem("0xbcc817f057950b0df41206c5d7125e6225cae18e", true, "yourSecretKey", "http://example.com"))) // NewConfirmedInternalTransactionsRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CreateSubscriptionsForApi.NewConfirmedInternalTransactions(context.Background(), blockchain, network).Context(context).NewConfirmedInternalTransactionsRB(newConfirmedInternalTransactionsRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateSubscriptionsForApi.NewConfirmedInternalTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewConfirmedInternalTransactions`: NewConfirmedInternalTransactionsR
    fmt.Fprintf(os.Stdout, "Response from `CreateSubscriptionsForApi.NewConfirmedInternalTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewConfirmedInternalTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **newConfirmedInternalTransactionsRB** | [**NewConfirmedInternalTransactionsRB**](NewConfirmedInternalTransactionsRB.md) |  | 

### Return type

[**NewConfirmedInternalTransactionsR**](NewConfirmedInternalTransactionsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewConfirmedInternalTransactionsAndEachConfirmation

> NewConfirmedInternalTransactionsAndEachConfirmationR NewConfirmedInternalTransactionsAndEachConfirmation(ctx, blockchain, network).Context(context).NewConfirmedInternalTransactionsAndEachConfirmationRB(newConfirmedInternalTransactionsAndEachConfirmationRB).Execute()

New confirmed internal transactions and each confirmation



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
    blockchain := "ethereum-classic" // string | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
    network := "mordor" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    newConfirmedInternalTransactionsAndEachConfirmationRB := *openapiclient.NewNewConfirmedInternalTransactionsAndEachConfirmationRB(*openapiclient.NewNewConfirmedInternalTransactionsAndEachConfirmationRBData(*openapiclient.NewNewConfirmedInternalTransactionsAndEachConfirmationRBDataItem("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", true, "yourSecretString", "https://example.com", int32(3)))) // NewConfirmedInternalTransactionsAndEachConfirmationRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CreateSubscriptionsForApi.NewConfirmedInternalTransactionsAndEachConfirmation(context.Background(), blockchain, network).Context(context).NewConfirmedInternalTransactionsAndEachConfirmationRB(newConfirmedInternalTransactionsAndEachConfirmationRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateSubscriptionsForApi.NewConfirmedInternalTransactionsAndEachConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewConfirmedInternalTransactionsAndEachConfirmation`: NewConfirmedInternalTransactionsAndEachConfirmationR
    fmt.Fprintf(os.Stdout, "Response from `CreateSubscriptionsForApi.NewConfirmedInternalTransactionsAndEachConfirmation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewConfirmedInternalTransactionsAndEachConfirmationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **newConfirmedInternalTransactionsAndEachConfirmationRB** | [**NewConfirmedInternalTransactionsAndEachConfirmationRB**](NewConfirmedInternalTransactionsAndEachConfirmationRB.md) |  | 

### Return type

[**NewConfirmedInternalTransactionsAndEachConfirmationR**](NewConfirmedInternalTransactionsAndEachConfirmationR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewConfirmedTokensTransactions

> NewConfirmedTokensTransactionsR NewConfirmedTokensTransactions(ctx, blockchain, network).Context(context).NewConfirmedTokensTransactionsRB(newConfirmedTokensTransactionsRB).Execute()

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
    blockchain := "ethereum" // string | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
    network := "ropsten" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    newConfirmedTokensTransactionsRB := *openapiclient.NewNewConfirmedTokensTransactionsRB(*openapiclient.NewNewConfirmedTokensTransactionsRBData(*openapiclient.NewNewConfirmedTokensTransactionsRBDataItem("0xbf16582e53d6fd892f11de8a3e29e8c3b65d77c2", "http://example.com"))) // NewConfirmedTokensTransactionsRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CreateSubscriptionsForApi.NewConfirmedTokensTransactions(context.Background(), blockchain, network).Context(context).NewConfirmedTokensTransactionsRB(newConfirmedTokensTransactionsRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateSubscriptionsForApi.NewConfirmedTokensTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewConfirmedTokensTransactions`: NewConfirmedTokensTransactionsR
    fmt.Fprintf(os.Stdout, "Response from `CreateSubscriptionsForApi.NewConfirmedTokensTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewConfirmedTokensTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **newConfirmedTokensTransactionsRB** | [**NewConfirmedTokensTransactionsRB**](NewConfirmedTokensTransactionsRB.md) |  | 

### Return type

[**NewConfirmedTokensTransactionsR**](NewConfirmedTokensTransactionsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewConfirmedTokensTransactionsAndEachConfirmation

> NewConfirmedTokensTransactionsAndEachConfirmationR NewConfirmedTokensTransactionsAndEachConfirmation(ctx, blockchain, network).Context(context).NewConfirmedTokensTransactionsAndEachConfirmationRB(newConfirmedTokensTransactionsAndEachConfirmationRB).Execute()

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
    blockchain := "ethereum" // string | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
    network := "ropsten" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    newConfirmedTokensTransactionsAndEachConfirmationRB := *openapiclient.NewNewConfirmedTokensTransactionsAndEachConfirmationRB(*openapiclient.NewNewConfirmedTokensTransactionsAndEachConfirmationRBData(*openapiclient.NewNewConfirmedTokensTransactionsAndEachConfirmationRBDataItem("0x033ef6db9fbd0ee60e2931906b987fe0280471a0", "http://example.com"))) // NewConfirmedTokensTransactionsAndEachConfirmationRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CreateSubscriptionsForApi.NewConfirmedTokensTransactionsAndEachConfirmation(context.Background(), blockchain, network).Context(context).NewConfirmedTokensTransactionsAndEachConfirmationRB(newConfirmedTokensTransactionsAndEachConfirmationRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateSubscriptionsForApi.NewConfirmedTokensTransactionsAndEachConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewConfirmedTokensTransactionsAndEachConfirmation`: NewConfirmedTokensTransactionsAndEachConfirmationR
    fmt.Fprintf(os.Stdout, "Response from `CreateSubscriptionsForApi.NewConfirmedTokensTransactionsAndEachConfirmation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewConfirmedTokensTransactionsAndEachConfirmationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **newConfirmedTokensTransactionsAndEachConfirmationRB** | [**NewConfirmedTokensTransactionsAndEachConfirmationRB**](NewConfirmedTokensTransactionsAndEachConfirmationRB.md) |  | 

### Return type

[**NewConfirmedTokensTransactionsAndEachConfirmationR**](NewConfirmedTokensTransactionsAndEachConfirmationR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewUnconfirmedCoinsTransactions

> NewUnconfirmedCoinsTransactionsR NewUnconfirmedCoinsTransactions(ctx, blockchain, network).Context(context).NewUnconfirmedCoinsTransactionsRB(newUnconfirmedCoinsTransactionsRB).Execute()

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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    newUnconfirmedCoinsTransactionsRB := *openapiclient.NewNewUnconfirmedCoinsTransactionsRB(*openapiclient.NewNewUnconfirmedCoinsTransactionsRBData(*openapiclient.NewNewUnconfirmedCoinsTransactionsRBDataItem("mho4jHBcrNCncKt38trJahXakuaBnS7LK5", "http://example.com"))) // NewUnconfirmedCoinsTransactionsRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CreateSubscriptionsForApi.NewUnconfirmedCoinsTransactions(context.Background(), blockchain, network).Context(context).NewUnconfirmedCoinsTransactionsRB(newUnconfirmedCoinsTransactionsRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateSubscriptionsForApi.NewUnconfirmedCoinsTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewUnconfirmedCoinsTransactions`: NewUnconfirmedCoinsTransactionsR
    fmt.Fprintf(os.Stdout, "Response from `CreateSubscriptionsForApi.NewUnconfirmedCoinsTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewUnconfirmedCoinsTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **newUnconfirmedCoinsTransactionsRB** | [**NewUnconfirmedCoinsTransactionsRB**](NewUnconfirmedCoinsTransactionsRB.md) |  | 

### Return type

[**NewUnconfirmedCoinsTransactionsR**](NewUnconfirmedCoinsTransactionsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewUnconfirmedTokensTransactions

> NewUnconfirmedTokensTransactionsR NewUnconfirmedTokensTransactions(ctx, blockchain, network).Context(context).NewUnconfirmedTokensTransactionsRB(newUnconfirmedTokensTransactionsRB).Execute()

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
    blockchain := "ethereum" // string | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
    network := "ropsten" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    newUnconfirmedTokensTransactionsRB := *openapiclient.NewNewUnconfirmedTokensTransactionsRB(*openapiclient.NewNewUnconfirmedTokensTransactionsRBData(*openapiclient.NewNewUnconfirmedTokensTransactionsRBDataItem("0x033ef6db9fbd0ee60e2931906b987fe0280471a0", "http://example.com"))) // NewUnconfirmedTokensTransactionsRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CreateSubscriptionsForApi.NewUnconfirmedTokensTransactions(context.Background(), blockchain, network).Context(context).NewUnconfirmedTokensTransactionsRB(newUnconfirmedTokensTransactionsRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateSubscriptionsForApi.NewUnconfirmedTokensTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewUnconfirmedTokensTransactions`: NewUnconfirmedTokensTransactionsR
    fmt.Fprintf(os.Stdout, "Response from `CreateSubscriptionsForApi.NewUnconfirmedTokensTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewUnconfirmedTokensTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **newUnconfirmedTokensTransactionsRB** | [**NewUnconfirmedTokensTransactionsRB**](NewUnconfirmedTokensTransactionsRB.md) |  | 

### Return type

[**NewUnconfirmedTokensTransactionsR**](NewUnconfirmedTokensTransactionsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


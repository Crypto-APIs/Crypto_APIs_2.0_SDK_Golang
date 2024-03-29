# \CallbackDataApi

All URIs are relative to *https://rest.cryptoapis.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddressDetailsFromCallback**](CallbackDataApi.md#GetAddressDetailsFromCallback) | **Get** /blockchain-events/{blockchain}/{network}/addresses/{address} | Get Address Details From Callback
[**GetBlockDetailsByBlockHashFromCallback**](CallbackDataApi.md#GetBlockDetailsByBlockHashFromCallback) | **Get** /blockchain-events/{blockchain}/{network}/blocks/hash/{blockHash} | Get Block Details By Block Hash From Callback
[**GetBlockDetailsByBlockHeightFromCallback**](CallbackDataApi.md#GetBlockDetailsByBlockHeightFromCallback) | **Get** /blockchain-events/{blockchain}/{network}/blocks/height/{blockHeight} | Get Block Details By Block Height From Callback
[**GetTransactionDetailsByTransactionIDFromCallback**](CallbackDataApi.md#GetTransactionDetailsByTransactionIDFromCallback) | **Get** /blockchain-events/{blockchain}/{network}/transactions/{transactionId} | Get Transaction Details By Transaction ID From Callback



## GetAddressDetailsFromCallback

> GetAddressDetailsFromCallbackR GetAddressDetailsFromCallback(ctx, blockchain, network, address).Context(context).Execute()

Get Address Details From Callback



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
    address := "mzYijhgmzZrmuB7wBDazRKirnChKyow4M3" // string | Represents the public address, which is a compressed and shortened form of a public key.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallbackDataApi.GetAddressDetailsFromCallback(context.Background(), blockchain, network, address).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbackDataApi.GetAddressDetailsFromCallback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddressDetailsFromCallback`: GetAddressDetailsFromCallbackR
    fmt.Fprintf(os.Stdout, "Response from `CallbackDataApi.GetAddressDetailsFromCallback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**address** | **string** | Represents the public address, which is a compressed and shortened form of a public key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressDetailsFromCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetAddressDetailsFromCallbackR**](GetAddressDetailsFromCallbackR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockDetailsByBlockHashFromCallback

> GetBlockDetailsByBlockHashFromCallbackR GetBlockDetailsByBlockHashFromCallback(ctx, blockchain, network, blockHash).Context(context).Execute()

Get Block Details By Block Hash From Callback



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
    blockHash := "000000000000000bdea8ba7df4bfd9f398e428fde8ee47152bcf93834ee48e8a" // string | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallbackDataApi.GetBlockDetailsByBlockHashFromCallback(context.Background(), blockchain, network, blockHash).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbackDataApi.GetBlockDetailsByBlockHashFromCallback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockDetailsByBlockHashFromCallback`: GetBlockDetailsByBlockHashFromCallbackR
    fmt.Fprintf(os.Stdout, "Response from `CallbackDataApi.GetBlockDetailsByBlockHashFromCallback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**blockHash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockDetailsByBlockHashFromCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetBlockDetailsByBlockHashFromCallbackR**](GetBlockDetailsByBlockHashFromCallbackR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockDetailsByBlockHeightFromCallback

> GetBlockDetailsByBlockHeightFromCallbackR GetBlockDetailsByBlockHeightFromCallback(ctx, blockchain, network, blockHeight).Context(context).Execute()

Get Block Details By Block Height From Callback



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
    blockHeight := "673852" // string | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \"Genesis block\".
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallbackDataApi.GetBlockDetailsByBlockHeightFromCallback(context.Background(), blockchain, network, blockHeight).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbackDataApi.GetBlockDetailsByBlockHeightFromCallback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockDetailsByBlockHeightFromCallback`: GetBlockDetailsByBlockHeightFromCallbackR
    fmt.Fprintf(os.Stdout, "Response from `CallbackDataApi.GetBlockDetailsByBlockHeightFromCallback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**blockHeight** | **string** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockDetailsByBlockHeightFromCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetBlockDetailsByBlockHeightFromCallbackR**](GetBlockDetailsByBlockHeightFromCallbackR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionDetailsByTransactionIDFromCallback

> GetTransactionDetailsByTransactionIDFromCallbackR GetTransactionDetailsByTransactionIDFromCallback(ctx, blockchain, network, transactionId).Context(context).Execute()

Get Transaction Details By Transaction ID From Callback



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
    transactionId := "8888f6c8168ff69aaf6438ab185c690e8c76c63e5f9c472c1c86f08406ea74f2" // string | Represents the unique identifier of a transaction, i.e. it could be transactionId in UTXO-based protocols like Bitcoin, and transaction hash in Ethereum blockchain.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallbackDataApi.GetTransactionDetailsByTransactionIDFromCallback(context.Background(), blockchain, network, transactionId).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallbackDataApi.GetTransactionDetailsByTransactionIDFromCallback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionDetailsByTransactionIDFromCallback`: GetTransactionDetailsByTransactionIDFromCallbackR
    fmt.Fprintf(os.Stdout, "Response from `CallbackDataApi.GetTransactionDetailsByTransactionIDFromCallback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**transactionId** | **string** | Represents the unique identifier of a transaction, i.e. it could be transactionId in UTXO-based protocols like Bitcoin, and transaction hash in Ethereum blockchain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionDetailsByTransactionIDFromCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetTransactionDetailsByTransactionIDFromCallbackR**](GetTransactionDetailsByTransactionIDFromCallbackR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


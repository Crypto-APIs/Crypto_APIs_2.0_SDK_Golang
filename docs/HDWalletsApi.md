# \HDWalletsApi

All URIs are relative to *https://rest.cryptoapis.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHDWalletXPubYPubZPubDetails**](HDWalletsApi.md#GetHDWalletXPubYPubZPubDetails) | **Get** /blockchain-data/{blockchain}/{network}/hd/{extendedPublicKey}/details | Get HD Wallet (xPub, yPub, zPub) Details
[**ListHDWalletXPubYPubZPubTransactions**](HDWalletsApi.md#ListHDWalletXPubYPubZPubTransactions) | **Get** /blockchain-data/{blockchain}/{network}/hd/{extendedPublicKey}/transactions | List HD Wallet (xPub, yPub, zPub) Transactions
[**SyncHDWalletXPubYPubZPub**](HDWalletsApi.md#SyncHDWalletXPubYPubZPub) | **Post** /blockchain-data/{blockchain}/{network}/hd/sync | Sync HD Wallet (xPub, yPub, zPub)



## GetHDWalletXPubYPubZPubDetails

> GetHDWalletXPubYPubZPubDetailsR GetHDWalletXPubYPubZPubDetails(ctx, blockchain, extendedPublicKey, network).Context(context).Derivation(derivation).Execute()

Get HD Wallet (xPub, yPub, zPub) Details



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
    extendedPublicKey := "upub5Ei6bRNneqozk6smK7dvtXHC5PjUyEL4ynCfMKvjznLcXi9DQaikETzQjHvJC43XexMvQs64jxB1njMjCHpRZ4xQWAmv3ge9cVtjfsHmbvQ" // string | Defines the account extended publicly known key which is used to derive all child public keys.
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    derivation := "derivation_example" // string | The way how the HD walled derives, for example when the type is ACCOUNT, it derives change and receive addresses while when the type is BIP32 it derives directly. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HDWalletsApi.GetHDWalletXPubYPubZPubDetails(context.Background(), blockchain, extendedPublicKey, network).Context(context).Derivation(derivation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HDWalletsApi.GetHDWalletXPubYPubZPubDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHDWalletXPubYPubZPubDetails`: GetHDWalletXPubYPubZPubDetailsR
    fmt.Fprintf(os.Stdout, "Response from `HDWalletsApi.GetHDWalletXPubYPubZPubDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**extendedPublicKey** | **string** | Defines the account extended publicly known key which is used to derive all child public keys. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHDWalletXPubYPubZPubDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **derivation** | **string** | The way how the HD walled derives, for example when the type is ACCOUNT, it derives change and receive addresses while when the type is BIP32 it derives directly. | 

### Return type

[**GetHDWalletXPubYPubZPubDetailsR**](GetHDWalletXPubYPubZPubDetailsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHDWalletXPubYPubZPubTransactions

> ListHDWalletXPubYPubZPubTransactionsR ListHDWalletXPubYPubZPubTransactions(ctx, blockchain, extendedPublicKey, network).Context(context).Derivation(derivation).Limit(limit).Offset(offset).Execute()

List HD Wallet (xPub, yPub, zPub) Transactions



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
    blockchain := "bitcoin" // string | Represents the specific blockchain.
    extendedPublicKey := "tpubD9GMECjiZHCaF9NHSMAeMbQMXnM7CviEJZsYBuztVwsUjPHWjxewWAUXWV2UExaAtoEvQGXDBmVWo6ZHGtj6TsH6Pop7D9DskQwGHA1gu1w" // string | Defines the master public key (xPub) of the account.
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    derivation := "derivation_example" // string | The way how the HD walled derives, for example when the type is ACCOUNT, it derives change and receive addresses while when the type is BIP32 it derives directly. (optional)
    limit := int32(50) // int32 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int32(10) // int32 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HDWalletsApi.ListHDWalletXPubYPubZPubTransactions(context.Background(), blockchain, extendedPublicKey, network).Context(context).Derivation(derivation).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HDWalletsApi.ListHDWalletXPubYPubZPubTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHDWalletXPubYPubZPubTransactions`: ListHDWalletXPubYPubZPubTransactionsR
    fmt.Fprintf(os.Stdout, "Response from `HDWalletsApi.ListHDWalletXPubYPubZPubTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain. | 
**extendedPublicKey** | **string** | Defines the master public key (xPub) of the account. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHDWalletXPubYPubZPubTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **derivation** | **string** | The way how the HD walled derives, for example when the type is ACCOUNT, it derives change and receive addresses while when the type is BIP32 it derives directly. | 
 **limit** | **int32** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListHDWalletXPubYPubZPubTransactionsR**](ListHDWalletXPubYPubZPubTransactionsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncHDWalletXPubYPubZPub

> SyncHDWalletXPubYPubZPubR SyncHDWalletXPubYPubZPub(ctx, blockchain, network).Context(context).SyncHDWalletXPubYPubZPubRB(syncHDWalletXPubYPubZPubRB).Execute()

Sync HD Wallet (xPub, yPub, zPub)



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
    syncHDWalletXPubYPubZPubRB := *openapiclient.NewSyncHDWalletXPubYPubZPubRB(*openapiclient.NewSyncHDWalletXPubYPubZPubRBData(*openapiclient.NewSyncHDWalletXPubYPubZPubRBDataItem("upub5Ei6bRNneqozk6smK7dvtXHC5PjUyEL4ynCfMKvjznLcXi9DQaikETzQjHvJC43XexMvQs64jxB1njMjCHpRZ4xQWAmv3ge9cVtjfsHmbvQ"))) // SyncHDWalletXPubYPubZPubRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HDWalletsApi.SyncHDWalletXPubYPubZPub(context.Background(), blockchain, network).Context(context).SyncHDWalletXPubYPubZPubRB(syncHDWalletXPubYPubZPubRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HDWalletsApi.SyncHDWalletXPubYPubZPub``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncHDWalletXPubYPubZPub`: SyncHDWalletXPubYPubZPubR
    fmt.Fprintf(os.Stdout, "Response from `HDWalletsApi.SyncHDWalletXPubYPubZPub`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncHDWalletXPubYPubZPubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **syncHDWalletXPubYPubZPubRB** | [**SyncHDWalletXPubYPubZPubRB**](SyncHDWalletXPubYPubZPubRB.md) |  | 

### Return type

[**SyncHDWalletXPubYPubZPubR**](SyncHDWalletXPubYPubZPubR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


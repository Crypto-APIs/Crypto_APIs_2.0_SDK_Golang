# \HDWalletsApi

All URIs are relative to *https://rest.cryptoapis.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeriveAndSyncNewChangeAddresses**](HDWalletsApi.md#DeriveAndSyncNewChangeAddresses) | **Post** /blockchain-data/{blockchain}/{network}/hd/derive-sync-change | Derive And Sync New Change Addresses
[**DeriveAndSyncNewReceivingAddresses**](HDWalletsApi.md#DeriveAndSyncNewReceivingAddresses) | **Post** /blockchain-data/{blockchain}/{network}/hd/derive-and-sync | Derive And Sync New Receiving Addresses
[**GetHDWalletXPubYPubZPubAssets**](HDWalletsApi.md#GetHDWalletXPubYPubZPubAssets) | **Get** /blockchain-data/{blockchain}/{network}/hd/{extendedPublicKey}/assets | Get HD Wallet (xPub, yPub, zPub) Assets
[**GetHDWalletXPubYPubZPubDetails**](HDWalletsApi.md#GetHDWalletXPubYPubZPubDetails) | **Get** /blockchain-data/{blockchain}/{network}/hd/{extendedPublicKey}/details | Get HD Wallet (xPub, yPub, zPub) Details
[**ListHDWalletXPubYPubZPubTransactions**](HDWalletsApi.md#ListHDWalletXPubYPubZPubTransactions) | **Get** /blockchain-data/{blockchain}/{network}/hd/{extendedPublicKey}/transactions | List HD Wallet (xPub, yPub, zPub) Transactions
[**ListHDWalletXPubYPubZPubUTXOs**](HDWalletsApi.md#ListHDWalletXPubYPubZPubUTXOs) | **Get** /blockchain-data/{blockchain}/{network}/hd/{extendedPublicKey}/utxos | List HD Wallet (xPub, yPub, zPub) UTXOs
[**ListSyncedAddresses**](HDWalletsApi.md#ListSyncedAddresses) | **Get** /blockchain-data/{blockchain}/{network}/hd/{extendedPublicKey}/synced-addresses | List Synced Addresses
[**PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPub**](HDWalletsApi.md#PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPub) | **Post** /blockchain-data/{blockchain}/{network}/transactions/prepare-utxo-transaction | Prepare A UTXO-Based Transaction From HD Wallet (xPub, yPub, zPub)
[**PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPub**](HDWalletsApi.md#PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPub) | **Post** /blockchain-data/{blockchain}/{network}/transactions/prepare-account-based-transaction | Prepare An Account-Based Transaction From HD Wallet (xPub, yPub, zPub)
[**SyncHDWalletXPubYPubZPub**](HDWalletsApi.md#SyncHDWalletXPubYPubZPub) | **Post** /blockchain-data/{blockchain}/{network}/hd/sync | Sync HD Wallet (xPub, yPub, zPub)
[**SyncNewHDWalletXPubYPubZPub**](HDWalletsApi.md#SyncNewHDWalletXPubYPubZPub) | **Post** /blockchain-data/{blockchain}/{network}/hd/sync-new | Sync New HD Wallet (xPub, yPub, zPub)



## DeriveAndSyncNewChangeAddresses

> DeriveAndSyncNewChangeAddressesR DeriveAndSyncNewChangeAddresses(ctx, blockchain, network).Context(context).DeriveAndSyncNewChangeAddressesRB(deriveAndSyncNewChangeAddressesRB).Execute()

Derive And Sync New Change Addresses



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
    deriveAndSyncNewChangeAddressesRB := *openapiclient.NewDeriveAndSyncNewChangeAddressesRB(*openapiclient.NewDeriveAndSyncNewChangeAddressesRBData(*openapiclient.NewDeriveAndSyncNewChangeAddressesRBDataItem("xpub6BuJ8T4xTEePRTWxEcyyZRHPRZw91GFRjuu4H1eNqNGDswpraD5Hthf7JBbK7iQayuLf2MbxT6MVrKGbY7HvBcQo937Qiwmxz7Fzy9S5y9q"))) // DeriveAndSyncNewChangeAddressesRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HDWalletsApi.DeriveAndSyncNewChangeAddresses(context.Background(), blockchain, network).Context(context).DeriveAndSyncNewChangeAddressesRB(deriveAndSyncNewChangeAddressesRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HDWalletsApi.DeriveAndSyncNewChangeAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeriveAndSyncNewChangeAddresses`: DeriveAndSyncNewChangeAddressesR
    fmt.Fprintf(os.Stdout, "Response from `HDWalletsApi.DeriveAndSyncNewChangeAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeriveAndSyncNewChangeAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **deriveAndSyncNewChangeAddressesRB** | [**DeriveAndSyncNewChangeAddressesRB**](DeriveAndSyncNewChangeAddressesRB.md) |  | 

### Return type

[**DeriveAndSyncNewChangeAddressesR**](DeriveAndSyncNewChangeAddressesR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeriveAndSyncNewReceivingAddresses

> DeriveAndSyncNewReceivingAddressesR DeriveAndSyncNewReceivingAddresses(ctx, blockchain, network).Context(context).DeriveAndSyncNewReceivingAddressesRB(deriveAndSyncNewReceivingAddressesRB).Execute()

Derive And Sync New Receiving Addresses



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
    deriveAndSyncNewReceivingAddressesRB := *openapiclient.NewDeriveAndSyncNewReceivingAddressesRB(*openapiclient.NewDeriveAndSyncNewReceivingAddressesRBData(*openapiclient.NewDeriveAndSyncNewReceivingAddressesRBDataItem("xpub6DSqNgZke91RZBXk9s8tBknGPiVB7AQqVyxHCLEM49D3VGeMWg6qmSDruSn7SgQApVH1M8cSvjXfDmhRysDt9hZWFAMcZf4X1DAzd23G4ZQ"))) // DeriveAndSyncNewReceivingAddressesRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HDWalletsApi.DeriveAndSyncNewReceivingAddresses(context.Background(), blockchain, network).Context(context).DeriveAndSyncNewReceivingAddressesRB(deriveAndSyncNewReceivingAddressesRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HDWalletsApi.DeriveAndSyncNewReceivingAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeriveAndSyncNewReceivingAddresses`: DeriveAndSyncNewReceivingAddressesR
    fmt.Fprintf(os.Stdout, "Response from `HDWalletsApi.DeriveAndSyncNewReceivingAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeriveAndSyncNewReceivingAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **deriveAndSyncNewReceivingAddressesRB** | [**DeriveAndSyncNewReceivingAddressesRB**](DeriveAndSyncNewReceivingAddressesRB.md) |  | 

### Return type

[**DeriveAndSyncNewReceivingAddressesR**](DeriveAndSyncNewReceivingAddressesR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHDWalletXPubYPubZPubAssets

> GetHDWalletXPubYPubZPubAssetsR GetHDWalletXPubYPubZPubAssets(ctx, blockchain, extendedPublicKey, network).Context(context).Derivation(derivation).Execute()

Get HD Wallet (xPub, yPub, zPub) Assets



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
    extendedPublicKey := "xpub68SyZPMPpZUy9QB2fk2J28b5Rwd6jeWKind3K8oziZuVcL7wWZiXZNCPKuh42ejSpTLYngQ9Gbzj9a1Ap2QQmoFs2sMSbUvkEr8D3GW7MrR" // string | Defines the account extended publicly known key which is used to derive all child public keys.
    network := "ropsten" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    derivation := "account" // string | The way how the HD walled derives, for example when the type is ACCOUNT, it derives change and receive addresses while when the type is BIP32 it derives directly. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HDWalletsApi.GetHDWalletXPubYPubZPubAssets(context.Background(), blockchain, extendedPublicKey, network).Context(context).Derivation(derivation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HDWalletsApi.GetHDWalletXPubYPubZPubAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHDWalletXPubYPubZPubAssets`: GetHDWalletXPubYPubZPubAssetsR
    fmt.Fprintf(os.Stdout, "Response from `HDWalletsApi.GetHDWalletXPubYPubZPubAssets`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetHDWalletXPubYPubZPubAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **derivation** | **string** | The way how the HD walled derives, for example when the type is ACCOUNT, it derives change and receive addresses while when the type is BIP32 it derives directly. | 

### Return type

[**GetHDWalletXPubYPubZPubAssetsR**](GetHDWalletXPubYPubZPubAssetsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    derivation := "derivation_example" // string | The way how the HD walled derives, for example when the type is ACCOUNT, it derives change and receive addresses while when the type is BIP32 it derives directly. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HDWalletsApi.GetHDWalletXPubYPubZPubDetails(context.Background(), blockchain, extendedPublicKey, network).Context(context).Derivation(derivation).Execute()
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
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    derivation := "derivation_example" // string | The way how the HD walled derives, for example when the type is ACCOUNT, it derives change and receive addresses while when the type is BIP32 it derives directly. (optional)
    limit := int64(50) // int64 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int64(0) // int64 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HDWalletsApi.ListHDWalletXPubYPubZPubTransactions(context.Background(), blockchain, extendedPublicKey, network).Context(context).Derivation(derivation).Limit(limit).Offset(offset).Execute()
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
 **limit** | **int64** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int64** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

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


## ListHDWalletXPubYPubZPubUTXOs

> ListHDWalletXPubYPubZPubUTXOsR ListHDWalletXPubYPubZPubUTXOs(ctx, blockchain, extendedPublicKey, network).Context(context).Derivation(derivation).Limit(limit).Offset(offset).Execute()

List HD Wallet (xPub, yPub, zPub) UTXOs



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
    extendedPublicKey := "tpubDDCs6jf3Tg8VTts6EBCNpibVanPQpSkmYRLAXVvuhcuC6ZcbYtEizqERj8D4TBukuvjNSjtjEbKYdtFuRG5WuisrirZG9m5L8wUvf4bHhgQ" // string | Defines the account extended publicly known key which is used to derive all child public keys.
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    derivation := "account" // string | The way how the HD walled derives, for example when the type is ACCOUNT, it derives change and receive addresses while when the type is BIP32 it derives directly. (optional)
    limit := int64(50) // int64 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int64(0) // int64 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HDWalletsApi.ListHDWalletXPubYPubZPubUTXOs(context.Background(), blockchain, extendedPublicKey, network).Context(context).Derivation(derivation).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HDWalletsApi.ListHDWalletXPubYPubZPubUTXOs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHDWalletXPubYPubZPubUTXOs`: ListHDWalletXPubYPubZPubUTXOsR
    fmt.Fprintf(os.Stdout, "Response from `HDWalletsApi.ListHDWalletXPubYPubZPubUTXOs`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiListHDWalletXPubYPubZPubUTXOsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **derivation** | **string** | The way how the HD walled derives, for example when the type is ACCOUNT, it derives change and receive addresses while when the type is BIP32 it derives directly. | 
 **limit** | **int64** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int64** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListHDWalletXPubYPubZPubUTXOsR**](ListHDWalletXPubYPubZPubUTXOsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncedAddresses

> ListSyncedAddressesR ListSyncedAddresses(ctx, blockchain, extendedPublicKey, network).Context(context).AddressFormat(addressFormat).IsChangeAddress(isChangeAddress).Limit(limit).Offset(offset).Execute()

List Synced Addresses



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
    extendedPublicKey := "tpubD9GMECjiZHCaF9NHSMAeMbQMXnM7CviEJZsYBuztVwsUjPHWjxewWAUXWV2UExaAtoEvQGXDBmVWo6ZHGtj6TsH6Pop7D9DskQwGHA1gu1w" // string | Defines the account extended publicly known key which is used to derive all child public keys.
    network := "ropsten" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    addressFormat := "standard" // string | Defines the address format value. (optional)
    isChangeAddress := false // bool | Defines if the address is change addres or not. (optional) (default to true)
    limit := int64(50) // int64 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int64(0) // int64 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HDWalletsApi.ListSyncedAddresses(context.Background(), blockchain, extendedPublicKey, network).Context(context).AddressFormat(addressFormat).IsChangeAddress(isChangeAddress).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HDWalletsApi.ListSyncedAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSyncedAddresses`: ListSyncedAddressesR
    fmt.Fprintf(os.Stdout, "Response from `HDWalletsApi.ListSyncedAddresses`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiListSyncedAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **addressFormat** | **string** | Defines the address format value. | 
 **isChangeAddress** | **bool** | Defines if the address is change addres or not. | [default to true]
 **limit** | **int64** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int64** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListSyncedAddressesR**](ListSyncedAddressesR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPub

> PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubR PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPub(ctx, blockchain, network).Context(context).PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRB(prepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRB).Execute()

Prepare A UTXO-Based Transaction From HD Wallet (xPub, yPub, zPub)



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
    blockchain := "bitcoin" // string | 
    network := "testnet" // string | 
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    prepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRB := *openapiclient.NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRB(*openapiclient.NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBData(*openapiclient.NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem(*openapiclient.NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee(), []openapiclient.PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItemRecipientsInner{*openapiclient.NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItemRecipientsInner("tb1q8wus03xdv3t6aknmsnpd0jmeu7dgh93j34pj5a", "0.00003")}, "tpubDCNoSqt3HF32yq8VU6mgapTuW1FzENZa3C5dKUF6WCQzubWz2nA1yxUhMQWkhhkD58Uc8YiuD8cmB3y5tihqAv4zT2GNyqKTNLchHJmsvt9"))) // PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HDWalletsApi.PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPub(context.Background(), blockchain, network).Context(context).PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRB(prepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HDWalletsApi.PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPub``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPub`: PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubR
    fmt.Fprintf(os.Stdout, "Response from `HDWalletsApi.PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPub`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** |  | 
**network** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **prepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRB** | [**PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRB**](PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRB.md) |  | 

### Return type

[**PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubR**](PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPub

> PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubR PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPub(ctx, blockchain, network).Context(context).PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRB(prepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRB).Execute()

Prepare An Account-Based Transaction From HD Wallet (xPub, yPub, zPub)



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
    blockchain := "ethereum" // string | 
    network := "ropsten" // string | 
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    prepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRB := *openapiclient.NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRB(*openapiclient.NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBData(*openapiclient.NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem("0.000003", *openapiclient.NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee(), "0x041c594a0cc194e826bef5411b29c7f27001b7e3", "0x03654A9E78771442CAdf8DB37ae60D6a12bAEa9f", "xpub6CsGdqTDEVRnLmpWN218HBwJqfhqSx46iA8ByzEA5Bz9jfwU3TSg9U7ambKgJyykvCraHQ6sAFAddMGFdPzhXrRanKbHnnkbDTyRPyn5gRJ"))) // PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HDWalletsApi.PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPub(context.Background(), blockchain, network).Context(context).PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRB(prepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HDWalletsApi.PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPub``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPub`: PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubR
    fmt.Fprintf(os.Stdout, "Response from `HDWalletsApi.PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPub`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** |  | 
**network** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **prepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRB** | [**PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRB**](PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRB.md) |  | 

### Return type

[**PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubR**](PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
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
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    syncHDWalletXPubYPubZPubRB := *openapiclient.NewSyncHDWalletXPubYPubZPubRB(*openapiclient.NewSyncHDWalletXPubYPubZPubRBData(*openapiclient.NewSyncHDWalletXPubYPubZPubRBDataItem("upub5Ei6bRNneqozk6smK7dvtXHC5PjUyEL4ynCfMKvjznLcXi9DQaikETzQjHvJC43XexMvQs64jxB1njMjCHpRZ4xQWAmv3ge9cVtjfsHmbvQ"))) // SyncHDWalletXPubYPubZPubRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HDWalletsApi.SyncHDWalletXPubYPubZPub(context.Background(), blockchain, network).Context(context).SyncHDWalletXPubYPubZPubRB(syncHDWalletXPubYPubZPubRB).Execute()
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


## SyncNewHDWalletXPubYPubZPub

> SyncNewHDWalletXPubYPubZPubR SyncNewHDWalletXPubYPubZPub(ctx, blockchain, network).Context(context).SyncNewHDWalletXPubYPubZPubRB(syncNewHDWalletXPubYPubZPubRB).Execute()

Sync New HD Wallet (xPub, yPub, zPub)



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
    syncNewHDWalletXPubYPubZPubRB := *openapiclient.NewSyncNewHDWalletXPubYPubZPubRB(*openapiclient.NewSyncHDWalletXPubYPubZPubRBData(*openapiclient.NewSyncHDWalletXPubYPubZPubRBDataItem("upub5Ei6bRNneqozk6smK7dvtXHC5PjUyEL4ynCfMKvjznLcXi9DQaikETzQjHvJC43XexMvQs64jxB1njMjCHpRZ4xQWAmv3ge9cVtjfsHmbvQ"))) // SyncNewHDWalletXPubYPubZPubRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HDWalletsApi.SyncNewHDWalletXPubYPubZPub(context.Background(), blockchain, network).Context(context).SyncNewHDWalletXPubYPubZPubRB(syncNewHDWalletXPubYPubZPubRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HDWalletsApi.SyncNewHDWalletXPubYPubZPub``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncNewHDWalletXPubYPubZPub`: SyncNewHDWalletXPubYPubZPubR
    fmt.Fprintf(os.Stdout, "Response from `HDWalletsApi.SyncNewHDWalletXPubYPubZPub`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncNewHDWalletXPubYPubZPubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **syncNewHDWalletXPubYPubZPubRB** | [**SyncNewHDWalletXPubYPubZPubRB**](SyncNewHDWalletXPubYPubZPubRB.md) |  | 

### Return type

[**SyncNewHDWalletXPubYPubZPubR**](SyncNewHDWalletXPubYPubZPubR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


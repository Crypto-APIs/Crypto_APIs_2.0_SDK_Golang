# \TransactionsApi

All URIs are relative to *https://rest.cryptoapis.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCoinsTransactionFromAddressForWholeAmount**](TransactionsApi.md#CreateCoinsTransactionFromAddressForWholeAmount) | **Post** /wallet-as-a-service/wallets/{walletId}/{blockchain}/{network}/addresses/{address}/all-transaction-requests | Create Coins Transaction From Address For Whole Amount
[**CreateCoinsTransactionRequestFromAddress**](TransactionsApi.md#CreateCoinsTransactionRequestFromAddress) | **Post** /wallet-as-a-service/wallets/{walletId}/{blockchain}/{network}/addresses/{address}/transaction-requests | Create Coins Transaction Request from Address
[**CreateCoinsTransactionRequestFromWallet**](TransactionsApi.md#CreateCoinsTransactionRequestFromWallet) | **Post** /wallet-as-a-service/wallets/{walletId}/{blockchain}/{network}/transaction-requests | Create Coins Transaction Request from Wallet
[**CreateFungibleTokensTransactionRequestFromAddress**](TransactionsApi.md#CreateFungibleTokensTransactionRequestFromAddress) | **Post** /wallet-as-a-service/wallets/{walletId}/{blockchain}/{network}/addresses/{senderAddress}/token-transaction-requests | Create Fungible Tokens Transaction Request from Address



## CreateCoinsTransactionFromAddressForWholeAmount

> CreateCoinsTransactionFromAddressForWholeAmountR CreateCoinsTransactionFromAddressForWholeAmount(ctx, address, blockchain, network, walletId).Context(context).CreateCoinsTransactionFromAddressForWholeAmountRB(createCoinsTransactionFromAddressForWholeAmountRB).Execute()

Create Coins Transaction From Address For Whole Amount



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
    address := "0x6f61e3c2fbb8c8be698bd0907ba6c04b62800fe5" // string | Defines the source address.
    blockchain := "ethereum" // string | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
    network := "ropsten" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    walletId := "609e221675d04500068718dc" // string | Represents the sender's specific and unique Wallet ID of the sender.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    createCoinsTransactionFromAddressForWholeAmountRB := *openapiclient.NewCreateCoinsTransactionFromAddressForWholeAmountRB(*openapiclient.NewCreateCoinsTransactionFromAddressForWholeAmountRBData(*openapiclient.NewCreateCoinsTransactionFromAddressForWholeAmountRBDataItem("standard", "0xc065b539490f81b6c297c37b1925c3be2f190732"))) // CreateCoinsTransactionFromAddressForWholeAmountRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.CreateCoinsTransactionFromAddressForWholeAmount(context.Background(), address, blockchain, network, walletId).Context(context).CreateCoinsTransactionFromAddressForWholeAmountRB(createCoinsTransactionFromAddressForWholeAmountRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.CreateCoinsTransactionFromAddressForWholeAmount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCoinsTransactionFromAddressForWholeAmount`: CreateCoinsTransactionFromAddressForWholeAmountR
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.CreateCoinsTransactionFromAddressForWholeAmount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Defines the source address. | 
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**walletId** | **string** | Represents the sender&#39;s specific and unique Wallet ID of the sender. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCoinsTransactionFromAddressForWholeAmountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **createCoinsTransactionFromAddressForWholeAmountRB** | [**CreateCoinsTransactionFromAddressForWholeAmountRB**](CreateCoinsTransactionFromAddressForWholeAmountRB.md) |  | 

### Return type

[**CreateCoinsTransactionFromAddressForWholeAmountR**](CreateCoinsTransactionFromAddressForWholeAmountR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCoinsTransactionRequestFromAddress

> CreateCoinsTransactionRequestFromAddressR CreateCoinsTransactionRequestFromAddress(ctx, address, blockchain, network, walletId).Context(context).CreateCoinsTransactionRequestFromAddressRB(createCoinsTransactionRequestFromAddressRB).Execute()

Create Coins Transaction Request from Address



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
    address := "0x6f61e3c2fbb8c8be698bd0907ba6c04b62800fe5" // string | Defines the specific source address for the transaction. For XRP we also support the X-address format.
    blockchain := "ethereum" // string | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
    network := "ropsten" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    walletId := "609e221675d04500068718dc" // string | Represents the sender's specific and unique Wallet ID of the sender.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    createCoinsTransactionRequestFromAddressRB := *openapiclient.NewCreateCoinsTransactionRequestFromAddressRB(*openapiclient.NewCreateCoinsTransactionRequestFromAddressRBData(*openapiclient.NewCreateCoinsTransactionRequestFromAddressRBDataItem("0.2", "slow", "0xc065b539490f81b6c297c37b1925c3be2f190732"))) // CreateCoinsTransactionRequestFromAddressRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.CreateCoinsTransactionRequestFromAddress(context.Background(), address, blockchain, network, walletId).Context(context).CreateCoinsTransactionRequestFromAddressRB(createCoinsTransactionRequestFromAddressRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.CreateCoinsTransactionRequestFromAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCoinsTransactionRequestFromAddress`: CreateCoinsTransactionRequestFromAddressR
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.CreateCoinsTransactionRequestFromAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Defines the specific source address for the transaction. For XRP we also support the X-address format. | 
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**walletId** | **string** | Represents the sender&#39;s specific and unique Wallet ID of the sender. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCoinsTransactionRequestFromAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **createCoinsTransactionRequestFromAddressRB** | [**CreateCoinsTransactionRequestFromAddressRB**](CreateCoinsTransactionRequestFromAddressRB.md) |  | 

### Return type

[**CreateCoinsTransactionRequestFromAddressR**](CreateCoinsTransactionRequestFromAddressR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCoinsTransactionRequestFromWallet

> CreateCoinsTransactionRequestFromWalletR CreateCoinsTransactionRequestFromWallet(ctx, blockchain, network, walletId).Context(context).CreateCoinsTransactionRequestFromWalletRB(createCoinsTransactionRequestFromWalletRB).Execute()

Create Coins Transaction Request from Wallet



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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks. (default to "testnet")
    walletId := "609e221675d04500068718dc" // string | Represents the sender's specific and unique Wallet ID of the sender.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    createCoinsTransactionRequestFromWalletRB := *openapiclient.NewCreateCoinsTransactionRequestFromWalletRB(*openapiclient.NewCreateCoinsTransactionRequestFromWalletRBData(*openapiclient.NewCreateCoinsTransactionRequestFromWalletRBDataItem("standard", []openapiclient.CreateCoinsTransactionRequestFromWalletRBDataItemRecipientsInner{*openapiclient.NewCreateCoinsTransactionRequestFromWalletRBDataItemRecipientsInner("2MtzNEqm2D9jcbPJ5mW7Z3AUNwqt3afZH66", "0.125")}))) // CreateCoinsTransactionRequestFromWalletRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.CreateCoinsTransactionRequestFromWallet(context.Background(), blockchain, network, walletId).Context(context).CreateCoinsTransactionRequestFromWalletRB(createCoinsTransactionRequestFromWalletRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.CreateCoinsTransactionRequestFromWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCoinsTransactionRequestFromWallet`: CreateCoinsTransactionRequestFromWalletR
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.CreateCoinsTransactionRequestFromWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | [default to &quot;testnet&quot;]
**walletId** | **string** | Represents the sender&#39;s specific and unique Wallet ID of the sender. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCoinsTransactionRequestFromWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **createCoinsTransactionRequestFromWalletRB** | [**CreateCoinsTransactionRequestFromWalletRB**](CreateCoinsTransactionRequestFromWalletRB.md) |  | 

### Return type

[**CreateCoinsTransactionRequestFromWalletR**](CreateCoinsTransactionRequestFromWalletR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFungibleTokensTransactionRequestFromAddress

> CreateFungibleTokensTransactionRequestFromAddressR CreateFungibleTokensTransactionRequestFromAddress(ctx, blockchain, network, senderAddress, walletId).Context(context).CreateFungibleTokensTransactionRequestFromAddressRB(createFungibleTokensTransactionRequestFromAddressRB).Execute()

Create Fungible Tokens Transaction Request from Address



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
    blockchain := "ethereum" // string | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. (default to "ethereum")
    network := "mainnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks. (default to "mainnet")
    senderAddress := "0x6f61e3c2fbb8c8be698bd0907ba6c04b62800fe5" // string | Defines the specific source address for the transaction.
    walletId := "609e221675d04500068718dc" // string | Defines the unique ID of the Wallet.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    createFungibleTokensTransactionRequestFromAddressRB := *openapiclient.NewCreateFungibleTokensTransactionRequestFromAddressRB(*openapiclient.NewCreateFungibleTokensTransactionRequestFromAddressRBData(*openapiclient.NewCreateFungibleTokensTransactionRequestFromAddressRBDataItem("0.2", "standard", "0xc065b539490f81b6c297c37b1925c3be2f190732", "0xdac17f958d2ee523a2206206994597c13d831ec7"))) // CreateFungibleTokensTransactionRequestFromAddressRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.CreateFungibleTokensTransactionRequestFromAddress(context.Background(), blockchain, network, senderAddress, walletId).Context(context).CreateFungibleTokensTransactionRequestFromAddressRB(createFungibleTokensTransactionRequestFromAddressRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.CreateFungibleTokensTransactionRequestFromAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFungibleTokensTransactionRequestFromAddress`: CreateFungibleTokensTransactionRequestFromAddressR
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.CreateFungibleTokensTransactionRequestFromAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | [default to &quot;ethereum&quot;]
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | [default to &quot;mainnet&quot;]
**senderAddress** | **string** | Defines the specific source address for the transaction. | 
**walletId** | **string** | Defines the unique ID of the Wallet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFungibleTokensTransactionRequestFromAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **createFungibleTokensTransactionRequestFromAddressRB** | [**CreateFungibleTokensTransactionRequestFromAddressRB**](CreateFungibleTokensTransactionRequestFromAddressRB.md) |  | 

### Return type

[**CreateFungibleTokensTransactionRequestFromAddressR**](CreateFungibleTokensTransactionRequestFromAddressR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


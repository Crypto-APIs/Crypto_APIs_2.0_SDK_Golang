# \TransactionsApi

All URIs are relative to *https://rest.cryptoapis.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCoinsTransactionRequestFromAddress**](TransactionsApi.md#CreateCoinsTransactionRequestFromAddress) | **Post** /wallet-as-a-service/wallets/{walletId}/{blockchain}/{network}/addresses/{address}/transaction-requests | Create Coins Transaction Request from Address
[**CreateCoinsTransactionRequestFromWallet**](TransactionsApi.md#CreateCoinsTransactionRequestFromWallet) | **Post** /wallet-as-a-service/wallets/{walletId}/{blockchain}/{network}/transaction-requests | Create Coins Transaction Request from Wallet
[**CreateTokensTransactionRequestFromAddress**](TransactionsApi.md#CreateTokensTransactionRequestFromAddress) | **Post** /wallet-as-a-service/wallets/{walletId}/{blockchain}/{network}/addresses/{senderAddress}/token-transaction-requests | Create Tokens Transaction Request from Address



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
    address := "0x6f61e3c2fbb8c8be698bd0907ba6c04b62800fe5" // string | Defines the specific source address for the transaction.
    blockchain := "ethereum" // string | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
    network := "network_example" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
    walletId := "609e221675d04500068718dc" // string | Represents the sender's specific and unique Wallet ID of the sender.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    createCoinsTransactionRequestFromAddressRB := *openapiclient.NewCreateCoinsTransactionRequestFromAddressRB(*openapiclient.NewCreateCoinsTransactionRequestFromAddressRBData(*openapiclient.NewCreateCoinsTransactionRequestFromAddressRBDataItem("0.2", "slow", "0xc065b539490f81b6c297c37b1925c3be2f190732"))) // CreateCoinsTransactionRequestFromAddressRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.CreateCoinsTransactionRequestFromAddress(context.Background(), address, blockchain, network, walletId).Context(context).CreateCoinsTransactionRequestFromAddressRB(createCoinsTransactionRequestFromAddressRB).Execute()
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
**address** | **string** | Defines the specific source address for the transaction. | 
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
    walletId := "609e221675d04500068718dc" // string | Represents the sender's specific and unique Wallet ID of the sender.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    createCoinsTransactionRequestFromWalletRB := *openapiclient.NewCreateCoinsTransactionRequestFromWalletRB(*openapiclient.NewCreateCoinsTransactionRequestFromWalletRBData(*openapiclient.NewCreateCoinsTransactionRequestFromWalletRBDataItem("standard", []openapiclient.CreateCoinsTransactionRequestFromWalletRBDataItemRecipients{*openapiclient.NewCreateCoinsTransactionRequestFromWalletRBDataItemRecipients("0x6f61e3c2fbb8c8be698bd0907ba6c04b62800fe5", "0.125")}))) // CreateCoinsTransactionRequestFromWalletRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.CreateCoinsTransactionRequestFromWallet(context.Background(), blockchain, network, walletId).Context(context).CreateCoinsTransactionRequestFromWalletRB(createCoinsTransactionRequestFromWalletRB).Execute()
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
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
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


## CreateTokensTransactionRequestFromAddress

> CreateTokensTransactionRequestFromAddressR CreateTokensTransactionRequestFromAddress(ctx, blockchain, network, senderAddress, walletId).Context(context).CreateTokensTransactionRequestFromAddressRB(createTokensTransactionRequestFromAddressRB).Execute()

Create Tokens Transaction Request from Address



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
    network := "network_example" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks. (default to "mainnet")
    senderAddress := "0x6f61e3c2fbb8c8be698bd0907ba6c04b62800fe5" // string | Defines the specific source address for the transaction.
    walletId := "609e221675d04500068718dc" // string | Defines the unique ID of the Wallet.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    createTokensTransactionRequestFromAddressRB := *openapiclient.NewCreateTokensTransactionRequestFromAddressRB(*openapiclient.NewCreateTokensTransactionRequestFromAddressRBData(*openapiclient.NewCreateTokensTransactionRequestFromAddressRBDataItem("0.2", "standard", "0xc065b539490f81b6c297c37b1925c3be2f190732", "1"))) // CreateTokensTransactionRequestFromAddressRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.CreateTokensTransactionRequestFromAddress(context.Background(), blockchain, network, senderAddress, walletId).Context(context).CreateTokensTransactionRequestFromAddressRB(createTokensTransactionRequestFromAddressRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.CreateTokensTransactionRequestFromAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTokensTransactionRequestFromAddress`: CreateTokensTransactionRequestFromAddressR
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.CreateTokensTransactionRequestFromAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | [default to &quot;ethereum&quot;]
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | [default to &quot;mainnet&quot;]
**senderAddress** | **string** | Defines the specific source address for the transaction. | 
**walletId** | **string** | Defines the unique ID of the Wallet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokensTransactionRequestFromAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **createTokensTransactionRequestFromAddressRB** | [**CreateTokensTransactionRequestFromAddressRB**](CreateTokensTransactionRequestFromAddressRB.md) |  | 

### Return type

[**CreateTokensTransactionRequestFromAddressR**](CreateTokensTransactionRequestFromAddressR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


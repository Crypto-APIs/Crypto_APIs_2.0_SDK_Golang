# \UnifiedEndpointsApi

All URIs are relative to *https://rest.cryptoapis.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EstimateTransactionSmartFee**](UnifiedEndpointsApi.md#EstimateTransactionSmartFee) | **Get** /blockchain-data/{blockchain}/{network}/estimate-transaction-smart-fee | Estimate Transaction Smart Fee
[**GetAddressDetails**](UnifiedEndpointsApi.md#GetAddressDetails) | **Get** /blockchain-data/{blockchain}/{network}/addresses/{address} | Get Address Details
[**GetBlockDetailsByBlockHash**](UnifiedEndpointsApi.md#GetBlockDetailsByBlockHash) | **Get** /blockchain-data/{blockchain}/{network}/blocks/hash/{blockHash} | Get Block Details By Block Hash
[**GetBlockDetailsByBlockHeight**](UnifiedEndpointsApi.md#GetBlockDetailsByBlockHeight) | **Get** /blockchain-data/{blockchain}/{network}/blocks/height/{height} | Get Block Details By Block Height
[**GetFeeRecommendations**](UnifiedEndpointsApi.md#GetFeeRecommendations) | **Get** /blockchain-data/{blockchain}/{network}/mempool/fees | Get Fee Recommendations
[**GetLastMinedBlock**](UnifiedEndpointsApi.md#GetLastMinedBlock) | **Get** /blockchain-data/{blockchain}/{network}/blocks/last | Get Last Mined Block
[**GetNextAvailableNonce**](UnifiedEndpointsApi.md#GetNextAvailableNonce) | **Get** /blockchain-data/{blockchain}/{network}/addresses/{address}/next-available-nonce | Get Next Available Nonce
[**GetRawTransactionData**](UnifiedEndpointsApi.md#GetRawTransactionData) | **Get** /blockchain-data/{blockchain}/{network}/transactions/{transactionId}/raw-data | Get Raw Transaction Data
[**GetTransactionDetailsByTransactionID**](UnifiedEndpointsApi.md#GetTransactionDetailsByTransactionID) | **Get** /blockchain-data/{blockchain}/{network}/transactions/{transactionId} | Get Transaction Details By Transaction ID
[**ListAllUnconfirmedTransactions**](UnifiedEndpointsApi.md#ListAllUnconfirmedTransactions) | **Get** /blockchain-data/{blockchain}/{network}/address-transactions-unconfirmed | List All Unconfirmed Transactions
[**ListConfirmedTokensTransfersByAddressAndTimeRange**](UnifiedEndpointsApi.md#ListConfirmedTokensTransfersByAddressAndTimeRange) | **Get** /blockchain-data/{blockchain}/{network}/addresses/{address}/tokens-transfers-by-time-range | List Confirmed Tokens Transfers By Address And Time Range
[**ListConfirmedTransactionsByAddress**](UnifiedEndpointsApi.md#ListConfirmedTransactionsByAddress) | **Get** /blockchain-data/{blockchain}/{network}/addresses/{address}/transactions | List Confirmed Transactions By Address
[**ListConfirmedTransactionsByAddressAndTimeRange**](UnifiedEndpointsApi.md#ListConfirmedTransactionsByAddressAndTimeRange) | **Get** /blockchain-data/{blockchain}/{network}/addresses/{address}/transactions-by-time-range | List Confirmed Transactions By Address And Time Range
[**ListInternalTransactionsByAddressAndTimeRange**](UnifiedEndpointsApi.md#ListInternalTransactionsByAddressAndTimeRange) | **Get** /blockchain-data/{blockchain}/{network}/addresses/{address}/internal-by-time-range | List Internal Transactions By Address And Time Range
[**ListLatestMinedBlocks**](UnifiedEndpointsApi.md#ListLatestMinedBlocks) | **Get** /blockchain-data/{blockchain}/{network}/blocks/last/{count} | List Latest Mined Blocks
[**ListTransactionsByBlockHash**](UnifiedEndpointsApi.md#ListTransactionsByBlockHash) | **Get** /blockchain-data/{blockchain}/{network}/blocks/hash/{blockHash}/transactions | List Transactions by Block Hash
[**ListTransactionsByBlockHeight**](UnifiedEndpointsApi.md#ListTransactionsByBlockHeight) | **Get** /blockchain-data/{blockchain}/{network}/blocks/height/{height}/transactions | List Transactions by Block Height
[**ListUnconfirmedTransactionsByAddress**](UnifiedEndpointsApi.md#ListUnconfirmedTransactionsByAddress) | **Get** /blockchain-data/{blockchain}/{network}/address-transactions-unconfirmed/{address} | List Unconfirmed Transactions by Address
[**ListUnspentTransactionOutputsByAddress**](UnifiedEndpointsApi.md#ListUnspentTransactionOutputsByAddress) | **Get** /blockchain-data/{blockchain}/{network}/addresses/{address}/unspent-outputs | List Unspent Transaction Outputs By Address



## EstimateTransactionSmartFee

> EstimateTransactionSmartFeeR EstimateTransactionSmartFee(ctx, blockchain, network).Context(context).ConfirmationTarget(confirmationTarget).EstimateMode(estimateMode).Execute()

Estimate Transaction Smart Fee



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
    confirmationTarget := int32(2) // int32 | Integer representation of confirmation target in blocks that estimation will be valid for (optional)
    estimateMode := "economical" // string | String representation of the address (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnifiedEndpointsApi.EstimateTransactionSmartFee(context.Background(), blockchain, network).Context(context).ConfirmationTarget(confirmationTarget).EstimateMode(estimateMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.EstimateTransactionSmartFee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EstimateTransactionSmartFee`: EstimateTransactionSmartFeeR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.EstimateTransactionSmartFee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** |  | 
**network** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEstimateTransactionSmartFeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **confirmationTarget** | **int32** | Integer representation of confirmation target in blocks that estimation will be valid for | 
 **estimateMode** | **string** | String representation of the address | 

### Return type

[**EstimateTransactionSmartFeeR**](EstimateTransactionSmartFeeR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddressDetails

> GetAddressDetailsR GetAddressDetails(ctx, blockchain, network, address).Context(context).Execute()

Get Address Details



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
    resp, r, err := apiClient.UnifiedEndpointsApi.GetAddressDetails(context.Background(), blockchain, network, address).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.GetAddressDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddressDetails`: GetAddressDetailsR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.GetAddressDetails`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAddressDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetAddressDetailsR**](GetAddressDetailsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockDetailsByBlockHash

> GetBlockDetailsByBlockHashR GetBlockDetailsByBlockHash(ctx, blockchain, network, blockHash).Context(context).Execute()

Get Block Details By Block Hash



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
    blockHash := "0000000006b3f483bec16b8a85c632bdd30a14a202c83a9148002c9ee441dd0c" // string | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnifiedEndpointsApi.GetBlockDetailsByBlockHash(context.Background(), blockchain, network, blockHash).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.GetBlockDetailsByBlockHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockDetailsByBlockHash`: GetBlockDetailsByBlockHashR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.GetBlockDetailsByBlockHash`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetBlockDetailsByBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetBlockDetailsByBlockHashR**](GetBlockDetailsByBlockHashR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockDetailsByBlockHeight

> GetBlockDetailsByBlockHeightR GetBlockDetailsByBlockHeight(ctx, blockchain, network, height).Context(context).Execute()

Get Block Details By Block Height



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
    height := int32(673852) // int32 | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \"Genesis block\".
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnifiedEndpointsApi.GetBlockDetailsByBlockHeight(context.Background(), blockchain, network, height).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.GetBlockDetailsByBlockHeight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockDetailsByBlockHeight`: GetBlockDetailsByBlockHeightR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.GetBlockDetailsByBlockHeight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**height** | **int32** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockDetailsByBlockHeightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetBlockDetailsByBlockHeightR**](GetBlockDetailsByBlockHeightR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeeRecommendations

> GetFeeRecommendationsR GetFeeRecommendations(ctx, blockchain, network).Context(context).Execute()

Get Fee Recommendations



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnifiedEndpointsApi.GetFeeRecommendations(context.Background(), blockchain, network).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.GetFeeRecommendations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeeRecommendations`: GetFeeRecommendationsR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.GetFeeRecommendations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeeRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetFeeRecommendationsR**](GetFeeRecommendationsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastMinedBlock

> GetLastMinedBlockR GetLastMinedBlock(ctx, blockchain, network).Context(context).Execute()

Get Last Mined Block



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnifiedEndpointsApi.GetLastMinedBlock(context.Background(), blockchain, network).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.GetLastMinedBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLastMinedBlock`: GetLastMinedBlockR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.GetLastMinedBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastMinedBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetLastMinedBlockR**](GetLastMinedBlockR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNextAvailableNonce

> GetNextAvailableNonceR GetNextAvailableNonce(ctx, blockchain, network, address).Context(context).Execute()

Get Next Available Nonce



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
    address := "0x626b046b0ce356f248b215b01b459f8b8d59e1a4" // string | Represents the public address, which is a compressed and shortened form of a public key.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnifiedEndpointsApi.GetNextAvailableNonce(context.Background(), blockchain, network, address).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.GetNextAvailableNonce``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNextAvailableNonce`: GetNextAvailableNonceR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.GetNextAvailableNonce`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetNextAvailableNonceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetNextAvailableNonceR**](GetNextAvailableNonceR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawTransactionData

> GetRawTransactionDataR GetRawTransactionData(ctx, blockchain, network, transactionId).Context(context).Execute()

Get Raw Transaction Data



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
    transactionId := "4b66461bf88b61e1e4326356534c135129defb504c7acb2fd6c92697d79eb250" // string | Represents the unique identifier of a transaction, i.e. it could be transactionId in UTXO-based protocols like Bitcoin, and transaction hash in Ethereum blockchain.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnifiedEndpointsApi.GetRawTransactionData(context.Background(), blockchain, network, transactionId).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.GetRawTransactionData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRawTransactionData`: GetRawTransactionDataR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.GetRawTransactionData`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetRawTransactionDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetRawTransactionDataR**](GetRawTransactionDataR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionDetailsByTransactionID

> GetTransactionDetailsByTransactionIDR GetTransactionDetailsByTransactionID(ctx, blockchain, network, transactionId).Context(context).Execute()

Get Transaction Details By Transaction ID



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
    transactionId := "4b66461bf88b61e1e4326356534c135129defb504c7acb2fd6c92697d79eb250" // string | Represents the unique identifier of a transaction, i.e. it could be `transactionId` in UTXO-based protocols like Bitcoin, and transaction `hash` in Ethereum blockchain.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnifiedEndpointsApi.GetTransactionDetailsByTransactionID(context.Background(), blockchain, network, transactionId).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.GetTransactionDetailsByTransactionID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionDetailsByTransactionID`: GetTransactionDetailsByTransactionIDR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.GetTransactionDetailsByTransactionID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**transactionId** | **string** | Represents the unique identifier of a transaction, i.e. it could be &#x60;transactionId&#x60; in UTXO-based protocols like Bitcoin, and transaction &#x60;hash&#x60; in Ethereum blockchain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionDetailsByTransactionIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetTransactionDetailsByTransactionIDR**](GetTransactionDetailsByTransactionIDR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllUnconfirmedTransactions

> ListAllUnconfirmedTransactionsR ListAllUnconfirmedTransactions(ctx, blockchain, network).Context(context).Limit(limit).Offset(offset).Execute()

List All Unconfirmed Transactions



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
    resp, r, err := apiClient.UnifiedEndpointsApi.ListAllUnconfirmedTransactions(context.Background(), blockchain, network).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.ListAllUnconfirmedTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllUnconfirmedTransactions`: ListAllUnconfirmedTransactionsR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.ListAllUnconfirmedTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllUnconfirmedTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int64** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int64** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListAllUnconfirmedTransactionsR**](ListAllUnconfirmedTransactionsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfirmedTokensTransfersByAddressAndTimeRange

> ListConfirmedTokensTransfersByAddressAndTimeRangeR ListConfirmedTokensTransfersByAddressAndTimeRange(ctx, blockchain, network, address).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).Context(context).Limit(limit).Offset(offset).Execute()

List Confirmed Tokens Transfers By Address And Time Range



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
    blockchain := "ethereum" // string | Represents the specific blockchain protocol name, e.g. Ethereum, Ethereum Classic, etc.
    network := "ropsten" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    address := "0x033ef6db9fbd0ee60e2931906b987fe0280471a0" // string | Represents the public address, which is a compressed and shortened form of a public key.
    fromTimestamp := int32(1236238648) // int32 | Defines the specific time/date from which the results will start being listed.
    toTimestamp := int32(1644417868) // int32 | Defines the specific time/date to which the results will be listed.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int32(50) // int32 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int32(0) // int32 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnifiedEndpointsApi.ListConfirmedTokensTransfersByAddressAndTimeRange(context.Background(), blockchain, network, address).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.ListConfirmedTokensTransfersByAddressAndTimeRange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConfirmedTokensTransfersByAddressAndTimeRange`: ListConfirmedTokensTransfersByAddressAndTimeRangeR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.ListConfirmedTokensTransfersByAddressAndTimeRange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Ethereum Classic, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**address** | **string** | Represents the public address, which is a compressed and shortened form of a public key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConfirmedTokensTransfersByAddressAndTimeRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fromTimestamp** | **int32** | Defines the specific time/date from which the results will start being listed. | 
 **toTimestamp** | **int32** | Defines the specific time/date to which the results will be listed. | 
 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int32** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListConfirmedTokensTransfersByAddressAndTimeRangeR**](ListConfirmedTokensTransfersByAddressAndTimeRangeR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfirmedTransactionsByAddress

> ListConfirmedTransactionsByAddressR ListConfirmedTransactionsByAddress(ctx, blockchain, network, address).Context(context).Limit(limit).Offset(offset).Execute()

List Confirmed Transactions By Address



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
    address := "mho4jHBcrNCncKt38trJahXakuaBnS7LK5" // string | Represents the public address, which is a compressed and shortened form of a public key.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int64(50) // int64 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int64(0) // int64 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnifiedEndpointsApi.ListConfirmedTransactionsByAddress(context.Background(), blockchain, network, address).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.ListConfirmedTransactionsByAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConfirmedTransactionsByAddress`: ListConfirmedTransactionsByAddressR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.ListConfirmedTransactionsByAddress`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiListConfirmedTransactionsByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int64** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int64** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListConfirmedTransactionsByAddressR**](ListConfirmedTransactionsByAddressR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfirmedTransactionsByAddressAndTimeRange

> ListConfirmedTransactionsByAddressAndTimeRangeR ListConfirmedTransactionsByAddressAndTimeRange(ctx, blockchain, network, address).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).Context(context).Limit(limit).Offset(offset).Execute()

List Confirmed Transactions By Address And Time Range



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
    address := "mho4jHBcrNCncKt38trJahXakuaBnS7LK5" // string | Represents the public address, which is a compressed and shortened form of a public key.
    fromTimestamp := int32(1236238648) // int32 | Defines the specific time/date from which the results will start being listed.
    toTimestamp := int32(1644417868) // int32 | Defines the specific time/date to which the results will be listed.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int64(50) // int64 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int64(0) // int64 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnifiedEndpointsApi.ListConfirmedTransactionsByAddressAndTimeRange(context.Background(), blockchain, network, address).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.ListConfirmedTransactionsByAddressAndTimeRange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConfirmedTransactionsByAddressAndTimeRange`: ListConfirmedTransactionsByAddressAndTimeRangeR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.ListConfirmedTransactionsByAddressAndTimeRange`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiListConfirmedTransactionsByAddressAndTimeRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fromTimestamp** | **int32** | Defines the specific time/date from which the results will start being listed. | 
 **toTimestamp** | **int32** | Defines the specific time/date to which the results will be listed. | 
 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int64** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int64** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListConfirmedTransactionsByAddressAndTimeRangeR**](ListConfirmedTransactionsByAddressAndTimeRangeR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInternalTransactionsByAddressAndTimeRange

> ListInternalTransactionsByAddressAndTimeRangeR ListInternalTransactionsByAddressAndTimeRange(ctx, blockchain, network, address).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).Context(context).Limit(limit).Offset(offset).Execute()

List Internal Transactions By Address And Time Range



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
    address := "0xc8fe2ceac93ad50e496b497357ae5385192dd28d" // string | String identifier of the address document represented in CryptoAPIs
    fromTimestamp := int32(1635979828) // int32 | Defines the specific time/date from which the results will start being listed.
    toTimestamp := int32(1643329896) // int32 | Defines the specific time/date to which the results will be listed.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int32(50) // int32 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int32(0) // int32 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnifiedEndpointsApi.ListInternalTransactionsByAddressAndTimeRange(context.Background(), blockchain, network, address).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.ListInternalTransactionsByAddressAndTimeRange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInternalTransactionsByAddressAndTimeRange`: ListInternalTransactionsByAddressAndTimeRangeR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.ListInternalTransactionsByAddressAndTimeRange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**address** | **string** | String identifier of the address document represented in CryptoAPIs | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInternalTransactionsByAddressAndTimeRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fromTimestamp** | **int32** | Defines the specific time/date from which the results will start being listed. | 
 **toTimestamp** | **int32** | Defines the specific time/date to which the results will be listed. | 
 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int32** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListInternalTransactionsByAddressAndTimeRangeR**](ListInternalTransactionsByAddressAndTimeRangeR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLatestMinedBlocks

> ListLatestMinedBlocksR ListLatestMinedBlocks(ctx, network, blockchain, count).Context(context).Execute()

List Latest Mined Blocks



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
    network := "ropsten" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks
    blockchain := "ethereum" // string | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
    count := int32(2) // int32 | Specifies how many records were requested.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnifiedEndpointsApi.ListLatestMinedBlocks(context.Background(), network, blockchain, count).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.ListLatestMinedBlocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLatestMinedBlocks`: ListLatestMinedBlocksR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.ListLatestMinedBlocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks | 
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**count** | **int32** | Specifies how many records were requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLatestMinedBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**ListLatestMinedBlocksR**](ListLatestMinedBlocksR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionsByBlockHash

> ListTransactionsByBlockHashR ListTransactionsByBlockHash(ctx, blockchain, network, blockHash).Context(context).Limit(limit).Offset(offset).Execute()

List Transactions by Block Hash



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
    blockHash := "00000000000000127080d8bcf84f4ad830a71ea0aadce3632579b6b2f26cd94b" // string | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int64(50) // int64 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int64(0) // int64 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnifiedEndpointsApi.ListTransactionsByBlockHash(context.Background(), blockchain, network, blockHash).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.ListTransactionsByBlockHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionsByBlockHash`: ListTransactionsByBlockHashR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.ListTransactionsByBlockHash`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiListTransactionsByBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int64** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int64** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListTransactionsByBlockHashR**](ListTransactionsByBlockHashR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionsByBlockHeight

> ListTransactionsByBlockHeightR ListTransactionsByBlockHeight(ctx, blockchain, network, height).Context(context).Limit(limit).Offset(offset).Execute()

List Transactions by Block Height



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
    height := int64(673852) // int64 | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \"Genesis block\".
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int64(50) // int64 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int64(0) // int64 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnifiedEndpointsApi.ListTransactionsByBlockHeight(context.Background(), blockchain, network, height).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.ListTransactionsByBlockHeight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionsByBlockHeight`: ListTransactionsByBlockHeightR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.ListTransactionsByBlockHeight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**height** | **int64** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsByBlockHeightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int64** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int64** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListTransactionsByBlockHeightR**](ListTransactionsByBlockHeightR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUnconfirmedTransactionsByAddress

> ListUnconfirmedTransactionsByAddressR ListUnconfirmedTransactionsByAddress(ctx, blockchain, network, address).Context(context).Limit(limit).Offset(offset).Execute()

List Unconfirmed Transactions by Address



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
    address := "2NDt6eztswLiVgVYaGUhkTPmugUGovVypAe" // string | Represents the public address, which is a compressed and shortened form of a public key.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int64(50) // int64 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int64(0) // int64 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnifiedEndpointsApi.ListUnconfirmedTransactionsByAddress(context.Background(), blockchain, network, address).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.ListUnconfirmedTransactionsByAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUnconfirmedTransactionsByAddress`: ListUnconfirmedTransactionsByAddressR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.ListUnconfirmedTransactionsByAddress`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiListUnconfirmedTransactionsByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int64** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int64** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListUnconfirmedTransactionsByAddressR**](ListUnconfirmedTransactionsByAddressR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUnspentTransactionOutputsByAddress

> ListUnspentTransactionOutputsByAddressR ListUnspentTransactionOutputsByAddress(ctx, blockchain, network, address).Context(context).Limit(limit).Offset(offset).Execute()

List Unspent Transaction Outputs By Address



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
    address := "2MtzNEqm2D9jcbPJ5mW7Z3AUNwqt3afZH66" // string | Represents the address that has unspend funds per which the result is returned.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int32(50) // int32 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int32(0) // int32 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnifiedEndpointsApi.ListUnspentTransactionOutputsByAddress(context.Background(), blockchain, network, address).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.ListUnspentTransactionOutputsByAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUnspentTransactionOutputsByAddress`: ListUnspentTransactionOutputsByAddressR
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.ListUnspentTransactionOutputsByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**address** | **string** | Represents the address that has unspend funds per which the result is returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUnspentTransactionOutputsByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int32** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListUnspentTransactionOutputsByAddressR**](ListUnspentTransactionOutputsByAddressR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


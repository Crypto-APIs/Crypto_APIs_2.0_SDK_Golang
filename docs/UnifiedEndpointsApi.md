# \UnifiedEndpointsApi

All URIs are relative to *https://rest.cryptoapis.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddressDetails**](UnifiedEndpointsApi.md#GetAddressDetails) | **Get** /blockchain-data/{blockchain}/{network}/addresses/{address} | Get Address Details
[**GetBlockDetailsByBlockHash**](UnifiedEndpointsApi.md#GetBlockDetailsByBlockHash) | **Get** /blockchain-data/{blockchain}/{network}/blocks/hash/{blockHash} | Get Block Details By Block Hash
[**GetBlockDetailsByBlockHeight**](UnifiedEndpointsApi.md#GetBlockDetailsByBlockHeight) | **Get** /blockchain-data/{blockchain}/{network}/blocks/height/{height} | Get Block Details By Block Height
[**GetFeeRecommendations**](UnifiedEndpointsApi.md#GetFeeRecommendations) | **Get** /blockchain-data/{blockchain}/{network}/mempool/fees | Get Fee Recommendations
[**GetLatestMinedBlock**](UnifiedEndpointsApi.md#GetLatestMinedBlock) | **Get** /blockchain-data/{blockchain}/{network}/blocks/last | Get Latest Mined Block
[**GetTransactionDetailsByTransactionID**](UnifiedEndpointsApi.md#GetTransactionDetailsByTransactionID) | **Get** /blockchain-data/{blockchain}/{network}/transactions/{transactionId} | Get Transaction Details By Transaction ID
[**ListTransactionsByAddress**](UnifiedEndpointsApi.md#ListTransactionsByAddress) | **Get** /blockchain-data/{blockchain}/{network}/addresses/{address}/transactions | List Transactions By Address
[**ListTransactionsByBlockHash**](UnifiedEndpointsApi.md#ListTransactionsByBlockHash) | **Get** /blockchain-data/{blockchain}/{network}/blocks/hash/{blockHash}/transactions | List Transactions by Block Hash
[**ListTransactionsByBlockHeight**](UnifiedEndpointsApi.md#ListTransactionsByBlockHeight) | **Get** /blockchain-data/{blockchain}/{network}/blocks/height/{height}/transactions | List Transactions by Block Height



## GetAddressDetails

> GetAddressDetailsResponse GetAddressDetails(ctx, blockchain, network, address).Context(context).Execute()

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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
    address := "mzYijhgmzZrmuB7wBDazRKirnChKyow4M3" // string | Represents the public address, which is a compressed and shortened form of a public key.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnifiedEndpointsApi.GetAddressDetails(context.Background(), blockchain, network, address).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.GetAddressDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddressDetails`: GetAddressDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.GetAddressDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**address** | **string** | Represents the public address, which is a compressed and shortened form of a public key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetAddressDetailsResponse**](GetAddressDetailsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockDetailsByBlockHash

> GetBlockDetailsByBlockHashResponse GetBlockDetailsByBlockHash(ctx, blockchain, network, blockHash).Context(context).Execute()

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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
    blockHash := "0000000006b3f483bec16b8a85c632bdd30a14a202c83a9148002c9ee441dd0c" // string | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnifiedEndpointsApi.GetBlockDetailsByBlockHash(context.Background(), blockchain, network, blockHash).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.GetBlockDetailsByBlockHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockDetailsByBlockHash`: GetBlockDetailsByBlockHashResponse
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.GetBlockDetailsByBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**blockHash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockDetailsByBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetBlockDetailsByBlockHashResponse**](GetBlockDetailsByBlockHashResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockDetailsByBlockHeight

> GetBlockDetailsByBlockHeightResponse GetBlockDetailsByBlockHeight(ctx, blockchain, network, height).Context(context).Execute()

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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
    height := int32(673852) // int32 | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \"Genesis block\".
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnifiedEndpointsApi.GetBlockDetailsByBlockHeight(context.Background(), blockchain, network, height).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.GetBlockDetailsByBlockHeight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockDetailsByBlockHeight`: GetBlockDetailsByBlockHeightResponse
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.GetBlockDetailsByBlockHeight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**height** | **int32** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockDetailsByBlockHeightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetBlockDetailsByBlockHeightResponse**](GetBlockDetailsByBlockHeightResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeeRecommendations

> GetFeeRecommendationsResponse GetFeeRecommendations(ctx, blockchain, network).Context(context).Execute()

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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnifiedEndpointsApi.GetFeeRecommendations(context.Background(), blockchain, network).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.GetFeeRecommendations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeeRecommendations`: GetFeeRecommendationsResponse
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.GetFeeRecommendations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeeRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetFeeRecommendationsResponse**](GetFeeRecommendationsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestMinedBlock

> GetLatestMinedBlockResponse GetLatestMinedBlock(ctx, blockchain, network).Context(context).Execute()

Get Latest Mined Block



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnifiedEndpointsApi.GetLatestMinedBlock(context.Background(), blockchain, network).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.GetLatestMinedBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestMinedBlock`: GetLatestMinedBlockResponse
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.GetLatestMinedBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestMinedBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetLatestMinedBlockResponse**](GetLatestMinedBlockResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionDetailsByTransactionID

> GetTransactionDetailsByTransactionIDResponse GetTransactionDetailsByTransactionID(ctx, blockchain, network, transactionId).Context(context).Execute()

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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
    transactionId := "4b66461bf88b61e1e4326356534c135129defb504c7acb2fd6c92697d79eb250" // string | Represents the unique identifier of a transaction, i.e. it could be `transactionId` in UTXO-based protocols like Bitcoin, and transaction `hash` in Ethereum blockchain.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnifiedEndpointsApi.GetTransactionDetailsByTransactionID(context.Background(), blockchain, network, transactionId).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.GetTransactionDetailsByTransactionID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionDetailsByTransactionID`: GetTransactionDetailsByTransactionIDResponse
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.GetTransactionDetailsByTransactionID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**transactionId** | **string** | Represents the unique identifier of a transaction, i.e. it could be &#x60;transactionId&#x60; in UTXO-based protocols like Bitcoin, and transaction &#x60;hash&#x60; in Ethereum blockchain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionDetailsByTransactionIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetTransactionDetailsByTransactionIDResponse**](GetTransactionDetailsByTransactionIDResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionsByAddress

> ListTransactionsByAddressResponse ListTransactionsByAddress(ctx, blockchain, network, address).Context(context).Limit(limit).Offset(offset).Execute()

List Transactions By Address



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
    address := "mho4jHBcrNCncKt38trJahXakuaBnS7LK5" // string | Represents the public address, which is a compressed and shortened form of a public key.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int32(50) // int32 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int32(10) // int32 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnifiedEndpointsApi.ListTransactionsByAddress(context.Background(), blockchain, network, address).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.ListTransactionsByAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionsByAddress`: ListTransactionsByAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.ListTransactionsByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**address** | **string** | Represents the public address, which is a compressed and shortened form of a public key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int32** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListTransactionsByAddressResponse**](ListTransactionsByAddressResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionsByBlockHash

> ListTransactionsByBlockHashResponse ListTransactionsByBlockHash(ctx, blockchain, network, blockHash).Context(context).Limit(limit).Offset(offset).Execute()

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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
    blockHash := "00000000000000127080d8bcf84f4ad830a71ea0aadce3632579b6b2f26cd94b" // string | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int32(50) // int32 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int32(10) // int32 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnifiedEndpointsApi.ListTransactionsByBlockHash(context.Background(), blockchain, network, blockHash).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.ListTransactionsByBlockHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionsByBlockHash`: ListTransactionsByBlockHashResponse
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.ListTransactionsByBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**blockHash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsByBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int32** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListTransactionsByBlockHashResponse**](ListTransactionsByBlockHashResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionsByBlockHeight

> ListTransactionsByBlockHeightResponse ListTransactionsByBlockHeight(ctx, blockchain, network, height).Context(context).Limit(limit).Offset(offset).Execute()

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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
    height := int32(673852) // int32 | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \"Genesis block\".
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int32(50) // int32 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int32(10) // int32 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnifiedEndpointsApi.ListTransactionsByBlockHeight(context.Background(), blockchain, network, height).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedEndpointsApi.ListTransactionsByBlockHeight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionsByBlockHeight`: ListTransactionsByBlockHeightResponse
    fmt.Fprintf(os.Stdout, "Response from `UnifiedEndpointsApi.ListTransactionsByBlockHeight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**height** | **int32** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsByBlockHeightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int32** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListTransactionsByBlockHeightResponse**](ListTransactionsByBlockHeightResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \XRPRippleApi

All URIs are relative to *https://rest.cryptoapis.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLatestMinedXRPRippleBlock**](XRPRippleApi.md#GetLatestMinedXRPRippleBlock) | **Get** /blockchain-data/xrp/{network}/blocks/last | Get Latest Mined XRP (Ripple) Block
[**GetXRPRippleAddressDetails**](XRPRippleApi.md#GetXRPRippleAddressDetails) | **Get** /blockchain-data/xrp/{network}/addresses/{address} | Get XRP (Ripple) Address Details
[**GetXRPRippleBlockDetailsByBlockHash**](XRPRippleApi.md#GetXRPRippleBlockDetailsByBlockHash) | **Get** /blockchain-data/xrp/{network}/blocks/hash/{blockHash} | Get XRP (Ripple) Block Details By Block Hash
[**GetXRPRippleBlockDetailsByBlockHeight**](XRPRippleApi.md#GetXRPRippleBlockDetailsByBlockHeight) | **Get** /blockchain-data/xrp/{network}/blocks/height/{height} | Get XRP (Ripple) Block Details By Block Height
[**GetXRPRippleTransactionDetailsByTransactionID**](XRPRippleApi.md#GetXRPRippleTransactionDetailsByTransactionID) | **Get** /blockchain-data/xrp/{network}/transactions/{transactionHash} | Get XRP (Ripple) Transaction Details By Transaction ID
[**ListXRPRippleTransactionsByAddress**](XRPRippleApi.md#ListXRPRippleTransactionsByAddress) | **Get** /blockchain-data/xrp/{network}/addresses/{address}/transactions | List XRP (Ripple) Transactions by Address
[**ListXRPRippleTransactionsByBlockHash**](XRPRippleApi.md#ListXRPRippleTransactionsByBlockHash) | **Get** /blockchain-data/xrp/{network}/blocks/hash/{blockHash}/transactions | List XRP (Ripple) Transactions By Block Hash



## GetLatestMinedXRPRippleBlock

> GetLatestMinedXRPRippleBlockResponse GetLatestMinedXRPRippleBlock(ctx, network).Context(context).Execute()

Get Latest Mined XRP (Ripple) Block



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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XRPRippleApi.GetLatestMinedXRPRippleBlock(context.Background(), network).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XRPRippleApi.GetLatestMinedXRPRippleBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestMinedXRPRippleBlock`: GetLatestMinedXRPRippleBlockResponse
    fmt.Fprintf(os.Stdout, "Response from `XRPRippleApi.GetLatestMinedXRPRippleBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestMinedXRPRippleBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetLatestMinedXRPRippleBlockResponse**](GetLatestMinedXRPRippleBlockResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetXRPRippleAddressDetails

> GetXRPRippleAddressDetailsResponse GetXRPRippleAddressDetails(ctx, network, address).Context(context).Execute()

Get XRP (Ripple) Address Details



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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\",  are test networks.
    address := "rA9bXGJcXvZKaWofrRphdJsBWzhyCfH3z" // string | Represents the public address, which is a compressed and shortened form of a public key.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XRPRippleApi.GetXRPRippleAddressDetails(context.Background(), network, address).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XRPRippleApi.GetXRPRippleAddressDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetXRPRippleAddressDetails`: GetXRPRippleAddressDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `XRPRippleApi.GetXRPRippleAddressDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;,  are test networks. | 
**address** | **string** | Represents the public address, which is a compressed and shortened form of a public key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetXRPRippleAddressDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetXRPRippleAddressDetailsResponse**](GetXRPRippleAddressDetailsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetXRPRippleBlockDetailsByBlockHash

> GetXRPRippleBlockDetailsByBlockHashResponse GetXRPRippleBlockDetailsByBlockHash(ctx, network, blockHash).Context(context).Execute()

Get XRP (Ripple) Block Details By Block Hash



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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\",  are test networks.
    blockHash := "1ab0614d2a438da8b23086cbceef7d443edbd295d9c7619fc8a19c7618bc22c9" // string | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XRPRippleApi.GetXRPRippleBlockDetailsByBlockHash(context.Background(), network, blockHash).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XRPRippleApi.GetXRPRippleBlockDetailsByBlockHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetXRPRippleBlockDetailsByBlockHash`: GetXRPRippleBlockDetailsByBlockHashResponse
    fmt.Fprintf(os.Stdout, "Response from `XRPRippleApi.GetXRPRippleBlockDetailsByBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;,  are test networks. | 
**blockHash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetXRPRippleBlockDetailsByBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetXRPRippleBlockDetailsByBlockHashResponse**](GetXRPRippleBlockDetailsByBlockHashResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetXRPRippleBlockDetailsByBlockHeight

> GetXRPRippleBlockDetailsByBlockHeightResponse GetXRPRippleBlockDetailsByBlockHeight(ctx, network, height).Context(context).Execute()

Get XRP (Ripple) Block Details By Block Height



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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\",  are test networks.
    height := "15886156" // string | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \"Genesis block\".
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XRPRippleApi.GetXRPRippleBlockDetailsByBlockHeight(context.Background(), network, height).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XRPRippleApi.GetXRPRippleBlockDetailsByBlockHeight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetXRPRippleBlockDetailsByBlockHeight`: GetXRPRippleBlockDetailsByBlockHeightResponse
    fmt.Fprintf(os.Stdout, "Response from `XRPRippleApi.GetXRPRippleBlockDetailsByBlockHeight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;,  are test networks. | 
**height** | **string** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetXRPRippleBlockDetailsByBlockHeightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetXRPRippleBlockDetailsByBlockHeightResponse**](GetXRPRippleBlockDetailsByBlockHeightResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetXRPRippleTransactionDetailsByTransactionID

> GetXRPRippleTransactionDetailsByTransactionIDResponse GetXRPRippleTransactionDetailsByTransactionID(ctx, network, transactionHash).Context(context).Execute()

Get XRP (Ripple) Transaction Details By Transaction ID



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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
    transactionHash := "36a1737481edec87bacc3101dfb752ae2c76f9171e7edebe587e330c1ea77c8d" // string | Represents the same as `transactionId` for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols `hash` is different from `transactionId` for SegWit transactions.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XRPRippleApi.GetXRPRippleTransactionDetailsByTransactionID(context.Background(), network, transactionHash).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XRPRippleApi.GetXRPRippleTransactionDetailsByTransactionID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetXRPRippleTransactionDetailsByTransactionID`: GetXRPRippleTransactionDetailsByTransactionIDResponse
    fmt.Fprintf(os.Stdout, "Response from `XRPRippleApi.GetXRPRippleTransactionDetailsByTransactionID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**transactionHash** | **string** | Represents the same as &#x60;transactionId&#x60; for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols &#x60;hash&#x60; is different from &#x60;transactionId&#x60; for SegWit transactions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetXRPRippleTransactionDetailsByTransactionIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetXRPRippleTransactionDetailsByTransactionIDResponse**](GetXRPRippleTransactionDetailsByTransactionIDResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListXRPRippleTransactionsByAddress

> ListXRPRippleTransactionsByAddressResponse ListXRPRippleTransactionsByAddress(ctx, network, address).Context(context).Limit(limit).Offset(offset).Execute()

List XRP (Ripple) Transactions by Address



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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\",  are test networks.
    address := "rA9bXGJcXvZKaWofrRphdJsBWzhyCfH3z" // string | Represents the public address, which is a compressed and shortened form of a public key.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int32(50) // int32 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int32(10) // int32 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XRPRippleApi.ListXRPRippleTransactionsByAddress(context.Background(), network, address).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XRPRippleApi.ListXRPRippleTransactionsByAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListXRPRippleTransactionsByAddress`: ListXRPRippleTransactionsByAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `XRPRippleApi.ListXRPRippleTransactionsByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;,  are test networks. | 
**address** | **string** | Represents the public address, which is a compressed and shortened form of a public key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListXRPRippleTransactionsByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int32** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListXRPRippleTransactionsByAddressResponse**](ListXRPRippleTransactionsByAddressResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListXRPRippleTransactionsByBlockHash

> ListXRPRippleTransactionsByBlockHashResponse ListXRPRippleTransactionsByBlockHash(ctx, network, blockHash).Context(context).Limit(limit).Offset(offset).Execute()

List XRP (Ripple) Transactions By Block Hash



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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
    blockHash := "14754656235f865a74eba27791fd41a47bdfe07fe811ff6d78f53db32e129e39" // string | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int32(50) // int32 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int32(10) // int32 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XRPRippleApi.ListXRPRippleTransactionsByBlockHash(context.Background(), network, blockHash).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XRPRippleApi.ListXRPRippleTransactionsByBlockHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListXRPRippleTransactionsByBlockHash`: ListXRPRippleTransactionsByBlockHashResponse
    fmt.Fprintf(os.Stdout, "Response from `XRPRippleApi.ListXRPRippleTransactionsByBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**blockHash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListXRPRippleTransactionsByBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int32** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListXRPRippleTransactionsByBlockHashResponse**](ListXRPRippleTransactionsByBlockHashResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


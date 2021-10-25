# \ZilliqaApi

All URIs are relative to *https://rest.cryptoapis.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLatestMinedZilliqaBlock**](ZilliqaApi.md#GetLatestMinedZilliqaBlock) | **Get** /blockchain-data/zilliqa-specific/{network}/blocks/last | Get Latest Mined Zilliqa Block
[**GetZilliqaAddressDetails**](ZilliqaApi.md#GetZilliqaAddressDetails) | **Get** /blockchain-data/zilliqa-specific/{network}/addresses/{address} | Get Zilliqa Address Details
[**GetZilliqaBlockDetailsByBlockHash**](ZilliqaApi.md#GetZilliqaBlockDetailsByBlockHash) | **Get** /blockchain-data/zilliqa-specific/{network}/blocks/hash/{blockHash} | Get Zilliqa Block Details By Block Hash
[**GetZilliqaBlockDetailsByBlockHeight**](ZilliqaApi.md#GetZilliqaBlockDetailsByBlockHeight) | **Get** /blockchain-data/zilliqa-specific/{network}/blocks/height/{blockHeight} | Get Zilliqa Block Details By Block Height
[**GetZilliqaTransactionDetailsByTransactionID**](ZilliqaApi.md#GetZilliqaTransactionDetailsByTransactionID) | **Get** /blockchain-data/zilliqa-specific/{network}/transactions/{transactionHash} | Get Zilliqa Transaction Details by Transaction ID
[**ListZilliqaTransactionsByAddress**](ZilliqaApi.md#ListZilliqaTransactionsByAddress) | **Get** /blockchain-data/zilliqa-specific/{network}/addresses/{address}/transactions | List Zilliqa Transactions by Address
[**ListZilliqaTransactionsByBlockHash**](ZilliqaApi.md#ListZilliqaTransactionsByBlockHash) | **Get** /blockchain-data/zilliqa-specific/{network}/blocks/hash/{blockHash}/transactions | List Zilliqa Transactions By Block Hash
[**ListZilliqaTransactionsByBlockHeight**](ZilliqaApi.md#ListZilliqaTransactionsByBlockHeight) | **Get** /blockchain-data/zilliqa-specific/{network}/blocks/height/{blockHeight}/transactions | List Zilliqa Transactions By Block Height



## GetLatestMinedZilliqaBlock

> GetLatestMinedZilliqaBlockR GetLatestMinedZilliqaBlock(ctx, network).Context(context).Execute()

Get Latest Mined Zilliqa Block



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
    network := "mainnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZilliqaApi.GetLatestMinedZilliqaBlock(context.Background(), network).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZilliqaApi.GetLatestMinedZilliqaBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestMinedZilliqaBlock`: GetLatestMinedZilliqaBlockR
    fmt.Fprintf(os.Stdout, "Response from `ZilliqaApi.GetLatestMinedZilliqaBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestMinedZilliqaBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetLatestMinedZilliqaBlockR**](GetLatestMinedZilliqaBlockR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZilliqaAddressDetails

> GetZilliqaAddressDetailsR GetZilliqaAddressDetails(ctx, network, address).Context(context).Execute()

Get Zilliqa Address Details



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
    network := "mainnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    address := "zil17v9gs6ctn9pq6wka8k668umdlaj88nrfej7rnm" // string | Defines the specific transaction's address.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZilliqaApi.GetZilliqaAddressDetails(context.Background(), network, address).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZilliqaApi.GetZilliqaAddressDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZilliqaAddressDetails`: GetZilliqaAddressDetailsR
    fmt.Fprintf(os.Stdout, "Response from `ZilliqaApi.GetZilliqaAddressDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**address** | **string** | Defines the specific transaction&#39;s address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZilliqaAddressDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetZilliqaAddressDetailsR**](GetZilliqaAddressDetailsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZilliqaBlockDetailsByBlockHash

> GetZilliqaBlockDetailsByBlockHashR GetZilliqaBlockDetailsByBlockHash(ctx, network, blockHash).Context(context).Execute()

Get Zilliqa Block Details By Block Hash



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
    network := "mainnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    blockHash := "0xf679d0b5387f0b0b3c3c1f368305512b23860888ba4415063d464a09b8bb6205" // string | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZilliqaApi.GetZilliqaBlockDetailsByBlockHash(context.Background(), network, blockHash).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZilliqaApi.GetZilliqaBlockDetailsByBlockHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZilliqaBlockDetailsByBlockHash`: GetZilliqaBlockDetailsByBlockHashR
    fmt.Fprintf(os.Stdout, "Response from `ZilliqaApi.GetZilliqaBlockDetailsByBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**blockHash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZilliqaBlockDetailsByBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetZilliqaBlockDetailsByBlockHashR**](GetZilliqaBlockDetailsByBlockHashR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZilliqaBlockDetailsByBlockHeight

> GetZilliqaBlockDetailsByBlockHeightR GetZilliqaBlockDetailsByBlockHeight(ctx, network, blockHeight).Context(context).Execute()

Get Zilliqa Block Details By Block Height



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
    network := "mainnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    blockHeight := int32(1244298) // int32 | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \"Genesis block\".
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZilliqaApi.GetZilliqaBlockDetailsByBlockHeight(context.Background(), network, blockHeight).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZilliqaApi.GetZilliqaBlockDetailsByBlockHeight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZilliqaBlockDetailsByBlockHeight`: GetZilliqaBlockDetailsByBlockHeightR
    fmt.Fprintf(os.Stdout, "Response from `ZilliqaApi.GetZilliqaBlockDetailsByBlockHeight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**blockHeight** | **int32** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZilliqaBlockDetailsByBlockHeightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetZilliqaBlockDetailsByBlockHeightR**](GetZilliqaBlockDetailsByBlockHeightR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZilliqaTransactionDetailsByTransactionID

> GetZilliqaTransactionDetailsByTransactionIDR GetZilliqaTransactionDetailsByTransactionID(ctx, network, transactionHash).Context(context).Execute()

Get Zilliqa Transaction Details by Transaction ID



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
    network := "mainnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    transactionHash := "0xe649afbe2849c70a7a8b3087bee2c7b2fa0b6b77ab658b4fff390a1d06f0a8bc" // string | String identifier of the transaction
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZilliqaApi.GetZilliqaTransactionDetailsByTransactionID(context.Background(), network, transactionHash).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZilliqaApi.GetZilliqaTransactionDetailsByTransactionID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZilliqaTransactionDetailsByTransactionID`: GetZilliqaTransactionDetailsByTransactionIDR
    fmt.Fprintf(os.Stdout, "Response from `ZilliqaApi.GetZilliqaTransactionDetailsByTransactionID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**transactionHash** | **string** | String identifier of the transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZilliqaTransactionDetailsByTransactionIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetZilliqaTransactionDetailsByTransactionIDR**](GetZilliqaTransactionDetailsByTransactionIDR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListZilliqaTransactionsByAddress

> ListZilliqaTransactionsByAddressR ListZilliqaTransactionsByAddress(ctx, network, address).Context(context).Limit(limit).Offset(offset).Execute()

List Zilliqa Transactions by Address



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
    network := "mainnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    address := "zil17v9gs6ctn9pq6wka8k668umdlaj88nrfej7rnm" // string | Defines the specific address of the sender.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int32(50) // int32 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int32(10) // int32 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZilliqaApi.ListZilliqaTransactionsByAddress(context.Background(), network, address).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZilliqaApi.ListZilliqaTransactionsByAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListZilliqaTransactionsByAddress`: ListZilliqaTransactionsByAddressR
    fmt.Fprintf(os.Stdout, "Response from `ZilliqaApi.ListZilliqaTransactionsByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**address** | **string** | Defines the specific address of the sender. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListZilliqaTransactionsByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int32** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListZilliqaTransactionsByAddressR**](ListZilliqaTransactionsByAddressR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListZilliqaTransactionsByBlockHash

> ListZilliqaTransactionsByBlockHashR ListZilliqaTransactionsByBlockHash(ctx, network, blockHash).Context(context).Limit(limit).Offset(offset).Execute()

List Zilliqa Transactions By Block Hash



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
    network := "mainnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    blockHash := "0x3160ca05ac12f60023de6b6987596eccbf084629f1a92b59534c5c021bf27041" // string | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int32(50) // int32 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int32(10) // int32 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZilliqaApi.ListZilliqaTransactionsByBlockHash(context.Background(), network, blockHash).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZilliqaApi.ListZilliqaTransactionsByBlockHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListZilliqaTransactionsByBlockHash`: ListZilliqaTransactionsByBlockHashR
    fmt.Fprintf(os.Stdout, "Response from `ZilliqaApi.ListZilliqaTransactionsByBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**blockHash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListZilliqaTransactionsByBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int32** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListZilliqaTransactionsByBlockHashR**](ListZilliqaTransactionsByBlockHashR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListZilliqaTransactionsByBlockHeight

> ListZilliqaTransactionsByBlockHeightR ListZilliqaTransactionsByBlockHeight(ctx, network, blockHeight).Context(context).Limit(limit).Offset(offset).Execute()

List Zilliqa Transactions By Block Height



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
    network := "mainnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    blockHeight := int32(1250392) // int32 | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \"Genesis block\".
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int32(50) // int32 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int32(10) // int32 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZilliqaApi.ListZilliqaTransactionsByBlockHeight(context.Background(), network, blockHeight).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZilliqaApi.ListZilliqaTransactionsByBlockHeight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListZilliqaTransactionsByBlockHeight`: ListZilliqaTransactionsByBlockHeightR
    fmt.Fprintf(os.Stdout, "Response from `ZilliqaApi.ListZilliqaTransactionsByBlockHeight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**blockHeight** | **int32** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListZilliqaTransactionsByBlockHeightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **limit** | **int32** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]

### Return type

[**ListZilliqaTransactionsByBlockHeightR**](ListZilliqaTransactionsByBlockHeightR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


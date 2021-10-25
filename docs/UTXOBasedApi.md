# \UTXOBasedApi

All URIs are relative to *https://rest.cryptoapis.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUnspentTransactionOutputsByAddress**](UTXOBasedApi.md#ListUnspentTransactionOutputsByAddress) | **Get** /blockchain-data/{blockchain}/{network}/addresses/{address}/unspent | List Unspent Transaction Outputs By Address



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
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
    address := "2MtzNEqm2D9jcbPJ5mW7Z3AUNwqt3afZH66" // string | Represents the public address, which is a compressed and shortened form of a public key.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    limit := int32(50) // int32 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int32(10) // int32 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UTXOBasedApi.ListUnspentTransactionOutputsByAddress(context.Background(), blockchain, network, address).Context(context).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UTXOBasedApi.ListUnspentTransactionOutputsByAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUnspentTransactionOutputsByAddress`: ListUnspentTransactionOutputsByAddressR
    fmt.Fprintf(os.Stdout, "Response from `UTXOBasedApi.ListUnspentTransactionOutputsByAddress`: %v\n", resp)
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


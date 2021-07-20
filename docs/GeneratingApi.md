# \GeneratingApi

All URIs are relative to *https://rest.cryptoapis.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateReceivingAddress**](GeneratingApi.md#GenerateReceivingAddress) | **Post** /wallet-as-a-service/wallets/{walletId}/{blockchain}/{network}/addresses | Generate Receiving Address



## GenerateReceivingAddress

> GenerateReceivingAddressR GenerateReceivingAddress(ctx, blockchain, network, walletId).Context(context).GenerateReceivingAddressRB(generateReceivingAddressRB).Execute()

Generate Receiving Address



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
    walletId := "60c9d9921c38030006675ff6" // string | Represents the unique ID of the specific Wallet.
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    generateReceivingAddressRB := *openapiclient.NewGenerateReceivingAddressRB(*openapiclient.NewGenerateReceivingAddressRBData(*openapiclient.NewGenerateReceivingAddressRBDataItem("yourLabelStringHere"))) // GenerateReceivingAddressRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneratingApi.GenerateReceivingAddress(context.Background(), blockchain, network, walletId).Context(context).GenerateReceivingAddressRB(generateReceivingAddressRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneratingApi.GenerateReceivingAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateReceivingAddress`: GenerateReceivingAddressR
    fmt.Fprintf(os.Stdout, "Response from `GeneratingApi.GenerateReceivingAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**walletId** | **string** | Represents the unique ID of the specific Wallet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateReceivingAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **generateReceivingAddressRB** | [**GenerateReceivingAddressRB**](GenerateReceivingAddressRB.md) |  | 

### Return type

[**GenerateReceivingAddressR**](GenerateReceivingAddressR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


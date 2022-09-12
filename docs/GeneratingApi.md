# \GeneratingApi

All URIs are relative to *https://rest.cryptoapis.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateDepositAddress**](GeneratingApi.md#GenerateDepositAddress) | **Post** /wallet-as-a-service/wallets/{walletId}/{blockchain}/{network}/addresses | Generate Deposit Address



## GenerateDepositAddress

> GenerateDepositAddressR GenerateDepositAddress(ctx, blockchain, network, walletId).Context(context).GenerateDepositAddressRB(generateDepositAddressRB).Execute()

Generate Deposit Address



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
    walletId := "60c9d9921c38030006675ff6" // string | Represents the unique ID of the specific Wallet.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    generateDepositAddressRB := *openapiclient.NewGenerateDepositAddressRB(*openapiclient.NewGenerateDepositAddressRBData(*openapiclient.NewGenerateDepositAddressRBDataItem("yourLabelStringHere"))) // GenerateDepositAddressRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneratingApi.GenerateDepositAddress(context.Background(), blockchain, network, walletId).Context(context).GenerateDepositAddressRB(generateDepositAddressRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneratingApi.GenerateDepositAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateDepositAddress`: GenerateDepositAddressR
    fmt.Fprintf(os.Stdout, "Response from `GeneratingApi.GenerateDepositAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**walletId** | **string** | Represents the unique ID of the specific Wallet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDepositAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **generateDepositAddressRB** | [**GenerateDepositAddressRB**](GenerateDepositAddressRB.md) |  | 

### Return type

[**GenerateDepositAddressR**](GenerateDepositAddressR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


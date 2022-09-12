# \ExchangeRatesApi

All URIs are relative to *https://rest.cryptoapis.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExchangeRateByAssetSymbols**](ExchangeRatesApi.md#GetExchangeRateByAssetSymbols) | **Get** /market-data/exchange-rates/by-symbols/{fromAssetSymbol}/{toAssetSymbol} | Get Exchange Rate By Asset Symbols
[**GetExchangeRateByAssetsIDs**](ExchangeRatesApi.md#GetExchangeRateByAssetsIDs) | **Get** /market-data/exchange-rates/by-asset-ids/{fromAssetId}/{toAssetId} | Get Exchange Rate By Assets IDs



## GetExchangeRateByAssetSymbols

> GetExchangeRateByAssetSymbolsR GetExchangeRateByAssetSymbols(ctx, fromAssetSymbol, toAssetSymbol).Context(context).CalculationTimestamp(calculationTimestamp).Execute()

Get Exchange Rate By Asset Symbols



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
    fromAssetSymbol := "btc" // string | Defines the base asset symbol to get a rate for.
    toAssetSymbol := "usd" // string | Defines the relation asset symbol in which the base asset rate will be displayed.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    calculationTimestamp := int32(1635514425) // int32 | Defines the time of the market data used to calculate the exchange rate in UNIX Timestamp. Oldest possible timestamp is 30 days. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExchangeRatesApi.GetExchangeRateByAssetSymbols(context.Background(), fromAssetSymbol, toAssetSymbol).Context(context).CalculationTimestamp(calculationTimestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangeRatesApi.GetExchangeRateByAssetSymbols``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeRateByAssetSymbols`: GetExchangeRateByAssetSymbolsR
    fmt.Fprintf(os.Stdout, "Response from `ExchangeRatesApi.GetExchangeRateByAssetSymbols`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromAssetSymbol** | **string** | Defines the base asset symbol to get a rate for. | 
**toAssetSymbol** | **string** | Defines the relation asset symbol in which the base asset rate will be displayed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeRateByAssetSymbolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **calculationTimestamp** | **int32** | Defines the time of the market data used to calculate the exchange rate in UNIX Timestamp. Oldest possible timestamp is 30 days. | 

### Return type

[**GetExchangeRateByAssetSymbolsR**](GetExchangeRateByAssetSymbolsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeRateByAssetsIDs

> GetExchangeRateByAssetsIDsR GetExchangeRateByAssetsIDs(ctx, fromAssetId, toAssetId).Context(context).CalculationTimestamp(calculationTimestamp).Execute()

Get Exchange Rate By Assets IDs



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
    fromAssetId := "5b1ea92e584bf50020130612" // string | Defines the base asset Reference ID to get a rate for.
    toAssetId := "5b1ea92e584bf50020130615" // string | Defines the relation asset Reference ID in which the base asset rate will be displayed.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    calculationTimestamp := int32(1618577849) // int32 | Defines the time of the market data used to calculate the exchange rate in UNIX Timestamp. Oldest possible timestamp is 30 days. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExchangeRatesApi.GetExchangeRateByAssetsIDs(context.Background(), fromAssetId, toAssetId).Context(context).CalculationTimestamp(calculationTimestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangeRatesApi.GetExchangeRateByAssetsIDs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeRateByAssetsIDs`: GetExchangeRateByAssetsIDsR
    fmt.Fprintf(os.Stdout, "Response from `ExchangeRatesApi.GetExchangeRateByAssetsIDs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromAssetId** | **string** | Defines the base asset Reference ID to get a rate for. | 
**toAssetId** | **string** | Defines the relation asset Reference ID in which the base asset rate will be displayed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeRateByAssetsIDsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **calculationTimestamp** | **int32** | Defines the time of the market data used to calculate the exchange rate in UNIX Timestamp. Oldest possible timestamp is 30 days. | 

### Return type

[**GetExchangeRateByAssetsIDsR**](GetExchangeRateByAssetsIDsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \AssetsApi

All URIs are relative to *https://rest.cryptoapis.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAssetDetailsByAssetID**](AssetsApi.md#GetAssetDetailsByAssetID) | **Get** /market-data/assets/assetId/{assetId} | Get Asset Details By Asset ID
[**GetAssetDetailsByAssetSymbol**](AssetsApi.md#GetAssetDetailsByAssetSymbol) | **Get** /market-data/assets/{assetSymbol} | Get Asset Details By Asset Symbol



## GetAssetDetailsByAssetID

> GetAssetDetailsByAssetIDR GetAssetDetailsByAssetID(ctx, assetId).Context(context).Execute()

Get Asset Details By Asset ID



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
    assetId := "5b1ea92e584bf50020130612" // string | Defines the unique ID of the specific asset.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.GetAssetDetailsByAssetID(context.Background(), assetId).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.GetAssetDetailsByAssetID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetDetailsByAssetID`: GetAssetDetailsByAssetIDR
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.GetAssetDetailsByAssetID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | Defines the unique ID of the specific asset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetDetailsByAssetIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetAssetDetailsByAssetIDR**](GetAssetDetailsByAssetIDR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetDetailsByAssetSymbol

> GetAssetDetailsByAssetSymbolR GetAssetDetailsByAssetSymbol(ctx, assetSymbol).Context(context).Execute()

Get Asset Details By Asset Symbol



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
    assetSymbol := "BTC" // string | Specifies the asset's unique symbol in the Crypto APIs listings.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.GetAssetDetailsByAssetSymbol(context.Background(), assetSymbol).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.GetAssetDetailsByAssetSymbol``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetDetailsByAssetSymbol`: GetAssetDetailsByAssetSymbolR
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.GetAssetDetailsByAssetSymbol`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetSymbol** | **string** | Specifies the asset&#39;s unique symbol in the Crypto APIs listings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetDetailsByAssetSymbolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetAssetDetailsByAssetSymbolR**](GetAssetDetailsByAssetSymbolR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


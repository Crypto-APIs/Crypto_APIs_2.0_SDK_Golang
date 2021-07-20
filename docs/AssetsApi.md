# \AssetsApi

All URIs are relative to *https://rest.cryptoapis.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAssetsDetails**](AssetsApi.md#ListAssetsDetails) | **Get** /market-data/assets/details | List Assets Details



## ListAssetsDetails

> ListAssetsDetailsR ListAssetsDetails(ctx).Context(context).AssetType(assetType).CryptoType(cryptoType).Limit(limit).Offset(offset).WaasEnabled(waasEnabled).Execute()

List Assets Details



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
    context := "context_example" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    assetType := "crypto" // string | Defines the type of the supported asset. This could be either \"crypto\" or \"fiat\". (optional)
    cryptoType := "coin" // string | Subtype of the crypto assets. Could be COIN or TOKEN (optional)
    limit := int32(50) // int32 | Defines how many items should be returned in the response per page basis. (optional) (default to 50)
    offset := int32(10) // int32 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional) (default to 0)
    waasEnabled := true // bool | Show only if WaaS is/not enabled (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetsApi.ListAssetsDetails(context.Background()).Context(context).AssetType(assetType).CryptoType(cryptoType).Limit(limit).Offset(offset).WaasEnabled(waasEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.ListAssetsDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAssetsDetails`: ListAssetsDetailsR
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.ListAssetsDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssetsDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **assetType** | **string** | Defines the type of the supported asset. This could be either \&quot;crypto\&quot; or \&quot;fiat\&quot;. | 
 **cryptoType** | **string** | Subtype of the crypto assets. Could be COIN or TOKEN | 
 **limit** | **int32** | Defines how many items should be returned in the response per page basis. | [default to 50]
 **offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | [default to 0]
 **waasEnabled** | **bool** | Show only if WaaS is/not enabled | 

### Return type

[**ListAssetsDetailsR**](ListAssetsDetailsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


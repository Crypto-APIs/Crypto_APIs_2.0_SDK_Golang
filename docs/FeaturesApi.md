# \FeaturesApi

All URIs are relative to *https://rest.cryptoapis.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BroadcastLocallySignedTransaction**](FeaturesApi.md#BroadcastLocallySignedTransaction) | **Post** /blockchain-tools/{blockchain}/{network}/transactions/broadcast | Broadcast Locally Signed Transaction
[**ConvertBitcoinCashAddress**](FeaturesApi.md#ConvertBitcoinCashAddress) | **Post** /blockchain-tools/{blockchain}/{network}/address/convert | Convert Bitcoin Cash Address
[**DecodeRawTransactionHex**](FeaturesApi.md#DecodeRawTransactionHex) | **Post** /blockchain-tools/{blockchain}/{network}/decode-raw-transaction | Decode Raw Transaction Hex
[**DecodeXAddress**](FeaturesApi.md#DecodeXAddress) | **Get** /blockchain-tools/{blockchain}/{network}/decode-x-address/{xAddress} | Decode X-Address
[**DeriveHDWalletXPubYPubZPubChangeOrReceivingAddresses**](FeaturesApi.md#DeriveHDWalletXPubYPubZPubChangeOrReceivingAddresses) | **Get** /blockchain-tools/{blockchain}/{network}/hd/{extendedPublicKey}/addresses/derive-address | Derive HD Wallet (xPub, yPub, zPub) Change Or Receiving Addresses
[**EncodeXAddress**](FeaturesApi.md#EncodeXAddress) | **Get** /blockchain-tools/{blockchain}/{network}/encode-x-address/{classicAddress}/{addressTag} | Encode X-Address
[**EstimateGasLimit**](FeaturesApi.md#EstimateGasLimit) | **Post** /blockchain-tools/{blockchain}/{network}/gas-limit | Estimate Gas Limit
[**EstimateTokenGasLimit**](FeaturesApi.md#EstimateTokenGasLimit) | **Post** /blockchain-tools/{blockchain}/{network}/gas-limit/contract | Estimate Token Gas Limit
[**GetEIP1559FeeRecommendations**](FeaturesApi.md#GetEIP1559FeeRecommendations) | **Get** /blockchain-tools/{blockchain}/{network}/fees/eip1559 | Get EIP 1559 Fee Recommendations
[**ValidateAddress**](FeaturesApi.md#ValidateAddress) | **Post** /blockchain-tools/{blockchain}/{network}/addresses/validate | Validate Address



## BroadcastLocallySignedTransaction

> BroadcastLocallySignedTransactionR BroadcastLocallySignedTransaction(ctx, blockchain, network).Context(context).BroadcastLocallySignedTransactionRB(broadcastLocallySignedTransactionRB).Execute()

Broadcast Locally Signed Transaction



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
    broadcastLocallySignedTransactionRB := *openapiclient.NewBroadcastLocallySignedTransactionRB(*openapiclient.NewBroadcastLocallySignedTransactionRBData(*openapiclient.NewBroadcastLocallySignedTransactionRBDataItem("0xf86a22827d00831e8480941b85a43e2e7f52e766ddfdfa2b901c42cb1201be8801b27f33b807c0008029a084ccbf02b27e0842fb1eda7a187a5589c3759be0e969e0ca989dc469a5e5e394a02e111e1156b197f1de4c1d9ba4af26e50665ea6d617d05b3e4047da12b915e69"))) // BroadcastLocallySignedTransactionRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.BroadcastLocallySignedTransaction(context.Background(), blockchain, network).Context(context).BroadcastLocallySignedTransactionRB(broadcastLocallySignedTransactionRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.BroadcastLocallySignedTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BroadcastLocallySignedTransaction`: BroadcastLocallySignedTransactionR
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.BroadcastLocallySignedTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastLocallySignedTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **broadcastLocallySignedTransactionRB** | [**BroadcastLocallySignedTransactionRB**](BroadcastLocallySignedTransactionRB.md) |  | 

### Return type

[**BroadcastLocallySignedTransactionR**](BroadcastLocallySignedTransactionR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConvertBitcoinCashAddress

> ConvertBitcoinCashAddressR ConvertBitcoinCashAddress(ctx, blockchain, network).Context(context).ConvertBitcoinCashAddressRB(convertBitcoinCashAddressRB).Execute()

Convert Bitcoin Cash Address



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
    blockchain := "bitcoin-cash" // string | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    convertBitcoinCashAddressRB := *openapiclient.NewConvertBitcoinCashAddressRB(*openapiclient.NewConvertBitcoinCashAddressRBData(*openapiclient.NewConvertBitcoinCashAddressRBDataItem("bchtest:qpcgz3zt5zp5dj7vd9ms24xquamncvhnxvlz97eee8"))) // ConvertBitcoinCashAddressRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.ConvertBitcoinCashAddress(context.Background(), blockchain, network).Context(context).ConvertBitcoinCashAddressRB(convertBitcoinCashAddressRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.ConvertBitcoinCashAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConvertBitcoinCashAddress`: ConvertBitcoinCashAddressR
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.ConvertBitcoinCashAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConvertBitcoinCashAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **convertBitcoinCashAddressRB** | [**ConvertBitcoinCashAddressRB**](ConvertBitcoinCashAddressRB.md) |  | 

### Return type

[**ConvertBitcoinCashAddressR**](ConvertBitcoinCashAddressR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DecodeRawTransactionHex

> DecodeRawTransactionHexR DecodeRawTransactionHex(ctx, blockchain, network).Context(context).DecodeRawTransactionHexRB(decodeRawTransactionHexRB).Execute()

Decode Raw Transaction Hex



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
    decodeRawTransactionHexRB := *openapiclient.NewDecodeRawTransactionHexRB(*openapiclient.NewDecodeRawTransactionHexRBData(*openapiclient.NewDecodeRawTransactionHexRBDataItem("0100000001f3f6a909f8521adb57d898d2985834e632374e770fd9e2b98656f1bf1fdfd427010000006b48304502203a776322ebf8eb8b58cc6ced4f2574f4c73aa664edce0b0022690f2f6f47c521022100b82353305988cb0ebd443089a173ceec93fe4dbfe98d74419ecc84a6a698e31d012103c5c1bc61f60ce3d6223a63cedbece03b12ef9f0068f2f3c4a7e7f06c523c3664ffffffff0260e31600000000001976a914977ae6e32349b99b72196cb62b5ef37329ed81b488ac063d1000000000001976a914f76bc4190f3d8e2315e5c11c59cfc8be9df747e388ac00000000"))) // DecodeRawTransactionHexRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.DecodeRawTransactionHex(context.Background(), blockchain, network).Context(context).DecodeRawTransactionHexRB(decodeRawTransactionHexRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.DecodeRawTransactionHex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DecodeRawTransactionHex`: DecodeRawTransactionHexR
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.DecodeRawTransactionHex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDecodeRawTransactionHexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **decodeRawTransactionHexRB** | [**DecodeRawTransactionHexRB**](DecodeRawTransactionHexRB.md) |  | 

### Return type

[**DecodeRawTransactionHexR**](DecodeRawTransactionHexR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DecodeXAddress

> DecodeXAddressR DecodeXAddress(ctx, blockchain, network, xAddress).Context(context).Execute()

Decode X-Address



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
    blockchain := "xrp" // string | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    xAddress := "TVTMSyg6nRscAm2JtRd8hnpF9nD21CgZx6ibb9iy3EWHotV" // string | Represents the encoded classic address with its destination tag.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.DecodeXAddress(context.Background(), blockchain, network, xAddress).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.DecodeXAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DecodeXAddress`: DecodeXAddressR
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.DecodeXAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**xAddress** | **string** | Represents the encoded classic address with its destination tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDecodeXAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**DecodeXAddressR**](DecodeXAddressR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeriveHDWalletXPubYPubZPubChangeOrReceivingAddresses

> DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesR DeriveHDWalletXPubYPubZPubChangeOrReceivingAddresses(ctx, blockchain, extendedPublicKey, network).Context(context).AddressFormat(addressFormat).AddressesCount(addressesCount).IsChange(isChange).StartIndex(startIndex).Execute()

Derive HD Wallet (xPub, yPub, zPub) Change Or Receiving Addresses



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
    extendedPublicKey := "upub5Ei6bRNneqozk6smK7dvtXHC5PjUyEL4ynCfMKvjznLcXi9DQaikETzQjHvJC43XexMvQs64jxB1njMjCHpRZ4xQWAmv3ge9cVtjfsHmbvQ" // string | Defines the account extended publicly known key which is used to derive all child public keys.
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    addressFormat := "p2sh" // string | Represents the format of the address. (optional)
    addressesCount := int64(2) // int64 | Represents the addresses count. (optional)
    isChange := true // bool | Defines if the specific address is a change or deposit address. If the value is True - it is a change address, if it is False - it is a Deposit address. (optional)
    startIndex := int64(3) // int64 | The starting index of the response items, i.e. where the response should start listing the returned items. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.DeriveHDWalletXPubYPubZPubChangeOrReceivingAddresses(context.Background(), blockchain, extendedPublicKey, network).Context(context).AddressFormat(addressFormat).AddressesCount(addressesCount).IsChange(isChange).StartIndex(startIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.DeriveHDWalletXPubYPubZPubChangeOrReceivingAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeriveHDWalletXPubYPubZPubChangeOrReceivingAddresses`: DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesR
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.DeriveHDWalletXPubYPubZPubChangeOrReceivingAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**extendedPublicKey** | **string** | Defines the account extended publicly known key which is used to derive all child public keys. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **addressFormat** | **string** | Represents the format of the address. | 
 **addressesCount** | **int64** | Represents the addresses count. | 
 **isChange** | **bool** | Defines if the specific address is a change or deposit address. If the value is True - it is a change address, if it is False - it is a Deposit address. | 
 **startIndex** | **int64** | The starting index of the response items, i.e. where the response should start listing the returned items. | 

### Return type

[**DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesR**](DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EncodeXAddress

> EncodeXAddressR EncodeXAddress(ctx, addressTag, blockchain, classicAddress, network).Context(context).Execute()

Encode X-Address



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
    addressTag := int32(3999472835) // int32 | Defines a specific Tag that is an additional XRP address feature. It helps identifying a transaction recipient beyond a wallet address.
    blockchain := "xrp" // string | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
    classicAddress := "rA9bXGJcXvZKaWofrRphdJsBWzhyCfH3z" // string | Represents the public address, which is a compressed and shortened form of a public key.
    network := "testnet" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.EncodeXAddress(context.Background(), addressTag, blockchain, classicAddress, network).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.EncodeXAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EncodeXAddress`: EncodeXAddressR
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.EncodeXAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addressTag** | **int32** | Defines a specific Tag that is an additional XRP address feature. It helps identifying a transaction recipient beyond a wallet address. | 
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**classicAddress** | **string** | Represents the public address, which is a compressed and shortened form of a public key. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEncodeXAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**EncodeXAddressR**](EncodeXAddressR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EstimateGasLimit

> EstimateGasLimitR EstimateGasLimit(ctx, blockchain, network).Context(context).EstimateGasLimitRB(estimateGasLimitRB).Execute()

Estimate Gas Limit



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
    blockchain := "ethereum" // string | Represents the specific blockchain protocol name, e.g. Ethereum.
    network := "ropsten" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    estimateGasLimitRB := *openapiclient.NewEstimateGasLimitRB(*openapiclient.NewEstimateGasLimitRBData(*openapiclient.NewEstimateGasLimitRBDataItem("0.002", "0xc065b539490f81b6c297c37b1925c3be2f190738", "0x6f61e3c2fbb8c8be698bd0907ba6c04b62800fe5"))) // EstimateGasLimitRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.EstimateGasLimit(context.Background(), blockchain, network).Context(context).EstimateGasLimitRB(estimateGasLimitRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.EstimateGasLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EstimateGasLimit`: EstimateGasLimitR
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.EstimateGasLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEstimateGasLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **estimateGasLimitRB** | [**EstimateGasLimitRB**](EstimateGasLimitRB.md) |  | 

### Return type

[**EstimateGasLimitR**](EstimateGasLimitR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EstimateTokenGasLimit

> EstimateTokenGasLimitR EstimateTokenGasLimit(ctx, blockchain, network).Context(context).EstimateTokenGasLimitRB(estimateTokenGasLimitRB).Execute()

Estimate Token Gas Limit



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
    blockchain := "ethereum" // string | Represents the specific blockchain protocol name, e.g. Ethereum.
    network := "ropsten" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)
    estimateTokenGasLimitRB := *openapiclient.NewEstimateTokenGasLimitRB(*openapiclient.NewEstimateTokenGasLimitRBData(*openapiclient.NewEstimateTokenGasLimitRBDataItem("0.12", "0x092de782a7e1e0a92991ad829a0a33aef3c7545e", "ERC-20", "0xc065b539490f81b6c297c37b1925c3be2f190738", "0x6f61e3c2fbb8c8be698bd0907ba6c04b62800fe5"))) // EstimateTokenGasLimitRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.EstimateTokenGasLimit(context.Background(), blockchain, network).Context(context).EstimateTokenGasLimitRB(estimateTokenGasLimitRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.EstimateTokenGasLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EstimateTokenGasLimit`: EstimateTokenGasLimitR
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.EstimateTokenGasLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEstimateTokenGasLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **estimateTokenGasLimitRB** | [**EstimateTokenGasLimitRB**](EstimateTokenGasLimitRB.md) |  | 

### Return type

[**EstimateTokenGasLimitR**](EstimateTokenGasLimitR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEIP1559FeeRecommendations

> GetEIP1559FeeRecommendationsR GetEIP1559FeeRecommendations(ctx, network, blockchain).Context(context).Execute()

Get EIP 1559 Fee Recommendations



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
    network := "ropsten" // string | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
    blockchain := "ethereum" // string | Represents the specific blockchain protocol name, e.g. Ethereum.
    context := "yourExampleString" // string | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.GetEIP1559FeeRecommendations(context.Background(), network, blockchain).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.GetEIP1559FeeRecommendations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEIP1559FeeRecommendations`: GetEIP1559FeeRecommendationsR
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.GetEIP1559FeeRecommendations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEIP1559FeeRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 

### Return type

[**GetEIP1559FeeRecommendationsR**](GetEIP1559FeeRecommendationsR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAddress

> ValidateAddressR ValidateAddress(ctx, blockchain, network).Context(context).ValidateAddressRB(validateAddressRB).Execute()

Validate Address



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
    validateAddressRB := *openapiclient.NewValidateAddressRB(*openapiclient.NewValidateAddressRBData(*openapiclient.NewValidateAddressRBDataItem("mho4jHBcrNCncKt38trJahXakuaBnS7LK5"))) // ValidateAddressRB |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.ValidateAddress(context.Background(), blockchain, network).Context(context).ValidateAddressRB(validateAddressRB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.ValidateAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateAddress`: ValidateAddressR
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.ValidateAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | 
 **validateAddressRB** | [**ValidateAddressRB**](ValidateAddressRB.md) |  | 

### Return type

[**ValidateAddressR**](ValidateAddressR.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


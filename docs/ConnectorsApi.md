# \ConnectorsApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ConnectJsGet**](ConnectorsApi.md#V1ConnectJsGet) | **Get** /v1/connect.js | Get embeddable connect javascript
[**V1ConnectMobileGet**](ConnectorsApi.md#V1ConnectMobileGet) | **Get** /v1/connect/mobile | Mobile connect page
[**V1ConnectorsConnectorConnectGet**](ConnectorsApi.md#V1ConnectorsConnectorConnectGet) | **Get** /v1/connectors/{connector}/connect | Obtain a token from 3rd party data source
[**V1ConnectorsConnectorConnectInstructionsGet**](ConnectorsApi.md#V1ConnectorsConnectorConnectInstructionsGet) | **Get** /v1/connectors/{connector}/connectInstructions | Connection Instructions
[**V1ConnectorsConnectorConnectParameterGet**](ConnectorsApi.md#V1ConnectorsConnectorConnectParameterGet) | **Get** /v1/connectors/{connector}/connectParameter | Connect Parameter
[**V1ConnectorsConnectorDisconnectGet**](ConnectorsApi.md#V1ConnectorsConnectorDisconnectGet) | **Get** /v1/connectors/{connector}/disconnect | Delete stored connection info
[**V1ConnectorsConnectorInfoGet**](ConnectorsApi.md#V1ConnectorsConnectorInfoGet) | **Get** /v1/connectors/{connector}/info | Get connector info for user
[**V1ConnectorsConnectorUpdateGet**](ConnectorsApi.md#V1ConnectorsConnectorUpdateGet) | **Get** /v1/connectors/{connector}/update | Sync with data source
[**V1ConnectorsListGet**](ConnectorsApi.md#V1ConnectorsListGet) | **Get** /v1/connectors/list | List of Connectors


# **V1ConnectJsGet**
> V1ConnectJsGet($accessToken)

Get embeddable connect javascript

Get embeddable connect javascript. Usage:    - Embedding in applications with popups for 3rd-party authentication windows.      Use `qmSetupInPopup` function after connecting `connect.js`.    - Embedding in applications with popups for 3rd-party authentication windows.      Requires a selector to block. It will be embedded in this block.      Use `qmSetupOnPage` function after connecting `connect.js`.    - Embedding in mobile applications without popups for 3rd-party authentication.      Use `qmSetupOnMobile` function after connecting `connect.js`.      if using the MoodiModo Clones, Use `qmSetupOnIonic` function after connecting `connect.js`. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

void (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [internalApiKey](../README.md#internalApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/x-javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ConnectMobileGet**
> V1ConnectMobileGet($accessToken)

Mobile connect page

This page is designed to be opened in a webview.  Instead of using popup authentication boxes, it uses redirection. You can include the user's access_token as a URL parameter like https://app.quantimo.do/api/v1/connect/mobile?access_token=123


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User OAuth access token | 

### Return type

void (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [internalApiKey](../README.md#internalApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ConnectorsConnectorConnectGet**
> V1ConnectorsConnectorConnectGet($connector, $accessToken)

Obtain a token from 3rd party data source

Attempt to obtain a token from the data provider, store it in the database. With this, the connector to continue to obtain new user data until the token is revoked.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connector** | **string**| Lowercase system name of the source application or device. Get a list of available connectors from the /connectors/list endpoint. | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

void (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ConnectorsConnectorConnectInstructionsGet**
> V1ConnectorsConnectorConnectInstructionsGet($connector, $parameters, $url, $usePopup, $accessToken)

Connection Instructions

Returns instructions that describe what parameters and endpoint to use to connect to the given data provider.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connector** | **string**| Lowercase system name of the source application or device. Get a list of available connectors from the /connectors/list endpoint. | 
 **parameters** | **string**| JSON Array of Parameters for the request to enable connector. | 
 **url** | **string**| URL which should be used to enable the connector. | 
 **usePopup** | **bool**| Should use popup when enabling connector | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

void (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ConnectorsConnectorConnectParameterGet**
> ConnectorInstruction V1ConnectorsConnectorConnectParameterGet($connector, $displayName, $key, $placeholder, $type_, $usePopup, $accessToken, $defaultValue)

Connect Parameter

Returns instructions that describe what parameters and endpoint to use to connect to the given data provider.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connector** | **string**| Lowercase system name of the source application or device. Get a list of available connectors from the /connectors/list endpoint. | 
 **displayName** | **string**| Name of the parameter that is user visible in the form | 
 **key** | **string**| Name of the property that the user has to enter such as username or password Connector (used in HTTP request) | 
 **placeholder** | **string**| Placeholder hint value for the parameter input tag. | 
 **type_** | **string**| Type of input field such as those found here http://www.w3schools.com/tags/tag_input.asp | 
 **usePopup** | **bool**| Should use popup when enabling connector | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **defaultValue** | **string**| Default parameter value | [optional] 

### Return type

[**ConnectorInstruction**](ConnectorInstruction.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ConnectorsConnectorDisconnectGet**
> V1ConnectorsConnectorDisconnectGet($connector)

Delete stored connection info

The disconnect method deletes any stored tokens or connection information from the connectors database.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connector** | **string**| Lowercase system name of the source application or device. Get a list of available connectors from the /connectors/list endpoint. | 

### Return type

void (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ConnectorsConnectorInfoGet**
> ConnectorInfo V1ConnectorsConnectorInfoGet($connector, $accessToken)

Get connector info for user

Returns information about the connector such as the connector id, whether or not is connected for this user (i.e. we have a token or credentials), and its update history for the user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connector** | **string**| Lowercase system name of the source application or device. Get a list of available connectors from the /connectors/list endpoint. | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

[**ConnectorInfo**](ConnectorInfo.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ConnectorsConnectorUpdateGet**
> V1ConnectorsConnectorUpdateGet($connector, $accessToken)

Sync with data source

The update method tells the QM Connector Framework to check with the data provider (such as Fitbit or MyFitnessPal) and retrieve any new measurements available.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connector** | **string**| Lowercase system name of the source application or device | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

void (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ConnectorsListGet**
> []Connector V1ConnectorsListGet()

List of Connectors

A connector pulls data from other data providers using their API or a screenscraper. Returns a list of all available connectors and information about them such as their id, name, whether the user has provided access, logo url, connection instructions, and the update history.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]Connector**](Connector.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


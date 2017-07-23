# \ConnectorsApi

All URIs are relative to *https://app.quantimo.do/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ConnectMobileGet**](ConnectorsApi.md#V1ConnectMobileGet) | **Get** /v1/connect/mobile | Mobile connect page
[**V1ConnectorsConnectorNameConnectGet**](ConnectorsApi.md#V1ConnectorsConnectorNameConnectGet) | **Get** /v1/connectors/{connectorName}/connect | Obtain a token from 3rd party data source
[**V1ConnectorsConnectorNameConnectInstructionsGet**](ConnectorsApi.md#V1ConnectorsConnectorNameConnectInstructionsGet) | **Get** /v1/connectors/{connectorName}/connectInstructions | Connection Instructions
[**V1ConnectorsConnectorNameConnectParameterGet**](ConnectorsApi.md#V1ConnectorsConnectorNameConnectParameterGet) | **Get** /v1/connectors/{connectorName}/connectParameter | Connect Parameter
[**V1ConnectorsConnectorNameDisconnectGet**](ConnectorsApi.md#V1ConnectorsConnectorNameDisconnectGet) | **Get** /v1/connectors/{connectorName}/disconnect | Delete stored connection info
[**V1ConnectorsConnectorNameInfoGet**](ConnectorsApi.md#V1ConnectorsConnectorNameInfoGet) | **Get** /v1/connectors/{connectorName}/info | Get connector info for user
[**V1ConnectorsConnectorNameUpdateGet**](ConnectorsApi.md#V1ConnectorsConnectorNameUpdateGet) | **Get** /v1/connectors/{connectorName}/update | Sync with data source
[**V1ConnectorsListGet**](ConnectorsApi.md#V1ConnectorsListGet) | **Get** /v1/connectors/list | List of Connectors
[**V1IntegrationJsGet**](ConnectorsApi.md#V1IntegrationJsGet) | **Get** /v1/integration.js | Get embeddable connect javascript


# **V1ConnectMobileGet**
> V1ConnectMobileGet($accessToken, $userId)

Mobile connect page

This page is designed to be opened in a webview.  Instead of using popup authentication boxes, it uses redirection. You can include the user's access_token as a URL parameter like https://app.quantimo.do/api/v1/connect/mobile?access_token=123


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User OAuth access token | 
 **userId** | **int32**| User&#39;s id | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ConnectorsConnectorNameConnectGet**
> V1ConnectorsConnectorNameConnectGet($connectorName, $accessToken, $userId)

Obtain a token from 3rd party data source

Attempt to obtain a token from the data provider, store it in the database. With this, the connector to continue to obtain new user data until the token is revoked.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectorName** | **string**| Lowercase system name of the source application or device. Get a list of available connectors from the /v1/connectors/list endpoint. | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **userId** | **int32**| User&#39;s id | [optional] 

### Return type

void (empty response body)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ConnectorsConnectorNameConnectInstructionsGet**
> V1ConnectorsConnectorNameConnectInstructionsGet($connectorName, $parameters, $url, $usePopup, $accessToken, $userId)

Connection Instructions

Returns instructions that describe what parameters and endpoint to use to connect to the given data provider.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectorName** | **string**| Lowercase system name of the source application or device. Get a list of available connectors from the /v1/connectors/list endpoint. | 
 **parameters** | **string**| JSON Array of Parameters for the request to enable connector. | 
 **url** | **string**| URL which should be used to enable the connector. | 
 **usePopup** | **bool**| Should use popup when enabling connector | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **userId** | **int32**| User&#39;s id | [optional] 

### Return type

void (empty response body)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ConnectorsConnectorNameConnectParameterGet**
> ConnectorInstruction V1ConnectorsConnectorNameConnectParameterGet($connectorName, $displayName, $key, $placeholder, $type_, $usePopup, $accessToken, $userId, $defaultValue)

Connect Parameter

Returns instructions that describe what parameters and endpoint to use to connect to the given data provider.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectorName** | **string**| Lowercase system name of the source application or device. Get a list of available connectors from the /v1/connectors/list endpoint. | 
 **displayName** | **string**| Name of the parameter that is user visible in the form | 
 **key** | **string**| Name of the property that the user has to enter such as username or password Connector (used in HTTP request) | 
 **placeholder** | **string**| Placeholder hint value for the parameter input tag. | 
 **type_** | **string**| Type of input field such as those found here http://www.w3schools.com/tags/tag_input.asp | 
 **usePopup** | **bool**| Should use popup when enabling connector | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **userId** | **int32**| User&#39;s id | [optional] 
 **defaultValue** | **string**| Default parameter value | [optional] 

### Return type

[**ConnectorInstruction**](ConnectorInstruction.md)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ConnectorsConnectorNameDisconnectGet**
> V1ConnectorsConnectorNameDisconnectGet($connectorName)

Delete stored connection info

The disconnect method deletes any stored tokens or connection information from the connectors database.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectorName** | **string**| Lowercase system name of the source application or device. Get a list of available connectors from the /v1/connectors/list endpoint. | 

### Return type

void (empty response body)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ConnectorsConnectorNameInfoGet**
> ConnectorInfo V1ConnectorsConnectorNameInfoGet($connectorName, $accessToken, $userId)

Get connector info for user

Returns information about the connector such as the connector id, whether or not is connected for this user (i.e. we have a token or credentials), and its update history for the user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectorName** | **string**| Lowercase system name of the source application or device. Get a list of available connectors from the /v1/connectors/list endpoint. | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **userId** | **int32**| User&#39;s id | [optional] 

### Return type

[**ConnectorInfo**](ConnectorInfo.md)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ConnectorsConnectorNameUpdateGet**
> V1ConnectorsConnectorNameUpdateGet($connectorName, $accessToken, $userId)

Sync with data source

The update method tells the QM Connector Framework to check with the data provider (such as Fitbit or MyFitnessPal) and retrieve any new measurements available.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectorName** | **string**| Lowercase system name of the source application or device | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **userId** | **int32**| User&#39;s id | [optional] 

### Return type

void (empty response body)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

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

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1IntegrationJsGet**
> V1IntegrationJsGet($accessToken)

Get embeddable connect javascript

Get embeddable connect javascript. Usage:   - Embedding in applications with popups for 3rd-party authentication windows.     Use `qmSetupInPopup` function after connecting `connect.js`.   - Embedding in applications with popups for 3rd-party authentication windows.     Requires a selector to block. It will be embedded in this block.     Use `qmSetupOnPage` function after connecting `connect.js`.   - Embedding in mobile applications without popups for 3rd-party authentication.     Use `qmSetupOnMobile` function after connecting `connect.js`.     If using in a Cordova application call  `qmSetupOnIonic` function after connecting `connect.js`.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/x-javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


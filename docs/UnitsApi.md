# \UnitsApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1UnitCategoriesGet**](UnitsApi.md#V1UnitCategoriesGet) | **Get** /v1/unitCategories | Get unit categories
[**V1UnitsGet**](UnitsApi.md#V1UnitsGet) | **Get** /v1/units | Get all available units
[**V1UnitsVariableGet**](UnitsApi.md#V1UnitsVariableGet) | **Get** /v1/unitsVariable | Units for Variable


# **V1UnitCategoriesGet**
> UnitCategory V1UnitCategoriesGet()

Get unit categories

Get a list of the categories of measurement units such as 'Distance', 'Duration', 'Energy', 'Frequency', 'Miscellany', 'Pressure', 'Proportion', 'Rating', 'Temperature', 'Volume', and 'Weight'.


### Parameters
This endpoint does not need any parameter.

### Return type

[**UnitCategory**](UnitCategory.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UnitsGet**
> []Unit V1UnitsGet($accessToken, $id, $unitName, $abbreviatedUnitName, $categoryName)

Get all available units

Get all available units


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **id** | **int32**| Unit id | [optional] 
 **unitName** | **string**| Unit name | [optional] 
 **abbreviatedUnitName** | **string**| Restrict the results to a specific unit by providing the unit abbreviation. | [optional] 
 **categoryName** | **string**| Restrict the results to a specific unit category by providing the unit category name. | [optional] 

### Return type

[**[]Unit**](Unit.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UnitsVariableGet**
> []Unit V1UnitsVariableGet($accessToken, $unitName, $abbreviatedUnitName, $categoryName, $variable)

Units for Variable

Get a list of all possible units to use for a given variable


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **unitName** | **string**| Name of Unit you want to retrieve | [optional] 
 **abbreviatedUnitName** | **string**| Abbreviated Unit Name of the unit you want | [optional] 
 **categoryName** | **string**| Name of the category you want units for | [optional] 
 **variable** | **string**| Name of the variable you want units for | [optional] 

### Return type

[**[]Unit**](Unit.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


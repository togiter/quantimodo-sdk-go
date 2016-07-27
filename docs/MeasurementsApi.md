# \MeasurementsApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MeasurementSourcesGet**](MeasurementsApi.md#V1MeasurementSourcesGet) | **Get** /v1/measurementSources | Get measurement sources
[**V1MeasurementSourcesPost**](MeasurementsApi.md#V1MeasurementSourcesPost) | **Post** /v1/measurementSources | Add a data source
[**V1MeasurementsDailyGet**](MeasurementsApi.md#V1MeasurementsDailyGet) | **Get** /v1/measurements/daily | Get daily measurements for this user
[**V1MeasurementsDeletePost**](MeasurementsApi.md#V1MeasurementsDeletePost) | **Post** /v1/measurements/delete | Delete a measurement
[**V1MeasurementsGet**](MeasurementsApi.md#V1MeasurementsGet) | **Get** /v1/measurements | Get measurements for this user
[**V1MeasurementsPost**](MeasurementsApi.md#V1MeasurementsPost) | **Post** /v1/measurements | Post a new set or update existing measurements to the database
[**V1MeasurementsRangeGet**](MeasurementsApi.md#V1MeasurementsRangeGet) | **Get** /v1/measurementsRange | Get measurements range for this user
[**V2MeasurementsCsvGet**](MeasurementsApi.md#V2MeasurementsCsvGet) | **Get** /v2/measurements/csv | Get Measurements CSV
[**V2MeasurementsIdDelete**](MeasurementsApi.md#V2MeasurementsIdDelete) | **Delete** /v2/measurements/{id} | Delete Measurement
[**V2MeasurementsIdGet**](MeasurementsApi.md#V2MeasurementsIdGet) | **Get** /v2/measurements/{id} | Get Measurement
[**V2MeasurementsIdPut**](MeasurementsApi.md#V2MeasurementsIdPut) | **Put** /v2/measurements/{id} | Update Measurement
[**V2MeasurementsRequestCsvPost**](MeasurementsApi.md#V2MeasurementsRequestCsvPost) | **Post** /v2/measurements/request_csv | Post Request for Measurements CSV
[**V2MeasurementsRequestPdfPost**](MeasurementsApi.md#V2MeasurementsRequestPdfPost) | **Post** /v2/measurements/request_pdf | Post Request for Measurements PDF
[**V2MeasurementsRequestXlsPost**](MeasurementsApi.md#V2MeasurementsRequestXlsPost) | **Post** /v2/measurements/request_xls | Post Request for Measurements XLS


# **V1MeasurementSourcesGet**
> MeasurementSource V1MeasurementSourcesGet()

Get measurement sources

Returns a list of all the apps from which measurement data is obtained.


### Parameters
This endpoint does not need any parameter.

### Return type

[**MeasurementSource**](MeasurementSource.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1MeasurementSourcesPost**
> V1MeasurementSourcesPost($body, $accessToken)

Add a data source

Add a life-tracking app or device to the QuantiModo list of data sources.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MeasurementSource**](MeasurementSource.md)| An array of names of data sources you want to add. | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

void (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1MeasurementsDailyGet**
> Measurement V1MeasurementsDailyGet($variableName, $accessToken, $abbreviatedUnitName, $startTime, $endTime, $groupingWidth, $groupingTimezone, $limit, $offset, $sort)

Get daily measurements for this user

Measurements are any value that can be recorded like daily steps, a mood rating, or apples eaten. <br>Supported filter parameters:<br><ul><li><b>value</b> - Value of measurement</li><li><b>lastUpdated</b> - The time that this measurement was created or last updated in the UTC format \"YYYY-MM-DDThh:mm:ss\"</li></ul><br>


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableName** | **string**| Name of the variable you want measurements for | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **abbreviatedUnitName** | **string**| The unit your want the measurements in | [optional] 
 **startTime** | **string**| The lower limit of measurements returned (Iso8601) | [optional] 
 **endTime** | **string**| The upper limit of measurements returned (Iso8601) | [optional] 
 **groupingWidth** | **int32**| The time (in seconds) over which measurements are grouped together | [optional] 
 **groupingTimezone** | **string**| The time (in seconds) over which measurements are grouped together | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. | [optional] 
 **offset** | **int32**| Now suppose you wanted to show results 11-20. You&#39;d set the offset to 10 and the limit to 10. | [optional] 
 **sort** | **int32**| Sort by given field. If the field is prefixed with &#x60;-, it will sort in descending order. | [optional] 

### Return type

[**Measurement**](Measurement.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1MeasurementsDeletePost**
> CommonResponse V1MeasurementsDeletePost($body)

Delete a measurement

Delete a previously submitted measurement


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MeasurementDelete**](MeasurementDelete.md)| The startTime and variableId of the measurement to be deleted. | 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1MeasurementsGet**
> Measurement V1MeasurementsGet($accessToken, $id, $variableName, $variableCategoryName, $source, $value, $lastUpdated, $unit, $startTime, $createdAt, $updatedAt, $endTime, $groupingWidth, $groupingTimezone, $limit, $offset, $sort)

Get measurements for this user

Measurements are any value that can be recorded like daily steps, a mood rating, or apples eaten. <br>Supported filter parameters:<br><ul><li><b>value</b> - Value of measurement</li><li><b>lastUpdated</b> - The time that this measurement was created or last updated in the UTC format \"YYYY-MM-DDThh:mm:ss\"</li></ul><br>


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **id** | **int32**| Measurement id | [optional] 
 **variableName** | **string**| Name of the variable you want measurements for | [optional] 
 **variableCategoryName** | **string**| Name of the variable category you want measurements for | [optional] 
 **source** | **string**| Name of the source you want measurements for (supports exact name match only) | [optional] 
 **value** | **string**| Value of measurement | [optional] 
 **lastUpdated** | **string**| The time that this measurement was created or last updated in the format \&quot;YYYY-MM-DDThh:mm:ss\&quot; | [optional] 
 **unit** | **string**| The unit you want the measurements returned in | [optional] 
 **startTime** | **string**| The lower limit of measurements returned (Epoch) | [optional] 
 **createdAt** | **string**| The time the measurement record was first created in the format YYYY-MM-DDThh:mm:ss. Time zone should be UTC and not local. | [optional] 
 **updatedAt** | **string**| The time the measurement record was last changed in the format YYYY-MM-DDThh:mm:ss. Time zone should be UTC and not local. | [optional] 
 **endTime** | **string**| The upper limit of measurements returned (Epoch) | [optional] 
 **groupingWidth** | **int32**| The time (in seconds) over which measurements are grouped together | [optional] 
 **groupingTimezone** | **string**| The time (in seconds) over which measurements are grouped together | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. | [optional] 
 **offset** | **int32**| Now suppose you wanted to show results 11-20. You&#39;d set the offset to 10 and the limit to 10. | [optional] 
 **sort** | **int32**| Sort by given field. If the field is prefixed with &#x60;-, it will sort in descending order. | [optional] 

### Return type

[**Measurement**](Measurement.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1MeasurementsPost**
> V1MeasurementsPost($body, $accessToken)

Post a new set or update existing measurements to the database

You can submit or update multiple measurements in a \"measurements\" sub-array.  If the variable these measurements correspond to does not already exist in the database, it will be automatically added.  The request body should look something like [{\"measurements\":[{\"startTime\":1439389320,\"value\":\"3\"}, {\"startTime\":1439389319,\"value\":\"2\"}],\"name\":\"Acne (out of 5)\",\"source\":\"QuantiModo\",\"category\":\"Symptoms\",\"combinationOperation\":\"MEAN\",\"unit\":\"/5\"}]


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MeasurementSet**](MeasurementSet.md)| An array of measurements you want to insert. | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

void (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1MeasurementsRangeGet**
> MeasurementRange V1MeasurementsRangeGet($sources, $user)

Get measurements range for this user

Get Unix time-stamp (epoch time) of the user's first and last measurements taken.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sources** | **string**| Enter source name to limit to specific source (varchar) | [optional] 
 **user** | **int32**| If not specified, uses currently logged in user (bigint) | [optional] 

### Return type

[**MeasurementRange**](MeasurementRange.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2MeasurementsCsvGet**
> *os.File V2MeasurementsCsvGet($accessToken)

Get Measurements CSV

Download a CSV containing all user measurements


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

[***os.File**](*os.File.md)

### Authorization

[quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2MeasurementsIdDelete**
> InlineResponse20012 V2MeasurementsIdDelete($id, $accessToken)

Delete Measurement

Delete Measurement


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| id of Measurement | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2MeasurementsIdGet**
> InlineResponse20011 V2MeasurementsIdGet($id, $accessToken)

Get Measurement

Get Measurement


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| id of Measurement | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2MeasurementsIdPut**
> InlineResponse20012 V2MeasurementsIdPut($id, $accessToken, $body)

Update Measurement

Update Measurement


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| id of Measurement | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **body** | [**Measurement**](Measurement.md)| Measurement that should be updated | [optional] 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2MeasurementsRequestCsvPost**
> int32 V2MeasurementsRequestCsvPost($accessToken)

Post Request for Measurements CSV

Use this endpoint to schedule a CSV export containing all user measurements to be emailed to the user within 24 hours.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

**int32**

### Authorization

[quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2MeasurementsRequestPdfPost**
> int32 V2MeasurementsRequestPdfPost($accessToken)

Post Request for Measurements PDF

Use this endpoint to schedule a PDF export containing all user measurements to be emailed to the user within 24 hours.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

**int32**

### Authorization

[quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2MeasurementsRequestXlsPost**
> int32 V2MeasurementsRequestXlsPost($accessToken)

Post Request for Measurements XLS

Use this endpoint to schedule a XLS export containing all user measurements to be emailed to the user within 24 hours.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

**int32**

### Authorization

[quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


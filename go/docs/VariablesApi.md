# \VariablesApi

All URIs are relative to *https://app.quantimo.do/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PublicVariablesGet**](VariablesApi.md#V1PublicVariablesGet) | **Get** /v1/public/variables | Get public variables
[**V1PublicVariablesSearchSearchGet**](VariablesApi.md#V1PublicVariablesSearchSearchGet) | **Get** /v1/public/variables/search/{search} | Get top 5 PUBLIC variables with the most correlations
[**V1UserVariablesDeletePost**](VariablesApi.md#V1UserVariablesDeletePost) | **Post** /v1/userVariables/delete | Delete All Measurements For Variable
[**V1UserVariablesPost**](VariablesApi.md#V1UserVariablesPost) | **Post** /v1/userVariables | Update User Settings for a Variable
[**V1UserVariablesResetPost**](VariablesApi.md#V1UserVariablesResetPost) | **Post** /v1/userVariables/reset | Reset user settings for a variable to defaults
[**V1VariableCategoriesGet**](VariablesApi.md#V1VariableCategoriesGet) | **Get** /v1/variableCategories | Variable categories
[**V1VariablesGet**](VariablesApi.md#V1VariablesGet) | **Get** /v1/variables | Get variables with user&#39;s settings
[**V1VariablesPost**](VariablesApi.md#V1VariablesPost) | **Post** /v1/variables | Create Variables
[**V1VariablesSearchSearchGet**](VariablesApi.md#V1VariablesSearchSearchGet) | **Get** /v1/variables/search/{search} | Get variables by search query
[**V1VariablesVariableNameGet**](VariablesApi.md#V1VariablesVariableNameGet) | **Get** /v1/variables/{variableName} | Get info about a variable


# **V1PublicVariablesGet**
> Variable V1PublicVariablesGet($accessToken, $userId, $id, $category, $name, $updatedAt, $source, $latestMeasurementTime, $numberOfRawMeasurements, $lastSource, $limit, $offset, $sort)

Get public variables

This endpoint retrieves an array of all public variables. Public variables are things like foods, medications, symptoms, conditions, and anything not unique to a particular user. For instance, a telephone number or name would not be a public variable.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **userId** | **int32**| User&#39;s id | [optional] 
 **id** | **int32**| Common variable id | [optional] 
 **category** | **string**| Filter data by category | [optional] 
 **name** | **string**| Original name of the variable (supports exact name match only) | [optional] 
 **updatedAt** | **string**| Filter by the last time any of the properties of the variable were changed. Uses UTC format \&quot;YYYY-MM-DDThh:mm:ss\&quot; | [optional] 
 **source** | **string**| The name of the data source that created the variable (supports exact name match only). So if you have a client application and you only want variables that were last updated by your app, you can include the name of your app here | [optional] 
 **latestMeasurementTime** | **string**| Filter variables based on the last time a measurement for them was created or updated in the UTC format \&quot;YYYY-MM-DDThh:mm:ss\&quot; | [optional] 
 **numberOfRawMeasurements** | **string**| Filter variables by the total number of measurements that they have. This could be used of you want to filter or sort by popularity. | [optional] 
 **lastSource** | **string**| Limit variables to those which measurements were last submitted by a specific source. So if you have a client application and you only want variables that were last updated by your app, you can include the name of your app here. (supports exact name match only) | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. | [optional] 
 **offset** | **int32**| Since the maximum limit is 200 records, to get more than that you&#39;ll have to make multiple API calls and page through the results. To retrieve all the data, you can iterate through data by using the &#x60;limit&#x60; and &#x60;offset&#x60; query parameters.  For example, if you want to retrieve data from 61-80 then you can use a query with the following parameters, &#x60;imit&#x3D;20&amp;offset&#x3D;60&#x60;. | [optional] 
 **sort** | **int32**| Sort by given field. If the field is prefixed with &#x60;-, it will sort in descending order. | [optional] 

### Return type

[**Variable**](Variable.md)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1PublicVariablesSearchSearchGet**
> Variable V1PublicVariablesSearchSearchGet($search, $accessToken, $userId, $variableCategoryName, $source, $effectOrCause, $publicEffectOrCause, $limit, $offset, $sort)

Get top 5 PUBLIC variables with the most correlations

Get top 5 PUBLIC variables with the most correlations containing the entered search characters. For example, search for 'mood' as an effect. Since 'Overall Mood' has a lot of correlations with other variables, it should be in the autocomplete list.Supported filter parameters:<ul><li><b>category</b> - Category of Variable</li></ul>


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string**| Search query can be some fraction of a variable name. | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **userId** | **int32**| User&#39;s id | [optional] 
 **variableCategoryName** | **string**| Filter variables by category name. The variable categories include Activity, Causes of Illness, Cognitive Performance, Conditions, Environment, Foods, Location, Miscellaneous, Mood, Nutrition, Physical Activity, Physique, Sleep, Social Interactions, Symptoms, Treatments, Vital Signs, and Work. | [optional] 
 **source** | **string**| Specify a data source name to only return variables from a specific data source. | [optional] 
 **effectOrCause** | **string**| Indicate if you only want variables that have user correlations. Possible values are effect and cause. | [optional] 
 **publicEffectOrCause** | **string**| Indicate if you only want variables that have aggregated correlations.  Possible values are effect and cause. | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. | [optional] 
 **offset** | **int32**| Since the maximum limit is 200 records, to get more than that you&#39;ll have to make multiple API calls and page through the results. To retrieve all the data, you can iterate through data by using the &#x60;limit&#x60; and &#x60;offset&#x60; query parameters.  For example, if you want to retrieve data from 61-80 then you can use a query with the following parameters, &#x60;imit&#x3D;20&amp;offset&#x3D;60&#x60;. | [optional] 
 **sort** | **int32**| Sort by given field. If the field is prefixed with &#x60;-, it will sort in descending order. | [optional] 

### Return type

[**Variable**](Variable.md)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UserVariablesDeletePost**
> V1UserVariablesDeletePost($variableId)

Delete All Measurements For Variable

Users can delete all of their measurements for a variable


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableId** | [**UserVariableDelete**](UserVariableDelete.md)| Id of the variable whose measurements should be deleted | 

### Return type

void (empty response body)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UserVariablesPost**
> V1UserVariablesPost($userVariables)

Update User Settings for a Variable

Users can change the parameters used in analysis of that variable such as the expected duration of action for a variable to have an effect, the estimated delay before the onset of action. In order to filter out erroneous data, they are able to set the maximum and minimum reasonable daily values for a variable.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userVariables** | [**UserVariables**](UserVariables.md)| Variable user settings data | 

### Return type

void (empty response body)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UserVariablesResetPost**
> V1UserVariablesResetPost($variableId)

Reset user settings for a variable to defaults

Reset user settings for a variable to defaults


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableId** | [**UserVariableDelete**](UserVariableDelete.md)| Id of the variable that should be reset | 

### Return type

void (empty response body)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1VariableCategoriesGet**
> []VariableCategory V1VariableCategoriesGet()

Variable categories

The variable categories include Activity, Causes of Illness, Cognitive Performance, Conditions, Environment, Foods, Location, Miscellaneous, Mood, Nutrition, Physical Activity, Physique, Sleep, Social Interactions, Symptoms, Treatments, Vital Signs, and Work.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]VariableCategory**](VariableCategory.md)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1VariablesGet**
> Variable V1VariablesGet($accessToken, $userId, $id, $category, $name, $updatedAt, $source, $latestMeasurementTime, $numberOfRawMeasurements, $lastSource, $limit, $offset, $sort)

Get variables with user's settings

Get variables for which the user has measurements. If the user has specified variable settings, these are provided instead of the common variable defaults.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **userId** | **int32**| User&#39;s id | [optional] 
 **id** | **int32**| Common variable id | [optional] 
 **category** | **string**| Filter data by category | [optional] 
 **name** | **string**| Original name of the variable (supports exact name match only) | [optional] 
 **updatedAt** | **string**| Filter by the last time any of the properties of the variable were changed. Uses UTC format \&quot;YYYY-MM-DDThh:mm:ss\&quot; | [optional] 
 **source** | **string**| The name of the data source that created the variable (supports exact name match only). So if you have a client application and you only want variables that were last updated by your app, you can include the name of your app here | [optional] 
 **latestMeasurementTime** | **string**| Filter variables based on the last time a measurement for them was created or updated in the UTC format \&quot;YYYY-MM-DDThh:mm:ss\&quot; | [optional] 
 **numberOfRawMeasurements** | **string**| Filter variables by the total number of measurements that they have. This could be used of you want to filter or sort by popularity. | [optional] 
 **lastSource** | **string**| Limit variables to those which measurements were last submitted by a specific source. So if you have a client application and you only want variables that were last updated by your app, you can include the name of your app here. (supports exact name match only) | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. | [optional] 
 **offset** | **int32**| Since the maximum limit is 200 records, to get more than that you&#39;ll have to make multiple API calls and page through the results. To retrieve all the data, you can iterate through data by using the &#x60;limit&#x60; and &#x60;offset&#x60; query parameters.  For example, if you want to retrieve data from 61-80 then you can use a query with the following parameters, &#x60;imit&#x3D;20&amp;offset&#x3D;60&#x60;. | [optional] 
 **sort** | **int32**| Sort by given field. If the field is prefixed with &#x60;-, it will sort in descending order. | [optional] 

### Return type

[**Variable**](Variable.md)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1VariablesPost**
> V1VariablesPost($body, $accessToken, $userId)

Create Variables

Allows the client to create a new variable in the `variables` table.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VariablesNew**](VariablesNew.md)| Original name for the variable. | 
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

# **V1VariablesSearchSearchGet**
> []Variable V1VariablesSearchSearchGet($search, $accessToken, $userId, $variableCategoryName, $includePublic, $manualTracking, $source, $effectOrCause, $publicEffectOrCause, $limit, $offset)

Get variables by search query

Get variables containing the search characters for which the currently logged in user has measurements. Used to provide auto-complete function in variable search boxes.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string**| Search query which may be an entire variable name or a fragment of one. | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **userId** | **int32**| User&#39;s id | [optional] 
 **variableCategoryName** | **string**| Filter variables by category name. The variable categories include Activity, Causes of Illness, Cognitive Performance, Conditions, Environment, Foods, Location, Miscellaneous, Mood, Nutrition, Physical Activity, Physique, Sleep, Social Interactions, Symptoms, Treatments, Vital Signs, and Work. | [optional] 
 **includePublic** | **bool**| Set to true if you would like to include public variables when no user variables are found. | [optional] 
 **manualTracking** | **bool**| Set to true if you would like to exlude variables like apps and website names. | [optional] 
 **source** | **string**| Specify a data source name to only return variables from a specific data source. | [optional] 
 **effectOrCause** | **string**| Indicate if you only want variables that have user correlations. Possible values are effect and cause. | [optional] 
 **publicEffectOrCause** | **string**| Indicate if you only want variables that have aggregated correlations.  Possible values are effect and cause. | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. | [optional] 
 **offset** | **int32**| Since the maximum limit is 200 records, to get more than that you&#39;ll have to make multiple API calls and page through the results. To retrieve all the data, you can iterate through data by using the &#x60;limit&#x60; and &#x60;offset&#x60; query parameters.  For example, if you want to retrieve data from 61-80 then you can use a query with the following parameters, &#x60;imit&#x3D;20&amp;offset&#x3D;60&#x60;. | [optional] 

### Return type

[**[]Variable**](Variable.md)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1VariablesVariableNameGet**
> Variable V1VariablesVariableNameGet($variableName, $accessToken, $userId)

Get info about a variable

Get all of the settings and information about a variable by its name. If the logged in user has modified the settings for the variable, these will be provided instead of the default settings for that variable.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableName** | **string**| Variable name | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **userId** | **int32**| User&#39;s id | [optional] 

### Return type

[**Variable**](Variable.md)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


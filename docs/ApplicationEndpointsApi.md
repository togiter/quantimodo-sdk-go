# \ApplicationEndpointsApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2ApplicationConnectionsGet**](ApplicationEndpointsApi.md#V2ApplicationConnectionsGet) | **Get** /v2/application/connections | Get all Connections
[**V2ApplicationCredentialsGet**](ApplicationEndpointsApi.md#V2ApplicationCredentialsGet) | **Get** /v2/application/credentials | Get all Credentials
[**V2ApplicationMeasurementsGet**](ApplicationEndpointsApi.md#V2ApplicationMeasurementsGet) | **Get** /v2/application/measurements | Get measurements for all users using your application
[**V2ApplicationTrackingRemindersGet**](ApplicationEndpointsApi.md#V2ApplicationTrackingRemindersGet) | **Get** /v2/application/trackingReminders | Get tracking reminders
[**V2ApplicationUpdatesGet**](ApplicationEndpointsApi.md#V2ApplicationUpdatesGet) | **Get** /v2/application/updates | Get all Updates
[**V2ApplicationUserVariableRelationshipsGet**](ApplicationEndpointsApi.md#V2ApplicationUserVariableRelationshipsGet) | **Get** /v2/application/userVariableRelationships | Get all UserVariableRelationships
[**V2ApplicationUserVariablesGet**](ApplicationEndpointsApi.md#V2ApplicationUserVariablesGet) | **Get** /v2/application/userVariables | Get all UserVariables
[**V2ApplicationVariableUserSourcesGet**](ApplicationEndpointsApi.md#V2ApplicationVariableUserSourcesGet) | **Get** /v2/application/variableUserSources | Get all VariableUserSources
[**V2ApplicationVotesGet**](ApplicationEndpointsApi.md#V2ApplicationVotesGet) | **Get** /v2/application/votes | Get all Votes


# **V2ApplicationConnectionsGet**
> InlineResponse2003 V2ApplicationConnectionsGet($accessToken, $connectorId, $connectStatus, $connectError, $updateRequestedAt, $updateStatus, $updateError, $lastSuccessfulUpdatedAt, $createdAt, $updatedAt, $limit, $offset, $sort)

Get all Connections

Get all Connections


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| Application&#39;s OAuth2 access token | [optional] 
 **connectorId** | **int32**| The id for the connector data source for which the connection is connected | [optional] 
 **connectStatus** | **string**| Indicates whether a connector is currently connected to a service for a user. | [optional] 
 **connectError** | **string**| Error message if there is a problem with authorizing this connection. | [optional] 
 **updateRequestedAt** | **string**| Time at which an update was requested by a user. | [optional] 
 **updateStatus** | **string**| Indicates whether a connector is currently updated. | [optional] 
 **updateError** | **string**| Indicates if there was an error during the update. | [optional] 
 **lastSuccessfulUpdatedAt** | **string**| The time at which the connector was last successfully updated. | [optional] 
 **createdAt** | **string**| When the record was first created. Use ISO 8601 datetime format  | [optional] 
 **updatedAt** | **string**| When the record was last updated. Use ISO 8601 datetime format  | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records. | [optional] 
 **offset** | **int32**| OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned. | [optional] 
 **sort** | **string**| Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order. | [optional] 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[oauth2](../README.md#oauth2), [internalApiKey](../README.md#internalApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2ApplicationCredentialsGet**
> InlineResponse2004 V2ApplicationCredentialsGet($accessToken, $connectorId, $attrKey, $attrValue, $createdAt, $updatedAt, $limit, $offset, $sort)

Get all Credentials

Get all Credentials


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| Application&#39;s OAuth2 access token | [optional] 
 **connectorId** | **int32**| The id for the connector data source from which the credential was obtained | [optional] 
 **attrKey** | **string**| Attribute name such as token, userid, username, or password | [optional] 
 **attrValue** | **string**| Encrypted value for the attribute specified | [optional] 
 **createdAt** | **string**| When the record was first created. Use ISO 8601 datetime format  | [optional] 
 **updatedAt** | **string**| When the record was last updated. Use ISO 8601 datetime format  | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records. | [optional] 
 **offset** | **int32**| OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned. | [optional] 
 **sort** | **string**| Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order. | [optional] 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[oauth2](../README.md#oauth2), [internalApiKey](../README.md#internalApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2ApplicationMeasurementsGet**
> InlineResponse2005 V2ApplicationMeasurementsGet($accessToken, $clientId, $connectorId, $variableId, $sourceId, $startTime, $value, $unitId, $originalValue, $originalUnitId, $duration, $note, $latitude, $longitude, $location, $createdAt, $updatedAt, $error_, $limit, $offset, $sort)

Get measurements for all users using your application

Measurements are any value that can be recorded like daily steps, a mood rating, or apples eaten.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| Application&#39;s OAuth2 access token | [optional] 
 **clientId** | **string**| The ID of the client application which originally stored the measurement | [optional] 
 **connectorId** | **int32**| The id for the connector data source from which the measurement was obtained | [optional] 
 **variableId** | **int32**| ID of the variable for which we are creating the measurement records | [optional] 
 **sourceId** | **int32**| Application or device used to record the measurement values | [optional] 
 **startTime** | **string**| start time for the measurement event. Use ISO 8601 datetime format  | [optional] 
 **value** | **float32**| The value of the measurement after conversion to the default unit for that variable | [optional] 
 **unitId** | **int32**| The default unit id for the variable | [optional] 
 **originalValue** | **float32**| Unconverted value of measurement as originally posted (before conversion to default unit) | [optional] 
 **originalUnitId** | **int32**| Unit id of the measurement as originally submitted | [optional] 
 **duration** | **int32**| Duration of the event being measurement in seconds | [optional] 
 **note** | **string**| An optional note the user may include with their measurement | [optional] 
 **latitude** | **float32**| Latitude at which the measurement was taken | [optional] 
 **longitude** | **float32**| Longitude at which the measurement was taken | [optional] 
 **location** | **string**| Optional human readable name for the location where the measurement was recorded | [optional] 
 **createdAt** | **string**| When the record was first created. Use ISO 8601 datetime format  | [optional] 
 **updatedAt** | **string**| When the record was last updated. Use ISO 8601 datetime format  | [optional] 
 **error_** | **string**| An error message if there is a problem with the measurement | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records. | [optional] 
 **offset** | **int32**| OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned. | [optional] 
 **sort** | **string**| Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order. | [optional] 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[oauth2](../README.md#oauth2), [internalApiKey](../README.md#internalApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2ApplicationTrackingRemindersGet**
> InlineResponse2001 V2ApplicationTrackingRemindersGet($accessToken, $clientId, $createdAt, $updatedAt, $limit, $offset, $sort)

Get tracking reminders

Get the variable id, frequency, and default value for the user tracking reminders 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **clientId** | **string**| The ID of the client application which last created or updated this trackingReminder | [optional] 
 **createdAt** | **string**| When the record was first created. Use ISO 8601 datetime format  | [optional] 
 **updatedAt** | **string**| When the record was last updated. Use ISO 8601 datetime format  | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records. | [optional] 
 **offset** | **int32**| OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned. | [optional] 
 **sort** | **string**| Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order. | [optional] 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[oauth2](../README.md#oauth2), [internalApiKey](../README.md#internalApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2ApplicationUpdatesGet**
> InlineResponse2006 V2ApplicationUpdatesGet($accessToken, $connectorId, $numberOfMeasurements, $success, $message, $createdAt, $updatedAt, $limit, $offset, $sort)

Get all Updates

Get all Updates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| Application&#39;s OAuth2 access token | [optional] 
 **connectorId** | **int32**| connector_id | [optional] 
 **numberOfMeasurements** | **int32**| number_of_measurements | [optional] 
 **success** | **bool**| success | [optional] 
 **message** | **string**| message | [optional] 
 **createdAt** | **string**| When the record was first created. Use ISO 8601 datetime format  | [optional] 
 **updatedAt** | **string**| When the record was last updated. Use ISO 8601 datetime format  | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records. | [optional] 
 **offset** | **int32**| OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned. | [optional] 
 **sort** | **string**| Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order. | [optional] 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[oauth2](../README.md#oauth2), [internalApiKey](../README.md#internalApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2ApplicationUserVariableRelationshipsGet**
> InlineResponse2007 V2ApplicationUserVariableRelationshipsGet($accessToken, $id, $confidenceLevel, $confidenceScore, $direction, $durationOfAction, $errorMessage, $onsetDelay, $outcomeVariableId, $predictorVariableId, $predictorUnitId, $sinnRank, $strengthLevel, $strengthScore, $vote, $valuePredictingHighOutcome, $valuePredictingLowOutcome, $limit, $offset, $sort)

Get all UserVariableRelationships

Get all UserVariableRelationships


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **id** | **int32**| id | [optional] 
 **confidenceLevel** | **string**| Our confidence that a consistent predictive relationship exists based on the amount of evidence, reproducibility, and other factors | [optional] 
 **confidenceScore** | **float32**| A quantitative representation of our confidence that a consistent predictive relationship exists based on the amount of evidence, reproducibility, and other factors | [optional] 
 **direction** | **string**| Direction is positive if higher predictor values generally precede higher outcome values. Direction is negative if higher predictor values generally precede lower outcome values. | [optional] 
 **durationOfAction** | **int32**| Estimated number of seconds following the onset delay in which a stimulus produces a perceivable effect | [optional] 
 **errorMessage** | **string**| error_message | [optional] 
 **onsetDelay** | **int32**| Estimated number of seconds that pass before a stimulus produces a perceivable effect | [optional] 
 **outcomeVariableId** | **int32**| Variable ID for the outcome variable | [optional] 
 **predictorVariableId** | **int32**| Variable ID for the predictor variable | [optional] 
 **predictorUnitId** | **int32**| ID for default unit of the predictor variable | [optional] 
 **sinnRank** | **float32**| A value representative of the relevance of this predictor relative to other predictors of this outcome.  Usually used for relevancy sorting. | [optional] 
 **strengthLevel** | **string**| Can be weak, medium, or strong based on the size of the effect which the predictor appears to have on the outcome relative to other variable relationship strength scores. | [optional] 
 **strengthScore** | **float32**| A value represented to the size of the effect which the predictor appears to have on the outcome. | [optional] 
 **vote** | **string**| vote | [optional] 
 **valuePredictingHighOutcome** | **float32**| Value for the predictor variable (in it&#39;s default unit) which typically precedes an above average outcome value | [optional] 
 **valuePredictingLowOutcome** | **float32**| Value for the predictor variable (in it&#39;s default unit) which typically precedes a below average outcome value | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records. | [optional] 
 **offset** | **int32**| OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned. | [optional] 
 **sort** | **string**| Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order. | [optional] 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[oauth2](../README.md#oauth2), [internalApiKey](../README.md#internalApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2ApplicationUserVariablesGet**
> InlineResponse2008 V2ApplicationUserVariablesGet($accessToken, $clientId, $parentId, $variableId, $defaultUnitId, $minimumAllowedValue, $maximumAllowedValue, $fillingValue, $joinWith, $onsetDelay, $durationOfAction, $variableCategoryId, $updated, $public, $causeOnly, $fillingType, $numberOfMeasurements, $numberOfProcessedMeasurements, $measurementsAtLastAnalysis, $lastUnitId, $lastOriginalUnitId, $lastOriginalValue, $lastValue, $lastSourceId, $numberOfCorrelations, $status, $errorMessage, $lastSuccessfulUpdateTime, $standardDeviation, $variance, $minimumRecordedValue, $maximumRecordedValue, $mean, $median, $mostCommonUnitId, $mostCommonValue, $numberOfUniqueDailyValues, $numberOfChanges, $skewness, $kurtosis, $latitude, $longitude, $location, $createdAt, $updatedAt, $outcome, $sources, $earliestSourceTime, $latestSourceTime, $earliestMeasurementTime, $latestMeasurementTime, $earliestFillingTime, $latestFillingTime, $limit, $offset, $sort)

Get all UserVariables

Get all UserVariables


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **clientId** | **string**| The ID of the client application which last created or updated this user variable | [optional] 
 **parentId** | **int32**| ID of the parent variable if this variable has any parent | [optional] 
 **variableId** | **int32**| ID of variable | [optional] 
 **defaultUnitId** | **int32**| D of unit to use for this variable | [optional] 
 **minimumAllowedValue** | **float32**| Minimum reasonable value for this variable (uses default unit) | [optional] 
 **maximumAllowedValue** | **float32**| Maximum reasonable value for this variable (uses default unit) | [optional] 
 **fillingValue** | **float32**| Value for replacing null measurements | [optional] 
 **joinWith** | **int32**| The Variable this Variable should be joined with. If the variable is joined with some other variable then it is not shown to user in the list of variables | [optional] 
 **onsetDelay** | **int32**| Estimated number of seconds that pass before a stimulus produces a perceivable effect | [optional] 
 **durationOfAction** | **int32**| Estimated duration of time following the onset delay in which a stimulus produces a perceivable effect | [optional] 
 **variableCategoryId** | **int32**| ID of variable category | [optional] 
 **updated** | **int32**| updated | [optional] 
 **public** | **int32**| Is variable public | [optional] 
 **causeOnly** | **bool**| A value of 1 indicates that this variable is generally a cause in a causal relationship.  An example of a causeOnly variable would be a variable such as Cloud Cover which would generally not be influenced by the behaviour of the user | [optional] 
 **fillingType** | **string**| 0 -&gt; No filling, 1 -&gt; Use filling-value | [optional] 
 **numberOfMeasurements** | **int32**| Number of measurements | [optional] 
 **numberOfProcessedMeasurements** | **int32**| Number of processed measurements | [optional] 
 **measurementsAtLastAnalysis** | **int32**| Number of measurements at last analysis | [optional] 
 **lastUnitId** | **int32**| ID of last Unit | [optional] 
 **lastOriginalUnitId** | **int32**| ID of last original Unit | [optional] 
 **lastOriginalValue** | **int32**| Last original value which is stored | [optional] 
 **lastValue** | **float32**| Last Value | [optional] 
 **lastSourceId** | **int32**| ID of last source | [optional] 
 **numberOfCorrelations** | **int32**| Number of correlations for this variable | [optional] 
 **status** | **string**| status | [optional] 
 **errorMessage** | **string**| error_message | [optional] 
 **lastSuccessfulUpdateTime** | **string**| When this variable or its settings were last updated | [optional] 
 **standardDeviation** | **float32**| Standard deviation | [optional] 
 **variance** | **float32**| Variance | [optional] 
 **minimumRecordedValue** | **float32**| Minimum recorded value of this variable | [optional] 
 **maximumRecordedValue** | **float32**| Maximum recorded value of this variable | [optional] 
 **mean** | **float32**| Mean | [optional] 
 **median** | **float32**| Median | [optional] 
 **mostCommonUnitId** | **int32**| Most common Unit ID | [optional] 
 **mostCommonValue** | **float32**| Most common value | [optional] 
 **numberOfUniqueDailyValues** | **float32**| Number of unique daily values | [optional] 
 **numberOfChanges** | **int32**| Number of changes | [optional] 
 **skewness** | **float32**| Skewness | [optional] 
 **kurtosis** | **float32**| Kurtosis | [optional] 
 **latitude** | **float32**| Latitude | [optional] 
 **longitude** | **float32**| Longitude | [optional] 
 **location** | **string**| Location | [optional] 
 **createdAt** | **string**| When the record was first created. Use ISO 8601 datetime format  | [optional] 
 **updatedAt** | **string**| When the record was last updated. Use ISO 8601 datetime format  | [optional] 
 **outcome** | **bool**| Outcome variables (those with &#x60;outcome&#x60; &#x3D;&#x3D; 1) are variables for which a human would generally want to identify the influencing factors.  These include symptoms of illness, physique, mood, cognitive performance, etc.  Generally correlation calculations are only performed on outcome variables | [optional] 
 **sources** | **string**| Comma-separated list of source names to limit variables to those sources | [optional] 
 **earliestSourceTime** | **int32**| Earliest source time | [optional] 
 **latestSourceTime** | **int32**| Latest source time | [optional] 
 **earliestMeasurementTime** | **int32**| Earliest measurement time | [optional] 
 **latestMeasurementTime** | **int32**| Latest measurement time | [optional] 
 **earliestFillingTime** | **int32**| Earliest filling time | [optional] 
 **latestFillingTime** | **int32**| Latest filling time | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records. | [optional] 
 **offset** | **int32**| OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned. | [optional] 
 **sort** | **string**| Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order. | [optional] 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[oauth2](../README.md#oauth2), [internalApiKey](../README.md#internalApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2ApplicationVariableUserSourcesGet**
> InlineResponse2009 V2ApplicationVariableUserSourcesGet($accessToken, $variableId, $timestamp, $earliestMeasurementTime, $latestMeasurementTime, $createdAt, $updatedAt, $limit, $offset, $sort)

Get all VariableUserSources

Get all VariableUserSources


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **variableId** | **int32**| ID of variable | [optional] 
 **timestamp** | **int32**| Time that this measurement occurred Uses epoch minute (epoch time divided by 60) | [optional] 
 **earliestMeasurementTime** | **int32**| Earliest measurement time | [optional] 
 **latestMeasurementTime** | **int32**| Latest measurement time | [optional] 
 **createdAt** | **string**| When the record was first created. Use ISO 8601 datetime format  | [optional] 
 **updatedAt** | **string**| When the record was last updated. Use ISO 8601 datetime format  | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records. | [optional] 
 **offset** | **int32**| OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned. | [optional] 
 **sort** | **string**| Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order. | [optional] 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[oauth2](../README.md#oauth2), [internalApiKey](../README.md#internalApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2ApplicationVotesGet**
> InlineResponse20010 V2ApplicationVotesGet($accessToken, $clientId, $causeId, $effectId, $value, $createdAt, $updatedAt, $limit, $offset, $sort)

Get all Votes

Get all Votes


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **clientId** | **string**| The ID of the client application which last created or updated this vote | [optional] 
 **causeId** | **int32**| ID of predictor variable | [optional] 
 **effectId** | **int32**| ID of outcome variable | [optional] 
 **value** | **int32**| Value of Vote. 1 is for upvote. 0 is for downvote.  Otherwise, there is no vote. | [optional] 
 **createdAt** | **string**| When the record was first created. Use ISO 8601 datetime format  | [optional] 
 **updatedAt** | **string**| When the record was last updated. Use ISO 8601 datetime format  | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records. | [optional] 
 **offset** | **int32**| OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned. | [optional] 
 **sort** | **string**| Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order. | [optional] 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[oauth2](../README.md#oauth2), [internalApiKey](../README.md#internalApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


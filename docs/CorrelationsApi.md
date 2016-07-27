# \CorrelationsApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AggregatedCorrelationsGet**](CorrelationsApi.md#V1AggregatedCorrelationsGet) | **Get** /v1/aggregatedCorrelations | Get aggregated correlations
[**V1AggregatedCorrelationsPost**](CorrelationsApi.md#V1AggregatedCorrelationsPost) | **Post** /v1/aggregatedCorrelations | Store or Update a Correlation
[**V1CorrelationsGet**](CorrelationsApi.md#V1CorrelationsGet) | **Get** /v1/correlations | Get correlations
[**V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameCausesGet**](CorrelationsApi.md#V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameCausesGet) | **Get** /v1/organizations/{organizationId}/users/{userId}/variables/{variableName}/causes | Search user correlations for a given cause
[**V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameEffectsGet**](CorrelationsApi.md#V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameEffectsGet) | **Get** /v1/organizations/{organizationId}/users/{userId}/variables/{variableName}/effects | Search user correlations for a given cause
[**V1PublicCorrelationsSearchSearchGet**](CorrelationsApi.md#V1PublicCorrelationsSearchSearchGet) | **Get** /v1/public/correlations/search/{search} | Get average correlations for variables containing search term
[**V1VariablesVariableNameCausesGet**](CorrelationsApi.md#V1VariablesVariableNameCausesGet) | **Get** /v1/variables/{variableName}/causes | Search user correlations for a given effect
[**V1VariablesVariableNameEffectsGet**](CorrelationsApi.md#V1VariablesVariableNameEffectsGet) | **Get** /v1/variables/{variableName}/effects | Search user correlations for a given cause
[**V1VariablesVariableNamePublicCausesGet**](CorrelationsApi.md#V1VariablesVariableNamePublicCausesGet) | **Get** /v1/variables/{variableName}/public/causes | Search public correlations for a given effect
[**V1VariablesVariableNamePublicEffectsGet**](CorrelationsApi.md#V1VariablesVariableNamePublicEffectsGet) | **Get** /v1/variables/{variableName}/public/effects | Search public correlations for a given cause
[**V1VotesDeletePost**](CorrelationsApi.md#V1VotesDeletePost) | **Post** /v1/votes/delete | Delete vote
[**V1VotesPost**](CorrelationsApi.md#V1VotesPost) | **Post** /v1/votes | Post or update vote


# **V1AggregatedCorrelationsGet**
> []Correlation V1AggregatedCorrelationsGet($accessToken, $effect, $cause, $correlationCoefficient, $onsetDelay, $durationOfAction, $lastUpdated, $limit, $offset, $sort)

Get aggregated correlations

Get correlations based on the anonymized aggregate data from all QuantiModo users.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **effect** | **string**| ORIGINAL variable name of the effect variable for which the user desires correlations | [optional] 
 **cause** | **string**| ORIGINAL variable name of the cause variable for which the user desires correlations | [optional] 
 **correlationCoefficient** | **string**| Pearson correlation coefficient between cause and effect after lagging by onset delay and grouping by duration of action | [optional] 
 **onsetDelay** | **string**| The number of seconds which pass following a cause measurement before an effect would likely be observed. | [optional] 
 **durationOfAction** | **string**| The time in seconds over which the cause would be expected to exert a measurable effect. We have selected a default value for each variable. This default value may be replaced by a user specified by adjusting their variable user settings. | [optional] 
 **lastUpdated** | **string**| The time that this measurement was last updated in the UTC format \&quot;YYYY-MM-DDThh:mm:ss\&quot; | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. | [optional] 
 **offset** | **int32**| Now suppose you wanted to show results 11-20. You&#39;d set the offset to 10 and the limit to 10. | [optional] 
 **sort** | **int32**| Sort by given field. If the field is prefixed with &#x60;-, it will sort in descending order. | [optional] 

### Return type

[**[]Correlation**](Correlation.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AggregatedCorrelationsPost**
> V1AggregatedCorrelationsPost($body, $accessToken)

Store or Update a Correlation

Add correlation


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PostCorrelation**](PostCorrelation.md)| Provides correlation data | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

void (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CorrelationsGet**
> []Correlation V1CorrelationsGet($accessToken, $effect, $cause, $correlationCoefficient, $onsetDelay, $durationOfAction, $lastUpdated, $limit, $offset, $sort)

Get correlations

Get correlations based on data from a single user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **effect** | **string**| ORIGINAL variable name of the effect variable for which the user desires correlations | [optional] 
 **cause** | **string**| ORIGINAL variable name of the cause variable for which the user desires correlations | [optional] 
 **correlationCoefficient** | **string**| Pearson correlation coefficient between cause and effect after lagging by onset delay and grouping by duration of action | [optional] 
 **onsetDelay** | **string**| The number of seconds which pass following a cause measurement before an effect would likely be observed. | [optional] 
 **durationOfAction** | **string**| The time in seconds over which the cause would be expected to exert a measurable effect. We have selected a default value for each variable. This default value may be replaced by a user specified by adjusting their variable user settings. | [optional] 
 **lastUpdated** | **string**| The time that this measurement was last updated in the UTC format \&quot;YYYY-MM-DDThh:mm:ss\&quot; | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. | [optional] 
 **offset** | **int32**| Now suppose you wanted to show results 11-20. You&#39;d set the offset to 10 and the limit to 10. | [optional] 
 **sort** | **int32**| Sort by given field. If the field is prefixed with &#x60;-, it will sort in descending order. | [optional] 

### Return type

[**[]Correlation**](Correlation.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameCausesGet**
> []Correlation V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameCausesGet($organizationId, $userId, $variableName, $organizationToken, $accessToken, $includePublic)

Search user correlations for a given cause

Returns average of all correlations and votes for all user cause variables for a given cause. If parameter \"include_public\" is used, it also returns public correlations. User correlation overwrites or supersedes public correlation.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32**| Organization ID | 
 **userId** | **int32**| User id | 
 **variableName** | **string**| Effect variable name | 
 **organizationToken** | **string**| Organization access token | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **includePublic** | **string**| Include public correlations, Can be \&quot;1\&quot; or empty. | [optional] 

### Return type

[**[]Correlation**](Correlation.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameEffectsGet**
> []CommonResponse V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameEffectsGet($organizationId, $userId, $variableName, $organizationToken, $accessToken, $includePublic)

Search user correlations for a given cause

Returns average of all correlations and votes for all user cause variables for a given effect. If parameter \"include_public\" is used, it also returns public correlations. User correlation overwrites or supersedes public correlation.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32**| Organization ID | 
 **userId** | **int32**| User id | 
 **variableName** | **string**| Cause variable name | 
 **organizationToken** | **string**| Organization access token | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **includePublic** | **string**| Include public correlations, Can be \&quot;1\&quot; or empty. | [optional] 

### Return type

[**[]CommonResponse**](CommonResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1PublicCorrelationsSearchSearchGet**
> []Correlation V1PublicCorrelationsSearchSearchGet($search, $effectOrCause, $accessToken)

Get average correlations for variables containing search term

Returns the average correlations from all users for all public variables that contain the characters in the search query. Returns average of all users public variable correlations with a specified cause or effect.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string**| Name of the variable that you want to know the causes or effects of. | 
 **effectOrCause** | **string**| Setting this to effect indicates that the searched variable is the effect and that the causes of this variable should be returned.  cause indicates that the searched variable is the cause and the effects should be returned. | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

[**[]Correlation**](Correlation.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1VariablesVariableNameCausesGet**
> []Correlation V1VariablesVariableNameCausesGet($variableName)

Search user correlations for a given effect

Returns average of all correlations and votes for all user cause variables for a given effect


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableName** | **string**| Effect variable name | 

### Return type

[**[]Correlation**](Correlation.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1VariablesVariableNameEffectsGet**
> []Correlation V1VariablesVariableNameEffectsGet($variableName, $accessToken, $correlationCoefficient)

Search user correlations for a given cause

Returns average of all correlations and votes for all user effect variables for a given cause


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableName** | **string**| Cause variable name | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **correlationCoefficient** | **string**| You can use this to get effects with correlations greater than or less than 0 | [optional] 

### Return type

[**[]Correlation**](Correlation.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1VariablesVariableNamePublicCausesGet**
> []Correlation V1VariablesVariableNamePublicCausesGet($variableName, $accessToken, $correlationCoefficient)

Search public correlations for a given effect

Returns average of all correlations and votes for all public cause variables for a given effect


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableName** | **string**| Effect variable name | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **correlationCoefficient** | **string**| You can use this to get causes with correlations greater than or less than 0 | [optional] 

### Return type

[**[]Correlation**](Correlation.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1VariablesVariableNamePublicEffectsGet**
> []Correlation V1VariablesVariableNamePublicEffectsGet($variableName, $accessToken)

Search public correlations for a given cause

Returns average of all correlations and votes for all public cause variables for a given cause


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableName** | **string**| Cause variable name | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

[**[]Correlation**](Correlation.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1VotesDeletePost**
> CommonResponse V1VotesDeletePost($body, $accessToken)

Delete vote

Delete previously posted vote


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VoteDelete**](VoteDelete.md)| The cause and effect variable names for the predictor vote to be deleted. | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1VotesPost**
> CommonResponse V1VotesPost($body, $accessToken)

Post or update vote

This is to enable users to indicate their opinion on the plausibility of a causal relationship between a treatment and outcome. QuantiModo incorporates crowd-sourced plausibility estimations into their algorithm. This is done allowing user to indicate their view of the plausibility of each relationship with thumbs up/down buttons placed next to each prediction.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PostVote**](PostVote.md)| Contains the cause variable, effect variable, and vote value. | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


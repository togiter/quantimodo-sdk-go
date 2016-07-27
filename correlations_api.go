/* 
 * QuantiModo
 *
 * Welcome to QuantiModo API! QuantiModo makes it easy to retrieve normalized user data from a wide array of devices and applications. [Learn about QuantiModo](https://quantimo.do) or contact us at <api@quantimo.do>.         Before you get started, you will need to: * Sign in/Sign up, and add some data at [https://app.quantimo.do/api/v2/account/connectors](https://app.quantimo.do/api/v2/account/connectors) to try out the API for yourself * Create an app to get your client id and secret at [https://app.quantimo.do/api/v2/apps](https://app.quantimo.do/api/v2/apps) * As long as you're signed in, it will use your browser's cookie for authentication.  However, client applications must use OAuth2 tokens to access the API.     ## Application Endpoints These endpoints give you access to all authorized users' data for that application. ### Getting Application Token Make a `POST` request to `/api/v2/oauth/access_token`         * `grant_type` Must be `client_credentials`.         * `clientId` Your application's clientId.         * `client_secret` Your application's client_secret.         * `redirect_uri` Your application's redirect url.                ## Example Queries ### Query Options The standard query options for QuantiModo API are as described in the table below. These are the available query options in QuantiModo API: <table>            <thead>                <tr>                    <th>Parameter</th>                    <th>Description</th>                </tr>            </thead>            <tbody>                <tr>                    <td>limit</td>                    <td>The LIMIT is used to limit the number of results returned.  So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records.</td>                </tr>                <tr>                    <td>offset</td>                    <td>Suppose you wanted to show results 11-20. You'd set the    offset to 10 and the limit to 10.</td>                </tr>                <tr>                    <td>sort</td>                    <td>Sort by given field. If the field is prefixed with '-', it    will sort in descending order.</td>                </tr>            </tbody>        </table>         ### Pagination Conventions Since the maximum limit is 200 records, to get more than that you'll have to make multiple API calls and page through the results. To retrieve all the data, you can iterate through data by using the `limit` and `offset` query parameters.For example, if you want to retrieve data from 61-80 then you can use a query with the following parameters,         `/v2/variables?limit=20&offset=60`         Generally, you'll be retrieving new or updated user data. To avoid unnecessary API calls, you'll want to store your last refresh time locally.  Initially, it should be set to 0. Then whenever you make a request to get new data, you should limit the returned results to those updated since your last refresh by appending append         `?lastUpdated=(ge)&quot2013-01-D01T01:01:01&quot`         to your request.         Also for better pagination, you can get link to the records of first, last, next and previous page from response headers: * `Total-Count` - Total number of results for given query * `Link-First` - Link to get first page records * `Link-Last` - Link to get last page records * `Link-Prev` - Link to get previous records set * `Link-Next` - Link to get next records set         Remember, response header will be only sent when the record set is available. e.g. You will not get a ```Link-Last``` & ```Link-Next``` when you query for the last page.         ### Filter operators support API supports the following operators with filter parameters: <br> **Comparison operators**         Comparison operators allow you to limit results to those greater than, less than, or equal to a specified value for a specified attribute. These operators can be used with strings, numbers, and dates. The following comparison operators are available: * `gt` for `greater than` comparison * `ge` for `greater than or equal` comparison * `lt` for `less than` comparison * `le` for `less than or equal` comparison         They are included in queries using the following format:         `(<operator>)<value>`         For example, in order to filter value which is greater than 21, the following query parameter should be used:         `?value=(gt)21` <br><br> **Equals/In Operators**         It also allows filtering by the exact value of an attribute or by a set of values, depending on the type of value passed as a query parameter. If the value contains commas, the parameter is split on commas and used as array input for `IN` filtering, otherwise the exact match is applied. In order to only show records which have the value 42, the following query should be used:         `?value=42`         In order to filter records which have value 42 or 43, the following query should be used:         `?value=42,43` <br><br> **Like operators**         Like operators allow filtering using `LIKE` query. This operator is triggered if exact match operator is used, but value contains `%` sign as the first or last character. In order to filter records which category that start with `Food`, the following query should be used:         `?category=Food%` <br><br> **Negation operator**         It is possible to get negated results of a query by prefixed the operator with `!`. Some examples:         `//filter records except those with value are not 42 or 43`<br> `?value=!42,43`         `//filter records with value not greater than 21`<br> `?value=!(ge)21` <br><br> **Multiple constraints for single attribute**         It is possible to apply multiple constraints by providing an array of query filters:         Filter all records which value is greater than 20.2 and less than 20.3<br> `?value[]=(gt)20.2&value[]=(lt)20.3`         Filter all records which value is greater than 20.2 and less than 20.3 but not 20.2778<br> `?value[]=(gt)20.2&value[]=(lt)20.3&value[]=!20.2778`<br><br> 
 *
 * OpenAPI spec version: 2.0.6
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
)

type CorrelationsApi struct {
	Configuration Configuration
}

func NewCorrelationsApi() *CorrelationsApi {
	configuration := NewConfiguration()
	return &CorrelationsApi{
		Configuration: *configuration,
	}
}

func NewCorrelationsApiWithBasePath(basePath string) *CorrelationsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &CorrelationsApi{
		Configuration: *configuration,
	}
}

/**
 * Get aggregated correlations
 * Get correlations based on the anonymized aggregate data from all QuantiModo users.
 *
 * @param accessToken User&#39;s OAuth2 access token
 * @param effect ORIGINAL variable name of the effect variable for which the user desires correlations
 * @param cause ORIGINAL variable name of the cause variable for which the user desires correlations
 * @param correlationCoefficient Pearson correlation coefficient between cause and effect after lagging by onset delay and grouping by duration of action
 * @param onsetDelay The number of seconds which pass following a cause measurement before an effect would likely be observed.
 * @param durationOfAction The time in seconds over which the cause would be expected to exert a measurable effect. We have selected a default value for each variable. This default value may be replaced by a user specified by adjusting their variable user settings.
 * @param lastUpdated The time that this measurement was last updated in the UTC format \&quot;YYYY-MM-DDThh:mm:ss\&quot;
 * @param limit The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0.
 * @param offset Now suppose you wanted to show results 11-20. You&#39;d set the offset to 10 and the limit to 10.
 * @param sort Sort by given field. If the field is prefixed with &#x60;-, it will sort in descending order.
 * @return []Correlation
 */
func (a CorrelationsApi) V1AggregatedCorrelationsGet(accessToken string, effect string, cause string, correlationCoefficient string, onsetDelay string, durationOfAction string, lastUpdated string, limit int32, offset int32, sort int32) ([]Correlation, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/aggregatedCorrelations"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(oauth2)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
			queryParams.Add("effect", a.Configuration.APIClient.ParameterToString(effect, ""))
			queryParams.Add("cause", a.Configuration.APIClient.ParameterToString(cause, ""))
			queryParams.Add("correlationCoefficient", a.Configuration.APIClient.ParameterToString(correlationCoefficient, ""))
			queryParams.Add("onsetDelay", a.Configuration.APIClient.ParameterToString(onsetDelay, ""))
			queryParams.Add("durationOfAction", a.Configuration.APIClient.ParameterToString(durationOfAction, ""))
			queryParams.Add("lastUpdated", a.Configuration.APIClient.ParameterToString(lastUpdated, ""))
			queryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
			queryParams.Add("offset", a.Configuration.APIClient.ParameterToString(offset, ""))
			queryParams.Add("sort", a.Configuration.APIClient.ParameterToString(sort, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]Correlation)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Store or Update a Correlation
 * Add correlation
 *
 * @param body Provides correlation data
 * @param accessToken User&#39;s OAuth2 access token
 * @return void
 */
func (a CorrelationsApi) V1AggregatedCorrelationsPost(body PostCorrelation, accessToken string) (*APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/aggregatedCorrelations"

	// verify the required parameter 'body' is set
	if &body == nil {
		return nil, errors.New("Missing required parameter 'body' when calling CorrelationsApi->V1AggregatedCorrelationsPost")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(oauth2)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body


	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return NewAPIResponse(httpResponse.RawResponse), err
	}

	return NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Get correlations
 * Get correlations based on data from a single user.
 *
 * @param accessToken User&#39;s OAuth2 access token
 * @param effect ORIGINAL variable name of the effect variable for which the user desires correlations
 * @param cause ORIGINAL variable name of the cause variable for which the user desires correlations
 * @param correlationCoefficient Pearson correlation coefficient between cause and effect after lagging by onset delay and grouping by duration of action
 * @param onsetDelay The number of seconds which pass following a cause measurement before an effect would likely be observed.
 * @param durationOfAction The time in seconds over which the cause would be expected to exert a measurable effect. We have selected a default value for each variable. This default value may be replaced by a user specified by adjusting their variable user settings.
 * @param lastUpdated The time that this measurement was last updated in the UTC format \&quot;YYYY-MM-DDThh:mm:ss\&quot;
 * @param limit The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0.
 * @param offset Now suppose you wanted to show results 11-20. You&#39;d set the offset to 10 and the limit to 10.
 * @param sort Sort by given field. If the field is prefixed with &#x60;-, it will sort in descending order.
 * @return []Correlation
 */
func (a CorrelationsApi) V1CorrelationsGet(accessToken string, effect string, cause string, correlationCoefficient string, onsetDelay string, durationOfAction string, lastUpdated string, limit int32, offset int32, sort int32) ([]Correlation, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/correlations"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(oauth2)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
			queryParams.Add("effect", a.Configuration.APIClient.ParameterToString(effect, ""))
			queryParams.Add("cause", a.Configuration.APIClient.ParameterToString(cause, ""))
			queryParams.Add("correlationCoefficient", a.Configuration.APIClient.ParameterToString(correlationCoefficient, ""))
			queryParams.Add("onsetDelay", a.Configuration.APIClient.ParameterToString(onsetDelay, ""))
			queryParams.Add("durationOfAction", a.Configuration.APIClient.ParameterToString(durationOfAction, ""))
			queryParams.Add("lastUpdated", a.Configuration.APIClient.ParameterToString(lastUpdated, ""))
			queryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
			queryParams.Add("offset", a.Configuration.APIClient.ParameterToString(offset, ""))
			queryParams.Add("sort", a.Configuration.APIClient.ParameterToString(sort, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]Correlation)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Search user correlations for a given cause
 * Returns average of all correlations and votes for all user cause variables for a given cause. If parameter \&quot;include_public\&quot; is used, it also returns public correlations. User correlation overwrites or supersedes public correlation.
 *
 * @param organizationId Organization ID
 * @param userId User id
 * @param variableName Effect variable name
 * @param organizationToken Organization access token
 * @param accessToken User&#39;s OAuth2 access token
 * @param includePublic Include public correlations, Can be \&quot;1\&quot; or empty.
 * @return []Correlation
 */
func (a CorrelationsApi) V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameCausesGet(organizationId int32, userId int32, variableName string, organizationToken string, accessToken string, includePublic string) ([]Correlation, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/organizations/{organizationId}/users/{userId}/variables/{variableName}/causes"
	path = strings.Replace(path, "{"+"organizationId"+"}", fmt.Sprintf("%v", organizationId), -1)
	path = strings.Replace(path, "{"+"userId"+"}", fmt.Sprintf("%v", userId), -1)
	path = strings.Replace(path, "{"+"variableName"+"}", fmt.Sprintf("%v", variableName), -1)

	// verify the required parameter 'organizationId' is set
	if &organizationId == nil {
		return new([]Correlation), nil, errors.New("Missing required parameter 'organizationId' when calling CorrelationsApi->V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameCausesGet")
	}
	// verify the required parameter 'userId' is set
	if &userId == nil {
		return new([]Correlation), nil, errors.New("Missing required parameter 'userId' when calling CorrelationsApi->V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameCausesGet")
	}
	// verify the required parameter 'variableName' is set
	if &variableName == nil {
		return new([]Correlation), nil, errors.New("Missing required parameter 'variableName' when calling CorrelationsApi->V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameCausesGet")
	}
	// verify the required parameter 'organizationToken' is set
	if &organizationToken == nil {
		return new([]Correlation), nil, errors.New("Missing required parameter 'organizationToken' when calling CorrelationsApi->V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameCausesGet")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(oauth2)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
			queryParams.Add("organizationToken", a.Configuration.APIClient.ParameterToString(organizationToken, ""))
			queryParams.Add("includePublic", a.Configuration.APIClient.ParameterToString(includePublic, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]Correlation)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Search user correlations for a given cause
 * Returns average of all correlations and votes for all user cause variables for a given effect. If parameter \&quot;include_public\&quot; is used, it also returns public correlations. User correlation overwrites or supersedes public correlation.
 *
 * @param organizationId Organization ID
 * @param userId User id
 * @param variableName Cause variable name
 * @param organizationToken Organization access token
 * @param accessToken User&#39;s OAuth2 access token
 * @param includePublic Include public correlations, Can be \&quot;1\&quot; or empty.
 * @return []CommonResponse
 */
func (a CorrelationsApi) V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameEffectsGet(organizationId int32, userId int32, variableName string, organizationToken string, accessToken string, includePublic string) ([]CommonResponse, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/organizations/{organizationId}/users/{userId}/variables/{variableName}/effects"
	path = strings.Replace(path, "{"+"organizationId"+"}", fmt.Sprintf("%v", organizationId), -1)
	path = strings.Replace(path, "{"+"userId"+"}", fmt.Sprintf("%v", userId), -1)
	path = strings.Replace(path, "{"+"variableName"+"}", fmt.Sprintf("%v", variableName), -1)

	// verify the required parameter 'organizationId' is set
	if &organizationId == nil {
		return new([]CommonResponse), nil, errors.New("Missing required parameter 'organizationId' when calling CorrelationsApi->V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameEffectsGet")
	}
	// verify the required parameter 'userId' is set
	if &userId == nil {
		return new([]CommonResponse), nil, errors.New("Missing required parameter 'userId' when calling CorrelationsApi->V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameEffectsGet")
	}
	// verify the required parameter 'variableName' is set
	if &variableName == nil {
		return new([]CommonResponse), nil, errors.New("Missing required parameter 'variableName' when calling CorrelationsApi->V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameEffectsGet")
	}
	// verify the required parameter 'organizationToken' is set
	if &organizationToken == nil {
		return new([]CommonResponse), nil, errors.New("Missing required parameter 'organizationToken' when calling CorrelationsApi->V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameEffectsGet")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(oauth2)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
			queryParams.Add("organizationToken", a.Configuration.APIClient.ParameterToString(organizationToken, ""))
			queryParams.Add("includePublic", a.Configuration.APIClient.ParameterToString(includePublic, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]CommonResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Get average correlations for variables containing search term
 * Returns the average correlations from all users for all public variables that contain the characters in the search query. Returns average of all users public variable correlations with a specified cause or effect.
 *
 * @param search Name of the variable that you want to know the causes or effects of.
 * @param effectOrCause Setting this to effect indicates that the searched variable is the effect and that the causes of this variable should be returned.  cause indicates that the searched variable is the cause and the effects should be returned.
 * @param accessToken User&#39;s OAuth2 access token
 * @return []Correlation
 */
func (a CorrelationsApi) V1PublicCorrelationsSearchSearchGet(search string, effectOrCause string, accessToken string) ([]Correlation, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/public/correlations/search/{search}"
	path = strings.Replace(path, "{"+"search"+"}", fmt.Sprintf("%v", search), -1)

	// verify the required parameter 'search' is set
	if &search == nil {
		return new([]Correlation), nil, errors.New("Missing required parameter 'search' when calling CorrelationsApi->V1PublicCorrelationsSearchSearchGet")
	}
	// verify the required parameter 'effectOrCause' is set
	if &effectOrCause == nil {
		return new([]Correlation), nil, errors.New("Missing required parameter 'effectOrCause' when calling CorrelationsApi->V1PublicCorrelationsSearchSearchGet")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(oauth2)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
			queryParams.Add("effectOrCause", a.Configuration.APIClient.ParameterToString(effectOrCause, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]Correlation)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Search user correlations for a given effect
 * Returns average of all correlations and votes for all user cause variables for a given effect
 *
 * @param variableName Effect variable name
 * @return []Correlation
 */
func (a CorrelationsApi) V1VariablesVariableNameCausesGet(variableName string) ([]Correlation, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/variables/{variableName}/causes"
	path = strings.Replace(path, "{"+"variableName"+"}", fmt.Sprintf("%v", variableName), -1)

	// verify the required parameter 'variableName' is set
	if &variableName == nil {
		return new([]Correlation), nil, errors.New("Missing required parameter 'variableName' when calling CorrelationsApi->V1VariablesVariableNameCausesGet")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(oauth2)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]Correlation)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Search user correlations for a given cause
 * Returns average of all correlations and votes for all user effect variables for a given cause
 *
 * @param variableName Cause variable name
 * @param accessToken User&#39;s OAuth2 access token
 * @param correlationCoefficient You can use this to get effects with correlations greater than or less than 0
 * @return []Correlation
 */
func (a CorrelationsApi) V1VariablesVariableNameEffectsGet(variableName string, accessToken string, correlationCoefficient string) ([]Correlation, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/variables/{variableName}/effects"
	path = strings.Replace(path, "{"+"variableName"+"}", fmt.Sprintf("%v", variableName), -1)

	// verify the required parameter 'variableName' is set
	if &variableName == nil {
		return new([]Correlation), nil, errors.New("Missing required parameter 'variableName' when calling CorrelationsApi->V1VariablesVariableNameEffectsGet")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(oauth2)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
			queryParams.Add("correlationCoefficient", a.Configuration.APIClient.ParameterToString(correlationCoefficient, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]Correlation)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Search public correlations for a given effect
 * Returns average of all correlations and votes for all public cause variables for a given effect
 *
 * @param variableName Effect variable name
 * @param accessToken User&#39;s OAuth2 access token
 * @param correlationCoefficient You can use this to get causes with correlations greater than or less than 0
 * @return []Correlation
 */
func (a CorrelationsApi) V1VariablesVariableNamePublicCausesGet(variableName string, accessToken string, correlationCoefficient string) ([]Correlation, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/variables/{variableName}/public/causes"
	path = strings.Replace(path, "{"+"variableName"+"}", fmt.Sprintf("%v", variableName), -1)

	// verify the required parameter 'variableName' is set
	if &variableName == nil {
		return new([]Correlation), nil, errors.New("Missing required parameter 'variableName' when calling CorrelationsApi->V1VariablesVariableNamePublicCausesGet")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(oauth2)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
			queryParams.Add("correlationCoefficient", a.Configuration.APIClient.ParameterToString(correlationCoefficient, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]Correlation)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Search public correlations for a given cause
 * Returns average of all correlations and votes for all public cause variables for a given cause
 *
 * @param variableName Cause variable name
 * @param accessToken User&#39;s OAuth2 access token
 * @return []Correlation
 */
func (a CorrelationsApi) V1VariablesVariableNamePublicEffectsGet(variableName string, accessToken string) ([]Correlation, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/variables/{variableName}/public/effects"
	path = strings.Replace(path, "{"+"variableName"+"}", fmt.Sprintf("%v", variableName), -1)

	// verify the required parameter 'variableName' is set
	if &variableName == nil {
		return new([]Correlation), nil, errors.New("Missing required parameter 'variableName' when calling CorrelationsApi->V1VariablesVariableNamePublicEffectsGet")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(oauth2)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]Correlation)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Delete vote
 * Delete previously posted vote
 *
 * @param body The cause and effect variable names for the predictor vote to be deleted.
 * @param accessToken User&#39;s OAuth2 access token
 * @return *CommonResponse
 */
func (a CorrelationsApi) V1VotesDeletePost(body VoteDelete, accessToken string) (*CommonResponse, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/votes/delete"

	// verify the required parameter 'body' is set
	if &body == nil {
		return new(CommonResponse), nil, errors.New("Missing required parameter 'body' when calling CorrelationsApi->V1VotesDeletePost")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(oauth2)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload = new(CommonResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Post or update vote
 * This is to enable users to indicate their opinion on the plausibility of a causal relationship between a treatment and outcome. QuantiModo incorporates crowd-sourced plausibility estimations into their algorithm. This is done allowing user to indicate their view of the plausibility of each relationship with thumbs up/down buttons placed next to each prediction.
 *
 * @param body Contains the cause variable, effect variable, and vote value.
 * @param accessToken User&#39;s OAuth2 access token
 * @return *CommonResponse
 */
func (a CorrelationsApi) V1VotesPost(body PostVote, accessToken string) (*CommonResponse, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/votes"

	// verify the required parameter 'body' is set
	if &body == nil {
		return new(CommonResponse), nil, errors.New("Missing required parameter 'body' when calling CorrelationsApi->V1VotesPost")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(oauth2)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload = new(CommonResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}


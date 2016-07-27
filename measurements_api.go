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
	"os"
"io/ioutil"
"encoding/json"
)

type MeasurementsApi struct {
	Configuration Configuration
}

func NewMeasurementsApi() *MeasurementsApi {
	configuration := NewConfiguration()
	return &MeasurementsApi{
		Configuration: *configuration,
	}
}

func NewMeasurementsApiWithBasePath(basePath string) *MeasurementsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &MeasurementsApi{
		Configuration: *configuration,
	}
}

/**
 * Get measurement sources
 * Returns a list of all the apps from which measurement data is obtained.
 *
 * @return *MeasurementSource
 */
func (a MeasurementsApi) V1MeasurementSourcesGet() (*MeasurementSource, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/measurementSources"


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
	var successPayload = new(MeasurementSource)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Add a data source
 * Add a life-tracking app or device to the QuantiModo list of data sources.
 *
 * @param body An array of names of data sources you want to add.
 * @param accessToken User&#39;s OAuth2 access token
 * @return void
 */
func (a MeasurementsApi) V1MeasurementSourcesPost(body MeasurementSource, accessToken string) (*APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/measurementSources"

	// verify the required parameter 'body' is set
	if &body == nil {
		return nil, errors.New("Missing required parameter 'body' when calling MeasurementsApi->V1MeasurementSourcesPost")
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
 * Get daily measurements for this user
 * Measurements are any value that can be recorded like daily steps, a mood rating, or apples eaten. &lt;br&gt;Supported filter parameters:&lt;br&gt;&lt;ul&gt;&lt;li&gt;&lt;b&gt;value&lt;/b&gt; - Value of measurement&lt;/li&gt;&lt;li&gt;&lt;b&gt;lastUpdated&lt;/b&gt; - The time that this measurement was created or last updated in the UTC format \&quot;YYYY-MM-DDThh:mm:ss\&quot;&lt;/li&gt;&lt;/ul&gt;&lt;br&gt;
 *
 * @param variableName Name of the variable you want measurements for
 * @param accessToken User&#39;s OAuth2 access token
 * @param abbreviatedUnitName The unit your want the measurements in
 * @param startTime The lower limit of measurements returned (Iso8601)
 * @param endTime The upper limit of measurements returned (Iso8601)
 * @param groupingWidth The time (in seconds) over which measurements are grouped together
 * @param groupingTimezone The time (in seconds) over which measurements are grouped together
 * @param limit The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0.
 * @param offset Now suppose you wanted to show results 11-20. You&#39;d set the offset to 10 and the limit to 10.
 * @param sort Sort by given field. If the field is prefixed with &#x60;-, it will sort in descending order.
 * @return *Measurement
 */
func (a MeasurementsApi) V1MeasurementsDailyGet(variableName string, accessToken string, abbreviatedUnitName string, startTime string, endTime string, groupingWidth int32, groupingTimezone string, limit int32, offset int32, sort int32) (*Measurement, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/measurements/daily"

	// verify the required parameter 'variableName' is set
	if &variableName == nil {
		return new(Measurement), nil, errors.New("Missing required parameter 'variableName' when calling MeasurementsApi->V1MeasurementsDailyGet")
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
			queryParams.Add("variableName", a.Configuration.APIClient.ParameterToString(variableName, ""))
			queryParams.Add("abbreviatedUnitName", a.Configuration.APIClient.ParameterToString(abbreviatedUnitName, ""))
			queryParams.Add("startTime", a.Configuration.APIClient.ParameterToString(startTime, ""))
			queryParams.Add("endTime", a.Configuration.APIClient.ParameterToString(endTime, ""))
			queryParams.Add("groupingWidth", a.Configuration.APIClient.ParameterToString(groupingWidth, ""))
			queryParams.Add("groupingTimezone", a.Configuration.APIClient.ParameterToString(groupingTimezone, ""))
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
	var successPayload = new(Measurement)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Delete a measurement
 * Delete a previously submitted measurement
 *
 * @param body The startTime and variableId of the measurement to be deleted.
 * @return *CommonResponse
 */
func (a MeasurementsApi) V1MeasurementsDeletePost(body MeasurementDelete) (*CommonResponse, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/measurements/delete"

	// verify the required parameter 'body' is set
	if &body == nil {
		return new(CommonResponse), nil, errors.New("Missing required parameter 'body' when calling MeasurementsApi->V1MeasurementsDeletePost")
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
 * Get measurements for this user
 * Measurements are any value that can be recorded like daily steps, a mood rating, or apples eaten. &lt;br&gt;Supported filter parameters:&lt;br&gt;&lt;ul&gt;&lt;li&gt;&lt;b&gt;value&lt;/b&gt; - Value of measurement&lt;/li&gt;&lt;li&gt;&lt;b&gt;lastUpdated&lt;/b&gt; - The time that this measurement was created or last updated in the UTC format \&quot;YYYY-MM-DDThh:mm:ss\&quot;&lt;/li&gt;&lt;/ul&gt;&lt;br&gt;
 *
 * @param accessToken User&#39;s OAuth2 access token
 * @param id Measurement id
 * @param variableName Name of the variable you want measurements for
 * @param variableCategoryName Name of the variable category you want measurements for
 * @param source Name of the source you want measurements for (supports exact name match only)
 * @param value Value of measurement
 * @param lastUpdated The time that this measurement was created or last updated in the format \&quot;YYYY-MM-DDThh:mm:ss\&quot;
 * @param unit The unit you want the measurements returned in
 * @param startTime The lower limit of measurements returned (Epoch)
 * @param createdAt The time the measurement record was first created in the format YYYY-MM-DDThh:mm:ss. Time zone should be UTC and not local.
 * @param updatedAt The time the measurement record was last changed in the format YYYY-MM-DDThh:mm:ss. Time zone should be UTC and not local.
 * @param endTime The upper limit of measurements returned (Epoch)
 * @param groupingWidth The time (in seconds) over which measurements are grouped together
 * @param groupingTimezone The time (in seconds) over which measurements are grouped together
 * @param limit The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0.
 * @param offset Now suppose you wanted to show results 11-20. You&#39;d set the offset to 10 and the limit to 10.
 * @param sort Sort by given field. If the field is prefixed with &#x60;-, it will sort in descending order.
 * @return *Measurement
 */
func (a MeasurementsApi) V1MeasurementsGet(accessToken string, id int32, variableName string, variableCategoryName string, source string, value string, lastUpdated string, unit string, startTime string, createdAt string, updatedAt string, endTime string, groupingWidth int32, groupingTimezone string, limit int32, offset int32, sort int32) (*Measurement, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/measurements"


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
			queryParams.Add("id", a.Configuration.APIClient.ParameterToString(id, ""))
			queryParams.Add("variableName", a.Configuration.APIClient.ParameterToString(variableName, ""))
			queryParams.Add("variableCategoryName", a.Configuration.APIClient.ParameterToString(variableCategoryName, ""))
			queryParams.Add("source", a.Configuration.APIClient.ParameterToString(source, ""))
			queryParams.Add("value", a.Configuration.APIClient.ParameterToString(value, ""))
			queryParams.Add("lastUpdated", a.Configuration.APIClient.ParameterToString(lastUpdated, ""))
			queryParams.Add("unit", a.Configuration.APIClient.ParameterToString(unit, ""))
			queryParams.Add("startTime", a.Configuration.APIClient.ParameterToString(startTime, ""))
			queryParams.Add("createdAt", a.Configuration.APIClient.ParameterToString(createdAt, ""))
			queryParams.Add("updatedAt", a.Configuration.APIClient.ParameterToString(updatedAt, ""))
			queryParams.Add("endTime", a.Configuration.APIClient.ParameterToString(endTime, ""))
			queryParams.Add("groupingWidth", a.Configuration.APIClient.ParameterToString(groupingWidth, ""))
			queryParams.Add("groupingTimezone", a.Configuration.APIClient.ParameterToString(groupingTimezone, ""))
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
	var successPayload = new(Measurement)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Post a new set or update existing measurements to the database
 * You can submit or update multiple measurements in a \&quot;measurements\&quot; sub-array.  If the variable these measurements correspond to does not already exist in the database, it will be automatically added.  The request body should look something like [{\&quot;measurements\&quot;:[{\&quot;startTime\&quot;:1439389320,\&quot;value\&quot;:\&quot;3\&quot;}, {\&quot;startTime\&quot;:1439389319,\&quot;value\&quot;:\&quot;2\&quot;}],\&quot;name\&quot;:\&quot;Acne (out of 5)\&quot;,\&quot;source\&quot;:\&quot;QuantiModo\&quot;,\&quot;category\&quot;:\&quot;Symptoms\&quot;,\&quot;combinationOperation\&quot;:\&quot;MEAN\&quot;,\&quot;unit\&quot;:\&quot;/5\&quot;}]
 *
 * @param body An array of measurements you want to insert.
 * @param accessToken User&#39;s OAuth2 access token
 * @return void
 */
func (a MeasurementsApi) V1MeasurementsPost(body MeasurementSet, accessToken string) (*APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/measurements"

	// verify the required parameter 'body' is set
	if &body == nil {
		return nil, errors.New("Missing required parameter 'body' when calling MeasurementsApi->V1MeasurementsPost")
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
 * Get measurements range for this user
 * Get Unix time-stamp (epoch time) of the user&#39;s first and last measurements taken.
 *
 * @param sources Enter source name to limit to specific source (varchar)
 * @param user If not specified, uses currently logged in user (bigint)
 * @return *MeasurementRange
 */
func (a MeasurementsApi) V1MeasurementsRangeGet(sources string, user int32) (*MeasurementRange, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/measurementsRange"


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
		queryParams.Add("sources", a.Configuration.APIClient.ParameterToString(sources, ""))
			queryParams.Add("user", a.Configuration.APIClient.ParameterToString(user, ""))
	

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
	var successPayload = new(MeasurementRange)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Get Measurements CSV
 * Download a CSV containing all user measurements
 *
 * @param accessToken User&#39;s OAuth2 access token
 * @return **os.File
 */
func (a MeasurementsApi) V2MeasurementsCsvGet(accessToken string) (**os.File, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v2/measurements/csv"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(quantimodo_oauth2)' required
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
		"text/csv",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(*os.File)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Delete Measurement
 * Delete Measurement
 *
 * @param id id of Measurement
 * @param accessToken User&#39;s OAuth2 access token
 * @return *InlineResponse20012
 */
func (a MeasurementsApi) V2MeasurementsIdDelete(id int32, accessToken string) (*InlineResponse20012, *APIResponse, error) {

	var httpMethod = "Delete"
	// create path and map variables
	path := a.Configuration.BasePath + "/v2/measurements/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	// verify the required parameter 'id' is set
	if &id == nil {
		return new(InlineResponse20012), nil, errors.New("Missing required parameter 'id' when calling MeasurementsApi->V2MeasurementsIdDelete")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(quantimodo_oauth2)' required
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
	var successPayload = new(InlineResponse20012)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Get Measurement
 * Get Measurement
 *
 * @param id id of Measurement
 * @param accessToken User&#39;s OAuth2 access token
 * @return *InlineResponse20011
 */
func (a MeasurementsApi) V2MeasurementsIdGet(id int32, accessToken string) (*InlineResponse20011, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v2/measurements/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	// verify the required parameter 'id' is set
	if &id == nil {
		return new(InlineResponse20011), nil, errors.New("Missing required parameter 'id' when calling MeasurementsApi->V2MeasurementsIdGet")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(quantimodo_oauth2)' required
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
	var successPayload = new(InlineResponse20011)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Update Measurement
 * Update Measurement
 *
 * @param id id of Measurement
 * @param accessToken User&#39;s OAuth2 access token
 * @param body Measurement that should be updated
 * @return *InlineResponse20012
 */
func (a MeasurementsApi) V2MeasurementsIdPut(id int32, accessToken string, body Measurement) (*InlineResponse20012, *APIResponse, error) {

	var httpMethod = "Put"
	// create path and map variables
	path := a.Configuration.BasePath + "/v2/measurements/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	// verify the required parameter 'id' is set
	if &id == nil {
		return new(InlineResponse20012), nil, errors.New("Missing required parameter 'id' when calling MeasurementsApi->V2MeasurementsIdPut")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(quantimodo_oauth2)' required
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

	var successPayload = new(InlineResponse20012)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Post Request for Measurements CSV
 * Use this endpoint to schedule a CSV export containing all user measurements to be emailed to the user within 24 hours.
 *
 * @param accessToken User&#39;s OAuth2 access token
 * @return *int32
 */
func (a MeasurementsApi) V2MeasurementsRequestCsvPost(accessToken string) (*int32, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/v2/measurements/request_csv"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(quantimodo_oauth2)' required
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
	var successPayload = new(int32)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Post Request for Measurements PDF
 * Use this endpoint to schedule a PDF export containing all user measurements to be emailed to the user within 24 hours.
 *
 * @param accessToken User&#39;s OAuth2 access token
 * @return *int32
 */
func (a MeasurementsApi) V2MeasurementsRequestPdfPost(accessToken string) (*int32, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/v2/measurements/request_pdf"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(quantimodo_oauth2)' required
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
	var successPayload = new(int32)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Post Request for Measurements XLS
 * Use this endpoint to schedule a XLS export containing all user measurements to be emailed to the user within 24 hours.
 *
 * @param accessToken User&#39;s OAuth2 access token
 * @return *int32
 */
func (a MeasurementsApi) V2MeasurementsRequestXlsPost(accessToken string) (*int32, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/v2/measurements/request_xls"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(quantimodo_oauth2)' required
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
	var successPayload = new(int32)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}


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

type ApplicationEndpointsApi struct {
	Configuration Configuration
}

func NewApplicationEndpointsApi() *ApplicationEndpointsApi {
	configuration := NewConfiguration()
	return &ApplicationEndpointsApi{
		Configuration: *configuration,
	}
}

func NewApplicationEndpointsApiWithBasePath(basePath string) *ApplicationEndpointsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &ApplicationEndpointsApi{
		Configuration: *configuration,
	}
}

/**
 * Get all Connections
 * Get all Connections
 *
 * @param accessToken Application&#39;s OAuth2 access token
 * @param connectorId The id for the connector data source for which the connection is connected
 * @param connectStatus Indicates whether a connector is currently connected to a service for a user.
 * @param connectError Error message if there is a problem with authorizing this connection.
 * @param updateRequestedAt Time at which an update was requested by a user.
 * @param updateStatus Indicates whether a connector is currently updated.
 * @param updateError Indicates if there was an error during the update.
 * @param lastSuccessfulUpdatedAt The time at which the connector was last successfully updated.
 * @param createdAt When the record was first created. Use ISO 8601 datetime format 
 * @param updatedAt When the record was last updated. Use ISO 8601 datetime format 
 * @param limit The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records.
 * @param offset OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned.
 * @param sort Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order.
 * @return *InlineResponse2003
 */
func (a ApplicationEndpointsApi) V2ApplicationConnectionsGet(accessToken string, connectorId int32, connectStatus string, connectError string, updateRequestedAt string, updateStatus string, updateError string, lastSuccessfulUpdatedAt string, createdAt string, updatedAt string, limit int32, offset int32, sort string) (*InlineResponse2003, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v2/application/connections"


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
	// authentication '(internalApiKey)' required
	// set key with prefix in header
	headerParams["api_key"] = a.Configuration.GetAPIKeyWithPrefix("api_key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
			queryParams.Add("connectorId", a.Configuration.APIClient.ParameterToString(connectorId, ""))
			queryParams.Add("connectStatus", a.Configuration.APIClient.ParameterToString(connectStatus, ""))
			queryParams.Add("connectError", a.Configuration.APIClient.ParameterToString(connectError, ""))
			queryParams.Add("updateRequestedAt", a.Configuration.APIClient.ParameterToString(updateRequestedAt, ""))
			queryParams.Add("updateStatus", a.Configuration.APIClient.ParameterToString(updateStatus, ""))
			queryParams.Add("updateError", a.Configuration.APIClient.ParameterToString(updateError, ""))
			queryParams.Add("lastSuccessfulUpdatedAt", a.Configuration.APIClient.ParameterToString(lastSuccessfulUpdatedAt, ""))
			queryParams.Add("createdAt", a.Configuration.APIClient.ParameterToString(createdAt, ""))
			queryParams.Add("updatedAt", a.Configuration.APIClient.ParameterToString(updatedAt, ""))
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
	var successPayload = new(InlineResponse2003)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Get all Credentials
 * Get all Credentials
 *
 * @param accessToken Application&#39;s OAuth2 access token
 * @param connectorId The id for the connector data source from which the credential was obtained
 * @param attrKey Attribute name such as token, userid, username, or password
 * @param attrValue Encrypted value for the attribute specified
 * @param createdAt When the record was first created. Use ISO 8601 datetime format 
 * @param updatedAt When the record was last updated. Use ISO 8601 datetime format 
 * @param limit The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records.
 * @param offset OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned.
 * @param sort Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order.
 * @return *InlineResponse2004
 */
func (a ApplicationEndpointsApi) V2ApplicationCredentialsGet(accessToken string, connectorId int32, attrKey string, attrValue string, createdAt string, updatedAt string, limit int32, offset int32, sort string) (*InlineResponse2004, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v2/application/credentials"


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
	// authentication '(internalApiKey)' required
	// set key with prefix in header
	headerParams["api_key"] = a.Configuration.GetAPIKeyWithPrefix("api_key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
			queryParams.Add("connectorId", a.Configuration.APIClient.ParameterToString(connectorId, ""))
			queryParams.Add("attrKey", a.Configuration.APIClient.ParameterToString(attrKey, ""))
			queryParams.Add("attrValue", a.Configuration.APIClient.ParameterToString(attrValue, ""))
			queryParams.Add("createdAt", a.Configuration.APIClient.ParameterToString(createdAt, ""))
			queryParams.Add("updatedAt", a.Configuration.APIClient.ParameterToString(updatedAt, ""))
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
	var successPayload = new(InlineResponse2004)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Get measurements for all users using your application
 * Measurements are any value that can be recorded like daily steps, a mood rating, or apples eaten.
 *
 * @param accessToken Application&#39;s OAuth2 access token
 * @param clientId The ID of the client application which originally stored the measurement
 * @param connectorId The id for the connector data source from which the measurement was obtained
 * @param variableId ID of the variable for which we are creating the measurement records
 * @param sourceId Application or device used to record the measurement values
 * @param startTime start time for the measurement event. Use ISO 8601 datetime format 
 * @param value The value of the measurement after conversion to the default unit for that variable
 * @param unitId The default unit id for the variable
 * @param originalValue Unconverted value of measurement as originally posted (before conversion to default unit)
 * @param originalUnitId Unit id of the measurement as originally submitted
 * @param duration Duration of the event being measurement in seconds
 * @param note An optional note the user may include with their measurement
 * @param latitude Latitude at which the measurement was taken
 * @param longitude Longitude at which the measurement was taken
 * @param location Optional human readable name for the location where the measurement was recorded
 * @param createdAt When the record was first created. Use ISO 8601 datetime format 
 * @param updatedAt When the record was last updated. Use ISO 8601 datetime format 
 * @param error_ An error message if there is a problem with the measurement
 * @param limit The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records.
 * @param offset OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned.
 * @param sort Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order.
 * @return *InlineResponse2005
 */
func (a ApplicationEndpointsApi) V2ApplicationMeasurementsGet(accessToken string, clientId string, connectorId int32, variableId int32, sourceId int32, startTime string, value float32, unitId int32, originalValue float32, originalUnitId int32, duration int32, note string, latitude float32, longitude float32, location string, createdAt string, updatedAt string, error_ string, limit int32, offset int32, sort string) (*InlineResponse2005, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v2/application/measurements"


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
	// authentication '(internalApiKey)' required
	// set key with prefix in header
	headerParams["api_key"] = a.Configuration.GetAPIKeyWithPrefix("api_key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
			queryParams.Add("clientId", a.Configuration.APIClient.ParameterToString(clientId, ""))
			queryParams.Add("connectorId", a.Configuration.APIClient.ParameterToString(connectorId, ""))
			queryParams.Add("variableId", a.Configuration.APIClient.ParameterToString(variableId, ""))
			queryParams.Add("sourceId", a.Configuration.APIClient.ParameterToString(sourceId, ""))
			queryParams.Add("startTime", a.Configuration.APIClient.ParameterToString(startTime, ""))
			queryParams.Add("value", a.Configuration.APIClient.ParameterToString(value, ""))
			queryParams.Add("unitId", a.Configuration.APIClient.ParameterToString(unitId, ""))
			queryParams.Add("originalValue", a.Configuration.APIClient.ParameterToString(originalValue, ""))
			queryParams.Add("originalUnitId", a.Configuration.APIClient.ParameterToString(originalUnitId, ""))
			queryParams.Add("duration", a.Configuration.APIClient.ParameterToString(duration, ""))
			queryParams.Add("note", a.Configuration.APIClient.ParameterToString(note, ""))
			queryParams.Add("latitude", a.Configuration.APIClient.ParameterToString(latitude, ""))
			queryParams.Add("longitude", a.Configuration.APIClient.ParameterToString(longitude, ""))
			queryParams.Add("location", a.Configuration.APIClient.ParameterToString(location, ""))
			queryParams.Add("createdAt", a.Configuration.APIClient.ParameterToString(createdAt, ""))
			queryParams.Add("updatedAt", a.Configuration.APIClient.ParameterToString(updatedAt, ""))
			queryParams.Add("error_", a.Configuration.APIClient.ParameterToString(error_, ""))
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
	var successPayload = new(InlineResponse2005)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Get tracking reminders
 * Get the variable id, frequency, and default value for the user tracking reminders 
 *
 * @param accessToken User&#39;s OAuth2 access token
 * @param clientId The ID of the client application which last created or updated this trackingReminder
 * @param createdAt When the record was first created. Use ISO 8601 datetime format 
 * @param updatedAt When the record was last updated. Use ISO 8601 datetime format 
 * @param limit The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records.
 * @param offset OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned.
 * @param sort Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order.
 * @return *InlineResponse2001
 */
func (a ApplicationEndpointsApi) V2ApplicationTrackingRemindersGet(accessToken string, clientId string, createdAt string, updatedAt string, limit int32, offset int32, sort string) (*InlineResponse2001, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v2/application/trackingReminders"


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
	// authentication '(internalApiKey)' required
	// set key with prefix in header
	headerParams["api_key"] = a.Configuration.GetAPIKeyWithPrefix("api_key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
			queryParams.Add("clientId", a.Configuration.APIClient.ParameterToString(clientId, ""))
			queryParams.Add("createdAt", a.Configuration.APIClient.ParameterToString(createdAt, ""))
			queryParams.Add("updatedAt", a.Configuration.APIClient.ParameterToString(updatedAt, ""))
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
	var successPayload = new(InlineResponse2001)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Get all Updates
 * Get all Updates
 *
 * @param accessToken Application&#39;s OAuth2 access token
 * @param connectorId connector_id
 * @param numberOfMeasurements number_of_measurements
 * @param success success
 * @param message message
 * @param createdAt When the record was first created. Use ISO 8601 datetime format 
 * @param updatedAt When the record was last updated. Use ISO 8601 datetime format 
 * @param limit The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records.
 * @param offset OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned.
 * @param sort Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order.
 * @return *InlineResponse2006
 */
func (a ApplicationEndpointsApi) V2ApplicationUpdatesGet(accessToken string, connectorId int32, numberOfMeasurements int32, success bool, message string, createdAt string, updatedAt string, limit int32, offset int32, sort string) (*InlineResponse2006, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v2/application/updates"


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
	// authentication '(internalApiKey)' required
	// set key with prefix in header
	headerParams["api_key"] = a.Configuration.GetAPIKeyWithPrefix("api_key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
			queryParams.Add("connectorId", a.Configuration.APIClient.ParameterToString(connectorId, ""))
			queryParams.Add("numberOfMeasurements", a.Configuration.APIClient.ParameterToString(numberOfMeasurements, ""))
			queryParams.Add("success", a.Configuration.APIClient.ParameterToString(success, ""))
			queryParams.Add("message", a.Configuration.APIClient.ParameterToString(message, ""))
			queryParams.Add("createdAt", a.Configuration.APIClient.ParameterToString(createdAt, ""))
			queryParams.Add("updatedAt", a.Configuration.APIClient.ParameterToString(updatedAt, ""))
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
	var successPayload = new(InlineResponse2006)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Get all UserVariableRelationships
 * Get all UserVariableRelationships
 *
 * @param accessToken User&#39;s OAuth2 access token
 * @param id id
 * @param confidenceLevel Our confidence that a consistent predictive relationship exists based on the amount of evidence, reproducibility, and other factors
 * @param confidenceScore A quantitative representation of our confidence that a consistent predictive relationship exists based on the amount of evidence, reproducibility, and other factors
 * @param direction Direction is positive if higher predictor values generally precede higher outcome values. Direction is negative if higher predictor values generally precede lower outcome values.
 * @param durationOfAction Estimated number of seconds following the onset delay in which a stimulus produces a perceivable effect
 * @param errorMessage error_message
 * @param onsetDelay Estimated number of seconds that pass before a stimulus produces a perceivable effect
 * @param outcomeVariableId Variable ID for the outcome variable
 * @param predictorVariableId Variable ID for the predictor variable
 * @param predictorUnitId ID for default unit of the predictor variable
 * @param sinnRank A value representative of the relevance of this predictor relative to other predictors of this outcome.  Usually used for relevancy sorting.
 * @param strengthLevel Can be weak, medium, or strong based on the size of the effect which the predictor appears to have on the outcome relative to other variable relationship strength scores.
 * @param strengthScore A value represented to the size of the effect which the predictor appears to have on the outcome.
 * @param vote vote
 * @param valuePredictingHighOutcome Value for the predictor variable (in it&#39;s default unit) which typically precedes an above average outcome value
 * @param valuePredictingLowOutcome Value for the predictor variable (in it&#39;s default unit) which typically precedes a below average outcome value
 * @param limit The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records.
 * @param offset OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned.
 * @param sort Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order.
 * @return *InlineResponse2007
 */
func (a ApplicationEndpointsApi) V2ApplicationUserVariableRelationshipsGet(accessToken string, id int32, confidenceLevel string, confidenceScore float32, direction string, durationOfAction int32, errorMessage string, onsetDelay int32, outcomeVariableId int32, predictorVariableId int32, predictorUnitId int32, sinnRank float32, strengthLevel string, strengthScore float32, vote string, valuePredictingHighOutcome float32, valuePredictingLowOutcome float32, limit int32, offset int32, sort string) (*InlineResponse2007, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v2/application/userVariableRelationships"


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
	// authentication '(internalApiKey)' required
	// set key with prefix in header
	headerParams["api_key"] = a.Configuration.GetAPIKeyWithPrefix("api_key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
			queryParams.Add("id", a.Configuration.APIClient.ParameterToString(id, ""))
			queryParams.Add("confidenceLevel", a.Configuration.APIClient.ParameterToString(confidenceLevel, ""))
			queryParams.Add("confidenceScore", a.Configuration.APIClient.ParameterToString(confidenceScore, ""))
			queryParams.Add("direction", a.Configuration.APIClient.ParameterToString(direction, ""))
			queryParams.Add("durationOfAction", a.Configuration.APIClient.ParameterToString(durationOfAction, ""))
			queryParams.Add("errorMessage", a.Configuration.APIClient.ParameterToString(errorMessage, ""))
			queryParams.Add("onsetDelay", a.Configuration.APIClient.ParameterToString(onsetDelay, ""))
			queryParams.Add("outcomeVariableId", a.Configuration.APIClient.ParameterToString(outcomeVariableId, ""))
			queryParams.Add("predictorVariableId", a.Configuration.APIClient.ParameterToString(predictorVariableId, ""))
			queryParams.Add("predictorUnitId", a.Configuration.APIClient.ParameterToString(predictorUnitId, ""))
			queryParams.Add("sinnRank", a.Configuration.APIClient.ParameterToString(sinnRank, ""))
			queryParams.Add("strengthLevel", a.Configuration.APIClient.ParameterToString(strengthLevel, ""))
			queryParams.Add("strengthScore", a.Configuration.APIClient.ParameterToString(strengthScore, ""))
			queryParams.Add("vote", a.Configuration.APIClient.ParameterToString(vote, ""))
			queryParams.Add("valuePredictingHighOutcome", a.Configuration.APIClient.ParameterToString(valuePredictingHighOutcome, ""))
			queryParams.Add("valuePredictingLowOutcome", a.Configuration.APIClient.ParameterToString(valuePredictingLowOutcome, ""))
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
	var successPayload = new(InlineResponse2007)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Get all UserVariables
 * Get all UserVariables
 *
 * @param accessToken User&#39;s OAuth2 access token
 * @param clientId The ID of the client application which last created or updated this user variable
 * @param parentId ID of the parent variable if this variable has any parent
 * @param variableId ID of variable
 * @param defaultUnitId D of unit to use for this variable
 * @param minimumAllowedValue Minimum reasonable value for this variable (uses default unit)
 * @param maximumAllowedValue Maximum reasonable value for this variable (uses default unit)
 * @param fillingValue Value for replacing null measurements
 * @param joinWith The Variable this Variable should be joined with. If the variable is joined with some other variable then it is not shown to user in the list of variables
 * @param onsetDelay Estimated number of seconds that pass before a stimulus produces a perceivable effect
 * @param durationOfAction Estimated duration of time following the onset delay in which a stimulus produces a perceivable effect
 * @param variableCategoryId ID of variable category
 * @param updated updated
 * @param public Is variable public
 * @param causeOnly A value of 1 indicates that this variable is generally a cause in a causal relationship.  An example of a causeOnly variable would be a variable such as Cloud Cover which would generally not be influenced by the behaviour of the user
 * @param fillingType 0 -&gt; No filling, 1 -&gt; Use filling-value
 * @param numberOfMeasurements Number of measurements
 * @param numberOfProcessedMeasurements Number of processed measurements
 * @param measurementsAtLastAnalysis Number of measurements at last analysis
 * @param lastUnitId ID of last Unit
 * @param lastOriginalUnitId ID of last original Unit
 * @param lastOriginalValue Last original value which is stored
 * @param lastValue Last Value
 * @param lastSourceId ID of last source
 * @param numberOfCorrelations Number of correlations for this variable
 * @param status status
 * @param errorMessage error_message
 * @param lastSuccessfulUpdateTime When this variable or its settings were last updated
 * @param standardDeviation Standard deviation
 * @param variance Variance
 * @param minimumRecordedValue Minimum recorded value of this variable
 * @param maximumRecordedValue Maximum recorded value of this variable
 * @param mean Mean
 * @param median Median
 * @param mostCommonUnitId Most common Unit ID
 * @param mostCommonValue Most common value
 * @param numberOfUniqueDailyValues Number of unique daily values
 * @param numberOfChanges Number of changes
 * @param skewness Skewness
 * @param kurtosis Kurtosis
 * @param latitude Latitude
 * @param longitude Longitude
 * @param location Location
 * @param createdAt When the record was first created. Use ISO 8601 datetime format 
 * @param updatedAt When the record was last updated. Use ISO 8601 datetime format 
 * @param outcome Outcome variables (those with &#x60;outcome&#x60; &#x3D;&#x3D; 1) are variables for which a human would generally want to identify the influencing factors.  These include symptoms of illness, physique, mood, cognitive performance, etc.  Generally correlation calculations are only performed on outcome variables
 * @param sources Comma-separated list of source names to limit variables to those sources
 * @param earliestSourceTime Earliest source time
 * @param latestSourceTime Latest source time
 * @param earliestMeasurementTime Earliest measurement time
 * @param latestMeasurementTime Latest measurement time
 * @param earliestFillingTime Earliest filling time
 * @param latestFillingTime Latest filling time
 * @param limit The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records.
 * @param offset OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned.
 * @param sort Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order.
 * @return *InlineResponse2008
 */
func (a ApplicationEndpointsApi) V2ApplicationUserVariablesGet(accessToken string, clientId string, parentId int32, variableId int32, defaultUnitId int32, minimumAllowedValue float32, maximumAllowedValue float32, fillingValue float32, joinWith int32, onsetDelay int32, durationOfAction int32, variableCategoryId int32, updated int32, public int32, causeOnly bool, fillingType string, numberOfMeasurements int32, numberOfProcessedMeasurements int32, measurementsAtLastAnalysis int32, lastUnitId int32, lastOriginalUnitId int32, lastOriginalValue int32, lastValue float32, lastSourceId int32, numberOfCorrelations int32, status string, errorMessage string, lastSuccessfulUpdateTime string, standardDeviation float32, variance float32, minimumRecordedValue float32, maximumRecordedValue float32, mean float32, median float32, mostCommonUnitId int32, mostCommonValue float32, numberOfUniqueDailyValues float32, numberOfChanges int32, skewness float32, kurtosis float32, latitude float32, longitude float32, location string, createdAt string, updatedAt string, outcome bool, sources string, earliestSourceTime int32, latestSourceTime int32, earliestMeasurementTime int32, latestMeasurementTime int32, earliestFillingTime int32, latestFillingTime int32, limit int32, offset int32, sort string) (*InlineResponse2008, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v2/application/userVariables"


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
	// authentication '(internalApiKey)' required
	// set key with prefix in header
	headerParams["api_key"] = a.Configuration.GetAPIKeyWithPrefix("api_key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
			queryParams.Add("clientId", a.Configuration.APIClient.ParameterToString(clientId, ""))
			queryParams.Add("parentId", a.Configuration.APIClient.ParameterToString(parentId, ""))
			queryParams.Add("variableId", a.Configuration.APIClient.ParameterToString(variableId, ""))
			queryParams.Add("defaultUnitId", a.Configuration.APIClient.ParameterToString(defaultUnitId, ""))
			queryParams.Add("minimumAllowedValue", a.Configuration.APIClient.ParameterToString(minimumAllowedValue, ""))
			queryParams.Add("maximumAllowedValue", a.Configuration.APIClient.ParameterToString(maximumAllowedValue, ""))
			queryParams.Add("fillingValue", a.Configuration.APIClient.ParameterToString(fillingValue, ""))
			queryParams.Add("joinWith", a.Configuration.APIClient.ParameterToString(joinWith, ""))
			queryParams.Add("onsetDelay", a.Configuration.APIClient.ParameterToString(onsetDelay, ""))
			queryParams.Add("durationOfAction", a.Configuration.APIClient.ParameterToString(durationOfAction, ""))
			queryParams.Add("variableCategoryId", a.Configuration.APIClient.ParameterToString(variableCategoryId, ""))
			queryParams.Add("updated", a.Configuration.APIClient.ParameterToString(updated, ""))
			queryParams.Add("public", a.Configuration.APIClient.ParameterToString(public, ""))
			queryParams.Add("causeOnly", a.Configuration.APIClient.ParameterToString(causeOnly, ""))
			queryParams.Add("fillingType", a.Configuration.APIClient.ParameterToString(fillingType, ""))
			queryParams.Add("numberOfMeasurements", a.Configuration.APIClient.ParameterToString(numberOfMeasurements, ""))
			queryParams.Add("numberOfProcessedMeasurements", a.Configuration.APIClient.ParameterToString(numberOfProcessedMeasurements, ""))
			queryParams.Add("measurementsAtLastAnalysis", a.Configuration.APIClient.ParameterToString(measurementsAtLastAnalysis, ""))
			queryParams.Add("lastUnitId", a.Configuration.APIClient.ParameterToString(lastUnitId, ""))
			queryParams.Add("lastOriginalUnitId", a.Configuration.APIClient.ParameterToString(lastOriginalUnitId, ""))
			queryParams.Add("lastOriginalValue", a.Configuration.APIClient.ParameterToString(lastOriginalValue, ""))
			queryParams.Add("lastValue", a.Configuration.APIClient.ParameterToString(lastValue, ""))
			queryParams.Add("lastSourceId", a.Configuration.APIClient.ParameterToString(lastSourceId, ""))
			queryParams.Add("numberOfCorrelations", a.Configuration.APIClient.ParameterToString(numberOfCorrelations, ""))
			queryParams.Add("status", a.Configuration.APIClient.ParameterToString(status, ""))
			queryParams.Add("errorMessage", a.Configuration.APIClient.ParameterToString(errorMessage, ""))
			queryParams.Add("lastSuccessfulUpdateTime", a.Configuration.APIClient.ParameterToString(lastSuccessfulUpdateTime, ""))
			queryParams.Add("standardDeviation", a.Configuration.APIClient.ParameterToString(standardDeviation, ""))
			queryParams.Add("variance", a.Configuration.APIClient.ParameterToString(variance, ""))
			queryParams.Add("minimumRecordedValue", a.Configuration.APIClient.ParameterToString(minimumRecordedValue, ""))
			queryParams.Add("maximumRecordedValue", a.Configuration.APIClient.ParameterToString(maximumRecordedValue, ""))
			queryParams.Add("mean", a.Configuration.APIClient.ParameterToString(mean, ""))
			queryParams.Add("median", a.Configuration.APIClient.ParameterToString(median, ""))
			queryParams.Add("mostCommonUnitId", a.Configuration.APIClient.ParameterToString(mostCommonUnitId, ""))
			queryParams.Add("mostCommonValue", a.Configuration.APIClient.ParameterToString(mostCommonValue, ""))
			queryParams.Add("numberOfUniqueDailyValues", a.Configuration.APIClient.ParameterToString(numberOfUniqueDailyValues, ""))
			queryParams.Add("numberOfChanges", a.Configuration.APIClient.ParameterToString(numberOfChanges, ""))
			queryParams.Add("skewness", a.Configuration.APIClient.ParameterToString(skewness, ""))
			queryParams.Add("kurtosis", a.Configuration.APIClient.ParameterToString(kurtosis, ""))
			queryParams.Add("latitude", a.Configuration.APIClient.ParameterToString(latitude, ""))
			queryParams.Add("longitude", a.Configuration.APIClient.ParameterToString(longitude, ""))
			queryParams.Add("location", a.Configuration.APIClient.ParameterToString(location, ""))
			queryParams.Add("createdAt", a.Configuration.APIClient.ParameterToString(createdAt, ""))
			queryParams.Add("updatedAt", a.Configuration.APIClient.ParameterToString(updatedAt, ""))
			queryParams.Add("outcome", a.Configuration.APIClient.ParameterToString(outcome, ""))
			queryParams.Add("sources", a.Configuration.APIClient.ParameterToString(sources, ""))
			queryParams.Add("earliestSourceTime", a.Configuration.APIClient.ParameterToString(earliestSourceTime, ""))
			queryParams.Add("latestSourceTime", a.Configuration.APIClient.ParameterToString(latestSourceTime, ""))
			queryParams.Add("earliestMeasurementTime", a.Configuration.APIClient.ParameterToString(earliestMeasurementTime, ""))
			queryParams.Add("latestMeasurementTime", a.Configuration.APIClient.ParameterToString(latestMeasurementTime, ""))
			queryParams.Add("earliestFillingTime", a.Configuration.APIClient.ParameterToString(earliestFillingTime, ""))
			queryParams.Add("latestFillingTime", a.Configuration.APIClient.ParameterToString(latestFillingTime, ""))
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
	var successPayload = new(InlineResponse2008)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Get all VariableUserSources
 * Get all VariableUserSources
 *
 * @param accessToken User&#39;s OAuth2 access token
 * @param variableId ID of variable
 * @param timestamp Time that this measurement occurred Uses epoch minute (epoch time divided by 60)
 * @param earliestMeasurementTime Earliest measurement time
 * @param latestMeasurementTime Latest measurement time
 * @param createdAt When the record was first created. Use ISO 8601 datetime format 
 * @param updatedAt When the record was last updated. Use ISO 8601 datetime format 
 * @param limit The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records.
 * @param offset OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned.
 * @param sort Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order.
 * @return *InlineResponse2009
 */
func (a ApplicationEndpointsApi) V2ApplicationVariableUserSourcesGet(accessToken string, variableId int32, timestamp int32, earliestMeasurementTime int32, latestMeasurementTime int32, createdAt string, updatedAt string, limit int32, offset int32, sort string) (*InlineResponse2009, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v2/application/variableUserSources"


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
	// authentication '(internalApiKey)' required
	// set key with prefix in header
	headerParams["api_key"] = a.Configuration.GetAPIKeyWithPrefix("api_key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
			queryParams.Add("variableId", a.Configuration.APIClient.ParameterToString(variableId, ""))
			queryParams.Add("timestamp", a.Configuration.APIClient.ParameterToString(timestamp, ""))
			queryParams.Add("earliestMeasurementTime", a.Configuration.APIClient.ParameterToString(earliestMeasurementTime, ""))
			queryParams.Add("latestMeasurementTime", a.Configuration.APIClient.ParameterToString(latestMeasurementTime, ""))
			queryParams.Add("createdAt", a.Configuration.APIClient.ParameterToString(createdAt, ""))
			queryParams.Add("updatedAt", a.Configuration.APIClient.ParameterToString(updatedAt, ""))
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
	var successPayload = new(InlineResponse2009)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Get all Votes
 * Get all Votes
 *
 * @param accessToken User&#39;s OAuth2 access token
 * @param clientId The ID of the client application which last created or updated this vote
 * @param causeId ID of predictor variable
 * @param effectId ID of outcome variable
 * @param value Value of Vote. 1 is for upvote. 0 is for downvote.  Otherwise, there is no vote.
 * @param createdAt When the record was first created. Use ISO 8601 datetime format 
 * @param updatedAt When the record was last updated. Use ISO 8601 datetime format 
 * @param limit The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records.
 * @param offset OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned.
 * @param sort Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order.
 * @return *InlineResponse20010
 */
func (a ApplicationEndpointsApi) V2ApplicationVotesGet(accessToken string, clientId string, causeId int32, effectId int32, value int32, createdAt string, updatedAt string, limit int32, offset int32, sort string) (*InlineResponse20010, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v2/application/votes"


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
	// authentication '(internalApiKey)' required
	// set key with prefix in header
	headerParams["api_key"] = a.Configuration.GetAPIKeyWithPrefix("api_key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("accessToken", a.Configuration.APIClient.ParameterToString(accessToken, ""))
			queryParams.Add("clientId", a.Configuration.APIClient.ParameterToString(clientId, ""))
			queryParams.Add("causeId", a.Configuration.APIClient.ParameterToString(causeId, ""))
			queryParams.Add("effectId", a.Configuration.APIClient.ParameterToString(effectId, ""))
			queryParams.Add("value", a.Configuration.APIClient.ParameterToString(value, ""))
			queryParams.Add("createdAt", a.Configuration.APIClient.ParameterToString(createdAt, ""))
			queryParams.Add("updatedAt", a.Configuration.APIClient.ParameterToString(updatedAt, ""))
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
	var successPayload = new(InlineResponse20010)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}


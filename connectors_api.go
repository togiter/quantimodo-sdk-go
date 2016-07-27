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

type ConnectorsApi struct {
	Configuration Configuration
}

func NewConnectorsApi() *ConnectorsApi {
	configuration := NewConfiguration()
	return &ConnectorsApi{
		Configuration: *configuration,
	}
}

func NewConnectorsApiWithBasePath(basePath string) *ConnectorsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &ConnectorsApi{
		Configuration: *configuration,
	}
}

/**
 * Get embeddable connect javascript
 * Get embeddable connect javascript. Usage:    - Embedding in applications with popups for 3rd-party authentication windows.      Use &#x60;qmSetupInPopup&#x60; function after connecting &#x60;connect.js&#x60;.    - Embedding in applications with popups for 3rd-party authentication windows.      Requires a selector to block. It will be embedded in this block.      Use &#x60;qmSetupOnPage&#x60; function after connecting &#x60;connect.js&#x60;.    - Embedding in mobile applications without popups for 3rd-party authentication.      Use &#x60;qmSetupOnMobile&#x60; function after connecting &#x60;connect.js&#x60;.      if using the MoodiModo Clones, Use &#x60;qmSetupOnIonic&#x60; function after connecting &#x60;connect.js&#x60;. 
 *
 * @param accessToken User&#39;s OAuth2 access token
 * @return void
 */
func (a ConnectorsApi) V1ConnectJsGet(accessToken string) (*APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/connect.js"


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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/x-javascript",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}

	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return NewAPIResponse(httpResponse.RawResponse), err
	}

	return NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Mobile connect page
 * This page is designed to be opened in a webview.  Instead of using popup authentication boxes, it uses redirection. You can include the user&#39;s access_token as a URL parameter like https://app.quantimo.do/api/v1/connect/mobile?access_token&#x3D;123
 *
 * @param accessToken User OAuth access token
 * @return void
 */
func (a ConnectorsApi) V1ConnectMobileGet(accessToken string) (*APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/connect/mobile"

	// verify the required parameter 'accessToken' is set
	if &accessToken == nil {
		return nil, errors.New("Missing required parameter 'accessToken' when calling ConnectorsApi->V1ConnectMobileGet")
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
	// authentication '(internalApiKey)' required
	// set key with prefix in header
	headerParams["api_key"] = a.Configuration.GetAPIKeyWithPrefix("api_key")
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
		"text/html",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}

	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return NewAPIResponse(httpResponse.RawResponse), err
	}

	return NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Obtain a token from 3rd party data source
 * Attempt to obtain a token from the data provider, store it in the database. With this, the connector to continue to obtain new user data until the token is revoked.
 *
 * @param connector Lowercase system name of the source application or device. Get a list of available connectors from the /connectors/list endpoint.
 * @param accessToken User&#39;s OAuth2 access token
 * @return void
 */
func (a ConnectorsApi) V1ConnectorsConnectorConnectGet(connector string, accessToken string) (*APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/connectors/{connector}/connect"
	path = strings.Replace(path, "{"+"connector"+"}", fmt.Sprintf("%v", connector), -1)

	// verify the required parameter 'connector' is set
	if &connector == nil {
		return nil, errors.New("Missing required parameter 'connector' when calling ConnectorsApi->V1ConnectorsConnectorConnectGet")
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

	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return NewAPIResponse(httpResponse.RawResponse), err
	}

	return NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Connection Instructions
 * Returns instructions that describe what parameters and endpoint to use to connect to the given data provider.
 *
 * @param connector Lowercase system name of the source application or device. Get a list of available connectors from the /connectors/list endpoint.
 * @param parameters JSON Array of Parameters for the request to enable connector.
 * @param url URL which should be used to enable the connector.
 * @param usePopup Should use popup when enabling connector
 * @param accessToken User&#39;s OAuth2 access token
 * @return void
 */
func (a ConnectorsApi) V1ConnectorsConnectorConnectInstructionsGet(connector string, parameters string, url string, usePopup bool, accessToken string) (*APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/connectors/{connector}/connectInstructions"
	path = strings.Replace(path, "{"+"connector"+"}", fmt.Sprintf("%v", connector), -1)

	// verify the required parameter 'connector' is set
	if &connector == nil {
		return nil, errors.New("Missing required parameter 'connector' when calling ConnectorsApi->V1ConnectorsConnectorConnectInstructionsGet")
	}
	// verify the required parameter 'parameters' is set
	if &parameters == nil {
		return nil, errors.New("Missing required parameter 'parameters' when calling ConnectorsApi->V1ConnectorsConnectorConnectInstructionsGet")
	}
	// verify the required parameter 'url' is set
	if &url == nil {
		return nil, errors.New("Missing required parameter 'url' when calling ConnectorsApi->V1ConnectorsConnectorConnectInstructionsGet")
	}
	// verify the required parameter 'usePopup' is set
	if &usePopup == nil {
		return nil, errors.New("Missing required parameter 'usePopup' when calling ConnectorsApi->V1ConnectorsConnectorConnectInstructionsGet")
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
			queryParams.Add("parameters", a.Configuration.APIClient.ParameterToString(parameters, ""))
			queryParams.Add("url", a.Configuration.APIClient.ParameterToString(url, ""))
			queryParams.Add("usePopup", a.Configuration.APIClient.ParameterToString(usePopup, ""))
	

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

	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return NewAPIResponse(httpResponse.RawResponse), err
	}

	return NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Connect Parameter
 * Returns instructions that describe what parameters and endpoint to use to connect to the given data provider.
 *
 * @param connector Lowercase system name of the source application or device. Get a list of available connectors from the /connectors/list endpoint.
 * @param displayName Name of the parameter that is user visible in the form
 * @param key Name of the property that the user has to enter such as username or password Connector (used in HTTP request)
 * @param placeholder Placeholder hint value for the parameter input tag.
 * @param type_ Type of input field such as those found here http://www.w3schools.com/tags/tag_input.asp
 * @param usePopup Should use popup when enabling connector
 * @param accessToken User&#39;s OAuth2 access token
 * @param defaultValue Default parameter value
 * @return *ConnectorInstruction
 */
func (a ConnectorsApi) V1ConnectorsConnectorConnectParameterGet(connector string, displayName string, key string, placeholder string, type_ string, usePopup bool, accessToken string, defaultValue string) (*ConnectorInstruction, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/connectors/{connector}/connectParameter"
	path = strings.Replace(path, "{"+"connector"+"}", fmt.Sprintf("%v", connector), -1)

	// verify the required parameter 'connector' is set
	if &connector == nil {
		return new(ConnectorInstruction), nil, errors.New("Missing required parameter 'connector' when calling ConnectorsApi->V1ConnectorsConnectorConnectParameterGet")
	}
	// verify the required parameter 'displayName' is set
	if &displayName == nil {
		return new(ConnectorInstruction), nil, errors.New("Missing required parameter 'displayName' when calling ConnectorsApi->V1ConnectorsConnectorConnectParameterGet")
	}
	// verify the required parameter 'key' is set
	if &key == nil {
		return new(ConnectorInstruction), nil, errors.New("Missing required parameter 'key' when calling ConnectorsApi->V1ConnectorsConnectorConnectParameterGet")
	}
	// verify the required parameter 'placeholder' is set
	if &placeholder == nil {
		return new(ConnectorInstruction), nil, errors.New("Missing required parameter 'placeholder' when calling ConnectorsApi->V1ConnectorsConnectorConnectParameterGet")
	}
	// verify the required parameter 'type_' is set
	if &type_ == nil {
		return new(ConnectorInstruction), nil, errors.New("Missing required parameter 'type_' when calling ConnectorsApi->V1ConnectorsConnectorConnectParameterGet")
	}
	// verify the required parameter 'usePopup' is set
	if &usePopup == nil {
		return new(ConnectorInstruction), nil, errors.New("Missing required parameter 'usePopup' when calling ConnectorsApi->V1ConnectorsConnectorConnectParameterGet")
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
			queryParams.Add("defaultValue", a.Configuration.APIClient.ParameterToString(defaultValue, ""))
			queryParams.Add("displayName", a.Configuration.APIClient.ParameterToString(displayName, ""))
			queryParams.Add("key", a.Configuration.APIClient.ParameterToString(key, ""))
			queryParams.Add("placeholder", a.Configuration.APIClient.ParameterToString(placeholder, ""))
			queryParams.Add("type_", a.Configuration.APIClient.ParameterToString(type_, ""))
			queryParams.Add("usePopup", a.Configuration.APIClient.ParameterToString(usePopup, ""))
	

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
	var successPayload = new(ConnectorInstruction)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Delete stored connection info
 * The disconnect method deletes any stored tokens or connection information from the connectors database.
 *
 * @param connector Lowercase system name of the source application or device. Get a list of available connectors from the /connectors/list endpoint.
 * @return void
 */
func (a ConnectorsApi) V1ConnectorsConnectorDisconnectGet(connector string) (*APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/connectors/{connector}/disconnect"
	path = strings.Replace(path, "{"+"connector"+"}", fmt.Sprintf("%v", connector), -1)

	// verify the required parameter 'connector' is set
	if &connector == nil {
		return nil, errors.New("Missing required parameter 'connector' when calling ConnectorsApi->V1ConnectorsConnectorDisconnectGet")
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

	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return NewAPIResponse(httpResponse.RawResponse), err
	}

	return NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Get connector info for user
 * Returns information about the connector such as the connector id, whether or not is connected for this user (i.e. we have a token or credentials), and its update history for the user.
 *
 * @param connector Lowercase system name of the source application or device. Get a list of available connectors from the /connectors/list endpoint.
 * @param accessToken User&#39;s OAuth2 access token
 * @return *ConnectorInfo
 */
func (a ConnectorsApi) V1ConnectorsConnectorInfoGet(connector string, accessToken string) (*ConnectorInfo, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/connectors/{connector}/info"
	path = strings.Replace(path, "{"+"connector"+"}", fmt.Sprintf("%v", connector), -1)

	// verify the required parameter 'connector' is set
	if &connector == nil {
		return new(ConnectorInfo), nil, errors.New("Missing required parameter 'connector' when calling ConnectorsApi->V1ConnectorsConnectorInfoGet")
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
	var successPayload = new(ConnectorInfo)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Sync with data source
 * The update method tells the QM Connector Framework to check with the data provider (such as Fitbit or MyFitnessPal) and retrieve any new measurements available.
 *
 * @param connector Lowercase system name of the source application or device
 * @param accessToken User&#39;s OAuth2 access token
 * @return void
 */
func (a ConnectorsApi) V1ConnectorsConnectorUpdateGet(connector string, accessToken string) (*APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/connectors/{connector}/update"
	path = strings.Replace(path, "{"+"connector"+"}", fmt.Sprintf("%v", connector), -1)

	// verify the required parameter 'connector' is set
	if &connector == nil {
		return nil, errors.New("Missing required parameter 'connector' when calling ConnectorsApi->V1ConnectorsConnectorUpdateGet")
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

	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return NewAPIResponse(httpResponse.RawResponse), err
	}

	return NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * List of Connectors
 * A connector pulls data from other data providers using their API or a screenscraper. Returns a list of all available connectors and information about them such as their id, name, whether the user has provided access, logo url, connection instructions, and the update history.
 *
 * @return []Connector
 */
func (a ConnectorsApi) V1ConnectorsListGet() ([]Connector, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/connectors/list"


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
	var successPayload = new([]Connector)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
}


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

type PairsApi struct {
	Configuration Configuration
}

func NewPairsApi() *PairsApi {
	configuration := NewConfiguration()
	return &PairsApi{
		Configuration: *configuration,
	}
}

func NewPairsApiWithBasePath(basePath string) *PairsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &PairsApi{
		Configuration: *configuration,
	}
}

/**
 * Get pairs
 * Pairs cause measurements with effect measurements grouped over the duration of action after the onset delay.
 *
 * @param cause Original variable name for the explanatory or independent variable
 * @param effect Original variable name for the outcome or dependent variable
 * @param accessToken User&#39;s OAuth2 access token
 * @param causeSource Name of data source that the cause measurements should come from
 * @param causeUnit Abbreviated name for the unit cause measurements to be returned in
 * @param delay Delay before onset of action (in seconds) from the cause variable settings.
 * @param duration Duration of action (in seconds) from the cause variable settings.
 * @param effectSource Name of data source that the effectmeasurements should come from
 * @param effectUnit Abbreviated name for the unit effect measurements to be returned in
 * @param endTime The most recent date (in epoch time) for which we should return measurements
 * @param startTime The earliest date (in epoch time) for which we should return measurements
 * @param limit The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0.
 * @param offset Now suppose you wanted to show results 11-20. You&#39;d set the offset to 10 and the limit to 10.
 * @param sort Sort by given field. If the field is prefixed with &#x60;-, it will sort in descending order.
 * @return []Pairs
 */
func (a PairsApi) V1PairsCsvGet(cause string, effect string, accessToken string, causeSource string, causeUnit string, delay string, duration string, effectSource string, effectUnit string, endTime string, startTime string, limit int32, offset int32, sort int32) ([]Pairs, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/pairsCsv"

	// verify the required parameter 'cause' is set
	if &cause == nil {
		return new([]Pairs), nil, errors.New("Missing required parameter 'cause' when calling PairsApi->V1PairsCsvGet")
	}
	// verify the required parameter 'effect' is set
	if &effect == nil {
		return new([]Pairs), nil, errors.New("Missing required parameter 'effect' when calling PairsApi->V1PairsCsvGet")
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
			queryParams.Add("cause", a.Configuration.APIClient.ParameterToString(cause, ""))
			queryParams.Add("causeSource", a.Configuration.APIClient.ParameterToString(causeSource, ""))
			queryParams.Add("causeUnit", a.Configuration.APIClient.ParameterToString(causeUnit, ""))
			queryParams.Add("delay", a.Configuration.APIClient.ParameterToString(delay, ""))
			queryParams.Add("duration", a.Configuration.APIClient.ParameterToString(duration, ""))
			queryParams.Add("effect", a.Configuration.APIClient.ParameterToString(effect, ""))
			queryParams.Add("effectSource", a.Configuration.APIClient.ParameterToString(effectSource, ""))
			queryParams.Add("effectUnit", a.Configuration.APIClient.ParameterToString(effectUnit, ""))
			queryParams.Add("endTime", a.Configuration.APIClient.ParameterToString(endTime, ""))
			queryParams.Add("startTime", a.Configuration.APIClient.ParameterToString(startTime, ""))
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
	var successPayload = new([]Pairs)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Get pairs
 * Pairs cause measurements with effect measurements grouped over the duration of action after the onset delay.
 *
 * @param cause Original variable name for the explanatory or independent variable
 * @param effect Original variable name for the outcome or dependent variable
 * @param accessToken User&#39;s OAuth2 access token
 * @param causeSource Name of data source that the cause measurements should come from
 * @param causeUnit Abbreviated name for the unit cause measurements to be returned in
 * @param delay Delay before onset of action (in seconds) from the cause variable settings.
 * @param duration Duration of action (in seconds) from the cause variable settings.
 * @param effectSource Name of data source that the effectmeasurements should come from
 * @param effectUnit Abbreviated name for the unit effect measurements to be returned in
 * @param endTime The most recent date (in epoch time) for which we should return measurements
 * @param startTime The earliest date (in epoch time) for which we should return measurements
 * @param limit The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0.
 * @param offset Now suppose you wanted to show results 11-20. You&#39;d set the offset to 10 and the limit to 10.
 * @param sort Sort by given field. If the field is prefixed with &#x60;-, it will sort in descending order.
 * @return []Pairs
 */
func (a PairsApi) V1PairsGet(cause string, effect string, accessToken string, causeSource string, causeUnit string, delay string, duration string, effectSource string, effectUnit string, endTime string, startTime string, limit int32, offset int32, sort int32) ([]Pairs, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/v1/pairs"

	// verify the required parameter 'cause' is set
	if &cause == nil {
		return new([]Pairs), nil, errors.New("Missing required parameter 'cause' when calling PairsApi->V1PairsGet")
	}
	// verify the required parameter 'effect' is set
	if &effect == nil {
		return new([]Pairs), nil, errors.New("Missing required parameter 'effect' when calling PairsApi->V1PairsGet")
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
			queryParams.Add("cause", a.Configuration.APIClient.ParameterToString(cause, ""))
			queryParams.Add("causeSource", a.Configuration.APIClient.ParameterToString(causeSource, ""))
			queryParams.Add("causeUnit", a.Configuration.APIClient.ParameterToString(causeUnit, ""))
			queryParams.Add("delay", a.Configuration.APIClient.ParameterToString(delay, ""))
			queryParams.Add("duration", a.Configuration.APIClient.ParameterToString(duration, ""))
			queryParams.Add("effect", a.Configuration.APIClient.ParameterToString(effect, ""))
			queryParams.Add("effectSource", a.Configuration.APIClient.ParameterToString(effectSource, ""))
			queryParams.Add("effectUnit", a.Configuration.APIClient.ParameterToString(effectUnit, ""))
			queryParams.Add("endTime", a.Configuration.APIClient.ParameterToString(endTime, ""))
			queryParams.Add("startTime", a.Configuration.APIClient.ParameterToString(startTime, ""))
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
	var successPayload = new([]Pairs)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
}


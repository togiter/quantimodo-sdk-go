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

type Variable struct {

	// Variable ID
	Id int32 `json:"id,omitempty"`

	// User-defined variable display name.
	Name string `json:"name,omitempty"`

	// Name used when the variable was originally created in the `variables` table.
	OriginalName string `json:"originalName,omitempty"`

	// Variable category like Mood, Sleep, Physical Activity, Treatment, Symptom, etc.
	Category string `json:"category,omitempty"`

	// Abbreviated name of the default unit for the variable
	AbbreviatedUnitName string `json:"abbreviatedUnitName,omitempty"`

	// Id of the default unit for the variable
	AbbreviatedUnitId int32 `json:"abbreviatedUnitId,omitempty"`

	// Comma-separated list of source names to limit variables to those sources
	Sources string `json:"sources,omitempty"`

	// Minimum reasonable value for this variable (uses default unit)
	MinimumValue float64 `json:"minimumValue,omitempty"`

	// Maximum reasonable value for this variable (uses default unit)
	MaximumValue float64 `json:"maximumValue,omitempty"`

	// Way to aggregate measurements over time. Options are \"MEAN\" or \"SUM\".  SUM should be used for things like minutes of exercise.  If you use MEAN for exercise, then a person might exercise more minutes in one day but add separate measurements that were smaller.  So when we are doing correlational analysis, we would think that the person exercised less that day even though they exercised more.  Conversely, we must use MEAN for things such as ratings which cannot be SUMMED.
	CombinationOperation string `json:"combinationOperation,omitempty"`

	// Value for replacing null measurements
	FillingValue float64 `json:"fillingValue,omitempty"`

	// The Variable this Variable should be joined with. If the variable is joined with some other variable then it is not shown to user in the list of variables.
	JoinWith string `json:"joinWith,omitempty"`

	// Array of Variables that are joined with this Variable
	JoinedVariables []Variable `json:"joinedVariables,omitempty"`

	// Id of the parent variable if this variable has any parent
	Parent int32 `json:"parent,omitempty"`

	// Array of Variables that are sub variables to this Variable
	SubVariables []Variable `json:"subVariables,omitempty"`

	// How long it takes for a measurement in this variable to take effect
	OnsetDelay int32 `json:"onsetDelay,omitempty"`

	// How long the effect of a measurement in this variable lasts
	DurationOfAction int32 `json:"durationOfAction,omitempty"`

	// Earliest measurement time
	EarliestMeasurementTime int32 `json:"earliestMeasurementTime,omitempty"`

	// Latest measurement time
	LatestMeasurementTime int32 `json:"latestMeasurementTime,omitempty"`

	// When this variable or its settings were last updated
	Updated int32 `json:"updated,omitempty"`

	// A value of 1 indicates that this variable is generally a cause in a causal relationship.  An example of a causeOnly variable would be a variable such as Cloud Cover which would generally not be influenced by the behaviour of the user.
	CauseOnly int32 `json:"causeOnly,omitempty"`

	// Number of correlations
	NumberOfCorrelations int32 `json:"numberOfCorrelations,omitempty"`

	// Outcome variables (those with `outcome` == 1) are variables for which a human would generally want to identify the influencing factors.  These include symptoms of illness, physique, mood, cognitive performance, etc.  Generally correlation calculations are only performed on outcome variables.
	Outcome int32 `json:"outcome,omitempty"`

	// The number of measurements that a given user had for this variable the last time a correlation calculation was performed. Generally correlation values are only updated once the current number of measurements for a variable is more than 10% greater than the measurementsAtLastAnalysis.  This avoids a computationally-demanding recalculation when there's not enough new data to make a significant difference in the correlation.
	MeasurementsAtLastAnalysis int32 `json:"measurementsAtLastAnalysis,omitempty"`

	// Number of measurements
	NumberOfMeasurements int32 `json:"numberOfMeasurements,omitempty"`

	// Last unit
	LastUnit string `json:"lastUnit,omitempty"`

	// Last value
	LastValue int32 `json:"lastValue,omitempty"`

	// Most common value
	MostCommonValue int32 `json:"mostCommonValue,omitempty"`

	// Most common unit
	MostCommonUnit string `json:"mostCommonUnit,omitempty"`

	// Last source
	LastSource int32 `json:"lastSource,omitempty"`
}

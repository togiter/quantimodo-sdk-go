# Go API client for swagger

Welcome to QuantiModo API! QuantiModo makes it easy to retrieve normalized user data from a wide array of devices and applications. [Learn about QuantiModo](https://quantimo.do) or contact us at <api@quantimo.do>.         Before you get started, you will need to: * Sign in/Sign up, and add some data at [https://app.quantimo.do/api/v2/account/connectors](https://app.quantimo.do/api/v2/account/connectors) to try out the API for yourself * Create an app to get your client id and secret at [https://app.quantimo.do/api/v2/apps](https://app.quantimo.do/api/v2/apps) * As long as you're signed in, it will use your browser's cookie for authentication.  However, client applications must use OAuth2 tokens to access the API.     ## Application Endpoints These endpoints give you access to all authorized users' data for that application. ### Getting Application Token Make a `POST` request to `/api/v2/oauth/access_token`         * `grant_type` Must be `client_credentials`.         * `clientId` Your application's clientId.         * `client_secret` Your application's client_secret.         * `redirect_uri` Your application's redirect url.                ## Example Queries ### Query Options The standard query options for QuantiModo API are as described in the table below. These are the available query options in QuantiModo API: <table>            <thead>                <tr>                    <th>Parameter</th>                    <th>Description</th>                </tr>            </thead>            <tbody>                <tr>                    <td>limit</td>                    <td>The LIMIT is used to limit the number of results returned.  So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records.</td>                </tr>                <tr>                    <td>offset</td>                    <td>Suppose you wanted to show results 11-20. You'd set the    offset to 10 and the limit to 10.</td>                </tr>                <tr>                    <td>sort</td>                    <td>Sort by given field. If the field is prefixed with '-', it    will sort in descending order.</td>                </tr>            </tbody>        </table>         ### Pagination Conventions Since the maximum limit is 200 records, to get more than that you'll have to make multiple API calls and page through the results. To retrieve all the data, you can iterate through data by using the `limit` and `offset` query parameters.For example, if you want to retrieve data from 61-80 then you can use a query with the following parameters,         `/v2/variables?limit=20&offset=60`         Generally, you'll be retrieving new or updated user data. To avoid unnecessary API calls, you'll want to store your last refresh time locally.  Initially, it should be set to 0. Then whenever you make a request to get new data, you should limit the returned results to those updated since your last refresh by appending append         `?lastUpdated=(ge)&quot2013-01-D01T01:01:01&quot`         to your request.         Also for better pagination, you can get link to the records of first, last, next and previous page from response headers: * `Total-Count` - Total number of results for given query * `Link-First` - Link to get first page records * `Link-Last` - Link to get last page records * `Link-Prev` - Link to get previous records set * `Link-Next` - Link to get next records set         Remember, response header will be only sent when the record set is available. e.g. You will not get a ```Link-Last``` & ```Link-Next``` when you query for the last page.         ### Filter operators support API supports the following operators with filter parameters: <br> **Comparison operators**         Comparison operators allow you to limit results to those greater than, less than, or equal to a specified value for a specified attribute. These operators can be used with strings, numbers, and dates. The following comparison operators are available: * `gt` for `greater than` comparison * `ge` for `greater than or equal` comparison * `lt` for `less than` comparison * `le` for `less than or equal` comparison         They are included in queries using the following format:         `(<operator>)<value>`         For example, in order to filter value which is greater than 21, the following query parameter should be used:         `?value=(gt)21` <br><br> **Equals/In Operators**         It also allows filtering by the exact value of an attribute or by a set of values, depending on the type of value passed as a query parameter. If the value contains commas, the parameter is split on commas and used as array input for `IN` filtering, otherwise the exact match is applied. In order to only show records which have the value 42, the following query should be used:         `?value=42`         In order to filter records which have value 42 or 43, the following query should be used:         `?value=42,43` <br><br> **Like operators**         Like operators allow filtering using `LIKE` query. This operator is triggered if exact match operator is used, but value contains `%` sign as the first or last character. In order to filter records which category that start with `Food`, the following query should be used:         `?category=Food%` <br><br> **Negation operator**         It is possible to get negated results of a query by prefixed the operator with `!`. Some examples:         `//filter records except those with value are not 42 or 43`<br> `?value=!42,43`         `//filter records with value not greater than 21`<br> `?value=!(ge)21` <br><br> **Multiple constraints for single attribute**         It is possible to apply multiple constraints by providing an array of query filters:         Filter all records which value is greater than 20.2 and less than 20.3<br> `?value[]=(gt)20.2&value[]=(lt)20.3`         Filter all records which value is greater than 20.2 and less than 20.3 but not 20.2778<br> `?value[]=(gt)20.2&value[]=(lt)20.3&value[]=!20.2778`<br><br> 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.0.6
- Package version: 1.0.0
- Build date: 2016-07-27T03:54:29.559Z
- Build package: class io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://localhost/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApplicationEndpointsApi* | [**V2ApplicationConnectionsGet**](docs/ApplicationEndpointsApi.md#v2applicationconnectionsget) | **Get** /v2/application/connections | Get all Connections
*ApplicationEndpointsApi* | [**V2ApplicationCredentialsGet**](docs/ApplicationEndpointsApi.md#v2applicationcredentialsget) | **Get** /v2/application/credentials | Get all Credentials
*ApplicationEndpointsApi* | [**V2ApplicationMeasurementsGet**](docs/ApplicationEndpointsApi.md#v2applicationmeasurementsget) | **Get** /v2/application/measurements | Get measurements for all users using your application
*ApplicationEndpointsApi* | [**V2ApplicationTrackingRemindersGet**](docs/ApplicationEndpointsApi.md#v2applicationtrackingremindersget) | **Get** /v2/application/trackingReminders | Get tracking reminders
*ApplicationEndpointsApi* | [**V2ApplicationUpdatesGet**](docs/ApplicationEndpointsApi.md#v2applicationupdatesget) | **Get** /v2/application/updates | Get all Updates
*ApplicationEndpointsApi* | [**V2ApplicationUserVariableRelationshipsGet**](docs/ApplicationEndpointsApi.md#v2applicationuservariablerelationshipsget) | **Get** /v2/application/userVariableRelationships | Get all UserVariableRelationships
*ApplicationEndpointsApi* | [**V2ApplicationUserVariablesGet**](docs/ApplicationEndpointsApi.md#v2applicationuservariablesget) | **Get** /v2/application/userVariables | Get all UserVariables
*ApplicationEndpointsApi* | [**V2ApplicationVariableUserSourcesGet**](docs/ApplicationEndpointsApi.md#v2applicationvariableusersourcesget) | **Get** /v2/application/variableUserSources | Get all VariableUserSources
*ApplicationEndpointsApi* | [**V2ApplicationVotesGet**](docs/ApplicationEndpointsApi.md#v2applicationvotesget) | **Get** /v2/application/votes | Get all Votes
*AuthenticationApi* | [**V2AuthSocialAuthorizeCodeGet**](docs/AuthenticationApi.md#v2authsocialauthorizecodeget) | **Get** /v2/auth/social/authorizeCode | Second Step in Social Authentication flow with JWT Token
*AuthenticationApi* | [**V2AuthSocialAuthorizeTokenGet**](docs/AuthenticationApi.md#v2authsocialauthorizetokenget) | **Get** /v2/auth/social/authorizeToken | Native Social Authentication
*AuthenticationApi* | [**V2AuthSocialLoginGet**](docs/AuthenticationApi.md#v2authsocialloginget) | **Get** /v2/auth/social/login | First Setp in Social Authentication flow with JWT Token
*AuthenticationApi* | [**V2Oauth2AccessTokenGet**](docs/AuthenticationApi.md#v2oauth2accesstokenget) | **Get** /v2/oauth2/access_token | Get a user access token
*AuthenticationApi* | [**V2OauthAuthorizeGet**](docs/AuthenticationApi.md#v2oauthauthorizeget) | **Get** /v2/oauth/authorize | Request Authorization Code
*ConnectorsApi* | [**V1ConnectJsGet**](docs/ConnectorsApi.md#v1connectjsget) | **Get** /v1/connect.js | Get embeddable connect javascript
*ConnectorsApi* | [**V1ConnectMobileGet**](docs/ConnectorsApi.md#v1connectmobileget) | **Get** /v1/connect/mobile | Mobile connect page
*ConnectorsApi* | [**V1ConnectorsConnectorConnectGet**](docs/ConnectorsApi.md#v1connectorsconnectorconnectget) | **Get** /v1/connectors/{connector}/connect | Obtain a token from 3rd party data source
*ConnectorsApi* | [**V1ConnectorsConnectorConnectInstructionsGet**](docs/ConnectorsApi.md#v1connectorsconnectorconnectinstructionsget) | **Get** /v1/connectors/{connector}/connectInstructions | Connection Instructions
*ConnectorsApi* | [**V1ConnectorsConnectorConnectParameterGet**](docs/ConnectorsApi.md#v1connectorsconnectorconnectparameterget) | **Get** /v1/connectors/{connector}/connectParameter | Connect Parameter
*ConnectorsApi* | [**V1ConnectorsConnectorDisconnectGet**](docs/ConnectorsApi.md#v1connectorsconnectordisconnectget) | **Get** /v1/connectors/{connector}/disconnect | Delete stored connection info
*ConnectorsApi* | [**V1ConnectorsConnectorInfoGet**](docs/ConnectorsApi.md#v1connectorsconnectorinfoget) | **Get** /v1/connectors/{connector}/info | Get connector info for user
*ConnectorsApi* | [**V1ConnectorsConnectorUpdateGet**](docs/ConnectorsApi.md#v1connectorsconnectorupdateget) | **Get** /v1/connectors/{connector}/update | Sync with data source
*ConnectorsApi* | [**V1ConnectorsListGet**](docs/ConnectorsApi.md#v1connectorslistget) | **Get** /v1/connectors/list | List of Connectors
*CorrelationsApi* | [**V1AggregatedCorrelationsGet**](docs/CorrelationsApi.md#v1aggregatedcorrelationsget) | **Get** /v1/aggregatedCorrelations | Get aggregated correlations
*CorrelationsApi* | [**V1AggregatedCorrelationsPost**](docs/CorrelationsApi.md#v1aggregatedcorrelationspost) | **Post** /v1/aggregatedCorrelations | Store or Update a Correlation
*CorrelationsApi* | [**V1CorrelationsGet**](docs/CorrelationsApi.md#v1correlationsget) | **Get** /v1/correlations | Get correlations
*CorrelationsApi* | [**V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameCausesGet**](docs/CorrelationsApi.md#v1organizationsorganizationidusersuseridvariablesvariablenamecausesget) | **Get** /v1/organizations/{organizationId}/users/{userId}/variables/{variableName}/causes | Search user correlations for a given cause
*CorrelationsApi* | [**V1OrganizationsOrganizationIdUsersUserIdVariablesVariableNameEffectsGet**](docs/CorrelationsApi.md#v1organizationsorganizationidusersuseridvariablesvariablenameeffectsget) | **Get** /v1/organizations/{organizationId}/users/{userId}/variables/{variableName}/effects | Search user correlations for a given cause
*CorrelationsApi* | [**V1PublicCorrelationsSearchSearchGet**](docs/CorrelationsApi.md#v1publiccorrelationssearchsearchget) | **Get** /v1/public/correlations/search/{search} | Get average correlations for variables containing search term
*CorrelationsApi* | [**V1VariablesVariableNameCausesGet**](docs/CorrelationsApi.md#v1variablesvariablenamecausesget) | **Get** /v1/variables/{variableName}/causes | Search user correlations for a given effect
*CorrelationsApi* | [**V1VariablesVariableNameEffectsGet**](docs/CorrelationsApi.md#v1variablesvariablenameeffectsget) | **Get** /v1/variables/{variableName}/effects | Search user correlations for a given cause
*CorrelationsApi* | [**V1VariablesVariableNamePublicCausesGet**](docs/CorrelationsApi.md#v1variablesvariablenamepubliccausesget) | **Get** /v1/variables/{variableName}/public/causes | Search public correlations for a given effect
*CorrelationsApi* | [**V1VariablesVariableNamePublicEffectsGet**](docs/CorrelationsApi.md#v1variablesvariablenamepubliceffectsget) | **Get** /v1/variables/{variableName}/public/effects | Search public correlations for a given cause
*CorrelationsApi* | [**V1VotesDeletePost**](docs/CorrelationsApi.md#v1votesdeletepost) | **Post** /v1/votes/delete | Delete vote
*CorrelationsApi* | [**V1VotesPost**](docs/CorrelationsApi.md#v1votespost) | **Post** /v1/votes | Post or update vote
*MeasurementsApi* | [**V1MeasurementSourcesGet**](docs/MeasurementsApi.md#v1measurementsourcesget) | **Get** /v1/measurementSources | Get measurement sources
*MeasurementsApi* | [**V1MeasurementSourcesPost**](docs/MeasurementsApi.md#v1measurementsourcespost) | **Post** /v1/measurementSources | Add a data source
*MeasurementsApi* | [**V1MeasurementsDailyGet**](docs/MeasurementsApi.md#v1measurementsdailyget) | **Get** /v1/measurements/daily | Get daily measurements for this user
*MeasurementsApi* | [**V1MeasurementsDeletePost**](docs/MeasurementsApi.md#v1measurementsdeletepost) | **Post** /v1/measurements/delete | Delete a measurement
*MeasurementsApi* | [**V1MeasurementsGet**](docs/MeasurementsApi.md#v1measurementsget) | **Get** /v1/measurements | Get measurements for this user
*MeasurementsApi* | [**V1MeasurementsPost**](docs/MeasurementsApi.md#v1measurementspost) | **Post** /v1/measurements | Post a new set or update existing measurements to the database
*MeasurementsApi* | [**V1MeasurementsRangeGet**](docs/MeasurementsApi.md#v1measurementsrangeget) | **Get** /v1/measurementsRange | Get measurements range for this user
*MeasurementsApi* | [**V2MeasurementsCsvGet**](docs/MeasurementsApi.md#v2measurementscsvget) | **Get** /v2/measurements/csv | Get Measurements CSV
*MeasurementsApi* | [**V2MeasurementsIdDelete**](docs/MeasurementsApi.md#v2measurementsiddelete) | **Delete** /v2/measurements/{id} | Delete Measurement
*MeasurementsApi* | [**V2MeasurementsIdGet**](docs/MeasurementsApi.md#v2measurementsidget) | **Get** /v2/measurements/{id} | Get Measurement
*MeasurementsApi* | [**V2MeasurementsIdPut**](docs/MeasurementsApi.md#v2measurementsidput) | **Put** /v2/measurements/{id} | Update Measurement
*MeasurementsApi* | [**V2MeasurementsRequestCsvPost**](docs/MeasurementsApi.md#v2measurementsrequestcsvpost) | **Post** /v2/measurements/request_csv | Post Request for Measurements CSV
*MeasurementsApi* | [**V2MeasurementsRequestPdfPost**](docs/MeasurementsApi.md#v2measurementsrequestpdfpost) | **Post** /v2/measurements/request_pdf | Post Request for Measurements PDF
*MeasurementsApi* | [**V2MeasurementsRequestXlsPost**](docs/MeasurementsApi.md#v2measurementsrequestxlspost) | **Post** /v2/measurements/request_xls | Post Request for Measurements XLS
*OrganizationsApi* | [**V1OrganizationsOrganizationIdUsersPost**](docs/OrganizationsApi.md#v1organizationsorganizationiduserspost) | **Post** /v1/organizations/{organizationId}/users | Get user tokens for existing users, create new users
*PairsApi* | [**V1PairsCsvGet**](docs/PairsApi.md#v1pairscsvget) | **Get** /v1/pairsCsv | Get pairs
*PairsApi* | [**V1PairsGet**](docs/PairsApi.md#v1pairsget) | **Get** /v1/pairs | Get pairs
*RemindersApi* | [**V1TrackingReminderNotificationsGet**](docs/RemindersApi.md#v1trackingremindernotificationsget) | **Get** /v1/trackingReminderNotifications | Get specific pending tracking reminders
*RemindersApi* | [**V1TrackingReminderNotificationsSkipPost**](docs/RemindersApi.md#v1trackingremindernotificationsskippost) | **Post** /v1/trackingReminderNotifications/skip | Skip a pending tracking reminder
*RemindersApi* | [**V1TrackingReminderNotificationsSnoozePost**](docs/RemindersApi.md#v1trackingremindernotificationssnoozepost) | **Post** /v1/trackingReminderNotifications/snooze | Snooze a pending tracking reminder
*RemindersApi* | [**V1TrackingReminderNotificationsTrackPost**](docs/RemindersApi.md#v1trackingremindernotificationstrackpost) | **Post** /v1/trackingReminderNotifications/track | Track a pending tracking reminder
*RemindersApi* | [**V1TrackingRemindersDeletePost**](docs/RemindersApi.md#v1trackingremindersdeletepost) | **Post** /v1/trackingReminders/delete | Delete tracking reminder
*RemindersApi* | [**V1TrackingRemindersGet**](docs/RemindersApi.md#v1trackingremindersget) | **Get** /v1/trackingReminders | Get repeating tracking reminder settings
*RemindersApi* | [**V1TrackingRemindersPost**](docs/RemindersApi.md#v1trackingreminderspost) | **Post** /v1/trackingReminders | Store a Tracking Reminder
*TagsApi* | [**V1UserTagsDeletePost**](docs/TagsApi.md#v1usertagsdeletepost) | **Post** /v1/userTags/delete | Delete user tag or ingredient
*TagsApi* | [**V1UserTagsPost**](docs/TagsApi.md#v1usertagspost) | **Post** /v1/userTags | Post or update user tags or ingredients
*UnitsApi* | [**V1UnitCategoriesGet**](docs/UnitsApi.md#v1unitcategoriesget) | **Get** /v1/unitCategories | Get unit categories
*UnitsApi* | [**V1UnitsGet**](docs/UnitsApi.md#v1unitsget) | **Get** /v1/units | Get all available units
*UnitsApi* | [**V1UnitsVariableGet**](docs/UnitsApi.md#v1unitsvariableget) | **Get** /v1/unitsVariable | Units for Variable
*UserApi* | [**V1OrganizationsOrganizationIdUsersPost**](docs/UserApi.md#v1organizationsorganizationiduserspost) | **Post** /v1/organizations/{organizationId}/users | Get user tokens for existing users, create new users
*UserApi* | [**V1UserMeGet**](docs/UserApi.md#v1usermeget) | **Get** /v1/user/me | Get all available units for variableGet authenticated user
*VariablesApi* | [**V1PublicVariablesGet**](docs/VariablesApi.md#v1publicvariablesget) | **Get** /v1/public/variables | Get public variables
*VariablesApi* | [**V1PublicVariablesSearchSearchGet**](docs/VariablesApi.md#v1publicvariablessearchsearchget) | **Get** /v1/public/variables/search/{search} | Get top 5 PUBLIC variables with the most correlations
*VariablesApi* | [**V1UserVariablesPost**](docs/VariablesApi.md#v1uservariablespost) | **Post** /v1/userVariables | Update User Settings for a Variable
*VariablesApi* | [**V1VariableCategoriesGet**](docs/VariablesApi.md#v1variablecategoriesget) | **Get** /v1/variableCategories | Variable categories
*VariablesApi* | [**V1VariablesGet**](docs/VariablesApi.md#v1variablesget) | **Get** /v1/variables | Get variables with user&#39;s settings
*VariablesApi* | [**V1VariablesPost**](docs/VariablesApi.md#v1variablespost) | **Post** /v1/variables | Create Variables
*VariablesApi* | [**V1VariablesSearchSearchGet**](docs/VariablesApi.md#v1variablessearchsearchget) | **Get** /v1/variables/search/{search} | Get variables by search query
*VariablesApi* | [**V1VariablesVariableNameGet**](docs/VariablesApi.md#v1variablesvariablenameget) | **Get** /v1/variables/{variableName} | Get info about a variable
*VotesApi* | [**V1VotesDeletePost**](docs/VotesApi.md#v1votesdeletepost) | **Post** /v1/votes/delete | Delete vote
*VotesApi* | [**V1VotesPost**](docs/VotesApi.md#v1votespost) | **Post** /v1/votes | Post or update vote


## Documentation For Models

 - [CommonResponse](docs/CommonResponse.md)
 - [Connection](docs/Connection.md)
 - [Connector](docs/Connector.md)
 - [ConnectorInfo](docs/ConnectorInfo.md)
 - [ConnectorInfoHistoryItem](docs/ConnectorInfoHistoryItem.md)
 - [ConnectorInstruction](docs/ConnectorInstruction.md)
 - [ConversionStep](docs/ConversionStep.md)
 - [Correlation](docs/Correlation.md)
 - [Credential](docs/Credential.md)
 - [HumanTime](docs/HumanTime.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse20010](docs/InlineResponse20010.md)
 - [InlineResponse20011](docs/InlineResponse20011.md)
 - [InlineResponse20012](docs/InlineResponse20012.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [InlineResponse2007](docs/InlineResponse2007.md)
 - [InlineResponse2008](docs/InlineResponse2008.md)
 - [InlineResponse2009](docs/InlineResponse2009.md)
 - [JsonErrorResponse](docs/JsonErrorResponse.md)
 - [Measurement](docs/Measurement.md)
 - [MeasurementDelete](docs/MeasurementDelete.md)
 - [MeasurementRange](docs/MeasurementRange.md)
 - [MeasurementSet](docs/MeasurementSet.md)
 - [MeasurementSource](docs/MeasurementSource.md)
 - [Pairs](docs/Pairs.md)
 - [Permission](docs/Permission.md)
 - [PostCorrelation](docs/PostCorrelation.md)
 - [PostVote](docs/PostVote.md)
 - [TrackingReminder](docs/TrackingReminder.md)
 - [TrackingReminderDelete](docs/TrackingReminderDelete.md)
 - [TrackingReminderNotification](docs/TrackingReminderNotification.md)
 - [TrackingReminderNotificationSkip](docs/TrackingReminderNotificationSkip.md)
 - [TrackingReminderNotificationSnooze](docs/TrackingReminderNotificationSnooze.md)
 - [TrackingReminderNotificationTrack](docs/TrackingReminderNotificationTrack.md)
 - [Unit](docs/Unit.md)
 - [UnitCategory](docs/UnitCategory.md)
 - [Update](docs/Update.md)
 - [User](docs/User.md)
 - [UserTag](docs/UserTag.md)
 - [UserTokenFailedResponse](docs/UserTokenFailedResponse.md)
 - [UserTokenRequest](docs/UserTokenRequest.md)
 - [UserTokenRequestInnerUserField](docs/UserTokenRequestInnerUserField.md)
 - [UserTokenSuccessfulResponse](docs/UserTokenSuccessfulResponse.md)
 - [UserTokenSuccessfulResponseInnerUserField](docs/UserTokenSuccessfulResponseInnerUserField.md)
 - [UserVariable](docs/UserVariable.md)
 - [UserVariableRelationship](docs/UserVariableRelationship.md)
 - [UserVariables](docs/UserVariables.md)
 - [ValueObject](docs/ValueObject.md)
 - [Variable](docs/Variable.md)
 - [VariableCategory](docs/VariableCategory.md)
 - [VariableNew](docs/VariableNew.md)
 - [VariableUserSource](docs/VariableUserSource.md)
 - [VariablesNew](docs/VariablesNew.md)
 - [Vote](docs/Vote.md)
 - [VoteDelete](docs/VoteDelete.md)


## Documentation For Authorization


## oauth2

- **Type**: OAuth
- **Flow**: implicit
- **Authorizatoin URL**: https://app.quantimo.do/api/v1/oauth2/authorize
- **Scopes**: 
 - **basic**: Basic authentication
 - **readmeasurements**: Grants read access to measurements and variables. Allows the client app to obtain the user's data.
 - **writemeasurements**: Grants write access to measurements and variables. Allows the client app to store user data.

## quantimodo_oauth2

- **Type**: OAuth
- **Flow**: accessCode
- **Authorizatoin URL**: /api/v2/oauth/authorize
- **Scopes**: 
 - **basic**: allows you to read user info (displayname, email, etc).
 - **readmeasurements**: allows one to read a user's data
 - **writemeasurements**: allows you to write user data

## basicAuth

- **Type**: HTTP basic authentication

## internalApiKey

- **Type**: API key 
- **API key parameter name**: api_key
- **Location**: HTTP header


## Author




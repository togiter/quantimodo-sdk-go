# Connection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | id | [optional] [default to null]
**UserId** | **int32** | ID of user that owns this correlation | [optional] [default to null]
**ConnectorId** | **int32** | The id for the connector data source for which the connection is connected | [default to null]
**ConnectStatus** | **string** | Indicates whether a connector is currently connected to a service for a user. | [optional] [default to null]
**ConnectError** | **string** | Error message if there is a problem with authorizing this connection. | [optional] [default to null]
**UpdateRequestedAt** | [**time.Time**](time.Time.md) | Time at which an update was requested by a user. | [optional] [default to null]
**UpdateStatus** | **string** | Indicates whether a connector is currently updated. | [optional] [default to null]
**UpdateError** | **string** | Indicates if there was an error during the update. | [optional] [default to null]
**LastSuccessfulUpdatedAt** | [**time.Time**](time.Time.md) | The time at which the connector was last successfully updated. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | When the record was first created. Use ISO 8601 datetime format | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | When the record in the database was last updated. Use ISO 8601 datetime format | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



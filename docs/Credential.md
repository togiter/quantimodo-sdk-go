# Credential

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | ID of user that owns this credential | [optional] [default to null]
**ConnectorId** | **int32** | The id for the connector data source from which the credential was obtained | [default to null]
**AttrKey** | **string** | Attribute name such as token, userid, username, or password | [default to null]
**AttrValue** | **string** | Encrypted value for the attribute specified | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | When the record was first created. Use ISO 8601 datetime format | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | When the record in the database was last updated. Use ISO 8601 datetime format | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



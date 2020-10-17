# API Endpoints


## GET /api/resolve/{id}
Possible Errors:
 * [Missing ID](#Missing-ID-Attribute)
 * [Invalid ID](#Invalid-ID)

This endpoint will get the Minecraft player UUID associated with a Discord ID given or vice-versa. If an ID isn't provided then "Missing ID" error is given, if the ID provided isn't found then "Invalid ID" error is given.

## Errors

### Missing ID Attribute
```json
{
    "errcode": "MISSING_ID",
    "message": "A resolvable ID is missing in the URL path."
}
```

### Invalid ID
```json
{
    "errcode": "INVALID_ID",
    "message": "The ID provided isn't in the database or is invalid."
}
```

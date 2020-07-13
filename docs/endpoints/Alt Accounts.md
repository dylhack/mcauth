# Alt Account Management
These endpoints are for enforcing the authenticator to allow players to join
. In particular administrator alt accounts.

## GET /getAltsOf/:owner
Possible Errors:
 * [Missing Owner Attribute](#Missing-Owner-Attribute)
 * [Invalid Owner Attribute Type](#Invalid-Owner-Attribute-Type)

### Request Body
Required Headers:
 1. Authorization: `Bearer <webserver token>` 

(empty) The owner attribute is provided in the URL path.
ie `GET http://127.0.0.1/getAltsOf/notch`

### Response Body
| Attribute | Type     | Description          |
|-----------|----------|----------------------|
| alt_accs  | AltAcc[] | An array of AltAcc's |

#### AltAcc typedef
| Attribute | Type   | Description                    |
|-----------|--------|--------------------------------|
| owner     | string | The person who claimed the alt |
| alt_name  | string | The Minecraft player name alt  |
| alt_id    | string | The Minecraft player UUID alt  |


## POST /newAlt
Possible Errors:
 * [Missing Owner Attribute](#Missing-Owner-Attribute)
 * [Invalid Owner Attribute Type](#Invalid-Owner-Attribute-Type)
 * [Invalid Owner](#Invalid-Owner)
 * [Missing Player Name Attribute](#Missing-Player-Name-Attribute)
 * [Invalid Player Name Type Attribute](#Invalid-Player-Name-Attribute-Type)
 * [Alt Already Added Error](#Alt-Already-Added-Error)

### Request Body
Required Headers:
 1. Content-Type: `application/json`
 2. Authorization: `Bearer <webserver token>` 

| Attribute   | Type   | Description                               |
|-------------|--------|-------------------------------------------|
| player_name | string | The name of the alt account being claimed |
| owner       | string | The owner of said alt account             |

### Response Body
Empty (200 OK)

## DELETE /delAlt
Possible Errors:
 * [Missing Player Name Attribute](#Missing-Player-Name-Attribute)
 * [Invalid Player Name Type Attribute](#Invalid-Player-Name-Attribute-Type)

### Request Body
Required Headers:
 1. Content-Type: `application/json`
 2. Authorization: `Bearer <webserver token>` 

| Attribute   | Type   | Description                               |
|-------------|--------|-------------------------------------------|
| player_name | string | The name of the alt account being removed |


### Response Body
| Attribute  | Type    | Description                                |
|------------|---------|--------------------------------------------|
| is_deleted | boolean | Whether or not it was successfully removed |


## Errors

#### Missing Player Name Attribute
```json
{
  "errcode": "NO_PLAYER_NAME",
  "message": "A player name wasn't provided"
}
```

#### Invalid Player Name Attribute Type
```json
{
  "errcode": "PLAYER_TYPE_ERROR",
  "message": "The player name attribute was not provided as a string"
}
```

#### Missing Owner Attribute
```json
{
  "errcode": "NO_OWNER",
  "message": "An owner attribute was not provided"
}
```

#### Invalid Owner Attribute Type
```json
{
  "errcode": "OWNER_TYPE_ERROR",
  "message": "The owner attribute provided is not a string type"
}
```

### Invalid Owner
```json
{
	"errcode": "INVALID_OWNER",
	"message": "The owner provided is not a valid player name"
}
```

### Alt Already Added Error
```json
{
  "errcode": "ALT_ALREADY_ADDED",
  "message": "The alt provided is already stored in the database"
}
```

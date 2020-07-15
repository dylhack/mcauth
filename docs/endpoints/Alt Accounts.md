# Alt Account Management
These endpoints are for enforcing the authenticator to allow players to join
. In particular administrator alt accounts.

## GET /alts/{owner}
Possible Errors:
 * [Missing Owner Attribute](#Missing-Owner-Attribute)

### Required Headers
 1. Authorization: `<webserver token>` 

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


## POST /alts/{owner}/{player name}
Add a new alt account associated with an owner

Possible Errors:
 * [Missing Owner Attribute](#Missing-Owner-Attribute)
 * [Invalid Owner](#Invalid-Owner)
 * [Missing Player Name Attribute](#Missing-Player-Name-Attribute)
 * [Invalid Player Name](#Invalid-Alt-Name)
 * [Alt Already Added Error](#Alt-Already-Added-Error)

### Required Headers
 1. Authorization: `<webserver token>` 

### Response Body
Empty (200 OK)

## DELETE /alts/{alt name}
Possible Errors:
 * [Missing Player Name Attribute](#Missing-Player-Name-Attribute)
 * [Invalid Player Name](#Invalid-Alt-Name)

### Required Headers
Required Headers:
 1. Authorization: `<webserver token>` 

### Response Body
Empty (200 OK)

## Errors

#### Missing Player Name Attribute
```json
{
  "errcode": "MISSING_ALT_NAME",
  "message": "An alt player name wasn't provided"
}
```

### Invalid Alt Name
```json
{
  "errcode": "INVALID_ALT_NAME",
  "message": "The alt account name provided is not a valid player name"
}
```

#### Missing Owner Attribute
```json
{
  "errcode": "MISSING_OWNER",
  "message": "An owner attribute was not provided"
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

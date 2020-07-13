## POST /isPlayerValid
Possible Errors:
 * [Missing Player UUID Attribute](#Missing-Player-UUID-Attribute)
 * [Invalid Player UUID Attribute Type](#Invalid-Player-UUID-Attribute-Type)

This endpoint checks if a player is allowed to join the Minecraft server.

### Request Body
Required Headers:
 1. Content-Type: `application/json`
 2. Authorization: `Bearer <webserver token>` 

| Attribute | Type   | Description             |
|-----------|--------|-------------------------|
| player_id | string | The Minecraft player ID |

The player_id is the Minecraft player UUID stripped of all the dashes. The
server will provide the following response if everything went alright, 
otherwise an error may occur.


### Response Body

#### Invalid Player - Response
This means they failed authentication

| Attribute | Type    | Description                                      |
|-----------|---------|--------------------------------------------------|
| valid     | boolean | Whether or not the given player is ready to play |
| reason    | string  | Possible reasons are described below             |

 - "no_link": The Minecraft player hasn't linked their Discord account.
 - "no_role": They fail to have the required roles on Discord to join the
  Minecraft server.
 - "maintenance": The bot is in maintenance mode meaning only admin's can join.
 - "banned": This person is banned from using the bot and auth server.
 - "auth_code": This code needs to authorize their linkage
 
 
 #### Invalid Player - Please Auth Response
 | Attribute | Type    | Description                                      |
 |-----------|---------|--------------------------------------------------|
 | valid     | boolean | Whether or not the given player is ready to play |
 | reason    | string  | Only is "auth_code"                              |
 | auth_code | string  | The auth code they must provide the Discord bot  |


#### Valid Player - Response
This means they can play on the Minecraft server.

| Attribute | Type    | Description                                      |
|-----------|---------|--------------------------------------------------|
| valid     | boolean | Whether or not the given player is ready to play |

The valid attribute is a boolean which represents whether the player can
play on the Minecraft server. This will always return a boolean whether or
not there was an issue getting the member associated with the provided
player ID.

An added "reason" attribute also exists. It will only be 'no_link' which
means the Minecraft player isn't linked with a Discord account and 'no_role'
which means they're not whitelisted 

An operator of the Minecraft server can enforce validation of a player as
well using the [alts endpoint](./Alt%20Accounts.md).

## Errors

### Missing Player UUID Attribute
```json
{
  "errcode": "NO_PLAYER_ID",
  "message": "There wasn't a player ID provided"
}
```

### Invalid Player UUID Attribute Type
```json
{
  "errcode": "PLAYER_ID_TYPE",
  "message": "The provided player ID wasn't a string"
}
```

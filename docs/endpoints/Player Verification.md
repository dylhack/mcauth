# Player Verification
This is where the magic happens. In the database Minecraft player UUID's (without hyphens) and
Discord user ID's (twitter snowflakes) are stored together on the same row. This is called a 
linked account. This one endpoint allows external clients to verify a given Minecraft player based
on their Minecraft UUID (without hyphens).

## GET /verify/{Player UUID}
Possible Errors:
 * [Missing Player UUID Attribute](#Missing-Player-UUID-Attribute)

This endpoint checks if a player is allowed to join the Minecraft server.

Required Headers:
 1. Content-Type: `application/json`
 2. Authorization: `<webserver token>` 

| Attribute   | Type   | Description             |
|-------------|--------|-------------------------|
| Player UUID | string | The Minecraft player ID |

The `player UUID` is the Minecraft player UUID stripped of all the dashes. The server will provide
the following response if everything went alright, otherwise an error may occur.


### Response Body

#### Invalid Player - Response
This means they failed authentication

| Attribute | Type    | Description                                      |
|-----------|---------|--------------------------------------------------|
| verified  | boolean | Whether or not the given player is ready to play |
| reason    | string  | Possible reasons are described below             |

 - "no_role": They fail to have the required roles on Discord to join the
  Minecraft server.
 - "maintenance": The bot is in maintenance mode meaning only admin's can join.
 - "banned": This person is banned from using the bot and auth server.
 - "auth_code": This code needs to authorize their linkage
 
 
#### Invalid Player - Please Auth Response
 | Attribute | Type    | Description                                      |
 |-----------|---------|--------------------------------------------------|
 | verified  | boolean | Whether or not the given player is ready to play |
 | reason    | string  | Only is "auth_code"                              |
 | auth_code | string  | The auth code they must provide the Discord bot  |


#### Valid Player - Response
This means they can play on the Minecraft server.

| Attribute | Type    | Description                                      |
|-----------|---------|--------------------------------------------------|
| verified  | boolean | Whether or not the given player is ready to play |

The valid attribute is a boolean which represents whether the player can play on the Minecraft 
server. This will always return a boolean whether or not there was an issue getting the member
associated with the provided player ID.

An added "reason" attribute also exists. It will only be 'no_link' which means the Minecraft player
isn't linked with a Discord account and 'no_role' which means they're not whitelisted 

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

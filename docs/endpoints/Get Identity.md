# Player Verification
This is where the client can retrieve the Discord identity of a player based on their Minecraft UUID without initiating an auth code check.
This can be useful for debugging the current state of the user.

## GET /player/{Player UUID}
Possible Errors:
 * [Missing Player UUID Attribute](#Missing-Player-UUID-Attribute)

This endpoint finds the discord identity of a player.

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
There was a problem with finding the identity of the player.

| Attribute   | Type    | Description                                      |
|-------------|---------|--------------------------------------------------|
| player_type | string  | Possible types are described below               |

 - "not_found": This player could not be found.
 - "pending_auth": An auth code has been generated for the player,
however the code has not been submitted in Discord.


#### Valid Player - Response
The player exists, here is the linked account

**Note: This is not finalised! The user#1234 data might be added following discussion on how the data should be obtained**

| Attribute   | Type    | Description                                      |
|-------------|---------|--------------------------------------------------|
| snowflake   | string  | The Discord snowflake (id) of the player         |
| player_type | string  | Possible types are described below               |
 - "player": This is a regular player who is allowed on the server.
 - "alt": This is an alt account of a staff member
 - "banned": This person is banned from using the bot and auth server.

## Errors

### Missing Player UUID Attribute
```json
{
  "errcode": "NO_PLAYER_ID",
  "message": "There wasn't a player ID provided"
}
```

# Player Verification
This is where the client can retrieve the Discord identity of a player based on their Minecraft Name without initiating an auth code check.
This can be useful for debugging the current state of the user.

## GET /player/{Player Name}
Possible Errors:
 * [Missing Player Name Attribute](#Missing-Player-Name-Attribute)

This endpoint finds the discord identity of a player.

Required Headers:
 1. Content-Type: `application/json`
 2. Authorization: `<webserver token>` 

| Attribute   | Type   | Description             |
|-------------|--------|-------------------------|
| Player Name | string | The Minecraft player name |

The `Player Name` is the Minecraft player Name stripped of all the dashes. The server will provide
the following response if everything went alright, otherwise an error may occur.


### Response Body

#### No Data - Response
The player does not exist

| Attribute   | Type    | Description                                      |
|-------------|---------|--------------------------------------------------|
| player_type | string  | Possible types are described below               |
 - "not_found": This player could not be found
 - "outdated_spec": special state indicating that the client must be updated
in order to account for a change that is not backwards compatible.

#### Unlinked Player - Response

| Attribute   | Type    | Description                                      |
|-------------|---------|--------------------------------------------------|
| player_id   | string  | The Minecraft UUID for the player                |
| player_type | string  | Will be "pending_auth" - An auth code has been generated for the player, however the code has not been submitted in Discord. |

#### Valid Player - Response
The player exists, here is the linked account

| Attribute   | Type    | Description                                      |
|-------------|---------|--------------------------------------------------|
| snowflake   | string  | The Discord snowflake (id) of the player         |
| player_id   | string  | The Minecraft UUID for the player                |
| player_type | string  | Possible types are described below               |
 - "player": This is a regular player who is allowed on the server.
 - "alt": This is an alt account of a staff member
 - "banned": This person is banned from using the bot and auth server.

## Errors

### Missing Player Name Attribute
```json
{
  "errcode": "NO_PLAYER_ID",
  "message": "There wasn't a player ID provided"
}
```

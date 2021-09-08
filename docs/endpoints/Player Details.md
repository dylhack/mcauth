# Player Details Endpoint
To retrieve a user's roles and such you can utilize this endpoint which will
allow MCAuth clients to evaluate players their own way rather than relying on
the central MCAuth server to do so.

## GET /details/{playerID}
 1. Authorization: `<webserver token>` 

### Response

| Name      | Type     | Description                                 |
|:----------|:---------|:--------------------------------------------|
| id        | string   | The Minecraft player's Discord ID           |
| roles     | string[] | An array of role IDs                        |
| state     | State    | The player's state according to MCAuth      |
| auth_code | string   | An auth code to use to have the player link |

#### State

| Value (str) | Description                                       |
|:------------|:--------------------------------------------------|
| whitelisted | The player is verified                            |
| admin       | The player is an admin                            |
| alt_acc     | The player is an alt of an admin                  |
| no_link     | The player hasn't linked their MC account         |
| no_role     | The player doesn't have a whitelisted role        |
| auth_code   | The player needs to link using the code provided. |

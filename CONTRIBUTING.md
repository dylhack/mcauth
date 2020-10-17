# Contributing
If you're interested in investing your time into MCAuth's development then at
least RTFM. Once you have make sure your pull requests are heavily described
and give it the "why" you made the pull request.


## RTFM	
 1. [The MCAuth Client](./docs/1.%20The%20MCAuth%20Client.md)
 2. [The Web Server](./docs/2.%20Web%20Server.md)
    1. [Verifying Players](./docs/2.1.%20Verifying%20Players.md)
    2. [Alt Account Management](./docs/2.2.%20Alt%20Account%20Management.md)
 3. [The Discord Bot](./docs/3.%20Discord%20Bot.md)


## Internal Workings
As you dig deep into the source-code there's going to be a lot of methodology
unexplained in the documentation above. Most of it will be covered below, but
if something needs clarification then leave an issue.

### Role Syncing
The bot has something called the Sync struct, since we can't rely on the 
Discord API so heavily (because rate-limits) the bot will stay in sync with
every guild member's roles with what's given in the websocket. This is all
stored in memory so when MCAuth goes offline all the roles are cleared.

This is done in [/internal/bot/sync.go](./internal/bot/sync.go)

### Mojang API
For some situations we might need to figure out what a UUID is associated to
what player name or vice-versa so we can use the Mojang API to resolve them.

 * [Getting Player name w/ UUID](https://github.com/dylhack/mcauth/blob/production/internal/common/minecraft.go#L29-L51)
 * [Getting Player UUID w/ name](https://github.com/dylhack/mcauth/blob/production/internal/common/minecraft.go#L53-L83)

### Request Authentication
The webserver has a token attribute in the config.yml. Requests are checked up
in [internal/webserver/routes/authenticator.go](./internal/webserver/routes/authenticator.go)


go nuts.

# MCAuth
For linking your Minecraft and Discord account together. MCAuth is for Minecraft server owners
to allow the right players in their Discord to play on their server. MCAuth uses authentication
codes which allows players to link their accounts with ease.

MCAuth is under the [GNU General Public License 3](./LICENSE)

## Requirements
 * [PostgreSQL](https://www.postgresql.org/)
 * [Go 1.14](https://golang.org/)
 * [Spigot Minecraft Server](https://www.spigotmc.org/)
   * **Working Alternatives**
   * [Paper MC](https://papermc.io/)
 * [The MCAuth Client Minecraft Plugin](https://github.com/dhghf/mcauth-client)


## Setup

### 1. Build
First we need to compile the code this can be done like this
```
$ cd ./cmd/mcauth
$ go build
```

### 2. Configure
An executable file will be created in the directory, run it once, and it will 
generate a default config file. Fill out the config file. A guide is provided 
[here](./docs/extra/Configuration.md)

### 3. Setup the Plugin
[Visit the plugin's README](https://github.com/dhghf/mcauth-client/blob/master/README.md)

### 4. Setup Complete
Once the plug in has been setup and running it should now be protecting your
Minecraft server and actively listening to the players joining and verifying them.

## Bot Usage
__Regular Commands__

 * commands: Display these commands.
 * auth `code`: For linking your Minecraft account (this is given by the server on first time joining)
 * help: Display help message from config.yml
 * whoami: Displays the account you're linked with
 * whois `player name or @user`: Displays the account linked with a given player
 * status: MCAuth status

__Admin Commands__
 * unlink `player name or @user`: To unlink another user

## Under The Hood

### Databasing
We use Postgres driver [github.com/lib/pq](https://github.com/lib/pq) and ORM
[GORM](https://gorm.io)

### Verifying
When a new player joins the Minecraft server it first checks to see if they're
linked with an account.

#### If they are linked
If they are linked then it will get their roles that are stored in memory, if
their roles aren't already synced then the bot will reach out and stay
up-to-date on their roles so that the bot won't have to request the guild
member's roles every time they join in the future.


#### If they're not linked
If they aren't linked then it gets their already pending authentication code
or it generates the new one.

### Benchmarks
Last tested `July 26th, 2020`

__Notes__

A "sync" is when the Discord bot fetchs a player's roles on Discord. It will
keep up-to-date on the player's roles so it will never have to fetch the roles
again in the future.
 * This is only using the same 50 player UUID's
 * The first post is the initial sync of those 50 players
 * The second post is the same 50 players after the initial sync
 * 50 concurrent transactions, 60 seconds
 * This is using GORM + Postgres

***Before** Initial Role Sync*
```
** SIEGE 4.0.4
Transactions:                277572 hits
Availability:                100.00 %
Elapsed time:                59.11 secs
Data transferred:            4.55 MB
Response time:               0.01 secs
Transaction rate:            4695.85 trans/sec
Throughput:                  0.08 MB/sec
Concurrency:                 49.72
Successful transactions:     277572
Failed transactions:         0
Longest transaction:         27.25
Shortest transaction:        0.00
```

***After** Initial Role Sync*
```
** SIEGE 4.0.4
Transactions:                391772 hits
Availability:                100.00 %
Elapsed time:                59.76 secs
Data transferred:            6.43 MB
Response time:               0.01 secs
Transaction rate:            6555.76 trans/sec
Throughput:                  0.11 MB/sec
Concurrency:                 49.52
Successful transactions:     391772
Failed transactions:         0
Longest transaction:         0.21
Shortest transaction:        0.00
```

### Versioning
`x.y.z`
 - `x` A major change, the plugin needs to update
 - `y` A minor change, an important server-only change was made
 - `z` A patch, an insignificant change

### Further Dev Information
See [Endpoints.md](./docs/Endpoints.md)

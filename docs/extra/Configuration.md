# config.yml
## 1. Setup the Database
The database section of the configuration file is for connecting to a 
PostgreSQL database to store linked accounts, alt accounts, bans, and
authentication codes. **Make sure that the username you enter, has has
permission to create schemas.**


If you need help setting up a PostgreSQL database here's some references
 * [How To Install and Use PostgreSQL on Ubuntu Server](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-postgresql-on-ubuntu-18-04)

Once you have postgres install enter `psql` again and enter the following
commands to allow MCAuth to manage it's own part in the database.

Create the user & password
```sql
CREATE USER mcauth WITH PASSWORD 'password'; # Change 'password'
```

Give mcauth access to create schemas
```sql
GRANT CREATE ON DATABASE postgres TO mcauth;
```

Now fill in your config.yml
```yaml
database:
  host:                 "localhost"
  port:                 5432
  username:             "mcauth"
  password:             "password"
  database_name:        "postgres" # this should be "postgres"
  max_connections:      50
  max_idle_connections: 50
  conn_lifespan:        1h0m0s
```


## 2. Setup the Discord Bot
Go-to Discord's [Developer Portal](https://discord.com/developers/applications)
and create a new application, create a new bot by click bot on the left side 
bar, and copy the token and put it in your config.yml.

```yaml
discord_bot:
  help_message:         "Join the server, enter you auth code by typing .mc auth <code>"
  token:                "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGbG9vciBHYW5nIjoiT3VoISJ9.29wjTwrUk4XWQ1X9n9lCOP5R0B0O6PF7wdgBwNDNzig"
  prefix:               ".mc"
```

__Visual Guides__

![new application screenshot](../res/new%20app.png)
![new bot screenshot](../res/new%20bot.png)
![copy token screenshot](../res/bot%20token.png)


## 3. Setup the plugin
You'll need the port and token from the webserver section of the config. Follow
the guide [here](https://github.com/dylhack/mcauth-client) to install the plugin 
and configure it.

```yaml
webserver:
  port:                 5353
  token:                ""
```
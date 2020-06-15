## Cacao Discord Bot

This is a personal Discord bot that can control a Minecraft server using RCON protocol. I'm sure there are probably other bots out there that can do the same thing. However, that's no fun for me :). This is my first serious attempt at building something in Go as I'm still very much learning the language. I wanted a useful project to use for learning so creating a bot seemed like a good candidate.

I used [GoMusicBot](https://github.com/ducc/GoMusicBot) as an example of how to use the DiscordGo library. I also took their approach in how to handle commands and configuration as I wasn't sure the best way to handle this in Go. 

### Libraries Used

- [DiscordGo](https://github.com/bwmarrin/discordgo)
- [Go-Mc](https://github.com/Tnze/go-mc/net)
- [Godotenv](https://github.com/joho/godotenv)

### Configuration 

The configuration is handled through environment variables. I could have used JSON for this, but this is a simple straightforward approach to handle configuration values. Minecraft server and RCON information are provided in a dot env file. So it's pretty limited to self-hosting the bot for your server.  

The .env file should contain below:

```
# Bot Information
BOT_PREFIX=!mc
BOT_NAME=Cacao
BOT_VERSION=1.0.0
BOT_TOKEN=<Token String>
BOT_DEFAULT_STATUS=Minecraft

# Minecraft Information
MC_SERVER=<Server IP>
MC_RCON_PORT=25575
MC_RCON_PASSWORD=<Password String>
```

### Limitations

Currently limited to self-hosting and hard coding configuration values. Additionally, there is no role support yet so anyone can run these commands on your server if they know what commands can be run.

### Todo

- Finish basic administration commands like ban management, whitelisting, blacklisting, etc.
- Add a more comprehensive help command (currently, it just tells you what version is running).
- Add more comprehensive help for each command.
- Add basic role support so that only users with specific roles can perform actions.

### Future Nice to Haves

- Flesh out the bot so it can be hosted for multiple guilds.
- Implement custom RCON commands that go beyond what is offered already.
- Make a better bot avatar.
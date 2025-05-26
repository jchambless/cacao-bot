# Cacao Discord Bot

A Discord bot for managing Minecraft servers with a modern web interface.

![Build Status](https://github.com/jchambless/Cacao-Bot/workflows/Build%20and%20Publish%20Docker%20Image/badge.svg)
![Docker Pulls](https://img.shields.io/docker/pulls/jchambless/cacao-bot)
![License](https://img.shields.io/github/license/jchambless/Cacao-Bot)

## Features

- 🎮 **Minecraft Server Management**: Control your Minecraft server through Discord commands
- 🌐 **Web Interface**: Modern React-based dashboard with real-time server status
- 🐳 **Docker Support**: Easy deployment with Docker and docker-compose
- 🚀 **Multi-Architecture**: Supports both AMD64 and ARM64 platforms
- 📊 **Health Monitoring**: Built-in health checks and status monitoring
- 🔧 **RCON Integration**: Direct server control via RCON protocol

## Quick Start

### Using Docker (Recommended)

1. **Pull the image from Docker Hub:**
   ```bash
   docker pull jchambless/cacao-bot:latest
   ```

2. **Create a `.env` file:**
   ```env
   BOT_TOKEN=your_discord_bot_token
   HTTP_PORT=8080
   PREFIX=!mc
   DEFAULT_STATUS=Managing Minecraft Server
   RCON_HOST=your_minecraft_server_ip
   RCON_PORT=25575
   RCON_PASSWORD=your_rcon_password
   ```

3. **Run with docker-compose:**
   ```bash
   curl -O https://raw.githubusercontent.com/jchambless/Cacao-Bot/main/docker-compose.yml
   docker-compose up -d
   ```

4. **Or run directly:**
   ```bash
   docker run -p 8080:8080 --env-file .env jchambless/cacao-bot:latest
   ```

### Using the Deploy Script

For easier local development and deployment:

```bash
# Clone the repository
git clone https://github.com/jchambless/Cacao-Bot.git
cd Cacao-Bot

# Make the script executable
chmod +x deploy.sh

# Build and run locally
./deploy.sh build
./deploy.sh run

# Or start development environment
./deploy.sh dev
```

## Discord Commands

| Command | Description | Usage |
|---------|-------------|-------|
| `help` | Show available commands | `!mc help` |
| `ping` | Check server status | `!mc ping` |
| `list` | List online players | `!mc list` |
| `kick` | Kick a player | `!mc kick <player>` |
| `ban` | Ban a player | `!mc ban <player>` |
| `banip` | Ban an IP address | `!mc banip <ip>` |
| `pardon` | Unban a player | `!mc pardon <player>` |
| `op` | Give operator status | `!mc op <player>` |
| `deop` | Remove operator status | `!mc deop <player>` |
| `stop` | Stop the server | `!mc stop` |
| `about` | Bot information | `!mc about` |

## Web Interface

The bot includes a modern React-based web interface accessible at `http://localhost:8080`:

- 📊 **Server Status Dashboard**: Real-time server information
- 👥 **Player Management**: View online players
- 🔧 **Bot Configuration**: Manage bot settings
- 📈 **Analytics**: Server usage statistics

## Development

### Prerequisites

- Go 1.22+
- Node.js 18+
- Docker (optional)

### Local Development

1. **Clone the repository:**
   ```bash
   git clone https://github.com/jchambless/Cacao-Bot.git
   cd Cacao-Bot
   ```

2. **Set up environment:**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Install dependencies:**
   ```bash
   # Backend
   go mod download
   
   # Frontend
   cd web && npm install && cd ..
   ```

4. **Build and run:**
   ```bash
   # Build frontend
   cd web && npm run build && cd ..
   
   # Run the bot
   go run main.go
   ```

### Project Structure

```
Cacao-Bot/
├── main.go              # Main application entry
├── commands/            # Discord command handlers
├── framework/           # Bot framework and utilities
├── server/              # HTTP server and API endpoints
├── util/                # Utility functions
├── web/                 # React frontend application
├── Dockerfile           # Multi-stage Docker build
├── .env                 # Environment configuration
├── .github/workflows/   # GitHub Actions CI/CD
├── docker-compose.yml   # Docker Compose configuration
├── deploy.sh           # Deployment helper script
└── README.md           # This file
```

## Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `BOT_TOKEN` | Discord bot token | Required |
| `HTTP_PORT` | Web server port | `8080` |
| `PREFIX` | Command prefix | `!mc` |
| `DEFAULT_STATUS` | Bot status message | `Managing Minecraft Server` |
| `RCON_HOST` | Minecraft server IP | `localhost` |
| `RCON_PORT` | RCON port | `25575` |
| `RCON_PASSWORD` | RCON password | Required |

### Discord Bot Setup

1. Go to the [Discord Developer Portal](https://discord.com/developers/applications)
2. Create a new application
3. Go to the "Bot" tab
4. Create a bot and copy the token
5. Enable "Message Content Intent" under Privileged Gateway Intents
6. Invite the bot to your server with appropriate permissions

## Deployment

### GitHub Actions

The repository includes automated CI/CD with GitHub Actions:

- **Builds**: Automatically builds Docker images on push
- **Tests**: Runs frontend and backend tests
- **Publishes**: Pushes images to Docker Hub
- **Multi-arch**: Supports AMD64 and ARM64 platforms

#### Required Secrets

Set these in your GitHub repository settings:

- `DOCKER_USERNAME`: Your Docker Hub username
- `DOCKER_PASSWORD`: Your Docker Hub password or access token

### Manual Deployment

```bash
# Build and push to Docker Hub
./deploy.sh build
./deploy.sh push -u jchambless

# Or use Docker directly
docker build -t jchambless/cacao-bot .
docker push jchambless/cacao-bot
```

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Make your changes and test them
4. Commit your changes: `git commit -am 'Add feature'`
5. Push to the branch: `git push origin feature-name`
6. Submit a pull request

## Libraries Used

### Backend (Go)
- [DiscordGo](https://github.com/bwmarrin/discordgo) - Discord API wrapper
- [Gorilla Mux](https://github.com/gorilla/mux) - HTTP router
- [Go-MC](https://github.com/Tnze/go-mc) - Minecraft protocol implementation

### Frontend (React)
- [React](https://reactjs.org/) - UI framework
- [TypeScript](https://www.typescriptlang.org/) - Type safety
- [Tailwind CSS](https://tailwindcss.com/) - Styling
- [TanStack Query](https://tanstack.com/query) - Data fetching
- [Axios](https://axios-http.com/) - HTTP client

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- 📖 [Documentation](https://github.com/jchambless/Cacao-Bot/wiki)
- 🐛 [Report Issues](https://github.com/jchambless/Cacao-Bot/issues)
- 💬 [Discussions](https://github.com/jchambless/Cacao-Bot/discussions)

## Acknowledgments

- Inspired by [GoMusicBot](https://github.com/ducc/GoMusicBot) for Discord bot structure
- Thanks to the Discord.js and DiscordGo communities
- Minecraft RCON protocol implementation guidance

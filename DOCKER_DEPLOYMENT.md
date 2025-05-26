# Docker Deployment Guide

This project includes automated Docker image building and publishing to Docker Hub using GitHub Actions.

## GitHub Actions Setup

### Required Secrets

To enable automatic Docker image publishing, you need to set up the following secrets in your GitHub repository:

1. Go to your repository on GitHub
2. Navigate to Settings → Secrets and variables → Actions
3. Add the following secrets:

- `DOCKER_USERNAME`: Your Docker Hub username
- `DOCKER_PASSWORD`: Your Docker Hub password or access token (recommended)

### Docker Hub Access Token (Recommended)

Instead of using your Docker Hub password, it's recommended to create an access token:

1. Log in to Docker Hub
2. Go to Account Settings → Security → Access Tokens
3. Click "New Access Token"
4. Give it a descriptive name (e.g., "GitHub Actions - Cacao Bot")
5. Select "Read, Write, Delete" permissions
6. Copy the generated token and use it as `DOCKER_PASSWORD`

## Workflow Triggers

The workflow runs on:
- **Push to main branch**: Builds and pushes with `latest` tag
- **Push to develop branch**: Builds and pushes with `develop` tag
- **Git tags**: Builds and pushes with version tags (e.g., `v1.0.0`)
- **Pull requests**: Builds only (doesn't push to Docker Hub)

## Manual Deployment

### Local Build and Run

```bash
# Build the image locally
docker build -t cacao-bot ./src

# Run with docker-compose
docker-compose up -d

# Or run directly
docker run -p 8080:8080 \
  -e BOT_TOKEN="your_bot_token" \
  -e HTTP_PORT="8080" \
  -e PREFIX="!mc" \
  -e DEFAULT_STATUS="Managing Minecraft Server" \
  cacao-bot
```

### Pull from Docker Hub

Once the GitHub Actions workflow has run, you can pull the image:

```bash
# Pull latest version
docker pull your-username/cacao-bot:latest

# Pull specific version
docker pull your-username/cacao-bot:v1.0.0
```

## Environment Variables

The following environment variables can be configured:

- `BOT_TOKEN`: Discord bot token (required)
- `HTTP_PORT`: Port for the web server (default: 8080)
- `PREFIX`: Command prefix for Discord (default: !mc)
- `DEFAULT_STATUS`: Bot's default status message

## Health Check

The Docker image includes a health check endpoint at `/api/health` that can be used with:
- Docker health checks
- Kubernetes liveness/readiness probes
- Load balancer health checks

## Multi-Architecture Support

The workflow builds images for both:
- `linux/amd64` (Intel/AMD 64-bit)
- `linux/arm64` (ARM 64-bit, including Apple Silicon and ARM servers)

This ensures compatibility across different deployment platforms.

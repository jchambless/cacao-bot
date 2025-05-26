# 🚀 Cacao Bot - Complete Setup Guide

## ✅ What's Ready

Your Cacao Bot now has a complete Docker deployment pipeline with:

### 🔧 GitHub Actions Workflow
- **File**: `.github/workflows/docker-build.yml`
- **Triggers**: Push to main/develop, tags, pull requests
- **Features**: 
  - Multi-architecture builds (AMD64 + ARM64)
  - Automatic frontend building
  - Docker Hub publishing
  - Build caching for faster builds

### 🐳 Docker Setup
- **Multi-stage Dockerfile** for optimized images
- **Docker Compose** for easy local deployment
- **Health checks** and proper security practices
- **Deploy script** with helpful commands

### 🌐 React Frontend
- **Modern React + TypeScript** setup
- **Tailwind CSS** for styling
- **API integration** with your Go backend
- **Production build** integrated into Docker

## 🚀 Quick Start Commands

```bash
# Build and test locally
./deploy.sh build
./deploy.sh test

# Run development environment
./deploy.sh dev

# Build and push to Docker Hub
./deploy.sh push -u your-dockerhub-username

# Clean up Docker resources
./deploy.sh clean
```

## 📋 Setup Checklist

### 1. GitHub Repository Setup
- [ ] Push your code to GitHub
- [ ] Go to Settings → Secrets and variables → Actions
- [ ] Add secrets:
  - `DOCKER_USERNAME`: Your Docker Hub username
  - `DOCKER_PASSWORD`: Your Docker Hub password/token

### 2. Discord Bot Configuration
- [ ] Copy `.env.example` to `src/.env`
- [ ] Set your `BOT_TOKEN` from Discord Developer Portal
- [ ] Configure Minecraft RCON settings
- [ ] Enable "Message Content Intent" in Discord Developer Portal

### 3. Docker Hub (Optional)
- [ ] Create repository: `your-username/cacao-bot`
- [ ] Update README.md with your Docker Hub username

## 🔄 Deployment Workflow

1. **Development**: Code changes → Push to branch → GitHub Actions builds and tests
2. **Staging**: Merge to `develop` → Auto-deploy with `develop` tag
3. **Production**: Create tag `v1.0.0` → Auto-deploy with version tags + `latest`

## 📦 What Happens on Each Push

1. **Frontend Build**: React app built with npm
2. **Docker Build**: Multi-stage build creates optimized image
3. **Testing**: Automated tests verify functionality
4. **Publishing**: Images pushed to Docker Hub with proper tags
5. **Multi-arch**: Works on Intel/AMD and ARM servers

## 🎯 Next Steps

1. **Configure your bot token** in `src/.env`
2. **Push to GitHub** to trigger the first build
3. **Set up GitHub secrets** for Docker Hub publishing
4. **Test the deployment** with `./deploy.sh dev`
5. **Create your first release tag** (e.g., `v1.0.0`)

## 🐛 Troubleshooting

- **Build fails**: Check GitHub Actions logs
- **Health check fails**: Ensure valid Discord token
- **Permission errors**: Verify Docker Hub credentials
- **Frontend issues**: Check React build in `src/web/`

Your bot is now ready for production deployment! 🎉

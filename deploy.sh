#!/bin/bash

# Cacao Bot Build and Deploy Script

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Default values
COMMAND=""
IMAGE_NAME="cacao-bot"
TAG="latest"
DOCKER_USERNAME=""

# Help function
show_help() {
    echo "Cacao Bot Build and Deploy Script"
    echo ""
    echo "Usage: $0 [COMMAND] [OPTIONS]"
    echo ""
    echo "Commands:"
    echo "  build              Build the Docker image locally"
    echo "  run                Run the container locally"
    echo "  push               Push image to Docker Hub (requires login)"
    echo "  clean              Remove local images and containers"
    echo "  dev                Start development environment"
    echo "  test               Run tests and build verification"
    echo ""
    echo "Options:"
    echo "  -t, --tag TAG      Docker image tag (default: latest)"
    echo "  -u, --user USER    Docker Hub username for push"
    echo "  -h, --help         Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 build                   # Build image locally"
    echo "  $0 run                     # Run container with default settings"
    echo "  $0 push -u myusername      # Push to Docker Hub"
    echo "  $0 dev                     # Start development environment"
}

# Parse arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        build|run|push|clean|dev|test)
            COMMAND="$1"
            shift
            ;;
        -t|--tag)
            TAG="$2"
            shift 2
            ;;
        -u|--user)
            DOCKER_USERNAME="$2"
            shift 2
            ;;
        -h|--help)
            show_help
            exit 0
            ;;
        *)
            print_error "Unknown option: $1"
            show_help
            exit 1
            ;;
    esac
done

# Check if command is provided
if [[ -z "$COMMAND" ]]; then
    print_error "No command provided"
    show_help
    exit 1
fi

# Change to script directory
cd "$(dirname "$0")"

case $COMMAND in
    build)
        print_status "Building Docker image: $IMAGE_NAME:$TAG"
        docker build -t "$IMAGE_NAME:$TAG" ./src
        print_success "Image built successfully"
        ;;
        
    run)
        print_status "Running container: $IMAGE_NAME:$TAG"
        
        # Check if .env file exists
        if [[ -f "./src/.env" ]]; then
            print_status "Loading environment from ./src/.env"
            docker run -p 8080:8080 --env-file ./src/.env "$IMAGE_NAME:$TAG"
        else
            print_warning "No .env file found. Running with default settings."
            print_warning "Make sure to set BOT_TOKEN environment variable."
            docker run -p 8080:8080 \
                -e HTTP_PORT=8080 \
                -e PREFIX="!mc" \
                -e DEFAULT_STATUS="Managing Minecraft Server" \
                "$IMAGE_NAME:$TAG"
        fi
        ;;
        
    push)
        if [[ -z "$DOCKER_USERNAME" ]]; then
            print_error "Docker username required for push. Use -u flag."
            exit 1
        fi
        
        print_status "Tagging image for Docker Hub"
        docker tag "$IMAGE_NAME:$TAG" "$DOCKER_USERNAME/$IMAGE_NAME:$TAG"
        
        print_status "Pushing to Docker Hub: $DOCKER_USERNAME/$IMAGE_NAME:$TAG"
        docker push "$DOCKER_USERNAME/$IMAGE_NAME:$TAG"
        
        print_success "Image pushed successfully"
        ;;
        
    clean)
        print_status "Cleaning up Docker resources"
        
        # Stop and remove containers
        docker ps -a -q --filter ancestor="$IMAGE_NAME" | xargs -r docker stop
        docker ps -a -q --filter ancestor="$IMAGE_NAME" | xargs -r docker rm
        
        # Remove images
        docker images -q "$IMAGE_NAME" | xargs -r docker rmi
        
        # Clean up build cache
        docker builder prune -f
        
        print_success "Cleanup completed"
        ;;
        
    dev)
        print_status "Starting development environment"
        
        # Build the image first
        docker build -t "$IMAGE_NAME:dev" ./src
        
        # Run with docker-compose if available, otherwise use docker run
        if [[ -f "docker-compose.yml" ]]; then
            print_status "Using docker-compose for development"
            docker-compose up
        else
            print_status "Running container in development mode"
            docker run -p 8080:8080 --env-file ./src/.env "$IMAGE_NAME:dev"
        fi
        ;;
        
    test)
        print_status "Running tests and build verification"
        
        # Test frontend build
        print_status "Testing frontend build"
        cd src/web && npm ci && npm run build && cd ../..
        
        # Test Docker build
        print_status "Testing Docker build"
        docker build -t "$IMAGE_NAME:test" ./src
        
        # Test container startup
        print_status "Testing container startup"
        CONTAINER_ID=$(docker run -d -p 8081:8080 \
            -e BOT_TOKEN="test_token" \
            -e HTTP_PORT=8080 \
            "$IMAGE_NAME:test")
        
        # Wait a moment for startup
        sleep 5
        
        # Test health endpoint
        if curl -f http://localhost:8081/api/health >/dev/null 2>&1; then
            print_success "Health check passed"
        else
            print_error "Health check failed"
        fi
        
        # Cleanup test container
        docker stop "$CONTAINER_ID" >/dev/null
        docker rm "$CONTAINER_ID" >/dev/null
        
        print_success "All tests passed"
        ;;
esac

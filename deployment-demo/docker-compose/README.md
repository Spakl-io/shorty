# Shorty Traefik Link Shortener Deployment Guide

## Overview

This guide provides instructions on deploying Shorty, a link shortening middleware for Traefik, using Docker Compose. Shorty enables the creation of short URLs to redirect to longer URLs within a Traefik setup.

## Prerequisites

- Docker and Docker Compose installed.
- Basic understanding of Docker, Docker Compose, and Traefik.
- A domain name (for production deployment).

## Deployment Steps

### 1. Clone the Repository

Clone the `deploymenmt-demos` repository to your local machine or server:

```bash
git clone [REPOSITORY_URL]
cd deploymenmt-demos/docker-compose
```

### 2. Configure the Docker Compose File

The provided `docker-compose.yml` file includes the necessary setup for deploying Traefik with the Shorty plugin. Review and modify the file according to your environment. 

- Uncomment and configure the `networks` section if you need to create a new proxy network.
- Update the `labels` section under the `shorty` service to define your desired short URLs and their corresponding redirection targets.
- Ensure the volume paths and environment variables are correctly set.

### 3. Start the Services

Run the following command to start the services defined in your Docker Compose file:

```bash
docker-compose up -d
```

### 4. Accessing Traefik Dashboard

You can access the Traefik dashboard at `http://traefik.localhost` (or your configured domain) to monitor your routers, services, and middlewares.

### 5. Using Shorty

After deployment, Shorty will start redirecting the configured short URLs to their target URLs. You can add or modify the links directly in the Docker Compose file under the `labels` section of the `shorty` service.

## Adding New Links

To add a new short link, you need to update the labels in the `shorty` service within the `docker-compose.yml` file. Each link is defined as a label following this format:

```yaml
- "traefik.http.middlewares.my-shorty.plugin.shorty.links.[SHORT_NAME]=[TARGET_URL]"
```

Replace `[SHORT_NAME]` with your desired short URL identifier and `[TARGET_URL]` with the full URL you want to redirect to.

For example, to add a short link for `https://example.com`, you would add the following label:

```yaml
- "traefik.http.middlewares.my-shorty.plugin.shorty.links.example=https://example.com"
```

After adding or modifying links, save the file and re-deploy the services:

```bash
docker-compose up -d
```

## How It Works

- **Traefik as Reverse Proxy**: Traefik is used to route requests based on hostnames and paths.
- **Shorty Middleware**: Shorty, implemented as a middleware in Traefik, intercepts requests and performs URL redirections.
- **Docker Labels for Configuration**: The redirection rules for Shorty are defined using Docker labels in the `docker-compose.yml` file. This approach allows easy management of URL mappings.
- **Automatic Updates on Re-deployment**: Changes made in the Docker Compose file (like adding new links) are applied when the services are re-deployed.


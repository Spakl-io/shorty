# Introduction

Shorty is a link shortening middleware for Traefik. It allows you to create short URLs that redirect to longer URLs. This README outlines how to install and configure Shorty in your Traefik setup.


## Installation

### Using Traefik Static Configuration

1. **YAML Configuration**:
   Add Shorty to your Traefik YAML configuration file (`traefik.yml`):

   ```yaml
   experimental:
     plugins:
       shorty:
         moduleName: "github.com/Spakl-io/shorty"
         version: "v1.0.0"
   ```

   Replace `github.com/Spakl-io/shorty` and `v1.0.0` with the appropriate module name and version of your plugin.

2. **Command-Line Configuration**:
   Alternatively, you can configure Traefik to use Shorty via command line:

   ```bash
   traefik --experimental.plugins.shorty.moduleName="github.com/Spakl-io/shorty" --experimental.plugins.shorty.version="v1.0.0"
   ```

## Configuration

### Defining Labels for Docker Services

Add the following labels to your Docker service to use Shorty. These labels define the router, service, middleware, and link mappings.

```yaml
services:
  your-service:
    labels:
      - "traefik.http.routers.shorty.rule=Host(`shorty.io`)"
      - "traefik.http.routers.shorty.service=traefik-shorty@docker"
      - "traefik.http.routers.shorty.middlewares=my-shorty@docker"
      - "traefik.http.middlewares.my-shorty.plugin.shorty.links.proxmox=https://proxmox.ops.spakl.io"
      - "traefik.http.middlewares.my-shorty.plugin.shorty.links.proxmox2=https://proxmox2.ops.spakl.io"
      - "traefik.http.middlewares.my-shorty.plugin.shorty.links.google=https://google.com"
```

### Setting Up a Domain Name

The `Host` rule in the router label defines the domain name. For example, `Host(`shorty.localhost`)` sets up `shorty.localhost` as the domain.

### Example Docker-Compose

Here's an example `docker-compose.yml` snippet:

```yaml
version: '3'

services:
  traefik:
    # ... Traefik configuration ...
    command:
      - "--experimental.plugins.shorty.moduleName=github.com/Spakl-io/shorty"
      - "--experimental.plugins.shorty.version=v1.0.0"
    # other configurations...

  your-service:
    image: your-service-image
    labels:
      - "traefik.http.routers.shorty.rule=Host(`shorty.localhost`)"
      - "traefik.http.routers.shorty.service=traefik-shorty@docker"
      - "traefik.http.routers.shorty.middlewares=shorty1@docker"
      - "traefik.http.middlewares.shorty1.plugin.shorty.links.proxmox=https://proxmox.ops.spakl.io"
      - "traefik.http.middlewares.shorty1.plugin.shorty.links.proxmox2=https://proxmox2.ops.spakl.io"
      - "traefik.http.middlewares.shorty1.plugin.shorty.links.google=https://google.com"
```

## Usage

Once configured, Shorty will intercept requests to the specified short URLs and redirect them to the corresponding target URLs.

## Contributing

Contributions to Shorty are welcome. Please refer to the project's GitHub repository for contribution guidelines.

## License

Specify your license or link to the license file.


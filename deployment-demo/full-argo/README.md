# Shorty Traefik Link Shortener Kubernetes Deployment Guide with Argo CD

## Overview

This guide covers deploying the Shorty Traefik Link Shortener in a Kubernetes environment using Argo CD for continuous deployment and Cert-Manager for SSL certificates.

## Prerequisites

- A Kubernetes cluster with Argo CD and Cert-Manager installed.
- Access to a Git repository (e.g., GitLab) for Argo CD synchronization.
- A registered domain name for your Shorty instance through cloudflare.
- Cloudflare token with DNS edit 

## Deployment Steps

### 1. Prepare Your Git Repository

Ensure your Git repository contains all necessary Kubernetes manifests, including Certificates, IngressRoute, Middleware, Cloudflare DNS solver, and any additional configurations for Shorty and Traefik.

### 2. Configure the Argo CD Application Manifest

Update the Argo CD `Application` manifest with the path to the Shorty configuration in your repository. Replace `'git@gitlab.com:spakl/lab/clusters/clstr.git'` with the URL of your Git repository:

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: shorty-sync
  namespace: argocd
spec:
  project: default
  source:
    repoURL: 'git@gitlab.com:spakl/lab/clusters/clstr.git' # Replace with your repository
    targetRevision: master
    path: 'apps/shorty' # Replace with the path to your Shorty configurations
  destination:
    server: 'https://kubernetes.default.svc'
    namespace: shorty
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
    syncOptions:
      - CreateNamespace=true
```

### 3. Deploy with Argo CD

Apply the Argo CD Application manifest to your Kubernetes cluster. Argo CD will automatically sync and deploy all configurations found in the specified path of your repository:

```bash
kubectl apply -f [PATH_TO_ARGOCD_APPLICATION_MANIFEST]
```

Replace `[PATH_TO_ARGOCD_APPLICATION_MANIFEST]` with the location of your Argo CD Application manifest file.

### 4. Verify Deployment

Once applied, Argo CD will create the necessary resources and deploy Shorty. You can monitor the deployment status and manage configurations directly from the Argo CD dashboard.

### 5. Access Shorty

After deployment, access Shorty via the configured domain. The Shorty middleware in Traefik will handle the redirection from short URLs to the specified long URLs.

## Adding New Links

To add new links, update the `Middleware` manifest in your Git repository with new entries under the `shorty.links` section and commit the changes. Argo CD will automatically apply the updates to your cluster.


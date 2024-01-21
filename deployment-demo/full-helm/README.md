## Prerequisites

- Helm and `kubectl` installed on your local machine.
- Access to the Kubernetes configurations.

## Deployment Steps

### 1. Deploy Traefik using Helm

The first part of the Argo CD application deploys Traefik using its official Helm chart. Follow these steps to replicate this:

1. **Add Traefik Helm Repository** (if not already added):

   ```bash
   helm repo add traefik https://helm.traefik.io/traefik
   helm repo update
   ```

2. **Install Traefik using Helm**:

   Assuming you have a `values.yaml` file for Traefik in your current directory, run:

   ```bash
   helm install traefik traefik/traefik --version v24.0.0 -f ./values.yaml --namespace shorty --create-namespace
   ```

   Replace `./values.yaml` with the path to your Traefik values file.

### 2. Apply Additional Kubernetes Configurations

**Apply Kubernetes Configurations**:

   Apply each of your Kubernetes manifest files using `kubectl`:

   ```bash
   kubectl apply -f ./resources --namespace shorty
   ```

### 3. Verify Deployment

After applying the Helm chart and Kubernetes manifests, verify that all components are correctly deployed:

```bash
kubectl get all -n shorty
```

### 4. Access and Manage Shorty

Once deployed, access Shorty via the configured domain. Any modifications, like adding new links, should be done by updating the respective Kubernetes manifests and re-applying them with `kubectl`.

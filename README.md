# Prerequisites

- gcloud ( https://cloud.google.com/sdk/docs/install )
- docker ( https://docs.docker.com/get-docker/ )

## On Windows

- WSL

# Changes

- Update top two lines in Makefile with service name and project-id

# GCP

- Acticvate Cloud Run: https://console.cloud.google.com/run
- Acticvate Container Registry: https://console.cloud.google.com/gcr
  - Click "Create service", but you don't need to actually make it, it's needed to enable it.
- gcloud auth login
- gcloud auth configure-docker

---
title: Cloud Run
keywords: [cloud run, deploy, google, docker, gcp]
description: Deploying to Google Cloud Run.
---

# Cloud Run Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/cloud-run) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/cloud-run)

This example demonstrates how to deploy a Go Fiber application to Google Cloud Run.

## Description

This project provides a starting point for deploying a Go Fiber application to Google Cloud Run. It includes necessary configuration files and scripts to build and deploy the application using Docker and Google Cloud Build.

## Requirements

- [Go](https://golang.org/dl/) 1.18 or higher
- [Docker](https://www.docker.com/get-started)
- [Google Cloud SDK](https://cloud.google.com/sdk/docs/install)
- [Git](https://git-scm.com/downloads)

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/gofiber/recipes.git
    cd recipes/cloud-run
    ```

2. Install the dependencies:
    ```bash
    go mod download
    ```

3. Build the Docker image:
    ```bash
    docker build -t cloud-run-example .
    ```

4. Run the Docker container:
    ```bash
    docker run -p 3000:3000 cloud-run-example
    ```

The application should now be running on `http://localhost:3000`.

## Deploy to Google Cloud Run

1. Set up Google Cloud SDK and authenticate:
    ```bash
    gcloud auth login
    gcloud config set project [YOUR_PROJECT_ID]
    ```

2. Build and push the Docker image using Google Cloud Build:
    ```bash
    gcloud builds submit --tag gcr.io/[YOUR_PROJECT_ID]/cloud-run-example
    ```

3. Deploy the image to Cloud Run:
    ```bash
    gcloud run deploy cloud-run-example --image gcr.io/[YOUR_PROJECT_ID]/cloud-run-example --platform managed --region [YOUR_REGION] --allow-unauthenticated
    ```

Replace `[YOUR_PROJECT_ID]` and `[YOUR_REGION]` with your Google Cloud project ID and desired region.

## Cloud Build Configuration

The `cloudbuild.yaml` file defines the steps to build and deploy the application using Google Cloud Build:

```yaml
steps:
  - name: 'gcr.io/kaniko-project/executor:latest'
    id: 'build-and-push'
    args:
      - '--destination=asia.gcr.io/$PROJECT_ID/$_SERVICE_NAME:$SHORT_SHA'
      - '--destination=asia.gcr.io/$PROJECT_ID/$_SERVICE_NAME:latest'
      - '--dockerfile=Dockerfile'
      - '--context=.'
      - '--cache=true'
      - '--cache-ttl=120h'

  - id: 'Deploy to Cloud Run'
    name: 'gcr.io/cloud-builders/gcloud'
    entrypoint: 'bash'
    args:
      - '-c'
      - |
        gcloud run deploy $_SERVICE_NAME \
        --image=asia.gcr.io/$PROJECT_ID/$_SERVICE_NAME:$SHORT_SHA \
        --region=$_REGION --platform managed --allow-unauthenticated \
        --port=3000
options:
  substitutionOption: ALLOW_LOOSE

substitutions:
  _SERVICE_NAME: cloud-run-example
  _REGION: asia-southeast1
```

## Example Usage

1. Open your browser and navigate to the Cloud Run service URL provided after deployment.

2. You should see the message: `Hello, World!`.

## Conclusion

This example provides a basic setup for deploying a Go Fiber application to Google Cloud Run. It can be extended and customized further to fit the needs of more complex applications.

## References

- [Google Cloud Run Documentation](https://cloud.google.com/run/docs)
- [Fiber Documentation](https://docs.gofiber.io)
- [Docker Documentation](https://docs.docker.com/)

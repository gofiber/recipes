---
title: Firebase GCloud
keywords: [firebase, gcloud, cloud run, cloud function, app engine]
---

# Deploy Fiber to Google Cloud with Firebase

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/gcloud) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/gcloud)

Examples on how to run an application using Fiber on Google Cloud.

## Running Locally

* Run on the command line:

```
go run cmd/main.go
```

## Deploy using Google Cloud Run

This step will build a Docker Image, publish to Google Cloud Registry and deploy on Cloud Run Managed enviroment.

[![Run on Google Cloud](https://storage.googleapis.com/cloudrun/button.svg)](https://console.cloud.google.com/cloudshell/editor?shellonly=true&cloudshell_image=gcr.io/cloudrun/button&cloudshell_git_repo=https://github.com/gofiber/recipes&cloudshell_working_dir=gcloud)

After deploying the server on Cloud Run, you can get it's url on GCP Console ([link](https://console.cloud.google.com/run)) and select the service `gcloud-fiber` that we just deployed. Them copy the URL will look like `https://{project-id}-{some-random-hash-string}.a.run.app`.

Or you can do it manually with those steps:

* Run on the command line:

```
export GCLOUD_PROJECT=[YOUR_PROJECT_ID]
gcloud builds submit — -tag gcr.io/$GCLOUD_PROJECT/gcloud-fiber .
gcloud beta run deploy --platform managed --image gcr.io/$GCLOUD_PROJECT/gcloud-fiber
```

## Deploy using Google App Engine

This step will deploy the app to Google App Engine Standard Go enviroment. The app configuration and additional configurations can be tweaked on the `app.yaml` file.

* Run on the command line:

```
gcloud app deploy
```

## Deploy using Google Cloud Function

This step will deploy a HTTP Cloud Function using Go enviroment. You can use the `deploy.sh` script. Just edit your project id on it.

For the Cloud Functions env, Google enforces us to deploy a function that is a `http.HandlerFunc`, so on the file `functions.go` there is a workaround to reroute the HTTP call to the Fiber app instance.

* Run on the command line:

```
gcloud functions deploy MyCloudFunction --runtime go111 --trigger-http
```

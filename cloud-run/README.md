# [Cloud Run](https://cloud.google.com/run)
Deploy and run serverless container application

## Run locally
```bash
go run main.go
```

## Deploy
Make sure you have the permission

#### Manually
```bash
gcloud builds submit . \
    --substitutions SHORT_SHA=$(git rev-parse --short HEAD)
```

#### Automated
Trigger currently supports source code from:
 - [Cloud Source Repositories](https://cloud.google.com/source-repositories)
 - [Github] (https://github.com)

 Learn more on [official docs](https://cloud.google.com/cloud-build/docs/automating-builds/create-manage-triggers)



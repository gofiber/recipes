---
title: AWS SAM
keywords: [aws, sam, serverless, lambda]
description: Serverless applications with AWS SAM.
---

# AWS SAM

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/aws-sam) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/aws-sam)

This is a sample template for sam-app - Below is a brief explanation of what we have generated for you:

```bash
.
├── README.md                   <-- This instructions file
├── app                         <-- Source code for a lambda function
│   └── main.go                 <-- Lambda function code
└── template.yaml
```

## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)
* SAM CLI - [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

## Setup process

### Installing dependencies & building the target

In this example we use the built-in `sam build` to automatically download all the dependencies and package our build target.
Read more about [SAM Build here](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-cli-command-reference-sam-build.html)

```shell
sam build --use-container
```

### Local development

**Invoking function locally through local lambda invoke**

```bash
sam local start-api

curl -XPOST "http://localhost:3001/2015-03-31/functions/sam-app/invocations"
{"statusCode":200,"headers":null,"multiValueHeaders":{"Content-Type":["application/json"]},"body":"{\"message\":\"Hello World\"}"}%
```




## Packaging and deployment

AWS Lambda Golang runtime requires a flat folder with the executable generated on build step. SAM will use `CodeUri` property to know where to look up for the application:

```yaml
...
    FirstFunction:
        Type: AWS::Serverless::Function
        Properties:
            CodeUri: app/
            ...
```

To deploy your application for the first time, run the following in your shell:

```bash
sam deploy --guided
```

The command will package and deploy your application to AWS, with a series of prompts:

* **Stack Name**: The name of the stack to deploy to CloudFormation. This should be unique to your account and region, and a good starting point would be something matching your project name.
* **AWS Region**: The AWS region you want to deploy your app to.
* **Confirm changes before deploy**: If set to yes, any change sets will be shown to you before execution for manual review. If set to no, the AWS SAM CLI will automatically deploy application changes.
* **Allow SAM CLI IAM role creation**: Many AWS SAM templates, including this example, create AWS IAM roles required for the AWS Lambda function(s) included to access AWS services. By default, these are scoped down to minimum required permissions. To deploy an AWS CloudFormation stack which creates or modifies IAM roles, the `CAPABILITY_IAM` value for `capabilities` must be provided. If permission isn't provided through this prompt, to deploy this example you must explicitly pass `--capabilities CAPABILITY_IAM` to the `sam deploy` command.
* **Save arguments to samconfig.toml**: If set to yes, your choices will be saved to a configuration file inside the project, so that in the future you can just re-run `sam deploy` without parameters to deploy changes to your application.

You can find your API Gateway Endpoint URL in the output values displayed after deployment.


## Add Permission to the Lambda Function for Public Access

After deploying your Lambda function with an associated function URL, you might encounter a scenario where the function URL is not accessible due to missing permissions for public access. This is common when the authentication type for the function URL is set to "None," indicating that the function is intended to be publicly accessible without authentication.

To ensure your Lambda function URL can be invoked publicly, you need to add the necessary permission that allows unauthenticated requests. This step is crucial when your function URL's authentication type is "None" but lacks the requisite permissions for public invocation.

Manually Configuring Permissions
You can manually configure permissions through the AWS Lambda console by creating a resource-based policy that grants the lambda:invokeFunctionUrl permission to all principals (*). This approach is straightforward but not suitable for automation within deployment pipelines.

Automating Permission Configuration
For a more automated approach, especially useful in CI/CD pipelines, you can use the AWS CLI or SDKs to add the necessary permissions after deploying your Lambda function. This can be incorporated into your deployment scripts or CI/CD workflows.

Here is an example AWS CLI command that adds the required permission for public access to your Lambda function URL:

```shell
aws lambda add-permission \
  --function-name <your-function-name> \
  --action lambda:InvokeFunctionUrl \
  --principal "*" \
  --function-url-auth-type "NONE" \
  --statement-id unique-statement-id
```

This command grants permission to all principals (*) to invoke your Lambda function URL, enabling public access as intended.

# Appendix

### Golang installation

Please ensure Go 1.x (where 'x' is the latest version) is installed as per the instructions on the official golang website: https://golang.org/doc/install

A quickstart way would be to use Homebrew, chocolatey or your linux package manager.

#### Homebrew (Mac)

Issue the following command from the terminal:

```shell
brew install golang
```

If it's already installed, run the following command to ensure it's the latest version:

```shell
brew update
brew upgrade golang
```

#### Chocolatey (Windows)

Issue the following command from the powershell:

```shell
choco install golang
```

If it's already installed, run the following command to ensure it's the latest version:

```shell
choco upgrade golang
```

## Bringing to the next level

Here are a few ideas that you can use to get more acquainted as to how this overall process works:

* Create an additional API resource (e.g. `/hello/{proxy+}`) and return the name requested through this new path
* Update unit test to capture that
* Package & Deploy

Next, you can use the following resources to know more about beyond hello world samples and how others structure their Serverless applications:

* [AWS Serverless Application Repository](https://aws.amazon.com/serverless/serverlessrepo/)

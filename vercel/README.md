# [Vercel](https://vercel.com)

Deploy and run a serverless go fiber application on vercel. With this recipie, you can also declare your routes using the fiber router.

## Config

Make sure to add the `vercel.json` file so the routing will work properly with vercel. This will rewrite all requests to the application to the `api/index.go` handler where the router will take over.

```json
{
  "rewrites": [
    { "source": "(.*)", "destination": "api/index.go" }
  ]
}
```

## Deploy

Deploy this application to vercel by clicking the button below.

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2Fgofiber%2Frecipes%2Fvercel)


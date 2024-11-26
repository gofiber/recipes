---
title: CSRF
keywords: [csrf, security, hacking, vulnerability]
---

# CSRF Examples

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/csrf) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/csrf)

Example Cross Site Request Forgery (CSRF) vulnerabilities in action.


## Requirements

* [git](https://git-scm.com/downloads)
* [Golang](https://golang.org/)


## Install Go Modules

Like any golang project, you will need to download and install the required modules for the project to run. Change into the "csrf" directory:
```bash
cd csrf
```

And then:
```bash
go mod vendor && go mod download && go mod tidy
```
This command installs the golang dependencies needed to run the project in a new directory named `vendor`.

Once the modules have finished installing, you can run the project like this:
```bash
go run main.go
```
OR
```bash
go run main.go withoutCsrf
```

You should see the following if everything is OK:
```
Server started and listening at localhost:3000
```

## Try the demo

Start the server without csrf, to see the dangers of these attacks
```bash
go run main.go withoutCsrf
```
Open your browser to and navigate to [localhost:3000](http://localhost:3000).

Login using the test account:
* Username: `bob`
* Password: `test`

In a new tab, navigate to [localhost:3001](http://localhost:3001) to view some examples of CSRF exploits. You will notice that the balance goes down everytime you load that page. This is because the page is successfully exploiting a CSRF vulnerability.


## See the "fixed" version

To see the csrf version of this demo, just stop the server by pressing __CTRL + C__ to kill the server process and then run
```bash
go run main.go
```

Navigate again to [localhost:3000](http://localhost:3000) and login to the test account.

And once more try the page with the CSRF exploits: [localhost:3001](http://localhost:3001).

You will notice now that the account balance is unchanged.


## Going further

Here are some useful links where you can learn more about this topic:
* https://en.wikipedia.org/wiki/Cross-site_request_forgery
* https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)

# CSRF-with-session Example

Example Cross Site Request Forgery (CSRF) middleware for GoFiber web applications using sessions.


## Requirements

* [git](https://git-scm.com/downloads)
* [Golang](https://golang.org/)


## Install Go Modules

Like any golang project, you will need to download and install the required modules for the project to run. Change into the "csrf-with-session" directory:
```bash
cd csrf-with-session
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

You should see the following if everything is OK:
```
Server started and listening at localhost:3000
```

## Try the demo

Start the server by running:
```bash
go run main.go
```
Open your browser to and navigate to [localhost:3000](http://localhost:3000).

Login using one of the test accounts:
* Username: `user1`
* Password: `password1`
OR
* Username: `user2`
* Password: `password2`

Once logged in, you will be able to see the /protected page.

## CSRF Protection

All methods except GET, HEAD, OPTIONS, and TRACE are checked for the CSRF token. If the token is not present or does not match the token in the session, the request is aborted with a 403 Forbidden error.

## Token Lyfecycle

The CSRF token is generated when the user visits any page on the site. The token is stored in the session and is valid for until it expires, of the users scope changes (e.g. they log in, or log out).

It is important that CSRF tokens do not presist beyond the scope of the user's session, that a new session is created when the user logs in, and that the session is destroyed when the user logs out.

## Session Storage

Sessions are stored in memory for this example, but you can use any session store you like. See the [Fiber session documentation](https://docs.gofiber.io/api/middleware/session) for more information.

### Note on pre-sessions

GoFiber's CSRF middleware will automatically create a session if one does not already exist. That means that we always have pre-sessions when using the CSRF middleware. Because we have pre-sessions: in this example we demonstrate that a session variable `loggedIn` is set to `true` when the user logs in, as to distinguish between logged in and logged out users.

## Going further

Here are some useful links where you can learn more about this topic:
* https://en.wikipedia.org/wiki/Cross-site_request_forgery
* https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)

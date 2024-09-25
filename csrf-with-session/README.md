# CSRF-with-session Example

Example GoFiber web app using Cross Site Request Forgery (CSRF) middleware with session.

This example impliments multiple best-practices for CSRF protection:

- CSRF Tokens are linked to the user's session.
- Pre-sessions are used, so that CSRF tokens are always available, even for anonymous users (eg for login forms).
- Cookies are set with a defense-in-depth approach:
    - Secure: true
    - HttpOnly: true
    - SameSite: Lax
    - IdleTimeout: 30 minutes (of inactivity)
    - Cookie names are prefixed with "__Host-" (see [MDN-Set-Cookie](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie) for more information))
- A custom Middleware is used to enforce an absolute timeout of 12 hours for the session.

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
Server started and listening at 127.0.0.1:8443
```

## Try the demo

Start the server by running:
```bash
go run main.go
```
Open your browser to and navigate to [127.0.0.1:8443](http://127.0.0.1:8443).


### Accept the self-signed certificate warning and visit the site.

In Chrome:

- Click on "Advanced"
- Click on "Proceed to 127.0.0.1:8443 (unsafe)"

In Firefox:

- Click on "Advanced"
- Click on "Accept the Risk and Continue"

In Safari:

- Click on "Show Details"
- Click on "visit this website"


### Try to access the /protected page

Login using one of the test accounts:
* Username: `user1`
* Password: `password1`
OR
* Username: `user2`
* Password: `password2`

Once logged in, you will be able to see the /protected page.


### Submit the form on the /protected page

Once logged in, you will be able to see the /protected page. The /protected page contains a form that submits to the /protected page. If you try to submit the form without a valid CSRF token, you will get a 403 Forbidden error.


## CSRF Protection

All methods except GET, HEAD, OPTIONS, and TRACE are checked for the CSRF token. If the token is not present or does not match the token in the session, the request is aborted with a 403 Forbidden error.


## Token Lifecycle

The CSRF token is generated when the user visits any page on the site. The token is stored in the session and is valid for until it expires, or the authorization scope changes (e.g. the user logs in, or logs out).

It is important that CSRF tokens do not persist beyond the scope of the user's session, that a new session is created when the user logs in, and that the session is destroyed when the user logs out.

The CSRF middleware has a `SingleUseToken` configuration option that can be used to generate a new token for each request. This is useful for some applications, but is not used in this example. Single use tokens have usability implications in scenarios where the user has multiple tabs open, or when the user uses the back button in their browser.


## Session Storage

Sessions are stored in memory for this example, but you can use any session store you like. See the [Fiber session documentation](https://docs.gofiber.io/api/middleware/session) for more information.


### Note on pre-sessions

GoFiber's CSRF middleware will automatically create a session if one does not exist. That means that we always have pre-sessions when using the CSRF middleware. In this example we set a session variable `loggedIn` 
to `true` when the user logs in, in order to distinguish between logged in and logged out users.

### Notes on the Session Timeout

The session absolute timeout is set to 12 hours in this example. This is a reasonable default, but you may want to adjust this value based on your application's requirements. The session absolute timeout is enforced by the custom `sessionExpirationMiddleware`, which resets the session if the user has been inactive for more than 12 hours.

The session idle timeout is set to 30 minutes in this example. This is a reasonable default, but you may want to adjust this value based on your application's requirements. The session idle timeout is enforced by the session middleware, which automatically deletes the session from the underlying store if the user has been inactive for more than 30 minutes.

### Notes on authentication and authorization

This example uses a simple username and password for authentication. In a real-world application, you should use a more secure authentication mechanism. Usernames and passwords should be securely stored, salted, and hashed. Implementing Multi-Factor Authentication (MFA) with username and password can significantly enhance security by requiring multiple forms of verification. Additionally, consider using OAuth2 or OpenID Connect, both of which support MFA for even greater security.

## Going further

Here are some useful links where you can learn more about this topic:
* https://en.wikipedia.org/wiki/Cross-site_request_forgery
* https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)

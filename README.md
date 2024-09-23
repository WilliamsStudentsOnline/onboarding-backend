# WSO Backend Onboarding

What you need:
- Golang
- A code editor (e.g. VSCode or GoLand)

To start the server, simply run:
```console
make run
```
and to quit, use the keyboard shortcut `control + C`

## Structure of this repository

- `cmd` This stores the entry point for the server under `main.go`
- `docs` This is where auto-generated documentation is stored
- `internal` This is where the actual private logic is stored
  - `router` Here we define the endpoint groups and set up the root router
  - `users` Here endpoints related to users and their respective handler function
  - `admin` Here are endpoints that require admin authorization and their respective handler function
  - `fakedatabase` This is where our (fake) user database logic is stored.
- `model`  This stores the defined structs for requests/responses and data models

## RESTful API

## TODO: this is not???? necessarily covered in the slides unless we are doing slide presentation during onboarding
## nonetheless i think we should def include something about rest here
Covered in our slideshow. 

## Gin

Gin is a framework that allows us to build highly efficient servers. It works around the idea of a `gin.Context`. A `context` is a wrapper for a http request instance. It stores all the information received from a request, and all the necessary information to respond to it as well. Every time a request is made, we call one of the functions that takes in a `gin.Context`. That funciton, uses the `context` to respond to the request.

## API Testing & Debugging

You can send a HTTP request to your local server by calling

```console
curl -X GET http://localhost:8080/<endpoint>

Here, `-X GET` specifies the HTTP method.

```
For example,
```console
curl -X GET http://localhost:8080/ping
```

## Tasks

### Part 1

Curretly if you run the server and make a GET request to the `ping` endpoint,

```console
curl -X GET http://localhost:8080/ping
```
You would get back pong. We want to change this, so that it returns `Hello World!`. You first task, is to find `ping` API handler, and modify it's response.

### Part 2

We have a database called `fakedatabase` that stores our WSO users and their favorite colors. Currently, there is no way for us to add any new users to the database. Luckily another WSO member has created an edit user endpoint, however it's not working correctly.

When we run
```console
curl -X POST \
  http://localhost:8080/api/v1/admin/edit \
  -H 'authorization: Basic Zm9vOmJhcg==' \
  -H 'content-type: application/json' \
  -d '{"user":"Ye","color":"Blue"}'
```
it is supposed to add the corresponding entry to the database.

Note that `Zm9vOmJhcg==` is `base64("foo:bar")` and we are using [Basic Auth](https://en.wikipedia.org/wiki/Basic_access_authentication) to authenticate users to this endpoint.
### TODO: make an actual task here

### Part 3

Now that we can add users, we want a way to query our database for a given user's favorite color. We want to make a request like so
```console
curl -X GET http://localhost:8080/api/v1/user/Ye
```
and get the information back in a JSON format like

```json
{
  "user": "Ye",
  "color": "Blue"
}
```
with a success http status code (200). If the user is not found, return an empty JSON with not found status code (404). Find the GET user endpoint, and modify it to query the database for the given name and return a JSON like above.

Note that you can use our pre-defined structs for constructing jsons (they have pre-defined json bindings)

### Part 4

Now that our server is fully implemented, let's generate the documentation for it. As you probabably noticed, there are a bunch of annotations above each `gin.Context` function. These annotations are used to autogenerate a documentation website for us.
To generate the documentation, run:
```console
make swag
```
### NOTE YOU MAKE HAVE TO INSTALL SWAG
And to see your documentation, start up the server and navigate to
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html).
From here, you can actually test out your endpoints right there in the browser! Play around with it, and try to make it execute the same endpoints as the ones you ran using `curl`.

# Welcome

## TOPIC

We're going to build an RSS feed aggregator in Go! It's a web server that allows clients to:

- Add RSS feeds to be collected


RSS feeds are a way for websites to publish updates to their content. You can use this project to keep up with your favorite blogs, news sites, podcasts, and more!

## Learning goals

- Learn how to integrate a Go server with PostgreSQL
- Learn about the basics of database migrations


## Setup

Before we dive into the project, let's make sure you have everything you'll need on your machine.

- An editor. I use VS code.
- A command line
- The latest Go toolchain.
- If you're in VS Code, I recommend the official Go extension.
- An HTTP client. I use Thunder Client.

If you're ready, move on to the next step!

## 2. Install packages

Install the following packages using `go get`:

- chi
- [cors](https://github.com/go-chi/cors)
- godotenv

## 3. Env

Create a gitignore'd `.env` file in the root of your project and add the following:

```bash
PORT="8080"
```

The `.env` file is a convenient way to store environment  variables.

- Use `godotenv.Load()` to load the variables from the file into your environment at the top of `main()`.
- Use `os.Getenv()` to get the value of `PORT`.

## 4. Create a router and server

- Create a `chi.NewRouter`.
- Use `router.Use` to add the built-in `cors.Handler` middleware.
- Create a sub-router for the `/v1` namespace and mount it to the main router.
- Create a new `http.Server` and add the port and the main router to it.
- Start the server.

## 5. Create some JSON helper functions

Create two functions:

- `respondWithJSON(w http.ResponseWriter, status int, payload interface{})`
- `respondWithError(w http.ResponseWriter, code int, msg string)` (which calls `respondWithJSON` with error-specific values)

 They're simply helper functions that write an HTTP response with:

- A status code
- An `application/json` content type
- A JSON body

## 6. Add a readiness handler

Add a handler for `GET /v1/readiness` requests. It should return a `200` status code and a JSON body:

```json
{
  "status": "ok"
}
```

## 7. Add an error handler

Add a handler for `GET /v1/err` requests. It should return a `500` status code and a JSON body:

```json
{
  "error": "Internal Server Error"
}
```

## 8. Run and test your server

```bash
$ go build -o Folder_name && ./Folder_name 
```

Once it's running, use an HTTP client to test your endpoints.



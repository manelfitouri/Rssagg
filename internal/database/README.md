# Create a Feed

An RSS feed is just a URL that points to some XML. Users will be able to add feeds to our database.

1. **Create a feeds table**

   Like any table in our DB, we'll need the standard `id`, `created_at`, and `updated_at` fields. We'll also need a few more:

   - `name`: The name of the feed 
   - `url`: The URL of the feed
   - `user_id`: The ID of the user who added this feed


   Write the appropriate migrations and run them.

2. **Add a query to create a feed**

   Add a new query to create a feed, then use `sqlc generate` to generate the Go code.

3. **Create some authentication middleware**

   Most of the endpoints going forward will require a user to be logged in. we created a middleware that will check for a valid API key.

    I prefer to create custom handlers that accept extra values. 

   - A custom type for handlers that require authentication

     ```go
     type authedHandler func(http.ResponseWriter, *http.Request, database.User)
     ```

   - Middleware that authenticates a request, gets the user and calls the next authed handler

     ```go
     func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
         ///
     }
     ```

   - Using the middleware

     ```go
     v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerUsersGet))
     ```

4. **Create a handler to create a feed**

   Create a handler that creates a feed. This handler and the "get user" handler should use the authentication middleware.

   - **Endpoint:** `POST /v1/feeds`

   - **Example request body:**

     ```json
     {
       "name": "manel",
       "url": "https://Rss/index.xml"
     }
     ```

   - **Example response body:**

     ```json
     {
       "id": "4a82b372-b0e2-45e3-956a-b9b83358f86b",
        "CreatedAt": "2024-04-12T18:57:25.587659Z",
        "UpdatedAt": "2024-04-12T18:57:25.587659Z",
       "name": "manel",
       "url": "https://Rss/index.xml"
       "user_id": "d6962597-f316-4306-a929-fe8c8651671e"
     }
     ```

5. **Test**

   Test your handler using an HTTP client, then use your database client to make sure the data was saved correctly.

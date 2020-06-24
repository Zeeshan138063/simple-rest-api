# Go Basic API Server

### Instructions

Make sure you have Go latest version installed

If something goes right, you should see a log like this:

`2020/05/11 15:29:02 running api server at http://0.0.0.0:5400`

That means the api is running

Here are the routes (you can find them in internal/api/routes.go)

1. `GET     http://0.0.0.0:5400/api/v1/user`
2. `POST    http://0.0.0.0:5400/api/v1/user`
3. `GET     http://0.0.0.0:5400/api/v1/user/{id}`
4. `PUT     http://0.0.0.0:5400/api/v1/user/{id}`
5. `DELETE  http://0.0.0.0:5400/api/v1/user/{id}`

Now open Postman and send request to any of the routes with the following body as an example:

`POST` -> `http://0.0.0.0:5400/api/v1/user`

`{
     "first_name": "John",
     "last_name": "Doe",
     "email": "john.doe@gmail.com"
 }`

That's it

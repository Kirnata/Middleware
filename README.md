# Build echo middleware

### Build a simple middleware using echo framework

A handler which sends how many minutes left until next New Year and response with HTTP 200 OK status code.

Building a middleware, which checks HTTP header `User-Role` presents and contains `admin` and prints `Hey, this is Santa! 🎅 ` to the console (using logrus package) if so.

## Run

```shell
go run github.com/Kirnata/Middleware/cmd/app
```

## Test

```shell
curl --location --request GET '127.0.0.1:8080/status' \
--header 'User-Role: admin'
```

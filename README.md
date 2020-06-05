# JWT-Test
JWT implementation with the repository https://github.com/dgrijalva/jwt-go

In this small repository you can see a general configuration to implement access and refresh JWT. If you want to see where I'm going to be able to implement it

  - https://www.youtube.com/watch?v=c03z2OSbkEo
  - https://levelup.gitconnected.com/crud-restful-api-with-go-gorm-jwt-postgres-mysql-and-testing-460a85ab7121

To make use of this repository, keys must be created at the root of the project. To create it you just have to run the following command:

```sh
$ openssl genrsa -out private.rsa 1024
```

With that we create the private key and it will be called private.rsa

```sh
$ openssl rsa -in private.rsa -pubout > public.rsa.pub
```
This generates the public key that allows validating that the token is valid together with the private key

In the main file, you will find the server startup and the HandleFunc to execute the functions that allow you to check that the JWT is well implemented.

In the drivers folder you will find the functions to log in and the one that will be executed as long as the token is valid

Gorilla Mux was used for the server

- https://www.gorillatoolkit.org/pkg/mux

para descargar todas las dependencias del repositorio ejecute el siguiente comando:

```sh
$ go mod tidy
```

To start the server run:
```sh
$ go run ./main.go
```

Once started the server can test with endpoints:
- /login
- /test-token

To test the login, you need to send a json with username: admin, password: admin. This in turn will generate an access token and a refresh token.

And to prove that the tokens are valid, test the endpoint /test-token placing as header (Authorization) the token generated with the prefix Bearer. Example: "Bearer __your_token__"



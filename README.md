# Golang microservices - API Gateway 

## Stack

 - Go 1.18
 - Docker
 - Docker Compose
 - Pattern: API Gateway

### Important

Docker, docker-compose, curl should be preinstalled to deploy and test application.
Used Docker version 20.10.17 and docker-compose version 1.29.0.

## Microservices entities

 - Authenticator-Microservice
 - User-Microservice
 - Proxy-Microservice

### Architecture

![Architecture](https://github.com/snsvistunov/go-microservices/blob/master/go_microservices.png?raw=true)

### Authenticator-Microservice
Authenticator-Microservice authenticates client HTTP request if the request is valid or not by analyzing the value in the user request header `Username`. If the user with this header is present in the database, then it's a valid request otherwise the user request is not a valid request.

Implement API named `/auth` which will perform the above task. For every authenticated user request, it sends HTTP code `200` to the Proxy-Microservice and for every
unauthenticated request, it sends HTTP status code `401 unauthorized` 

### User-Microservice

Implements:
1. **API**: `/user/profile` **Method**: `GET`, **Secured**: `true`
This API will return the username, dob (date of birth), age, email, and phone number.

**Example**

```
{
    "UserProfile": {
        "username": "Uncle Bob",
        "date_of_birth": "05-12-1952",
        "age": 69,
        "e-mail": "unclebob@clean.architecture",
        "phone_number": "760-474-1420"
    }
}
```

2. **API**: `/microservice/name` **Method**: `GET` , **Secured**: `false`
This API will return the name of microservice (for our use case user-microservice ).

**Example**

```
{
    "MicroserviceName": "user-microservice"
}
```

### Proxy-Microservice

The Proxy-Microservice is the contact point for the client. It does conditional routing based on these services: User-Microservice, Authenticator-Microservice.

1. If authentication is required for the request API, it routes it to Authenticator-Microservice and if the authenticator microservice returns status code 200, it then routes to User-Microservice otherwise it sends the request back to the user as the user request is unauthorized.

2. If the authentication is not required, then it routes directly to User-Microservice. 


## Configuration

1. **Authenticator-Microservice** - HTTP-port listening `8081` by default. You can change cofiguration params in config file `config/envs/app.env`. All testing data saved in directory `data`.

2. **User-Microservice** - HTTP-port listening `8082` by default. You can change cofiguration params in config file `config/envs/app.env`. All testing data saved in directory `data`.   

3. **Proxy-Microservice** - HTTP-port listening `8888` by default. You can change cofiguration params in config file `config/envs/app.env`.

## How to deploy 

Deploy application:

```
docker-compose up --build 
```

It will run Authenticator-Microservice, User-Microservice, Proxy-Microservice.
You can access to the app via Proxy-Microservice:   

http://127.0.0.1:8888/user/profile
http://127.0.0.1:8888/microservice/name

Stop application:

```
docker-compose down
```

## Testing

### Makefile

`Makefile` - contains curl commands to send.

`make test` - send GET requests to retrieve data.

Example:
```
$ make test

```

## Demo Data

All demo data is included in the application code (hardcoded).

The data for the User-Service service is located in the directory `go-user-service/data/users.go`.

```
var Users = map[string]models.User{
	"Uncle Bob": {
		Username:    "Uncle Bob",
		DateOfBirth: "05-12-1952",
		Age:         69,
		Email:       "unclebob@clean.architecture",
		PhoneNumber: "760-474-1420",
	},
	"Gopher": {
		Username:    "Gopher",
		DateOfBirth: "10-11-2009",
		Age:         13,
		Email:       "go@golang.google.com",
		PhoneNumber: "304-786-2880",
	},
	"Deve": {
		Username:    "Deve",
		DateOfBirth: "29-04-2003",
		Age:         19,
		Email:       "deve@example.com",
		PhoneNumber: "304-786-2880",
	},
	"Steve": {
		Username:    "Steve",
		DateOfBirth: "29-04-2000",
		Age:         22,
		Email:       "steve@example.com",
		PhoneNumber: "304-786-2880",
	},
	"Molly": {
		Username:    "Molly",
		DateOfBirth: "29-04-1992",
		Age:         30,
		Email:       "molly@example.com",
		PhoneNumber: "304-786-2880",
	},
}

```

The data for the Auth-Service service is located in the directory `go-auth-service/data/users.go`.

```
var Users = map[string]models.User{
	"Uncle Bob": {
		Username:   "Uncle Bob",
		AuthStatus: true,
	},
	"Gopher": {
		Username:   "Gopher",
		AuthStatus: true,
	},
	"Deve": {
		Username:   "Deve",
		AuthStatus: false,
	},
	"Steve": {
		Username:   "Steve",
		AuthStatus: true,
	},
	"Molly": {
		Username:   "Molly",
		AuthStatus: false,
	},
}

```
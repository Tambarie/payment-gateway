# Payment-gateway-API

Rentals-API is a `RESTful-API` will allow a
merchant to offer a way for their `shoppers` to pay for their `product`.




## Using Payment-API

To use Payment-API, clone this repository and follow one of two methods shown below:

### Using Docker Compose

1. [Install Docker Compose](https://docs.docker.com/compose/install/)
2. Run all containers with `make up`.


### Using Your Local Machine

1. [Install Go](https://golang.org/doc/install)
2. [Install Mongo DB](https://docs.mongodb.com/manual/installation/)
3. Create a database named `payment-gateway`.
4. Install all dependencies using:


```
go mod tidy
```
5. Run the application using:

```
make run
```


## Setting Environmental Variables
An environment variable is a text file containing ``KEY=value`` pairs of your secret keys and other private information. For security purposes, it is ignored using ``.gitignore`` and not committed with the rest of your codebase.

To create, ensure you are in the root directory of the project then on your terminal type:
```
touch .env
```
All the variables used within the project can now be added within the ``.env`` file in the following format:
```
SERVICE_MODE="dev"
DB_TYPE = "mongodb"
MONGO_DB_HOST="localhost"
MONGO_DB_NAME="payment-gateway"
MONGO_DB_PORT="27017"
SERVICE_PORT=8080
DB_PASS=<your db password>
```

### Postman Collection

1. [Get the postman collection link here ](https://www.getpostman.com/collections/db0e41771a301a73740e)


### Postman Collection

1. [Get the postman collection link here ](https://www.getpostman.com/collections/db0e41771a301a73740e)

## Tests
Testing is done using the GoMock framework. The ``gomock`` package and the ``mockgen``code generation tool are used for this purpose.
If you installed the dependencies using the command given above, then the packages would have been installed. Otherwise, installation can be done using the following commands:
```
go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen
```

After installing the packages, run:
```
make mock-service
```

The command above helps to generate mock interfaces from a source file.

To run tests, run:
```
go test
```

## Author

* Emmanuel Gbaragbo ([Tambarie](https://github.com/Tambarie)) 🐛





# Money Tracker GO

## Description

This is a Golang (Gin) BE application that helps you track your spending and understand where your money goes.

## Highlight

1. View Transaction List with Filters
2. Insert Transaction Data
3. View Transaction Summary by Period

## Project Structure

The project is organized as follows:

```
├── assets/ # Contain files for bufio scanner (notably the common password validation)
├── constants/ # Contain shared constant values used across the project
├── docs/ # Contain the project documentations (such as postman and erd)
├── dto/ # Defines data transfer objects
├── handlers/ # Functions to process HTTP requests and responses
├── helpers/ # Utility functions and general purpose code used across the app
├── initializers/ # Initialize components like DB connection
├── middlewares/ # Functions to process requests before they reach handlers
├── migrate/ # Functions that will maintain DB schemas
├── models/ # Describes DB schemas
├── repositories/ # Contains logic for querying and persisting data
├── router/ # Routes setup and handler mappings
├── usecases/ # Specifies application business logic
├── .dockerignore # Excluded files and directories from Docker build context
├── .env.backup # Example environment variables file
├── .gitignore # Files and directories to ignore in Git
├── Dockerfile # Docker image build specifications
├── go.mod # Go modules file
├── go.sum # Go checksum file
├── LICENSE # App license info
├── main.go # App entry point
├── makefile # Command shortcuts
└── README.md # Project documentation
```

## Prerequisites

Before running this project, you need to have the following installed:

- **Go 1.23 or higher**: For running the Go application
- **MySQL**: DB used by the app
- **Docker**: To run the project as this project implements containerization
- **Make (Optional)**: To simplify common tasks by creating shortcuts
- **PlantUML VSCode Extension (Optional)**: To view erd.puml file defined in this project

## How to Run the Project

1. **Clone the repository**:

```
git clone https://github.com/ertantorizkyf/money-tracker-go
```

2. **Navigate into the project folder**:

```
cd money-tracker-go
```

3. **Adjust environment variables**:

- Copy .env.example to .env file and adjust the value to fit your local machine configuration. It will be copied to the docker container

4. **Run the application**:

- Using docker:

```
docker build -t money-tracker-go . && \
	docker rm -f money-tracker-go 2>/dev/null || true && \
	docker run --name money-tracker-go -d -p 3000:3000 --env-file .env money-tracker-go
```

- Using make:

```
make up
```

5. **The application should now be running (at port 3000 by default)**:

- Hit the `/ping` endpoint to test if the app runs properly

## Documentation

1. DB schemas can be viewed under the `erd.puml` file
2. API collection can be viewed under the `money_tracker.postman_collection.json` file

## License

This project is licensed under the MIT License.

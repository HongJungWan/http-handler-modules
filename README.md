# http-handler-modules

# Project Overview

Hi, I'm Hong JungWan, a second-year Go developer. I worked on a boilerplate project for team members who are new to the Go language.

This project demonstrates various aspects of web application development, including creating an HTTP server, using middleware, implementing HTTP endpoints, Test Code and integrating JWT with an RDBMS for authentication and authorization. Below is a detailed overview of what is covered in this project.

<br><br>

# Table of Contents

### Creating an HTTP Server
- Initializing the Project
- Running the Web Server
- Writing Code for Middleware Tests
- Making Ports Configurable

<br>

### Preparing the Development Environment
- Running the Application with Docker
- Writing a Makefile

<br>

### Changing HTTP Server Configuration for Weak Coupling
- Bringing in Settings from Environment Variables
- Handling Signals
- Defining Server Structure
- Defining Routes Separately
- Refactoring the run Function

<br>

### Adding Endpoints
- Implementing entity.Task Type Definition and Temporary Storage
- Implementing Helper Functions
- Implementing Endpoints to Register Tasks
- Implementing Tests Combining Table-Driven Tests and Golden Tests
- Implementing Endpoints to Return Task List
- Defining HTTP Handlers for Routes

<br>

### Implementing Database Processing Using RDBMS
- Setting Up MySQL Execution Environment
- Implementing RDBMS Processing
- Writing Tests for RDBMS-Related Features

<br>

### Dividing HTTP Handlers by Function
- Writing Repository Layer for RDBMS in HTTP Handlers
- Dividing HTTP Handlers
- Automatically Generating Mock Code
- Writing User Registration Features
- Verifying Functionality

<br>

### Using Redis and JWT for Authentication and Authorization
- Functions Implemented in This Chapter
- Preparing Redis
- Preparing JWT Signing
- Writing Access Tokens Using JWT
- Implementing User Login Endpoints
- Implementing Authentication Features Using Middleware
- Implementing Endpoints to Protect Routes with Authentication Information
- Verifying Functionality

<br><br>

# Test Code Run

```go
go test ./...

or

go test -v ./...
```
# http-handler-modules

# Project Overview

Hi, I'm Hong JungWan, a second-year Go developer. I worked on a boilerplate project for team members who are new to the Go language.

This project demonstrates various aspects of web application development, including creating an HTTP server, using middleware, implementing HTTP endpoints, Test Code and integrating JWT with an RDBMS for authentication and authorization. Below is a detailed overview of what is covered in this project.

<br>

# Table of Contents

### Creating an HTTP Server
- Initializing the Project
    - Writing a basic main function
    - Creating a Go project


- Running the Web Server
    - Running the web server


- Writing Code for Middleware Tests
    - Writing simple tests
    - Handling run failures
    - Writing teardown code for errors


- Making Ports Configurable
    - Making port numbers configurable for execution environments

<br>

### Preparing the Development Environment
- Running the Application with Docker
    - Running with auto-reload
    - Configuring for container environment


- Writing a Makefile

<br>

### Changing HTTP Server Configuration for Weak Coupling
- Bringing in Settings from Environment Variables
    - Implementing Config package


- Handling Signals
    - Using signal.NotifyContext to wait for signals


- Defining Server Structure
    - Using NewMux to define routing


- Defining Routes Separately
    - Using httptest package for testing


- Refactoring the run Function
    - Organizing the code

<br>

### Adding Endpoints
- Implementing entity.Task Type Definition and Temporary Storage
    - Temporary storage implementation


- Implementing Helper Functions
    - Implementing helper functions for tests


- Implementing Endpoints to Register Tasks
    - Validating request bodies


- Implementing Tests Combining Table-Driven Tests and Golden Tests
    - Writing tests for endpoints


- Implementing Endpoints to Return Task List


- Defining HTTP Handlers for Routes
    - Setting up unique routes using github.com/go-chi/chi

<br>

### Implementing Database Processing Using RDBMS
- Setting Up MySQL Execution Environment
    - Deciding table definition and migration methods
    - Running MySQL container in workflow


- Implementing RDBMS Processing
    - Using database/sql and github.com/jmoiron/sqlx packages
    - Connecting to the database
    - Writing code to handle task operations


- Writing Tests for RDBMS-Related Features
    - Writing tests for database access

<br>

### Dividing HTTP Handlers by Function
- Writing Repository Layer for RDBMS in HTTP Handlers


- Dividing HTTP Handlers
    - Using go generate command for automatic code generation


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
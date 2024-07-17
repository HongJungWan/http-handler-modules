# http-handler-modules

# Project Overview

Hi, I'm Hong JungWan, a second-year Go developer. I worked on a boilerplate project for team members who are new to the Go language.

This project demonstrates various aspects of web application development, including creating an HTTP server, using middleware, implementing HTTP endpoints, Test Code and integrating JWT with an RDBMS for authentication and authorization. Below is a detailed overview of what is covered in this project.

<br><br>

# Contents

| Section                                         | Core Points |
|-------------------------------------------------|-------------|
| **Creating an HTTP Server**                     | - Running the Web Server and Middleware Tests<br>- Making Ports Configurable |
| **Changing HTTP Server Configuration for Weak Coupling** | - Bringing in Settings from Environment Variables<br>- Defining Server Structure<br>- Defining Routes Separately |
| **Adding Endpoints**                            | - Implementing entity.Task Type Definition and Temporary Storage<br>- Implementing Endpoints to Register Tasks<br>- Implementing Endpoints to Return Task List |
| **Implementing Database Processing Using RDBMS**| - Setting Up MySQL Execution Environment<br>- Implementing RDBMS Processing<br>- Writing Tests for RDBMS-Related Features |
| **Dividing HTTP Handlers by Function**          | - Writing Repository Layer for RDBMS in HTTP Handlers<br>- Dividing HTTP Handlers<br>- Writing User Registration Features |
| **Using Redis and JWT for Authentication and Authorization** | - Preparing Redis<br>- Writing Access Tokens Using JWT<br>- Implementing Authentication Features Using Middleware |

<br><br>

# Test Code Run

```go
go test ./...

or

go test -v ./...
```
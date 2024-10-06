# DistributeX
DistributeX is basically a Distributed File Storage Server built on Golang with Multithreading Concepts

Here the file size more than 1 MB will be splitted and will be stored in parts.

# multi threading - Golang

```markdown
DistributeX - Distributed File Storage Server
```

# 1. Requirement

1. Should use Golang
2. Should use RDB
3. Should use Thread
4. API documentation
5. Should Dockerize
    1. Should be run all the program by one “docker compose up (—build)” command

# 2. Features

1. API 1 - upload file
    1. Separate file to multiple files
    2. Upload separated files to Database in parallel using threads
    3. Return file id.
2. API 2 - get uploaded Files data
3. API 3 - download file by id
    1. Get files parallel using threads and merge back to one file.
    2. Return merged original file

# How to Run 
```
docker compose up --build
```
# API Documentation 
You can find the API Documentation here ```http://localhost:8080/swagger/index.html#/``` 

# Test 
### You can test the application in Three Ways
1. Postman collection file is added to the repository you can import the file and test the API's there.

2. You can use Swagger API Doc to try the API's ```http://localhost:8080/swagger/index.html#/``` 
3. Frontend is build already Open this URL in browser
   ``` http://localhost:3000 ```
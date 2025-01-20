# Simple Web Server in Go

This is a basic web server built using Go's standard `net/http` package. It demonstrates how to handle HTTP requests and serve responses on a specified port.

---

## Features

- Responds to HTTP requests on the root (`/`) endpoint with a simple greeting message.  
- Easily extendable to add more routes and features.

---

## Prerequisites

- **Go (version 1.16 or later):** Ensure Go is installed on your machine.  
  You can download it from [golang.org](https://golang.org).

---

## Running the Web Server

1. Clone or download this repository to your local machine.  
2. Open a terminal in the project directory.  
3. Run the following command to start the server:

   ```bash
   go run webserver.go
   ```

4. The server will start listening on port `8080`. You should see a message like:

   ```
   Server listening to: 8080
   ```

5. Open a browser or use `curl` to test the server:
   - **Browser**: Visit [http://localhost:8080](http://localhost:8080).  
   - **Curl**: Run:
     ```bash
     curl http://localhost:8080
     ```

6. You can also use **Postman** to test the endpoints.

---

## Snapshots

_Attaching screenshots of the working application here for better visualization._

## Creating an user
![image](https://github.com/user-attachments/assets/bbbef2ad-d0dd-4aab-8095-7116b6a589a0)

## Retrieving an existing user
![image](https://github.com/user-attachments/assets/059218f0-9b89-4f4d-98b7-29ea9a901249)

## Deleting an existing user
![image](https://github.com/user-attachments/assets/49e03809-cc55-4f74-99db-e466572e16a1)

## Trying to retrieve an non-existing user
![image](https://github.com/user-attachments/assets/c08c816b-b166-48ce-b253-62930805a8c8)

---

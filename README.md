# Clients API - Go & Gin

A RESTful API developed in **Go (Golang)** for managing client registrations. This project was built to consolidate backend web development concepts, HTTP requests, in-memory data structuring, and local file persistence using the **Gin** framework.

## Technologies Used

* **Language:** Go (Golang)
* **Web Framework:** Gin (`github.com/gin-gonic/gin`)
* **Storage:** Data persistence in a local `.json` file.

## Features (CRUD)

The system allows complete management of a client workflow:
- **C**reate: Registration of new clients.
- **R**ead: Listing of all clients or specific search by ID.
- **U**pdate: Updating data for an existing client.
- **D**elete: Removal of a client from the system.

> **Architecture Note:** To ensure high performance and fast responses, the API works with an in-memory client *slice* for operations, synchronizing the data with the `dados/clients.json` file upon every change (Post, Put, Delete).

## API Endpoints

The API runs by default on port `:8080`. Below are the available routes:

| HTTP Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/clients` | Returns the complete list of registered clients. |
| `GET` | `/clients/:id` | Returns the details of a specific client by ID. |
| `POST` | `/clients` | Creates a new client record. |
| `PUT` | `/clients/:id` | Updates the data of an existing client by ID. |
| `DELETE` | `/clients/:id` | Removes a client from the database by ID. |

Request Body Example (POST / PUT)
```json
  {
    "nome": "John Doe",
    "email": "johndoe@email.com",
    "telefone": "5551234567"
  }
````

### How to Run the Project Locally
Clone the repository to your machine:
`
Bash
git clone [https://github.com/chelzzzs/clientsAPI.git](https://github.com/chelzzzs/clientsAPI.git)`

Enter the project folder:

`Bash
cd clientsAPI
`
Download the module dependencies (Gin Framework):
`Bash
go mod tidy`

Run the application:
`Bash
go run main.go
`

# Author
Michel Zulszeski


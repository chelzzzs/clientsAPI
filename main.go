package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var clients = []models.Client{}

func main() {
	loadClients()
	fmt.Println("Clientes:", clients)

	router := gin.Default()
	router.GET("/clients", getClients)
	router.POST("/clients", postClient)
	router.GET("/clients/:id", getByClientID)
	router.DELETE("/clients/:id", deleteClient)
	router.PUT("/clients/:id", putClient)
	router.Run()

}

func putClient(c *gin.Context) {
	var id string = c.Param("id")
	var updateClient models.Client

	if err := c.BindJSON(&updateClient); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	for i, client := range clients {
		if fmt.Sprint(client.ID) == id {
			client.ID = updateClient.ID
			clients[i] = updateClient
			saveClientes()
			c.JSON(200, gin.H{
				"message":       "Cliente updated sucessfully",
				"id_atualizado": id,
				"new_client":    updateClient,
			})
			return
		}
	}
	c.JSON(404, gin.H{"error": "client not found"})
}

func deleteClient(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	for i, client := range clients {
		if client.ID == id {
			clients = append(clients[:i], clients[i+1:]...)
			saveClientes()

			c.JSON(200, gin.H{"message": "Client deleted successfully"})
			return
		}
	}
	c.JSON(400, gin.H{"error": "Client not found"})

}

func getByClientID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, client := range clients {
		if client.ID == id {
			c.JSON(200, gin.H{"client": client})
			return
		}
	}
	c.JSON(400, gin.H{"error": "Client not found"})
}

func postClient(c *gin.Context) {
	var newClient models.Client
	if err := c.ShouldBindJSON(&newClient); err != nil {
		c.JSON(404, gin.H{"error": "Invalid JSON"})
		return
	}

	newClient.ID = len(clients) + 1
	clients = append(clients, newClient)
	saveClientes()
	c.JSON(201, gin.H{"message": "Client added sucessfully", "client": newClient, "id": newClient.ID})
}

func getClients(c *gin.Context) {
	c.JSON(200, gin.H{
		"clients": clients,
	})

}

func loadClients() {
	file, err := os.Open("dados/clients.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&clients); err != nil {
		fmt.Println("error decoding JSON:", err)
	}
}

func saveClientes() {

	file, err := os.Create("dados/clients.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(clients); err != nil {
		fmt.Println("error encoding JSON", err)
	}

}

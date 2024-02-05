package DB

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

type Item struct {
	Name string `json:"name"`
}

type infoJSON struct {
	Info string `json:"info"`
}

type newItem struct {
	NewItemName        string `json:"NewItemName"`
	NewItemModel       string `json:"NewItemModel"`
	NewItemProdYear    string `json:"NewItemProdYear"`
	NewItemDescription string `json:"NewItemDescription"`
}

var (
	db  *sql.DB
	err error
)

func OpenDatabase() (*sql.DB, error) {

	db, err = sql.Open("mysql", "root:October2022@tcp(localhost:3306)/item_Collection")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to DB.")

	return db, err
}

func GetDB() *sql.DB {
	// Ensure that the database is opened before returning it
	if db == nil {
		fmt.Println("Database is not yet opened. Call OpenDatabase first.")
		return nil
	}

	return db
}

func GetDBItem(context *gin.Context) {

	if GetDB() == nil {
		_, err := OpenDatabase()
		if err != nil {
			return
		}
	}

	result, err := db.Query("SELECT ItemProdYear FROM item")
	if err != nil {
		panic(err.Error())
	}

	var results []string

	for result.Next() {
		var item Item

		err = result.Scan(&item.Name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("Item name: " + item.Name)
		results = append(results, item.Name)
	}

	chosenItemName := results[rand.Intn(len(results))]

	context.JSON(http.StatusOK, gin.H{"message": "Message received.", "itemname": chosenItemName})

}

func GetFrontPageInfo(context *gin.Context) {

	file, err := os.Open("./Assets/info.json")
	if err != nil {
		fmt.Println("Failed to get config file. Error: " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get config file."})
		context.Abort()
		return
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	JSONObject := infoJSON{}
	err = decoder.Decode(&JSONObject)
	if err != nil {
		fmt.Println("Failed to parse config file. Error: " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse config file."})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Message received.", "info": JSONObject.Info})

}

func AddItem(context *gin.Context) {

	if GetDB() == nil {
		_, err := OpenDatabase()
		if err != nil {
			return
		}
	}

	var newItem = newItem{}

	if err := context.ShouldBindJSON(&newItem); err != nil {
		fmt.Println("Invalid post request. Error: " + err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post request."})
		context.Abort()
		return
	}

	// Convert string to integer
	ProdYearInteger, err := strconv.Atoi(newItem.NewItemProdYear)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return
	}

	// Prepare the SQL statement
	query := "INSERT INTO item (ItemName, ItemModel, ItemProdYear, ItemDescription) VALUES (?, ?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Provide values for the placeholders and execute the statement
	_, err = stmt.Exec(newItem.NewItemName, newItem.NewItemModel, ProdYearInteger, newItem.NewItemDescription)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New record inserted successfully!")

}

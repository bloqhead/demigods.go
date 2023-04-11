package handler

import (
	"embed"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type Category struct {
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}

type Scaling struct {
	Str string `json:"str,omitempty"`
	Int string `json:"int,omitempty"`
	Fai string `json:"fai,omitempty"`
	Arc string `json:"arc,omitempty"`
	Dex string `json:"dex,omitempty"`
}

type Stat struct {
	Weight float64 `json:"weight,omitempty"`
	Physical float64 `json:"physical,omitempty"`
	Magic float64 `json:"magic,omitempty"`
	Fire float64 `json:"fire,omitempty"`
	Light float64 `json:"light,omitempty"`
	Holy float64 `json:"holy,omitempty"`
}

type Weapon struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Tier string `json:"tier"`
	Category string `json:"category"`
	Scaling Scaling `json:"scaling"`
	Skill string `json:"skill"`
	Stats Stat `json:"stats"`
}

var Weapons []Weapon
var Categories []Category

//go:embed data.json
//go:embed categories.json

var data embed.FS

func fetchData() {
	content, err := data.ReadFile("data.json")

	if err != nil {
		log.Fatal("Error opening data file: ", err)
	}

	err = json.Unmarshal(content, &Weapons)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
}

// fetch all items
func FetchAll(c *gin.Context) {
	fetchData()
	c.JSON(http.StatusOK, Weapons)
}

// find by ID
func FetchById(c *gin.Context) {
	fetchData()

	id := c.Param("id")
	items := make([]Weapon, 0, len(Weapons))
	
	for _, v := range Weapons {
		if strconv.Itoa(v.ID) == id {
			items = append(items, v)
		}
	}

	if len(items) > 0 {
		c.JSON(http.StatusOK, items)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"status": "nodata",
			"message": "No items found for ID " + id,
		})
	}
}

// find by category
func FetchByCategory(c *gin.Context) {
	fetchData()

	category := c.Param("category")
	items := make([]Weapon, 0, len(Weapons))
	
	for _, v := range Weapons {
		cat := slug.Make(v.Category)
		slug := slug.Make(category)

		if cat == slug {
			items = append(items, v)
		}
	}

	if len(items) > 0 {
		c.JSON(http.StatusOK, items)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"status": "nodata",
			"message": "No items found for category " + category,
		})
	}
}

// fetch category list
func FetchCategories(c *gin.Context) {
	content, err := data.ReadFile("categories.json")

	if err != nil {
		log.Fatal("Error opening data file: ", err)
	}

	err = json.Unmarshal(content, &Categories)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	items := make([]Category, 0, len(Categories))
	
	for _, v := range Categories {
		items = append(items, v)
	}

	if len(items) > 0 {
		c.JSON(http.StatusOK, items)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"status": "nodata",
			"message": "No categories found",
		})
	}
}

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}
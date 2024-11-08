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
	Label string `json:"label"`
	Value string `json:"value"`
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
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tier string `json:"tier,omitempty"`
	Category string `json:"category,omitempty"`
	Scaling Scaling `json:"scaling,omitempty"`
	Skill string `json:"skill,omitempty"`
	Stats Stat `json:"stats,omitempty"`
}

var Weapons []Weapon
var Categories []Category
var EmptyData = make([]string, 0, 0)

// embed our data (DO NOT MODIFY)
//go:embed data/data.json
//go:embed data/categories.json

var data embed.FS

func FetchData() {
	dataset1, err1 := data.ReadFile("data/data.json")

	if err1 != nil {
		log.Fatal("Error opening data file: ", err1)
	}

	err1 = json.Unmarshal(dataset1, &Weapons)

	if err1 != nil {
		log.Fatal("Error during item Unmarshal(): ", err1)
	}

	dataset2, err2 := data.ReadFile("data/categories.json")

	if err2 != nil {
		log.Fatal("Error opening categories file: ", err2)
	}

	err2 = json.Unmarshal(dataset2, &Categories)

	if err2 != nil {
		log.Fatal("Error during categories Unmarshal(): ", err2)
	}
}

// fetch all items
func FetchAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": Weapons,
	})
}

// find by ID
func FetchById(c *gin.Context) {
	id := c.Param("id")
	items := make([]Weapon, 0, len(Weapons))
	
	for _, v := range Weapons {
		if strconv.Itoa(v.ID) == id {
			items = append(items, v)
		}
	}

	if len(items) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": items,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "No items found for ID " + id,
			"data": EmptyData,
		})
	}
}

// find by category
func FetchByCategory(c *gin.Context) {
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
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": items,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"status": "nodata",
			"message": "No items found for category " + category,
			"data": EmptyData,
		})
	}
}

// fetch category list
func FetchCategories(c *gin.Context) {
	items := make([]Category, 0, len(Categories))
	
	for _, v := range Categories {
		items = append(items, v)
	}

	if len(items) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": items,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"status": "nodata",
			"message": "No categories found",
			"data": EmptyData,
		})
	}
}

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Max-Age", "3600")
	c.Next()
}
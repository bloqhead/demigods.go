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

//go:embed data.json
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

	c.JSON(http.StatusOK, items)
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

	c.JSON(http.StatusOK, items)
}

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}
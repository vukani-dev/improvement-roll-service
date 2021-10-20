package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"os"

	"encoding/json"

	"io/ioutil"

	"sort"

	"sync"

	"time"
)

type SharedCategory struct {
	Category Category  `json:"category"`
	Tags     []string  `json:"tags"`
	Author   string    `json:"author"`
	Date     time.Time `json:"date"`
}

type SharedCategoryMem struct {
	sync.RWMutex
	categories []SharedCategory
}

type Category struct {
	Name          string `json:"name"`
	TimeSensitive bool   `json:"timeSensitive"`
	Description   string `json:"description"`
	Tasks         []Task `json:"tasks"`
}

type Task struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	Time        int    `json:"time"`
}

var mem = &SharedCategoryMem{}

func (m *SharedCategoryMem) Get() *SharedCategoryMem {
	m.RLock()
	m.RUnlock()
	return m
}

func (m *SharedCategoryMem) Set(categories []SharedCategory) {
	m.Lock()
	m.categories = categories
	m.Unlock()
}

func main() {
	app := fiber.New()
	mem.Set(initCategories("categories/"))

	app.Get("/", func(c *fiber.Ctx) error {
		var categories SharedCategoryMem
		categories = *mem.Get()
		order := c.Query("order")
		categories.orderCategories(order)

		tags := c.Query("tags")
		if tags != "" {
			categories.filterCategoriesByTag(tags)
		}

		b, err := json.Marshal(categories.categories)
		if err != nil {
			fmt.Println(err)
		}

		return c.SendString(string(b))
	})
	app.Listen(":3000")
}

func (categories *SharedCategoryMem) filterCategoriesByTag(tags string) {
	fmt.Println(tags)
}

func (categories *SharedCategoryMem) orderCategories(order string) {
	sort.Slice(categories.categories, func(i, j int) bool {
		switch order {
		case "desc":
			return categories.categories[i].Date.After(categories.categories[j].Date)
		case "asc":
			return categories.categories[i].Date.Before(categories.categories[j].Date)
		default:
			return categories.categories[i].Date.After(categories.categories[j].Date)
		}
	})
}

func initCategories(dir string) []SharedCategory {
	items, _ := ioutil.ReadDir(dir)
	tmp := make([]SharedCategory, 0, len(items))

	for _, item := range items {
		filePath := dir + item.Name()
		fmt.Println(filePath)
		tmp = append(tmp, parseCategory(filePath, item.Name()))
	}

	return tmp
}

func parseCategory(filePath string, categorName string) SharedCategory {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened " + categorName)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var sharedCategory SharedCategory
	json.Unmarshal(byteValue, &sharedCategory)

	defer jsonFile.Close()
	return sharedCategory
}

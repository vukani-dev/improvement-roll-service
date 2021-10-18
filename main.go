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

func (m *SharedCategoryMem) Get() []SharedCategory {
	m.RLock()
	m.RUnlock()
	return m.categories
}

func (m *SharedCategoryMem) Set(categories []SharedCategory) {
	m.Lock()
	m.categories = categories
	m.Unlock()
}

func main() {
	app := fiber.New()
	mem.Set(initCategories())

	app.Get("/", func(c *fiber.Ctx) error {
		b, err := json.Marshal(mem.Get())
		if err != nil {
			fmt.Println(err)
		}

		return c.SendString(string(b))
	})

	app.Get("/time", func(c *fiber.Ctx) error {
		var categories []SharedCategory
		categories = mem.Get()

		sort.Slice(categories, func(i, j int) bool {
			return categories[i].Date.After(categories[j].Date)
		})

		b, err := json.Marshal(mem.Get())
		if err != nil {
			fmt.Println(err)
		}
		return c.SendString(string(b))
	})

	app.Listen(":3000")
}

func initCategories() []SharedCategory {
	dir := "categories/"
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

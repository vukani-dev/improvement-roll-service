package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"os"

	"encoding/json"

	"io/ioutil"
)

type SharedCategory struct {
	Category Category `json:"category"`
	Tags     []string `json:"tags"`
	Author   string   `json:"author"`
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

func main() {
	app := fiber.New()
	categories := initCategories()
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(categories)
		b, err := json.Marshal(categories)
		if err != nil {
			fmt.Println(err)
		}

		return c.SendString(string(b))
	})

	app.Listen(":3000")
}

func initCategories() CategoryMemory {
	dir := "categories/"
	items, _ := ioutil.ReadDir(dir)
	tmp := make([]SharedCategory, 0, len(items))

	for _, item := range items {
		filePath := dir + item.Name()
		fmt.Println(filePath)
		tmp = append(tmp, parseCategory(filePath, item.Name()))
	}

	return CategoryMemory{Data: tmp}
}

type CategoryMemory struct {
	Data []SharedCategory
}

func NewCategoryMemory() CategoryMemory {
	categories := initCategories()
	return categories
}

func parseCategory(filePath string, categorName string) SharedCategory {
	jsonFile, err := os.Open("categories/test.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened test.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var sharedCategory SharedCategory
	json.Unmarshal(byteValue, &sharedCategory)

	defer jsonFile.Close()
	return sharedCategory
}

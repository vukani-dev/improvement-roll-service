package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
)

type SharedCategory struct {
	Category Category  `json:"category" toml:"category" yaml:"category"`
	Tags     []string  `json:"tags" toml:"tags" yaml:"tags"`
	Author   string    `json:"author" toml:"author" yaml:"author"`
	Date     time.Time `json:"date" toml:"date" yaml:"date"`
}

type SharedCategoryPaged struct {
	SharedCategories []SharedCategory `json:"sharedCategories"`
	Page             int              `json:"page"`
	TotalPages       int              `json:"totalPages"`
}

type SharedCategoryMem struct {
	sync.RWMutex
	categories []SharedCategory
}

type Category struct {
	Name          string `json:"name" toml:"name" yaml:"name"`
	TimeSensitive bool   `json:"timeSensitive" toml:"time_sensitive" yaml:"timeSensitive"`
	Description   string `json:"description" toml:"description" yaml:"description"`
	Tasks         []Task `json:"tasks" toml:"tasks" yaml:"tasks"`
}

type Task struct {
	Name        string `json:"name" toml:"name" yaml:"name"`
	Description string `json:"desc" toml:"desc" yaml:"desc"`
	Minutes     int    `json:"minutes" toml:"minutes" yaml:"minutes"`
}

var mem = &SharedCategoryMem{}

func (m *SharedCategoryMem) Get() *SharedCategoryMem {
	m.RLock()
	valSharedMem := new(SharedCategoryMem)
	valSharedMem.categories = make([]SharedCategory, len(m.categories))
	copy(valSharedMem.categories, m.categories)
	m.RUnlock()
	return valSharedMem
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
		categories := *mem.Get()

		order := c.Query("order")
		categories.orderCategories(order)

		searchQuery := c.Query("search")
		if searchQuery != "" {
			categories.filterCategoriesByName(searchQuery)
		}

		tags := c.Query("tags")
		if tags != "" {
			categories.filterCategoriesByTag(tags)
		}

		author := c.Query("author")
		if author != "" {
			categories.filterCategoriesByAuthor(author)
		}

		pageStr := c.Query("page")
		page := tryParsePageToInt(pageStr)
		pagedCategories := categories.getPage(page)

		b, err := json.Marshal(pagedCategories)
		if err != nil {
			fmt.Println(err)
		}

		return c.SendString(string(b))
	})
	app.Listen(":3000")
}

func tryParsePageToInt(str string) int {
	intVar, err := strconv.Atoi(str)
	if err != nil {
		return 1
	}
	return intVar
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

func (categories *SharedCategoryMem) filterCategoriesByTag(tags string) {
	tagSlice := strings.Split(tags, ",")
	n := 0
	for _, category := range categories.categories {
		tagIncluded := false
		for _, tag := range tagSlice {
			tagIncluded = containsTag(category.Tags, tag)
			if containsTag(category.Tags, tag) {
				tagIncluded = true
				break
			}
		}
		if tagIncluded {
			categories.categories[n] = category
			n++
		}

	}
	categories.categories = categories.categories[:n]
}

func containsTag(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(a, e) {
			return true
		}
	}
	return false
}

func (categories *SharedCategoryMem) filterCategoriesByName(search string) {
	n := 0
	for _, category := range categories.categories {
		lowerName := strings.ToLower(category.Category.Name)
		if strings.Contains(lowerName, search) {
			categories.categories[n] = category
			n++
		}
	}
	categories.categories = categories.categories[:n]
}

func (categories *SharedCategoryMem) filterCategoriesByAuthor(author string) {
	n := 0
	for _, category := range categories.categories {
		lowerName := strings.ToLower(category.Author)
		if strings.Contains(lowerName, author) {
			categories.categories[n] = category
			n++
		}
	}
	categories.categories = categories.categories[:n]
}

func (categories *SharedCategoryMem) getPage(pageNum int) SharedCategoryPaged {
	pageSize := 10
	indexStart := pageSize * (pageNum - 1)

	categoriesSize := len(categories.categories)

	categoriesPaged := SharedCategoryPaged{}
	categoriesPaged.Page = pageNum
	categoriesPaged.TotalPages = (((categoriesSize - 1) / pageSize) + 1)

	for i := indexStart; len(categoriesPaged.SharedCategories) < 10; i++ {
		if i >= categoriesSize {
			break
		}
		categoriesPaged.SharedCategories = append(categoriesPaged.SharedCategories, categories.categories[i])
	}
	return categoriesPaged
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
	fileExtension := filepath.Ext(filePath)
	if fileExtension == ".toml" {
		tomlFile, err := toml.LoadFile(filePath)
		if err != nil {
			fmt.Println(err)
			return SharedCategory{}
		}
		byteValue, _ := toml.Marshal(tomlFile)
		var sharedCategoriesMap map[string]any
		toml.Unmarshal(byteValue, &sharedCategoriesMap)
		sharedCategoriesMap["date"], _ = time.Parse("01-02-2006", sharedCategoriesMap["date"].(string))
		byteValue, _ = toml.Marshal(sharedCategoriesMap)
		var sharedCategory SharedCategory
		toml.Unmarshal(byteValue, &sharedCategory)
		return sharedCategory
	}
	if fileExtension == ".yaml" {
		yamlFile, err := os.Open(filePath)
		byteValue, _ := ioutil.ReadAll(yamlFile)
		if err != nil {
			fmt.Println(err)
			return SharedCategory{}
		}
		var sharedCategoriesMap map[string]any
		yaml.Unmarshal(byteValue, &sharedCategoriesMap)
		sharedCategoriesMap["date"], _ = time.Parse("01-02-2006", sharedCategoriesMap["date"].(string))
		byteValue, _ = yaml.Marshal(sharedCategoriesMap)
		var sharedCategory SharedCategory
		yaml.Unmarshal(byteValue, &sharedCategory)
		return sharedCategory
	}
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var sharedCategoriesMap map[string]any
	json.Unmarshal(byteValue, &sharedCategoriesMap)
	sharedCategoriesMap["date"], _ = time.Parse("01-02-2006", sharedCategoriesMap["date"].(string))
	byteValue, _ = json.Marshal(sharedCategoriesMap)
	var sharedCategory SharedCategory
	json.Unmarshal(byteValue, &sharedCategory)
	defer jsonFile.Close()
	return sharedCategory
}

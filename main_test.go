package main

import (
	"io/ioutil"

	"testing"
)

func TestInitCategories(t *testing.T) {
	sharedCategories := initCategories("categories/")
	items, _ := ioutil.ReadDir("categories/")

	categoryLength := len(sharedCategories);
	itemLength := len(items);

	// testing to see if parsing is pulling back the same amt of object that are in the dir
	if categoryLength != itemLength {
		t.Errorf("Error parsing categories. Expected %v but got %v", itemLength, categoryLength)
		return
	}
	
	// general checks
	for i := 0; i < len(sharedCategories); i++ {

		// categories should have a name and description
		if sharedCategories[i].Category.Name == "" || sharedCategories[i].Category.Description == ""{
			t.Errorf("Category with filename %q should have a name and description", items[i].Name())	
			return
		}

		// community categories must have at least 5 tasks
		if len(sharedCategories[i].Category.Tasks) < 5 {
			// t.Errorf("Category %q must have at least 5 tasks", sharedCategories[i].Category.Name)	
			// return
		}

		// sharedCategories must have an author
		if sharedCategories[i].Author == "" {
			t.Errorf("Category %q must have an author", sharedCategories[i].Category.Name)	
			return
		}

		// sharedCategories must have a date added
		if sharedCategories[i].Date.IsZero() {
			t.Errorf("Category %q must have a date in MM-DD-YYYY Format. Go to: %q for help", sharedCategories[i].Category.Name, "https://www.unixtimestamp.com/")	
			return
		}

		// sharedCategories must have tags
		if len(sharedCategories[i].Tags) < 1 {
			t.Errorf("Category %q must have at least one tag", sharedCategories[i].Category.Name)	
			return
		}

		// if a category is time sensiitive it needs to have minutes in the tasks
		if sharedCategories[i].Category.TimeSensitive {
			for _, task := range sharedCategories[i].Category.Tasks {
				if task.Minutes < 1 {
					t.Errorf("Category %q is set to be time sensitive yet one of its tasks doesnt have time or the time is set to 0", sharedCategories[i].Category.Name)	
					return
				}
			}
		}
	}

	// making sure no categories have the same name
	for i := 0; i < len(sharedCategories); i++ {
		matches := 0
		for j := 0; j < len(sharedCategories); j++ {
			if sharedCategories[i].Category.Name == sharedCategories[j].Category.Name {
				matches = matches + 1
				if matches > 1 {
					t.Errorf("There are two categories with the same name: %q", sharedCategories[i].Category.Name)	
					return
				}
			}
		}
	}
}

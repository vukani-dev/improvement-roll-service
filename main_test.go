package main

import "testing"

func TestInitCategories(t *testing.T) {
	sharedCategories := initCategories("test_categories/")

	if len(sharedCategories) != 2 {
		t.Errorf("Error grabbing categories from dir. Expected 2 got %q", len(sharedCategories))
	}
	for i := 0; i < len(sharedCategories); i++ {
		if sharedCategories[i].Author != "hero" {
			t.Errorf("Error getting author. Expected hero go %q", sharedCategories[i].Author)
		}
		if len(sharedCategories[i].Category.Tasks) != 3 {
			t.Errorf("Error parsing tasks. Expected 3 got %q", len(sharedCategories[i].Category.Tasks))
		}
		if sharedCategories[i].Category.Tasks[0].Name != "task1" {
			t.Errorf("Error parsing tasks. Expected taskname of task1 got %q", sharedCategories[i].Category.Tasks[0].Name)
		}
	}

}

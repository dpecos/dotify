package tests

import (
	"testing"

	"github.com/dpecos/dotback/models"

	"reflect"
)

func loadConfig(t *testing.T) models.Config {
	var c models.Config
	if err := c.Load("config"); err != nil {
		t.Errorf("Error loading config: %s", err)
	}
	return c
}

func TestRetrieveRecipes(t *testing.T) {
	c := loadConfig(t)

	expected := []string{"ssh", "git", "mercurial", "tmux", "bash", "conky", "nodejs", "vim", "go", "install-packages-deb"}
	recipes := c.RecipeNames()

	if !reflect.DeepEqual(recipes, expected) {
		t.Errorf("Recipe list not matching expected (got %q, expected %q)", recipes, expected)
	}

}
func TestRetrieveActionsInRecipe(t *testing.T) {
	c := loadConfig(t)

	recipe, err := c.GetRecipe("bash")
	if err != nil {
		t.Fatalf("Recipe not found")
	}

	expected := []string{"link", "include"}
	actions := recipe.ActionNames()

	if !reflect.DeepEqual(actions, expected) {
		t.Errorf("Action list not matching expected (got %q, expected %q)", actions, expected)
	}
}

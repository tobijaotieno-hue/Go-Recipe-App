package main

import (
	"fmt"
	"net/http"
)

type Recipe struct {
	ID          int
	Name        string
	Ingredients []string
	Description string
}

var recipes = []Recipe{
	{
		ID:          1,
		Name:        "Pancakes",
		Ingredients: []string{"2 cups flour", "2 eggs", "1 cup milk", "2 tbsp sugar"},
		Description: "Fluffy breakfast pancakes. Mix ingredients and cook on griddle until golden brown.",
	},
	{
		ID:          2,
		Name:        "Pasta Carbonara",
		Ingredients: []string{"400g spaghetti", "200g bacon", "3 eggs", "100g parmesan"},
		Description: "Classic Italian pasta. Cook pasta, fry bacon, mix with eggs and cheese.",
	},
	{
		ID:          3,
		Name:        "Caesar Salad",
		Ingredients: []string{"1 romaine lettuce", "croutons", "parmesan", "caesar dressing"},
		Description: "Fresh and crispy salad. Toss lettuce with dressing, top with croutons and cheese.",
	},
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/recipe", recipeHandler)
	http.HandleFunc("/static/style.css", styleHandler)
	
	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><head><title>Recipe App</title><link rel='stylesheet' href='/static/style.css'></head><body>")
	fmt.Fprintf(w, "<div class='container'>")
	fmt.Fprintf(w, "<h1>Recipe App</h1>")
	fmt.Fprintf(w, "<h2>All Recipes</h2>")
	
	for _, recipe := range recipes {
		fmt.Fprintf(w, "<div class='recipe-card'>")
		fmt.Fprintf(w, "<h3><a href='/recipe?id=%d'>%s</a></h3>", recipe.ID, recipe.Name)
		fmt.Fprintf(w, "<p>%s</p>", recipe.Description)
		fmt.Fprintf(w, "</div>")
	}
	
	fmt.Fprintf(w, "</div></body></html>")
}

func recipeHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	var selectedRecipe *Recipe
	
	for _, recipe := range recipes {
		if fmt.Sprintf("%d", recipe.ID) == idStr {
			selectedRecipe = &recipe
			break
		}
	}
	
	if selectedRecipe == nil {
		http.NotFound(w, r)
		return
	}
	
	fmt.Fprintf(w, "<html><head><title>%s</title><link rel='stylesheet' href='/static/style.css'></head><body>", selectedRecipe.Name)
	fmt.Fprintf(w, "<div class='container'>")
	fmt.Fprintf(w, "<a href='/' class='back-link'>← Back to all recipes</a>")
	fmt.Fprintf(w, "<h1>%s</h1>", selectedRecipe.Name)
	fmt.Fprintf(w, "<h3>Ingredients:</h3><ul>")
	
	for _, ingredient := range selectedRecipe.Ingredients {
		fmt.Fprintf(w, "<li>%s</li>", ingredient)
	}
	
	fmt.Fprintf(w, "</ul><h3>Instructions:</h3>")
	fmt.Fprintf(w, "<p>%s</p>", selectedRecipe.Description)
	fmt.Fprintf(w, "</div></body></html>")
}

func styleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	css := `
body {
	font-family: Arial, sans-serif;
	background-color: #f5f5f5;
	margin: 0;
	padding: 20px;
}

.container {
	max-width: 800px;
	margin: 0 auto;
	background: white;
	padding: 30px;
	border-radius: 8px;
	box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

h1 {
	color: #333;
	border-bottom: 3px solid #4CAF50;
	padding-bottom: 10px;
}

h2 {
	color: #666;
}

.recipe-card {
	margin: 20px 0;
	padding: 20px;
	border: 1px solid #ddd;
	border-radius: 5px;
	transition: box-shadow 0.3s;
}

.recipe-card:hover {
	box-shadow: 0 4px 8px rgba(0,0,0,0.1);
}

.recipe-card h3 {
	margin-top: 0;
}

.recipe-card a {
	color: #4CAF50;
	text-decoration: none;
}

.recipe-card a:hover {
	text-decoration: underline;
}

.back-link {
	display: inline-block;
	margin-bottom: 20px;
	color: #4CAF50;
	text-decoration: none;
}

.back-link:hover {
	text-decoration: underline;
}

ul {
	line-height: 1.8;
}

li {
	margin: 5px 0;
}
`
	fmt.Fprintf(w, css)
}

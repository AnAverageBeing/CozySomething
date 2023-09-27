package main

import (
	"bytes"
	"fmt"
	"html/template"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./web")

	// Define the API endpoint to fetch blog previews as HTML
	app.Get("/api/blog-previews", func(c *fiber.Ctx) error {
		// Generate random blog previews (for demonstration)
		previews := generateRandomBlogPreviews(5) // Change 5 to the desired number of previews

		// Define the template
		tmpl, err := template.ParseFiles("./web/templates/blog-preview.html")
		if err != nil {
			return err
		}

		// Create a buffer to store the rendered template
		var buffer bytes.Buffer

		// Render the template with the data and write it to the buffer
		err = tmpl.Execute(&buffer, struct{ BlogPreviews []BlogPreview }{BlogPreviews: previews})
		if err != nil {
			return err
		}

		// Send the contents of the buffer as the response
		return c.SendString(buffer.String())
	})

	app.Listen(":3000")
}

type BlogPreview struct {
	BlogPageUrl  string `json:"blogPageUrl"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

// Function to generate random blog previews (for demonstration)
func generateRandomBlogPreviews(count int) []BlogPreview {
	var previews []BlogPreview

	for i := 0; i < count; i++ {
		preview := BlogPreview{
			BlogPageUrl:  fmt.Sprintf("/blog/%d", i+1),
			ThumbnailUrl: "https://cataas.com/cat",
			Title:        getRandomTitle(),
			Description:  getRandomDescription(),
		}
		previews = append(previews, preview)
	}

	return previews
}

// Function to generate random titles (for demonstration)
func getRandomTitle() string {
	titles := []string{
		"Exciting News",
		"Tips and Tricks",
		"Product Update",
		"New Release",
		"Latest Trends",
	}

	return titles[rand.Intn(len(titles))]
}

// Function to generate random descriptions (for demonstration)
func getRandomDescription() string {
	descriptions := []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		"Sed eget mauris ac urna fermentum condimentum.",
		"Vestibulum at rhoncus odio. Nullam bibendum varius ex.",
		"Proin non felis a turpis venenatis fringilla.",
		"Etiam sit amet turpis et enim commodo varius a id nisi.",
	}

	return descriptions[rand.Intn(len(descriptions))]
}

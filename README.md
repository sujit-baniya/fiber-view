# Installation

> go get -u github.com/sujit-baniya/fiber-view



## Usage

### Using Default instance

```go
package main

import (
	"github.com/gofiber/fiber/v2"
	view "github.com/sujit-baniya/fiber-view"
	"log"
)

func main() {
	view.Default(view.Config{
		Path:       "./",
		Extension:  ".html",
		Global: []string{"auth"},
	})
	app := fiber.New(fiber.Config{
		Views: view.Template(),
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return view.Render(c, "index", fiber.Map{
			"auth": "Hello",
		})
	})
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("auth", "World")
		return c.Next()
	})
	app.Get("/world", func(c *fiber.Ctx) error {
		return view.Render(c, "index", fiber.Map{})
	})
	log.Fatal(app.Listen(":8080"))
}
```

### Using individual instance

```go
package main

import (
	"github.com/gofiber/fiber/v2"
	view "github.com/sujit-baniya/fiber-view"
	"log"
)

func main() {
	vw := view.New(view.Config{
		Path:       "./",
		Extension:  ".html",
		Global: []string{"auth"},
	})
	app := fiber.New(fiber.Config{
		Views: vw.Template(),
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return vw.Render(c, "index", fiber.Map{
			"auth": "Hello",
		})
	})
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("auth", "World")
		return c.Next()
	})
	app.Get("/world", func(c *fiber.Ctx) error {
		return vw.Render(c, "index", fiber.Map{})
	})
	log.Fatal(app.Listen(":8080"))
}

```

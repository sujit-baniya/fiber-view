package view

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"io/fs"
	"net/http"
)

type Config struct {
	Path       string `yaml:"path" env-default:"web/views"`
	EmbeddedFS embed.FS
	Extension  string `yaml:"extension" env-default:".html"`
	Global     []string
}

type View struct {
	engine *html.Engine
	global []string
}

func (v *View) Append(c *fiber.Ctx, data fiber.Map) fiber.Map {
	if v.global == nil {
		v.global = []string{}
	}
	for _, global := range v.global {
		dt := c.Locals(global)
		if dt != nil {
			data[global] = dt
		}
	}
	return data
}

func (v *View) Render(c *fiber.Ctx, viewFile string, data fiber.Map, layout ...string) error {
	return c.Render(viewFile, v.Append(c, data), layout...)
}

func (v *View) Template() *html.Engine {
	return v.engine
}

var DefaultView *View

func DefaultFS(cfg Config) {
	DefaultView = NewFS(cfg)
}

func Default(cfg Config) {
	DefaultView = New(cfg)
}

func NewFS(cfg Config) *View {
	d, _ := fs.Sub(cfg.EmbeddedFS, cfg.Path)
	view := &View{
		engine: html.NewFileSystem(http.FS(d), cfg.Extension),
		global: cfg.Global,
	}
	return view
}

func New(cfg Config) *View {
	view := &View{
		engine: html.New(cfg.Path, cfg.Extension),
		global: cfg.Global,
	}
	return view
}

func Render(c *fiber.Ctx, viewFile string, data fiber.Map, layout ...string) error {
	if err := DefaultView.Render(c, viewFile, data, layout...); err != nil { //nolint:wsl
		if err := DefaultView.Render(c, "errors/404", data, layout...); err != nil { //nolint:wsl
			panic(err.Error())
		}
	}
	return nil
}

func Template() *html.Engine {
	return DefaultView.engine
}

func Append(c *fiber.Ctx, data fiber.Map) fiber.Map {
	return DefaultView.Append(c, data)
}

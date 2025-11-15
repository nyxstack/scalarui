# ScalarUI Go Package

A Go package for embedding and serving [Scalar UI](https://scalar.com/) with custom configurations. This package renders the full Scalar UI HTML template directly in Go and supports both URL-based and embedded OpenAPI specs.

## Features

* 🎨 **Beautiful UI** – Modern API documentation
* ⚡ **Easy to Use** – Render a full HTML page with one call
* 🎯 **Fully Configurable** – Themes, layout, auth, variables, CSS, servers, etc.
* 📦 **Embedded Template** – No external HTML required
* 🔗 **Supports URL + Inline Specs**
* 🔥 **Optional Hot Reload** – Auto-refresh when your spec changes
* 🛠️ **Developer Tools** – Built-in “Try It” request interface

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/nyxstack/scalarui"
)

func main() {
    config := scalarui.NewConfig().
        WithTitle("My API Documentation").
        WithURL("http://localhost:8080/openapi.yaml")

    ui := scalarui.New(config)

    http.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
        html, _ := ui.Render()
        w.Header().Set("Content-Type", "text/html")
        fmt.Fprint(w, html)
    })

    http.HandleFunc("/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "openapi.yaml")
    })

    log.Println("Docs at http://localhost:8080/docs")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Configuration Options

### Basic Configuration

```go
config := scalarui.NewConfig().
    WithTitle("My API").
    WithDescription("API documentation").
    WithURL("http://localhost:8080/openapi.yaml")
```

### Theming & Styling

```go
config := scalarui.NewConfig().
    WithTheme("purple").
    WithDarkMode(true).
    WithCustomCSS(`.scalar-app { --scalar-color-accent: #7c3aed }`).
    WithVariable("primary-color", "#7c3aed")
```

### UI Customization

```go
config.
    WithSidebar(true).
    WithDeveloperTools(true). // always | never | localhost
    WithInteractive(true).
    HideHTTPMethods().
    HideModelsSection().
    HideDownload()
```

### Authentication

```go
config.WithAuthentication(map[string]any{
    "bearerAuth": map[string]any{
        "type":  "bearer",
        "token": "your-token",
    },
})
```

### Multiple Servers

```go
config.
    WithServer("https://api.prod.com", "Production").
    WithServer("https://api.staging.com", "Staging")
```

### Embedded Spec

```go
config.WithContent(`{"openapi":"3.0.0","info":{"title":"X"}}`)
```

## Hot Reload (Optional)

### 1. Add the endpoint

```go
var startTime = time.Now()

func hotReload(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    fmt.Fprintf(w, "%d", startTime.Unix())
}
```

### 2. Register it

```go
http.HandleFunc("/hot-reload", hotReload)
```

### 3. Set in config

```go
config.HotReloadURL = "/hot-reload"
```

The UI will auto-refresh when the timestamp changes.

## API

### Core Methods

* `NewConfig()`
* `New(config)`
* `NewWithDefaults()`
* `Render()`

---

## Framework Integration

### Gin

```go
r.GET("/docs", func(c *gin.Context) {
    out, _ := ui.Render()
    c.Header("Content-Type", "text/html")
    c.String(200, out)
})
```

### Echo

```go
e.GET("/docs", func(c echo.Context) error {
    out, _ := ui.Render()
    return c.HTML(200, out)
})
```

---

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
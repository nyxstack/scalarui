package scalarui

// Config represents the complete Scalar Universal Configuration
type Config struct {
	/* ------------------------------------------------------------- */
	/* Basic Configuration */
	/* ------------------------------------------------------------- */

	URL          string      `json:"url,omitempty"`          // URL to the OpenAPI/Swagger document
	Content      interface{} `json:"content,omitempty"`      // Direct OpenAPI/Swagger content (YAML/JSON/string/object)
	ProxyURL     string      `json:"proxyUrl,omitempty"`     // CORS proxy URL used for fetching specs
	HotReloadURL string      `json:"hotReloadUrl,omitempty"` // Hot-reload URL for development
	/* ------------------------------------------------------------- */
	/* Display Options */
	/* ------------------------------------------------------------- */

	Title       string `json:"title,omitempty"`       // Document/page title
	Description string `json:"description,omitempty"` // Document/page description
	Favicon     string `json:"favicon,omitempty"`     // Favicon URL

	/* ------------------------------------------------------------- */
	/* Theme & Layout */
	/* ------------------------------------------------------------- */

	Theme              string            `json:"theme,omitempty"`              // Theme name (default, alternate, moon, purple, etc.)
	Layout             string            `json:"layout,omitempty"`             // Layout type (modern, classic)
	DarkMode           bool              `json:"darkMode,omitempty"`           // Enable dark mode
	ForceDarkModeState string            `json:"forceDarkModeState,omitempty"` // Force dark/light/system mode
	CustomCSS          string            `json:"customCss,omitempty"`          // Custom CSS injected into UI
	Variables          map[string]string `json:"variables,omitempty"`          // CSS custom properties
	WithDefaultFonts   bool              `json:"withDefaultFonts,omitempty"`   // Load Scalar's default web fonts

	/* ------------------------------------------------------------- */
	/* Sidebar / Visibility Toggles */
	/* ------------------------------------------------------------- */

	ShowSidebar             bool   `json:"showSidebar,omitempty"`                  // Show/hide sidebar
	HideMethods             bool   `json:"hideMethods,omitempty"`                  // Hide HTTP method badges
	HideModels              bool   `json:"hideModels,omitempty"`                   // Hide models panel
	HideSearch              bool   `json:"hideSearch,omitempty"`                   // Hide search bar
	HideTestRequestButton   bool   `json:"hideTestRequestButton,omitempty"`        // Hide "Test Request" button
	HideClientButton        bool   `json:"hideClientButton,omitempty"`             // Hide "Client SDK" button
	HideDownloadButton      bool   `json:"hideDownloadButton,omitempty"`           // Hide "Download OpenAPI" button
	HideDarkModeToggle      bool   `json:"hideDarkModeToggle,omitempty"`           // Hide dark mode toggle
	DocumentDownloadType    string `json:"documentDownloadType,omitempty"`         // yaml|json|both|none|direct
	OperationTitleSource    string `json:"operationTitleSource,omitempty"`         // summary|id
	OrderRequiredFirst      bool   `json:"orderRequiredPropertiesFirst,omitempty"` // Required props first
	OrderSchemaPropertiesBy string `json:"orderSchemaPropertiesBy,omitempty"`      // alpha|custom

	/* ------------------------------------------------------------- */
	/* Auto Expansion */
	/* ------------------------------------------------------------- */

	ExpandAllResponses     bool `json:"expandAllResponses,omitempty"`     // Expand all responses
	ExpandAllModelSections bool `json:"expandAllModelSections,omitempty"` // Expand all models
	DefaultOpenAllTags     bool `json:"defaultOpenAllTags,omitempty"`     // Expand all tag groups

	/* ------------------------------------------------------------- */
	/* Developer Tools */
	/* ------------------------------------------------------------- */

	ShowDeveloperTools string `json:"showDeveloperTools,omitempty"` // always|never|localhost
	Interactive        bool   `json:"interactive,omitempty"`        // Enable "Try It" features

	/* ------------------------------------------------------------- */
	/* Authentication */
	/* ------------------------------------------------------------- */

	PersistAuth     bool                   `json:"persistAuth,omitempty"`     // Persist auth information
	Authentication  map[string]interface{} `json:"authentication,omitempty"`  // Auth schemes
	WithCredentials bool                   `json:"withCredentials,omitempty"` // Send cookies on requests

	/* ------------------------------------------------------------- */
	/* Servers / Path Routing */
	/* ------------------------------------------------------------- */

	Servers       []Server               `json:"servers,omitempty"`       // Server overrides
	PathRouting   map[string]interface{} `json:"pathRouting,omitempty"`   // Base path routing config
	BaseServerURL string                 `json:"baseServerURL,omitempty"` // Override server base URL

	/* ------------------------------------------------------------- */
	/* Multi-Document Support */
	/* ------------------------------------------------------------- */

	Sources []SourceConfig `json:"sources,omitempty"` // Multiple OpenAPI docs

	/* ------------------------------------------------------------- */
	/* Metadata */
	/* ------------------------------------------------------------- */

	MetaData map[string]interface{} `json:"metaData,omitempty"` // title, description, og:image, etc.

	/* ------------------------------------------------------------- */
	/* Client Configuration */
	/* ------------------------------------------------------------- */

	DefaultHttpClient map[string]interface{} `json:"defaultHttpClient,omitempty"` // HTTP client defaults
	HiddenClients     interface{}            `json:"hiddenClients,omitempty"`     // Hide specific clients

	/* ------------------------------------------------------------- */
	/* Telemetry / Loading State */
	/* ------------------------------------------------------------- */

	Telemetry bool `json:"telemetry,omitempty"` // Anonymous usage telemetry
	IsLoading bool `json:"isLoading,omitempty"` // Force UI into loading state

	/* ------------------------------------------------------------- */
	/* Callback Hooks (JS functions encoded as interface{}) */
	/* ------------------------------------------------------------- */

	OnSpecUpdate     interface{} `json:"onSpecUpdate,omitempty"`     // When spec updates
	OnLoaded         interface{} `json:"onLoaded,omitempty"`         // UI loaded
	OnBeforeRequest  interface{} `json:"onBeforeRequest,omitempty"`  // Before sending Try-It request
	OnRequestSent    interface{} `json:"onRequestSent,omitempty"`    // After Try-It request
	OnDocumentSelect interface{} `json:"onDocumentSelect,omitempty"` // Selecting doc in multi-doc mode
	OnServerChange   interface{} `json:"onServerChange,omitempty"`   // When server dropdown changes
	OnShowMore       interface{} `json:"onShowMore,omitempty"`       // "Show more" click
	OnSidebarClick   interface{} `json:"onSidebarClick,omitempty"`   // Sidebar interaction

	/* ------------------------------------------------------------- */
	/* Slug Generators */
	/* ------------------------------------------------------------- */

	GenerateHeadingSlug   interface{} `json:"generateHeadingSlug,omitempty"`   // Custom slug function
	GenerateModelSlug     interface{} `json:"generateModelSlug,omitempty"`     // Custom model slug
	GenerateOperationSlug interface{} `json:"generateOperationSlug,omitempty"` // Custom op slug
	GenerateTagSlug       interface{} `json:"generateTagSlug,omitempty"`       // Custom tag slug
	GenerateWebhookSlug   interface{} `json:"generateWebhookSlug,omitempty"`   // Custom webhook slug

	/* ------------------------------------------------------------- */
	/* Sorting Options */
	/* ------------------------------------------------------------- */

	TagsSorter       interface{} `json:"tagsSorter,omitempty"`       // Sort tag groups
	OperationsSorter interface{} `json:"operationsSorter,omitempty"` // Sort operations within tags

	/* ------------------------------------------------------------- */
	/* Redirect */
	/* ------------------------------------------------------------- */

	Redirect interface{} `json:"redirect,omitempty"` // Redirect rules

	/* ------------------------------------------------------------- */
	/* Plugins */
	/* ------------------------------------------------------------- */

	Plugins []Plugin `json:"plugins,omitempty"` // Plugin definitions
}

/* ------------------------------------------------------------- */
/* Server Configuration */
/* ------------------------------------------------------------- */

type Server struct {
	URL         string                 `json:"url"`                   // Server URL
	Description string                 `json:"description,omitempty"` // Description of server
	Variables   map[string]interface{} `json:"variables,omitempty"`   // Server variables
}

/* ------------------------------------------------------------- */
/* Multi-Document Source Configuration */
/* ------------------------------------------------------------- */

type SourceConfig struct {
	Title   string      `json:"title,omitempty"`   // UI title for this document
	Slug    string      `json:"slug,omitempty"`    // URL slug
	URL     string      `json:"url,omitempty"`     // URL to spec
	Content interface{} `json:"content,omitempty"` // Spec inline
	Default bool        `json:"default,omitempty"` // Default selected doc
}

/* ------------------------------------------------------------- */
/* Plugins */
/* ------------------------------------------------------------- */

type Plugin struct {
	Name    string                 `json:"name"`              // Plugin name
	Config  map[string]interface{} `json:"config,omitempty"`  // Plugin-specific config
	Enabled bool                   `json:"enabled,omitempty"` // Enabled state
}

// NewConfig creates a new config with sensible defaults
func NewConfig() *Config {
	return &Config{
		Theme:              "default",
		Layout:             "modern",
		ShowSidebar:        true,
		ShowDeveloperTools: "always",
		Interactive:        true,
		ProxyURL:           "https://proxy.scalar.com",
		Variables:          make(map[string]string),
		Authentication:     make(map[string]interface{}),
		MetaData:           make(map[string]interface{}),
		DefaultHttpClient:  make(map[string]interface{}),
	}
}

// WithURL sets the OpenAPI document URL
func (c *Config) WithURL(url string) *Config {
	c.URL = url
	return c
}

// WithContent sets the OpenAPI document content directly
func (c *Config) WithContent(content string) *Config {
	c.Content = content
	return c
}

// WithTitle sets the page title
func (c *Config) WithTitle(title string) *Config {
	c.Title = title
	return c
}

// WithDescription sets the page description
func (c *Config) WithDescription(description string) *Config {
	c.Description = description
	return c
}

// WithTheme sets the theme
func (c *Config) WithTheme(theme string) *Config {
	c.Theme = theme
	return c
}

// WithDarkMode enables or disables dark mode
func (c *Config) WithDarkMode(enabled bool) *Config {
	c.DarkMode = enabled
	return c
}

// WithProxyURL sets the proxy URL
func (c *Config) WithProxyURL(proxyURL string) *Config {
	c.ProxyURL = proxyURL
	return c
}

// WithCustomCSS sets custom CSS
func (c *Config) WithCustomCSS(css string) *Config {
	c.CustomCSS = css
	return c
}

// WithVariable sets a CSS custom property
func (c *Config) WithVariable(name, value string) *Config {
	if c.Variables == nil {
		c.Variables = make(map[string]string)
	}
	c.Variables[name] = value
	return c
}

// WithServer adds a server configuration
func (c *Config) WithServer(url, description string) *Config {
	c.Servers = append(c.Servers, Server{
		URL:         url,
		Description: description,
	})
	return c
}

// WithAuthentication sets authentication configuration
func (c *Config) WithAuthentication(auth map[string]interface{}) *Config {
	c.Authentication = auth
	return c
}

// WithInteractive enables or disables interactive mode
func (c *Config) WithInteractive(interactive bool) *Config {
	c.Interactive = interactive
	return c
}

// WithSidebar shows or hides the sidebar
func (c *Config) WithSidebar(show bool) *Config {
	c.ShowSidebar = show
	return c
}

// WithDeveloperTools shows or hides developer tools
func (c *Config) WithDeveloperTools(show bool) *Config {
	if show {
		c.ShowDeveloperTools = "always"
	} else {
		c.ShowDeveloperTools = "never"
	}
	return c
}

// HideHTTPMethods hides HTTP methods in sidebar
func (c *Config) HideHTTPMethods() *Config {
	c.HideMethods = true
	return c
}

// HideModelsSection hides the models section
func (c *Config) HideModelsSection() *Config {
	c.HideModels = true
	return c
}

// HideDownload hides the download button
func (c *Config) HideDownload() *Config {
	c.HideDownloadButton = true
	return c
}

// WithForceDarkModeState forces dark/light mode
func (c *Config) WithForceDarkModeState(state string) *Config {
	c.ForceDarkModeState = state
	return c
}

// WithHideSearch hides the search bar
func (c *Config) WithHideSearch(hide bool) *Config {
	c.HideSearch = hide
	return c
}

// WithHideTestRequestButton hides the "Test Request" button
func (c *Config) WithHideTestRequestButton(hide bool) *Config {
	c.HideTestRequestButton = hide
	return c
}

// WithHideClientButton hides the Client SDK button
func (c *Config) WithHideClientButton(hide bool) *Config {
	c.HideClientButton = hide
	return c
}

// WithHideDarkModeToggle hides the dark mode toggle
func (c *Config) WithHideDarkModeToggle(hide bool) *Config {
	c.HideDarkModeToggle = hide
	return c
}

// WithDocumentDownloadType sets yaml/json/both/none/direct
func (c *Config) WithDocumentDownloadType(t string) *Config {
	c.DocumentDownloadType = t
	return c
}

// WithOperationTitleSource sets summary|id
func (c *Config) WithOperationTitleSource(src string) *Config {
	c.OperationTitleSource = src
	return c
}

// WithOrderRequiredFirst enables required-properties-first ordering
func (c *Config) WithOrderRequiredFirst(enabled bool) *Config {
	c.OrderRequiredFirst = enabled
	return c
}

// WithOrderSchemaPropertiesBy sets alpha|custom|preserve
func (c *Config) WithOrderSchemaPropertiesBy(mode string) *Config {
	c.OrderSchemaPropertiesBy = mode
	return c
}

// WithExpandAllResponses expands all responses by default
func (c *Config) WithExpandAllResponses(enabled bool) *Config {
	c.ExpandAllResponses = enabled
	return c
}

// WithExpandAllModelSections expands all models
func (c *Config) WithExpandAllModelSections(enabled bool) *Config {
	c.ExpandAllModelSections = enabled
	return c
}

// WithDefaultOpenAllTags expands all tag groups
func (c *Config) WithDefaultOpenAllTags(enabled bool) *Config {
	c.DefaultOpenAllTags = enabled
	return c
}

// WithPersistAuth enables localStorage auth persistence
func (c *Config) WithPersistAuth(enabled bool) *Config {
	c.PersistAuth = enabled
	return c
}

// WithWithCredentials sets credentials=true for fetch
func (c *Config) WithWithCredentials(enabled bool) *Config {
	c.WithCredentials = enabled
	return c
}

// WithBaseServerURL overrides base server prefix
func (c *Config) WithBaseServerURL(url string) *Config {
	c.BaseServerURL = url
	return c
}

// WithPathRouting sets path routing config
func (c *Config) WithPathRouting(cfg map[string]interface{}) *Config {
	c.PathRouting = cfg
	return c
}

// WithSource adds a multi-document source
func (c *Config) WithSource(src SourceConfig) *Config {
	c.Sources = append(c.Sources, src)
	return c
}

// WithMetaData sets meta tag configuration
func (c *Config) WithMetaData(md map[string]interface{}) *Config {
	c.MetaData = md
	return c
}

// WithDefaultHttpClient sets default HTTP client state
func (c *Config) WithDefaultHttpClient(cfg map[string]interface{}) *Config {
	c.DefaultHttpClient = cfg
	return c
}

// WithHiddenClients configures excluded HTTP clients
func (c *Config) WithHiddenClients(h interface{}) *Config {
	c.HiddenClients = h
	return c
}

// WithTelemetry enables/disables telemetry
func (c *Config) WithTelemetry(enabled bool) *Config {
	c.Telemetry = enabled
	return c
}

// WithIsLoading forces loading state
func (c *Config) WithIsLoading(loading bool) *Config {
	c.IsLoading = loading
	return c
}

// --- Callbacks -------------------------------------------------

func (c *Config) WithOnSpecUpdate(fn interface{}) *Config {
	c.OnSpecUpdate = fn
	return c
}

func (c *Config) WithOnLoaded(fn interface{}) *Config {
	c.OnLoaded = fn
	return c
}

func (c *Config) WithOnBeforeRequest(fn interface{}) *Config {
	c.OnBeforeRequest = fn
	return c
}

func (c *Config) WithOnRequestSent(fn interface{}) *Config {
	c.OnRequestSent = fn
	return c
}

func (c *Config) WithOnDocumentSelect(fn interface{}) *Config {
	c.OnDocumentSelect = fn
	return c
}

func (c *Config) WithOnServerChange(fn interface{}) *Config {
	c.OnServerChange = fn
	return c
}

func (c *Config) WithOnShowMore(fn interface{}) *Config {
	c.OnShowMore = fn
	return c
}

func (c *Config) WithOnSidebarClick(fn interface{}) *Config {
	c.OnSidebarClick = fn
	return c
}

// --- Slug Generators ------------------------------------------

func (c *Config) WithGenerateHeadingSlug(fn interface{}) *Config {
	c.GenerateHeadingSlug = fn
	return c
}

func (c *Config) WithGenerateModelSlug(fn interface{}) *Config {
	c.GenerateModelSlug = fn
	return c
}

func (c *Config) WithGenerateOperationSlug(fn interface{}) *Config {
	c.GenerateOperationSlug = fn
	return c
}

func (c *Config) WithGenerateTagSlug(fn interface{}) *Config {
	c.GenerateTagSlug = fn
	return c
}

func (c *Config) WithGenerateWebhookSlug(fn interface{}) *Config {
	c.GenerateWebhookSlug = fn
	return c
}

// --- Sorting ---------------------------------------------------

func (c *Config) WithTagsSorter(fn interface{}) *Config {
	c.TagsSorter = fn
	return c
}

func (c *Config) WithOperationsSorter(fn interface{}) *Config {
	c.OperationsSorter = fn
	return c
}

// --- Redirect --------------------------------------------------

func (c *Config) WithRedirect(fn interface{}) *Config {
	c.Redirect = fn
	return c
}

// --- Plugins ---------------------------------------------------

func (c *Config) WithPlugin(p Plugin) *Config {
	c.Plugins = append(c.Plugins, p)
	return c
}

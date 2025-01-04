package pesto

//////////////////////////////////////
/// PestoContentType(s)
//////////////////////////////////////
// PestoContentType -
type PestoContentType struct {
	ID                     string `json:"_id"`
	Name                   string `json:"name"`
	Project_id             string `json:"project_id"`
	Frontmatter_definition string `json:"frontmatter_definition"`
	Description            string `json:"description"`
	// Ingredient  []CoffeeIngredient `json:"ingredients"`
}
type CreatePestoContentTypePayload struct {
	// ID                     string `json:"_id"`
	Name                   string `json:"name"`
	Project_id             string `json:"project_id"`
	Frontmatter_definition string `json:"frontmatter_definition"`
	Description            string `json:"description"`
	// Ingredient  []CoffeeIngredient `json:"ingredients"`
}
type UpdatePestoContentTypePayload struct {
	ID                     string `json:"_id"`
	Name                   string `json:"name"`
	Project_id             string `json:"project_id"`
	Frontmatter_definition string `json:"frontmatter_definition"`
	Description            string `json:"description"`
	// Ingredient  []CoffeeIngredient `json:"ingredients"`
}

//////////////////////////////////////
/// PestoProject(s)
//////////////////////////////////////
// PestoProject -
type PestoProject struct {
	ID                   string `json:"_id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	Git_ssh_uri          string `json:"git_ssh_uri"`
	Git_service_provider string `json:"git_service_provider"`
	// Ingredient  []CoffeeIngredient `json:"ingredients"`
}
type CreatePestoProjectPayload struct {
	//ID                   string    `json:"_id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	Git_ssh_uri          string `json:"git_ssh_uri"`
	Git_service_provider string `json:"git_service_provider"`
	// Ingredient  []CoffeeIngredient `json:"ingredients"`
}
type UpdatePestoProjectPayload struct {
	ID                   string `json:"_id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	Git_ssh_uri          string `json:"git_ssh_uri"`
	Git_service_provider string `json:"git_service_provider"`
	// Ingredient  []CoffeeIngredient `json:"ingredients"`
}

//////////////////////////////////////
/// Below are types defined by the original hashicups
//////////////////////////////////////

// Order -
type Order struct {
	ID    int         `json:"id,omitempty"`
	Items []OrderItem `json:"items,omitempty"`
}

// OrderItem -
type OrderItem struct {
	Coffee   Coffee `json:"coffee"`
	Quantity int    `json:"quantity"`
}

// Coffee -
type Coffee struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Teaser      string             `json:"teaser"`
	Collection  string             `json:"collection"`
	Origin      string             `json:"origin"`
	Color       string             `json:"color"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Image       string             `json:"image"`
	Ingredient  []CoffeeIngredient `json:"ingredients"`
}

// Ingredient -
type CoffeeIngredient struct {
	ID       int    `json:"ingredient_id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

// Ingredient -
type Ingredient struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

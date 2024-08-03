package pesto

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetPestoProjects - Returns list of PestoProjects (no auth required)
func (c *Client) GetPestoProjects() ([]PestoProject, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/pesto-project", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	PestoProjects := []PestoProject{}
	err = json.Unmarshal(body, &PestoProjects)
	if err != nil {
		return nil, err
	}

	return PestoProjects, nil
}

// GetPestoProject - Returns specific PestoProject (no auth required)
func (c *Client) GetPestoProject(PestoProjectID string) ([]PestoProject, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/pesto-project/%s", c.HostURL, PestoProjectID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	PestoProjects := []PestoProject{}
	err = json.Unmarshal(body, &PestoProjects)
	if err != nil {
		return nil, err
	}

	return PestoProjects, nil
}

/*
// GetPestoProjectIngredients - Returns list of PestoProject ingredients (no auth required)
func (c *Client) GetPestoProjectIngredients(PestoProjectID string) ([]Ingredient, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/PestoProjects/%s/ingredients", c.HostURL, PestoProjectID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	ingredients := []Ingredient{}
	err = json.Unmarshal(body, &ingredients)
	if err != nil {
		return nil, err
	}

	return ingredients, nil
}

*/
// CreatePestoProject - Create new PestoProject
func (c *Client) CreatePestoProject(project PestoProject, authToken *string) (*PestoProject, error) {
	rb, err := json.Marshal(project)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/pesto-project", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	newPestoProject := PestoProject{}
	err = json.Unmarshal(body, &newPestoProject)
	if err != nil {
		return nil, err
	}
	return &newPestoProject, nil
}

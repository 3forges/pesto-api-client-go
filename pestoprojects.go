package pesto

import (
	"context"
	"encoding/json"
	// "errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
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
func (c *Client) GetPestoProject(PestoProjectID string) (PestoProject, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/pesto-project/%s", c.HostURL, PestoProjectID), nil)
	if err != nil {
		return PestoProject{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return PestoProject{}, err
	}

	returnedPestoProject := PestoProject{}
	err = json.Unmarshal(body, &returnedPestoProject)
	if err != nil {
		return PestoProject{}, err
	}

	return returnedPestoProject, nil
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
func (c *Client) CreatePestoProject(ctx context.Context, project CreatePestoProjectPayload, authToken *string) (*PestoProject, error) {
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

	tflog.Debug(ctx, fmt.Sprintf("PESTO API CLIENT GO - CREATE PESTO PROJECT - here is the API Response Body returned from Pesto API: %v ", body))
	fmt.Printf("PESTO API CLIENT GO - CREATE PESTO PROJECT - here is the API Response Body returned from Pesto API: %v \n", body)
	var isAPIResponseBodyNil string

	if body != nil {
		isAPIResponseBodyNil = "NO API Response Body object is not NIL"
	} else {
		isAPIResponseBodyNil = "YES API Response Body object is NIL!"
	}
	tflog.Debug(ctx, fmt.Sprintf("PESTO API CLIENT GO - CREATE PESTO PROJECT - Is the API Response Body returned from Pesto API NIL ?: %v", isAPIResponseBodyNil))
	fmt.Printf("PESTO API CLIENT GO - CREATE PESTO PROJECT - Is the API Response Body returned from Pesto API NIL ?: %v \n", isAPIResponseBodyNil)

	newPestoProject := PestoProject{}

	err = json.Unmarshal(body, &newPestoProject)
	if err != nil {
		return nil, err
	}
	return &newPestoProject, nil
}


// UpdatePestoProject - Update new PestoProject
func (c *Client) UpdatePestoProject(ctx context.Context, project UpdatePestoProjectPayload, authToken *string) (*PestoProject, error) {
	rb, err := json.Marshal(project)

	if err != nil {
		return nil, err
	}
	tflog.Debug(ctx, fmt.Sprintf("PESTO API CLIENT GO - UPDATE PESTO PROJECT - here is the payload which will be used to Update thePesto Project in the Pesto API: %v ", project))
	fmt.Printf("PESTO API CLIENT GO - UPDATE PESTO PROJECT - here is the payload which will be used to Update thePesto Project in the Pesto API: %v \n", project)
	
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/pesto-project/%s", c.HostURL, project.ID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, fmt.Sprintf("PESTO API CLIENT GO - UPDATE PESTO PROJECT - here is the API Response Body returned from Pesto API: %v ", body))
	fmt.Printf("PESTO API CLIENT GO - UPDATE PESTO PROJECT - here is the API Response Body returned from Pesto API: %v \n", body)
	var isAPIResponseBodyNil string

	if body != nil {
		isAPIResponseBodyNil = "NO API Response Body object is not NIL"
	} else {
		isAPIResponseBodyNil = "YES API Response Body object is NIL!"
	}
	tflog.Debug(ctx, fmt.Sprintf("PESTO API CLIENT GO - UPDATE PESTO PROJECT - Is the API Response Body returned from Pesto API NIL ?: %v", isAPIResponseBodyNil))
	fmt.Printf("PESTO API CLIENT GO - UPDATE PESTO PROJECT - Is the API Response Body returned from Pesto API NIL ?: %v \n", isAPIResponseBodyNil)

	newPestoProject := PestoProject{}

	err = json.Unmarshal(body, &newPestoProject)
	if err != nil {
		return nil, err
	}
	return &newPestoProject, nil
}



// DeletePestoProject - Deletes a Pesto Project
func (c *Client) DeletePestoProject(ctx context.Context, PestoProjectID string, authToken *string) (*PestoProject, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/pesto-project/%s", c.HostURL, PestoProjectID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, fmt.Sprintf("PESTO API CLIENT GO - DELETE PESTO PROJECT - here is the API Response Body returned from Pesto API: %v ", body))
	fmt.Printf("PESTO API CLIENT GO - DELETE PESTO PROJECT - here is the API Response Body returned from Pesto API: %v \n", body)
	var isAPIResponseBodyNil string

	if body != nil {
		isAPIResponseBodyNil = "NO API Response Body object is not NIL"
	} else {
		isAPIResponseBodyNil = "YES API Response Body object is NIL!"
	}
	tflog.Debug(ctx, fmt.Sprintf("PESTO API CLIENT GO - DELETE PESTO PROJECT - Is the API Response Body returned from Pesto API NIL ?: %v", isAPIResponseBodyNil))
	fmt.Printf("PESTO API CLIENT GO - DELETE PESTO PROJECT - Is the API Response Body returned from Pesto API NIL ?: %v \n", isAPIResponseBodyNil)

	// if string(body) != "Deleted order" {
	// 	return errors.New(string(body))
	// }

	deletedPestoProject := PestoProject{}

	err = json.Unmarshal(body, &deletedPestoProject)
	if err != nil {
		return nil, err
	}
	return &deletedPestoProject, nil
}

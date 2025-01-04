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

// GetPestoContentTypes - Returns list of PestoContentTypes (no auth required)
func (c *Client) GetPestoContentTypes() ([]PestoContentType, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/pesto-content-type", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	PestoContentTypes := []PestoContentType{}
	err = json.Unmarshal(body, &PestoContentTypes)
	if err != nil {
		return nil, err
	}

	return PestoContentTypes, nil
}

// GetPestoContentType - Returns specific PestoContentType (no auth required)
func (c *Client) GetPestoContentType(PestoContentTypeID string) (PestoContentType, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/pesto-content-type/%s", c.HostURL, PestoContentTypeID), nil)
	if err != nil {
		return PestoContentType{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return PestoContentType{}, err
	}

	returnedPestoContentType := PestoContentType{}
	err = json.Unmarshal(body, &returnedPestoContentType)
	if err != nil {
		return PestoContentType{}, err
	}

	return returnedPestoContentType, nil
}

// GetPestoContentTypesByProjectID - Returns all PestoContentTypes for a given Project ID (no auth required)
func (c *Client) GetPestoContentTypesByProjectID(PestoProjectID string) ([]PestoContentType, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/pesto-content-type/project/%s", c.HostURL, PestoProjectID), nil)
	if err != nil {
		return []PestoContentType{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return []PestoContentType{}, err
	}

	//returnedPestoContentType := PestoContentType{}
	//err = json.Unmarshal(body, &returnedPestoContentType)
	returnedPestoContentTypes := []PestoContentType{}
	err = json.Unmarshal(body, &returnedPestoContentTypes)
	if err != nil {
		return []PestoContentType{}, err
	}

	return returnedPestoContentTypes, nil
}

/*
// GetPestoContentTypeIngredients - Returns list of PestoContentType ingredients (no auth required)
func (c *Client) GetPestoContentTypeIngredients(PestoContentTypeID string) ([]Ingredient, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/PestoContentTypes/%s/ingredients", c.HostURL, PestoContentTypeID), nil)
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
// CreatePestoContentType - Create new PestoContentType
func (c *Client) CreatePestoContentType(ctx context.Context, contentType CreatePestoContentTypePayload, authToken *string) (*PestoContentType, error) {
	rb, err := json.Marshal(contentType)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/pesto-content-type", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, fmt.Sprintf("PESTO API CLIENT GO - CREATE PESTO CONTENT TYPE - here is the API Response Body returned from Pesto API: %v ", body))
	fmt.Printf("PESTO API CLIENT GO - CREATE PESTO CONTENT TYPE - here is the API Response Body returned from Pesto API: %v \n", body)
	var isAPIResponseBodyNil string

	if body != nil {
		isAPIResponseBodyNil = "NO API Response Body object is not NIL"
	} else {
		isAPIResponseBodyNil = "YES API Response Body object is NIL!"
	}
	tflog.Debug(ctx, fmt.Sprintf("PESTO API CLIENT GO - CREATE PESTO CONTENT TYPE - Is the API Response Body returned from Pesto API NIL ?: %v", isAPIResponseBodyNil))
	fmt.Printf("PESTO API CLIENT GO - CREATE PESTO CONTENT TYPE - Is the API Response Body returned from Pesto API NIL ?: %v \n", isAPIResponseBodyNil)

	newPestoContentType := PestoContentType{}

	err = json.Unmarshal(body, &newPestoContentType)
	if err != nil {
		return nil, err
	}
	return &newPestoContentType, nil
}

// UpdatePestoContentType - Update new PestoContentType
func (c *Client) UpdatePestoContentType(ctx context.Context, contentType UpdatePestoContentTypePayload, authToken *string) (*PestoContentType, error) {
	rb, err := json.Marshal(contentType)

	if err != nil {
		return nil, err
	}
	tflog.Debug(ctx, fmt.Sprintf("PESTO API CLIENT GO - UPDATE PESTO CONTENT TYPE - here is the payload which will be used to Update the Pesto ContentType in the Pesto API: %v ", contentType))
	fmt.Printf("PESTO API CLIENT GO - UPDATE PESTO CONTENT TYPE - here is the payload which will be used to Update the Pesto ContentType in the Pesto API: %v \n", contentType)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/pesto-content-type/%s", c.HostURL, contentType.ID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, fmt.Sprintf("PESTO API CLIENT GO - UPDATE PESTO CONTENT TYPE - here is the API Response Body returned from Pesto API: %v ", body))
	fmt.Printf("PESTO API CLIENT GO - UPDATE PESTO CONTENT TYPE - here is the API Response Body returned from Pesto API: %v \n", body)
	var isAPIResponseBodyNil string

	if body != nil {
		isAPIResponseBodyNil = "NO API Response Body object is not NIL"
	} else {
		isAPIResponseBodyNil = "YES API Response Body object is NIL!"
	}
	tflog.Debug(ctx, fmt.Sprintf("PESTO API CLIENT GO - UPDATE PESTO CONTENT TYPE - Is the API Response Body returned from Pesto API NIL ?: %v", isAPIResponseBodyNil))
	fmt.Printf("PESTO API CLIENT GO - UPDATE PESTO CONTENT TYPE - Is the API Response Body returned from Pesto API NIL ?: %v \n", isAPIResponseBodyNil)

	newPestoContentType := PestoContentType{}

	err = json.Unmarshal(body, &newPestoContentType)
	if err != nil {
		return nil, err
	}
	return &newPestoContentType, nil
}

// DeletePestoContentType - Deletes a Pesto ContentType
func (c *Client) DeletePestoContentType(ctx context.Context, PestoContentTypeID string, authToken *string) (*PestoContentType, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/pesto-content-type/%s", c.HostURL, PestoContentTypeID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, fmt.Sprintf("PESTO API CLIENT GO - DELETE PESTO CONTENT TYPE - here is the API Response Body returned from Pesto API: %v ", body))
	fmt.Printf("PESTO API CLIENT GO - DELETE PESTO CONTENT TYPE - here is the API Response Body returned from Pesto API: %v \n", body)
	var isAPIResponseBodyNil string

	if body != nil {
		isAPIResponseBodyNil = "NO API Response Body object is not NIL"
	} else {
		isAPIResponseBodyNil = "YES API Response Body object is NIL!"
	}
	tflog.Debug(ctx, fmt.Sprintf("PESTO API CLIENT GO - DELETE PESTO CONTENT TYPE - Is the API Response Body returned from Pesto API NIL ?: %v", isAPIResponseBodyNil))
	fmt.Printf("PESTO API CLIENT GO - DELETE PESTO CONTENT TYPE - Is the API Response Body returned from Pesto API NIL ?: %v \n", isAPIResponseBodyNil)

	// if string(body) != "Deleted order" {
	// 	return errors.New(string(body))
	// }

	deletedPestoContentType := PestoContentType{}

	err = json.Unmarshal(body, &deletedPestoContentType)
	if err != nil {
		return nil, err
	}
	return &deletedPestoContentType, nil
}

package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"flowmoney/bot/internal/errs"
	"flowmoney/bot/internal/models"
	"io"
	"net/http"
	"strconv"
)

func (c *Client) CreateCategory(title string, userId int) (models.Category, error) {
	request := map[string]any{"title": title, "user_id": userId}

	toJSON, err := json.Marshal(request)
	if err != nil {
		return models.Category{}, errs.FailedMarshall
	}

	req, err := http.NewRequest("POST", c.apiUrl+"/category", bytes.NewBuffer(toJSON))
	if err != nil {
		return models.Category{}, errs.RequestFailed
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return models.Category{}, errs.RequestFailed
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Category{}, errors.New(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Category{}, errs.RequestFailed
	}

	var cat models.Category

	err = json.Unmarshal(body, &cat)
	if err != nil {
		return models.Category{}, errs.RequestFailed
	}
	return cat, nil
}

func (c *Client) GetCategoryById(id int) (models.Category, error) {

	req, err := http.NewRequest("GET", c.apiUrl+"/category/"+strconv.Itoa(id), nil)
	if err != nil {
		return models.Category{}, errs.RequestFailed
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return models.Category{}, errs.RequestFailed
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Category{}, errors.New(resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Category{}, errs.RequestFailed
	}
	var cat models.Category
	err = json.Unmarshal(body, &cat)
	if err != nil {
		return models.Category{}, errs.RequestFailed
	}
	return cat, nil
}

func (c *Client) GetByUserId(id int) ([]models.Category, error) {

	req, err := http.NewRequest("GET", c.apiUrl+"/category/user/"+strconv.Itoa(id), nil)
	if err != nil {
		return []models.Category{}, errs.RequestFailed
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []models.Category{}, errs.RequestFailed
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []models.Category{}, errors.New(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []models.Category{}, errs.RequestFailed
	}

	var cat []models.Category

	err = json.Unmarshal(body, &cat)
	if err != nil {
		return []models.Category{}, errs.RequestFailed
	}
	return cat, nil
}

func (c *Client) UpdateCategory(id int, title string) (models.Category, error) {
	request := map[string]any{"title": title}

	toJSON, err := json.Marshal(request)
	if err != nil {
		return models.Category{}, errs.FailedMarshall
	}

	req, err := http.NewRequest("PUT", c.apiUrl+"/category/update/"+strconv.Itoa(id), bytes.NewBuffer(toJSON))
	if err != nil {
		return models.Category{}, errs.RequestFailed
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return models.Category{}, errs.RequestFailed
	}

	if resp.StatusCode != http.StatusOK {
		return models.Category{}, errors.New(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Category{}, errs.RequestFailed
	}

	var cat models.Category

	err = json.Unmarshal(body, &cat)
	if err != nil {
		return models.Category{}, errs.RequestFailed
	}
	return cat, nil
}

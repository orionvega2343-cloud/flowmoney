package client

import (
	"bytes"
	"encoding/json"
	"flowmoney/bot/internal/errs"
	"flowmoney/bot/internal/models"
	"io"
	"net/http"
	"strconv"
)

func (c *Client) CreateBudget(UserId int, CategoryId int, Amount float64, month int, year int) (models.Budget, error) {
	request := map[string]any{"user_id": UserId, "category_id": CategoryId, "amount": Amount, "month": month, "year": year}

	toJSON, err := json.Marshal(request)
	if err != nil {
		return models.Budget{}, errs.FailedMarshall
	}

	req, err := http.NewRequest("POST", c.apiUrl+"/budget", bytes.NewBuffer(toJSON))
	if err != nil {
		return models.Budget{}, errs.RequestFailed
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return models.Budget{}, errs.RequestFailed
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Budget{}, errs.RequestFailed
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Budget{}, errs.RequestFailed
	}

	var budget models.Budget

	err = json.Unmarshal(body, &budget)
	if err != nil {
		return models.Budget{}, errs.FailedMarshall
	}
	return budget, nil
}

func (c *Client) GetBudgetById(id int) (models.Budget, error) {
	req, err := http.NewRequest("GET", c.apiUrl+"/budget/"+strconv.Itoa(id), nil)
	if err != nil {
		return models.Budget{}, errs.RequestFailed
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return models.Budget{}, errs.RequestFailed
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Budget{}, errs.RequestFailed
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Budget{}, errs.RequestFailed
	}

	var budget models.Budget

	err = json.Unmarshal(body, &budget)
	if err != nil {
		return models.Budget{}, errs.FailedMarshall
	}

	return budget, nil
}

func (c *Client) GetBudgetByCategoryId(catId int) (models.Budget, error) {

	req, err := http.NewRequest("GET", c.apiUrl+"/budget/"+strconv.Itoa(catId), nil)
	if err != nil {
		return models.Budget{}, errs.RequestFailed
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return models.Budget{}, errs.RequestFailed
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Budget{}, errs.RequestFailed
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Budget{}, errs.RequestFailed
	}

	var budget models.Budget

	err = json.Unmarshal(body, &budget)
	if err != nil {
		return models.Budget{}, errs.FailedUnmarshall
	}

	return budget, nil
}

func (c *Client) GetByUserIdAndMonth(userId int, month int, year int) (models.Budget, error) {
	req, err := http.NewRequest("GET", c.apiUrl+"/budget/"+strconv.Itoa(userId)+"/monthly/"+strconv.Itoa(month)+"/year/"+strconv.Itoa(year), nil)
	if err != nil {
		return models.Budget{}, errs.RequestFailed
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return models.Budget{}, errs.RequestFailed
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Budget{}, errs.RequestFailed
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Budget{}, errs.RequestFailed
	}

	var budget models.Budget

	err = json.Unmarshal(body, &budget)
	if err != nil {
		return models.Budget{}, errs.FailedUnmarshall
	}

	return budget, nil

}

func (c *Client) UpdateBudget(amount float64, id int) (models.Budget, error) {
	request := map[string]any{"amount": amount, "id": id}

	toJSON, err := json.Marshal(request)
	if err != nil {
		return models.Budget{}, errs.FailedMarshall
	}

	req, err := http.NewRequest("PUT", c.apiUrl+"/budget/"+strconv.Itoa(id), bytes.NewBuffer(toJSON))
	if err != nil {
		return models.Budget{}, errs.RequestFailed
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return models.Budget{}, errs.RequestFailed
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Budget{}, errs.RequestFailed
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Budget{}, errs.RequestFailed
	}

	var budget models.Budget

	err = json.Unmarshal(body, &budget)
	if err != nil {
		return models.Budget{}, errs.FailedUnmarshall
	}

	return budget, nil
}

func (c *Client) DeleteBudgetById(id int) error {
	req, err := http.NewRequest("DELETE", c.apiUrl+"/budget/delete"+strconv.Itoa(id), nil)
	if err != nil {
		return errs.RequestFailed
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return errs.RequestFailed
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errs.RequestFailed
	}

	return nil
}

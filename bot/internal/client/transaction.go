package client

import (
	"bytes"
	"encoding/json"
	"flowmoney/bot/internal/errs"
	"flowmoney/bot/internal/models"
	"io"
	"net/http"
	"strconv"
	"time"
)

func (c *Client) CreateTransaction(UserId int, amount float64, Type string, date time.Time, categoryId int) (models.Transaction, error) {
	request := map[string]any{"user_id": UserId, "amount": amount, "type": Type, "date": date, "category_id": categoryId}

	toJSON, err := json.Marshal(request)
	if err != nil {
		return models.Transaction{}, errs.FailedMarshall
	}

	req, err := http.NewRequest("POST", c.apiUrl+"/transactions", bytes.NewBuffer(toJSON))
	if err != nil {
		return models.Transaction{}, errs.RequestFailed
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return models.Transaction{}, errs.RequestFailed
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return models.Transaction{}, errs.FailedResponse
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Transaction{}, errs.RequestFailed
	}

	var transaction models.Transaction

	err = json.Unmarshal(body, &transaction)
	if err != nil {
		return models.Transaction{}, errs.RequestFailed
	}
	return transaction, nil
}

func (c *Client) GetTransactionById(id int) (models.Transaction, error) {

	req, err := http.NewRequest("GET", c.apiUrl+"/transactions/"+strconv.Itoa(id), nil)
	if err != nil {
		return models.Transaction{}, errs.RequestFailed
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return models.Transaction{}, errs.RequestFailed
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Transaction{}, errs.FailedResponse
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Transaction{}, errs.RequestFailed
	}

	var transaction models.Transaction

	err = json.Unmarshal(body, &transaction)
	if err != nil {
		return models.Transaction{}, errs.RequestFailed
	}

	return transaction, nil
}

func (c *Client) GetTransactionByUserId(userId int) ([]models.Transaction, error) {

	req, err := http.NewRequest("GET", c.apiUrl+"/transactions/user/"+strconv.Itoa(userId), nil)
	if err != nil {
		return []models.Transaction{}, errs.RequestFailed
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []models.Transaction{}, errs.RequestFailed
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []models.Transaction{}, errs.FailedResponse
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []models.Transaction{}, errs.RequestFailed
	}

	var transactions []models.Transaction

	err = json.Unmarshal(body, &transactions)
	if err != nil {
		return []models.Transaction{}, errs.RequestFailed
	}

	return transactions, nil
}

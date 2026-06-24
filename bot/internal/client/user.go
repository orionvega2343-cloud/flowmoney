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

func (c *Client) Register(name, email, password string) error {
	//Структура запроса
	request := map[string]string{"name": name, "email": email, "password": password}
	//Парсинг в JSON
	toJSON, err := json.Marshal(request)
	if err != nil {
		return errs.FailedMarshall
	}
	//Создание нового запроса
	req, err := http.NewRequest("POST", c.apiUrl+"/auth/register", bytes.NewBuffer(toJSON))
	if err != nil {
		return errs.RequestFailed
	}
	req.Header.Set("Content-Type", "application/json")
	//Реализация запроса
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return errs.RequestFailed
	}
	//Ожидаем закрытия тела
	defer resp.Body.Close()
	//Проверка статуса
	if resp.StatusCode != http.StatusCreated {
		return errs.FailedResponse
	}

	return nil
}

func (c *Client) Login(email, password string) error {
	request := map[string]string{"email": email, "password": password}

	toJSON, err := json.Marshal(request)
	if err != nil {
		return errs.FailedMarshall
	}
	req, err := http.NewRequest("POST", c.apiUrl+"/auth/login", bytes.NewBuffer(toJSON))
	if err != nil {
		return errs.RequestFailed
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return errs.RequestFailed
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errs.FailedResponse
	}

	read, err := io.ReadAll(resp.Body)
	if err != nil {
		return errs.FailedRead
	}

	var result map[string]string

	err = json.Unmarshal(read, &result)
	if err != nil {
		return errs.FailedUnmarshall
	}

	c.token = result["token"]
	return nil
}

func (c *Client) GetUserById(id int) (*models.User, error) {
	req, err := http.NewRequest("GET", c.apiUrl+"/user/"+strconv.Itoa(id), nil)
	if err != nil {
		return nil, errs.RequestFailed
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errs.RequestFailed
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errs.FailedResponse
	}

	var user models.User

	read, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errs.FailedRead
	}
	err = json.Unmarshal(read, &user)
	if err != nil {
		return nil, errs.FailedUnmarshall
	}
	return &user, nil
}

func (c *Client) UpdateBalance(id int, balance float64) (models.User, error) {
	//Created request map
	request := map[string]any{"balance": balance}

	//Parse to JSON
	toJSON, err := json.Marshal(request)
	if err != nil {
		return models.User{}, errs.FailedMarshall
	}

	//Created request
	req, err := http.NewRequest("PUT", c.apiUrl+"/update/"+strconv.Itoa(id), bytes.NewBuffer(toJSON))
	if err != nil {
		return models.User{}, errs.RequestFailed
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	//Request approved
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return models.User{}, errs.FailedResponse
	}
	defer resp.Body.Close()
	//Check response status
	if resp.StatusCode != http.StatusOK {
		return models.User{}, errs.FailedResponse
	}

	var user models.User
	//Read response body
	read, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.User{}, errs.FailedRead
	}

	err = json.Unmarshal(read, &user)
	if err != nil {
		return models.User{}, errs.FailedUnmarshall
	}
	return user, nil
}

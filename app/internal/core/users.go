package core

import (
	"EM-TASK/app/internal/domain"
	"EM-TASK/app/internal/domain/errcore"
	"EM-TASK/app/internal/repository/database"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func getDataFromAPI(url string) (map[string]interface{}, error) {
	var data map[string]interface{}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func countryWithHighestProbability(nation map[string]interface{}) string {
	countries, _ := nation["country"].([]interface{})
	max := 0.0
	res := ""

	for _, country := range countries {
		countryInfo, _ := country.(map[string]interface{})
		countryID, ok := countryInfo["country_id"].(string)
		probability, probOk := countryInfo["probability"].(float64)

		if ok && probOk && probability > max {
			max = probability
			res = countryID
		}
	}

	return res
}

func (c *core) CreateUser(ctx context.Context, req *domain.CreateUserRequest) (*domain.UserResponse, error) {
	data := domain.UserResponse{
		Name:       req.Name,
		Surname:    req.Surname,
		Patronymic: req.Patronymic,
	}

	//Обогащение возрастом
	age, err := getDataFromAPI("https://api.agify.io/?name=" + req.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get age drom API: %v", err)
	}

	if ageData, ok := age["age"].(float64); ok {
		data.Age = int(ageData)
	}

	//Обогащение полом
	gender, err := getDataFromAPI("https://api.genderize.io/?name=" + req.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get gender drom API: %v", err)
	}
	if genderStr, ok := gender["gender"].(string); ok {
		data.Gender = genderStr
	}

	//Обогащение национальностью
	nation, err := getDataFromAPI("https://api.nationalize.io/?name=" + req.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get nation drom API: %v", err)
	}
	data.Nation = countryWithHighestProbability(nation)

	user, err := c.repo.CreateUser(ctx, &data)
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	return user, nil
}

func (c *core) GetUser(ctx context.Context, req *domain.GetUserByIdRequest) (*domain.UserResponse, error) {
	user, err := c.repo.GetUser(ctx, req)
	if err != nil {
		if errors.Is(err, database.ErrObjectNotFound) {
			return nil, errcore.UserNotFoundError
		}
		return nil, err
	}

	return user, nil
}

func (c *core) GetAllUsers(ctx context.Context) ([]domain.UserResponse, error) {
	users, err := c.repo.GetAllUsers(ctx)
	if err != nil {
		if errors.Is(err, database.ErrObjectNotFound) {
			return nil, errcore.UserNotFoundError
		}
		return nil, err
	}
	return users, nil
}

func (c *core) GetAllSortUsers(ctx context.Context, req *domain.UserSort) ([]domain.UserResponse, error) {
	users, err := c.repo.GetAllSortUsers(ctx, req)
	if err != nil {
		if errors.Is(err, database.ErrObjectNotFound) {
			return nil, errcore.UserNotFoundError
		}
		return nil, err
	}
	return users, nil
}

func (c *core) UpdateUser(ctx context.Context, req *domain.UpdateUserRequest) (*domain.UserResponse, error) {

	err := c.repo.UpdateUser(ctx, req)
	if err != nil {
		if errors.Is(err, database.ErrObjectNotFound) {
			return nil, errcore.UserNotFoundError
		}
		return nil, err
	}

	user, err := c.GetUser(ctx, &domain.GetUserByIdRequest{ID: req.ID})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	return user, nil
}

func (c *core) DeleteUser(ctx context.Context, req *domain.GetUserByIdRequest) error {

	_, err := c.GetUser(ctx, &domain.GetUserByIdRequest{ID: req.ID})
	if err != nil {
		return errcore.NewInternalError(err)
	}

	err = c.repo.DeleteUser(ctx, req)
	if err != nil {
		return errcore.NewInternalError(err)
	}

	return nil
}

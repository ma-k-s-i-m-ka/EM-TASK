package http

import (
	"EM-TASK/app/internal/domain"
	"EM-TASK/app/internal/infrastructure/logger"
	"fmt"
	"net/http"
)

// @Summary Создать пользователя
// @Description Создает нового пользователя
// @Accept  json
// @Produce  json
// @Param input body domain.CreateUserRequest true "Данные для создания пользователя"
// @Success 201 {object} domain.UserResponse
// @Router /api/user [post]
func (s *server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateUserRequest

	if err := s.readJSON(&req, r); err != nil {
		s.sendError(err, w)
		return
	}

	user, err := s.core.CreateUser(r.Context(), &req)
	if err != nil {
		s.sendError(fmt.Errorf("cannot create user: %w", err), w)
		logger.Logger.Error().Msgf("The CreateUser request failed: %v", err)
		return
	}
	s.sendJSON(http.StatusCreated, user, w)
}

// @Summary Получить пользователя по идентификатору
// @Description Получает пользователя по заданному идентификатору
// @Accept json
// @Produce json
// @Param id path int true "Идентификатор пользователя"
// @Success 200 {object} domain.UserResponse
// @Router /api/user/{user_id} [get]
func (s *server) GetUser(w http.ResponseWriter, r *http.Request) {
	req := domain.GetUserByIdRequest{ID: s.parseParamInt64("user_id", r)}
	user, err := s.core.GetUser(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		logger.Logger.Error().Msgf("The GetUser request failed: %v", err)
		return
	}

	s.sendJSON(http.StatusOK, user, w)
}

// @Summary Получить всех пользователей
// @Description Получает список всех пользователей
// @Accept json
// @Produce json
// @Success 200 {array} domain.UserResponse
// @Router /api/all_users [get]
func (s *server) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := s.core.GetAllUsers(r.Context())
	if err != nil {
		s.sendError(err, w)
		logger.Logger.Error().Msgf("The GetAllUsers request failed: %v", err)
		return
	}

	s.sendJSON(http.StatusOK, users, w)
}

// @Summary Получить всех пользователей с сортировкой
// @Description Получает список всех пользователей с сортировкой
// @Accept json
// @Produce json
// @Param input body domain.UserSort true "Данные для сортировки пользователя"
// @Param name query string false "Имя пользователя"
// @Param surname query string false "Фамилия пользователя"
// @Param patronymic query string false "Отчество пользователя"
// @Param age query int false "Возраст пользователя"
// @Param gender query string false "Пол пользователя"
// @Param nation query string false "Страна пользователя"
// @Success 200 {array} domain.UserResponse
// @Router /api/sort_users [get]
func (s *server) GetAllSortUsers(w http.ResponseWriter, r *http.Request) {

	var req domain.UserSort

	if err := s.readJSON(&req, r); err != nil {
		s.sendError(err, w)
		logger.Logger.Error().Msgf("The body request failed: %v", err)
		return
	}

	users, err := s.core.GetAllSortUsers(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		logger.Logger.Error().Msgf("The GetAllSortUsers request failed: %v", err)
		return
	}

	s.sendJSON(http.StatusOK, users, w)
}

// @Summary Обновить пользователя
// @Description Обновляет существующего пользователя
// @Accept json
// @Produce json
// @Param id path int true "Идентификатор пользователя"
// @Param input body domain.UpdateUserRequest true "Данные для обновления пользователя"
// @Success 200 {object} domain.UserResponse
// @Router /api/user/{user_id} [patch]
func (s *server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := s.parseParamInt64("user_id", r)

	var req domain.UpdateUserRequest
	req.ID = id

	if err := s.readJSON(&req, r); err != nil {
		s.sendError(err, w)
		logger.Logger.Error().Msgf("The body request failed: %v", err)
		return
	}

	user, err := s.core.UpdateUser(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		logger.Logger.Error().Msgf("The UpdateUser request failed: %v", err)
		return
	}

	s.sendJSON(http.StatusOK, user, w)
}

// @Summary Удалить пользователя
// @Description Удаляет пользователя по заданному идентификатору
// @Accept json
// @Produce json
// @Param id path int true "Идентификатор пользователя"
// @Success 200 {string} string
// @Router /api/user/{user_id} [delete]
func (s *server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	req := domain.GetUserByIdRequest{ID: s.parseParamInt64("user_id", r)}

	err := s.core.DeleteUser(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		logger.Logger.Error().Msgf("The DeleteUser request failed: %v", err)
		return
	}

	s.sendJSON(http.StatusOK, "USER DELETED", w)
}

package postgresql

import (
	"EM-TASK/app/internal/domain"
	"EM-TASK/app/internal/domain/errcore"
	"EM-TASK/app/internal/repository/database"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"strings"
)

func (d *driver) CreateUser(ctx context.Context, req *domain.UserResponse) (*domain.UserResponse, error) {

	row := d.conn.QueryRow(ctx, `INSERT INTO users(name, surname, patronymic, age, gender, nation)
                                     VALUES($1, $2, $3, $4, $5, $6)
                                     returning id`,
		req.Name,
		req.Surname,
		req.Patronymic,
		req.Age,
		req.Gender,
		req.Nation,
	)

	err := row.Scan(&req.ID)
	if err != nil {
		err = fmt.Errorf("failed to execute create user query: %v", err)
		return nil, err
	}

	return req, nil
}

func (d *driver) GetUser(ctx context.Context, req *domain.GetUserByIdRequest) (*domain.UserResponse, error) {
	row := d.conn.QueryRow(ctx,
		`SELECT * FROM users
			 WHERE id = $1`, req.ID)

	user := &domain.UserResponse{}
	err := row.Scan(&user.ID, &user.Name, &user.Surname, &user.Patronymic,
		&user.Age, &user.Gender, &user.Nation)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, database.ErrObjectNotFound
		}
		return nil, fmt.Errorf("failed to execute find user by id query: %v", err)
	}

	return user, nil
}

func (d *driver) GetAllUsers(ctx context.Context) ([]domain.UserResponse, error) {
	rows, err := d.conn.Query(ctx,
		`SELECT * FROM users`)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, database.ErrObjectNotFound
	}
	defer rows.Close()

	var users []domain.UserResponse
	for rows.Next() {
		user := domain.UserResponse{}
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Surname,
			&user.Patronymic,
			&user.Age,
			&user.Gender,
			&user.Nation,
		)
		users = append(users, user)
	}
	return users, nil
}

func (d *driver) GetAllSortUsers(ctx context.Context, req *domain.UserSort) ([]domain.UserResponse, error) {
	query := `SELECT * FROM users WHERE true`

	if req.Name != nil {
		query += fmt.Sprintf(" AND name = '%s'", *req.Name)
	}
	if req.Surname != nil {
		query += fmt.Sprintf(" AND surname = '%s'", *req.Surname)
	}
	if req.Patronymic != nil {
		query += fmt.Sprintf(" AND patronymic = '%s'", *req.Patronymic)
	}
	if req.Gender != nil {
		query += fmt.Sprintf(" AND gender = '%s'", *req.Gender)
	}
	if req.Nation != nil {
		query += fmt.Sprintf(" AND nation = '%s'", *req.Nation)
	}

	query += " ORDER BY"
	if req.Name != nil {
		query += " name,"
	}
	if req.Surname != nil {
		query += " surname,"
	}
	if req.Patronymic != nil {
		query += " patronymic,"
	}
	if req.Gender != nil {
		query += " gender,"
	}
	if req.Nation != nil {
		query += " nation,"
	}
	query = strings.TrimRight(query, ",")

	rows, err := d.conn.Query(ctx, query)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, database.ErrObjectNotFound
	}
	defer rows.Close()

	var users []domain.UserResponse
	for rows.Next() {
		user := domain.UserResponse{}
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Surname,
			&user.Patronymic,
			&user.Age,
			&user.Gender,
			&user.Nation,
		)
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, errcore.UsersNotFoundError
	}

	return users, nil
}

func (d *driver) UpdateUser(ctx context.Context, req *domain.UpdateUserRequest) error {

	values := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if req.Name != nil {
		values = append(values, fmt.Sprintf("name=$%d", argId))
		args = append(args, *req.Name)
		argId++
	}
	if req.Surname != nil {
		values = append(values, fmt.Sprintf("surname=$%d", argId))
		args = append(args, *req.Surname)
		argId++
	}
	if req.Patronymic != nil {
		values = append(values, fmt.Sprintf("patronymic=$%d", argId))
		args = append(args, *req.Patronymic)
		argId++
	}
	if req.Age != nil {
		values = append(values, fmt.Sprintf("age=$%d", argId))
		args = append(args, *req.Age)
		argId++
	}
	if req.Gender != nil {
		values = append(values, fmt.Sprintf("gender=$%d", argId))
		args = append(args, *req.Gender)
		argId++
	}
	if req.Nation != nil {
		values = append(values, fmt.Sprintf("nation=$%d", argId))
		args = append(args, *req.Nation)
		argId++
	}

	valuesQuery := strings.Join(values, ", ")
	query := fmt.Sprintf("UPDATE users  SET %s WHERE id = $%d", valuesQuery, argId)
	args = append(args, req.ID)

	_, err := d.conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

func (d *driver) DeleteUser(ctx context.Context, req *domain.GetUserByIdRequest) error {
	result, err := d.conn.Exec(ctx,
		`DELETE FROM users
       		 WHERE id = $1`, req.ID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("deleting user: %v", pgx.ErrNoRows)
	}

	return nil
}

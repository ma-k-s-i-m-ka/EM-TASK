package domain

// @Example CreateUserRequest
type CreateUserRequest struct {
	Name       string  `json:"name" example:"Dmitriy"`
	Surname    string  `json:"surname" example:"Ushakov"`
	Patronymic *string `json:"patronymic,omitempty" example:"Vasilevich"`
}

// @Example UserResponse
type UserResponse struct {
	ID         int64   `json:"id" example:"1567"`
	Name       string  `json:"name" example:"Dmitriy"`
	Surname    string  `json:"surname" example:"Ushakov"`
	Patronymic *string `json:"patronymic,omitempty" example:"Vasilevich"`
	Age        int     `json:"age" example:"28"`
	Gender     string  `json:"gender" example:"male"`
	Nation     string  `json:"nation" example:"slav"`
}

// @Example GetUserByIdRequest
type GetUserByIdRequest struct {
	ID int64 `json:"id" example:"1567"`
}

// @Example UpdateUserRequest
type UpdateUserRequest struct {
	ID         int64   `json:"id" example:"1567"`
	Name       *string `json:"name,omitempty" example:"Dmitriy"`
	Surname    *string `json:"surname,omitempty" example:"Ushakov"`
	Patronymic *string `json:"patronymic,omitempty" example:"Vasilevich"`
	Age        *int    `json:"age,omitempty" example:"28"`
	Gender     *string `json:"gender,omitempty" example:"male"`
	Nation     *string `json:"nation,omitempty" example:"slav"`
}

// @Example UserSort
type UserSort struct {
	Name       *string `json:"name,omitempty" example:"Dmitriy"`
	Surname    *string `json:"surname,omitempty" example:"Ushakov"`
	Patronymic *string `json:"patronymic,omitempty" example:"Vasilevich"`
	Gender     *string `json:"gender,omitempty" example:"male"`
	Nation     *string `json:"nation,omitempty" example:"slav"`
}

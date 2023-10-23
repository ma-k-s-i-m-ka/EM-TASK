package errcore

var (
	UserNotFoundError = &CoreError{
		Message: "User not found.",
		Code:    "User.not_found",
		Type:    NotFoundType,
	}

	UsersNotFoundError = &CoreError{
		Message: "Users not found.",
		Code:    "Users.not_found",
		Type:    NotFoundType,
	}
)

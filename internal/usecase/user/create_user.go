package user

import (
	"context"

	"github.com/geraldobl58/gobid/internal/validator"
)

type CreateUserReq struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

func (req CreateUserReq) Valid(ctx context.Context) validator.Evaluator {

	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.UserName), "user_name", "this field is required")
	eval.CheckField(validator.NotBlank(req.Email), "email", "this field is required")
	eval.CheckField(validator.NotBlank(req.Bio), "bio", "this field is required")

	eval.CheckField(validator.MaxChars(req.Bio, 10) && validator.MinChars(req.Bio, 255), "bio", "this field must be between 10 and 255 characters")

	eval.CheckField(validator.MinChars(req.Password, 8), "password", "this field must be at least 8 characters")

	eval.CheckField(validator.Matches(req.Email, validator.EmailRX), "email", "this field is invalid")

	return eval
}

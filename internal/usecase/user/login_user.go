package user

import (
	"context"

	"github.com/geraldobl58/gobid/internal/validator"
)

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req LoginUserReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.Matches(req.Email, validator.EmailRX), "email", "invalid email")
	eval.CheckField(validator.NotBlank(req.Password), "password", "this field is required")

	return eval
}

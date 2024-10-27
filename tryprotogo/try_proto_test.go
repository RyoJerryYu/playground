package tryprotogo

import (
	"testing"

	"github.com/RyoJerryYu/playground/proto/gengo/api/v1/auth"
	"github.com/RyoJerryYu/playground/proto/gengo/api/v1/users"
)

func TestFieldMask(t *testing.T) {
	auther := auth.GetAuthStatusResponse{
		User: &users.User{
			Name:     "Ryo",
			Nickname: "Jerry",
		},
	}

	auther.FieldPath().User().UpdateTime()
}

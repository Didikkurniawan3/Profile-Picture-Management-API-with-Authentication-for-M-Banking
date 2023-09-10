package user

import "github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/models"

type UserResponse struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}

type UserResponseWithToken struct {
    UserResponse
    Token string `json:"token"`
}

func FormatUserResponse(user models.User, token string) interface{} {
    var formatter interface{}

    userResponse := UserResponse{
        ID:       user.ID,
        Username: user.Username,
        Email:    user.Email,
    }

    if token == "" {
        formatter = userResponse
    } else {
        formatter = UserResponseWithToken{
            UserResponse: userResponse,
            Token:        token,
        }
    }

    return formatter
}

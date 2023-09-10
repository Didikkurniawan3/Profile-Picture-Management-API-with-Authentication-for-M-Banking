package photo

import (
    "github.com/Didik2584/task-5-pbi-btpns-Didik_Kurniawan/models"
    "errors"
)

type PhotoResponse struct {
    ID       int    `json:"id"`
    Title    string `json:"title"`
    Caption  string `json:"caption"`
    PhotoURL string `json:"photo_url"`
    UserID   int    `json:"user_id"`
    User     models.User
}

type PhotoRegularResponse struct {
    ID       int    `json:"id"`
    Title    string `json:"title"`
    Caption  string `json:"caption"`
    PhotoURL string `json:"photo_url"`
}

func FormatPhoto(photo *models.Photo, typeRes string) (interface{}, error) {
    var formatter interface{}

    switch typeRes {
    case "regular":
        formatter = PhotoRegularResponse{
            ID:       photo.ID,
            Title:    photo.Title,
            Caption:  photo.Caption,
            PhotoURL: photo.PhotoURL,
        }
    case "full":
        if photo.User == nil {
            return nil, errors.New("User kosong")
        }
        formatter = PhotoResponse{
            ID:       photo.ID,
            Title:    photo.Title,
            Caption:  photo.Caption,
            PhotoURL: photo.PhotoURL,
            UserID:   photo.User.ID,
            User:     *photo.User,
        }
    default:
        return nil, errors.New("typeRes tidak valid")
    }

    return formatter, nil
}

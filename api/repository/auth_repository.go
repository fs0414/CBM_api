package repository

import (
	"os"
	"time"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/soramar/CBM_api/model/database"
	"github.com/soramar/CBM_api/model/schema"
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

var jwtSecretKey = []byte(os.Getenv("ACCESS_SECRET_KEY"))

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func CreateUser(user *schema.User) error {
	validate := validator.New()
	err := validate.Struct(user)

	if err != nil {
    var errorMessages []string // バリデーションエラーのメッセージを格納

    for _, err := range err.(validator.ValidationErrors) {
        var errorMessage string
        fieldName := err.Field() // NGになったフィールド名を取得
        typ := err.Tag()         // NGになったバリデーションタグを取得
        param := err.Param()     // バリデーションのパラメータ（例: min=8）

        switch fieldName {
        case "Name":
            errorMessage = "名前は必須です"
        case "Email":
            switch typ {
            case "required":
                errorMessage = "メールアドレスは必須です"
            case "email":
                errorMessage = "メールアドレスのフォーマットで登録してください"
            }
        case "Password":
            switch typ {
            case "required":
                errorMessage = "パスワードは必須です"
            case "min":
                // パラメータをメッセージに組み込む
                errorMessage = fmt.Sprintf("パスワードは%s文字以上で登録してください", param)
            }
        case "Role":
            errorMessage = "権限は必須です"
        }
        errorMessages = append(errorMessages, errorMessage)
    }
		// エラーメッセージを結合して返す
		return fmt.Errorf("%v", errorMessages)
	}

	// データベースへの登録処理
	err = database.Db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func GetUserByEmail(email string) (*schema.User, error) {
	var user schema.User
	err := database.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func IsEmailUnregistered(email string) bool {
	var user schema.User
	err := database.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return true
		}
		return true
	}
	return false
}

func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)

	return tokenString, err
}

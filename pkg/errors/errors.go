package errors

import (
	"errors"
	"fmt"
)

// 汎用エラー
var (
	ErrInvalidRequest        = errors.New("リクエストが不正です")
	ErrInternalServerError   = errors.New("サーバーエラーが発生しました")
	ErrInvalidRequestPayload = errors.New("リクエストペイロードが不正です")
	ErrInvalidParameter      = errors.New("パラメータが不正です")
	ErrInvalidToken          = errors.New("トークンが不正です")
)

// ユーザー関連エラー
var (
	ErrFailedCreateUser          = errors.New("ユーザーの作成が失敗しました")
	ErrUserNotFound              = errors.New("ユーザーが見つかりません")
	ErrUserAlreadyExists         = errors.New("ユーザーは既に存在します")
	ErrInvalidUsernameOrPassword = errors.New("ユーザー名またはパスワードが不正です")
)

// 認証関連エラー
var (
	ErrUnAuthorized = errors.New("認証が失敗しました")
)

func New(msg string) error {
	return errors.New(msg)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func NewNotFoundError(entity, id string) error {
	return fmt.Errorf("指定された%sが見つかりません: %s", entity, id)
}

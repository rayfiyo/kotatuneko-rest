package errors

import (
	"errors"
	"fmt"
)

// 汎用エラー
var (
	// リクエストが不正です
	ErrInvalidRequest = errors.New("リクエストが不正です")

	// サーバーエラーが発生しました
	ErrInternalServerError = errors.New("サーバーエラーが発生しました")

	// リクエストペイロードが不正です
	ErrInvalidRequestPayload = errors.New("リクエストペイロードが不正です")

	// パラメータが不正です
	ErrInvalidParameter = errors.New("パラメータが不正です")

	// トークンが不正です
	ErrInvalidToken = errors.New("トークンが不正です")
)

// ユーザー関連エラー
var (
	// ユーザーの作成が失敗しました
	ErrFailedCreateUser = errors.New("ユーザーの作成が失敗しました")

	// ユーザーが見つかりません
	ErrUserNotFound = errors.New("ユーザーが見つかりません")

	// ユーザーは既に存在します
	ErrUserAlreadyExists = errors.New("ユーザーは既に存在します")

	// ユーザー名またはパスワードが不正です
	ErrInvalidUsernameOrPassword = errors.New("ユーザー名またはパスワードが不正です")
)

// 認証関連エラー
var (
	// 認証が失敗しました
	ErrUnAuthorized = errors.New("認証が失敗しました")
)

func New(msg string) error {
	return errors.New(msg)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

// 指定された entity が見つかりません: id
func NewNotFoundError(entity, id string) error {
	return fmt.Errorf("指定された%sが見つかりません: %s", entity, id)
}

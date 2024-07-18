package errors

import (
	"errors"
	"runtime/debug"
)

type Error struct {
	Status     int
	Err        error
	Msg        string
	StackTrace string
}

func New(status int, err error, msg string) *Error {
	return &Error{
		Status:     status,
		Err:        err,
		Msg:        msg,
		StackTrace: string(debug.Stack()),
	}
}

func (e *Error) Error() string {
	return e.Err.Error()
}

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

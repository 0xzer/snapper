package snaperror

import "fmt"


type AuthenticationError struct {
    ErrorMessage string
}

func (e *AuthenticationError) Error() string {
    return fmt.Sprintf("failed to authenticate: %s", e.ErrorMessage)
}

type GetFriendsError struct {
    ErrorMessage string
}

func (e *GetFriendsError) Error() string {
    return fmt.Sprintf("failed to get snapchat friends: %s", e.ErrorMessage)
}

type InvalidUUIDError struct {
    ErrorMessage string
}

func (e *InvalidUUIDError) Error() string {
    return fmt.Sprintf("invalid uuid passed to function: %s", e.ErrorMessage)
}

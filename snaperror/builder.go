package snaperror

import "fmt"

type CreateMessageBuilderError struct {
    ErrorMessage string
}

func (e *CreateMessageBuilderError) Error() string {
    return fmt.Sprintf("failed to build CreateMessageBuilder: %s", e.ErrorMessage)
}

type UpdateMessageBuilderError struct {
    ErrorMessage string
}

func (e *UpdateMessageBuilderError) Error() string {
    return fmt.Sprintf("failed to build UpdateMessageBuilder: %s", e.ErrorMessage)
}

type UpdateConversationBuilderError struct {
    ErrorMessage string
}

func (e *UpdateConversationBuilderError) Error() string {
    return fmt.Sprintf("failed to build UpdateConversationBuilder: %s", e.ErrorMessage)
}

type CreateConversationBuilderError struct {
    ErrorMessage string
}

func (e *CreateConversationBuilderError) Error() string {
    return fmt.Sprintf("failed to build CreateConversationBuilder: %s", e.ErrorMessage)
}
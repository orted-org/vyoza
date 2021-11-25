package httperror

import "fmt"

type CError struct {
	Status int
	Message string
}

func (err *CError) Error() string {
    return fmt.Sprintf("Error Code: %v, %v", err.Status, err.Message)
}
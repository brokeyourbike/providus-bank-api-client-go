package providusbank

import "fmt"

type ErrResponse struct {
	Status    int    `json:"status"`
	Err       string `json:"error"`
	Message   string `json:"message"`
	Path      string `json:"path"`
	Timestamp Time   `json:"timestamp"`
}

func (e ErrResponse) Error() string {
	return fmt.Sprintf("status: %d error: %s msg: %s", e.Status, e.Err, e.Message)
}

package v1

// Message describes a message entry.
type Message struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// User message
type User struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Age        string `json:"age"`
}

// Operate message
type Operate struct {
	Message string `json:"message"`
}

package model

type Item interface {
	GetID() int
	GetName() string
	GetMessage() string
}

type ItemStruct struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (i ItemStruct) GetID() int {
	return i.ID
}

func (i ItemStruct) GetName() string {
	return i.Name
}

func (i ItemStruct) GetMessage() string {
	return i.Message
}

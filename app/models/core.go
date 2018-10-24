package models

type Model interface {
	ConvertToMap() map[string]interface{}
	TableName() string
	FetchAll() []map[string]interface{}
	FetchById() error
	Save() error
	Delete() error
}

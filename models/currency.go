package models

// Currency represents currency data
type Currency struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `currencies`
func (Currency) TableName() string {
	return "currencies"
}

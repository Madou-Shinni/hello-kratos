package model

type Stock struct {
	ID    int `gorm:"primaryKey"`
	Stock int `gorm:"column:stock"`
}

func (i *Stock) TableName() string {
	return "stock"
}

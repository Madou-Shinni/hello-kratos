package model

type Integral struct {
	ID       int `gorm:"primaryKey"`
	Integral int `gorm:"column:integral"`
}

func (i *Integral) TableName() string {
	return "integral"
}

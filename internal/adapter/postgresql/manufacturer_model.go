package postgresql

type ManufacturerModel struct {
	PostgresModel

	ID   uint32 `db:"id" gorm:"primaryKey;autoIncrement"`
	Name string `db:"name" gorm:"type:text;not null;"`
}

func (ManufacturerModel) TableName() string {
	return "manufacturers"
}

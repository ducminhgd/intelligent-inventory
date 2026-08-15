package postgresql

// ModelModel is the GORM model for the "models" table.
type ModelModel struct {
	PostgresModel

	ID             uint32 `db:"id" gorm:"primaryKey;autoIncrement"`
	ManufacturerID uint32 `db:"manufacturer_id" gorm:"not null;index"`
	Name           string `db:"name" gorm:"type:text;not null"`

	// Manufacturer is a BelongsTo association so AutoMigrate creates the foreign
	// key on manufacturer_id. RESTRICT is the safe default for both actions.
	Manufacturer ManufacturerModel `gorm:"foreignKey:ManufacturerID;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
}

func (ModelModel) TableName() string {
	return "models"
}

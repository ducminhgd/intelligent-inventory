package postgresql

import "time"

// StockedVehicleModel is the GORM model for the "stocked_vehicles" table.
type StockedVehicleModel struct {
	PostgresModel

	ID  uint32 `db:"id" gorm:"primaryKey;autoIncrement"`
	VIN string `db:"vin" gorm:"type:text;not null;uniqueIndex"`

	ModelID uint32 `db:"model_id" gorm:"not null;index"`

	Name  string  `db:"name" gorm:"type:text;not null"`
	Price float64 `db:"price" gorm:"type:decimal(16,4);not null"`

	Action string `db:"action" gorm:"type:text;not null;default:NONE;check:chk_stocked_vehicles_action,action IN ('NONE','PRICE_REDUCTION_PLANNED','PRICE_REDUCED','DESTROYED')"`

	// CreatedAt is re-declared here to add an index: it is the basis for the
	// aging computation and is queried by range.
	CreatedAt time.Time `db:"created_at" gorm:"autoCreateTime;type:timestamptz;index"`

	// Model is a BelongsTo association so AutoMigrate creates the foreign key on
	// model_id. RESTRICT is the safe default for both actions.
	Model ModelModel `gorm:"foreignKey:ModelID;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
}

func (StockedVehicleModel) TableName() string {
	return "stocked_vehicles"
}

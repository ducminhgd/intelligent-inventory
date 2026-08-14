package postgresql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresModel struct {
	gorm.Model
	CreatedAt time.Time `db:"created_at" gorm:"autoCreateTime;type:timestamptz"`
	CreatedBy uuid.UUID `db:"created_by" gorm:"type:uuid"`

	UpdatedAt time.Time `db:"updated_at" gorm:"autoUpdateTime;type:timestamptz"`
	UpdatedBy uuid.UUID `db:"updated_by" gorm:"type:uuid"`

	DeletedAt *time.Time `db:"deleted_at" gorm:"type:timestamptz"`
	DeletedBy *uuid.UUID `db:"deleted_by" gorm:"type:uuid"`
}

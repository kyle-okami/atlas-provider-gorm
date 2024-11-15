package customjointable

import (
	"time"

	"gorm.io/gorm"

	"github.com/kyle-okami/atlas-provider-gorm/gormschema"
)

type Person struct {
	ID        int
	Name      string
	Addresses []Address `gorm:"many2many:person_addresses;"`
}

type Address struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type PersonAddress struct {
	PersonID  int `gorm:"primaryKey"`
	AddressID int `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type TopCrowdedAddresses struct {
	AddressID int
	Count     int
}

func (TopCrowdedAddresses) ViewDef(dialect string) []gormschema.ViewOption {
	return []gormschema.ViewOption{
		gormschema.CreateStmt("CREATE VIEW top_crowded_addresses AS SELECT address_id, COUNT(person_id) AS count FROM person_addresses GROUP BY address_id ORDER BY count DESC LIMIT 10"),
	}
}

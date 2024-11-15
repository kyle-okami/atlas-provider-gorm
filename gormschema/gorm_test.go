package gormschema_test

import (
	"os"
	"testing"

	"github.com/kyle-okami/atlas-provider-gorm/gormschema"

	"ariga.io/atlas-go-sdk/recordriver"
	ckmodels "github.com/kyle-okami/atlas-provider-gorm/internal/testdata/circularfks"
	"github.com/kyle-okami/atlas-provider-gorm/internal/testdata/customjointable"
	"github.com/kyle-okami/atlas-provider-gorm/internal/testdata/models"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestPostgreSQLConfig(t *testing.T) {
	resetSession()
	l := gormschema.New("postgres")
	sql, err := l.Load(
		models.WorkingAgedUsers{},
		ckmodels.Location{},
		ckmodels.Event{},
		models.UserPetHistory{},
		models.User{},
		models.Pet{},
		models.TopPetOwner{},
	)
	require.NoError(t, err)
	requireEqualContent(t, sql, "testdata/postgresql_default")
	resetSession()
	l = gormschema.New("postgres", gormschema.WithConfig(
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		}))
	sql, err = l.Load(ckmodels.Location{}, ckmodels.Event{})
	require.NoError(t, err)
	requireEqualContent(t, sql, "testdata/postgresql_no_fk")
}

func TestMySQLConfig(t *testing.T) {
	resetSession()
	l := gormschema.New("mysql")
	sql, err := l.Load(
		models.WorkingAgedUsers{},
		ckmodels.Location{},
		ckmodels.Event{},
		models.UserPetHistory{},
		models.User{},
		models.Pet{},
		models.TopPetOwner{},
	)
	require.NoError(t, err)
	requireEqualContent(t, sql, "testdata/mysql_default")
	resetSession()
	l = gormschema.New("mysql", gormschema.WithConfig(
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		},
	))
	sql, err = l.Load(ckmodels.Location{}, ckmodels.Event{})
	require.NoError(t, err)
	requireEqualContent(t, sql, "testdata/mysql_no_fk")
	resetSession()
	l = gormschema.New("mysql", gormschema.WithJoinTable(&customjointable.Person{}, "Addresses", &customjointable.PersonAddress{}))
	sql, err = l.Load(customjointable.Address{}, customjointable.Person{}, customjointable.TopCrowdedAddresses{})
	require.NoError(t, err)
	requireEqualContent(t, sql, "testdata/mysql_custom_join_table")
	resetSession()
	l = gormschema.New("mysql")
	sql, err = l.Load(customjointable.PersonAddress{}, customjointable.Address{}, customjointable.Person{}, customjointable.TopCrowdedAddresses{})
	require.NoError(t, err)
	requireEqualContent(t, sql, "testdata/mysql_custom_join_table")
	resetSession()
	l = gormschema.New("mysql")
	sql, err = l.Load(customjointable.Address{}, customjointable.PersonAddress{}, customjointable.Person{}, customjointable.TopCrowdedAddresses{})
	require.NoError(t, err)
	requireEqualContent(t, sql, "testdata/mysql_custom_join_table") // position of tables should not matter
}

func TestSQLServerConfig(t *testing.T) {
	resetSession()
	l := gormschema.New("sqlserver")
	sql, err := l.Load(
		models.WorkingAgedUsers{},
		ckmodels.Location{},
		ckmodels.Event{},
		models.UserPetHistory{},
		models.User{},
		models.Pet{},
		models.TopPetOwner{},
	)
	require.NoError(t, err)
	requireEqualContent(t, sql, "testdata/sqlserver_default")
	resetSession()
	l = gormschema.New("sqlserver", gormschema.WithConfig(
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		}))
	sql, err = l.Load(ckmodels.Location{}, ckmodels.Event{})
	require.NoError(t, err)
	requireEqualContent(t, sql, "testdata/sqlserver_no_fk")
}

func resetSession() {
	sess, ok := recordriver.Session("gorm")
	if ok {
		sess.Statements = nil
	}
}

func requireEqualContent(t *testing.T, actual, fileName string) {
	buf, err := os.ReadFile(fileName)
	require.NoError(t, err)
	require.Equal(t, string(buf), actual)
}

package repos_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/userq11/grpc-test/repos"
	"xorm.io/xorm"
)

var (
	err   error
	db    *xorm.Engine
	dbSql *sql.DB
	mock  sqlmock.Sqlmock

	gr            repos.GlobalRepository
	truncateUsers = func() {
		mock.ExpectQuery("TRUNCATE users").WillReturnRows(sqlmock.NewRows([]string{}))

		_, err := db.Query("TRUNCATE users")
		Ω(err).To(BeNil())
	}
	clearDatabase = func() {
		if db == nil {
			Fail("Unable to run tests. Db is nil")
		}
		truncateUsers()
	}
)

var _ = BeforeSuite(func() {
	db, err = xorm.NewEngine("mysql", "")
	Ω(err).To(BeNil())
	dbSql, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	Ω(err).To(BeNil())
	db.DB().DB = dbSql

	gr = repos.GlobalRepo(db)
})

var _ = AfterSuite(func() {
	err = mock.ExpectationsWereMet()
	Ω(err).To(BeNil())
})

func TestRepos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repos Suite")
}

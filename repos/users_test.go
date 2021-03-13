package repos_test

import (
	"errors"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/userq11/grpc-test/models"
)

var _ = Describe("UsersRepo", func() {
	var (
		user1 *models.User

		setupFn = func() {
			user1, err = models.NewUser(&models.TempUser{
				Email:           "test@test.com",
				Password:        "test12345",
				ConfirmPassword: "test12345",
				FirstName:       "Some",
				LastName:        "Name",
			})
			Ω(err).To(BeNil())
		}
	)

	BeforeEach(func() {
		clearDatabase()
		setupFn()
	})

	Describe("Create", func() {
		Context("Failures", func() {
			It("should fail with a nil user", func() {
				err := gr.Users().Create(nil)
				Ω(err).NotTo(BeNil())
				Ω(err.Error()).To(Equal("validator: (nil *models.User)"))
			})
			It("should fail with missing values", func() {
				err := gr.Users().Create(&models.User{Password: user1.Password})
				Ω(err).NotTo(BeNil())
				Ω(err.Error()).To(Equal("Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag"))
			})
			It("should fail with db error", func() {
				errMsg := "database error"

				mock.ExpectExec("INSERT INTO `users` (`email`,`password`,`first_name`,`last_name`,`visible`) VALUES (?,?,?,?,?)").
					WithArgs(user1.Email, user1.Password, user1.FirstName, user1.LastName, user1.Visible).
					WillReturnError(errors.New(errMsg))

				err := gr.Users().Create(user1)
				Ω(err).NotTo(BeNil())
				Ω(err.Error()).To(Equal(errMsg))
			})
		})
		Context("Success", func() {
			It("successfully saves a user in db", func() {
				mock.ExpectExec("INSERT INTO `users` (`email`,`password`,`first_name`,`last_name`,`visible`) VALUES (?,?,?,?,?)").
					WithArgs(user1.Email, user1.Password, user1.FirstName, user1.LastName, user1.Visible).
					WillReturnResult(sqlmock.NewResult(1, 1))

				err := gr.Users().Create(user1)
				Ω(err).To(BeNil())
			})
		})
	})

	Describe("FindByID", func() {
		Context("Failures", func() {
			It("should fail with a bad ID", func() {
				_, err := gr.Users().FindByID(0)
				Ω(err).NotTo(BeNil())
				Ω(err.Error()).To(Equal("bad ID"))
			})
			It("should fail with db error", func() {
				user1.ID = 1
				errMsg := "database error"

				mock.ExpectQuery("SELECT `id`, `email`, `password`, `first_name`, `last_name`, `visible` FROM `users` WHERE `id`=? LIMIT 1").
					WithArgs(user1.ID).
					WillReturnError(errors.New(errMsg))

				_, err := gr.Users().FindByID(1)
				Ω(err).NotTo(BeNil())
				Ω(err.Error()).To(Equal(errMsg))
			})
			It("should unable to find user", func() {
				user1.ID = 1
				errMsg := "could not find user"

				mock.ExpectQuery("SELECT `id`, `email`, `password`, `first_name`, `last_name`, `visible` FROM `users` WHERE `id`=? LIMIT 1").
					WithArgs(user1.ID).
					WillReturnRows(sqlmock.NewRows([]string{}))

				_, err := gr.Users().FindByID(1)
				Ω(err).NotTo(BeNil())
				Ω(err.Error()).To(Equal(errMsg))
			})
		})
		Context("Success", func() {
			It("should successfully add user to db", func() {
				user1.ID = 1

				mock.ExpectQuery("SELECT `id`, `email`, `password`, `first_name`, `last_name`, `visible` FROM `users` WHERE `id`=? LIMIT 1").
					WithArgs(user1.ID).
					WillReturnRows(sqlmock.NewRows(
						[]string{"id", "email", "password", "first_name", "last_name", "visible"}).
						AddRow(user1.ID, user1.Email, user1.Password, user1.FirstName, user1.LastName, user1.Visible),
					)

				found, err := gr.Users().FindByID(user1.ID)
				Ω(err).To(BeNil())
				Ω(found).To(BeEquivalentTo(user1))
			})
		})
	})

	Describe("FindByEmail", func() {
		Context("Failures", func() {
			It("should fail with a bad email", func() {
				_, err := gr.Users().FindByEmail("")
				fmt.Println(err)
				Ω(err).NotTo(BeNil())
				Ω(err.Error()).To(Equal("bad email"))
			})
			It("should fail with db error", func() {
				errMsg := "database error"

				mock.ExpectQuery("SELECT `id`, `email`, `password`, `first_name`, `last_name`, `visible` FROM `users` WHERE `email`=? LIMIT 1").
					WithArgs(user1.Email).
					WillReturnError(errors.New(errMsg))

				_, err := gr.Users().FindByEmail(user1.Email)
				Ω(err).NotTo(BeNil())
				Ω(err.Error()).To(Equal(errMsg))
			})
			It("should unable to find user", func() {
				errMsg := "could not find user"

				mock.ExpectQuery("SELECT `id`, `email`, `password`, `first_name`, `last_name`, `visible` FROM `users` WHERE `email`=? LIMIT 1").
					WithArgs(user1.Email).
					WillReturnRows(sqlmock.NewRows([]string{}))

				_, err := gr.Users().FindByEmail(user1.Email)
				Ω(err).NotTo(BeNil())
				Ω(err.Error()).To(Equal(errMsg))
			})
		})
		Context("Success", func() {
			It("should successfully add user to db", func() {
				mock.ExpectQuery("SELECT `id`, `email`, `password`, `first_name`, `last_name`, `visible` FROM `users` WHERE `email`=? LIMIT 1").
					WithArgs(user1.Email).
					WillReturnRows(sqlmock.NewRows(
						[]string{"id", "email", "password", "first_name", "last_name", "visible"}).
						AddRow(user1.ID, user1.Email, user1.Password, user1.FirstName, user1.LastName, user1.Visible),
					)

				found, err := gr.Users().FindByEmail(user1.Email)
				fmt.Println(err)
				Ω(err).To(BeNil())
				Ω(found).To(BeEquivalentTo(user1))
			})
		})
	})

	Describe("Update", func() {
		Context("Failure", func() {
			It("should fail with nil parameter", func() {
				err := gr.Users().Update(nil)
				Ω(err).NotTo(BeNil())
				Ω(err.Error()).To(Equal("user values invalid"))
			})
			It("should fail with an invalid user (requires ID)", func() {
				err := gr.Users().Update(user1)
				Ω(err).NotTo(BeNil())
				Ω(err.Error()).To(Equal("user values invalid"))
			})
			It("should fail with a database error", func() {
				errMsg := "database error"
				user1.ID = 1

				mock.ExpectExec("UPDATE `users` SET `email` = ?, `password` = ?, `first_name` = ?, `last_name` = ? WHERE `id`=?").
					WithArgs(user1.Email, user1.Password, user1.FirstName, user1.LastName, user1.ID).
					WillReturnError(errors.New(errMsg))

				err := gr.Users().Update(user1)
				fmt.Println(err)
				Ω(err).NotTo(BeNil())
				Ω(err.Error()).To(Equal(errMsg))
			})
		})
		Context("Success", func() {
			It("should update a user", func() {
				user1.ID = 1

				mock.ExpectExec("UPDATE `users` SET `email` = ?, `password` = ?, `first_name` = ?, `last_name` = ? WHERE `id`=?").
					WithArgs(user1.Email, user1.Password, user1.FirstName, user1.LastName, user1.ID).
					WillReturnResult(sqlmock.NewResult(0, 1))

				err := gr.Users().Update(user1)
				Ω(err).To(BeNil())
			})
		})
	})

	Describe("Delete", func() {
		Context("Failure", func() {
			It("should fail with nil parameter", func() {
				err := gr.Users().Delete(nil)
				Ω(err).NotTo(BeNil())
				Ω(err.Error()).To(Equal("user values invalid"))
			})
			It("should fail with invalid id", func() {
				err := gr.Users().Delete(user1)
				Ω(err).NotTo(BeNil())
				Ω(err.Error()).To(Equal("user values invalid"))
			})
			It("should fail with db error", func() {
				user1.ID = 1
				errMsg := "db error"

				mock.ExpectExec("DELETE FROM `users` WHERE `id`=? AND `email`=? AND `password`=? AND `first_name`=? AND `last_name`=? AND `id`=?").
					WithArgs(user1.ID, user1.Email, user1.Password, user1.FirstName, user1.LastName, user1.ID).
					WillReturnError(errors.New(errMsg))

				err := gr.Users().Delete(user1)
				Ω(err).NotTo(BeNil())
				Ω(err.Error()).To(Equal(errMsg))
			})
		})
		Context("Success", func() {
			It("should delete a user", func() {
				user1.ID = 1

				mock.ExpectExec("DELETE FROM `users` WHERE `id`=? AND `email`=? AND `password`=? AND `first_name`=? AND `last_name`=? AND `id`=?").
					WithArgs(user1.ID, user1.Email, user1.Password, user1.FirstName, user1.LastName, user1.ID).
					WillReturnResult(sqlmock.NewResult(0, 1))

				err := gr.Users().Delete(user1)
				Ω(err).To(BeNil())
			})
		})
	})
})

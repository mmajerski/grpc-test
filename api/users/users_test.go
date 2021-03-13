package users

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	"github.com/userq11/grpc-test/models"
	grpc "github.com/userq11/grpc-test/protobufs"
	reposMocks "github.com/userq11/grpc-test/repos/mocks"
	"github.com/userq11/grpc-test/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("grpc", func() {
	var (
		globalRepo *reposMocks.MockGlobalRepository
		usersRepo  *reposMocks.MockUsersRepo
		mockCtrl   *gomock.Controller
		router     grpc.UsersServer
		ctx        context.Context
	)

	setup := func() {
		mockCtrl = gomock.NewController(GinkgoT())
		globalRepo = reposMocks.NewMockGlobalRepository(mockCtrl)
		usersRepo = reposMocks.NewMockUsersRepo(mockCtrl)
		router = GetRoutes()

		ctx = utils.SetGlobalRepoOnContext(context.Background(), globalRepo)

		globalRepo.EXPECT().Users().Return(usersRepo).AnyTimes()
	}

	BeforeEach(func() {
		setup()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Create", func() {
		It("returns error on empty request", func() {
			errMsg := "Key: 'CreateUserRequest.newuser' Error:Field validation for 'newuser' failed on the 'valid-newUser' tag"
			_, err := router.Create(ctx, &grpc.CreateUserRequest{})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("returns error on missing email", func() {
			errMsg := "Key: 'CreateUserRequest.email' Error:Field validation for 'email' failed on the 'valid-email' tag\n" +
				"Key: 'CreateUserRequest.password' Error:Field validation for 'password' failed on the 'valid-password' tag\n" +
				"Key: 'CreateUserRequest.confirmPassword' Error:Field validation for 'confirmPassword' failed on the " + "'valid-confirmPassword' tag\n" +
				"Key: 'CreateUserRequest.firstName' Error:Field validation for 'firstName' failed on the 'valid-firstName' tag\n" +
				"Key: 'CreateUserRequest.lastName' Error:Field validation for 'lastName' failed on the 'valid-lastName' tag"
			_, err := router.Create(ctx, &grpc.CreateUserRequest{NewUser: &grpc.CreateUser{}})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("returns error on missing global repo in context", func() {
			errMsg := "missing global repo in context"

			user, err := models.NewUser(&models.TempUser{
				Email:           "test@test.com",
				Password:        "test12345",
				ConfirmPassword: "test12345",
				FirstName:       "Some",
				LastName:        "Name",
			})
			Ω(err).To(BeNil())

			_, err = router.Create(context.Background(), &grpc.CreateUserRequest{NewUser: &grpc.CreateUser{
				Email:           user.Email,
				Password:        "test12345",
				ConfirmPassword: "test12345",
				FirstName:       user.FirstName,
				LastName:        user.LastName,
			}})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("returns error when cannot create new user", func() {
			errMsg := "password and confirm password do not match"

			_, err := router.Create(ctx, &grpc.CreateUserRequest{NewUser: &grpc.CreateUser{
				Email:           "test@test.com",
				Password:        "test12345",
				ConfirmPassword: "badconfirm",
				FirstName:       "Some",
				LastName:        "Name",
			}})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("should fail on db test", func() {
			errMsg := "database err"

			user, err := models.NewUser(&models.TempUser{
				Email:           "test@test.com",
				Password:        "test12345",
				ConfirmPassword: "test12345",
				FirstName:       "Some",
				LastName:        "Name",
			})
			Ω(err).To(BeNil())

			usersRepo.EXPECT().Create(gomock.AssignableToTypeOf(user)).
				Return(errors.New(errMsg)).Times(1).Do(func(*models.User) {
				defer GinkgoRecover()
			})

			_, err = router.Create(ctx, &grpc.CreateUserRequest{NewUser: &grpc.CreateUser{
				Email:           user.Email,
				Password:        "test12345",
				ConfirmPassword: "test12345",
				FirstName:       user.FirstName,
				LastName:        user.LastName,
			}})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("should be successful", func() {
			user, err := models.NewUser(&models.TempUser{
				Email:           "test@test.com",
				Password:        "test12345",
				ConfirmPassword: "test12345",
				FirstName:       "Some",
				LastName:        "Name",
			})
			Ω(err).To(BeNil())

			usersRepo.EXPECT().Create(gomock.AssignableToTypeOf(user)).
				Return(nil).Times(1).Do(func(*models.User) {
				defer GinkgoRecover()
			})

			res, err := router.Create(ctx, &grpc.CreateUserRequest{NewUser: &grpc.CreateUser{
				Email:           user.Email,
				Password:        "test12345",
				ConfirmPassword: "test12345",
				FirstName:       user.FirstName,
				LastName:        user.LastName,
			}})
			Ω(err).To(BeNil())
			Ω(res.GetUser().GetEmail()).To(Equal(user.Email))
		})
	})

	Describe("FindByID", func() {
		It("returns error on empty request", func() {
			errMsg := "Key: 'FindByIDRequest.id' Error:Field validation for 'id' failed on the 'valid-id' tag"
			_, err := router.FindByID(ctx, &grpc.FindByIDRequest{})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("returns error on missing global repo in context", func() {
			errMsg := "missing global repo in context"

			_, err := router.FindByID(context.Background(), &grpc.FindByIDRequest{Id: 1})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("should fail on db test", func() {
			errMsg := "database err"

			usersRepo.EXPECT().FindByID(int64(1)).Return(nil, errors.New(errMsg)).Times(1).Do(func(int64) {
				defer GinkgoRecover()
			})

			_, err := router.FindByID(ctx, &grpc.FindByIDRequest{Id: 1})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("should be successful", func() {
			usersRepo.EXPECT().FindByID(int64(1)).Return(&models.User{
				ID: 1,
			}, nil).Times(1).Do(func(int64) {
				defer GinkgoRecover()
			})

			res, err := router.FindByID(ctx, &grpc.FindByIDRequest{Id: 1})
			Ω(err).To(BeNil())
			Ω(res.GetUser().GetId()).To(Equal(int64(1)))
		})
	})

	Describe("FindByEmail", func() {
		It("returns error on empty request", func() {
			errMsg := "Key: 'FindByEmailRequest.email' Error:Field validation for 'email' failed on the 'valid-email' tag"
			_, err := router.FindByEmail(ctx, &grpc.FindByEmailRequest{})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("returns error on missing global repo in context", func() {
			errMsg := "missing global repo in context"

			_, err := router.FindByEmail(context.Background(), &grpc.FindByEmailRequest{Email: "mail@mail.com"})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("should fail on db test", func() {
			errMsg := "database err"

			usersRepo.EXPECT().FindByEmail("mail@mail.com").Return(nil, errors.New(errMsg)).Times(1).Do(func(string) {
				defer GinkgoRecover()
			})

			_, err := router.FindByEmail(ctx, &grpc.FindByEmailRequest{Email: "mail@mail.com"})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("should be successful", func() {
			usersRepo.EXPECT().FindByEmail("mail@mail.com").Return(&models.User{
				Email: "mail@mail.com",
			}, nil).Times(1).Do(func(string) {
				defer GinkgoRecover()
			})

			res, err := router.FindByEmail(ctx, &grpc.FindByEmailRequest{Email: "mail@mail.com"})
			Ω(err).To(BeNil())
			Ω(res.GetUser().GetEmail()).To(Equal("mail@mail.com"))
		})
	})

	Describe("Update", func() {
		It("returns error on empty request", func() {
			errMsg := "Key: 'UpdateUserRequest.id' Error:Field validation for 'id' failed on the 'valid-id' tag\n" +
				"Key: 'UpdateUserRequest.password' Error:Field validation for 'password' failed on the 'valid-password' tag\n" +
				"Key: 'UpdateUserRequest.firstName' Error:Field validation for 'firstName' failed on the 'valid-firstName' tag\n" +
				"Key: 'UpdateUserRequest.lastName' Error:Field validation for 'lastName' failed on the 'valid-lastName' tag"
			_, err := router.Update(ctx, &grpc.UpdateUserRequest{})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("returns error on missing global repo in context", func() {
			errMsg := "missing global repo in context"

			_, err := router.Update(context.Background(), &grpc.UpdateUserRequest{
				Id:          1,
				NewPassword: "newpass",
				FirstName:   "FirstName",
				LastName:    "LastName",
			})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("should fail on FindByID", func() {
			errMsg := "database err"

			usersRepo.EXPECT().FindByID(int64(1)).Return(nil, errors.New(errMsg)).Times(1).Do(func(int64) {
				defer GinkgoRecover()
			})

			_, err := router.Update(ctx, &grpc.UpdateUserRequest{
				Id:          1,
				NewPassword: "newpass",
				FirstName:   "FirstName",
				LastName:    "LastName",
			})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("should fail on Update", func() {
			errMsg := "update err"

			user := &models.User{
				ID:        1,
				Email:     "mail@mail.com",
				FirstName: "Some",
				LastName:  "Name",
			}

			usersRepo.EXPECT().FindByID(int64(1)).Return(user, nil).Times(1).Do(func(int64) {
				defer GinkgoRecover()
			})

			usersRepo.EXPECT().Update(user).Return(errors.New(errMsg)).Times(1).Do(func(*models.User) {
				defer GinkgoRecover()
			})

			_, err := router.Update(ctx, &grpc.UpdateUserRequest{
				Id:          1,
				NewPassword: "newpass",
				FirstName:   "FirstName",
				LastName:    "LastName",
			})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("should be successful", func() {
			user := &models.User{
				ID:        1,
				Email:     "mail@mail.com",
				FirstName: "Some",
				LastName:  "Name",
			}

			usersRepo.EXPECT().FindByID(int64(1)).Return(user, nil).Times(1).Do(func(int64) {
				defer GinkgoRecover()
			})

			usersRepo.EXPECT().Update(user).Return(nil).Times(1).Do(func(*models.User) {
				defer GinkgoRecover()
			})

			_, err := router.Update(ctx, &grpc.UpdateUserRequest{
				Id:          1,
				NewPassword: "newpass",
				FirstName:   "FirstName",
				LastName:    "LastName",
			})
			Ω(err).To(BeNil())
		})
	})

	Describe("Delete", func() {
		It("returns error on empty request", func() {
			errMsg := "Key: 'DeleteUserRequest.id' Error:Field validation for 'id' failed on the 'valid-id' tag"
			_, err := router.Delete(ctx, &grpc.DeleteUserRequest{})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("returns error on missing global repo in context", func() {
			errMsg := "missing global repo in context"

			_, err := router.Delete(context.Background(), &grpc.DeleteUserRequest{Id: 1})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("should fail on FindByID", func() {
			errMsg := "database err"

			usersRepo.EXPECT().FindByID(int64(1)).Return(nil, errors.New(errMsg)).Times(1).Do(func(int64) {
				defer GinkgoRecover()
			})

			_, err := router.Delete(ctx, &grpc.DeleteUserRequest{
				Id: int64(1),
			})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("should fail on Delete", func() {
			errMsg := "delete err"

			user := &models.User{
				ID:        1,
				Email:     "mail@mail.com",
				FirstName: "Some",
				LastName:  "Name",
			}

			usersRepo.EXPECT().FindByID(int64(1)).Return(user, nil).Times(1).Do(func(int64) {
				defer GinkgoRecover()
			})

			usersRepo.EXPECT().Delete(user).Return(errors.New(errMsg)).Times(1).Do(func(*models.User) {
				defer GinkgoRecover()
			})

			_, err := router.Delete(ctx, &grpc.DeleteUserRequest{
				Id: 1,
			})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("should be successful", func() {
			user := &models.User{
				ID:        1,
				Email:     "mail@mail.com",
				FirstName: "Some",
				LastName:  "Name",
			}

			usersRepo.EXPECT().FindByID(int64(1)).Return(user, nil).Times(1).Do(func(int64) {
				defer GinkgoRecover()
			})

			usersRepo.EXPECT().Delete(user).Return(nil).Times(1).Do(func(*models.User) {
				defer GinkgoRecover()
			})

			_, err := router.Delete(ctx, &grpc.DeleteUserRequest{
				Id: 1,
			})
			Ω(err).To(BeNil())
		})
	})
})

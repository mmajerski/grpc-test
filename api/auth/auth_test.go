package auth

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	"github.com/pascaldekloe/jwt"
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
		authRepo   *reposMocks.MockAuthRepo
		usersRepo  *reposMocks.MockUsersRepo
		mockCtrl   *gomock.Controller
		router     grpc.AuthServer
		ctx        context.Context
	)

	setup := func() {
		mockCtrl = gomock.NewController(GinkgoT())
		globalRepo = reposMocks.NewMockGlobalRepository(mockCtrl)
		authRepo = reposMocks.NewMockAuthRepo(mockCtrl)
		usersRepo = reposMocks.NewMockUsersRepo(mockCtrl)
		router = GetRoutes()

		ctx = utils.SetGlobalRepoOnContext(context.Background(), globalRepo)

		globalRepo.EXPECT().Auth().Return(authRepo).AnyTimes()
		globalRepo.EXPECT().Users().Return(usersRepo).AnyTimes()
	}

	BeforeEach(func() {
		setup()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Login", func() {
		It("should fail on invalid request", func() {
			_, err := router.Login(ctx, &grpc.LoginRequest{})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(
				"Key: 'LoginRequest.email' Error:Field validation for 'email' failed on the 'valid-email' tag\n" +
					"Key: 'LoginRequest.password' Error:Field validation for 'password' failed on the 'valid-password' tag"))
		})
		It("should fail on invalid request", func() {
			_, err := router.Login(ctx, &grpc.LoginRequest{Email: "no@mail.com"})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(
				"Key: 'LoginRequest.password' Error:Field validation for 'password' failed on the 'valid-password' tag"))
		})
		It("should fail on missing global repo", func() {
			errMsg := "missing global repo in context"

			_, err := router.Login(context.Background(), &grpc.LoginRequest{Email: "no@mail.com", Password: "test12345"})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("should fail on db error when finding by email", func() {
			errMsg := "error db"

			usersRepo.EXPECT().FindByEmail("no@mail.com").
				Return(nil, errors.New(errMsg)).
				Times(1).Do(func(string) {
				defer GinkgoRecover()
			})

			_, err := router.Login(ctx, &grpc.LoginRequest{Email: "no@mail.com", Password: "test12345"})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})
		It("should fail on bad password", func() {
			tmpUser := &models.TempUser{
				Email:           "no@mail.com",
				Password:        "test12345",
				ConfirmPassword: "test12345",
				FirstName:       "Some",
				LastName:        "Name",
			}
			user, err := models.NewUser(tmpUser)
			Ω(err).To(BeNil())

			usersRepo.EXPECT().FindByEmail("no@mail.com").
				Return(user, nil).
				Times(1).Do(func(string) {
				defer GinkgoRecover()
			})

			_, err = router.Login(ctx, &grpc.LoginRequest{Email: "no@mail.com", Password: "badpass"})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal("invalid email or password"))
		})
		It("should fail on unable to get signed token", func() {
			tmpUser := &models.TempUser{
				Email:           "no@mail.com",
				Password:        "test12345",
				ConfirmPassword: "test12345",
				FirstName:       "Some",
				LastName:        "Name",
			}
			user, err := models.NewUser(tmpUser)
			Ω(err).To(BeNil())

			usersRepo.EXPECT().FindByEmail("no@mail.com").
				Return(user, nil).
				Times(1).Do(func(string) {
				defer GinkgoRecover()
			})

			authRepo.EXPECT().GetNewClaims("no@mail.com", map[string]interface{}{
				"user": user,
			}).
				Return(&jwt.Claims{
					Registered: jwt.Registered{
						Subject: "no@mail.com",
					},
					Set: map[string]interface{}{
						"user": user,
					},
				}).
				Times(1).Do(func(string, map[string]interface{}) {
				defer GinkgoRecover()
			})

			authRepo.EXPECT().GetSignedToken(&jwt.Claims{
				Registered: jwt.Registered{
					Subject: "no@mail.com",
				},
				Set: map[string]interface{}{
					"user": user,
				},
			}).
				Return("", errors.New("invalid key")).
				Times(1).Do(func(*jwt.Claims) {
				defer GinkgoRecover()
			})

			_, err = router.Login(ctx, &grpc.LoginRequest{Email: "no@mail.com", Password: "test12345"})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal("invalid key"))
		})
		It("should returns successfully", func() {
			tmpUser := &models.TempUser{
				Email:           "no@mail.com",
				Password:        "test12345",
				ConfirmPassword: "test12345",
				FirstName:       "Some",
				LastName:        "Name",
			}
			user, err := models.NewUser(tmpUser)
			Ω(err).To(BeNil())

			usersRepo.EXPECT().FindByEmail("no@mail.com").
				Return(user, nil).
				Times(1).Do(func(string) {
				defer GinkgoRecover()
			})

			authRepo.EXPECT().GetNewClaims("no@mail.com", map[string]interface{}{
				"user": user,
			}).
				Return(&jwt.Claims{
					Registered: jwt.Registered{
						Subject: "no@mail.com",
					},
					Set: map[string]interface{}{
						"user": user,
					},
				}).
				Times(1).Do(func(string, map[string]interface{}) {
				defer GinkgoRecover()
			})

			authRepo.EXPECT().GetSignedToken(&jwt.Claims{
				Registered: jwt.Registered{
					Subject: "no@mail.com",
				},
				Set: map[string]interface{}{
					"user": user,
				},
			}).
				Return("tokentoken", nil).
				Times(1).Do(func(*jwt.Claims) {
				defer GinkgoRecover()
			})

			res, err := router.Login(ctx, &grpc.LoginRequest{Email: "no@mail.com", Password: "test12345"})
			Ω(err).To(BeNil())
			Ω(res.GetToken()).To(Equal("tokentoken"))
		})
	})
})

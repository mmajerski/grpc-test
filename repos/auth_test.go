package repos_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AuthRepo", func() {

	Describe("GetNewClaims", func() {
		It("should returns claims object successfully", func() {
			claims := gr.Auth().GetNewClaims("test", map[string]interface{}{
				"user": map[string]interface{}{
					"id":        1,
					"email":     "mail@mail.com",
					"password":  "test12345",
					"lastName":  "Some",
					"firstName": "Name",
					"visible":   true,
				},
			})
			Ω(claims).NotTo(BeNil())
			Ω(claims.Set["user"]).NotTo(BeNil())
		})
	})

	Describe("GetSignedToken", func() {
		It("should returns signed token successfully", func() {
			claims := gr.Auth().GetNewClaims("test", map[string]interface{}{
				"user": map[string]interface{}{
					"id":        1,
					"email":     "mail@mail.com",
					"password":  "test12345",
					"lastName":  "Some",
					"firstName": "Name",
					"visible":   true,
				},
			})
			token, err := gr.Auth().GetSignedToken(claims)
			Ω(err).To(BeNil())
			Ω(token).NotTo(Equal(0))

		})
	})

	Describe("GetDataFromToken", func() {
		It("should fail because user data is missing visible", func() {
			claims := gr.Auth().GetNewClaims("test", map[string]interface{}{})

			token, err := gr.Auth().GetSignedToken(claims)
			Ω(err).To(BeNil())
			Ω(token).NotTo(Equal(0))

			user, err := gr.Auth().GetDataFromToken(token)
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal("user data missing in token"))
			Ω(user).To(BeNil())
		})
		It("should fail because user data is missing visible", func() {
			claims := gr.Auth().GetNewClaims("test", map[string]interface{}{
				"user": map[string]interface{}{
					"email":     "mail@mail.com",
					"password":  "test12345",
					"lastName":  "Some",
					"firstName": "Name",
					"visible":   true,
				},
			})

			token, err := gr.Auth().GetSignedToken(claims)
			Ω(err).To(BeNil())
			Ω(token).NotTo(Equal(0))

			user, err := gr.Auth().GetDataFromToken(token)
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal("user data missing in token"))
			Ω(user).To(BeNil())
		})
		It("should fail because user data is missing visible", func() {
			claims := gr.Auth().GetNewClaims("test", map[string]interface{}{
				"user": map[string]interface{}{
					"id":        1,
					"password":  "test12345",
					"lastName":  "Some",
					"firstName": "Name",
					"visible":   true,
				},
			})

			token, err := gr.Auth().GetSignedToken(claims)
			Ω(err).To(BeNil())
			Ω(token).NotTo(Equal(0))

			user, err := gr.Auth().GetDataFromToken(token)
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal("user data missing in token"))
			Ω(user).To(BeNil())
		})
		It("should fail because user data is missing visible", func() {
			claims := gr.Auth().GetNewClaims("test", map[string]interface{}{
				"user": map[string]interface{}{
					"id":        1,
					"email":     "mail@mail.com",
					"password":  "test12345",
					"lastName":  "Some",
					"firstName": "Name",
				},
			})

			token, err := gr.Auth().GetSignedToken(claims)
			Ω(err).To(BeNil())
			Ω(token).NotTo(Equal(0))

			user, err := gr.Auth().GetDataFromToken(token)
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal("user data missing in token"))
			Ω(user).To(BeNil())
		})
		It("should returns user data successfully", func() {
			claims := gr.Auth().GetNewClaims("test", map[string]interface{}{
				"user": map[string]interface{}{
					"id":        1,
					"email":     "mail@mail.com",
					"password":  "test12345",
					"lastName":  "Some",
					"firstName": "Name",
					"visible":   true,
				},
			})
			token, err := gr.Auth().GetSignedToken(claims)
			Ω(err).To(BeNil())
			Ω(token).NotTo(Equal(0))

			user, err := gr.Auth().GetDataFromToken(token)
			Ω(err).To(BeNil())
			Ω(user).To(BeEquivalentTo(user))
		})
	})
})

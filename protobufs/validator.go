package grpc

import valid "github.com/go-playground/validator/v10"

var validator *valid.Validate

func init() {
	validator = valid.New()

	validator.RegisterStructValidation(func(sl valid.StructLevel) {
		r := sl.Current().Interface().(CreateUserRequest)

		if r.GetNewUser() == nil {
			sl.ReportError("NewUser", "newuser", "NewUser", "valid-newUser", "")
		} else {
			if len(r.GetNewUser().GetEmail()) == 0 {
				sl.ReportError("Email", "email", "Email", "valid-email", "")
			}
			if len(r.GetNewUser().GetPassword()) == 0 {
				sl.ReportError("Password", "password", "Password", "valid-password", "")
			}
			if len(r.GetNewUser().GetConfirmPassword()) == 0 {
				sl.ReportError("ConfirmPassword", "confirmPassword", "ConfirmPassword", "valid-confirmPassword", "")
			}
			if len(r.GetNewUser().GetFirstName()) == 0 {
				sl.ReportError("FirstName", "firstName", "FirstName", "valid-firstName", "")
			}
			if len(r.GetNewUser().GetLastName()) == 0 {
				sl.ReportError("LastName", "lastName", "LastName", "valid-lastName", "")
			}
		}
	}, CreateUserRequest{})

	validator.RegisterStructValidation(func(sl valid.StructLevel) {
		r := sl.Current().Interface().(FindByIDRequest)

		if r.GetId() == 0 {
			sl.ReportError("ID", "id", "ID", "valid-id", "")
		}
	}, FindByIDRequest{})

	validator.RegisterStructValidation(func(sl valid.StructLevel) {
		r := sl.Current().Interface().(FindByEmailRequest)

		if len(r.GetEmail()) == 0 {
			sl.ReportError("Email", "email", "Email", "valid-email", "")
		}
	}, FindByEmailRequest{})

	validator.RegisterStructValidation(func(sl valid.StructLevel) {
		r := sl.Current().Interface().(UpdateUserRequest)

		if r.GetId() == 0 {
			sl.ReportError("ID", "id", "ID", "valid-id", "")
		}
		if len(r.GetNewPassword()) == 0 {
			sl.ReportError("Password", "password", "Password", "valid-password", "")
		}
		if len(r.GetFirstName()) == 0 {
			sl.ReportError("FirstName", "firstName", "FirstName", "valid-firstName", "")
		}
		if len(r.GetLastName()) == 0 {
			sl.ReportError("LastName", "lastName", "LastName", "valid-lastName", "")
		}
	}, UpdateUserRequest{})

	validator.RegisterStructValidation(func(sl valid.StructLevel) {
		r := sl.Current().Interface().(DeleteUserRequest)

		if r.GetId() == 0 {
			sl.ReportError("ID", "id", "ID", "valid-id", "")
		}
	}, DeleteUserRequest{})

	validator.RegisterStructValidation(func(sl valid.StructLevel) {
		r := sl.Current().Interface().(LoginRequest)

		if len(r.GetEmail()) == 0 {
			sl.ReportError("Email", "email", "Email", "valid-email", "")
		}

		if len(r.GetPassword()) == 0 {
			sl.ReportError("Password", "password", "Password", "valid-password", "")
		}
	}, LoginRequest{})
}

// Validate validates tags on object
func Validate(t interface{}) error {
	return validator.Struct(t)
}

### Golang gRPC API

Written in Go, functionalities:

- CRUD user,
- hashes password with bcrypt,
- validates fileds with validator package,
- signs in new users.

Uses xorm as orm, tests are written in Ginkgo. Primary goal of this project was to learn gRPC and Ginkgo so test coverage is ~100%.

##

**Stack**

- Go,
- MySQL(during development it was running inside docker container),
- golang-migrate for db migrations,
- validator for validate fields on struct,
- xorm as orm,
- ginkgo for BDD tests,
- sqlmock for sql driver behavior in tests,
- protobufs,
- gRPC for communication,
- gomock for testing api,
- dummy react client just for demonstrating purposes.

**How to run**

- `go run main.go`
- `go make proxy`
- `cd client && npm start`

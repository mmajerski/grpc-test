package repos

import "xorm.io/xorm"

var globalRepositoryInst *globalRepository

// GlobalRepository describes interface for global repository
type GlobalRepository interface {
	Users() UsersRepo
	Auth() AuthRepo
}

// GlobalRepo sets globalRepository
func GlobalRepo(db *xorm.Engine) GlobalRepository {
	if globalRepositoryInst != nil {
		return globalRepositoryInst
	}

	globalRepositoryInst = &globalRepository{db: db, repos: make(map[string]interface{})}
	return globalRepositoryInst
}

type globalRepository struct {
	db    *xorm.Engine
	repos map[string]interface{}
}

func (r *globalRepository) repoFactory(key string, factory func() interface{}) interface{} {
	if iface, exists := r.repos[key]; exists {
		return iface
	}
	iface := factory()
	r.repos[key] = iface
	return iface
}

func (r globalRepository) Users() UsersRepo {
	return r.repoFactory("Users", func() interface{} { return NewUsersRepo(r.db) }).(UsersRepo)
}

func (r globalRepository) Auth() AuthRepo {
	return r.repoFactory("Auth", func() interface{} { return NewAuthRepo(r.db) }).(AuthRepo)
}

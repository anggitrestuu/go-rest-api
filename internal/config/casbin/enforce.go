package casbin

import casbin "github.com/casbin/casbin/v2"

func setupCasbinEnforcer() (*casbin.Enforcer, error) {
	return casbin.NewEnforcer("internal/config/casbin/rbac_model.conf", "internal/config/casbin/policy.csv")
}

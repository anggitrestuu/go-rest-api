package v1

import (
	"context"
	"net/http"

	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
)

type roleAuthorizationUseCase struct {
	repo V1Domains.RoleAuthorizationRepository
}

func NewRoleAuthorizationUseCase(repo V1Domains.RoleAuthorizationRepository) V1Domains.RoleAuthorizationUseCase {
	return &roleAuthorizationUseCase{
		repo: repo,
	}
}

func (r *roleAuthorizationUseCase) AssignAuthorizationToRole(ctx context.Context, roleID, authorizationID int) (outDom V1Domains.RoleAuthorizationsDomain, statusCode int, err error) {
	err = r.repo.AssignAuthorizationToRole(ctx, roleID, authorizationID)
	if err != nil {
		return V1Domains.RoleAuthorizationsDomain{}, http.StatusBadRequest, err
	}

	outDom = V1Domains.RoleAuthorizationsDomain{
		RolesID:          roleID,
		AuthorizationsID: authorizationID,
	}

	return outDom, http.StatusCreated, nil
}

func (r *roleAuthorizationUseCase) RemoveAuthorizationFromRole(ctx context.Context, roleID, authorizationID int) (outDom V1Domains.RoleAuthorizationsDomain, statusCode int, err error) {
	err = r.repo.RemoveAuthorizationFromRole(ctx, roleID, authorizationID)
	if err != nil {
		return V1Domains.RoleAuthorizationsDomain{}, http.StatusBadRequest, err
	}

	outDom = V1Domains.RoleAuthorizationsDomain{
		RolesID:          roleID,
		AuthorizationsID: authorizationID,
	}

	return outDom, http.StatusOK, nil
}

func (r *roleAuthorizationUseCase) GetAuthorizationsByRoleID(ctx context.Context, roleID int) (outDom []V1Domains.AuthorizationDomain, statusCode int, err error) {
	outDom, err = r.repo.GetAuthorizationsByRoleID(ctx, roleID)
	if err != nil {
		return []V1Domains.AuthorizationDomain{}, http.StatusBadRequest, err
	}

	return outDom, http.StatusOK, nil
}

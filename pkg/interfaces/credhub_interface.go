package interfaces

import (
	"code.cloudfoundry.org/credhub-cli/credhub"
	"code.cloudfoundry.org/credhub-cli/credhub/auth"
	"code.cloudfoundry.org/credhub-cli/credhub/permissions"
)

type CredHubApi struct {
	CredHub *credhub.CredHub
}

//go:generate counterfeiter ./ CredHubAPI
type CredHubAPI interface {
	//methods to implement
	GetPermissionByPathActor(path string, actor string) (*permissions.Permission, error)
	AddPermission(path string, actor string, ops []string) (*permissions.Permission, error)
	UpdatePermission(uuid string, path string, actor string, ops []string) (*permissions.Permission, error)
	DeletePermission(uuid string) (*permissions.Permission, error)
}

func NewCredHubApi(target string, skipTLS bool, clientID string, clientPwd string, uaaEndpoint string) (CredHubAPI, error) {

	api, err := credhub.New(target,
		credhub.SkipTLSValidation(skipTLS),
		credhub.Auth(auth.UaaClientCredentials(clientID, clientPwd)),
		credhub.AuthURL(uaaEndpoint))
	return &CredHubApi{api}, err
}

func (c *CredHubApi) GetPermissionByPathActor(path string, actor string) (*permissions.Permission, error) {
	return c.CredHub.GetPermissionByPathActor(path, actor)
}

func (c *CredHubApi) AddPermission(path string, actor string, ops []string) (*permissions.Permission, error) {
	return c.CredHub.AddPermission(path, actor, ops)
}

func (c *CredHubApi) UpdatePermission(uuid string, path string, actor string, ops []string) (*permissions.Permission, error) {
	return c.CredHub.UpdatePermission(uuid, path, actor, ops)
}

func (c *CredHubApi) DeletePermission(uuid string) (*permissions.Permission, error) {
	return c.CredHub.DeletePermission(uuid)
}

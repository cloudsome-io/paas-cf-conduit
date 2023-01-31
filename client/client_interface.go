// Code generated by ifacemaker; DO NOT EDIT.

package client

import (
	gocfclient "github.com/cloudfoundry-community/go-cfclient"
)

// Client ...
type Client interface {
	RefreshAccessToken() error
	GetAppEnv(appGuid string) (*Env, error)
	GetSpaceByName(orgGuid string, name string) (*gocfclient.Space, error)
	GetOrgByName(name string) (*gocfclient.Org, error)
	GetAppByName(orgGuid, spaceGuid, appName string) (*gocfclient.App, error)
	GetServiceBindings(filters ...string) (map[string]*gocfclient.ServiceBinding, error)
	GetServiceInstances(filters ...string) (map[string]*gocfclient.ServiceInstance, error)
	BindService(appGuid string, serviceInstanceGuid string, parameters map[string]interface{}) (*Credentials, error)
	UploadStaticAppBits(appGuid string) error
	DestroyApp(appGuid string) error
	CreateApp(name string, spaceGUID string) (guid string, err error)
	StartApp(appGuid string) error
	PollForAppState(appGuid string, state string, maxRetries int) error
	AppSSHEndpoint() string
	AppSSHHostKeyFingerprint() string
	SSHCode() (string, error)
}

package rest

import (
	"github.com/Aptomi/aptomi/pkg/client"
	"github.com/Aptomi/aptomi/pkg/client/rest/http"
	"github.com/Aptomi/aptomi/pkg/config"
)

type coreClient struct {
	cfg        *config.Client
	httpClient http.Client
}

// New returns new instance of the Core API client http rest implementation
func New(cfg *config.Client, httpClient http.Client) client.Core {
	return &coreClient{cfg, httpClient}
}

func (client *coreClient) Policy() client.Policy {
	return &policyClient{client.cfg, client.httpClient}
}

func (client *coreClient) Dependency() client.Dependency {
	return &dependencyClient{client.cfg, client.httpClient}
}

func (client *coreClient) Revision() client.Revision {
	return &revisionClient{client.cfg, client.httpClient}
}

func (client *coreClient) State() client.State {
	return &stateClient{client.cfg, client.httpClient}
}

func (client *coreClient) User() client.User {
	return &userClient{client.cfg, client.httpClient}
}

func (client *coreClient) Version() client.Version {
	return &versionClient{client.cfg, client.httpClient}
}

package client

const (
	KUBERNETES_SERVICE_PORT_TYPE = "kubernetesServicePort"
)

type KubernetesServicePort struct {
	Resource

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	NodePort int64 `json:"nodePort,omitempty" yaml:"node_port,omitempty"`

	Port int64 `json:"port,omitempty" yaml:"port,omitempty"`

	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	TargetPort int64 `json:"targetPort,omitempty" yaml:"target_port,omitempty"`
}

type KubernetesServicePortCollection struct {
	Collection
	Data []KubernetesServicePort `json:"data,omitempty"`
}

type KubernetesServicePortClient struct {
	rancherClient *RancherClient
}

type KubernetesServicePortOperations interface {
	List(opts *ListOpts) (*KubernetesServicePortCollection, error)
	Create(opts *KubernetesServicePort) (*KubernetesServicePort, error)
	Update(existing *KubernetesServicePort, updates interface{}) (*KubernetesServicePort, error)
	ById(id string) (*KubernetesServicePort, error)
	Delete(container *KubernetesServicePort) error
}

func newKubernetesServicePortClient(rancherClient *RancherClient) *KubernetesServicePortClient {
	return &KubernetesServicePortClient{
		rancherClient: rancherClient,
	}
}

func (c *KubernetesServicePortClient) Create(container *KubernetesServicePort) (*KubernetesServicePort, error) {
	resp := &KubernetesServicePort{}
	err := c.rancherClient.doCreate(KUBERNETES_SERVICE_PORT_TYPE, container, resp)
	return resp, err
}

func (c *KubernetesServicePortClient) Update(existing *KubernetesServicePort, updates interface{}) (*KubernetesServicePort, error) {
	resp := &KubernetesServicePort{}
	err := c.rancherClient.doUpdate(KUBERNETES_SERVICE_PORT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *KubernetesServicePortClient) List(opts *ListOpts) (*KubernetesServicePortCollection, error) {
	resp := &KubernetesServicePortCollection{}
	err := c.rancherClient.doList(KUBERNETES_SERVICE_PORT_TYPE, opts, resp)
	return resp, err
}

func (c *KubernetesServicePortClient) ById(id string) (*KubernetesServicePort, error) {
	resp := &KubernetesServicePort{}
	err := c.rancherClient.doById(KUBERNETES_SERVICE_PORT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *KubernetesServicePortClient) Delete(container *KubernetesServicePort) error {
	return c.rancherClient.doResourceDelete(KUBERNETES_SERVICE_PORT_TYPE, &container.Resource)
}

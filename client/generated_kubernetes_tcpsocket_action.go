package client

const (
	KUBERNETES_TCPSOCKET_ACTION_TYPE = "kubernetesTCPSocketAction"
)

type KubernetesTCPSocketAction struct {
	Resource

	Port string `json:"port,omitempty" yaml:"port,omitempty"`
}

type KubernetesTCPSocketActionCollection struct {
	Collection
	Data []KubernetesTCPSocketAction `json:"data,omitempty"`
}

type KubernetesTCPSocketActionClient struct {
	rancherClient *RancherClient
}

type KubernetesTCPSocketActionOperations interface {
	List(opts *ListOpts) (*KubernetesTCPSocketActionCollection, error)
	Create(opts *KubernetesTCPSocketAction) (*KubernetesTCPSocketAction, error)
	Update(existing *KubernetesTCPSocketAction, updates interface{}) (*KubernetesTCPSocketAction, error)
	ById(id string) (*KubernetesTCPSocketAction, error)
	Delete(container *KubernetesTCPSocketAction) error
}

func newKubernetesTCPSocketActionClient(rancherClient *RancherClient) *KubernetesTCPSocketActionClient {
	return &KubernetesTCPSocketActionClient{
		rancherClient: rancherClient,
	}
}

func (c *KubernetesTCPSocketActionClient) Create(container *KubernetesTCPSocketAction) (*KubernetesTCPSocketAction, error) {
	resp := &KubernetesTCPSocketAction{}
	err := c.rancherClient.doCreate(KUBERNETES_TCPSOCKET_ACTION_TYPE, container, resp)
	return resp, err
}

func (c *KubernetesTCPSocketActionClient) Update(existing *KubernetesTCPSocketAction, updates interface{}) (*KubernetesTCPSocketAction, error) {
	resp := &KubernetesTCPSocketAction{}
	err := c.rancherClient.doUpdate(KUBERNETES_TCPSOCKET_ACTION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *KubernetesTCPSocketActionClient) List(opts *ListOpts) (*KubernetesTCPSocketActionCollection, error) {
	resp := &KubernetesTCPSocketActionCollection{}
	err := c.rancherClient.doList(KUBERNETES_TCPSOCKET_ACTION_TYPE, opts, resp)
	return resp, err
}

func (c *KubernetesTCPSocketActionClient) ById(id string) (*KubernetesTCPSocketAction, error) {
	resp := &KubernetesTCPSocketAction{}
	err := c.rancherClient.doById(KUBERNETES_TCPSOCKET_ACTION_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *KubernetesTCPSocketActionClient) Delete(container *KubernetesTCPSocketAction) error {
	return c.rancherClient.doResourceDelete(KUBERNETES_TCPSOCKET_ACTION_TYPE, &container.Resource)
}

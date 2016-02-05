package client

const (
	KUBERNETES_HANDLER_TYPE = "kubernetesHandler"
)

type KubernetesHandler struct {
	Resource

	Exec KubernetesExecAction `json:"exec,omitempty" yaml:"exec,omitempty"`

	HttpGet KubernetesHTTPGetAction `json:"httpGet,omitempty" yaml:"http_get,omitempty"`

	TcpSocket KubernetesTCPSocketAction `json:"tcpSocket,omitempty" yaml:"tcp_socket,omitempty"`
}

type KubernetesHandlerCollection struct {
	Collection
	Data []KubernetesHandler `json:"data,omitempty"`
}

type KubernetesHandlerClient struct {
	rancherClient *RancherClient
}

type KubernetesHandlerOperations interface {
	List(opts *ListOpts) (*KubernetesHandlerCollection, error)
	Create(opts *KubernetesHandler) (*KubernetesHandler, error)
	Update(existing *KubernetesHandler, updates interface{}) (*KubernetesHandler, error)
	ById(id string) (*KubernetesHandler, error)
	Delete(container *KubernetesHandler) error
}

func newKubernetesHandlerClient(rancherClient *RancherClient) *KubernetesHandlerClient {
	return &KubernetesHandlerClient{
		rancherClient: rancherClient,
	}
}

func (c *KubernetesHandlerClient) Create(container *KubernetesHandler) (*KubernetesHandler, error) {
	resp := &KubernetesHandler{}
	err := c.rancherClient.doCreate(KUBERNETES_HANDLER_TYPE, container, resp)
	return resp, err
}

func (c *KubernetesHandlerClient) Update(existing *KubernetesHandler, updates interface{}) (*KubernetesHandler, error) {
	resp := &KubernetesHandler{}
	err := c.rancherClient.doUpdate(KUBERNETES_HANDLER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *KubernetesHandlerClient) List(opts *ListOpts) (*KubernetesHandlerCollection, error) {
	resp := &KubernetesHandlerCollection{}
	err := c.rancherClient.doList(KUBERNETES_HANDLER_TYPE, opts, resp)
	return resp, err
}

func (c *KubernetesHandlerClient) ById(id string) (*KubernetesHandler, error) {
	resp := &KubernetesHandler{}
	err := c.rancherClient.doById(KUBERNETES_HANDLER_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *KubernetesHandlerClient) Delete(container *KubernetesHandler) error {
	return c.rancherClient.doResourceDelete(KUBERNETES_HANDLER_TYPE, &container.Resource)
}

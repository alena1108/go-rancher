package client

const (
	KUBERNETES_HTTPGET_ACTION_TYPE = "kubernetesHTTPGetAction"
)

type KubernetesHTTPGetAction struct {
	Resource

	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	Port string `json:"port,omitempty" yaml:"port,omitempty"`

	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
}

type KubernetesHTTPGetActionCollection struct {
	Collection
	Data []KubernetesHTTPGetAction `json:"data,omitempty"`
}

type KubernetesHTTPGetActionClient struct {
	rancherClient *RancherClient
}

type KubernetesHTTPGetActionOperations interface {
	List(opts *ListOpts) (*KubernetesHTTPGetActionCollection, error)
	Create(opts *KubernetesHTTPGetAction) (*KubernetesHTTPGetAction, error)
	Update(existing *KubernetesHTTPGetAction, updates interface{}) (*KubernetesHTTPGetAction, error)
	ById(id string) (*KubernetesHTTPGetAction, error)
	Delete(container *KubernetesHTTPGetAction) error
}

func newKubernetesHTTPGetActionClient(rancherClient *RancherClient) *KubernetesHTTPGetActionClient {
	return &KubernetesHTTPGetActionClient{
		rancherClient: rancherClient,
	}
}

func (c *KubernetesHTTPGetActionClient) Create(container *KubernetesHTTPGetAction) (*KubernetesHTTPGetAction, error) {
	resp := &KubernetesHTTPGetAction{}
	err := c.rancherClient.doCreate(KUBERNETES_HTTPGET_ACTION_TYPE, container, resp)
	return resp, err
}

func (c *KubernetesHTTPGetActionClient) Update(existing *KubernetesHTTPGetAction, updates interface{}) (*KubernetesHTTPGetAction, error) {
	resp := &KubernetesHTTPGetAction{}
	err := c.rancherClient.doUpdate(KUBERNETES_HTTPGET_ACTION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *KubernetesHTTPGetActionClient) List(opts *ListOpts) (*KubernetesHTTPGetActionCollection, error) {
	resp := &KubernetesHTTPGetActionCollection{}
	err := c.rancherClient.doList(KUBERNETES_HTTPGET_ACTION_TYPE, opts, resp)
	return resp, err
}

func (c *KubernetesHTTPGetActionClient) ById(id string) (*KubernetesHTTPGetAction, error) {
	resp := &KubernetesHTTPGetAction{}
	err := c.rancherClient.doById(KUBERNETES_HTTPGET_ACTION_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *KubernetesHTTPGetActionClient) Delete(container *KubernetesHTTPGetAction) error {
	return c.rancherClient.doResourceDelete(KUBERNETES_HTTPGET_ACTION_TYPE, &container.Resource)
}

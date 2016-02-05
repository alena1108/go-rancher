package client

const (
	KUBERNETES_EXEC_ACTION_TYPE = "kubernetesExecAction"
)

type KubernetesExecAction struct {
	Resource

	Command []string `json:"command,omitempty" yaml:"command,omitempty"`
}

type KubernetesExecActionCollection struct {
	Collection
	Data []KubernetesExecAction `json:"data,omitempty"`
}

type KubernetesExecActionClient struct {
	rancherClient *RancherClient
}

type KubernetesExecActionOperations interface {
	List(opts *ListOpts) (*KubernetesExecActionCollection, error)
	Create(opts *KubernetesExecAction) (*KubernetesExecAction, error)
	Update(existing *KubernetesExecAction, updates interface{}) (*KubernetesExecAction, error)
	ById(id string) (*KubernetesExecAction, error)
	Delete(container *KubernetesExecAction) error
}

func newKubernetesExecActionClient(rancherClient *RancherClient) *KubernetesExecActionClient {
	return &KubernetesExecActionClient{
		rancherClient: rancherClient,
	}
}

func (c *KubernetesExecActionClient) Create(container *KubernetesExecAction) (*KubernetesExecAction, error) {
	resp := &KubernetesExecAction{}
	err := c.rancherClient.doCreate(KUBERNETES_EXEC_ACTION_TYPE, container, resp)
	return resp, err
}

func (c *KubernetesExecActionClient) Update(existing *KubernetesExecAction, updates interface{}) (*KubernetesExecAction, error) {
	resp := &KubernetesExecAction{}
	err := c.rancherClient.doUpdate(KUBERNETES_EXEC_ACTION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *KubernetesExecActionClient) List(opts *ListOpts) (*KubernetesExecActionCollection, error) {
	resp := &KubernetesExecActionCollection{}
	err := c.rancherClient.doList(KUBERNETES_EXEC_ACTION_TYPE, opts, resp)
	return resp, err
}

func (c *KubernetesExecActionClient) ById(id string) (*KubernetesExecAction, error) {
	resp := &KubernetesExecAction{}
	err := c.rancherClient.doById(KUBERNETES_EXEC_ACTION_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *KubernetesExecActionClient) Delete(container *KubernetesExecAction) error {
	return c.rancherClient.doResourceDelete(KUBERNETES_EXEC_ACTION_TYPE, &container.Resource)
}

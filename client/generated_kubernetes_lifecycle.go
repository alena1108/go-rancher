package client

const (
	KUBERNETES_LIFECYCLE_TYPE = "kubernetesLifecycle"
)

type KubernetesLifecycle struct {
	Resource

	PostStart KubernetesHandler `json:"postStart,omitempty" yaml:"post_start,omitempty"`

	PreStart KubernetesHandler `json:"preStart,omitempty" yaml:"pre_start,omitempty"`
}

type KubernetesLifecycleCollection struct {
	Collection
	Data []KubernetesLifecycle `json:"data,omitempty"`
}

type KubernetesLifecycleClient struct {
	rancherClient *RancherClient
}

type KubernetesLifecycleOperations interface {
	List(opts *ListOpts) (*KubernetesLifecycleCollection, error)
	Create(opts *KubernetesLifecycle) (*KubernetesLifecycle, error)
	Update(existing *KubernetesLifecycle, updates interface{}) (*KubernetesLifecycle, error)
	ById(id string) (*KubernetesLifecycle, error)
	Delete(container *KubernetesLifecycle) error
}

func newKubernetesLifecycleClient(rancherClient *RancherClient) *KubernetesLifecycleClient {
	return &KubernetesLifecycleClient{
		rancherClient: rancherClient,
	}
}

func (c *KubernetesLifecycleClient) Create(container *KubernetesLifecycle) (*KubernetesLifecycle, error) {
	resp := &KubernetesLifecycle{}
	err := c.rancherClient.doCreate(KUBERNETES_LIFECYCLE_TYPE, container, resp)
	return resp, err
}

func (c *KubernetesLifecycleClient) Update(existing *KubernetesLifecycle, updates interface{}) (*KubernetesLifecycle, error) {
	resp := &KubernetesLifecycle{}
	err := c.rancherClient.doUpdate(KUBERNETES_LIFECYCLE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *KubernetesLifecycleClient) List(opts *ListOpts) (*KubernetesLifecycleCollection, error) {
	resp := &KubernetesLifecycleCollection{}
	err := c.rancherClient.doList(KUBERNETES_LIFECYCLE_TYPE, opts, resp)
	return resp, err
}

func (c *KubernetesLifecycleClient) ById(id string) (*KubernetesLifecycle, error) {
	resp := &KubernetesLifecycle{}
	err := c.rancherClient.doById(KUBERNETES_LIFECYCLE_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *KubernetesLifecycleClient) Delete(container *KubernetesLifecycle) error {
	return c.rancherClient.doResourceDelete(KUBERNETES_LIFECYCLE_TYPE, &container.Resource)
}

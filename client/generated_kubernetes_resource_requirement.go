package client

const (
	KUBERNETES_RESOURCE_REQUIREMENT_TYPE = "kubernetesResourceRequirement"
)

type KubernetesResourceRequirement struct {
	Resource

	Limits map[string]interface{} `json:"limits,omitempty" yaml:"limits,omitempty"`

	Requests map[string]interface{} `json:"requests,omitempty" yaml:"requests,omitempty"`
}

type KubernetesResourceRequirementCollection struct {
	Collection
	Data []KubernetesResourceRequirement `json:"data,omitempty"`
}

type KubernetesResourceRequirementClient struct {
	rancherClient *RancherClient
}

type KubernetesResourceRequirementOperations interface {
	List(opts *ListOpts) (*KubernetesResourceRequirementCollection, error)
	Create(opts *KubernetesResourceRequirement) (*KubernetesResourceRequirement, error)
	Update(existing *KubernetesResourceRequirement, updates interface{}) (*KubernetesResourceRequirement, error)
	ById(id string) (*KubernetesResourceRequirement, error)
	Delete(container *KubernetesResourceRequirement) error
}

func newKubernetesResourceRequirementClient(rancherClient *RancherClient) *KubernetesResourceRequirementClient {
	return &KubernetesResourceRequirementClient{
		rancherClient: rancherClient,
	}
}

func (c *KubernetesResourceRequirementClient) Create(container *KubernetesResourceRequirement) (*KubernetesResourceRequirement, error) {
	resp := &KubernetesResourceRequirement{}
	err := c.rancherClient.doCreate(KUBERNETES_RESOURCE_REQUIREMENT_TYPE, container, resp)
	return resp, err
}

func (c *KubernetesResourceRequirementClient) Update(existing *KubernetesResourceRequirement, updates interface{}) (*KubernetesResourceRequirement, error) {
	resp := &KubernetesResourceRequirement{}
	err := c.rancherClient.doUpdate(KUBERNETES_RESOURCE_REQUIREMENT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *KubernetesResourceRequirementClient) List(opts *ListOpts) (*KubernetesResourceRequirementCollection, error) {
	resp := &KubernetesResourceRequirementCollection{}
	err := c.rancherClient.doList(KUBERNETES_RESOURCE_REQUIREMENT_TYPE, opts, resp)
	return resp, err
}

func (c *KubernetesResourceRequirementClient) ById(id string) (*KubernetesResourceRequirement, error) {
	resp := &KubernetesResourceRequirement{}
	err := c.rancherClient.doById(KUBERNETES_RESOURCE_REQUIREMENT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *KubernetesResourceRequirementClient) Delete(container *KubernetesResourceRequirement) error {
	return c.rancherClient.doResourceDelete(KUBERNETES_RESOURCE_REQUIREMENT_TYPE, &container.Resource)
}

package client

const (
	KUBERNETES_PROBE_TYPE = "kubernetesProbe"
)

type KubernetesProbe struct {
	Resource

	Exec KubernetesExecAction `json:"exec,omitempty" yaml:"exec,omitempty"`

	HttpGet KubernetesHTTPGetAction `json:"httpGet,omitempty" yaml:"http_get,omitempty"`

	InitialDelaySeconds int64 `json:"initialDelaySeconds,omitempty" yaml:"initial_delay_seconds,omitempty"`

	TcpSocket KubernetesTCPSocketAction `json:"tcpSocket,omitempty" yaml:"tcp_socket,omitempty"`

	TimeoutSeconds int64 `json:"timeoutSeconds,omitempty" yaml:"timeout_seconds,omitempty"`
}

type KubernetesProbeCollection struct {
	Collection
	Data []KubernetesProbe `json:"data,omitempty"`
}

type KubernetesProbeClient struct {
	rancherClient *RancherClient
}

type KubernetesProbeOperations interface {
	List(opts *ListOpts) (*KubernetesProbeCollection, error)
	Create(opts *KubernetesProbe) (*KubernetesProbe, error)
	Update(existing *KubernetesProbe, updates interface{}) (*KubernetesProbe, error)
	ById(id string) (*KubernetesProbe, error)
	Delete(container *KubernetesProbe) error
}

func newKubernetesProbeClient(rancherClient *RancherClient) *KubernetesProbeClient {
	return &KubernetesProbeClient{
		rancherClient: rancherClient,
	}
}

func (c *KubernetesProbeClient) Create(container *KubernetesProbe) (*KubernetesProbe, error) {
	resp := &KubernetesProbe{}
	err := c.rancherClient.doCreate(KUBERNETES_PROBE_TYPE, container, resp)
	return resp, err
}

func (c *KubernetesProbeClient) Update(existing *KubernetesProbe, updates interface{}) (*KubernetesProbe, error) {
	resp := &KubernetesProbe{}
	err := c.rancherClient.doUpdate(KUBERNETES_PROBE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *KubernetesProbeClient) List(opts *ListOpts) (*KubernetesProbeCollection, error) {
	resp := &KubernetesProbeCollection{}
	err := c.rancherClient.doList(KUBERNETES_PROBE_TYPE, opts, resp)
	return resp, err
}

func (c *KubernetesProbeClient) ById(id string) (*KubernetesProbe, error) {
	resp := &KubernetesProbe{}
	err := c.rancherClient.doById(KUBERNETES_PROBE_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *KubernetesProbeClient) Delete(container *KubernetesProbe) error {
	return c.rancherClient.doResourceDelete(KUBERNETES_PROBE_TYPE, &container.Resource)
}

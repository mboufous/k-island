package kube

import "time"

type Namespace struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Namespaces []Namespace

type Pod struct {
	Name           string     `json:"name"`
	Hostname       string     `json:"hostname"`
	HostIP         string     `json:"host_ip"`
	PodIP          string     `json:"pod_ip"`
	Message        string     `json:"message,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	InitContainers Containers `json:"init_containers,omitempty"`
	Containers     Containers `json:"containers,omitempty"`
	Status         string     `json:"status"`
}

type Pods []Pod

// use ContainerState
type Container struct {
	Name    string         `json:"name"`
	State   ContainerState `json:"state"`
	Ready   bool           `json:"ready"`
	Image   string         `json:"image"`
	ImageID string         `json:"imageID"`
	Started bool           `json:"started"`
}

type Containers []Container

type ContainerState struct {
	Waiting    ContainerStateWaiting    `json:"container_waiting"`
	Running    ContainerStateRunning    `json:"container_running"`
	Terminated ContainerStateTerminated `json:"container_terminated"`
}

type ContainerStateWaiting struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

type ContainerStateRunning struct {
	StartedAt time.Time `json:"startedAt"`
}

type ContainerStateTerminated struct {
	ExitCode   int32     `json:"exit_code"`
	Reason     string    `json:"reason"`
	Message    string    `json:"message"`
	StartedAt  time.Time `json:"startedAt"`
	FinishedAt time.Time `json:"finishedAt"`
}

type Ingress struct {
	Name string      `json:"name"`
	Urls IngressUrls `json:"urls"`
}

type Ingresses []Ingress

type IngressUrl struct {
	ServiceName string `json:"service_name"`
	Url         string `json:"url"`
}

type IngressUrls []IngressUrl

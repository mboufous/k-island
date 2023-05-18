package kube

import (
	"fmt"

	"github.com/mboufous/k-island/utils/log"
	coreV1 "k8s.io/api/core/v1"
	netV1 "k8s.io/api/networking/v1"
)

// TODO: pass context
type KubeService interface {
	GetNamespaces() (Namespaces, error)
	GetPods(namespace string) (Pods, error)
	GetIngresses(namespace string) (Ingresses, error)
}

type kubeService struct {
	kubeClient Client
}

func NewService() (KubeService, error) {
	client, err := NewKubeClient()
	if err != nil {
		return nil, err
	}
	return &kubeService{
		kubeClient: client,
	}, nil
}

func (s *kubeService) GetNamespaces() (Namespaces, error) {
	namespaces, err := s.kubeClient.Namespaces()
	if err != nil {
		return nil, err
	}

	return toNamespaces(namespaces), nil
}

func (s *kubeService) GetIngresses(namespace string) (Ingresses, error) {
	ingresses, err := s.kubeClient.Ingresses(namespace)
	if err != nil {
		return nil, err
	}

	if len(ingresses) == 0 {
		return nil, ErrNoIngressFound
	}

	return toIngresses(ingresses), nil
}
func toUrls(ingress netV1.Ingress) IngressUrls {
	urls := make(IngressUrls, 0, len(ingress.Spec.Rules))
	for _, ingressHost := range ingress.Spec.Rules {
		for _, ingressPath := range ingressHost.HTTP.Paths {
			url := IngressUrl{
				ServiceName: ingressPath.Backend.Service.Name,
				Url:         fmt.Sprintf("https://%s:%d%s", ingressHost.Host, ingressPath.Backend.Service.Port.Number, ingressPath.Path),
			}
			urls = append(urls, url)
		}
	}
	return urls
}

func toIngresses(ingresses []netV1.Ingress) Ingresses {
	appIngresses := make(Ingresses, 0, len(ingresses))
	for _, ingress := range ingresses {
		appIngress := Ingress{
			Name: ingress.Name,
			Urls: toUrls(ingress),
		}
		appIngresses = append(appIngresses, appIngress)
	}
	return appIngresses
}

func (s *kubeService) GetPods(namespace string) (Pods, error) {
	pods, err := s.kubeClient.Pods(namespace)
	if err != nil {
		return nil, err
	}

	if len(pods) == 0 {
		return nil, ErrNoPodFound
	}

	return toPods(pods), nil
}

func toNamespaces(namespaces []coreV1.Namespace) Namespaces {
	appNamespaces := make(Namespaces, 0, len(namespaces))
	for _, ns := range namespaces {
		appNamespace := Namespace{
			Name:      ns.Name,
			CreatedAt: ns.CreationTimestamp.Time,
		}
		appNamespaces = append(appNamespaces, appNamespace)
	}
	return appNamespaces
}

func toContainerState(containerState coreV1.ContainerState) ContainerState {

	appContainerState := ContainerState{
		ContainerStateWaiting{},
		ContainerStateRunning{},
		ContainerStateTerminated{},
	}

	if containerState.Waiting != nil {
		appContainerState.Waiting = ContainerStateWaiting{
			Reason:  containerState.Waiting.Reason,
			Message: containerState.Waiting.Message,
		}
	}

	if containerState.Running != nil {
		appContainerState.Running = ContainerStateRunning{
			StartedAt: containerState.Running.StartedAt.Time,
		}
	}

	if containerState.Terminated != nil {
		appContainerState.Terminated = ContainerStateTerminated{
			ExitCode:   containerState.Terminated.ExitCode,
			Reason:     containerState.Terminated.Reason,
			Message:    containerState.Terminated.Message,
			StartedAt:  containerState.Terminated.StartedAt.Time,
			FinishedAt: containerState.Terminated.FinishedAt.Time,
		}
	}

	return appContainerState
}

func toContainers(containers []coreV1.ContainerStatus) Containers {
	appContainers := make(Containers, 0, len(containers))
	for _, container := range containers {
		appContainer := Container{
			Name:    container.Name,
			Image:   container.Image,
			ImageID: container.ImageID,
			Ready:   container.Ready,
			Started: *container.Started,
			State:   toContainerState(container.State),
		}
		appContainers = append(appContainers, appContainer)
	}
	return appContainers
}

func toPods(pods []coreV1.Pod) Pods {
	appPods := make(Pods, 0, len(pods))
	for _, pod := range pods {
		log.Debug(pod.Status)
		appPod := Pod{
			Name:       pod.Name,
			Hostname:   pod.Spec.Hostname,
			HostIP:     pod.Status.HostIP,
			PodIP:      pod.Status.PodIP,
			Message:    pod.Status.Message,
			CreatedAt:  pod.CreationTimestamp.Time,
			Containers: toContainers(pod.Status.ContainerStatuses),
			Status:     string(pod.Status.Phase),
		}
		appPods = append(appPods, appPod)
	}
	return appPods
}

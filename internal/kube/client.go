package kube

import (
	"context"
	"fmt"

	"github.com/mboufous/k-island/utils/log"
	coreV1 "k8s.io/api/core/v1"
	netV1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
)

type Client interface {
	Namespaces() ([]coreV1.Namespace, error)
	Pods(ns string) ([]coreV1.Pod, error)
	Ingresses(ns string) ([]netV1.Ingress, error)
}

type KubeClient struct {
	clientset *kubernetes.Clientset
}

func NewKubeClient() (*KubeClient, error) {
	config, err := newKubeConfig()
	if err != nil {
		return nil, fmt.Errorf("error creating config: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("NewKubeClient: %w", err)
	}

	return &KubeClient{
		clientset: clientset,
	}, nil
}

func newKubeConfig() (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Warnf("in-cluster: %s", err.Error())
		config, err = outOfClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("out-cluster: %w", err)
		}
	}
	return config, nil
}

func outOfClusterConfig() (*rest.Config, error) {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	rules.DefaultClientConfig = &clientcmd.DefaultClientConfig
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, configOverrides)
	return kubeConfig.ClientConfig()
}

func (c *KubeClient) Namespaces() ([]coreV1.Namespace, error) {
	namespaces, err := c.clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("kube client: %w", err)
	}
	log.Debugf("Found %d namespace", len(namespaces.Items))
	return namespaces.Items, nil
}

func (c *KubeClient) Pods(ns string) ([]coreV1.Pod, error) {
	pods, err := c.clientset.CoreV1().Pods(ns).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("kube client: %w", err)
	}
	log.Debugf("Found %d Pods in namespace %s", len(pods.Items), ns)
	return pods.Items, nil
}

func (c *KubeClient) Ingresses(ns string) ([]netV1.Ingress, error) {
	ingresses, _ := c.clientset.NetworkingV1().Ingresses(ns).List(context.TODO(), metav1.ListOptions{})
	return ingresses.Items, nil
	// fmt.Println("------- Found", len(ingresses.Items))
	// for _, ingress := range ingresses.Items {
	// 	fmt.Println(" ---- ", ingress.Name)
	// 	for _, rule := range ingress.Spec.Rules {
	// 		fmt.Println(" ------ ", rule.Host)
	// 		for _, path := range rule.HTTP.Paths {
	// 			fmt.Println(" ------ ", path.Path)

	// 		}
	// 	}

	// }
}

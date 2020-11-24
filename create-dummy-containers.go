package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" { // check if machine has home directory.
		// read kubeconfig flag. if not provided use config file $HOME/.kube/config
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 50; i++ {
		// build the pod defination we want to deploy
		pod := getPodObject(strconv.Itoa(i))

		// now create the pod in kubernetes cluster using the clientset
		pod, err = clientset.CoreV1().Pods(pod.Namespace).Create(pod)
		if err != nil {
			panic(err)
		}
		fmt.Println("Pod %d created successfully...", i)
		time.Sleep(5 * time.Second)

	}
}

func getPodObject(number string) *core.Pod {
        return &core.Pod{
                ObjectMeta: metav1.ObjectMeta{
                        Name:      "my-test-pod-" + number,
                        Namespace: "default",
                        Labels: map[string]string{
                                "app": "demo",
                                "number": number,
                        },
                },
                Spec: core.PodSpec{
                        Containers: []core.Container{
                                {
                                        Name:            "busybox",
                                        Image:           "busybox",
                                        ImagePullPolicy: core.PullIfNotPresent,
                                        Command: []string{
                                                "sleep",
                                                "3600",
                                        },
                                },
                        },
                },
        }
}


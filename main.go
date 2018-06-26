/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"k8s.io/client-go/rest"
	clientset "k8s.io/sample-controller/pkg/client/clientset/versioned"
	informers "k8s.io/sample-controller/pkg/client/informers/externalversions"
	"k8s.io/sample-controller/pkg/signals"
	"github.com/mitchellh/go-homedir"

	"path/filepath"
)

var (
	masterURL  string
	kubeconfig string
)

// retrieve the Kubernetes config
func getKubernetesConfig(masterURL string, kubeconfig string) (*rest.Config, error) {
	if kubeconfig == "" {
		// construct the path to resolve to `~/.kube/config`
		userHomeDirectory, err := homedir.Dir()
		if err == nil {
			kubeconfig = filepath.Join(userHomeDirectory, ".kube/config")
		}
	}

	// create the config from the path (from outside of the cluster)
	config, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err == nil {
		return config, nil
	}

	// try to get the config from inside the cluster
	config, err2 := rest.InClusterConfig()
	if err2 != nil {
		//Return original error as we are not inside a container running on the cluster
		return config, err
	}

	return config, nil
}

func main() {
	flag.Parse()

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	cfg, err := getKubernetesConfig(masterURL, kubeconfig)
	if err != nil {
		glog.Fatalf("Error building kubernetes config: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	exampleClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %s", err.Error())
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	exampleInformerFactory := informers.NewSharedInformerFactory(exampleClient, time.Second*30)

	controller := NewController(kubeClient, exampleClient,
		kubeInformerFactory.Apps().V1().Deployments(),
		exampleInformerFactory.Samplecontroller().V1alpha1().Foos())

	go kubeInformerFactory.Start(stopCh)
	go exampleInformerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}

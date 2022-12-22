package utils

import (
	"context"
	"flag"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodHelper struct {
	k8sclient *kubernetes.Clientset
}

func NewPodHelper() *PodHelper {
	config, err := rest.InClusterConfig()
	if err != nil {
		// try using kube config
		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return &PodHelper{k8sclient: client}
}

func (kh *PodHelper) CreateOracleFeeder(name string, owner string, oracle string, chainID string) error {
	fields := make(map[string]string)
	fields["oracle"] = oracle
	fields["chainID"] = chainID
	// fields["publickey"] = "0xc"
	fields["owner"] = owner
	privatekeyenv := corev1.EnvVar{Name: "PRIVATE_KEY", ValueFrom: &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{Key: ".private", LocalObjectReference: corev1.LocalObjectReference{Name: name}}}}
	deployedcontractenv := corev1.EnvVar{Name: "DEPLOYED_CONTRACT", Value: oracle}
	chainidenv := corev1.EnvVar{Name: "CHAIN_ID", Value: chainID}
	sleepsecondenv := corev1.EnvVar{Name: "SLEEP_SECONDS", Value: "2"}
	deviationenv := corev1.EnvVar{Name: "DEVIATION_PERMILLE", Value: "0"}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: fields,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  name,
					Image: "",
					Env:   []corev1.EnvVar{privatekeyenv, deployedcontractenv, chainidenv, sleepsecondenv, deviationenv},
				},
			},
		},
	}

	result, err := kh.k8sclient.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	log.Infof("Pod %s started\n", result.GetObjectMeta().GetName())
	return err

}

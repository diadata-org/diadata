package k8util

import (
	"context"

	"github.com/diadata-org/diadata/pkg/utils"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type K8Bridge struct {
	clientset *kubernetes.Clientset
	context   context.Context
	namespace string
}

func New(clientset *kubernetes.Clientset, namespace string) *K8Bridge {
	return &K8Bridge{clientset: clientset, context: context.TODO(), namespace: namespace}
}

func (k *K8Bridge) GetKeys(keyname string) (*v1.Secret, error) {
	return k.clientset.CoreV1().Secrets(k.namespace).Get(k.context, keyname, metav1.GetOptions{})
}

func (k *K8Bridge) GenerateKey(keyname string) (publickey string, err error) {
	{
		var privatekey string
		data := make(map[string][]byte)
		publickey, privatekey = utils.NewKeyPair()
		data[".private"] = []byte(privatekey)
		data[".public"] = []byte(publickey)

		secret := &v1.Secret{}
		secret.APIVersion = "v1"
		secret.Name = keyname
		secret.Data = data
		secret.Type = "Opaque"
		secret.Kind = "Secret"
		_, err = k.clientset.CoreV1().Secrets(k.namespace).Create(context.TODO(), secret, metav1.CreateOptions{})
		return
	}

}

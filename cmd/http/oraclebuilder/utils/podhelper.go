package utils

import (
	"context"
	"flag"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var log = logrus.New()

type PodHelper struct {
	k8sclient *kubernetes.Clientset
	Image     string
	NameSpace string
}

func NewPodHelper(image, namespace string) *PodHelper {
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
	return &PodHelper{k8sclient: client, Image: image, NameSpace: namespace}
}

func (kh *PodHelper) CreateOracleFeeder(feederID string, owner string, oracle string, chainID string, symbols, blockchainnode string) error {
	fields := make(map[string]string)
	fields["oracle"] = oracle
	fields["chainID"] = chainID
	// fields["publickey"] = "0xc"
	fields["owner"] = owner

	// -- oracle config
	publickeyenv := corev1.EnvVar{Name: "ORACLE_PUBLICKEY", ValueFrom: &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{Key: ".public", LocalObjectReference: corev1.LocalObjectReference{Name: feederID}}}}
	deployedcontractenv := corev1.EnvVar{Name: "DEPLOYED_CONTRACT", Value: oracle}
	chainidenv := corev1.EnvVar{Name: "ORACLE_CHAINID", Value: chainID}

	signerservice := corev1.EnvVar{Name: "ORACLE_SIGNER", Value: "signer.dia-oracle-feeder:50052"}

	sleepsecondenv := corev1.EnvVar{Name: "ORACLE_SLEEPSECONDS", Value: "2"}
	deviationenv := corev1.EnvVar{Name: "DEVIATION_PERMILLE", Value: "0"}
	frequencyseconds := corev1.EnvVar{Name: "ORACLE_FREQUENCYSECONDS", Value: "1"}
	oracletype := corev1.EnvVar{Name: "ORACLE_ORACLETYPE", Value: "3"}
	oraclesymbols := corev1.EnvVar{Name: "ORACLE_SYMBOLS", Value: symbols}
	oraclefeederid := corev1.EnvVar{Name: "ORACLE_FEEDERID", Value: feederID}
	// -- oracle config ends here

	// ---
	postgreshost := corev1.EnvVar{Name: "POSTGRES_HOST", Value: "dia-postgresql.dia-db"}
	postgresuser := corev1.EnvVar{Name: "POSTGRES_USER", ValueFrom: &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{Key: "user", LocalObjectReference: corev1.LocalObjectReference{Name: "user.graphqlserver"}}}}
	postgrespassword := corev1.EnvVar{Name: "POSTGRES_PASSWORD", ValueFrom: &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{Key: "password", LocalObjectReference: corev1.LocalObjectReference{Name: "user.graphqlserver"}}}}
	postgresdb := corev1.EnvVar{Name: "POSTGRES_DB", Value: "postgres"}
	updateconfigseconds := corev1.EnvVar{Name: "ORACLE_UPDATECONFIGSECONDS", Value: "120"}
	useenv := corev1.EnvVar{Name: "USE_ENV", Value: "true"}
	//---

	imagepullrequest := corev1.LocalObjectReference{Name: "all-icr-io"}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:   feederID,
			Labels: fields,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  feederID,
					Image: kh.Image,
					Env: []corev1.EnvVar{publickeyenv, deployedcontractenv, chainidenv,
						sleepsecondenv, deviationenv, frequencyseconds, oracletype,
						oraclesymbols, oraclefeederid, postgreshost, postgresuser, signerservice,
						postgrespassword, postgresdb, updateconfigseconds, useenv},
				},
			},
			ImagePullSecrets: []corev1.LocalObjectReference{imagepullrequest},
		},
	}

	result, err := kh.k8sclient.CoreV1().Pods(kh.NameSpace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	log.Infof("Pod %s started\n", result.GetObjectMeta().GetName())
	return err

}

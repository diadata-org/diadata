package k8util

import (
	"context"
	"fmt"
	"time"

	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

var log = logrus.New()

type K8Bridge struct {
	clientSet     *kubernetes.Clientset
	context       context.Context
	namespace     string
	SignerURL     string
	DiaRestURL    string
	DiaGraphqlURL string
	PostgresHost  string
	Affinity      string
	Image         string
}

func New(clientset *kubernetes.Clientset, namespace string, signerurl, diaresturl, diagraphqlurl, postgreshost, affinity, image string) *K8Bridge {
	return &K8Bridge{clientSet: clientset, context: context.Background(), namespace: namespace, SignerURL: signerurl, DiaRestURL: diaresturl, DiaGraphqlURL: diagraphqlurl, PostgresHost: postgreshost, Affinity: affinity, Image: image}
}

func (k *K8Bridge) GetKeys(keyname string) (*v1.Secret, error) {
	return k.clientSet.CoreV1().Secrets(k.namespace).Get(k.context, keyname, metav1.GetOptions{})
}

func (k *K8Bridge) GenerateKey(keyName string) (publicKey string, err error) {
	{
		var privatekey string
		data := make(map[string][]byte)
		publicKey, privatekey = utils.NewKeyPair()
		data[".private"] = []byte(privatekey)
		data[".public"] = []byte(publicKey)

		secret := &v1.Secret{}
		secret.APIVersion = "v1"
		secret.Name = keyName
		secret.Data = data
		secret.Type = "Opaque"
		secret.Kind = "Secret"
		_, err = k.clientSet.CoreV1().Secrets(k.namespace).Create(context.TODO(), secret, metav1.CreateOptions{})
		return
	}

}

func (kh *K8Bridge) CreateOracleFeeder(ctx context.Context, feederID, creator, feederAddress, oracle string, chainID string, symbols, feedSelection, blockchainnode string, frequency, sleepSeconds, deviationPermille, mandatoryFrequency string) error {

	envvars := kh.PodEnvironmentVariables(feederID, creator, oracle, chainID, symbols, feedSelection, blockchainnode, frequency, sleepSeconds, deviationPermille, mandatoryFrequency)

	// Define an image pull request
	imagepullrequest := v1.LocalObjectReference{Name: "all-icr-io"}

	// Create a Pod configuration
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: feederID,
			Labels: map[string]string{
				"oracle":        oracle,
				"chainID":       chainID,
				"owner":         creator,
				"feederAddress": feederAddress,
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  feederID,
					Image: kh.Image,
					Env:   envvars,
				},
			},
			ImagePullSecrets: []v1.LocalObjectReference{imagepullrequest},
			Affinity: &v1.Affinity{
				NodeAffinity: &v1.NodeAffinity{
					RequiredDuringSchedulingIgnoredDuringExecution: &v1.NodeSelector{
						NodeSelectorTerms: []v1.NodeSelectorTerm{
							{
								MatchExpressions: []v1.NodeSelectorRequirement{
									{
										Key:      "ibm-cloud.kubernetes.io/worker-pool-name",
										Operator: v1.NodeSelectorOpIn,
										Values: []string{
											kh.Affinity,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Create the Pod

	result, err := kh.clientSet.CoreV1().Pods(kh.namespace).Create(ctx, pod, metav1.CreateOptions{})
	if err != nil {
		log.Errorf("Failed to start pod %s: %v", feederID, err)
		return err
	}
	log.Infof("Pod %s started\n", result.GetObjectMeta().GetName())
	return err

}

func (kh *K8Bridge) DeleteOracleFeeder(ctx context.Context, feederID string) error {
	_, err := kh.clientSet.CoreV1().Pods(kh.namespace).Get(ctx, feederID, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			log.Infof("Pod %s not found, no need to delete\n", feederID)
			return nil
		}
		return err // return error if it's other than not found
	}

	// if no error, that means pod exists and needs to be deleted.
	deletePolicy := metav1.DeletePropagationForeground
	err = kh.clientSet.CoreV1().Pods(kh.namespace).Delete(ctx, feederID, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})

	if err != nil {
		return err
	}

	log.Infof("Pod %s deleted\n", feederID)
	return nil
}

// RestartOracleFeeder restarts an Oracle Feeder Pod in Kubernetes.
func (kh *K8Bridge) RestartOracleFeeder(ctx context.Context, feederID, creator, feederAddress, oracle string, chainID string, symbols, feedSelection, blockchainnode string, frequency, sleepSeconds, deviationPermille, mandatoryFrequency string) (err error) {

	_, err = kh.clientSet.CoreV1().Pods(kh.namespace).Get(ctx, feederID, metav1.GetOptions{})
	if err != nil {
		log.Infof("Pod  RestartOracleFeeder %s process %v \n", feederID, err)

		if errors.IsNotFound(err) {
			log.Infof("Pod %s not found, no need to delete\n", feederID)
		} else {
			return err
		}
		err = kh.CreateOracleFeeder(ctx, feederID, creator, feederAddress, oracle, chainID, symbols, feedSelection, blockchainnode, frequency, sleepSeconds, deviationPermille, mandatoryFrequency)
		if err != nil {
			log.Errorf("Pod %s start err\n", err)
			return
		}
		log.Infof("Pod %s started\n", feederID)
	} else {
		err = kh.clientSet.CoreV1().Pods(kh.namespace).Delete(ctx, feederID, metav1.DeleteOptions{})
		//if err != nil {
		//	return err
		//}
		deleteerr := kh.waitPodDeleted(ctx, oracle, func() {
			time.Sleep(1000 * time.Millisecond)
			err = kh.CreateOracleFeeder(ctx, feederID, creator, feederAddress, oracle, chainID, symbols, feedSelection, blockchainnode, frequency, sleepSeconds, deviationPermille, mandatoryFrequency)
			if err != nil {
				log.Errorf("Pod %s start err\n", err)
				return
			}
			log.Infof("Pod %s started\n", feederID)
		})
		if deleteerr != nil {
			log.Infof("Pod %s delete err\n", deleteerr.Error())
		}
		log.Infof("Pod %s deleted\n", feederID)
	}

	return err
}

// podWatcher creates a Pod watcher for a given Oracle address.
func (kh *K8Bridge) podWatcher(ctx context.Context, oracleaddress string) (watch.Interface, error) {
	labelSelector := fmt.Sprintf("oracle=%s", oracleaddress)
	log.Infof("Creating watcher for POD with label: %s", labelSelector)

	opts := metav1.ListOptions{
		TypeMeta:      metav1.TypeMeta{},
		LabelSelector: labelSelector,
		FieldSelector: "",
	}

	return kh.clientSet.CoreV1().Pods(kh.namespace).Watch(ctx, opts)
}

// waitPodDeleted waits for a Pod to be deleted.
func (kh *K8Bridge) waitPodDeleted(ctx context.Context, resName string, run func()) error {
	watcher, err := kh.podWatcher(ctx, resName)
	if err != nil {
		return err
	}

	defer watcher.Stop()

	for {
		select {
		case event := <-watcher.ResultChan():

			if event.Type == watch.Deleted {
				log.Infof("The POD \"%s\" is deleted", resName)
				run()

				return nil
			}

		case <-ctx.Done():
			log.Infof("Exit from waitPodDeleted for POD \"%s\" because the context is done", resName)
			return nil
		}
	}
}

// PodEnvironmentVariables generates environment variables for the Oracle Feeder Pod.
func (kh *K8Bridge) PodEnvironmentVariables(feederID string, owner string, oracle string, chainID string, symbols, feedSelection, blockchainnode string, frequency, sleepSeconds, deviationPermille, mandatoryFrequency string) (vars []v1.EnvVar) {

	// Oracle configuration
	publickeyenv := v1.EnvVar{Name: "ORACLE_PUBLICKEY", ValueFrom: &v1.EnvVarSource{SecretKeyRef: &v1.SecretKeySelector{Key: ".public", LocalObjectReference: v1.LocalObjectReference{Name: feederID}}}}
	deployedcontractenv := v1.EnvVar{Name: "DEPLOYED_CONTRACT", Value: oracle}
	chainidenv := v1.EnvVar{Name: "ORACLE_CHAINID", Value: chainID}
	signerservice := v1.EnvVar{Name: "ORACLE_SIGNER", Value: kh.SignerURL}
	sleepsecondenv := v1.EnvVar{Name: "ORACLE_SLEEPSECONDS", Value: sleepSeconds}
	deviationenv := v1.EnvVar{Name: "ORACLE_DEVIATIONPERMILLE", Value: deviationPermille}
	frequencyseconds := v1.EnvVar{Name: "ORACLE_FREQUENCYSECONDS", Value: frequency}
	oracletype := v1.EnvVar{Name: "ORACLE_ORACLETYPE", Value: "3"}
	oraclesymbols := v1.EnvVar{Name: "ORACLE_SYMBOLS", Value: symbols}
	oraclefeederid := v1.EnvVar{Name: "ORACLE_FEEDERID", Value: feederID}
	blockchainnodeenv := v1.EnvVar{Name: "ORACLE_BLOCKCHAINNODE", Value: blockchainnode}
	mandatoryfrequencyenv := v1.EnvVar{Name: "ORACLE_MANDATORYFREQUENCY", Value: mandatoryFrequency}
	feedSelectionenv := v1.EnvVar{Name: "ORACLE_FEEDSELECTION", Value: feedSelection}

	diaRestAPIenv := v1.EnvVar{Name: "DIA_REST_URL", Value: kh.DiaRestURL}

	diaGraphqlenv := v1.EnvVar{Name: "DIA_GRAPHQL_URL", Value: kh.DiaGraphqlURL}

	// Append all corev1.EnvVar instances to vars
	vars = append(vars, diaRestAPIenv, diaGraphqlenv, publickeyenv, deployedcontractenv, chainidenv, signerservice, sleepsecondenv, deviationenv, frequencyseconds, oracletype, oraclesymbols, oraclefeederid, blockchainnodeenv, mandatoryfrequencyenv, feedSelectionenv)

	// ---
	postgreshost := v1.EnvVar{Name: "POSTGRES_HOST", Value: kh.PostgresHost}
	postgresuser := v1.EnvVar{Name: "POSTGRES_USER", ValueFrom: &v1.EnvVarSource{SecretKeyRef: &v1.SecretKeySelector{Key: "user", LocalObjectReference: v1.LocalObjectReference{Name: "user.graphqlserver"}}}}
	postgrespassword := v1.EnvVar{Name: "POSTGRES_PASSWORD", ValueFrom: &v1.EnvVarSource{SecretKeyRef: &v1.SecretKeySelector{Key: "password", LocalObjectReference: v1.LocalObjectReference{Name: "user.graphqlserver"}}}}
	postgresdb := v1.EnvVar{Name: "POSTGRES_DB", Value: "postgres"}
	updateconfigseconds := v1.EnvVar{Name: "ORACLE_UPDATECONFIGSECONDS", Value: "120"}
	useenv := v1.EnvVar{Name: "USE_ENV", Value: "true"}

	vars = append(vars, postgreshost, postgresuser, postgrespassword, postgresdb, updateconfigseconds, useenv)
	//---
	// -- oracle config ends here

	return
}

package utils

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

var log = logrus.New()

// PodHelper is a utility struct for working with Kubernetes Pods.
type PodHelper struct {
	k8sclient     *kubernetes.Clientset
	Image         string
	NameSpace     string
	Affinity      string
	SignerURL     string
	DiaRestURL    string
	DiaGraphqlURL string
	PostgresHost  string
}

// NewPodHelper creates a new instance of PodHelper.
func NewPodHelper(image, namespace, affinity, signerURL, diaRestURL, diaGraphqlURL, postgresHost string) *PodHelper {
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
	return &PodHelper{k8sclient: client, Image: image, NameSpace: namespace, Affinity: affinity, SignerURL: signerURL, DiaRestURL: diaRestURL, DiaGraphqlURL: diaGraphqlURL, PostgresHost: postgresHost}
}

// CreateOracleFeeder creates a new Oracle Feeder Pod in Kubernetes.
func (kh *PodHelper) CreateOracleFeeder(ctx context.Context, feederID, creator, feederAddress, oracle string, chainID string, symbols, feedSelection, blockchainnode string, frequency, sleepSeconds, deviationPermille, mandatoryFrequency string) error {

	envvars := kh.PodEnvironmentVariables(feederID, creator, oracle, chainID, symbols, feedSelection, blockchainnode, frequency, sleepSeconds, deviationPermille, mandatoryFrequency)

	// Define an image pull request
	imagepullrequest := corev1.LocalObjectReference{Name: "all-icr-io"}

	// Create a Pod configuration
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: feederID,
			Labels: map[string]string{
				"oracle":        oracle,
				"chainID":       chainID,
				"owner":         creator,
				"feederAddress": feederAddress,
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  feederID,
					Image: kh.Image,
					Env:   envvars,
				},
			},
			ImagePullSecrets: []corev1.LocalObjectReference{imagepullrequest},
			Affinity: &corev1.Affinity{
				NodeAffinity: &corev1.NodeAffinity{
					RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{
						NodeSelectorTerms: []corev1.NodeSelectorTerm{
							{
								MatchExpressions: []corev1.NodeSelectorRequirement{
									{
										Key:      "ibm-cloud.kubernetes.io/worker-pool-name",
										Operator: corev1.NodeSelectorOpIn,
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

	result, err := kh.k8sclient.CoreV1().Pods(kh.NameSpace).Create(ctx, pod, metav1.CreateOptions{})
	if err != nil {
		log.Errorf("Failed to start pod %s: %v", feederID, err)
		return err
	}
	log.Infof("Pod %s started\n", result.GetObjectMeta().GetName())
	return err

}

// PodEnvironmentVariables generates environment variables for the Oracle Feeder Pod.
func (kh *PodHelper) PodEnvironmentVariables(feederID string, owner string, oracle string, chainID string, symbols, feedSelection, blockchainnode string, frequency, sleepSeconds, deviationPermille, mandatoryFrequency string) (vars []corev1.EnvVar) {

	// Oracle configuration
	publickeyenv := corev1.EnvVar{Name: "ORACLE_PUBLICKEY", ValueFrom: &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{Key: ".public", LocalObjectReference: corev1.LocalObjectReference{Name: feederID}}}}
	deployedcontractenv := corev1.EnvVar{Name: "DEPLOYED_CONTRACT", Value: oracle}
	chainidenv := corev1.EnvVar{Name: "ORACLE_CHAINID", Value: chainID}
	signerservice := corev1.EnvVar{Name: "ORACLE_SIGNER", Value: kh.SignerURL}
	sleepsecondenv := corev1.EnvVar{Name: "ORACLE_SLEEPSECONDS", Value: sleepSeconds}
	deviationenv := corev1.EnvVar{Name: "ORACLE_DEVIATIONPERMILLE", Value: deviationPermille}
	frequencyseconds := corev1.EnvVar{Name: "ORACLE_FREQUENCYSECONDS", Value: frequency}
	oracletype := corev1.EnvVar{Name: "ORACLE_ORACLETYPE", Value: "3"}
	oraclesymbols := corev1.EnvVar{Name: "ORACLE_SYMBOLS", Value: symbols}
	oraclefeederid := corev1.EnvVar{Name: "ORACLE_FEEDERID", Value: feederID}
	blockchainnodeenv := corev1.EnvVar{Name: "ORACLE_BLOCKCHAINNODE", Value: blockchainnode}
	mandatoryfrequencyenv := corev1.EnvVar{Name: "ORACLE_MANDATORYFREQUENCY", Value: mandatoryFrequency}
	feedSelectionenv := corev1.EnvVar{Name: "ORACLE_FEEDSELECTION", Value: feedSelection}

	diaRestAPIenv := corev1.EnvVar{Name: "DIA_REST_URL", Value: kh.DiaRestURL}

	diaGraphqlenv := corev1.EnvVar{Name: "DIA_GRAPHQL_URL", Value: kh.DiaGraphqlURL}

	// Append all corev1.EnvVar instances to vars
	vars = append(vars, diaRestAPIenv, diaGraphqlenv, publickeyenv, deployedcontractenv, chainidenv, signerservice, sleepsecondenv, deviationenv, frequencyseconds, oracletype, oraclesymbols, oraclefeederid, blockchainnodeenv, mandatoryfrequencyenv, feedSelectionenv)

	// ---
	postgreshost := corev1.EnvVar{Name: "POSTGRES_HOST", Value: kh.PostgresHost}
	postgresuser := corev1.EnvVar{Name: "POSTGRES_USER", ValueFrom: &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{Key: "user", LocalObjectReference: corev1.LocalObjectReference{Name: "user.graphqlserver"}}}}
	postgrespassword := corev1.EnvVar{Name: "POSTGRES_PASSWORD", ValueFrom: &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{Key: "password", LocalObjectReference: corev1.LocalObjectReference{Name: "user.graphqlserver"}}}}
	postgresdb := corev1.EnvVar{Name: "POSTGRES_DB", Value: "postgres"}
	updateconfigseconds := corev1.EnvVar{Name: "ORACLE_UPDATECONFIGSECONDS", Value: "120"}
	useenv := corev1.EnvVar{Name: "USE_ENV", Value: "true"}

	vars = append(vars, postgreshost, postgresuser, postgrespassword, postgresdb, updateconfigseconds, useenv)
	//---
	// -- oracle config ends here

	return
}

// UpdateOracleFeeder updates an existing Oracle Feeder Pod in Kubernetes.
func (kh *PodHelper) UpdateOracleFeeder(ctx context.Context, feederID string, owner string, oracle string, chainID string, symbols, feedSelection, blockchainnode string, frequency, sleepSeconds, deviationPermille, mandatoryFrequency string) error {
	fields := make(map[string]string)
	fields["oracle"] = oracle
	fields["chainID"] = chainID
	fields["owner"] = owner

	// Generate environment variables for the updated Pod
	envvars := kh.PodEnvironmentVariables(feederID, owner, oracle, chainID, symbols, feedSelection, blockchainnode, frequency, sleepSeconds, deviationPermille, mandatoryFrequency)

	// Define an image pull request
	imagepullrequest := corev1.LocalObjectReference{Name: "all-icr-io"}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: feederID,
			Labels: map[string]string{
				"oracle":  oracle,
				"chainID": chainID,
				"owner":   owner,
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  feederID,
					Image: kh.Image,
					Env:   envvars,
				},
			},
			ImagePullSecrets: []corev1.LocalObjectReference{imagepullrequest},
			Affinity: &corev1.Affinity{
				NodeAffinity: &corev1.NodeAffinity{
					RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{
						NodeSelectorTerms: []corev1.NodeSelectorTerm{
							{
								MatchExpressions: []corev1.NodeSelectorRequirement{
									{
										Key:      "ibm-cloud.kubernetes.io/worker-pool-name",
										Operator: corev1.NodeSelectorOpIn,
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

	result, err := kh.k8sclient.CoreV1().Pods(kh.NameSpace).Update(ctx, pod, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	log.Infof("Pod %s started\n", result.GetObjectMeta().GetName())
	return err

}

// DeleteOracleFeeder deletes an Oracle Feeder Pod in Kubernetes.
func (kh *PodHelper) DeleteOracleFeeder(ctx context.Context, feederID string) error {
	_, err := kh.k8sclient.CoreV1().Pods(kh.NameSpace).Get(ctx, feederID, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			log.Infof("Pod %s not found, no need to delete\n", feederID)
			return nil
		}
		return err // return error if it's other than not found
	}

	// if no error, that means pod exists and needs to be deleted.
	deletePolicy := metav1.DeletePropagationForeground
	err = kh.k8sclient.CoreV1().Pods(kh.NameSpace).Delete(ctx, feederID, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})

	if err != nil {
		return err
	}

	log.Infof("Pod %s deleted\n", feederID)
	return nil
}

// RestartOracleFeeder restarts an Oracle Feeder Pod in Kubernetes.
func (kh *PodHelper) RestartOracleFeeder(ctx context.Context, feederID string, oracleconfig dia.OracleConfig) (err error) {

	_, err = kh.k8sclient.CoreV1().Pods(kh.NameSpace).Get(ctx, feederID, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			log.Infof("Pod %s not found, no need to delete\n", feederID)
		} else {
			return err
		}
		err = kh.CreateOracleFeeder(ctx, feederID, oracleconfig.Owner, oracleconfig.FeederAddress, oracleconfig.Address, oracleconfig.ChainID, strings.Join(oracleconfig.Symbols[:], ","), oracleconfig.FeederSelection, oracleconfig.BlockchainNode, oracleconfig.Frequency, oracleconfig.SleepSeconds, oracleconfig.DeviationPermille, oracleconfig.MandatoryFrequency)
		if err != nil {
			log.Errorf("Pod %s start err\n", err)
			return
		}
		log.Infof("Pod %s started\n", feederID)
	} else {
		err = kh.k8sclient.CoreV1().Pods(kh.NameSpace).Delete(ctx, feederID, metav1.DeleteOptions{})
		//if err != nil {
		//	return err
		//}
		deleteerr := kh.waitPodDeleted(ctx, oracleconfig.Address, func() {
			time.Sleep(1000 * time.Millisecond)
			err = kh.CreateOracleFeeder(ctx, feederID, oracleconfig.Owner, oracleconfig.FeederAddress, oracleconfig.Address, oracleconfig.ChainID, strings.Join(oracleconfig.Symbols[:], ","), oracleconfig.FeederSelection, oracleconfig.BlockchainNode, oracleconfig.Frequency, oracleconfig.SleepSeconds, oracleconfig.DeviationPermille, oracleconfig.MandatoryFrequency)
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
func (kh *PodHelper) podWatcher(ctx context.Context, oracleaddress string) (watch.Interface, error) {
	labelSelector := fmt.Sprintf("oracle=%s", oracleaddress)
	log.Infof("Creating watcher for POD with label: %s", labelSelector)

	opts := metav1.ListOptions{
		TypeMeta:      metav1.TypeMeta{},
		LabelSelector: labelSelector,
		FieldSelector: "",
	}

	return kh.k8sclient.CoreV1().Pods(kh.NameSpace).Watch(ctx, opts)
}

// waitPodDeleted waits for a Pod to be deleted.
func (kh *PodHelper) waitPodDeleted(ctx context.Context, resName string, run func()) error {
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

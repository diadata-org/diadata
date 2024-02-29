package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

type Configs []struct {
	Metadata struct {
		Name              string    `json:"name"`
		Namespace         string    `json:"namespace"`
		UID               string    `json:"uid"`
		ResourceVersion   string    `json:"resourceVersion"`
		Generation        int       `json:"generation"`
		CreationTimestamp time.Time `json:"creationTimestamp"`
		Labels            struct {
			App                      string `json:"app"`
			AppKubernetesIoManagedBy string `json:"app.kubernetes.io/managed-by"`
			Chart                    string `json:"chart"`
			Heritage                 string `json:"heritage"`
			Release                  string `json:"release"`
		} `json:"labels"`
		Annotations struct {
			DeploymentKubernetesIoRevision string `json:"deployment.kubernetes.io/revision"`
			MetaHelmShReleaseName          string `json:"meta.helm.sh/release-name"`
			MetaHelmShReleaseNamespace     string `json:"meta.helm.sh/release-namespace"`
		} `json:"annotations"`
		ManagedFields []struct {
			Manager    string    `json:"manager"`
			Operation  string    `json:"operation"`
			APIVersion string    `json:"apiVersion"`
			Time       time.Time `json:"time"`
			FieldsType string    `json:"fieldsType"`
			FieldsV1   struct {
				FMetadata struct {
					FAnnotations struct {
						NAMING_FAILED struct {
						} `json:"."`
						FMetaHelmShReleaseName struct {
						} `json:"f:meta.helm.sh/release-name"`
						FMetaHelmShReleaseNamespace struct {
						} `json:"f:meta.helm.sh/release-namespace"`
					} `json:"f:annotations"`
					FLabels struct {
						NAMING_FAILED struct {
						} `json:"."`
						FApp struct {
						} `json:"f:app"`
						FAppKubernetesIoManagedBy struct {
						} `json:"f:app.kubernetes.io/managed-by"`
						FChart struct {
						} `json:"f:chart"`
						FHeritage struct {
						} `json:"f:heritage"`
						FRelease struct {
						} `json:"f:release"`
					} `json:"f:labels"`
				} `json:"f:metadata"`
				FSpec struct {
					FProgressDeadlineSeconds struct {
					} `json:"f:progressDeadlineSeconds"`
					FRevisionHistoryLimit struct {
					} `json:"f:revisionHistoryLimit"`
					FSelector struct {
					} `json:"f:selector"`
					FStrategy struct {
						FRollingUpdate struct {
							NAMING_FAILED struct {
							} `json:"."`
							FMaxSurge struct {
							} `json:"f:maxSurge"`
							FMaxUnavailable struct {
							} `json:"f:maxUnavailable"`
						} `json:"f:rollingUpdate"`
						FType struct {
						} `json:"f:type"`
					} `json:"f:strategy"`
					FTemplate struct {
						FMetadata struct {
							FAnnotations struct {
								NAMING_FAILED struct {
								} `json:"."`
								FRollme struct {
								} `json:"f:rollme"`
							} `json:"f:annotations"`
							FLabels struct {
								NAMING_FAILED struct {
								} `json:"."`
								FApp struct {
								} `json:"f:app"`
								FChart struct {
								} `json:"f:chart"`
								FHeritage struct {
								} `json:"f:heritage"`
								FRelease struct {
								} `json:"f:release"`
							} `json:"f:labels"`
						} `json:"f:metadata"`
						FSpec struct {
							FAffinity struct {
								NAMING_FAILED struct {
								} `json:"."`
								FNodeAffinity struct {
									NAMING_FAILED struct {
									} `json:"."`
									FRequiredDuringSchedulingIgnoredDuringExecution struct {
									} `json:"f:requiredDuringSchedulingIgnoredDuringExecution"`
								} `json:"f:nodeAffinity"`
							} `json:"f:affinity"`
							FContainers struct {
								KNameDeribitSolanaOracle struct {
									NAMING_FAILED struct {
									} `json:"."`
									FEnv struct {
										NAMING_FAILED struct {
										} `json:"."`
										KNameEXECMODE struct {
											NAMING_FAILED struct {
											} `json:"."`
											FName struct {
											} `json:"f:name"`
											FValue struct {
											} `json:"f:value"`
										} `json:"k:{"name":"EXEC_MODE"}"`
										KNameLISTENPORT struct {
											NAMING_FAILED struct {
											} `json:"."`
											FName struct {
											} `json:"f:name"`
											FValue struct {
											} `json:"f:value"`
										} `json:"k:{"name":"LISTEN_PORT"}"`
										KNamePROCESSNAME struct {
											NAMING_FAILED struct {
											} `json:"."`
											FName struct {
											} `json:"f:name"`
											FValue struct {
											} `json:"f:value"`
										} `json:"k:{"name":"PROCESS_NAME"}"`
										KNameUSEENV struct {
											NAMING_FAILED struct {
											} `json:"."`
											FName struct {
											} `json:"f:name"`
											FValue struct {
											} `json:"f:value"`
										} `json:"k:{"name":"USE_ENV"}"`
									} `json:"f:env"`
									FImage struct {
									} `json:"f:image"`
									FImagePullPolicy struct {
									} `json:"f:imagePullPolicy"`
									FName struct {
									} `json:"f:name"`
									FResources struct {
										NAMING_FAILED struct {
										} `json:"."`
										FLimits struct {
											NAMING_FAILED struct {
											} `json:"."`
											FCPU struct {
											} `json:"f:cpu"`
											FMemory struct {
											} `json:"f:memory"`
										} `json:"f:limits"`
										FRequests struct {
											NAMING_FAILED struct {
											} `json:"."`
											FCPU struct {
											} `json:"f:cpu"`
											FMemory struct {
											} `json:"f:memory"`
										} `json:"f:requests"`
									} `json:"f:resources"`
									FTerminationMessagePath struct {
									} `json:"f:terminationMessagePath"`
									FTerminationMessagePolicy struct {
									} `json:"f:terminationMessagePolicy"`
								} `json:"k:{"name":"deribit-solana-oracle"}"`
							} `json:"f:containers"`
							FDNSPolicy struct {
							} `json:"f:dnsPolicy"`
							FImagePullSecrets struct {
								NAMING_FAILED struct {
								} `json:"."`
								KNameAllIcrIo struct {
								} `json:"k:{"name":"all-icr-io"}"`
							} `json:"f:imagePullSecrets"`
							FRestartPolicy struct {
							} `json:"f:restartPolicy"`
							FSchedulerName struct {
							} `json:"f:schedulerName"`
							FSecurityContext struct {
							} `json:"f:securityContext"`
							FTerminationGracePeriodSeconds struct {
							} `json:"f:terminationGracePeriodSeconds"`
						} `json:"f:spec"`
					} `json:"f:template"`
				} `json:"f:spec"`
			} `json:"fieldsV1"`
			Subresource string `json:"subresource,omitempty"`
		} `json:"managedFields"`
	} `json:"metadata"`
	Spec struct {
		Replicas int `json:"replicas"`
		Selector struct {
			MatchLabels struct {
				App      string `json:"app"`
				Heritage string `json:"heritage"`
				Release  string `json:"release"`
			} `json:"matchLabels"`
		} `json:"selector"`
		Template struct {
			Metadata struct {
				CreationTimestamp interface{} `json:"creationTimestamp"`
				Labels            struct {
					App      string `json:"app"`
					Chart    string `json:"chart"`
					Heritage string `json:"heritage"`
					Release  string `json:"release"`
				} `json:"labels"`
				Annotations struct {
					Rollme string `json:"rollme"`
				} `json:"annotations"`
			} `json:"metadata"`
			Spec struct {
				Containers []struct {
					Name  string `json:"name"`
					Image string `json:"image"`
					Env   []struct {
						Name  string `json:"name"`
						Value string `json:"value"`
					} `json:"env"`
					Resources struct {
						Limits struct {
							CPU    string `json:"cpu"`
							Memory string `json:"memory"`
						} `json:"limits"`
						Requests struct {
							CPU    string `json:"cpu"`
							Memory string `json:"memory"`
						} `json:"requests"`
					} `json:"resources"`
					TerminationMessagePath   string `json:"terminationMessagePath"`
					TerminationMessagePolicy string `json:"terminationMessagePolicy"`
					ImagePullPolicy          string `json:"imagePullPolicy"`
				} `json:"containers"`
				RestartPolicy                 string `json:"restartPolicy"`
				TerminationGracePeriodSeconds int    `json:"terminationGracePeriodSeconds"`
				DNSPolicy                     string `json:"dnsPolicy"`
				SecurityContext               struct {
				} `json:"securityContext"`
				ImagePullSecrets []struct {
					Name string `json:"name"`
				} `json:"imagePullSecrets"`
				Affinity struct {
					NodeAffinity struct {
						RequiredDuringSchedulingIgnoredDuringExecution struct {
							NodeSelectorTerms []struct {
								MatchExpressions []struct {
									Key      string   `json:"key"`
									Operator string   `json:"operator"`
									Values   []string `json:"values"`
								} `json:"matchExpressions"`
							} `json:"nodeSelectorTerms"`
						} `json:"requiredDuringSchedulingIgnoredDuringExecution"`
					} `json:"nodeAffinity"`
				} `json:"affinity"`
				SchedulerName string `json:"schedulerName"`
			} `json:"spec"`
		} `json:"template"`
		Strategy struct {
			Type          string `json:"type"`
			RollingUpdate struct {
				MaxUnavailable string `json:"maxUnavailable"`
				MaxSurge       string `json:"maxSurge"`
			} `json:"rollingUpdate"`
		} `json:"strategy"`
		RevisionHistoryLimit    int `json:"revisionHistoryLimit"`
		ProgressDeadlineSeconds int `json:"progressDeadlineSeconds"`
	} `json:"spec"`
	Status struct {
		ObservedGeneration int `json:"observedGeneration"`
		Conditions         []struct {
			Type               string    `json:"type"`
			Status             string    `json:"status"`
			LastUpdateTime     time.Time `json:"lastUpdateTime"`
			LastTransitionTime time.Time `json:"lastTransitionTime"`
			Reason             string    `json:"reason"`
			Message            string    `json:"message"`
		} `json:"conditions"`
	} `json:"status"`
}

type SymbolFeed struct {
	LiquidityThreshold string          `json:"LiquidityThreshold"`
	Methodology        string          `json:"Methodology"`
	Symbol             string          `json:"Symbol"`
	FeedSelection      []FeedSelection `json:"FeedSelection"`
}

type Exchangepair struct {
	Exchange string   `json:"Exchange"`
	Pairs    []string `json:"Pairs"`
}

type FeedSelection struct {
	Address       string         `json:"Address"`
	Blockchain    string         `json:"Blockchain"`
	Exchangepairs []Exchangepair `json:"Exchangepairs"`
}

func main() {
	var cf Configs

	log := logrus.New()
	var err error
	var datastore *models.DB
	datastore, err = models.NewDataStore()
	if err != nil {
		log.Errorln("Error connecting to asset DB: ", err)
		return
	}
	req, err := http.NewRequest("GET", "https://rest.diadata.org/monitoringNs/dia-oracles-prod", nil)
	if err != nil {
		fmt.Println(err)
	}

	user := utils.Getenv("USER", "")
	password := utils.Getenv("PASSWORD", "")

	req.SetBasicAuth(user, password)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(contents, &cf)

	for _, m := range cf {
		c := make(map[string]string)
		oc := dia.OracleConfig{}

		if m.Metadata.Namespace == "dia-oracles-prod" {
			container := m.Spec.Template.Spec.Containers[0]
			for _, env := range container.Env {

				switch env.Name {
				case "DEPLOYED_CONTRACT", "ORACLE_CONTRACTADDRESS":
					c["DEPLOYED_CONTRACT"] = env.Value
					oc.Address = env.Value

					if strings.Contains(env.Value, "0x") {
						oc.Address = common.HexToAddress(env.Value).Hex()

					}

				case "SLEEP_SECONDS":
					c["SLEEP_SECONDS"] = env.Value
					oc.SleepSeconds = env.Value

				case "DEVIATION_PERMILLE":
					c["DEVIATION_PERMILLE"] = env.Value
					oc.DeviationPermille = env.Value

				case "ASSETS":
					c["ASSETS"] = env.Value

					symbols := strings.Split(env.Value, ",")

					oc.Symbols = symbols

				case "GQL_ASSETS":
					c["GQL_ASSETS"] = env.Value

					symbols := strings.Split(env.Value, ";")

					oc.Symbols = symbols

					var sfs []SymbolFeed
					if len(symbols) > 0 {

						for _, value := range symbols {
							maybefeed := strings.Split(value, "-")
							if len(maybefeed) >= 4 {

								var sf SymbolFeed
								json.Unmarshal([]byte(maybefeed[3]), &sf)
								sf.Methodology = "mair"
								sf.Symbol = maybefeed[0] + "-" + maybefeed[1]

								sfs = append(sfs, sf)

							}

						}

					}

					fs, _ := json.Marshal(sfs)

					oc.FeederSelection = string(fs[:])

				case "FREQUENCY_SECONDS", "ORACLE_FREQUENCYSECONDS":
					c["FREQUENCY_SECONDS"] = env.Value
					oc.Frequency = env.Value
				case "ORACLE_CHAINID", "CHAIN_ID":
					c["CHAINID"] = env.Value
					oc.ChainID = env.Value
				case "ORACLE_ORACLETYPE":
					c["ORACLE_ORACLETYPE"] = env.Value

				}

				// frequence, assets, contract address, deviation, assets

			}

		}

		err := datastore.SetOracleConfigCache(oc)
		if err != nil {
			fmt.Println("erron on SetOracleConfigCache", err)
		} else {
			fmt.Println("oracleconfig updated")
		}

	}

}

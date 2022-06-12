package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	openapi "github.com/cgascoig/intersight-go-sdk/intersight"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var ()

type iksConfig struct {
	client *openapi.APIClient
}

func newCmdIKS(client *openapi.APIClient) *cobra.Command {
	log.Trace("Running iks cmd generator")
	config := iksConfig{
		client: client,
	}

	cmd := &cobra.Command{
		Use:   "iks",
		Short: "Intersight Kubernetes Service commands",
	}

	clusterCmd := &cobra.Command{
		Use:   "cluster",
		Short: "Intersight Kubernetes Service cluster commands",
	}

	createCmd := getCreateCommand(config)

	getCredentialsCmd := &cobra.Command{
		Use:   "get-credentials NAME",
		Short: "Fetch credentials for an IKS cluster",
		Run:   config.clusterGetCredentials,
		Args:  cobra.ExactArgs(1),
	}

	getCredentialsCmd.Flags().StringP("filename", "f", "-", "Filename to store the kubeconfig file")

	deployCmd := &cobra.Command{
		Use:   "deploy NAME",
		Short: "Start deployment workflow for IKS cluster",
		Run:   config.clusterDeploy,
		Args:  cobra.ExactArgs(1),
	}

	deployCmd.Flags().BoolP("no-wait", "w", false, "Do not wait for deploy workflow to complete")
	deployCmd.Flags().Int("wait-timeout", 20, "Maximum time (in minutes) to wait for workflow to complete")

	undeployCmd := &cobra.Command{
		Use:   "undeploy NAME",
		Short: "Start undeployment workflow for IKS cluster",
		Run:   config.clusterUndeploy,
		Args:  cobra.ExactArgs(1),
	}

	undeployCmd.Flags().BoolP("no-wait", "w", false, "Do not wait for undeploy workflow to complete")
	undeployCmd.Flags().Int("wait-timeout", 20, "Maximum time (in minutes) to wait for workflow to complete")

	statusCmd := &cobra.Command{
		Use:   "status NAME",
		Short: "Get cluster status",
		Run:   config.clusterStatus,
		Args:  cobra.ExactArgs(1),
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "Get clusters",
		Run:   config.clusterList,
	}

	createCmd.PreRun = config.preRun
	getCredentialsCmd.PreRun = config.preRun
	deployCmd.PreRun = config.preRun
	undeployCmd.PreRun = config.preRun
	statusCmd.PreRun = config.preRun
	listCmd.PreRun = config.preRun

	clusterCmd.AddCommand(createCmd, getCredentialsCmd, deployCmd, undeployCmd, statusCmd, listCmd)
	cmd.AddCommand(clusterCmd)

	return cmd
}

const (
	organizationFlag = "organization"

	ipPoolFlag = "ip-pool"

	networkPolicyFlag          = "network-policy"
	sysConfigPolicyFlag        = "sys-config-policy"
	versionPolicyFlag          = "version-policy"
	addonPolicyFlag            = "addon-policy"
	containerRuntimePolicyFlag = "container-runtime-policy"
	trustedCAPolicyFlag        = "trusted-ca-policy"

	vmInfraConfigFlag  = "vm-infra-config-policy"
	vmInstanceTypeFlag = "vm-instance-type-policy"

	sshUserFlag    = "ssh-user"
	sshKeyFileFlag = "ssk-key-file"
	sshKeyFlag     = "ssh-key"

	loadBalancerCountFlag = "load-balancer-count"

	numControlNodesFlag = "num-master-nodes"
	numWorkerNodesFlag  = "num-worker-nodes"
)

func getCreateCommand(config iksConfig) *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "create NAME",
		Short: "Create IKS cluster",
		Run:   config.clusterCreate,
		Args:  cobra.ExactArgs(1),
	}

	flags := createCmd.Flags()

	flags.String(organizationFlag, "default", "Name of organization to create cluster within")

	flags.String(ipPoolFlag, "", "Name of existing IP Pool to use")

	flags.String(networkPolicyFlag, "", "Name of existing Network CIDR Policy")
	flags.String(sysConfigPolicyFlag, "", "Name of existing DNS, NTP and Timezone Policy")
	flags.String(versionPolicyFlag, "", "Name of existing Kubernetes Version Policy")
	flags.String(addonPolicyFlag, "", "Name of existing Addon Policy")
	flags.String(containerRuntimePolicyFlag, "", "Name of existing Container Runtime Policy")
	flags.String(trustedCAPolicyFlag, "", "Name of existing Trusted Certificate Authority Policy")

	flags.String(vmInfraConfigFlag, "", "Name of VM Infrastructure Config Policy ")
	flags.String(vmInstanceTypeFlag, "", "Name of VM Instance Type Policy ")

	var homeDir string = ""
	homeDir, err := homedir.Dir()
	if err != nil {
		log.Fatalln(err)
	}

	flags.String(sshUserFlag, "iksadmin", "Name of user account to create on Kubernetes nodes for SSH access")
	flags.String(sshKeyFileFlag, filepath.Join(homeDir, ".ssh/id_ed25519.pub"), "Filename containing the public key for SSH authentication on Kubernetes nodes")
	flags.String(sshKeyFlag, "", "The public key for SSH authentication on Kubernetes nodes (this takes precedence over ssh-key-file if set)")

	flags.Int(loadBalancerCountFlag, 1, "The desired number of loadbalancers in the cluster")

	flags.Int(numControlNodesFlag, 1, "The desired number of master nodes in the cluster - must be 1 or 3")
	flags.Int(numWorkerNodesFlag, 1, "The desired number of worker nodes in the cluster")

	return createCmd
}

func init() {
	auxCommandsGenerators = append(auxCommandsGenerators, newCmdIKS)
}

func (config *iksConfig) clusterCreate(cmd *cobra.Command, args []string) {
	log.Debug("Executing IKS cluster create")

	flags := cmd.Flags()

	organizationName, _ := flags.GetString(organizationFlag)

	ipPoolName, _ := flags.GetString(ipPoolFlag)

	networkPolicyName, _ := flags.GetString(networkPolicyFlag)
	sysConfigPolicyName, _ := flags.GetString(sysConfigPolicyFlag)
	versionPolicyName, _ := flags.GetString(versionPolicyFlag)
	addonPolicyName, _ := flags.GetString(addonPolicyFlag)
	containerRuntimePolicyName, _ := flags.GetString(containerRuntimePolicyFlag)
	trustedCAPolicyName, _ := flags.GetString(trustedCAPolicyFlag)
	vmInfraConfigName, _ := flags.GetString(vmInfraConfigFlag)
	vmInstanceTypeName, _ := flags.GetString(vmInstanceTypeFlag)

	loadBalancerCount, _ := flags.GetInt(loadBalancerCountFlag)

	numControlNodes, _ := flags.GetInt(numControlNodesFlag)
	numWorkerNodes, _ := flags.GetInt(numWorkerNodesFlag)
	if numControlNodes != 1 && numControlNodes != 3 {
		log.Errorf("%s must be 1 or 3", numControlNodesFlag)
		return
	}
	if numWorkerNodes < 1 {
		log.Errorf("%s must be 1 or more", numWorkerNodesFlag)
	}

	sshUser, _ := flags.GetString(sshUserFlag)
	sshKeys := resolveSSHKeys(flags)

	// TODO: Validation of networkPolicyName, ..., vmInfraProviderName

	var op Operation = &CreateKubernetesClusterProfile{}

	body := map[string]interface{}{
		"Name":         args[0],
		"Organization": organizationName,

		"ClusterIpPools": []interface{}{fmt.Sprintf("MoRef[Name:%s]", ipPoolName)},

		"NetConfig": fmt.Sprintf("MoRef[Name:%s]", networkPolicyName),
		"SysConfig": fmt.Sprintf("MoRef[Name:%s]", sysConfigPolicyName),

		"ManagementConfig": map[string]interface{}{
			"LoadBalancerCount": loadBalancerCount,
			"SshUser":           sshUser,
			"SshKeys":           sshKeys,
		},
	}

	if addonPolicyName != "" {
		body["Addons"] = []interface{}{fmt.Sprintf("MoRef[Name:%s]", addonPolicyName)}
	}
	if containerRuntimePolicyName != "" {
		body["ContainerRuntimeConfig"] = fmt.Sprintf("MoRef[Name:%s]", containerRuntimePolicyName)
	}
	if trustedCAPolicyName != "" {
		body["TrustedRegistries"] = fmt.Sprintf("MoRef[Name:%s]", trustedCAPolicyName)
	}

	log.Info("Creating Kubernetes Cluster Profile")

	err := op.SetBodyParams(config.client, body)
	if err != nil {
		log.Errorf("Error preparing request to create cluster profile: %v", err)
		return
	}

	_, _, err = op.Execute(config.client, nil, nil)
	if err != nil {
		log.Errorf("Error creating cluster profile: %v", err)
		return
	}

	log.Info("Kubernetes Cluster Profile created successfully")

	log.Info("Creating Kubernetes Control Plane Node Group Profile")

	err = config.createNodeGroupProfile(args[0], true, organizationName, numControlNodes, vmInfraConfigName, ipPoolName, versionPolicyName, vmInstanceTypeName)
	if err != nil {
		log.Errorf("Control Plane Node Group: %v", err)
		return
	}

	log.Info("Kubernetes Control Plane Node Group Profile created successfully")

	log.Info("Creating Kubernetes Worker Node Group Profile")

	config.createNodeGroupProfile(args[0], false, organizationName, numWorkerNodes, vmInfraConfigName, ipPoolName, versionPolicyName, vmInstanceTypeName)
	if err != nil {
		log.Errorf("Control Plane Node Group: %v", err)
		return
	}

	log.Info("Kubernetes Worker Node Group Profile created successfully")
}

func (config *iksConfig) createNodeGroupProfile(clusterName string, controlPlane bool, organizationName string, numNodes int, vmInfraConfigName, ipPoolName, versionPolicyName, vmInstanceTypeName string) error {
	var body map[string]interface{}
	var name, nodeType string
	var op Operation
	var err error

	if controlPlane {
		name = fmt.Sprintf("%s-ctrl", clusterName)
		nodeType = "ControlPlane"
	} else {
		name = fmt.Sprintf("%s-worker", clusterName)
		nodeType = "Worker"
	}

	//Create Node Pool
	body = map[string]interface{}{
		"Name":              name,
		"NodeType":          nodeType,
		"Desiredsize":       numNodes,
		"IpPools":           []interface{}{fmt.Sprintf("MoRef[Name:%s]", ipPoolName)},
		"KubernetesVersion": fmt.Sprintf("MoRef[Name:%s]", versionPolicyName),
		"ClusterProfile":    fmt.Sprintf("MoRef[Name:%s]", clusterName),
	}
	log.Debug("Node pool body", body)

	op = &CreateKubernetesNodeGroupProfile{}
	err = op.SetBodyParams(config.client, body)
	if err != nil {
		return fmt.Errorf("Error preparing request to create Master Node Group Profile: %v", err)
	}

	_, _, err = op.Execute(config.client, nil, nil)
	if err != nil {
		return fmt.Errorf("Error creating Master Node Group Profile: %v", err)
	}

	// Create VMInfra provider
	body = map[string]interface{}{
		"Name":              name,
		"InfraConfigPolicy": fmt.Sprintf("MoRef[Name:%s]", vmInfraConfigName),
		"InstanceType":      fmt.Sprintf("MoRef[Name:%s]", vmInstanceTypeName),
		"NodeGroup":         fmt.Sprintf("MoRef[Name:%s]", name),
	}

	op = &CreateKubernetesVirtualMachineInfrastructureProvider{}
	err = op.SetBodyParams(config.client, body)
	if err != nil {
		return fmt.Errorf("Error preparing request to create VirtualMachineInfraProvider: %v", err)
	}

	_, _, err = op.Execute(config.client, nil, nil)
	if err != nil {
		return fmt.Errorf("Error creating VirtualMachineInfraProvider: %v", err)
	}

	return nil
}

func resolveSSHKeys(flags *pflag.FlagSet) []interface{} {
	sshKey, _ := flags.GetString(sshKeyFlag)
	if sshKey != "" {
		return []interface{}{sshKey}
	}

	sshKeyFilename, _ := flags.GetString(sshKeyFileFlag)
	sshKeyBytes, err := ioutil.ReadFile(sshKeyFilename)
	if err != nil {
		log.Errorf("Unable to read SSH key from %s: %v", sshKeyFilename, err)
		return []interface{}{}
	}

	return []interface{}{string(sshKeyBytes)}
}

func (config *iksConfig) getClusterProfileByName(name string) (*openapi.KubernetesClusterProfile, error) {
	var op Operation = &GetKubernetesClusterProfileList{}

	res, _, err := op.Execute(config.client, []string{}, map[string]string{"filter": fmt.Sprintf("Name eq '%s'", name)})
	if err != nil {
		return nil, fmt.Errorf("Error retreiving cluster information: %v", err)
	}

	switch result := res.(type) {
	case openapi.KubernetesClusterProfileResponse:
		if len(result.KubernetesClusterProfileList.Results) != 1 {
			return nil, fmt.Errorf("Query did not return exactly 1 cluster")
		}

		profile := result.KubernetesClusterProfileList.Results[0]

		return &profile, nil
	default:
		return nil, fmt.Errorf("Malformed response")
	}
}

func (config *iksConfig) clusterGetCredentials(cmd *cobra.Command, args []string) {
	log.Debug("Executing IKS cluster get-credentials")

	filename, err := cmd.Flags().GetString("filename")
	if err != nil {
		log.Fatal("Filename required")
	}

	var op Operation = &GetKubernetesClusterProfileList{}

	res, _, err := op.Execute(config.client, []string{}, map[string]string{"filter": fmt.Sprintf("Name eq '%s'", args[0])})
	if err != nil {
		log.Errorf("Error retreiving cluster information: %v", err)
		return
	}

	switch result := res.(type) {
	case openapi.KubernetesClusterProfileResponse:
		if len(result.KubernetesClusterProfileList.Results) != 1 {
			log.Errorf("Query did not return exactly 1 cluster")
			return
		}

		profile := result.KubernetesClusterProfileList.Results[0]
		if !profile.KubeConfig.IsSet() {
			log.Errorf("Cluster does not have Kubeconfig available")
			return
		}

		kubeconfig := *profile.KubeConfig.Get().KubeConfig

		if kubeconfig == "" {
			log.Error("Cluster has empty Kubeconfig")
			return
		}

		var out io.Writer

		if filename == "-" {
			log.Debug("Outputting cluster credentials to stdout")
			out = os.Stdout
		} else {
			log.Debugf("Outputting cluster credentials into file: %s", filename)
			out, err = os.Create(filename)
			if err != nil {
				log.Errorf("Unable to open %s for writing: %v", filename, err)
				return
			}
		}

		fmt.Fprintln(out, kubeconfig)

	default:
		log.Errorf("Malformed response")
	}
}

func (config *iksConfig) clusterDeploy(cmd *cobra.Command, args []string) {
	var err error

	log.Debug("Executing IKS cluster deploy")

	var noWait bool
	var waitTimeout int

	noWait, _ = cmd.Flags().GetBool("no-wait")
	waitTimeout, _ = cmd.Flags().GetInt("wait-timeout")

	log.Debug("Getting cluster Moid")
	profile, err := config.getClusterProfileByName(args[0])
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	moid := *profile.Moid
	log.Debugf("Found cluster Moid: %s", moid)

	var op Operation = &UpdateKubernetesClusterProfile{}
	var (
		wait          = !noWait
		maxWaitTime   = time.Duration(waitTimeout) * time.Minute // default 20 mins
		checkInterval = 30 * time.Second
	)

	body := map[string]interface{}{
		"Action": "Deploy",
	}

	err = op.SetBodyParams(config.client, body)
	if err != nil {
		log.Errorf("Error preparing request for cluster deploy action: %v", err)
		return
	}

	log.Info("Starting cluster deployment ...")

	_, _, err = op.Execute(config.client, []string{moid}, nil)
	if err != nil {
		log.Errorf("Error starting deploy for Cluster Profile: %v", err)
		return
	}

	if wait {
		log.Info("Waiting for cluster deployment ...")
		startTime := time.Now()
		endTime := startTime.Add(maxWaitTime)
		for time.Now().Before(endTime) {
			time.Sleep(checkInterval)

			profile, err := config.getClusterProfileByName(args[0])
			if err != nil {
				log.Errorf("Error checking cluster status, will continue waiting: %v", err)
				continue
			}

			status := *profile.Status
			log.Debugf("Cluster status is %s", status)

			if status == "Deploying" || status == "Deployed" || status == "NotReady" || status == "Configuring" {
				log.Infof("Still waiting for cluster deployment (current status %s) ...", status)
				continue
			} else if status == "Ready" {
				log.Info("Cluster deployment complete. ")
				return
			} else {
				log.Errorf("Cluster deployment failed: %s", status)
				return
			}

		}
		log.Warn("Timed out waiting for cluster deployment")
	}
}

func (config *iksConfig) clusterUndeploy(cmd *cobra.Command, args []string) {
	var err error

	log.Debug("Executing IKS cluster undeploy")

	var noWait bool
	var waitTimeout int

	noWait, _ = cmd.Flags().GetBool("no-wait")
	waitTimeout, _ = cmd.Flags().GetInt("wait-timeout")

	log.Debug("Getting cluster Moid")
	profile, err := config.getClusterProfileByName(args[0])
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	moid := *profile.Moid
	log.Debugf("Found cluster Moid: %s", moid)

	var op Operation = &UpdateKubernetesClusterProfile{}
	var (
		wait          = !noWait
		maxWaitTime   = time.Duration(waitTimeout) * time.Minute // default 20 mins
		checkInterval = 30 * time.Second
	)

	body := map[string]interface{}{
		"Action": "Undeploy",
	}

	err = op.SetBodyParams(config.client, body)
	if err != nil {
		log.Errorf("Error preparing request for cluster undeploy action: %v", err)
		return
	}

	log.Info("Starting cluster undeployment ...")

	_, _, err = op.Execute(config.client, []string{moid}, nil)
	if err != nil {
		log.Errorf("Error starting undeploy for Cluster Profile: %v", err)
		return
	}

	if wait {
		log.Info("Waiting for cluster undeployment ...")
		startTime := time.Now()
		endTime := startTime.Add(maxWaitTime)
		for time.Now().Before(endTime) {
			time.Sleep(checkInterval)

			profile, err := config.getClusterProfileByName(args[0])
			if err != nil {
				log.Errorf("Error checking cluster status, will continue waiting: %v", err)
				continue
			}

			status := *profile.Status
			log.Debugf("Cluster status is %s", status)

			if status == "Undeploying" || status == "Deployed" || status == "NotReady" || status == "Deleting" {
				log.Infof("Still waiting for cluster undeployment (current status %s) ...", status)
				continue
			} else if status == "Configuring" || status == "Terminated" || status == "Undeployed" {
				log.Info("Cluster undeployment complete. ")
				return
			} else {
				log.Errorf("Cluster undeployment failed: %s", status)
				return
			}

		}
		log.Warn("Timed out waiting for cluster undeployment")
	}
}

func (config *iksConfig) clusterStatus(cmd *cobra.Command, args []string) {
	log.Debug("Executing IKS cluster status")

	profile, err := config.getClusterProfileByName(args[0])
	if err != nil {
		log.Errorf("Error checking cluster status: %v", err)
		return
	}

	fmt.Printf("Cluster Name: %s\n", safeStringP(profile.Name))
	fmt.Printf("Status: %s\n", safeStringP(profile.Status))

	fmt.Printf("Node Groups:\n")
	for _, ngpRel := range profile.NodeGroups {
		// fmt.Printf("Node Group Moid: %s\n", *ngpRel.MoMoRef.Moid)
		op := &GetKubernetesNodeGroupProfileByMoid{}
		res, _, err := op.Execute(config.client, []string{*ngpRel.MoMoRef.Moid}, nil)
		if err != nil {
			log.Errorf("Error getting Node Group Profile: %v", err)
			return
		}

		switch result := res.(type) {
		case openapi.KubernetesNodeGroupProfile:
			fmt.Printf("  Node Group Name: %s\n", safeStringP(result.Name))
			fmt.Printf("  Type: %s\n", safeStringP(result.NodeType))
			fmt.Printf("  Desired Size: %d\n", *result.Desiredsize)

			fmt.Printf("  Nodes: \n")
			for _, npRel := range result.Nodes {
				if npRel.MoMoRef.ObjectType == "kubernetes.VirtualMachineNodeProfile" {
					// fmt.Printf("    Node Moid: %s\n", *npRel.MoMoRef.Moid)

					op := &GetKubernetesVirtualMachineNodeProfileByMoid{}
					res, _, err := op.Execute(config.client, []string{*npRel.MoMoRef.Moid}, nil)
					if err != nil {
						log.Errorf("Error getting Node Profile: %v", err)
						return
					}

					switch result := res.(type) {
					case openapi.KubernetesVirtualMachineNodeProfile:
						fmt.Printf("    Node Name: %s\n", safeStringP(result.Name))
						fmt.Printf("    Cloud Provider: %s\n", safeStringP(result.CloudProvider))

						op := &GetIppoolIpLeaseByMoid{}
						res, _, err := op.Execute(config.client, []string{*result.NodeIp.MoMoRef.Moid}, nil)
						if err != nil {
							log.Errorf("Error getting IpLease: %v", err)
							return
						}

						switch result := res.(type) {
						case openapi.IppoolIpLease:
							fmt.Printf("    IP Address: %s\n", safeStringP(result.IpV4Address))
						}

						fmt.Printf("    -----------------\n")
					}
				} else {
					log.Warnf("Unsupported concreate Node Profile: %s", npRel.MoMoRef.ObjectType)
				}
			}
			fmt.Println("  -----------------")
		default:
			log.Warn("Unexpected response for Node Group Profile")
		}
	}
}

func (config *iksConfig) clusterList(cmd *cobra.Command, args []string) {
	var op Operation = &GetKubernetesClusterProfileList{}

	res, _, err := op.Execute(config.client, []string{}, nil)
	if err != nil {
		log.Errorf("Error retreiving cluster information: %v", err)
		return
	}

	switch result := res.(type) {
	case openapi.KubernetesClusterProfileResponse:
		for _, profile := range result.KubernetesClusterProfileList.Results {
			fmt.Printf("%s\n", safeStringP(profile.Name))
		}
	}
}

func (config *iksConfig) preRun(cmd *cobra.Command, args []string) {
	if verbose {
		config.client.GetConfig().Debug = true
	}
}

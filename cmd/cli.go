package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cgascoig/isctl/openapi"
	"github.com/spf13/cobra"
)

func runCmd(cmd *cobra.Command, args []string) {
	log.Printf("Running command %s with args %v", cmd.Use, args)
}

// ResultHandler is the function signature to handle API results
type ResultHandler = func(result interface{}, httpResponse *http.Response, err error)

// GetCommands returns the cobra command tree for the API
func GetCommands(client *openapi.APIClient, resultHandler ResultHandler) *cobra.Command {
	rootCmd :=
		func() *cobra.Command {
			cmd := &cobra.Command{
				Use: "",

				Short: "",
				Long: `
`,
			}

			cmd.AddCommand(
				func() *cobra.Command {
					cmd := &cobra.Command{
						Use: "create",

						Short: "Create resouce(s)",
						Long: `Create resouce(s)
`,
					}

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "adapter",

								Short: "Create Adapter resource(s)",
								Long: `Create Adapter resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.AdapterApi.CreateAdapterConfigPolicy(authCtx)

											body := openapi.NewAdapterConfigPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.AdapterConfigPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'adapter.ConfigPolicy' resource.",
										Long: `Create a 'adapter.ConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateadapterconfigpolicy",

										Short: "Update a 'adapter.ConfigPolicy' resource.",
										Long: `Update a 'adapter.ConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AdapterApi.UpdateAdapterConfigPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "appliance",

								Short: "Create Appliance resource(s)",
								Long: `Create Appliance resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backup",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.CreateApplianceBackup(authCtx)

											body := openapi.NewApplianceBackup()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.ApplianceBackup(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'appliance.Backup' resource.",
										Long: `Create a 'appliance.Backup' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backuppolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.CreateApplianceBackupPolicy(authCtx)

											body := openapi.NewApplianceBackupPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.ApplianceBackupPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'appliance.BackupPolicy' resource.",
										Long: `Create a 'appliance.BackupPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "dataexportpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.CreateApplianceDataExportPolicy(authCtx)

											body := openapi.NewApplianceDataExportPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.ApplianceDataExportPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'appliance.DataExportPolicy' resource.",
										Long: `Create a 'appliance.DataExportPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deviceclaim",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.CreateApplianceDeviceClaim(authCtx)

											body := openapi.NewApplianceDeviceClaim()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.ApplianceDeviceClaim(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'appliance.DeviceClaim' resource.",
										Long: `Create a 'appliance.DeviceClaim' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "diagsetting",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.CreateApplianceDiagSetting(authCtx)

											body := openapi.NewApplianceDiagSetting()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.ApplianceDiagSetting(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'appliance.DiagSetting' resource.",
										Long: `Create a 'appliance.DiagSetting' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "restore",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.CreateApplianceRestore(authCtx)

											body := openapi.NewApplianceRestore()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.ApplianceRestore(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'appliance.Restore' resource.",
										Long: `Create a 'appliance.Restore' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateappliancebackuppolicy",

										Short: "Update a 'appliance.BackupPolicy' resource.",
										Long: `Update a 'appliance.BackupPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.UpdateApplianceBackupPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateappliancecertificatesetting",

										Short: "Update a 'appliance.CertificateSetting' resource.",
										Long: `Update a 'appliance.CertificateSetting' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.UpdateApplianceCertificateSetting(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateappliancedataexportpolicy",

										Short: "Update a 'appliance.DataExportPolicy' resource.",
										Long: `Update a 'appliance.DataExportPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.UpdateApplianceDataExportPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateappliancediagsetting",

										Short: "Update a 'appliance.DiagSetting' resource.",
										Long: `Update a 'appliance.DiagSetting' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.UpdateApplianceDiagSetting(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateappliancesetupinfo",

										Short: "Update a 'appliance.SetupInfo' resource.",
										Long: `Update a 'appliance.SetupInfo' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.UpdateApplianceSetupInfo(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateapplianceupgrade",

										Short: "Update a 'appliance.Upgrade' resource.",
										Long: `Update a 'appliance.Upgrade' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.UpdateApplianceUpgrade(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateapplianceupgradepolicy",

										Short: "Update a 'appliance.UpgradePolicy' resource.",
										Long: `Update a 'appliance.UpgradePolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.UpdateApplianceUpgradePolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "asset",

								Short: "Create Asset resource(s)",
								Long: `Create Asset resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deviceclaim",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.AssetApi.CreateAssetDeviceClaim(authCtx)

											body := openapi.NewAssetDeviceClaim()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.AssetDeviceClaim(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'asset.DeviceClaim' resource.",
										Long: `Create a 'asset.DeviceClaim' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "manageddevice",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.AssetApi.CreateAssetManagedDevice(authCtx)

											body := openapi.NewAssetManagedDevice()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.AssetManagedDevice(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'asset.ManagedDevice' resource.",
										Long: `Create a 'asset.ManagedDevice' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateassetdeviceconfiguration",

										Short: "Update a 'asset.DeviceConfiguration' resource.",
										Long: `Update a 'asset.DeviceConfiguration' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AssetApi.UpdateAssetDeviceConfiguration(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateassetdevicecontractinformation",

										Short: "Update a 'asset.DeviceContractInformation' resource.",
										Long: `Update a 'asset.DeviceContractInformation' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AssetApi.UpdateAssetDeviceContractInformation(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateassetdeviceregistration",

										Short: "Updates the resource representing the device connector. For example, this can be used to annotate the device connector resource with user-specified tags.",
										Long: `Updates the resource representing the device connector. For example, this can be used to annotate the device connector resource with user-specified tags.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AssetApi.UpdateAssetDeviceRegistration(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateassetmanageddevice",

										Short: "Update a 'asset.ManagedDevice' resource.",
										Long: `Update a 'asset.ManagedDevice' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AssetApi.UpdateAssetManagedDevice(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "bios",

								Short: "Create Bios resource(s)",
								Long: `Create Bios resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.BiosApi.CreateBiosPolicy(authCtx)

											body := openapi.NewBiosPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.BiosPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'bios.Policy' resource.",
										Long: `Create a 'bios.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatebiosbootmode",

										Short: "Update a 'bios.BootMode' resource.",
										Long: `Update a 'bios.BootMode' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.BiosApi.UpdateBiosBootMode(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatebiospolicy",

										Short: "Update a 'bios.Policy' resource.",
										Long: `Update a 'bios.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.BiosApi.UpdateBiosPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatebiosunit",

										Short: "Update a 'bios.Unit' resource.",
										Long: `Update a 'bios.Unit' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.BiosApi.UpdateBiosUnit(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "boot",

								Short: "Create Boot resource(s)",
								Long: `Create Boot resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "precisionpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.BootApi.CreateBootPrecisionPolicy(authCtx)

											body := openapi.NewBootPrecisionPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.BootPrecisionPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'boot.PrecisionPolicy' resource.",
										Long: `Create a 'boot.PrecisionPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatebootdevicebootmode",

										Short: "Update a 'boot.DeviceBootMode' resource.",
										Long: `Update a 'boot.DeviceBootMode' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.BootApi.UpdateBootDeviceBootMode(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatebootprecisionpolicy",

										Short: "Update a 'boot.PrecisionPolicy' resource.",
										Long: `Update a 'boot.PrecisionPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.BootApi.UpdateBootPrecisionPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "compute",

								Short: "Create Compute resource(s)",
								Long: `Create Compute resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatecomputeblade",

										Short: "Update a 'compute.Blade' resource.",
										Long: `Update a 'compute.Blade' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ComputeApi.UpdateComputeBlade(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatecomputeboard",

										Short: "Update a 'compute.Board' resource.",
										Long: `Update a 'compute.Board' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ComputeApi.UpdateComputeBoard(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatecomputerackunit",

										Short: "Update a 'compute.RackUnit' resource.",
										Long: `Update a 'compute.RackUnit' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ComputeApi.UpdateComputeRackUnit(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatecomputeserversetting",

										Short: "Update a 'compute.ServerSetting' resource.",
										Long: `Update a 'compute.ServerSetting' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ComputeApi.UpdateComputeServerSetting(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "deviceconnector",

								Short: "Create Deviceconnector resource(s)",
								Long: `Create Deviceconnector resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.DeviceconnectorApi.CreateDeviceconnectorPolicy(authCtx)

											body := openapi.NewDeviceconnectorPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.DeviceconnectorPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'deviceconnector.Policy' resource.",
										Long: `Create a 'deviceconnector.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatedeviceconnectorpolicy",

										Short: "Update a 'deviceconnector.Policy' resource.",
										Long: `Update a 'deviceconnector.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.DeviceconnectorApi.UpdateDeviceconnectorPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "equipment",

								Short: "Create Equipment resource(s)",
								Long: `Create Equipment resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentchassis",

										Short: "Update a 'equipment.Chassis' resource.",
										Long: `Update a 'equipment.Chassis' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.UpdateEquipmentChassis(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentfan",

										Short: "Update a 'equipment.Fan' resource.",
										Long: `Update a 'equipment.Fan' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.UpdateEquipmentFan(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentfanmodule",

										Short: "Update a 'equipment.FanModule' resource.",
										Long: `Update a 'equipment.FanModule' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.UpdateEquipmentFanModule(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentfex",

										Short: "Update a 'equipment.Fex' resource.",
										Long: `Update a 'equipment.Fex' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.UpdateEquipmentFex(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentiocard",

										Short: "Update a 'equipment.IoCard' resource.",
										Long: `Update a 'equipment.IoCard' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.UpdateEquipmentIoCard(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentioexpander",

										Short: "Update a 'equipment.IoExpander' resource.",
										Long: `Update a 'equipment.IoExpander' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.UpdateEquipmentIoExpander(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentlocatorled",

										Short: "Update a 'equipment.LocatorLed' resource.",
										Long: `Update a 'equipment.LocatorLed' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.UpdateEquipmentLocatorLed(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentpsu",

										Short: "Update a 'equipment.Psu' resource.",
										Long: `Update a 'equipment.Psu' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.UpdateEquipmentPsu(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentrackenclosure",

										Short: "Update a 'equipment.RackEnclosure' resource.",
										Long: `Update a 'equipment.RackEnclosure' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.UpdateEquipmentRackEnclosure(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentrackenclosureslot",

										Short: "Update a 'equipment.RackEnclosureSlot' resource.",
										Long: `Update a 'equipment.RackEnclosureSlot' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.UpdateEquipmentRackEnclosureSlot(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentsharediomodule",

										Short: "Update a 'equipment.SharedIoModule' resource.",
										Long: `Update a 'equipment.SharedIoModule' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.UpdateEquipmentSharedIoModule(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentswitchcard",

										Short: "Update a 'equipment.SwitchCard' resource.",
										Long: `Update a 'equipment.SwitchCard' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.UpdateEquipmentSwitchCard(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentsystemiocontroller",

										Short: "Update a 'equipment.SystemIoController' resource.",
										Long: `Update a 'equipment.SystemIoController' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.UpdateEquipmentSystemIoController(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmenttpm",

										Short: "Update a 'equipment.Tpm' resource.",
										Long: `Update a 'equipment.Tpm' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.UpdateEquipmentTpm(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ether",

								Short: "Create Ether resource(s)",
								Long: `Create Ether resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateetherphysicalport",

										Short: "Update a 'ether.PhysicalPort' resource.",
										Long: `Update a 'ether.PhysicalPort' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EtherApi.UpdateEtherPhysicalPort(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "externalsite",

								Short: "Create Externalsite resource(s)",
								Long: `Create Externalsite resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "authorization",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ExternalsiteApi.CreateExternalsiteAuthorization(authCtx)

											body := openapi.NewExternalsiteAuthorization()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.ExternalsiteAuthorization(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'externalsite.Authorization' resource.",
										Long: `Create a 'externalsite.Authorization' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateexternalsiteauthorization",

										Short: "Update a 'externalsite.Authorization' resource.",
										Long: `Update a 'externalsite.Authorization' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ExternalsiteApi.UpdateExternalsiteAuthorization(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "fault",

								Short: "Create Fault resource(s)",
								Long: `Create Fault resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatefaultinstance",

										Short: "Update a 'fault.Instance' resource.",
										Long: `Update a 'fault.Instance' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FaultApi.UpdateFaultInstance(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "fc",

								Short: "Create Fc resource(s)",
								Long: `Create Fc resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatefcphysicalport",

										Short: "Update a 'fc.PhysicalPort' resource.",
										Long: `Update a 'fc.PhysicalPort' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FcApi.UpdateFcPhysicalPort(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "feedback",

								Short: "Create Feedback resource(s)",
								Long: `Create Feedback resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "feedbackpost",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.FeedbackApi.CreateFeedbackFeedbackPost(authCtx)

											body := openapi.NewFeedbackFeedbackPost()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.FeedbackFeedbackPost(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'feedback.FeedbackPost' resource.",
										Long: `Create a 'feedback.FeedbackPost' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "firmware",

								Short: "Create Firmware resource(s)",
								Long: `Create Firmware resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "distributable",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.FirmwareApi.CreateFirmwareDistributable(authCtx)

											body := openapi.NewFirmwareDistributable()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.FirmwareDistributable(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'firmware.Distributable' resource.",
										Long: `Create a 'firmware.Distributable' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "driverdistributable",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.FirmwareApi.CreateFirmwareDriverDistributable(authCtx)

											body := openapi.NewFirmwareDriverDistributable()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.FirmwareDriverDistributable(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'firmware.DriverDistributable' resource.",
										Long: `Create a 'firmware.DriverDistributable' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "eula",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.FirmwareApi.CreateFirmwareEula(authCtx)

											body := openapi.NewFirmwareEula()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.FirmwareEula(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'firmware.Eula' resource.",
										Long: `Create a 'firmware.Eula' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "serverconfigurationutilitydistributable",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.FirmwareApi.CreateFirmwareServerConfigurationUtilityDistributable(authCtx)

											body := openapi.NewFirmwareServerConfigurationUtilityDistributable()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.FirmwareServerConfigurationUtilityDistributable(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'firmware.ServerConfigurationUtilityDistributable' resource.",
										Long: `Create a 'firmware.ServerConfigurationUtilityDistributable' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatefirmwaredistributable",

										Short: "Update a 'firmware.Distributable' resource.",
										Long: `Update a 'firmware.Distributable' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.UpdateFirmwareDistributable(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatefirmwaredriverdistributable",

										Short: "Update a 'firmware.DriverDistributable' resource.",
										Long: `Update a 'firmware.DriverDistributable' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.UpdateFirmwareDriverDistributable(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatefirmwarerunningfirmware",

										Short: "Update a 'firmware.RunningFirmware' resource.",
										Long: `Update a 'firmware.RunningFirmware' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.UpdateFirmwareRunningFirmware(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatefirmwareserverconfigurationutilitydistributable",

										Short: "Update a 'firmware.ServerConfigurationUtilityDistributable' resource.",
										Long: `Update a 'firmware.ServerConfigurationUtilityDistributable' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.UpdateFirmwareServerConfigurationUtilityDistributable(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "upgrade",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.FirmwareApi.CreateFirmwareUpgrade(authCtx)

											body := openapi.NewFirmwareUpgrade()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.FirmwareUpgrade(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'firmware.Upgrade' resource.",
										Long: `Create a 'firmware.Upgrade' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "graphics",

								Short: "Create Graphics resource(s)",
								Long: `Create Graphics resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updategraphicscard",

										Short: "Update a 'graphics.Card' resource.",
										Long: `Update a 'graphics.Card' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.GraphicsApi.UpdateGraphicsCard(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updategraphicscontroller",

										Short: "Update a 'graphics.Controller' resource.",
										Long: `Update a 'graphics.Controller' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.GraphicsApi.UpdateGraphicsController(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "hcl",

								Short: "Create Hcl resource(s)",
								Long: `Create Hcl resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "compatibilitystatus",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HclApi.CreateHclCompatibilityStatus(authCtx)

											body := openapi.NewHclCompatibilityStatus()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HclCompatibilityStatus(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hcl.CompatibilityStatus' resource.",
										Long: `Create a 'hcl.CompatibilityStatus' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hyperflexsoftwarecompatibilityinfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HclApi.CreateHclHyperflexSoftwareCompatibilityInfo(authCtx)

											body := openapi.NewHclHyperflexSoftwareCompatibilityInfo()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HclHyperflexSoftwareCompatibilityInfo(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hcl.HyperflexSoftwareCompatibilityInfo' resource.",
										Long: `Create a 'hcl.HyperflexSoftwareCompatibilityInfo' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "supporteddrivername",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HclApi.CreateHclSupportedDriverName(authCtx)

											body := openapi.NewHclSupportedDriverName()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HclSupportedDriverName(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hcl.SupportedDriverName' resource.",
										Long: `Create a 'hcl.SupportedDriverName' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehclhyperflexsoftwarecompatibilityinfo",

										Short: "Update a 'hcl.HyperflexSoftwareCompatibilityInfo' resource.",
										Long: `Update a 'hcl.HyperflexSoftwareCompatibilityInfo' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HclApi.UpdateHclHyperflexSoftwareCompatibilityInfo(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "hyperflex",

								Short: "Create Hyperflex resource(s)",
								Long: `Create Hyperflex resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "appcatalog",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexAppCatalog(authCtx)

											body := openapi.NewHyperflexAppCatalog()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexAppCatalog(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.AppCatalog' resource.",
										Long: `Create a 'hyperflex.AppCatalog' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "autosupportpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexAutoSupportPolicy(authCtx)

											body := openapi.NewHyperflexAutoSupportPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexAutoSupportPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.AutoSupportPolicy' resource.",
										Long: `Create a 'hyperflex.AutoSupportPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "capabilityinfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexCapabilityInfo(authCtx)

											body := openapi.NewHyperflexCapabilityInfo()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexCapabilityInfo(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.CapabilityInfo' resource.",
										Long: `Create a 'hyperflex.CapabilityInfo' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "clusternetworkpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexClusterNetworkPolicy(authCtx)

											body := openapi.NewHyperflexClusterNetworkPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexClusterNetworkPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.ClusterNetworkPolicy' resource.",
										Long: `Create a 'hyperflex.ClusterNetworkPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "clusterprofile",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexClusterProfile(authCtx)

											body := openapi.NewHyperflexClusterProfile()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexClusterProfile(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.ClusterProfile' resource.",
										Long: `Create a 'hyperflex.ClusterProfile' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "clusterstoragepolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexClusterStoragePolicy(authCtx)

											body := openapi.NewHyperflexClusterStoragePolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexClusterStoragePolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.ClusterStoragePolicy' resource.",
										Long: `Create a 'hyperflex.ClusterStoragePolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "extfcstoragepolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexExtFcStoragePolicy(authCtx)

											body := openapi.NewHyperflexExtFcStoragePolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexExtFcStoragePolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.ExtFcStoragePolicy' resource.",
										Long: `Create a 'hyperflex.ExtFcStoragePolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "extiscsistoragepolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexExtIscsiStoragePolicy(authCtx)

											body := openapi.NewHyperflexExtIscsiStoragePolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexExtIscsiStoragePolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.ExtIscsiStoragePolicy' resource.",
										Long: `Create a 'hyperflex.ExtIscsiStoragePolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "featurelimitexternal",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexFeatureLimitExternal(authCtx)

											body := openapi.NewHyperflexFeatureLimitExternal()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexFeatureLimitExternal(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.FeatureLimitExternal' resource.",
										Long: `Create a 'hyperflex.FeatureLimitExternal' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "featurelimitinternal",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexFeatureLimitInternal(authCtx)

											body := openapi.NewHyperflexFeatureLimitInternal()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexFeatureLimitInternal(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.FeatureLimitInternal' resource.",
										Long: `Create a 'hyperflex.FeatureLimitInternal' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hxdpversion",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexHxdpVersion(authCtx)

											body := openapi.NewHyperflexHxdpVersion()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexHxdpVersion(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.HxdpVersion' resource.",
										Long: `Create a 'hyperflex.HxdpVersion' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "localcredentialpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexLocalCredentialPolicy(authCtx)

											body := openapi.NewHyperflexLocalCredentialPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexLocalCredentialPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.LocalCredentialPolicy' resource.",
										Long: `Create a 'hyperflex.LocalCredentialPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "nodeconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexNodeConfigPolicy(authCtx)

											body := openapi.NewHyperflexNodeConfigPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexNodeConfigPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.NodeConfigPolicy' resource.",
										Long: `Create a 'hyperflex.NodeConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "nodeprofile",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexNodeProfile(authCtx)

											body := openapi.NewHyperflexNodeProfile()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexNodeProfile(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.NodeProfile' resource.",
										Long: `Create a 'hyperflex.NodeProfile' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "proxysettingpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexProxySettingPolicy(authCtx)

											body := openapi.NewHyperflexProxySettingPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexProxySettingPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.ProxySettingPolicy' resource.",
										Long: `Create a 'hyperflex.ProxySettingPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "serverfirmwareversion",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexServerFirmwareVersion(authCtx)

											body := openapi.NewHyperflexServerFirmwareVersion()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexServerFirmwareVersion(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.ServerFirmwareVersion' resource.",
										Long: `Create a 'hyperflex.ServerFirmwareVersion' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "servermodel",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexServerModel(authCtx)

											body := openapi.NewHyperflexServerModel()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexServerModel(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.ServerModel' resource.",
										Long: `Create a 'hyperflex.ServerModel' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "softwareversionpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexSoftwareVersionPolicy(authCtx)

											body := openapi.NewHyperflexSoftwareVersionPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexSoftwareVersionPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.SoftwareVersionPolicy' resource.",
										Long: `Create a 'hyperflex.SoftwareVersionPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "sysconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexSysConfigPolicy(authCtx)

											body := openapi.NewHyperflexSysConfigPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexSysConfigPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.SysConfigPolicy' resource.",
										Long: `Create a 'hyperflex.SysConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ucsmconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexUcsmConfigPolicy(authCtx)

											body := openapi.NewHyperflexUcsmConfigPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexUcsmConfigPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.UcsmConfigPolicy' resource.",
										Long: `Create a 'hyperflex.UcsmConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexappcatalog",

										Short: "Update a 'hyperflex.AppCatalog' resource.",
										Long: `Update a 'hyperflex.AppCatalog' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexAppCatalog(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexautosupportpolicy",

										Short: "Update a 'hyperflex.AutoSupportPolicy' resource.",
										Long: `Update a 'hyperflex.AutoSupportPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexAutoSupportPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexcapabilityinfo",

										Short: "Update a 'hyperflex.CapabilityInfo' resource.",
										Long: `Update a 'hyperflex.CapabilityInfo' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexCapabilityInfo(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexcluster",

										Short: "Update a 'hyperflex.Cluster' resource.",
										Long: `Update a 'hyperflex.Cluster' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexCluster(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexclusternetworkpolicy",

										Short: "Update a 'hyperflex.ClusterNetworkPolicy' resource.",
										Long: `Update a 'hyperflex.ClusterNetworkPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexClusterNetworkPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexclusterprofile",

										Short: "Update a 'hyperflex.ClusterProfile' resource.",
										Long: `Update a 'hyperflex.ClusterProfile' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexClusterProfile(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexclusterstoragepolicy",

										Short: "Update a 'hyperflex.ClusterStoragePolicy' resource.",
										Long: `Update a 'hyperflex.ClusterStoragePolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexClusterStoragePolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexextfcstoragepolicy",

										Short: "Update a 'hyperflex.ExtFcStoragePolicy' resource.",
										Long: `Update a 'hyperflex.ExtFcStoragePolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexExtFcStoragePolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexextiscsistoragepolicy",

										Short: "Update a 'hyperflex.ExtIscsiStoragePolicy' resource.",
										Long: `Update a 'hyperflex.ExtIscsiStoragePolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexExtIscsiStoragePolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexfeaturelimitexternal",

										Short: "Update a 'hyperflex.FeatureLimitExternal' resource.",
										Long: `Update a 'hyperflex.FeatureLimitExternal' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexFeatureLimitExternal(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexfeaturelimitinternal",

										Short: "Update a 'hyperflex.FeatureLimitInternal' resource.",
										Long: `Update a 'hyperflex.FeatureLimitInternal' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexFeatureLimitInternal(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexhxdpversion",

										Short: "Update a 'hyperflex.HxdpVersion' resource.",
										Long: `Update a 'hyperflex.HxdpVersion' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexHxdpVersion(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexlocalcredentialpolicy",

										Short: "Update a 'hyperflex.LocalCredentialPolicy' resource.",
										Long: `Update a 'hyperflex.LocalCredentialPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexLocalCredentialPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexnodeconfigpolicy",

										Short: "Update a 'hyperflex.NodeConfigPolicy' resource.",
										Long: `Update a 'hyperflex.NodeConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexNodeConfigPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexnodeprofile",

										Short: "Update a 'hyperflex.NodeProfile' resource.",
										Long: `Update a 'hyperflex.NodeProfile' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexNodeProfile(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexproxysettingpolicy",

										Short: "Update a 'hyperflex.ProxySettingPolicy' resource.",
										Long: `Update a 'hyperflex.ProxySettingPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexProxySettingPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexserverfirmwareversion",

										Short: "Update a 'hyperflex.ServerFirmwareVersion' resource.",
										Long: `Update a 'hyperflex.ServerFirmwareVersion' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexServerFirmwareVersion(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexservermodel",

										Short: "Update a 'hyperflex.ServerModel' resource.",
										Long: `Update a 'hyperflex.ServerModel' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexServerModel(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexsoftwareversionpolicy",

										Short: "Update a 'hyperflex.SoftwareVersionPolicy' resource.",
										Long: `Update a 'hyperflex.SoftwareVersionPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexSoftwareVersionPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexsysconfigpolicy",

										Short: "Update a 'hyperflex.SysConfigPolicy' resource.",
										Long: `Update a 'hyperflex.SysConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexSysConfigPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexucsmconfigpolicy",

										Short: "Update a 'hyperflex.UcsmConfigPolicy' resource.",
										Long: `Update a 'hyperflex.UcsmConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexUcsmConfigPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexvcenterconfigpolicy",

										Short: "Update a 'hyperflex.VcenterConfigPolicy' resource.",
										Long: `Update a 'hyperflex.VcenterConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.UpdateHyperflexVcenterConfigPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vcenterconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.CreateHyperflexVcenterConfigPolicy(authCtx)

											body := openapi.NewHyperflexVcenterConfigPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.HyperflexVcenterConfigPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.VcenterConfigPolicy' resource.",
										Long: `Create a 'hyperflex.VcenterConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "iaas",

								Short: "Create Iaas resource(s)",
								Long: `Create Iaas resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiaasucsdinfo",

										Short: "Update a 'iaas.UcsdInfo' resource.",
										Long: `Update a 'iaas.UcsdInfo' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IaasApi.UpdateIaasUcsdInfo(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "iam",

								Short: "Create Iam resource(s)",
								Long: `Create Iam resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "account",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamAccount(authCtx)

											body := openapi.NewIamAccount()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamAccount(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.Account' resource.",
										Long: `Create a 'iam.Account' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "apikey",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamApiKey(authCtx)

											body := openapi.NewIamApiKey()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamApiKey(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.ApiKey' resource.",
										Long: `Create a 'iam.ApiKey' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "appregistration",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamAppRegistration(authCtx)

											body := openapi.NewIamAppRegistration()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamAppRegistration(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.AppRegistration' resource.",
										Long: `Create a 'iam.AppRegistration' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "certificate",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamCertificate(authCtx)

											body := openapi.NewIamCertificate()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamCertificate(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.Certificate' resource.",
										Long: `Create a 'iam.Certificate' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "certificaterequest",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamCertificateRequest(authCtx)

											body := openapi.NewIamCertificateRequest()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamCertificateRequest(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.CertificateRequest' resource.",
										Long: `Create a 'iam.CertificateRequest' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "endpointuser",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamEndPointUser(authCtx)

											body := openapi.NewIamEndPointUser()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamEndPointUser(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.EndPointUser' resource.",
										Long: `Create a 'iam.EndPointUser' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "endpointuserpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamEndPointUserPolicy(authCtx)

											body := openapi.NewIamEndPointUserPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamEndPointUserPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.EndPointUserPolicy' resource.",
										Long: `Create a 'iam.EndPointUserPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "endpointuserrole",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamEndPointUserRole(authCtx)

											body := openapi.NewIamEndPointUserRole()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamEndPointUserRole(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.EndPointUserRole' resource.",
										Long: `Create a 'iam.EndPointUserRole' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "idp",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamIdp(authCtx)

											body := openapi.NewIamIdp()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamIdp(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.Idp' resource.",
										Long: `Create a 'iam.Idp' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ldapgroup",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamLdapGroup(authCtx)

											body := openapi.NewIamLdapGroup()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamLdapGroup(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.LdapGroup' resource.",
										Long: `Create a 'iam.LdapGroup' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ldappolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamLdapPolicy(authCtx)

											body := openapi.NewIamLdapPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamLdapPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.LdapPolicy' resource.",
										Long: `Create a 'iam.LdapPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ldapprovider",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamLdapProvider(authCtx)

											body := openapi.NewIamLdapProvider()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamLdapProvider(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.LdapProvider' resource.",
										Long: `Create a 'iam.LdapProvider' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "permission",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamPermission(authCtx)

											body := openapi.NewIamPermission()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamPermission(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.Permission' resource.",
										Long: `Create a 'iam.Permission' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "privatekeyspec",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamPrivateKeySpec(authCtx)

											body := openapi.NewIamPrivateKeySpec()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamPrivateKeySpec(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.PrivateKeySpec' resource.",
										Long: `Create a 'iam.PrivateKeySpec' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "qualifier",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamQualifier(authCtx)

											body := openapi.NewIamQualifier()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamQualifier(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.Qualifier' resource.",
										Long: `Create a 'iam.Qualifier' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "resourceroles",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamResourceRoles(authCtx)

											body := openapi.NewIamResourceRoles()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamResourceRoles(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.ResourceRoles' resource.",
										Long: `Create a 'iam.ResourceRoles' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "sessionlimits",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamSessionLimits(authCtx)

											body := openapi.NewIamSessionLimits()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamSessionLimits(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.SessionLimits' resource.",
										Long: `Create a 'iam.SessionLimits' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "trustpoint",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamTrustPoint(authCtx)

											body := openapi.NewIamTrustPoint()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamTrustPoint(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.TrustPoint' resource.",
										Long: `Create a 'iam.TrustPoint' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamaccount",

										Short: "Update a 'iam.Account' resource.",
										Long: `Update a 'iam.Account' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamAccount(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamapikey",

										Short: "Update a 'iam.ApiKey' resource.",
										Long: `Update a 'iam.ApiKey' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamApiKey(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamappregistration",

										Short: "Update a 'iam.AppRegistration' resource.",
										Long: `Update a 'iam.AppRegistration' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamAppRegistration(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamcertificate",

										Short: "Update a 'iam.Certificate' resource.",
										Long: `Update a 'iam.Certificate' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamCertificate(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamcertificaterequest",

										Short: "Update a 'iam.CertificateRequest' resource.",
										Long: `Update a 'iam.CertificateRequest' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamCertificateRequest(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamendpointuser",

										Short: "Update a 'iam.EndPointUser' resource.",
										Long: `Update a 'iam.EndPointUser' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamEndPointUser(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamendpointuserpolicy",

										Short: "Update a 'iam.EndPointUserPolicy' resource.",
										Long: `Update a 'iam.EndPointUserPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamEndPointUserPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamendpointuserrole",

										Short: "Update a 'iam.EndPointUserRole' resource.",
										Long: `Update a 'iam.EndPointUserRole' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamEndPointUserRole(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamidp",

										Short: "Update a 'iam.Idp' resource.",
										Long: `Update a 'iam.Idp' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamIdp(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamidpreference",

										Short: "Update a 'iam.IdpReference' resource.",
										Long: `Update a 'iam.IdpReference' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamIdpReference(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamldapgroup",

										Short: "Update a 'iam.LdapGroup' resource.",
										Long: `Update a 'iam.LdapGroup' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamLdapGroup(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamldappolicy",

										Short: "Update a 'iam.LdapPolicy' resource.",
										Long: `Update a 'iam.LdapPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamLdapPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamldapprovider",

										Short: "Update a 'iam.LdapProvider' resource.",
										Long: `Update a 'iam.LdapProvider' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamLdapProvider(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamlocaluserpassword",

										Short: "Update a 'iam.LocalUserPassword' resource.",
										Long: `Update a 'iam.LocalUserPassword' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamLocalUserPassword(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiampermission",

										Short: "Update a 'iam.Permission' resource.",
										Long: `Update a 'iam.Permission' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamPermission(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamprivatekeyspec",

										Short: "Update a 'iam.PrivateKeySpec' resource.",
										Long: `Update a 'iam.PrivateKeySpec' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamPrivateKeySpec(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamqualifier",

										Short: "Update a 'iam.Qualifier' resource.",
										Long: `Update a 'iam.Qualifier' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamQualifier(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamresourceroles",

										Short: "Update a 'iam.ResourceRoles' resource.",
										Long: `Update a 'iam.ResourceRoles' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamResourceRoles(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamsessionlimits",

										Short: "Update a 'iam.SessionLimits' resource.",
										Long: `Update a 'iam.SessionLimits' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamSessionLimits(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamuser",

										Short: "Update a 'iam.User' resource.",
										Long: `Update a 'iam.User' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamUser(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamusergroup",

										Short: "Update a 'iam.UserGroup' resource.",
										Long: `Update a 'iam.UserGroup' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamUserGroup(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamuserpreference",

										Short: "Update a 'iam.UserPreference' resource.",
										Long: `Update a 'iam.UserPreference' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.UpdateIamUserPreference(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "user",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamUser(authCtx)

											body := openapi.NewIamUser()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamUser(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.User' resource.",
										Long: `Create a 'iam.User' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "usergroup",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.CreateIamUserGroup(authCtx)

											body := openapi.NewIamUserGroup()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IamUserGroup(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.UserGroup' resource.",
										Long: `Create a 'iam.UserGroup' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "infra",

								Short: "Create Infra resource(s)",
								Long: `Create Infra resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "accountexperience",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.InfraApi.CreateInfraAccountExperience(authCtx)

											body := openapi.NewInfraAccountExperience()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.InfraAccountExperience(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'infra.AccountExperience' resource.",
										Long: `Create a 'infra.AccountExperience' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateinfraaccountexperience",

										Short: "Update a 'infra.AccountExperience' resource.",
										Long: `Update a 'infra.AccountExperience' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.InfraApi.UpdateInfraAccountExperience(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "inventory",

								Short: "Create Inventory resource(s)",
								Long: `Create Inventory resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "request",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.InventoryApi.CreateInventoryRequest(authCtx)

											body := openapi.NewInventoryRequest()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.InventoryRequest(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'inventory.Request' resource.",
										Long: `Create a 'inventory.Request' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateinventorygenericinventory",

										Short: "Update a 'inventory.GenericInventory' resource.",
										Long: `Update a 'inventory.GenericInventory' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.InventoryApi.UpdateInventoryGenericInventory(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateinventorygenericinventoryholder",

										Short: "Update a 'inventory.GenericInventoryHolder' resource.",
										Long: `Update a 'inventory.GenericInventoryHolder' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.InventoryApi.UpdateInventoryGenericInventoryHolder(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ipmioverlan",

								Short: "Create Ipmioverlan resource(s)",
								Long: `Create Ipmioverlan resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IpmioverlanApi.CreateIpmioverlanPolicy(authCtx)

											body := openapi.NewIpmioverlanPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.IpmioverlanPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'ipmioverlan.Policy' resource.",
										Long: `Create a 'ipmioverlan.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateipmioverlanpolicy",

										Short: "Update a 'ipmioverlan.Policy' resource.",
										Long: `Update a 'ipmioverlan.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IpmioverlanApi.UpdateIpmioverlanPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "kvm",

								Short: "Create Kvm resource(s)",
								Long: `Create Kvm resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.KvmApi.CreateKvmPolicy(authCtx)

											body := openapi.NewKvmPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.KvmPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'kvm.Policy' resource.",
										Long: `Create a 'kvm.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatekvmpolicy",

										Short: "Update a 'kvm.Policy' resource.",
										Long: `Update a 'kvm.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.KvmApi.UpdateKvmPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "license",

								Short: "Create License resource(s)",
								Long: `Create License resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "licenseinfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.LicenseApi.CreateLicenseLicenseInfo(authCtx)

											body := openapi.NewLicenseLicenseInfo()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.LicenseLicenseInfo(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'license.LicenseInfo' resource.",
										Long: `Create a 'license.LicenseInfo' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatelicenseaccountlicensedata",

										Short: "Update a 'license.AccountLicenseData' resource.",
										Long: `Update a 'license.AccountLicenseData' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.LicenseApi.UpdateLicenseAccountLicenseData(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatelicensecustomerop",

										Short: "Update a 'license.CustomerOp' resource.",
										Long: `Update a 'license.CustomerOp' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.LicenseApi.UpdateLicenseCustomerOp(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatelicenselicenseinfo",

										Short: "Update a 'license.LicenseInfo' resource.",
										Long: `Update a 'license.LicenseInfo' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.LicenseApi.UpdateLicenseLicenseInfo(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatelicensesmartlicensetoken",

										Short: "Update a 'license.SmartlicenseToken' resource.",
										Long: `Update a 'license.SmartlicenseToken' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.LicenseApi.UpdateLicenseSmartlicenseToken(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ls",

								Short: "Create Ls resource(s)",
								Long: `Create Ls resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatelsserviceprofile",

										Short: "Update a 'ls.ServiceProfile' resource.",
										Long: `Update a 'ls.ServiceProfile' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.LsApi.UpdateLsServiceProfile(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "management",

								Short: "Create Management resource(s)",
								Long: `Create Management resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatemanagementcontroller",

										Short: "Update a 'management.Controller' resource.",
										Long: `Update a 'management.Controller' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ManagementApi.UpdateManagementController(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatemanagemententity",

										Short: "Update a 'management.Entity' resource.",
										Long: `Update a 'management.Entity' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ManagementApi.UpdateManagementEntity(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatemanagementinterface",

										Short: "Update a 'management.Interface' resource.",
										Long: `Update a 'management.Interface' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ManagementApi.UpdateManagementInterface(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "memory",

								Short: "Create Memory resource(s)",
								Long: `Create Memory resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "persistentmemorypolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.MemoryApi.CreateMemoryPersistentMemoryPolicy(authCtx)

											body := openapi.NewMemoryPersistentMemoryPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.MemoryPersistentMemoryPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'memory.PersistentMemoryPolicy' resource.",
										Long: `Create a 'memory.PersistentMemoryPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememoryarray",

										Short: "Update a 'memory.Array' resource.",
										Long: `Update a 'memory.Array' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.UpdateMemoryArray(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememorypersistentmemoryconfigresult",

										Short: "Update a 'memory.PersistentMemoryConfigResult' resource.",
										Long: `Update a 'memory.PersistentMemoryConfigResult' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.UpdateMemoryPersistentMemoryConfigResult(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememorypersistentmemoryconfiguration",

										Short: "Update a 'memory.PersistentMemoryConfiguration' resource.",
										Long: `Update a 'memory.PersistentMemoryConfiguration' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.UpdateMemoryPersistentMemoryConfiguration(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememorypersistentmemorynamespace",

										Short: "Update a 'memory.PersistentMemoryNamespace' resource.",
										Long: `Update a 'memory.PersistentMemoryNamespace' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.UpdateMemoryPersistentMemoryNamespace(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememorypersistentmemorynamespaceconfigresult",

										Short: "Update a 'memory.PersistentMemoryNamespaceConfigResult' resource.",
										Long: `Update a 'memory.PersistentMemoryNamespaceConfigResult' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.UpdateMemoryPersistentMemoryNamespaceConfigResult(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememorypersistentmemorypolicy",

										Short: "Update a 'memory.PersistentMemoryPolicy' resource.",
										Long: `Update a 'memory.PersistentMemoryPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.UpdateMemoryPersistentMemoryPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememorypersistentmemoryregion",

										Short: "Update a 'memory.PersistentMemoryRegion' resource.",
										Long: `Update a 'memory.PersistentMemoryRegion' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.UpdateMemoryPersistentMemoryRegion(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememorypersistentmemoryunit",

										Short: "Update a 'memory.PersistentMemoryUnit' resource.",
										Long: `Update a 'memory.PersistentMemoryUnit' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.UpdateMemoryPersistentMemoryUnit(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememoryunit",

										Short: "Update a 'memory.Unit' resource.",
										Long: `Update a 'memory.Unit' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.UpdateMemoryUnit(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "network",

								Short: "Create Network resource(s)",
								Long: `Create Network resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatenetworkelement",

										Short: "Update a 'network.Element' resource.",
										Long: `Update a 'network.Element' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NetworkApi.UpdateNetworkElement(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "networkconfig",

								Short: "Create Networkconfig resource(s)",
								Long: `Create Networkconfig resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NetworkconfigApi.CreateNetworkconfigPolicy(authCtx)

											body := openapi.NewNetworkconfigPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.NetworkconfigPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'networkconfig.Policy' resource.",
										Long: `Create a 'networkconfig.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatenetworkconfigpolicy",

										Short: "Update a 'networkconfig.Policy' resource.",
										Long: `Update a 'networkconfig.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NetworkconfigApi.UpdateNetworkconfigPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ntp",

								Short: "Create Ntp resource(s)",
								Long: `Create Ntp resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NtpApi.CreateNtpPolicy(authCtx)

											body := openapi.NewNtpPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.NtpPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'ntp.Policy' resource.",
										Long: `Create a 'ntp.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatentppolicy",

										Short: "Update a 'ntp.Policy' resource.",
										Long: `Update a 'ntp.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NtpApi.UpdateNtpPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "organization",

								Short: "Create Organization resource(s)",
								Long: `Create Organization resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "organization",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.OrganizationApi.CreateOrganizationOrganization(authCtx)

											body := openapi.NewOrganizationOrganization()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.OrganizationOrganization(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'organization.Organization' resource.",
										Long: `Create a 'organization.Organization' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateorganizationorganization",

										Short: "Update a 'organization.Organization' resource.",
										Long: `Update a 'organization.Organization' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.OrganizationApi.UpdateOrganizationOrganization(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "os",

								Short: "Create Os resource(s)",
								Long: `Create Os resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configurationfile",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.OsApi.CreateOsConfigurationFile(authCtx)

											body := openapi.NewOsConfigurationFile()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.OsConfigurationFile(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'os.ConfigurationFile' resource.",
										Long: `Create a 'os.ConfigurationFile' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "install",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.OsApi.CreateOsInstall(authCtx)

											body := openapi.NewOsInstall()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.OsInstall(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'os.Install' resource.",
										Long: `Create a 'os.Install' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ossupport",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.OsApi.CreateOsOsSupport(authCtx)

											body := openapi.NewOsOsSupport()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.OsOsSupport(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'os.OsSupport' resource.",
										Long: `Create a 'os.OsSupport' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "templatefile",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.OsApi.CreateOsTemplateFile(authCtx)

											body := openapi.NewOsTemplateFile()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.OsTemplateFile(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'os.TemplateFile' resource.",
										Long: `Create a 'os.TemplateFile' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "pci",

								Short: "Create Pci resource(s)",
								Long: `Create Pci resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatepcidevice",

										Short: "Update a 'pci.Device' resource.",
										Long: `Update a 'pci.Device' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.PciApi.UpdatePciDevice(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatepcilink",

										Short: "Update a 'pci.Link' resource.",
										Long: `Update a 'pci.Link' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.PciApi.UpdatePciLink(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatepciswitch",

										Short: "Update a 'pci.Switch' resource.",
										Long: `Update a 'pci.Switch' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.PciApi.UpdatePciSwitch(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "port",

								Short: "Create Port resource(s)",
								Long: `Create Port resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateportgroup",

										Short: "Update a 'port.Group' resource.",
										Long: `Update a 'port.Group' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.PortApi.UpdatePortGroup(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateportsubgroup",

										Short: "Update a 'port.SubGroup' resource.",
										Long: `Update a 'port.SubGroup' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.PortApi.UpdatePortSubGroup(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "processor",

								Short: "Create Processor resource(s)",
								Long: `Create Processor resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateprocessorunit",

										Short: "Update a 'processor.Unit' resource.",
										Long: `Update a 'processor.Unit' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ProcessorApi.UpdateProcessorUnit(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "recovery",

								Short: "Create Recovery resource(s)",
								Long: `Create Recovery resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backupconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.RecoveryApi.CreateRecoveryBackupConfigPolicy(authCtx)

											body := openapi.NewRecoveryBackupConfigPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.RecoveryBackupConfigPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'recovery.BackupConfigPolicy' resource.",
										Long: `Create a 'recovery.BackupConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backupprofile",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.RecoveryApi.CreateRecoveryBackupProfile(authCtx)

											body := openapi.NewRecoveryBackupProfile()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.RecoveryBackupProfile(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'recovery.BackupProfile' resource.",
										Long: `Create a 'recovery.BackupProfile' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ondemandbackup",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.RecoveryApi.CreateRecoveryOnDemandBackup(authCtx)

											body := openapi.NewRecoveryOnDemandBackup()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.RecoveryOnDemandBackup(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'recovery.OnDemandBackup' resource.",
										Long: `Create a 'recovery.OnDemandBackup' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "restore",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.RecoveryApi.CreateRecoveryRestore(authCtx)

											body := openapi.NewRecoveryRestore()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.RecoveryRestore(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'recovery.Restore' resource.",
										Long: `Create a 'recovery.Restore' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "scheduleconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.RecoveryApi.CreateRecoveryScheduleConfigPolicy(authCtx)

											body := openapi.NewRecoveryScheduleConfigPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.RecoveryScheduleConfigPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'recovery.ScheduleConfigPolicy' resource.",
										Long: `Create a 'recovery.ScheduleConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updaterecoverybackupconfigpolicy",

										Short: "Update a 'recovery.BackupConfigPolicy' resource.",
										Long: `Update a 'recovery.BackupConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.UpdateRecoveryBackupConfigPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updaterecoverybackupprofile",

										Short: "Update a 'recovery.BackupProfile' resource.",
										Long: `Update a 'recovery.BackupProfile' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.UpdateRecoveryBackupProfile(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updaterecoveryondemandbackup",

										Short: "Update a 'recovery.OnDemandBackup' resource.",
										Long: `Update a 'recovery.OnDemandBackup' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.UpdateRecoveryOnDemandBackup(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updaterecoveryscheduleconfigpolicy",

										Short: "Update a 'recovery.ScheduleConfigPolicy' resource.",
										Long: `Update a 'recovery.ScheduleConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.UpdateRecoveryScheduleConfigPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "resource",

								Short: "Create Resource resource(s)",
								Long: `Create Resource resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "group",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ResourceApi.CreateResourceGroup(authCtx)

											body := openapi.NewResourceGroup()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.ResourceGroup(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'resource.Group' resource.",
										Long: `Create a 'resource.Group' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateresourcegroup",

										Short: "Update a 'resource.Group' resource.",
										Long: `Update a 'resource.Group' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ResourceApi.UpdateResourceGroup(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "sdcard",

								Short: "Create Sdcard resource(s)",
								Long: `Create Sdcard resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SdcardApi.CreateSdcardPolicy(authCtx)

											body := openapi.NewSdcardPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.SdcardPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'sdcard.Policy' resource.",
										Long: `Create a 'sdcard.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesdcardpolicy",

										Short: "Update a 'sdcard.Policy' resource.",
										Long: `Update a 'sdcard.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdcardApi.UpdateSdcardPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "sdwan",

								Short: "Create Sdwan resource(s)",
								Long: `Create Sdwan resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "profile",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SdwanApi.CreateSdwanProfile(authCtx)

											body := openapi.NewSdwanProfile()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.SdwanProfile(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'sdwan.Profile' resource.",
										Long: `Create a 'sdwan.Profile' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "routernode",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SdwanApi.CreateSdwanRouterNode(authCtx)

											body := openapi.NewSdwanRouterNode()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.SdwanRouterNode(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'sdwan.RouterNode' resource.",
										Long: `Create a 'sdwan.RouterNode' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "routerpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SdwanApi.CreateSdwanRouterPolicy(authCtx)

											body := openapi.NewSdwanRouterPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.SdwanRouterPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'sdwan.RouterPolicy' resource.",
										Long: `Create a 'sdwan.RouterPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesdwanprofile",

										Short: "Update a 'sdwan.Profile' resource.",
										Long: `Update a 'sdwan.Profile' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdwanApi.UpdateSdwanProfile(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesdwanrouternode",

										Short: "Update a 'sdwan.RouterNode' resource.",
										Long: `Update a 'sdwan.RouterNode' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdwanApi.UpdateSdwanRouterNode(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesdwanrouterpolicy",

										Short: "Update a 'sdwan.RouterPolicy' resource.",
										Long: `Update a 'sdwan.RouterPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdwanApi.UpdateSdwanRouterPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesdwanvmanageaccountpolicy",

										Short: "Update a 'sdwan.VmanageAccountPolicy' resource.",
										Long: `Update a 'sdwan.VmanageAccountPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdwanApi.UpdateSdwanVmanageAccountPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vmanageaccountpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SdwanApi.CreateSdwanVmanageAccountPolicy(authCtx)

											body := openapi.NewSdwanVmanageAccountPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.SdwanVmanageAccountPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'sdwan.VmanageAccountPolicy' resource.",
										Long: `Create a 'sdwan.VmanageAccountPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "search",

								Short: "Create Search resource(s)",
								Long: `Create Search resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "suggestitem",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SearchApi.CreateSearchSuggestItem(authCtx)

											body := openapi.NewSearchSuggestItem()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.SearchSuggestItem(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'search.SuggestItem' resource.",
										Long: `Create a 'search.SuggestItem' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "security",

								Short: "Create Security resource(s)",
								Long: `Create Security resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesecurityunit",

										Short: "Update a 'security.Unit' resource.",
										Long: `Update a 'security.Unit' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SecurityApi.UpdateSecurityUnit(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "server",

								Short: "Create Server resource(s)",
								Long: `Create Server resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configimport",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ServerApi.CreateServerConfigImport(authCtx)

											body := openapi.NewServerConfigImport()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.ServerConfigImport(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'server.ConfigImport' resource.",
										Long: `Create a 'server.ConfigImport' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "profile",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ServerApi.CreateServerProfile(authCtx)

											body := openapi.NewServerProfile()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.ServerProfile(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'server.Profile' resource.",
										Long: `Create a 'server.Profile' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateserverprofile",

										Short: "Update a 'server.Profile' resource.",
										Long: `Update a 'server.Profile' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ServerApi.UpdateServerProfile(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "smtp",

								Short: "Create Smtp resource(s)",
								Long: `Create Smtp resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SmtpApi.CreateSmtpPolicy(authCtx)

											body := openapi.NewSmtpPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.SmtpPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'smtp.Policy' resource.",
										Long: `Create a 'smtp.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesmtppolicy",

										Short: "Update a 'smtp.Policy' resource.",
										Long: `Update a 'smtp.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SmtpApi.UpdateSmtpPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "snmp",

								Short: "Create Snmp resource(s)",
								Long: `Create Snmp resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SnmpApi.CreateSnmpPolicy(authCtx)

											body := openapi.NewSnmpPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.SnmpPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'snmp.Policy' resource.",
										Long: `Create a 'snmp.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesnmppolicy",

										Short: "Update a 'snmp.Policy' resource.",
										Long: `Update a 'snmp.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SnmpApi.UpdateSnmpPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "software",

								Short: "Create Software resource(s)",
								Long: `Create Software resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hclmeta",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SoftwareApi.CreateSoftwareHclMeta(authCtx)

											body := openapi.NewSoftwareHclMeta()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.SoftwareHclMeta(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'software.HclMeta' resource.",
										Long: `Create a 'software.HclMeta' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hyperflexdistributable",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SoftwareApi.CreateSoftwareHyperflexDistributable(authCtx)

											body := openapi.NewSoftwareHyperflexDistributable()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.SoftwareHyperflexDistributable(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'software.HyperflexDistributable' resource.",
										Long: `Create a 'software.HyperflexDistributable' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "solutiondistributable",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SoftwareApi.CreateSoftwareSolutionDistributable(authCtx)

											body := openapi.NewSoftwareSolutionDistributable()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.SoftwareSolutionDistributable(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'software.SolutionDistributable' resource.",
										Long: `Create a 'software.SolutionDistributable' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesoftwarehclmeta",

										Short: "Update a 'software.HclMeta' resource.",
										Long: `Update a 'software.HclMeta' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwareApi.UpdateSoftwareHclMeta(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesoftwarehyperflexdistributable",

										Short: "Update a 'software.HyperflexDistributable' resource.",
										Long: `Update a 'software.HyperflexDistributable' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwareApi.UpdateSoftwareHyperflexDistributable(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesoftwaresolutiondistributable",

										Short: "Update a 'software.SolutionDistributable' resource.",
										Long: `Update a 'software.SolutionDistributable' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwareApi.UpdateSoftwareSolutionDistributable(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "softwarerepository",

								Short: "Create Softwarerepository resource(s)",
								Long: `Create Softwarerepository resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "authorization",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SoftwarerepositoryApi.CreateSoftwarerepositoryAuthorization(authCtx)

											body := openapi.NewSoftwarerepositoryAuthorization()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.SoftwarerepositoryAuthorization(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'softwarerepository.Authorization' resource.",
										Long: `Create a 'softwarerepository.Authorization' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "operatingsystemfile",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SoftwarerepositoryApi.CreateSoftwarerepositoryOperatingSystemFile(authCtx)

											body := openapi.NewSoftwarerepositoryOperatingSystemFile()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.SoftwarerepositoryOperatingSystemFile(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'softwarerepository.OperatingSystemFile' resource.",
										Long: `Create a 'softwarerepository.OperatingSystemFile' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesoftwarerepositoryauthorization",

										Short: "Update a 'softwarerepository.Authorization' resource.",
										Long: `Update a 'softwarerepository.Authorization' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwarerepositoryApi.UpdateSoftwarerepositoryAuthorization(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesoftwarerepositoryoperatingsystemfile",

										Short: "Update a 'softwarerepository.OperatingSystemFile' resource.",
										Long: `Update a 'softwarerepository.OperatingSystemFile' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwarerepositoryApi.UpdateSoftwarerepositoryOperatingSystemFile(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "sol",

								Short: "Create Sol resource(s)",
								Long: `Create Sol resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SolApi.CreateSolPolicy(authCtx)

											body := openapi.NewSolPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.SolPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'sol.Policy' resource.",
										Long: `Create a 'sol.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesolpolicy",

										Short: "Update a 'sol.Policy' resource.",
										Long: `Update a 'sol.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SolApi.UpdateSolPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ssh",

								Short: "Create Ssh resource(s)",
								Long: `Create Ssh resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SshApi.CreateSshPolicy(authCtx)

											body := openapi.NewSshPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.SshPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'ssh.Policy' resource.",
										Long: `Create a 'ssh.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesshpolicy",

										Short: "Update a 'ssh.Policy' resource.",
										Long: `Update a 'ssh.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SshApi.UpdateSshPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "storage",

								Short: "Create Storage resource(s)",
								Long: `Create Storage resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "diskgrouppolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.CreateStorageDiskGroupPolicy(authCtx)

											body := openapi.NewStorageDiskGroupPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.StorageDiskGroupPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'storage.DiskGroupPolicy' resource.",
										Long: `Create a 'storage.DiskGroupPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "storagepolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.CreateStorageStoragePolicy(authCtx)

											body := openapi.NewStorageStoragePolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.StorageStoragePolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'storage.StoragePolicy' resource.",
										Long: `Create a 'storage.StoragePolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragecontroller",

										Short: "Update a 'storage.Controller' resource.",
										Long: `Update a 'storage.Controller' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStorageController(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragediskgrouppolicy",

										Short: "Update a 'storage.DiskGroupPolicy' resource.",
										Long: `Update a 'storage.DiskGroupPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStorageDiskGroupPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageenclosure",

										Short: "Update a 'storage.Enclosure' resource.",
										Long: `Update a 'storage.Enclosure' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStorageEnclosure(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageenclosuredisk",

										Short: "Update a 'storage.EnclosureDisk' resource.",
										Long: `Update a 'storage.EnclosureDisk' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStorageEnclosureDisk(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageenclosurediskslotep",

										Short: "Update a 'storage.EnclosureDiskSlotEp' resource.",
										Long: `Update a 'storage.EnclosureDiskSlotEp' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStorageEnclosureDiskSlotEp(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageflexflashcontroller",

										Short: "Update a 'storage.FlexFlashController' resource.",
										Long: `Update a 'storage.FlexFlashController' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStorageFlexFlashController(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageflexflashcontrollerprops",

										Short: "Update a 'storage.FlexFlashControllerProps' resource.",
										Long: `Update a 'storage.FlexFlashControllerProps' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStorageFlexFlashControllerProps(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageflexflashphysicaldrive",

										Short: "Update a 'storage.FlexFlashPhysicalDrive' resource.",
										Long: `Update a 'storage.FlexFlashPhysicalDrive' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStorageFlexFlashPhysicalDrive(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageflexflashvirtualdrive",

										Short: "Update a 'storage.FlexFlashVirtualDrive' resource.",
										Long: `Update a 'storage.FlexFlashVirtualDrive' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStorageFlexFlashVirtualDrive(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageflexutilcontroller",

										Short: "Update a 'storage.FlexUtilController' resource.",
										Long: `Update a 'storage.FlexUtilController' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStorageFlexUtilController(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageflexutilphysicaldrive",

										Short: "Update a 'storage.FlexUtilPhysicalDrive' resource.",
										Long: `Update a 'storage.FlexUtilPhysicalDrive' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStorageFlexUtilPhysicalDrive(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageflexutilvirtualdrive",

										Short: "Update a 'storage.FlexUtilVirtualDrive' resource.",
										Long: `Update a 'storage.FlexUtilVirtualDrive' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStorageFlexUtilVirtualDrive(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragephysicaldisk",

										Short: "Update a 'storage.PhysicalDisk' resource.",
										Long: `Update a 'storage.PhysicalDisk' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStoragePhysicalDisk(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragephysicaldiskextension",

										Short: "Update a 'storage.PhysicalDiskExtension' resource.",
										Long: `Update a 'storage.PhysicalDiskExtension' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStoragePhysicalDiskExtension(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragephysicaldiskusage",

										Short: "Update a 'storage.PhysicalDiskUsage' resource.",
										Long: `Update a 'storage.PhysicalDiskUsage' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStoragePhysicalDiskUsage(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragepurearray",

										Short: "Update a 'storage.PureArray' resource.",
										Long: `Update a 'storage.PureArray' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStoragePureArray(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragesasexpander",

										Short: "Update a 'storage.SasExpander' resource.",
										Long: `Update a 'storage.SasExpander' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStorageSasExpander(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragesasport",

										Short: "Update a 'storage.SasPort' resource.",
										Long: `Update a 'storage.SasPort' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStorageSasPort(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragestoragepolicy",

										Short: "Update a 'storage.StoragePolicy' resource.",
										Long: `Update a 'storage.StoragePolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStorageStoragePolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragevdmemberep",

										Short: "Update a 'storage.VdMemberEp' resource.",
										Long: `Update a 'storage.VdMemberEp' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStorageVdMemberEp(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragevirtualdrive",

										Short: "Update a 'storage.VirtualDrive' resource.",
										Long: `Update a 'storage.VirtualDrive' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStorageVirtualDrive(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragevirtualdriveextension",

										Short: "Update a 'storage.VirtualDriveExtension' resource.",
										Long: `Update a 'storage.VirtualDriveExtension' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.UpdateStorageVirtualDriveExtension(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "syslog",

								Short: "Create Syslog resource(s)",
								Long: `Create Syslog resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SyslogApi.CreateSyslogPolicy(authCtx)

											body := openapi.NewSyslogPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.SyslogPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'syslog.Policy' resource.",
										Long: `Create a 'syslog.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesyslogpolicy",

										Short: "Update a 'syslog.Policy' resource.",
										Long: `Update a 'syslog.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SyslogApi.UpdateSyslogPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "tam",

								Short: "Create Tam resource(s)",
								Long: `Create Tam resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "advisorycount",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.TamApi.CreateTamAdvisoryCount(authCtx)

											body := openapi.NewTamAdvisoryCount()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.TamAdvisoryCount(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'tam.AdvisoryCount' resource.",
										Long: `Create a 'tam.AdvisoryCount' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "advisoryinfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.TamApi.CreateTamAdvisoryInfo(authCtx)

											body := openapi.NewTamAdvisoryInfo()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.TamAdvisoryInfo(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'tam.AdvisoryInfo' resource.",
										Long: `Create a 'tam.AdvisoryInfo' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "advisoryinstance",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.TamApi.CreateTamAdvisoryInstance(authCtx)

											body := openapi.NewTamAdvisoryInstance()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.TamAdvisoryInstance(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'tam.AdvisoryInstance' resource.",
										Long: `Create a 'tam.AdvisoryInstance' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "securityadvisory",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.TamApi.CreateTamSecurityAdvisory(authCtx)

											body := openapi.NewTamSecurityAdvisory()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.TamSecurityAdvisory(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'tam.SecurityAdvisory' resource.",
										Long: `Create a 'tam.SecurityAdvisory' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatetamadvisorycount",

										Short: "Update a 'tam.AdvisoryCount' resource.",
										Long: `Update a 'tam.AdvisoryCount' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TamApi.UpdateTamAdvisoryCount(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatetamadvisoryinfo",

										Short: "Update a 'tam.AdvisoryInfo' resource.",
										Long: `Update a 'tam.AdvisoryInfo' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TamApi.UpdateTamAdvisoryInfo(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatetamadvisoryinstance",

										Short: "Update a 'tam.AdvisoryInstance' resource.",
										Long: `Update a 'tam.AdvisoryInstance' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TamApi.UpdateTamAdvisoryInstance(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatetamsecurityadvisory",

										Short: "Update a 'tam.SecurityAdvisory' resource.",
										Long: `Update a 'tam.SecurityAdvisory' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TamApi.UpdateTamSecurityAdvisory(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "task",

								Short: "Create Task resource(s)",
								Long: `Create Task resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "purescopedinventory",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.TaskApi.CreateTaskPureScopedInventory(authCtx)

											body := openapi.NewTaskPureScopedInventory()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.TaskPureScopedInventory(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'task.PureScopedInventory' resource.",
										Long: `Create a 'task.PureScopedInventory' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "top",

								Short: "Create Top resource(s)",
								Long: `Create Top resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatetopsystem",

										Short: "Update a 'top.System' resource.",
										Long: `Update a 'top.System' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TopApi.UpdateTopSystem(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "virtualization",

								Short: "Create Virtualization resource(s)",
								Long: `Create Virtualization resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevirtualizationvmwarecluster",

										Short: "Update a 'virtualization.VmwareCluster' resource.",
										Long: `Update a 'virtualization.VmwareCluster' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VirtualizationApi.UpdateVirtualizationVmwareCluster(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevirtualizationvmwaredatacenter",

										Short: "Update a 'virtualization.VmwareDatacenter' resource.",
										Long: `Update a 'virtualization.VmwareDatacenter' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VirtualizationApi.UpdateVirtualizationVmwareDatacenter(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevirtualizationvmwaredatastore",

										Short: "Update a 'virtualization.VmwareDatastore' resource.",
										Long: `Update a 'virtualization.VmwareDatastore' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VirtualizationApi.UpdateVirtualizationVmwareDatastore(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevirtualizationvmwarehost",

										Short: "Update a 'virtualization.VmwareHost' resource.",
										Long: `Update a 'virtualization.VmwareHost' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VirtualizationApi.UpdateVirtualizationVmwareHost(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevirtualizationvmwarevirtualmachine",

										Short: "Update a 'virtualization.VmwareVirtualMachine' resource.",
										Long: `Update a 'virtualization.VmwareVirtualMachine' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VirtualizationApi.UpdateVirtualizationVmwareVirtualMachine(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "vmedia",

								Short: "Create Vmedia resource(s)",
								Long: `Create Vmedia resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VmediaApi.CreateVmediaPolicy(authCtx)

											body := openapi.NewVmediaPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.VmediaPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vmedia.Policy' resource.",
										Long: `Create a 'vmedia.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevmediapolicy",

										Short: "Update a 'vmedia.Policy' resource.",
										Long: `Update a 'vmedia.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VmediaApi.UpdateVmediaPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "vnic",

								Short: "Create Vnic resource(s)",
								Long: `Create Vnic resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethadapterpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.CreateVnicEthAdapterPolicy(authCtx)

											body := openapi.NewVnicEthAdapterPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.VnicEthAdapterPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.EthAdapterPolicy' resource.",
										Long: `Create a 'vnic.EthAdapterPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethif",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.CreateVnicEthIf(authCtx)

											body := openapi.NewVnicEthIf()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.VnicEthIf(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.EthIf' resource.",
										Long: `Create a 'vnic.EthIf' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethnetworkpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.CreateVnicEthNetworkPolicy(authCtx)

											body := openapi.NewVnicEthNetworkPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.VnicEthNetworkPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.EthNetworkPolicy' resource.",
										Long: `Create a 'vnic.EthNetworkPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethqospolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.CreateVnicEthQosPolicy(authCtx)

											body := openapi.NewVnicEthQosPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.VnicEthQosPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.EthQosPolicy' resource.",
										Long: `Create a 'vnic.EthQosPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fcadapterpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.CreateVnicFcAdapterPolicy(authCtx)

											body := openapi.NewVnicFcAdapterPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.VnicFcAdapterPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.FcAdapterPolicy' resource.",
										Long: `Create a 'vnic.FcAdapterPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fcif",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.CreateVnicFcIf(authCtx)

											body := openapi.NewVnicFcIf()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.VnicFcIf(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.FcIf' resource.",
										Long: `Create a 'vnic.FcIf' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fcnetworkpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.CreateVnicFcNetworkPolicy(authCtx)

											body := openapi.NewVnicFcNetworkPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.VnicFcNetworkPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.FcNetworkPolicy' resource.",
										Long: `Create a 'vnic.FcNetworkPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fcqospolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.CreateVnicFcQosPolicy(authCtx)

											body := openapi.NewVnicFcQosPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.VnicFcQosPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.FcQosPolicy' resource.",
										Long: `Create a 'vnic.FcQosPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "lanconnectivitypolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.CreateVnicLanConnectivityPolicy(authCtx)

											body := openapi.NewVnicLanConnectivityPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.VnicLanConnectivityPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.LanConnectivityPolicy' resource.",
										Long: `Create a 'vnic.LanConnectivityPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "sanconnectivitypolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.CreateVnicSanConnectivityPolicy(authCtx)

											body := openapi.NewVnicSanConnectivityPolicy()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.VnicSanConnectivityPolicy(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.SanConnectivityPolicy' resource.",
										Long: `Create a 'vnic.SanConnectivityPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicethadapterpolicy",

										Short: "Update a 'vnic.EthAdapterPolicy' resource.",
										Long: `Update a 'vnic.EthAdapterPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.UpdateVnicEthAdapterPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicethif",

										Short: "Update a 'vnic.EthIf' resource.",
										Long: `Update a 'vnic.EthIf' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.UpdateVnicEthIf(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicethnetworkpolicy",

										Short: "Update a 'vnic.EthNetworkPolicy' resource.",
										Long: `Update a 'vnic.EthNetworkPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.UpdateVnicEthNetworkPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicethqospolicy",

										Short: "Update a 'vnic.EthQosPolicy' resource.",
										Long: `Update a 'vnic.EthQosPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.UpdateVnicEthQosPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicfcadapterpolicy",

										Short: "Update a 'vnic.FcAdapterPolicy' resource.",
										Long: `Update a 'vnic.FcAdapterPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.UpdateVnicFcAdapterPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicfcif",

										Short: "Update a 'vnic.FcIf' resource.",
										Long: `Update a 'vnic.FcIf' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.UpdateVnicFcIf(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicfcnetworkpolicy",

										Short: "Update a 'vnic.FcNetworkPolicy' resource.",
										Long: `Update a 'vnic.FcNetworkPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.UpdateVnicFcNetworkPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicfcqospolicy",

										Short: "Update a 'vnic.FcQosPolicy' resource.",
										Long: `Update a 'vnic.FcQosPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.UpdateVnicFcQosPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevniclanconnectivitypolicy",

										Short: "Update a 'vnic.LanConnectivityPolicy' resource.",
										Long: `Update a 'vnic.LanConnectivityPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.UpdateVnicLanConnectivityPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicsanconnectivitypolicy",

										Short: "Update a 'vnic.SanConnectivityPolicy' resource.",
										Long: `Update a 'vnic.SanConnectivityPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.UpdateVnicSanConnectivityPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "workflow",

								Short: "Create Workflow resource(s)",
								Long: `Create Workflow resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "batchapiexecutor",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.WorkflowApi.CreateWorkflowBatchApiExecutor(authCtx)

											body := openapi.NewWorkflowBatchApiExecutor()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.WorkflowBatchApiExecutor(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'workflow.BatchApiExecutor' resource.",
										Long: `Create a 'workflow.BatchApiExecutor' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "customdatatypedefinition",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.WorkflowApi.CreateWorkflowCustomDataTypeDefinition(authCtx)

											body := openapi.NewWorkflowCustomDataTypeDefinition()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.WorkflowCustomDataTypeDefinition(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'workflow.CustomDataTypeDefinition' resource.",
										Long: `Create a 'workflow.CustomDataTypeDefinition' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "taskdefinition",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.WorkflowApi.CreateWorkflowTaskDefinition(authCtx)

											body := openapi.NewWorkflowTaskDefinition()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.WorkflowTaskDefinition(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'workflow.TaskDefinition' resource.",
										Long: `Create a 'workflow.TaskDefinition' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateworkflowbatchapiexecutor",

										Short: "Update a 'workflow.BatchApiExecutor' resource.",
										Long: `Update a 'workflow.BatchApiExecutor' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.UpdateWorkflowBatchApiExecutor(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateworkflowcustomdatatypedefinition",

										Short: "Update a 'workflow.CustomDataTypeDefinition' resource.",
										Long: `Update a 'workflow.CustomDataTypeDefinition' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.UpdateWorkflowCustomDataTypeDefinition(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateworkflowtaskdefinition",

										Short: "Update a 'workflow.TaskDefinition' resource.",
										Long: `Update a 'workflow.TaskDefinition' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.UpdateWorkflowTaskDefinition(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateworkflowtaskinfo",

										Short: "Update a 'workflow.TaskInfo' resource.",
										Long: `Update a 'workflow.TaskInfo' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.UpdateWorkflowTaskInfo(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateworkflowworkflowdefinition",

										Short: "Update a 'workflow.WorkflowDefinition' resource.",
										Long: `Update a 'workflow.WorkflowDefinition' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.UpdateWorkflowWorkflowDefinition(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateworkflowworkflowinfo",

										Short: "Update a 'workflow.WorkflowInfo' resource.",
										Long: `Update a 'workflow.WorkflowInfo' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.UpdateWorkflowWorkflowInfo(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "workflowdefinition",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.WorkflowApi.CreateWorkflowWorkflowDefinition(authCtx)

											body := openapi.NewWorkflowWorkflowDefinition()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.WorkflowWorkflowDefinition(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'workflow.WorkflowDefinition' resource.",
										Long: `Create a 'workflow.WorkflowDefinition' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "workflowinfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.WorkflowApi.CreateWorkflowWorkflowInfo(authCtx)

											body := openapi.NewWorkflowWorkflowInfo()

											// Gather body from JSON on stdin.
											err := json.NewDecoder(os.Stdin).Decode(&body)
											if err != nil {
												resultHandler(nil, nil, fmt.Errorf("Decoding JSON: %v", err))
												return
											}

											req = req.WorkflowWorkflowInfo(*body)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'workflow.WorkflowInfo' resource.",
										Long: `Create a 'workflow.WorkflowInfo' resource.

Provide resource body as JSON on standard input`,
									}

									return cmd
								}())

							return cmd
						}())

					return cmd
				}())

			cmd.AddCommand(
				func() *cobra.Command {
					cmd := &cobra.Command{
						Use: "delete",

						Short: "Delete resouce(s)",
						Long: `Delete resouce(s)
`,
					}

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "adapter",

								Short: "Delete Adapter resource(s)",
								Long: `Delete Adapter resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configpolicy",

										Short: "Delete a 'adapter.ConfigPolicy' resource.",
										Long: `Delete a 'adapter.ConfigPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AdapterApi.DeleteAdapterConfigPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "appliance",

								Short: "Delete Appliance resource(s)",
								Long: `Delete Appliance resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backup",

										Short: "Delete a 'appliance.Backup' resource.",
										Long: `Delete a 'appliance.Backup' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.DeleteApplianceBackup(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "restore",

										Short: "Delete a 'appliance.Restore' resource.",
										Long: `Delete a 'appliance.Restore' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.DeleteApplianceRestore(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "asset",

								Short: "Delete Asset resource(s)",
								Long: `Delete Asset resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deviceclaim",

										Short: "Delete a 'asset.DeviceClaim' resource.",
										Long: `Delete a 'asset.DeviceClaim' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AssetApi.DeleteAssetDeviceClaim(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deviceregistration",

										Short: "Deletes the resource representing the device connector. All associated REST resources will be deleted. In particular, inventory and operational data associated with this device will be deleted.",
										Long: `Deletes the resource representing the device connector. All associated REST resources will be deleted. In particular, inventory and operational data associated with this device will be deleted.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AssetApi.DeleteAssetDeviceRegistration(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "manageddevice",

										Short: "Delete a 'asset.ManagedDevice' resource.",
										Long: `Delete a 'asset.ManagedDevice' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AssetApi.DeleteAssetManagedDevice(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "bios",

								Short: "Delete Bios resource(s)",
								Long: `Delete Bios resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'bios.Policy' resource.",
										Long: `Delete a 'bios.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.BiosApi.DeleteBiosPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "boot",

								Short: "Delete Boot resource(s)",
								Long: `Delete Boot resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "precisionpolicy",

										Short: "Delete a 'boot.PrecisionPolicy' resource.",
										Long: `Delete a 'boot.PrecisionPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.BootApi.DeleteBootPrecisionPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "deviceconnector",

								Short: "Delete Deviceconnector resource(s)",
								Long: `Delete Deviceconnector resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'deviceconnector.Policy' resource.",
										Long: `Delete a 'deviceconnector.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.DeviceconnectorApi.DeleteDeviceconnectorPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "firmware",

								Short: "Delete Firmware resource(s)",
								Long: `Delete Firmware resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "distributable",

										Short: "Delete a 'firmware.Distributable' resource.",
										Long: `Delete a 'firmware.Distributable' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.DeleteFirmwareDistributable(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "driverdistributable",

										Short: "Delete a 'firmware.DriverDistributable' resource.",
										Long: `Delete a 'firmware.DriverDistributable' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.DeleteFirmwareDriverDistributable(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "serverconfigurationutilitydistributable",

										Short: "Delete a 'firmware.ServerConfigurationUtilityDistributable' resource.",
										Long: `Delete a 'firmware.ServerConfigurationUtilityDistributable' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.DeleteFirmwareServerConfigurationUtilityDistributable(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "upgrade",

										Short: "Delete a 'firmware.Upgrade' resource.",
										Long: `Delete a 'firmware.Upgrade' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.DeleteFirmwareUpgrade(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "hcl",

								Short: "Delete Hcl resource(s)",
								Long: `Delete Hcl resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hyperflexsoftwarecompatibilityinfo",

										Short: "Delete a 'hcl.HyperflexSoftwareCompatibilityInfo' resource.",
										Long: `Delete a 'hcl.HyperflexSoftwareCompatibilityInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HclApi.DeleteHclHyperflexSoftwareCompatibilityInfo(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "hyperflex",

								Short: "Delete Hyperflex resource(s)",
								Long: `Delete Hyperflex resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "appcatalog",

										Short: "Delete a 'hyperflex.AppCatalog' resource.",
										Long: `Delete a 'hyperflex.AppCatalog' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexAppCatalog(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "autosupportpolicy",

										Short: "Delete a 'hyperflex.AutoSupportPolicy' resource.",
										Long: `Delete a 'hyperflex.AutoSupportPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexAutoSupportPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "capabilityinfo",

										Short: "Delete a 'hyperflex.CapabilityInfo' resource.",
										Long: `Delete a 'hyperflex.CapabilityInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexCapabilityInfo(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "clusternetworkpolicy",

										Short: "Delete a 'hyperflex.ClusterNetworkPolicy' resource.",
										Long: `Delete a 'hyperflex.ClusterNetworkPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexClusterNetworkPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "clusterprofile",

										Short: "Delete a 'hyperflex.ClusterProfile' resource.",
										Long: `Delete a 'hyperflex.ClusterProfile' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexClusterProfile(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "clusterstoragepolicy",

										Short: "Delete a 'hyperflex.ClusterStoragePolicy' resource.",
										Long: `Delete a 'hyperflex.ClusterStoragePolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexClusterStoragePolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "extfcstoragepolicy",

										Short: "Delete a 'hyperflex.ExtFcStoragePolicy' resource.",
										Long: `Delete a 'hyperflex.ExtFcStoragePolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexExtFcStoragePolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "extiscsistoragepolicy",

										Short: "Delete a 'hyperflex.ExtIscsiStoragePolicy' resource.",
										Long: `Delete a 'hyperflex.ExtIscsiStoragePolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexExtIscsiStoragePolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "featurelimitexternal",

										Short: "Delete a 'hyperflex.FeatureLimitExternal' resource.",
										Long: `Delete a 'hyperflex.FeatureLimitExternal' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexFeatureLimitExternal(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "featurelimitinternal",

										Short: "Delete a 'hyperflex.FeatureLimitInternal' resource.",
										Long: `Delete a 'hyperflex.FeatureLimitInternal' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexFeatureLimitInternal(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hxdpversion",

										Short: "Delete a 'hyperflex.HxdpVersion' resource.",
										Long: `Delete a 'hyperflex.HxdpVersion' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexHxdpVersion(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "localcredentialpolicy",

										Short: "Delete a 'hyperflex.LocalCredentialPolicy' resource.",
										Long: `Delete a 'hyperflex.LocalCredentialPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexLocalCredentialPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "nodeconfigpolicy",

										Short: "Delete a 'hyperflex.NodeConfigPolicy' resource.",
										Long: `Delete a 'hyperflex.NodeConfigPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexNodeConfigPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "nodeprofile",

										Short: "Delete a 'hyperflex.NodeProfile' resource.",
										Long: `Delete a 'hyperflex.NodeProfile' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexNodeProfile(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "proxysettingpolicy",

										Short: "Delete a 'hyperflex.ProxySettingPolicy' resource.",
										Long: `Delete a 'hyperflex.ProxySettingPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexProxySettingPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "serverfirmwareversion",

										Short: "Delete a 'hyperflex.ServerFirmwareVersion' resource.",
										Long: `Delete a 'hyperflex.ServerFirmwareVersion' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexServerFirmwareVersion(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "servermodel",

										Short: "Delete a 'hyperflex.ServerModel' resource.",
										Long: `Delete a 'hyperflex.ServerModel' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexServerModel(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "softwareversionpolicy",

										Short: "Delete a 'hyperflex.SoftwareVersionPolicy' resource.",
										Long: `Delete a 'hyperflex.SoftwareVersionPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexSoftwareVersionPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "sysconfigpolicy",

										Short: "Delete a 'hyperflex.SysConfigPolicy' resource.",
										Long: `Delete a 'hyperflex.SysConfigPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexSysConfigPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ucsmconfigpolicy",

										Short: "Delete a 'hyperflex.UcsmConfigPolicy' resource.",
										Long: `Delete a 'hyperflex.UcsmConfigPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexUcsmConfigPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vcenterconfigpolicy",

										Short: "Delete a 'hyperflex.VcenterConfigPolicy' resource.",
										Long: `Delete a 'hyperflex.VcenterConfigPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.DeleteHyperflexVcenterConfigPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "iaas",

								Short: "Delete Iaas resource(s)",
								Long: `Delete Iaas resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ucsdinfo",

										Short: "Delete a 'iaas.UcsdInfo' resource.",
										Long: `Delete a 'iaas.UcsdInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IaasApi.DeleteIaasUcsdInfo(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "iam",

								Short: "Delete Iam resource(s)",
								Long: `Delete Iam resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "account",

										Short: "Delete a 'iam.Account' resource.",
										Long: `Delete a 'iam.Account' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamAccount(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "apikey",

										Short: "Delete a 'iam.ApiKey' resource.",
										Long: `Delete a 'iam.ApiKey' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamApiKey(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "appregistration",

										Short: "Delete a 'iam.AppRegistration' resource.",
										Long: `Delete a 'iam.AppRegistration' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamAppRegistration(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "certificate",

										Short: "Delete a 'iam.Certificate' resource.",
										Long: `Delete a 'iam.Certificate' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamCertificate(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "certificaterequest",

										Short: "Delete a 'iam.CertificateRequest' resource.",
										Long: `Delete a 'iam.CertificateRequest' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamCertificateRequest(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "endpointuser",

										Short: "Delete a 'iam.EndPointUser' resource.",
										Long: `Delete a 'iam.EndPointUser' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamEndPointUser(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "endpointuserpolicy",

										Short: "Delete a 'iam.EndPointUserPolicy' resource.",
										Long: `Delete a 'iam.EndPointUserPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamEndPointUserPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "endpointuserrole",

										Short: "Delete a 'iam.EndPointUserRole' resource.",
										Long: `Delete a 'iam.EndPointUserRole' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamEndPointUserRole(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "idp",

										Short: "Delete a 'iam.Idp' resource.",
										Long: `Delete a 'iam.Idp' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamIdp(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ldapgroup",

										Short: "Delete a 'iam.LdapGroup' resource.",
										Long: `Delete a 'iam.LdapGroup' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamLdapGroup(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ldappolicy",

										Short: "Delete a 'iam.LdapPolicy' resource.",
										Long: `Delete a 'iam.LdapPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamLdapPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ldapprovider",

										Short: "Delete a 'iam.LdapProvider' resource.",
										Long: `Delete a 'iam.LdapProvider' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamLdapProvider(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "oauthtoken",

										Short: "Delete a 'iam.OAuthToken' resource.",
										Long: `Delete a 'iam.OAuthToken' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamOAuthToken(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "permission",

										Short: "Delete a 'iam.Permission' resource.",
										Long: `Delete a 'iam.Permission' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamPermission(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "privatekeyspec",

										Short: "Delete a 'iam.PrivateKeySpec' resource.",
										Long: `Delete a 'iam.PrivateKeySpec' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamPrivateKeySpec(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "qualifier",

										Short: "Delete a 'iam.Qualifier' resource.",
										Long: `Delete a 'iam.Qualifier' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamQualifier(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "resourceroles",

										Short: "Delete a 'iam.ResourceRoles' resource.",
										Long: `Delete a 'iam.ResourceRoles' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamResourceRoles(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "session",

										Short: "Delete a 'iam.Session' resource.",
										Long: `Delete a 'iam.Session' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamSession(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "sessionlimits",

										Short: "Delete a 'iam.SessionLimits' resource.",
										Long: `Delete a 'iam.SessionLimits' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamSessionLimits(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "trustpoint",

										Short: "Delete a 'iam.TrustPoint' resource.",
										Long: `Delete a 'iam.TrustPoint' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamTrustPoint(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "user",

										Short: "Delete a 'iam.User' resource.",
										Long: `Delete a 'iam.User' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamUser(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "usergroup",

										Short: "Delete a 'iam.UserGroup' resource.",
										Long: `Delete a 'iam.UserGroup' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.DeleteIamUserGroup(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "infra",

								Short: "Delete Infra resource(s)",
								Long: `Delete Infra resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "accountexperience",

										Short: "Delete a 'infra.AccountExperience' resource.",
										Long: `Delete a 'infra.AccountExperience' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.InfraApi.DeleteInfraAccountExperience(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ipmioverlan",

								Short: "Delete Ipmioverlan resource(s)",
								Long: `Delete Ipmioverlan resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'ipmioverlan.Policy' resource.",
										Long: `Delete a 'ipmioverlan.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IpmioverlanApi.DeleteIpmioverlanPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "kvm",

								Short: "Delete Kvm resource(s)",
								Long: `Delete Kvm resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'kvm.Policy' resource.",
										Long: `Delete a 'kvm.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.KvmApi.DeleteKvmPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "memory",

								Short: "Delete Memory resource(s)",
								Long: `Delete Memory resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "persistentmemorypolicy",

										Short: "Delete a 'memory.PersistentMemoryPolicy' resource.",
										Long: `Delete a 'memory.PersistentMemoryPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.DeleteMemoryPersistentMemoryPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "meta",

								Short: "Delete Meta resource(s)",
								Long: `Delete Meta resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "definition",

										Short: "Delete a 'meta.Definition' resource.",
										Long: `Delete a 'meta.Definition' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MetaApi.DeleteMetaDefinition(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "networkconfig",

								Short: "Delete Networkconfig resource(s)",
								Long: `Delete Networkconfig resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'networkconfig.Policy' resource.",
										Long: `Delete a 'networkconfig.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NetworkconfigApi.DeleteNetworkconfigPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ntp",

								Short: "Delete Ntp resource(s)",
								Long: `Delete Ntp resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'ntp.Policy' resource.",
										Long: `Delete a 'ntp.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NtpApi.DeleteNtpPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "organization",

								Short: "Delete Organization resource(s)",
								Long: `Delete Organization resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "organization",

										Short: "Delete a 'organization.Organization' resource.",
										Long: `Delete a 'organization.Organization' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.OrganizationApi.DeleteOrganizationOrganization(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "os",

								Short: "Delete Os resource(s)",
								Long: `Delete Os resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configurationfile",

										Short: "Delete a 'os.ConfigurationFile' resource.",
										Long: `Delete a 'os.ConfigurationFile' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.OsApi.DeleteOsConfigurationFile(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "recovery",

								Short: "Delete Recovery resource(s)",
								Long: `Delete Recovery resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backupconfigpolicy",

										Short: "Delete a 'recovery.BackupConfigPolicy' resource.",
										Long: `Delete a 'recovery.BackupConfigPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.DeleteRecoveryBackupConfigPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backupprofile",

										Short: "Delete a 'recovery.BackupProfile' resource.",
										Long: `Delete a 'recovery.BackupProfile' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.DeleteRecoveryBackupProfile(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ondemandbackup",

										Short: "Delete a 'recovery.OnDemandBackup' resource.",
										Long: `Delete a 'recovery.OnDemandBackup' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.DeleteRecoveryOnDemandBackup(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "restore",

										Short: "Delete a 'recovery.Restore' resource.",
										Long: `Delete a 'recovery.Restore' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.DeleteRecoveryRestore(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "scheduleconfigpolicy",

										Short: "Delete a 'recovery.ScheduleConfigPolicy' resource.",
										Long: `Delete a 'recovery.ScheduleConfigPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.DeleteRecoveryScheduleConfigPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "resource",

								Short: "Delete Resource resource(s)",
								Long: `Delete Resource resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "group",

										Short: "Delete a 'resource.Group' resource.",
										Long: `Delete a 'resource.Group' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ResourceApi.DeleteResourceGroup(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "sdcard",

								Short: "Delete Sdcard resource(s)",
								Long: `Delete Sdcard resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'sdcard.Policy' resource.",
										Long: `Delete a 'sdcard.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdcardApi.DeleteSdcardPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "sdwan",

								Short: "Delete Sdwan resource(s)",
								Long: `Delete Sdwan resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "profile",

										Short: "Delete a 'sdwan.Profile' resource.",
										Long: `Delete a 'sdwan.Profile' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdwanApi.DeleteSdwanProfile(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "routernode",

										Short: "Delete a 'sdwan.RouterNode' resource.",
										Long: `Delete a 'sdwan.RouterNode' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdwanApi.DeleteSdwanRouterNode(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "routerpolicy",

										Short: "Delete a 'sdwan.RouterPolicy' resource.",
										Long: `Delete a 'sdwan.RouterPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdwanApi.DeleteSdwanRouterPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vmanageaccountpolicy",

										Short: "Delete a 'sdwan.VmanageAccountPolicy' resource.",
										Long: `Delete a 'sdwan.VmanageAccountPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdwanApi.DeleteSdwanVmanageAccountPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "server",

								Short: "Delete Server resource(s)",
								Long: `Delete Server resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "profile",

										Short: "Delete a 'server.Profile' resource.",
										Long: `Delete a 'server.Profile' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ServerApi.DeleteServerProfile(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "smtp",

								Short: "Delete Smtp resource(s)",
								Long: `Delete Smtp resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'smtp.Policy' resource.",
										Long: `Delete a 'smtp.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SmtpApi.DeleteSmtpPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "snmp",

								Short: "Delete Snmp resource(s)",
								Long: `Delete Snmp resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'snmp.Policy' resource.",
										Long: `Delete a 'snmp.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SnmpApi.DeleteSnmpPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "software",

								Short: "Delete Software resource(s)",
								Long: `Delete Software resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hclmeta",

										Short: "Delete a 'software.HclMeta' resource.",
										Long: `Delete a 'software.HclMeta' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwareApi.DeleteSoftwareHclMeta(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hyperflexdistributable",

										Short: "Delete a 'software.HyperflexDistributable' resource.",
										Long: `Delete a 'software.HyperflexDistributable' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwareApi.DeleteSoftwareHyperflexDistributable(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "solutiondistributable",

										Short: "Delete a 'software.SolutionDistributable' resource.",
										Long: `Delete a 'software.SolutionDistributable' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwareApi.DeleteSoftwareSolutionDistributable(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "softwarerepository",

								Short: "Delete Softwarerepository resource(s)",
								Long: `Delete Softwarerepository resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "operatingsystemfile",

										Short: "Delete a 'softwarerepository.OperatingSystemFile' resource.",
										Long: `Delete a 'softwarerepository.OperatingSystemFile' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwarerepositoryApi.DeleteSoftwarerepositoryOperatingSystemFile(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "sol",

								Short: "Delete Sol resource(s)",
								Long: `Delete Sol resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'sol.Policy' resource.",
										Long: `Delete a 'sol.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SolApi.DeleteSolPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ssh",

								Short: "Delete Ssh resource(s)",
								Long: `Delete Ssh resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'ssh.Policy' resource.",
										Long: `Delete a 'ssh.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SshApi.DeleteSshPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "storage",

								Short: "Delete Storage resource(s)",
								Long: `Delete Storage resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "diskgrouppolicy",

										Short: "Delete a 'storage.DiskGroupPolicy' resource.",
										Long: `Delete a 'storage.DiskGroupPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.DeleteStorageDiskGroupPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "storagepolicy",

										Short: "Delete a 'storage.StoragePolicy' resource.",
										Long: `Delete a 'storage.StoragePolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.DeleteStorageStoragePolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "syslog",

								Short: "Delete Syslog resource(s)",
								Long: `Delete Syslog resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'syslog.Policy' resource.",
										Long: `Delete a 'syslog.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SyslogApi.DeleteSyslogPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "tam",

								Short: "Delete Tam resource(s)",
								Long: `Delete Tam resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "advisorycount",

										Short: "Delete a 'tam.AdvisoryCount' resource.",
										Long: `Delete a 'tam.AdvisoryCount' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TamApi.DeleteTamAdvisoryCount(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "advisoryinfo",

										Short: "Delete a 'tam.AdvisoryInfo' resource.",
										Long: `Delete a 'tam.AdvisoryInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TamApi.DeleteTamAdvisoryInfo(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "advisoryinstance",

										Short: "Delete a 'tam.AdvisoryInstance' resource.",
										Long: `Delete a 'tam.AdvisoryInstance' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TamApi.DeleteTamAdvisoryInstance(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "securityadvisory",

										Short: "Delete a 'tam.SecurityAdvisory' resource.",
										Long: `Delete a 'tam.SecurityAdvisory' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TamApi.DeleteTamSecurityAdvisory(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ucsd",

								Short: "Delete Ucsd resource(s)",
								Long: `Delete Ucsd resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backupinfo",

										Short: "Delete a 'ucsd.BackupInfo' resource.",
										Long: `Delete a 'ucsd.BackupInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.UcsdApi.DeleteUcsdBackupInfo(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "vmedia",

								Short: "Delete Vmedia resource(s)",
								Long: `Delete Vmedia resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'vmedia.Policy' resource.",
										Long: `Delete a 'vmedia.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VmediaApi.DeleteVmediaPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "vnic",

								Short: "Delete Vnic resource(s)",
								Long: `Delete Vnic resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethadapterpolicy",

										Short: "Delete a 'vnic.EthAdapterPolicy' resource.",
										Long: `Delete a 'vnic.EthAdapterPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.DeleteVnicEthAdapterPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethif",

										Short: "Delete a 'vnic.EthIf' resource.",
										Long: `Delete a 'vnic.EthIf' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.DeleteVnicEthIf(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethnetworkpolicy",

										Short: "Delete a 'vnic.EthNetworkPolicy' resource.",
										Long: `Delete a 'vnic.EthNetworkPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.DeleteVnicEthNetworkPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethqospolicy",

										Short: "Delete a 'vnic.EthQosPolicy' resource.",
										Long: `Delete a 'vnic.EthQosPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.DeleteVnicEthQosPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fcadapterpolicy",

										Short: "Delete a 'vnic.FcAdapterPolicy' resource.",
										Long: `Delete a 'vnic.FcAdapterPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.DeleteVnicFcAdapterPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fcif",

										Short: "Delete a 'vnic.FcIf' resource.",
										Long: `Delete a 'vnic.FcIf' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.DeleteVnicFcIf(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fcnetworkpolicy",

										Short: "Delete a 'vnic.FcNetworkPolicy' resource.",
										Long: `Delete a 'vnic.FcNetworkPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.DeleteVnicFcNetworkPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fcqospolicy",

										Short: "Delete a 'vnic.FcQosPolicy' resource.",
										Long: `Delete a 'vnic.FcQosPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.DeleteVnicFcQosPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "lanconnectivitypolicy",

										Short: "Delete a 'vnic.LanConnectivityPolicy' resource.",
										Long: `Delete a 'vnic.LanConnectivityPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.DeleteVnicLanConnectivityPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "sanconnectivitypolicy",

										Short: "Delete a 'vnic.SanConnectivityPolicy' resource.",
										Long: `Delete a 'vnic.SanConnectivityPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.DeleteVnicSanConnectivityPolicy(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "workflow",

								Short: "Delete Workflow resource(s)",
								Long: `Delete Workflow resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "batchapiexecutor",

										Short: "Delete a 'workflow.BatchApiExecutor' resource.",
										Long: `Delete a 'workflow.BatchApiExecutor' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.DeleteWorkflowBatchApiExecutor(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "customdatatypedefinition",

										Short: "Delete a 'workflow.CustomDataTypeDefinition' resource.",
										Long: `Delete a 'workflow.CustomDataTypeDefinition' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.DeleteWorkflowCustomDataTypeDefinition(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "taskdefinition",

										Short: "Delete a 'workflow.TaskDefinition' resource.",
										Long: `Delete a 'workflow.TaskDefinition' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.DeleteWorkflowTaskDefinition(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "workflowdefinition",

										Short: "Delete a 'workflow.WorkflowDefinition' resource.",
										Long: `Delete a 'workflow.WorkflowDefinition' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.DeleteWorkflowWorkflowDefinition(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "workflowinfo",

										Short: "Delete a 'workflow.WorkflowInfo' resource.",
										Long: `Delete a 'workflow.WorkflowInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.DeleteWorkflowWorkflowInfo(authCtx, args[0])

													httpResponse, err := req.Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					return cmd
				}())

			cmd.AddCommand(
				func() *cobra.Command {
					cmd := &cobra.Command{
						Use: "get",

						Short: "Get or list resouce(s)",
						Long: `Get or list resouce(s)
`,
					}

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "aaa",

								Short: "Get or list Aaa resource(s)",
								Long: `Get or list Aaa resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "auditrecord",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.AaaApi.GetAaaAuditRecordList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'aaa.AuditRecord' resource.",
										Long: `Read a 'aaa.AuditRecord' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AaaApi.GetAaaAuditRecordByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "adapter",

								Short: "Get or list Adapter resource(s)",
								Long: `Get or list Adapter resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.AdapterApi.GetAdapterConfigPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'adapter.ConfigPolicy' resource.",
										Long: `Read a 'adapter.ConfigPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AdapterApi.GetAdapterConfigPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "extethinterface",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.AdapterApi.GetAdapterExtEthInterfaceList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'adapter.ExtEthInterface' resource.",
										Long: `Read a 'adapter.ExtEthInterface' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AdapterApi.GetAdapterExtEthInterfaceByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hostethinterface",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.AdapterApi.GetAdapterHostEthInterfaceList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'adapter.HostEthInterface' resource.",
										Long: `Read a 'adapter.HostEthInterface' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AdapterApi.GetAdapterHostEthInterfaceByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hostfcinterface",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.AdapterApi.GetAdapterHostFcInterfaceList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'adapter.HostFcInterface' resource.",
										Long: `Read a 'adapter.HostFcInterface' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AdapterApi.GetAdapterHostFcInterfaceByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hostiscsiinterface",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.AdapterApi.GetAdapterHostIscsiInterfaceList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'adapter.HostIscsiInterface' resource.",
										Long: `Read a 'adapter.HostIscsiInterface' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AdapterApi.GetAdapterHostIscsiInterfaceByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "unit",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.AdapterApi.GetAdapterUnitList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'adapter.Unit' resource.",
										Long: `Read a 'adapter.Unit' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AdapterApi.GetAdapterUnitByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "appliance",

								Short: "Get or list Appliance resource(s)",
								Long: `Get or list Appliance resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backup",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.GetApplianceBackupList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.Backup' resource.",
										Long: `Read a 'appliance.Backup' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.GetApplianceBackupByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backuppolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.GetApplianceBackupPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.BackupPolicy' resource.",
										Long: `Read a 'appliance.BackupPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.GetApplianceBackupPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "certificatesetting",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.GetApplianceCertificateSettingList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.CertificateSetting' resource.",
										Long: `Read a 'appliance.CertificateSetting' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.GetApplianceCertificateSettingByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "dataexportpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.GetApplianceDataExportPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.DataExportPolicy' resource.",
										Long: `Read a 'appliance.DataExportPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.GetApplianceDataExportPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deviceclaim",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.GetApplianceDeviceClaimList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.DeviceClaim' resource.",
										Long: `Read a 'appliance.DeviceClaim' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.GetApplianceDeviceClaimByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "diagsetting",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.GetApplianceDiagSettingList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.DiagSetting' resource.",
										Long: `Read a 'appliance.DiagSetting' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.GetApplianceDiagSettingByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "imagebundle",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.GetApplianceImageBundleList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.ImageBundle' resource.",
										Long: `Read a 'appliance.ImageBundle' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.GetApplianceImageBundleByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "nodeinfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.GetApplianceNodeInfoList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.NodeInfo' resource.",
										Long: `Read a 'appliance.NodeInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.GetApplianceNodeInfoByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "releasenote",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.GetApplianceReleaseNoteList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.ReleaseNote' resource.",
										Long: `Read a 'appliance.ReleaseNote' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.GetApplianceReleaseNoteByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "restore",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.GetApplianceRestoreList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.Restore' resource.",
										Long: `Read a 'appliance.Restore' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.GetApplianceRestoreByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "setupinfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.GetApplianceSetupInfoList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.SetupInfo' resource.",
										Long: `Read a 'appliance.SetupInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.GetApplianceSetupInfoByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "systeminfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.GetApplianceSystemInfoList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.SystemInfo' resource.",
										Long: `Read a 'appliance.SystemInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.GetApplianceSystemInfoByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "upgrade",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.GetApplianceUpgradeList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.Upgrade' resource.",
										Long: `Read a 'appliance.Upgrade' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.GetApplianceUpgradeByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "upgradepolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ApplianceApi.GetApplianceUpgradePolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.UpgradePolicy' resource.",
										Long: `Read a 'appliance.UpgradePolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.GetApplianceUpgradePolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "asset",

								Short: "Get or list Asset resource(s)",
								Long: `Get or list Asset resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "clustermember",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.AssetApi.GetAssetClusterMemberList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'asset.ClusterMember' resource.",
										Long: `Read a 'asset.ClusterMember' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AssetApi.GetAssetClusterMemberByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deviceconfiguration",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.AssetApi.GetAssetDeviceConfigurationList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'asset.DeviceConfiguration' resource.",
										Long: `Read a 'asset.DeviceConfiguration' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AssetApi.GetAssetDeviceConfigurationByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deviceconnectormanager",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.AssetApi.GetAssetDeviceConnectorManagerList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'asset.DeviceConnectorManager' resource.",
										Long: `Read a 'asset.DeviceConnectorManager' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AssetApi.GetAssetDeviceConnectorManagerByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "devicecontractinformation",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.AssetApi.GetAssetDeviceContractInformationList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'asset.DeviceContractInformation' resource.",
										Long: `Read a 'asset.DeviceContractInformation' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AssetApi.GetAssetDeviceContractInformationByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deviceregistration",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.AssetApi.GetAssetDeviceRegistrationList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'asset.DeviceRegistration' resource.",
										Long: `Read a 'asset.DeviceRegistration' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AssetApi.GetAssetDeviceRegistrationByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "manageddevice",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.AssetApi.GetAssetManagedDeviceList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'asset.ManagedDevice' resource.",
										Long: `Read a 'asset.ManagedDevice' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AssetApi.GetAssetManagedDeviceByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "bios",

								Short: "Get or list Bios resource(s)",
								Long: `Get or list Bios resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "bootmode",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.BiosApi.GetBiosBootModeList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'bios.BootMode' resource.",
										Long: `Read a 'bios.BootMode' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.BiosApi.GetBiosBootModeByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.BiosApi.GetBiosPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'bios.Policy' resource.",
										Long: `Read a 'bios.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.BiosApi.GetBiosPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "unit",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.BiosApi.GetBiosUnitList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'bios.Unit' resource.",
										Long: `Read a 'bios.Unit' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.BiosApi.GetBiosUnitByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "boot",

								Short: "Get or list Boot resource(s)",
								Long: `Get or list Boot resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "devicebootmode",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.BootApi.GetBootDeviceBootModeList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'boot.DeviceBootMode' resource.",
										Long: `Read a 'boot.DeviceBootMode' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.BootApi.GetBootDeviceBootModeByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "precisionpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.BootApi.GetBootPrecisionPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'boot.PrecisionPolicy' resource.",
										Long: `Read a 'boot.PrecisionPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.BootApi.GetBootPrecisionPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "compute",

								Short: "Get or list Compute resource(s)",
								Long: `Get or list Compute resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "blade",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ComputeApi.GetComputeBladeList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'compute.Blade' resource.",
										Long: `Read a 'compute.Blade' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ComputeApi.GetComputeBladeByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "board",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ComputeApi.GetComputeBoardList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'compute.Board' resource.",
										Long: `Read a 'compute.Board' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ComputeApi.GetComputeBoardByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "physicalsummary",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ComputeApi.GetComputePhysicalSummaryList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'compute.PhysicalSummary' resource.",
										Long: `Read a 'compute.PhysicalSummary' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ComputeApi.GetComputePhysicalSummaryByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "rackunit",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ComputeApi.GetComputeRackUnitList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'compute.RackUnit' resource.",
										Long: `Read a 'compute.RackUnit' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ComputeApi.GetComputeRackUnitByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "serversetting",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ComputeApi.GetComputeServerSettingList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'compute.ServerSetting' resource.",
										Long: `Read a 'compute.ServerSetting' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ComputeApi.GetComputeServerSettingByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "cond",

								Short: "Get or list Cond resource(s)",
								Long: `Get or list Cond resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "alarm",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.CondApi.GetCondAlarmList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'cond.Alarm' resource.",
										Long: `Read a 'cond.Alarm' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.CondApi.GetCondAlarmByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hclstatus",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.CondApi.GetCondHclStatusList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'cond.HclStatus' resource.",
										Long: `Read a 'cond.HclStatus' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.CondApi.GetCondHclStatusByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hclstatusdetail",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.CondApi.GetCondHclStatusDetailList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'cond.HclStatusDetail' resource.",
										Long: `Read a 'cond.HclStatusDetail' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.CondApi.GetCondHclStatusDetailByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hclstatusjob",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.CondApi.GetCondHclStatusJobList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'cond.HclStatusJob' resource.",
										Long: `Read a 'cond.HclStatusJob' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.CondApi.GetCondHclStatusJobByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "deviceconnector",

								Short: "Get or list Deviceconnector resource(s)",
								Long: `Get or list Deviceconnector resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.DeviceconnectorApi.GetDeviceconnectorPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'deviceconnector.Policy' resource.",
										Long: `Read a 'deviceconnector.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.DeviceconnectorApi.GetDeviceconnectorPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "equipment",

								Short: "Get or list Equipment resource(s)",
								Long: `Get or list Equipment resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "chassis",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.EquipmentApi.GetEquipmentChassisList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.Chassis' resource.",
										Long: `Read a 'equipment.Chassis' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.GetEquipmentChassisByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "devicesummary",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.EquipmentApi.GetEquipmentDeviceSummaryList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.DeviceSummary' resource.",
										Long: `Read a 'equipment.DeviceSummary' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.GetEquipmentDeviceSummaryByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fan",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.EquipmentApi.GetEquipmentFanList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.Fan' resource.",
										Long: `Read a 'equipment.Fan' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.GetEquipmentFanByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fanmodule",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.EquipmentApi.GetEquipmentFanModuleList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.FanModule' resource.",
										Long: `Read a 'equipment.FanModule' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.GetEquipmentFanModuleByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fex",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.EquipmentApi.GetEquipmentFexList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.Fex' resource.",
										Long: `Read a 'equipment.Fex' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.GetEquipmentFexByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "iocard",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.EquipmentApi.GetEquipmentIoCardList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.IoCard' resource.",
										Long: `Read a 'equipment.IoCard' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.GetEquipmentIoCardByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ioexpander",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.EquipmentApi.GetEquipmentIoExpanderList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.IoExpander' resource.",
										Long: `Read a 'equipment.IoExpander' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.GetEquipmentIoExpanderByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "locatorled",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.EquipmentApi.GetEquipmentLocatorLedList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.LocatorLed' resource.",
										Long: `Read a 'equipment.LocatorLed' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.GetEquipmentLocatorLedByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "psu",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.EquipmentApi.GetEquipmentPsuList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.Psu' resource.",
										Long: `Read a 'equipment.Psu' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.GetEquipmentPsuByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "rackenclosure",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.EquipmentApi.GetEquipmentRackEnclosureList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.RackEnclosure' resource.",
										Long: `Read a 'equipment.RackEnclosure' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.GetEquipmentRackEnclosureByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "rackenclosureslot",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.EquipmentApi.GetEquipmentRackEnclosureSlotList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.RackEnclosureSlot' resource.",
										Long: `Read a 'equipment.RackEnclosureSlot' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.GetEquipmentRackEnclosureSlotByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "sharediomodule",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.EquipmentApi.GetEquipmentSharedIoModuleList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.SharedIoModule' resource.",
										Long: `Read a 'equipment.SharedIoModule' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.GetEquipmentSharedIoModuleByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "switchcard",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.EquipmentApi.GetEquipmentSwitchCardList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.SwitchCard' resource.",
										Long: `Read a 'equipment.SwitchCard' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.GetEquipmentSwitchCardByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "systemiocontroller",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.EquipmentApi.GetEquipmentSystemIoControllerList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.SystemIoController' resource.",
										Long: `Read a 'equipment.SystemIoController' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.GetEquipmentSystemIoControllerByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "tpm",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.EquipmentApi.GetEquipmentTpmList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.Tpm' resource.",
										Long: `Read a 'equipment.Tpm' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.GetEquipmentTpmByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ether",

								Short: "Get or list Ether resource(s)",
								Long: `Get or list Ether resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "physicalport",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.EtherApi.GetEtherPhysicalPortList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'ether.PhysicalPort' resource.",
										Long: `Read a 'ether.PhysicalPort' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EtherApi.GetEtherPhysicalPortByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "externalsite",

								Short: "Get or list Externalsite resource(s)",
								Long: `Get or list Externalsite resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "authorization",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ExternalsiteApi.GetExternalsiteAuthorizationList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'externalsite.Authorization' resource.",
										Long: `Read a 'externalsite.Authorization' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ExternalsiteApi.GetExternalsiteAuthorizationByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "fault",

								Short: "Get or list Fault resource(s)",
								Long: `Get or list Fault resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "instance",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.FaultApi.GetFaultInstanceList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'fault.Instance' resource.",
										Long: `Read a 'fault.Instance' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FaultApi.GetFaultInstanceByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "fc",

								Short: "Get or list Fc resource(s)",
								Long: `Get or list Fc resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "physicalport",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.FcApi.GetFcPhysicalPortList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'fc.PhysicalPort' resource.",
										Long: `Read a 'fc.PhysicalPort' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FcApi.GetFcPhysicalPortByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "firmware",

								Short: "Get or list Firmware resource(s)",
								Long: `Get or list Firmware resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "distributable",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.FirmwareApi.GetFirmwareDistributableList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'firmware.Distributable' resource.",
										Long: `Read a 'firmware.Distributable' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.GetFirmwareDistributableByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "driverdistributable",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.FirmwareApi.GetFirmwareDriverDistributableList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'firmware.DriverDistributable' resource.",
										Long: `Read a 'firmware.DriverDistributable' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.GetFirmwareDriverDistributableByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "eula",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.FirmwareApi.GetFirmwareEulaList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'firmware.Eula' resource.",
										Long: `Read a 'firmware.Eula' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.GetFirmwareEulaByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "runningfirmware",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.FirmwareApi.GetFirmwareRunningFirmwareList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'firmware.RunningFirmware' resource.",
										Long: `Read a 'firmware.RunningFirmware' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.GetFirmwareRunningFirmwareByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "serverconfigurationutilitydistributable",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.FirmwareApi.GetFirmwareServerConfigurationUtilityDistributableList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'firmware.ServerConfigurationUtilityDistributable' resource.",
										Long: `Read a 'firmware.ServerConfigurationUtilityDistributable' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.GetFirmwareServerConfigurationUtilityDistributableByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "upgrade",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.FirmwareApi.GetFirmwareUpgradeList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'firmware.Upgrade' resource.",
										Long: `Read a 'firmware.Upgrade' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.GetFirmwareUpgradeByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "upgradestatus",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.FirmwareApi.GetFirmwareUpgradeStatusList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'firmware.UpgradeStatus' resource.",
										Long: `Read a 'firmware.UpgradeStatus' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.GetFirmwareUpgradeStatusByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "forecast",

								Short: "Get or list Forecast resource(s)",
								Long: `Get or list Forecast resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "catalog",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ForecastApi.GetForecastCatalogList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'forecast.Catalog' resource.",
										Long: `Read a 'forecast.Catalog' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ForecastApi.GetForecastCatalogByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "definition",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ForecastApi.GetForecastDefinitionList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'forecast.Definition' resource.",
										Long: `Read a 'forecast.Definition' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ForecastApi.GetForecastDefinitionByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "instance",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ForecastApi.GetForecastInstanceList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'forecast.Instance' resource.",
										Long: `Read a 'forecast.Instance' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ForecastApi.GetForecastInstanceByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "graphics",

								Short: "Get or list Graphics resource(s)",
								Long: `Get or list Graphics resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "card",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.GraphicsApi.GetGraphicsCardList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'graphics.Card' resource.",
										Long: `Read a 'graphics.Card' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.GraphicsApi.GetGraphicsCardByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "controller",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.GraphicsApi.GetGraphicsControllerList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'graphics.Controller' resource.",
										Long: `Read a 'graphics.Controller' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.GraphicsApi.GetGraphicsControllerByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "hcl",

								Short: "Get or list Hcl resource(s)",
								Long: `Get or list Hcl resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "driverimage",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HclApi.GetHclDriverImageList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hcl.DriverImage' resource.",
										Long: `Read a 'hcl.DriverImage' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HclApi.GetHclDriverImageByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "exemptedcatalog",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HclApi.GetHclExemptedCatalogList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hcl.ExemptedCatalog' resource.",
										Long: `Read a 'hcl.ExemptedCatalog' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HclApi.GetHclExemptedCatalogByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hyperflexsoftwarecompatibilityinfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HclApi.GetHclHyperflexSoftwareCompatibilityInfoList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hcl.HyperflexSoftwareCompatibilityInfo' resource.",
										Long: `Read a 'hcl.HyperflexSoftwareCompatibilityInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HclApi.GetHclHyperflexSoftwareCompatibilityInfoByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "operatingsystem",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HclApi.GetHclOperatingSystemList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hcl.OperatingSystem' resource.",
										Long: `Read a 'hcl.OperatingSystem' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HclApi.GetHclOperatingSystemByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "operatingsystemvendor",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HclApi.GetHclOperatingSystemVendorList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hcl.OperatingSystemVendor' resource.",
										Long: `Read a 'hcl.OperatingSystemVendor' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HclApi.GetHclOperatingSystemVendorByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "servicestatus",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HclApi.GetHclServiceStatusList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hcl.ServiceStatus' resource.",
										Long: `Read a 'hcl.ServiceStatus' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HclApi.GetHclServiceStatusByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "hyperflex",

								Short: "Get or list Hyperflex resource(s)",
								Long: `Get or list Hyperflex resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "alarm",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexAlarmList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.Alarm' resource.",
										Long: `Read a 'hyperflex.Alarm' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexAlarmByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "appcatalog",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexAppCatalogList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.AppCatalog' resource.",
										Long: `Read a 'hyperflex.AppCatalog' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexAppCatalogByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "autosupportpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexAutoSupportPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.AutoSupportPolicy' resource.",
										Long: `Read a 'hyperflex.AutoSupportPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexAutoSupportPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "capabilityinfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexCapabilityInfoList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.CapabilityInfo' resource.",
										Long: `Read a 'hyperflex.CapabilityInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexCapabilityInfoByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "cluster",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexClusterList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.Cluster' resource.",
										Long: `Read a 'hyperflex.Cluster' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexClusterByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "clusternetworkpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexClusterNetworkPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ClusterNetworkPolicy' resource.",
										Long: `Read a 'hyperflex.ClusterNetworkPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexClusterNetworkPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "clusterprofile",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexClusterProfileList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ClusterProfile' resource.",
										Long: `Read a 'hyperflex.ClusterProfile' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexClusterProfileByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "clusterstoragepolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexClusterStoragePolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ClusterStoragePolicy' resource.",
										Long: `Read a 'hyperflex.ClusterStoragePolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexClusterStoragePolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configresult",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexConfigResultList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ConfigResult' resource.",
										Long: `Read a 'hyperflex.ConfigResult' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexConfigResultByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configresultentry",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexConfigResultEntryList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ConfigResultEntry' resource.",
										Long: `Read a 'hyperflex.ConfigResultEntry' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexConfigResultEntryByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "extfcstoragepolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexExtFcStoragePolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ExtFcStoragePolicy' resource.",
										Long: `Read a 'hyperflex.ExtFcStoragePolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexExtFcStoragePolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "extiscsistoragepolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexExtIscsiStoragePolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ExtIscsiStoragePolicy' resource.",
										Long: `Read a 'hyperflex.ExtIscsiStoragePolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexExtIscsiStoragePolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "featurelimitexternal",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexFeatureLimitExternalList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.FeatureLimitExternal' resource.",
										Long: `Read a 'hyperflex.FeatureLimitExternal' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexFeatureLimitExternalByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "featurelimitinternal",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexFeatureLimitInternalList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.FeatureLimitInternal' resource.",
										Long: `Read a 'hyperflex.FeatureLimitInternal' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexFeatureLimitInternalByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "health",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexHealthList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.Health' resource.",
										Long: `Read a 'hyperflex.Health' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexHealthByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hxdpversion",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexHxdpVersionList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.HxdpVersion' resource.",
										Long: `Read a 'hyperflex.HxdpVersion' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexHxdpVersionByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "localcredentialpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexLocalCredentialPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.LocalCredentialPolicy' resource.",
										Long: `Read a 'hyperflex.LocalCredentialPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexLocalCredentialPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "node",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexNodeList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.Node' resource.",
										Long: `Read a 'hyperflex.Node' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexNodeByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "nodeconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexNodeConfigPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.NodeConfigPolicy' resource.",
										Long: `Read a 'hyperflex.NodeConfigPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexNodeConfigPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "nodeprofile",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexNodeProfileList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.NodeProfile' resource.",
										Long: `Read a 'hyperflex.NodeProfile' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexNodeProfileByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "proxysettingpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexProxySettingPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ProxySettingPolicy' resource.",
										Long: `Read a 'hyperflex.ProxySettingPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexProxySettingPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "serverfirmwareversion",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexServerFirmwareVersionList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ServerFirmwareVersion' resource.",
										Long: `Read a 'hyperflex.ServerFirmwareVersion' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexServerFirmwareVersionByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "servermodel",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexServerModelList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ServerModel' resource.",
										Long: `Read a 'hyperflex.ServerModel' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexServerModelByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "softwareversionpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexSoftwareVersionPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.SoftwareVersionPolicy' resource.",
										Long: `Read a 'hyperflex.SoftwareVersionPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexSoftwareVersionPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "sysconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexSysConfigPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.SysConfigPolicy' resource.",
										Long: `Read a 'hyperflex.SysConfigPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexSysConfigPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ucsmconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexUcsmConfigPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.UcsmConfigPolicy' resource.",
										Long: `Read a 'hyperflex.UcsmConfigPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexUcsmConfigPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vcenterconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.HyperflexApi.GetHyperflexVcenterConfigPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.VcenterConfigPolicy' resource.",
										Long: `Read a 'hyperflex.VcenterConfigPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.GetHyperflexVcenterConfigPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "iaas",

								Short: "Get or list Iaas resource(s)",
								Long: `Get or list Iaas resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "connectorpack",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IaasApi.GetIaasConnectorPackList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iaas.ConnectorPack' resource.",
										Long: `Read a 'iaas.ConnectorPack' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IaasApi.GetIaasConnectorPackByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "devicestatus",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IaasApi.GetIaasDeviceStatusList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iaas.DeviceStatus' resource.",
										Long: `Read a 'iaas.DeviceStatus' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IaasApi.GetIaasDeviceStatusByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "licenseinfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IaasApi.GetIaasLicenseInfoList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iaas.LicenseInfo' resource.",
										Long: `Read a 'iaas.LicenseInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IaasApi.GetIaasLicenseInfoByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "mostruntasks",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IaasApi.GetIaasMostRunTasksList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iaas.MostRunTasks' resource.",
										Long: `Read a 'iaas.MostRunTasks' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IaasApi.GetIaasMostRunTasksByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ucsdinfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IaasApi.GetIaasUcsdInfoList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iaas.UcsdInfo' resource.",
										Long: `Read a 'iaas.UcsdInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IaasApi.GetIaasUcsdInfoByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ucsdmanagedinfra",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IaasApi.GetIaasUcsdManagedInfraList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iaas.UcsdManagedInfra' resource.",
										Long: `Read a 'iaas.UcsdManagedInfra' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IaasApi.GetIaasUcsdManagedInfraByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "iam",

								Short: "Get or list Iam resource(s)",
								Long: `Get or list Iam resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "account",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamAccountList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.Account' resource.",
										Long: `Read a 'iam.Account' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamAccountByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "apikey",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamApiKeyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.ApiKey' resource.",
										Long: `Read a 'iam.ApiKey' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamApiKeyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "appregistration",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamAppRegistrationList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.AppRegistration' resource.",
										Long: `Read a 'iam.AppRegistration' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamAppRegistrationByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "certificate",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamCertificateList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.Certificate' resource.",
										Long: `Read a 'iam.Certificate' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamCertificateByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "certificaterequest",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamCertificateRequestList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.CertificateRequest' resource.",
										Long: `Read a 'iam.CertificateRequest' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamCertificateRequestByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "domaingroup",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamDomainGroupList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.DomainGroup' resource.",
										Long: `Read a 'iam.DomainGroup' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamDomainGroupByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "endpointprivilege",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamEndPointPrivilegeList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.EndPointPrivilege' resource.",
										Long: `Read a 'iam.EndPointPrivilege' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamEndPointPrivilegeByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "endpointrole",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamEndPointRoleList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.EndPointRole' resource.",
										Long: `Read a 'iam.EndPointRole' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamEndPointRoleByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "endpointuser",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamEndPointUserList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.EndPointUser' resource.",
										Long: `Read a 'iam.EndPointUser' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamEndPointUserByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "endpointuserpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamEndPointUserPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.EndPointUserPolicy' resource.",
										Long: `Read a 'iam.EndPointUserPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamEndPointUserPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "endpointuserrole",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamEndPointUserRoleList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.EndPointUserRole' resource.",
										Long: `Read a 'iam.EndPointUserRole' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamEndPointUserRoleByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "idp",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamIdpList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.Idp' resource.",
										Long: `Read a 'iam.Idp' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamIdpByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "idpreference",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamIdpReferenceList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.IdpReference' resource.",
										Long: `Read a 'iam.IdpReference' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamIdpReferenceByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ldapgroup",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamLdapGroupList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.LdapGroup' resource.",
										Long: `Read a 'iam.LdapGroup' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamLdapGroupByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ldappolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamLdapPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.LdapPolicy' resource.",
										Long: `Read a 'iam.LdapPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamLdapPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ldapprovider",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamLdapProviderList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.LdapProvider' resource.",
										Long: `Read a 'iam.LdapProvider' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamLdapProviderByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "oauthtoken",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamOAuthTokenList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.OAuthToken' resource.",
										Long: `Read a 'iam.OAuthToken' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamOAuthTokenByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "permission",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamPermissionList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.Permission' resource.",
										Long: `Read a 'iam.Permission' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamPermissionByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "privatekeyspec",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamPrivateKeySpecList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.PrivateKeySpec' resource.",
										Long: `Read a 'iam.PrivateKeySpec' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamPrivateKeySpecByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "privilege",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamPrivilegeList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.Privilege' resource.",
										Long: `Read a 'iam.Privilege' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamPrivilegeByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "privilegeset",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamPrivilegeSetList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.PrivilegeSet' resource.",
										Long: `Read a 'iam.PrivilegeSet' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamPrivilegeSetByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "qualifier",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamQualifierList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.Qualifier' resource.",
										Long: `Read a 'iam.Qualifier' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamQualifierByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "resourcelimits",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamResourceLimitsList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.ResourceLimits' resource.",
										Long: `Read a 'iam.ResourceLimits' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamResourceLimitsByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "resourcepermission",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamResourcePermissionList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.ResourcePermission' resource.",
										Long: `Read a 'iam.ResourcePermission' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamResourcePermissionByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "resourceroles",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamResourceRolesList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.ResourceRoles' resource.",
										Long: `Read a 'iam.ResourceRoles' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamResourceRolesByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "role",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamRoleList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.Role' resource.",
										Long: `Read a 'iam.Role' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamRoleByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "securityholder",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamSecurityHolderList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.SecurityHolder' resource.",
										Long: `Read a 'iam.SecurityHolder' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamSecurityHolderByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "serviceprovider",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamServiceProviderList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.ServiceProvider' resource.",
										Long: `Read a 'iam.ServiceProvider' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamServiceProviderByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "session",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamSessionList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.Session' resource.",
										Long: `Read a 'iam.Session' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamSessionByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "sessionlimits",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamSessionLimitsList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.SessionLimits' resource.",
										Long: `Read a 'iam.SessionLimits' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamSessionLimitsByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "system",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamSystemList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.System' resource.",
										Long: `Read a 'iam.System' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamSystemByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "trustpoint",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamTrustPointList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.TrustPoint' resource.",
										Long: `Read a 'iam.TrustPoint' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamTrustPointByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "user",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamUserList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.User' resource.",
										Long: `Read a 'iam.User' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamUserByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "usergroup",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamUserGroupList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.UserGroup' resource.",
										Long: `Read a 'iam.UserGroup' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamUserGroupByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "userpreference",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IamApi.GetIamUserPreferenceList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.UserPreference' resource.",
										Long: `Read a 'iam.UserPreference' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.GetIamUserPreferenceByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "infra",

								Short: "Get or list Infra resource(s)",
								Long: `Get or list Infra resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "accountexperience",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.InfraApi.GetInfraAccountExperienceList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'infra.AccountExperience' resource.",
										Long: `Read a 'infra.AccountExperience' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.InfraApi.GetInfraAccountExperienceByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "inventory",

								Short: "Get or list Inventory resource(s)",
								Long: `Get or list Inventory resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deviceinfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.InventoryApi.GetInventoryDeviceInfoList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'inventory.DeviceInfo' resource.",
										Long: `Read a 'inventory.DeviceInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.InventoryApi.GetInventoryDeviceInfoByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "dnmobinding",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.InventoryApi.GetInventoryDnMoBindingList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'inventory.DnMoBinding' resource.",
										Long: `Read a 'inventory.DnMoBinding' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.InventoryApi.GetInventoryDnMoBindingByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "genericinventory",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.InventoryApi.GetInventoryGenericInventoryList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'inventory.GenericInventory' resource.",
										Long: `Read a 'inventory.GenericInventory' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.InventoryApi.GetInventoryGenericInventoryByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "genericinventoryholder",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.InventoryApi.GetInventoryGenericInventoryHolderList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'inventory.GenericInventoryHolder' resource.",
										Long: `Read a 'inventory.GenericInventoryHolder' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.InventoryApi.GetInventoryGenericInventoryHolderByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ipmioverlan",

								Short: "Get or list Ipmioverlan resource(s)",
								Long: `Get or list Ipmioverlan resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.IpmioverlanApi.GetIpmioverlanPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'ipmioverlan.Policy' resource.",
										Long: `Read a 'ipmioverlan.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IpmioverlanApi.GetIpmioverlanPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "kvm",

								Short: "Get or list Kvm resource(s)",
								Long: `Get or list Kvm resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "kvmsession",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.KvmApi.GetKvmKvmSessionList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'kvm.KvmSession' resource.",
										Long: `Read a 'kvm.KvmSession' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.KvmApi.GetKvmKvmSessionByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.KvmApi.GetKvmPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'kvm.Policy' resource.",
										Long: `Read a 'kvm.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.KvmApi.GetKvmPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "license",

								Short: "Get or list License resource(s)",
								Long: `Get or list License resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "accountlicensedata",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.LicenseApi.GetLicenseAccountLicenseDataList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'license.AccountLicenseData' resource.",
										Long: `Read a 'license.AccountLicenseData' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.LicenseApi.GetLicenseAccountLicenseDataByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "customerop",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.LicenseApi.GetLicenseCustomerOpList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'license.CustomerOp' resource.",
										Long: `Read a 'license.CustomerOp' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.LicenseApi.GetLicenseCustomerOpByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "licenseinfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.LicenseApi.GetLicenseLicenseInfoList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'license.LicenseInfo' resource.",
										Long: `Read a 'license.LicenseInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.LicenseApi.GetLicenseLicenseInfoByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "smartlicensetoken",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.LicenseApi.GetLicenseSmartlicenseTokenList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'license.SmartlicenseToken' resource.",
										Long: `Read a 'license.SmartlicenseToken' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.LicenseApi.GetLicenseSmartlicenseTokenByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ls",

								Short: "Get or list Ls resource(s)",
								Long: `Get or list Ls resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "serviceprofile",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.LsApi.GetLsServiceProfileList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'ls.ServiceProfile' resource.",
										Long: `Read a 'ls.ServiceProfile' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.LsApi.GetLsServiceProfileByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "management",

								Short: "Get or list Management resource(s)",
								Long: `Get or list Management resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "controller",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ManagementApi.GetManagementControllerList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'management.Controller' resource.",
										Long: `Read a 'management.Controller' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ManagementApi.GetManagementControllerByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "entity",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ManagementApi.GetManagementEntityList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'management.Entity' resource.",
										Long: `Read a 'management.Entity' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ManagementApi.GetManagementEntityByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "interface",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ManagementApi.GetManagementInterfaceList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'management.Interface' resource.",
										Long: `Read a 'management.Interface' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ManagementApi.GetManagementInterfaceByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "memory",

								Short: "Get or list Memory resource(s)",
								Long: `Get or list Memory resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "array",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.MemoryApi.GetMemoryArrayList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'memory.Array' resource.",
										Long: `Read a 'memory.Array' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.GetMemoryArrayByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "persistentmemoryconfigresult",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.MemoryApi.GetMemoryPersistentMemoryConfigResultList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'memory.PersistentMemoryConfigResult' resource.",
										Long: `Read a 'memory.PersistentMemoryConfigResult' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.GetMemoryPersistentMemoryConfigResultByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "persistentmemoryconfiguration",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.MemoryApi.GetMemoryPersistentMemoryConfigurationList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'memory.PersistentMemoryConfiguration' resource.",
										Long: `Read a 'memory.PersistentMemoryConfiguration' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.GetMemoryPersistentMemoryConfigurationByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "persistentmemorynamespace",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.MemoryApi.GetMemoryPersistentMemoryNamespaceList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'memory.PersistentMemoryNamespace' resource.",
										Long: `Read a 'memory.PersistentMemoryNamespace' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.GetMemoryPersistentMemoryNamespaceByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "persistentmemorynamespaceconfigresult",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.MemoryApi.GetMemoryPersistentMemoryNamespaceConfigResultList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'memory.PersistentMemoryNamespaceConfigResult' resource.",
										Long: `Read a 'memory.PersistentMemoryNamespaceConfigResult' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.GetMemoryPersistentMemoryNamespaceConfigResultByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "persistentmemorypolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.MemoryApi.GetMemoryPersistentMemoryPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'memory.PersistentMemoryPolicy' resource.",
										Long: `Read a 'memory.PersistentMemoryPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.GetMemoryPersistentMemoryPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "persistentmemoryregion",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.MemoryApi.GetMemoryPersistentMemoryRegionList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'memory.PersistentMemoryRegion' resource.",
										Long: `Read a 'memory.PersistentMemoryRegion' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.GetMemoryPersistentMemoryRegionByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "persistentmemoryunit",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.MemoryApi.GetMemoryPersistentMemoryUnitList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'memory.PersistentMemoryUnit' resource.",
										Long: `Read a 'memory.PersistentMemoryUnit' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.GetMemoryPersistentMemoryUnitByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "unit",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.MemoryApi.GetMemoryUnitList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'memory.Unit' resource.",
										Long: `Read a 'memory.Unit' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.GetMemoryUnitByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "meta",

								Short: "Get or list Meta resource(s)",
								Long: `Get or list Meta resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "definition",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.MetaApi.GetMetaDefinitionList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'meta.Definition' resource.",
										Long: `Read a 'meta.Definition' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MetaApi.GetMetaDefinitionByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "network",

								Short: "Get or list Network resource(s)",
								Long: `Get or list Network resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "element",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NetworkApi.GetNetworkElementList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'network.Element' resource.",
										Long: `Read a 'network.Element' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NetworkApi.GetNetworkElementByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "elementsummary",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NetworkApi.GetNetworkElementSummaryList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'network.ElementSummary' resource.",
										Long: `Read a 'network.ElementSummary' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NetworkApi.GetNetworkElementSummaryByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "networkconfig",

								Short: "Get or list Networkconfig resource(s)",
								Long: `Get or list Networkconfig resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NetworkconfigApi.GetNetworkconfigPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'networkconfig.Policy' resource.",
										Long: `Read a 'networkconfig.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NetworkconfigApi.GetNetworkconfigPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "niaapi",

								Short: "Get or list Niaapi resource(s)",
								Long: `Get or list Niaapi resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "apicccopost",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NiaapiApi.GetNiaapiApicCcoPostList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.ApicCcoPost' resource.",
										Long: `Read a 'niaapi.ApicCcoPost' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NiaapiApi.GetNiaapiApicCcoPostByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "apicfieldnotice",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NiaapiApi.GetNiaapiApicFieldNoticeList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.ApicFieldNotice' resource.",
										Long: `Read a 'niaapi.ApicFieldNotice' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NiaapiApi.GetNiaapiApicFieldNoticeByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "apichweol",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NiaapiApi.GetNiaapiApicHweolList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.ApicHweol' resource.",
										Long: `Read a 'niaapi.ApicHweol' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NiaapiApi.GetNiaapiApicHweolByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "apiclatestmaintainedrelease",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NiaapiApi.GetNiaapiApicLatestMaintainedReleaseList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.ApicLatestMaintainedRelease' resource.",
										Long: `Read a 'niaapi.ApicLatestMaintainedRelease' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NiaapiApi.GetNiaapiApicLatestMaintainedReleaseByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "apicreleaserecommend",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NiaapiApi.GetNiaapiApicReleaseRecommendList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.ApicReleaseRecommend' resource.",
										Long: `Read a 'niaapi.ApicReleaseRecommend' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NiaapiApi.GetNiaapiApicReleaseRecommendByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "apicsweol",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NiaapiApi.GetNiaapiApicSweolList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.ApicSweol' resource.",
										Long: `Read a 'niaapi.ApicSweol' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NiaapiApi.GetNiaapiApicSweolByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "dcnmccopost",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NiaapiApi.GetNiaapiDcnmCcoPostList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.DcnmCcoPost' resource.",
										Long: `Read a 'niaapi.DcnmCcoPost' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NiaapiApi.GetNiaapiDcnmCcoPostByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "dcnmfieldnotice",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NiaapiApi.GetNiaapiDcnmFieldNoticeList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.DcnmFieldNotice' resource.",
										Long: `Read a 'niaapi.DcnmFieldNotice' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NiaapiApi.GetNiaapiDcnmFieldNoticeByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "dcnmhweol",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NiaapiApi.GetNiaapiDcnmHweolList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.DcnmHweol' resource.",
										Long: `Read a 'niaapi.DcnmHweol' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NiaapiApi.GetNiaapiDcnmHweolByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "dcnmlatestmaintainedrelease",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NiaapiApi.GetNiaapiDcnmLatestMaintainedReleaseList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.DcnmLatestMaintainedRelease' resource.",
										Long: `Read a 'niaapi.DcnmLatestMaintainedRelease' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NiaapiApi.GetNiaapiDcnmLatestMaintainedReleaseByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "dcnmreleaserecommend",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NiaapiApi.GetNiaapiDcnmReleaseRecommendList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.DcnmReleaseRecommend' resource.",
										Long: `Read a 'niaapi.DcnmReleaseRecommend' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NiaapiApi.GetNiaapiDcnmReleaseRecommendByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "dcnmsweol",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NiaapiApi.GetNiaapiDcnmSweolList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.DcnmSweol' resource.",
										Long: `Read a 'niaapi.DcnmSweol' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NiaapiApi.GetNiaapiDcnmSweolByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "filedownloader",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NiaapiApi.GetNiaapiFileDownloaderList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.FileDownloader' resource.",
										Long: `Read a 'niaapi.FileDownloader' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NiaapiApi.GetNiaapiFileDownloaderByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "niametadata",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NiaapiApi.GetNiaapiNiaMetadataList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.NiaMetadata' resource.",
										Long: `Read a 'niaapi.NiaMetadata' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NiaapiApi.GetNiaapiNiaMetadataByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "versionregex",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NiaapiApi.GetNiaapiVersionRegexList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.VersionRegex' resource.",
										Long: `Read a 'niaapi.VersionRegex' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NiaapiApi.GetNiaapiVersionRegexByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "niatelemetry",

								Short: "Get or list Niatelemetry resource(s)",
								Long: `Get or list Niatelemetry resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "niainventory",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NiatelemetryApi.GetNiatelemetryNiaInventoryList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niatelemetry.NiaInventory' resource.",
										Long: `Read a 'niatelemetry.NiaInventory' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NiatelemetryApi.GetNiatelemetryNiaInventoryByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "nialicensestate",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NiatelemetryApi.GetNiatelemetryNiaLicenseStateList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niatelemetry.NiaLicenseState' resource.",
										Long: `Read a 'niatelemetry.NiaLicenseState' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NiatelemetryApi.GetNiatelemetryNiaLicenseStateByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ntp",

								Short: "Get or list Ntp resource(s)",
								Long: `Get or list Ntp resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.NtpApi.GetNtpPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'ntp.Policy' resource.",
										Long: `Read a 'ntp.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NtpApi.GetNtpPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "organization",

								Short: "Get or list Organization resource(s)",
								Long: `Get or list Organization resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "organization",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.OrganizationApi.GetOrganizationOrganizationList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'organization.Organization' resource.",
										Long: `Read a 'organization.Organization' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.OrganizationApi.GetOrganizationOrganizationByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "os",

								Short: "Get or list Os resource(s)",
								Long: `Get or list Os resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "catalog",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.OsApi.GetOsCatalogList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'os.Catalog' resource.",
										Long: `Read a 'os.Catalog' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.OsApi.GetOsCatalogByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configurationfile",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.OsApi.GetOsConfigurationFileList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'os.ConfigurationFile' resource.",
										Long: `Read a 'os.ConfigurationFile' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.OsApi.GetOsConfigurationFileByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "install",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.OsApi.GetOsInstallList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'os.Install' resource.",
										Long: `Read a 'os.Install' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.OsApi.GetOsInstallByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "pci",

								Short: "Get or list Pci resource(s)",
								Long: `Get or list Pci resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "coprocessorcard",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.PciApi.GetPciCoprocessorCardList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'pci.CoprocessorCard' resource.",
										Long: `Read a 'pci.CoprocessorCard' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.PciApi.GetPciCoprocessorCardByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "device",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.PciApi.GetPciDeviceList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'pci.Device' resource.",
										Long: `Read a 'pci.Device' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.PciApi.GetPciDeviceByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "link",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.PciApi.GetPciLinkList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'pci.Link' resource.",
										Long: `Read a 'pci.Link' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.PciApi.GetPciLinkByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "switch",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.PciApi.GetPciSwitchList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'pci.Switch' resource.",
										Long: `Read a 'pci.Switch' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.PciApi.GetPciSwitchByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "port",

								Short: "Get or list Port resource(s)",
								Long: `Get or list Port resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "group",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.PortApi.GetPortGroupList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'port.Group' resource.",
										Long: `Read a 'port.Group' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.PortApi.GetPortGroupByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "subgroup",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.PortApi.GetPortSubGroupList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'port.SubGroup' resource.",
										Long: `Read a 'port.SubGroup' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.PortApi.GetPortSubGroupByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "processor",

								Short: "Get or list Processor resource(s)",
								Long: `Get or list Processor resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "unit",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ProcessorApi.GetProcessorUnitList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'processor.Unit' resource.",
										Long: `Read a 'processor.Unit' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ProcessorApi.GetProcessorUnitByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "recovery",

								Short: "Get or list Recovery resource(s)",
								Long: `Get or list Recovery resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backupconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.RecoveryApi.GetRecoveryBackupConfigPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'recovery.BackupConfigPolicy' resource.",
										Long: `Read a 'recovery.BackupConfigPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.GetRecoveryBackupConfigPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backupprofile",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.RecoveryApi.GetRecoveryBackupProfileList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'recovery.BackupProfile' resource.",
										Long: `Read a 'recovery.BackupProfile' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.GetRecoveryBackupProfileByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configresult",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.RecoveryApi.GetRecoveryConfigResultList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'recovery.ConfigResult' resource.",
										Long: `Read a 'recovery.ConfigResult' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.GetRecoveryConfigResultByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configresultentry",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.RecoveryApi.GetRecoveryConfigResultEntryList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'recovery.ConfigResultEntry' resource.",
										Long: `Read a 'recovery.ConfigResultEntry' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.GetRecoveryConfigResultEntryByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ondemandbackup",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.RecoveryApi.GetRecoveryOnDemandBackupList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'recovery.OnDemandBackup' resource.",
										Long: `Read a 'recovery.OnDemandBackup' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.GetRecoveryOnDemandBackupByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "restore",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.RecoveryApi.GetRecoveryRestoreList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'recovery.Restore' resource.",
										Long: `Read a 'recovery.Restore' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.GetRecoveryRestoreByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "scheduleconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.RecoveryApi.GetRecoveryScheduleConfigPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'recovery.ScheduleConfigPolicy' resource.",
										Long: `Read a 'recovery.ScheduleConfigPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.GetRecoveryScheduleConfigPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "resource",

								Short: "Get or list Resource resource(s)",
								Long: `Get or list Resource resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "group",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ResourceApi.GetResourceGroupList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'resource.Group' resource.",
										Long: `Read a 'resource.Group' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ResourceApi.GetResourceGroupByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "groupmember",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ResourceApi.GetResourceGroupMemberList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'resource.GroupMember' resource.",
										Long: `Read a 'resource.GroupMember' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ResourceApi.GetResourceGroupMemberByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "licenseresourcecount",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ResourceApi.GetResourceLicenseResourceCountList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'resource.LicenseResourceCount' resource.",
										Long: `Read a 'resource.LicenseResourceCount' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ResourceApi.GetResourceLicenseResourceCountByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "membership",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ResourceApi.GetResourceMembershipList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'resource.Membership' resource.",
										Long: `Read a 'resource.Membership' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ResourceApi.GetResourceMembershipByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "membershipholder",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ResourceApi.GetResourceMembershipHolderList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'resource.MembershipHolder' resource.",
										Long: `Read a 'resource.MembershipHolder' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ResourceApi.GetResourceMembershipHolderByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "sdcard",

								Short: "Get or list Sdcard resource(s)",
								Long: `Get or list Sdcard resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SdcardApi.GetSdcardPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'sdcard.Policy' resource.",
										Long: `Read a 'sdcard.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdcardApi.GetSdcardPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "sdwan",

								Short: "Get or list Sdwan resource(s)",
								Long: `Get or list Sdwan resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "profile",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SdwanApi.GetSdwanProfileList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'sdwan.Profile' resource.",
										Long: `Read a 'sdwan.Profile' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdwanApi.GetSdwanProfileByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "routernode",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SdwanApi.GetSdwanRouterNodeList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'sdwan.RouterNode' resource.",
										Long: `Read a 'sdwan.RouterNode' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdwanApi.GetSdwanRouterNodeByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "routerpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SdwanApi.GetSdwanRouterPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'sdwan.RouterPolicy' resource.",
										Long: `Read a 'sdwan.RouterPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdwanApi.GetSdwanRouterPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vmanageaccountpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SdwanApi.GetSdwanVmanageAccountPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'sdwan.VmanageAccountPolicy' resource.",
										Long: `Read a 'sdwan.VmanageAccountPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdwanApi.GetSdwanVmanageAccountPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "search",

								Short: "Get or list Search resource(s)",
								Long: `Get or list Search resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "searchitem",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SearchApi.GetSearchSearchItemList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'search.SearchItem' resource.",
										Long: `Read a 'search.SearchItem' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SearchApi.GetSearchSearchItemByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "tagitem",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SearchApi.GetSearchTagItemList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'search.TagItem' resource.",
										Long: `Read a 'search.TagItem' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SearchApi.GetSearchTagItemByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "security",

								Short: "Get or list Security resource(s)",
								Long: `Get or list Security resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "unit",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SecurityApi.GetSecurityUnitList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'security.Unit' resource.",
										Long: `Read a 'security.Unit' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SecurityApi.GetSecurityUnitByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "server",

								Short: "Get or list Server resource(s)",
								Long: `Get or list Server resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configchangedetail",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ServerApi.GetServerConfigChangeDetailList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'server.ConfigChangeDetail' resource.",
										Long: `Read a 'server.ConfigChangeDetail' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ServerApi.GetServerConfigChangeDetailByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configimport",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ServerApi.GetServerConfigImportList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'server.ConfigImport' resource.",
										Long: `Read a 'server.ConfigImport' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ServerApi.GetServerConfigImportByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configresult",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ServerApi.GetServerConfigResultList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'server.ConfigResult' resource.",
										Long: `Read a 'server.ConfigResult' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ServerApi.GetServerConfigResultByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configresultentry",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ServerApi.GetServerConfigResultEntryList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'server.ConfigResultEntry' resource.",
										Long: `Read a 'server.ConfigResultEntry' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ServerApi.GetServerConfigResultEntryByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "profile",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.ServerApi.GetServerProfileList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'server.Profile' resource.",
										Long: `Read a 'server.Profile' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ServerApi.GetServerProfileByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "smtp",

								Short: "Get or list Smtp resource(s)",
								Long: `Get or list Smtp resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SmtpApi.GetSmtpPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'smtp.Policy' resource.",
										Long: `Read a 'smtp.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SmtpApi.GetSmtpPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "snmp",

								Short: "Get or list Snmp resource(s)",
								Long: `Get or list Snmp resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SnmpApi.GetSnmpPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'snmp.Policy' resource.",
										Long: `Read a 'snmp.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SnmpApi.GetSnmpPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "software",

								Short: "Get or list Software resource(s)",
								Long: `Get or list Software resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hclmeta",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SoftwareApi.GetSoftwareHclMetaList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'software.HclMeta' resource.",
										Long: `Read a 'software.HclMeta' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwareApi.GetSoftwareHclMetaByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hyperflexdistributable",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SoftwareApi.GetSoftwareHyperflexDistributableList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'software.HyperflexDistributable' resource.",
										Long: `Read a 'software.HyperflexDistributable' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwareApi.GetSoftwareHyperflexDistributableByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "solutiondistributable",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SoftwareApi.GetSoftwareSolutionDistributableList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'software.SolutionDistributable' resource.",
										Long: `Read a 'software.SolutionDistributable' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwareApi.GetSoftwareSolutionDistributableByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "softwarerepository",

								Short: "Get or list Softwarerepository resource(s)",
								Long: `Get or list Softwarerepository resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "authorization",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SoftwarerepositoryApi.GetSoftwarerepositoryAuthorizationList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'softwarerepository.Authorization' resource.",
										Long: `Read a 'softwarerepository.Authorization' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwarerepositoryApi.GetSoftwarerepositoryAuthorizationByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "catalog",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SoftwarerepositoryApi.GetSoftwarerepositoryCatalogList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'softwarerepository.Catalog' resource.",
										Long: `Read a 'softwarerepository.Catalog' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwarerepositoryApi.GetSoftwarerepositoryCatalogByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "operatingsystemfile",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SoftwarerepositoryApi.GetSoftwarerepositoryOperatingSystemFileList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'softwarerepository.OperatingSystemFile' resource.",
										Long: `Read a 'softwarerepository.OperatingSystemFile' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwarerepositoryApi.GetSoftwarerepositoryOperatingSystemFileByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "sol",

								Short: "Get or list Sol resource(s)",
								Long: `Get or list Sol resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SolApi.GetSolPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'sol.Policy' resource.",
										Long: `Read a 'sol.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SolApi.GetSolPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ssh",

								Short: "Get or list Ssh resource(s)",
								Long: `Get or list Ssh resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SshApi.GetSshPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'ssh.Policy' resource.",
										Long: `Read a 'ssh.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SshApi.GetSshPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "storage",

								Short: "Get or list Storage resource(s)",
								Long: `Get or list Storage resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "controller",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStorageControllerList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.Controller' resource.",
										Long: `Read a 'storage.Controller' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStorageControllerByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "diskgrouppolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStorageDiskGroupPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.DiskGroupPolicy' resource.",
										Long: `Read a 'storage.DiskGroupPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStorageDiskGroupPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "enclosure",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStorageEnclosureList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.Enclosure' resource.",
										Long: `Read a 'storage.Enclosure' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStorageEnclosureByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "enclosuredisk",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStorageEnclosureDiskList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.EnclosureDisk' resource.",
										Long: `Read a 'storage.EnclosureDisk' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStorageEnclosureDiskByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "enclosurediskslotep",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStorageEnclosureDiskSlotEpList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.EnclosureDiskSlotEp' resource.",
										Long: `Read a 'storage.EnclosureDiskSlotEp' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStorageEnclosureDiskSlotEpByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "flexflashcontroller",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStorageFlexFlashControllerList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.FlexFlashController' resource.",
										Long: `Read a 'storage.FlexFlashController' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStorageFlexFlashControllerByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "flexflashcontrollerprops",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStorageFlexFlashControllerPropsList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.FlexFlashControllerProps' resource.",
										Long: `Read a 'storage.FlexFlashControllerProps' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStorageFlexFlashControllerPropsByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "flexflashphysicaldrive",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStorageFlexFlashPhysicalDriveList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.FlexFlashPhysicalDrive' resource.",
										Long: `Read a 'storage.FlexFlashPhysicalDrive' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStorageFlexFlashPhysicalDriveByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "flexflashvirtualdrive",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStorageFlexFlashVirtualDriveList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.FlexFlashVirtualDrive' resource.",
										Long: `Read a 'storage.FlexFlashVirtualDrive' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStorageFlexFlashVirtualDriveByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "flexutilcontroller",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStorageFlexUtilControllerList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.FlexUtilController' resource.",
										Long: `Read a 'storage.FlexUtilController' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStorageFlexUtilControllerByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "flexutilphysicaldrive",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStorageFlexUtilPhysicalDriveList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.FlexUtilPhysicalDrive' resource.",
										Long: `Read a 'storage.FlexUtilPhysicalDrive' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStorageFlexUtilPhysicalDriveByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "flexutilvirtualdrive",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStorageFlexUtilVirtualDriveList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.FlexUtilVirtualDrive' resource.",
										Long: `Read a 'storage.FlexUtilVirtualDrive' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStorageFlexUtilVirtualDriveByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "physicaldisk",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStoragePhysicalDiskList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PhysicalDisk' resource.",
										Long: `Read a 'storage.PhysicalDisk' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStoragePhysicalDiskByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "physicaldiskextension",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStoragePhysicalDiskExtensionList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PhysicalDiskExtension' resource.",
										Long: `Read a 'storage.PhysicalDiskExtension' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStoragePhysicalDiskExtensionByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "physicaldiskusage",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStoragePhysicalDiskUsageList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PhysicalDiskUsage' resource.",
										Long: `Read a 'storage.PhysicalDiskUsage' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStoragePhysicalDiskUsageByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "purearray",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStoragePureArrayList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureArray' resource.",
										Long: `Read a 'storage.PureArray' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStoragePureArrayByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "purecontroller",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStoragePureControllerList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureController' resource.",
										Long: `Read a 'storage.PureController' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStoragePureControllerByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "puredisk",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStoragePureDiskList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureDisk' resource.",
										Long: `Read a 'storage.PureDisk' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStoragePureDiskByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "purehost",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStoragePureHostList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureHost' resource.",
										Long: `Read a 'storage.PureHost' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStoragePureHostByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "purehostgroup",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStoragePureHostGroupList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureHostGroup' resource.",
										Long: `Read a 'storage.PureHostGroup' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStoragePureHostGroupByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "purehostlun",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStoragePureHostLunList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureHostLun' resource.",
										Long: `Read a 'storage.PureHostLun' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStoragePureHostLunByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "pureport",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStoragePurePortList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PurePort' resource.",
										Long: `Read a 'storage.PurePort' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStoragePurePortByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "pureprotectiongroup",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStoragePureProtectionGroupList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureProtectionGroup' resource.",
										Long: `Read a 'storage.PureProtectionGroup' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStoragePureProtectionGroupByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "pureprotectiongroupsnapshot",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStoragePureProtectionGroupSnapshotList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureProtectionGroupSnapshot' resource.",
										Long: `Read a 'storage.PureProtectionGroupSnapshot' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStoragePureProtectionGroupSnapshotByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "purereplicationschedule",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStoragePureReplicationScheduleList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureReplicationSchedule' resource.",
										Long: `Read a 'storage.PureReplicationSchedule' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStoragePureReplicationScheduleByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "puresnapshotschedule",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStoragePureSnapshotScheduleList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureSnapshotSchedule' resource.",
										Long: `Read a 'storage.PureSnapshotSchedule' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStoragePureSnapshotScheduleByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "purevolume",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStoragePureVolumeList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureVolume' resource.",
										Long: `Read a 'storage.PureVolume' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStoragePureVolumeByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "purevolumesnapshot",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStoragePureVolumeSnapshotList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureVolumeSnapshot' resource.",
										Long: `Read a 'storage.PureVolumeSnapshot' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStoragePureVolumeSnapshotByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "sasexpander",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStorageSasExpanderList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.SasExpander' resource.",
										Long: `Read a 'storage.SasExpander' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStorageSasExpanderByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "sasport",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStorageSasPortList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.SasPort' resource.",
										Long: `Read a 'storage.SasPort' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStorageSasPortByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "storagepolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStorageStoragePolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.StoragePolicy' resource.",
										Long: `Read a 'storage.StoragePolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStorageStoragePolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vdmemberep",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStorageVdMemberEpList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.VdMemberEp' resource.",
										Long: `Read a 'storage.VdMemberEp' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStorageVdMemberEpByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "virtualdrive",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStorageVirtualDriveList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.VirtualDrive' resource.",
										Long: `Read a 'storage.VirtualDrive' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStorageVirtualDriveByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "virtualdriveextension",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.StorageApi.GetStorageVirtualDriveExtensionList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.VirtualDriveExtension' resource.",
										Long: `Read a 'storage.VirtualDriveExtension' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.GetStorageVirtualDriveExtensionByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "syslog",

								Short: "Get or list Syslog resource(s)",
								Long: `Get or list Syslog resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.SyslogApi.GetSyslogPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'syslog.Policy' resource.",
										Long: `Read a 'syslog.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SyslogApi.GetSyslogPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "tam",

								Short: "Get or list Tam resource(s)",
								Long: `Get or list Tam resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "advisorycount",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.TamApi.GetTamAdvisoryCountList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'tam.AdvisoryCount' resource.",
										Long: `Read a 'tam.AdvisoryCount' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TamApi.GetTamAdvisoryCountByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "advisoryinfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.TamApi.GetTamAdvisoryInfoList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'tam.AdvisoryInfo' resource.",
										Long: `Read a 'tam.AdvisoryInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TamApi.GetTamAdvisoryInfoByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "advisoryinstance",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.TamApi.GetTamAdvisoryInstanceList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'tam.AdvisoryInstance' resource.",
										Long: `Read a 'tam.AdvisoryInstance' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TamApi.GetTamAdvisoryInstanceByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "securityadvisory",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.TamApi.GetTamSecurityAdvisoryList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'tam.SecurityAdvisory' resource.",
										Long: `Read a 'tam.SecurityAdvisory' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TamApi.GetTamSecurityAdvisoryByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "terminal",

								Short: "Get or list Terminal resource(s)",
								Long: `Get or list Terminal resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "auditlog",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.TerminalApi.GetTerminalAuditLogList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'terminal.AuditLog' resource.",
										Long: `Read a 'terminal.AuditLog' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TerminalApi.GetTerminalAuditLogByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "top",

								Short: "Get or list Top resource(s)",
								Long: `Get or list Top resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "system",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.TopApi.GetTopSystemList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'top.System' resource.",
										Long: `Read a 'top.System' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TopApi.GetTopSystemByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ucsd",

								Short: "Get or list Ucsd resource(s)",
								Long: `Get or list Ucsd resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backupinfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.UcsdApi.GetUcsdBackupInfoList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'ucsd.BackupInfo' resource.",
										Long: `Read a 'ucsd.BackupInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.UcsdApi.GetUcsdBackupInfoByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "virtualization",

								Short: "Get or list Virtualization resource(s)",
								Long: `Get or list Virtualization resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vmwarecluster",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VirtualizationApi.GetVirtualizationVmwareClusterList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'virtualization.VmwareCluster' resource.",
										Long: `Read a 'virtualization.VmwareCluster' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VirtualizationApi.GetVirtualizationVmwareClusterByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vmwaredatacenter",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VirtualizationApi.GetVirtualizationVmwareDatacenterList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'virtualization.VmwareDatacenter' resource.",
										Long: `Read a 'virtualization.VmwareDatacenter' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VirtualizationApi.GetVirtualizationVmwareDatacenterByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vmwaredatastore",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VirtualizationApi.GetVirtualizationVmwareDatastoreList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'virtualization.VmwareDatastore' resource.",
										Long: `Read a 'virtualization.VmwareDatastore' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VirtualizationApi.GetVirtualizationVmwareDatastoreByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vmwarehost",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VirtualizationApi.GetVirtualizationVmwareHostList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'virtualization.VmwareHost' resource.",
										Long: `Read a 'virtualization.VmwareHost' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VirtualizationApi.GetVirtualizationVmwareHostByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vmwarevcenter",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VirtualizationApi.GetVirtualizationVmwareVcenterList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'virtualization.VmwareVcenter' resource.",
										Long: `Read a 'virtualization.VmwareVcenter' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VirtualizationApi.GetVirtualizationVmwareVcenterByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vmwarevirtualmachine",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VirtualizationApi.GetVirtualizationVmwareVirtualMachineList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'virtualization.VmwareVirtualMachine' resource.",
										Long: `Read a 'virtualization.VmwareVirtualMachine' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VirtualizationApi.GetVirtualizationVmwareVirtualMachineByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "vmedia",

								Short: "Get or list Vmedia resource(s)",
								Long: `Get or list Vmedia resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VmediaApi.GetVmediaPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vmedia.Policy' resource.",
										Long: `Read a 'vmedia.Policy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VmediaApi.GetVmediaPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "vnic",

								Short: "Get or list Vnic resource(s)",
								Long: `Get or list Vnic resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethadapterpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.GetVnicEthAdapterPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.EthAdapterPolicy' resource.",
										Long: `Read a 'vnic.EthAdapterPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.GetVnicEthAdapterPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethif",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.GetVnicEthIfList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.EthIf' resource.",
										Long: `Read a 'vnic.EthIf' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.GetVnicEthIfByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethnetworkpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.GetVnicEthNetworkPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.EthNetworkPolicy' resource.",
										Long: `Read a 'vnic.EthNetworkPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.GetVnicEthNetworkPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethqospolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.GetVnicEthQosPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.EthQosPolicy' resource.",
										Long: `Read a 'vnic.EthQosPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.GetVnicEthQosPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fcadapterpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.GetVnicFcAdapterPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.FcAdapterPolicy' resource.",
										Long: `Read a 'vnic.FcAdapterPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.GetVnicFcAdapterPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fcif",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.GetVnicFcIfList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.FcIf' resource.",
										Long: `Read a 'vnic.FcIf' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.GetVnicFcIfByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fcnetworkpolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.GetVnicFcNetworkPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.FcNetworkPolicy' resource.",
										Long: `Read a 'vnic.FcNetworkPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.GetVnicFcNetworkPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fcqospolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.GetVnicFcQosPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.FcQosPolicy' resource.",
										Long: `Read a 'vnic.FcQosPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.GetVnicFcQosPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "lanconnectivitypolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.GetVnicLanConnectivityPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.LanConnectivityPolicy' resource.",
										Long: `Read a 'vnic.LanConnectivityPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.GetVnicLanConnectivityPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "sanconnectivitypolicy",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.VnicApi.GetVnicSanConnectivityPolicyList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.SanConnectivityPolicy' resource.",
										Long: `Read a 'vnic.SanConnectivityPolicy' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.GetVnicSanConnectivityPolicyByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "workflow",

								Short: "Get or list Workflow resource(s)",
								Long: `Get or list Workflow resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "batchapiexecutor",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.WorkflowApi.GetWorkflowBatchApiExecutorList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.BatchApiExecutor' resource.",
										Long: `Read a 'workflow.BatchApiExecutor' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.GetWorkflowBatchApiExecutorByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "buildtaskmeta",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.WorkflowApi.GetWorkflowBuildTaskMetaList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.BuildTaskMeta' resource.",
										Long: `Read a 'workflow.BuildTaskMeta' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.GetWorkflowBuildTaskMetaByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "buildtaskmetaowner",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.WorkflowApi.GetWorkflowBuildTaskMetaOwnerList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.BuildTaskMetaOwner' resource.",
										Long: `Read a 'workflow.BuildTaskMetaOwner' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.GetWorkflowBuildTaskMetaOwnerByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "catalog",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.WorkflowApi.GetWorkflowCatalogList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.Catalog' resource.",
										Long: `Read a 'workflow.Catalog' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.GetWorkflowCatalogByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "customdatatypedefinition",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.WorkflowApi.GetWorkflowCustomDataTypeDefinitionList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.CustomDataTypeDefinition' resource.",
										Long: `Read a 'workflow.CustomDataTypeDefinition' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.GetWorkflowCustomDataTypeDefinitionByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "pendingdynamicworkflowinfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.WorkflowApi.GetWorkflowPendingDynamicWorkflowInfoList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.PendingDynamicWorkflowInfo' resource.",
										Long: `Read a 'workflow.PendingDynamicWorkflowInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.GetWorkflowPendingDynamicWorkflowInfoByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "taskdefinition",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.WorkflowApi.GetWorkflowTaskDefinitionList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.TaskDefinition' resource.",
										Long: `Read a 'workflow.TaskDefinition' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.GetWorkflowTaskDefinitionByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "taskinfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.WorkflowApi.GetWorkflowTaskInfoList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.TaskInfo' resource.",
										Long: `Read a 'workflow.TaskInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.GetWorkflowTaskInfoByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "taskmeta",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.WorkflowApi.GetWorkflowTaskMetaList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.TaskMeta' resource.",
										Long: `Read a 'workflow.TaskMeta' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.GetWorkflowTaskMetaByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "workflowdefinition",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.WorkflowApi.GetWorkflowWorkflowDefinitionList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.WorkflowDefinition' resource.",
										Long: `Read a 'workflow.WorkflowDefinition' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.GetWorkflowWorkflowDefinitionByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "workflowinfo",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.WorkflowApi.GetWorkflowWorkflowInfoList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.WorkflowInfo' resource.",
										Long: `Read a 'workflow.WorkflowInfo' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.GetWorkflowWorkflowInfoByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "workflowmeta",

										Run: func(cmd *cobra.Command, args []string) {
											client.GetConfig().Debug = verbose

											req := client.WorkflowApi.GetWorkflowWorkflowMetaList(authCtx)

											res, httpResponse, err := req.Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.WorkflowMeta' resource.",
										Long: `Read a 'workflow.WorkflowMeta' resource.
`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.GetWorkflowWorkflowMetaByMoid(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					return cmd
				}())

			cmd.AddCommand(
				func() *cobra.Command {
					cmd := &cobra.Command{
						Use: "update",

						Short: "Update resouce(s)",
						Long: `Update resouce(s)
`,
					}

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "adapter",

								Short: "Update Adapter resource(s)",
								Long: `Update Adapter resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configpolicy",

										Short: "Update a 'adapter.ConfigPolicy' resource.",
										Long: `Update a 'adapter.ConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AdapterApi.PatchAdapterConfigPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "appliance",

								Short: "Update Appliance resource(s)",
								Long: `Update Appliance resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backuppolicy",

										Short: "Update a 'appliance.BackupPolicy' resource.",
										Long: `Update a 'appliance.BackupPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.PatchApplianceBackupPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "certificatesetting",

										Short: "Update a 'appliance.CertificateSetting' resource.",
										Long: `Update a 'appliance.CertificateSetting' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.PatchApplianceCertificateSetting(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "dataexportpolicy",

										Short: "Update a 'appliance.DataExportPolicy' resource.",
										Long: `Update a 'appliance.DataExportPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.PatchApplianceDataExportPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "diagsetting",

										Short: "Update a 'appliance.DiagSetting' resource.",
										Long: `Update a 'appliance.DiagSetting' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.PatchApplianceDiagSetting(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "setupinfo",

										Short: "Update a 'appliance.SetupInfo' resource.",
										Long: `Update a 'appliance.SetupInfo' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.PatchApplianceSetupInfo(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "upgrade",

										Short: "Update a 'appliance.Upgrade' resource.",
										Long: `Update a 'appliance.Upgrade' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.PatchApplianceUpgrade(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "upgradepolicy",

										Short: "Update a 'appliance.UpgradePolicy' resource.",
										Long: `Update a 'appliance.UpgradePolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ApplianceApi.PatchApplianceUpgradePolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "asset",

								Short: "Update Asset resource(s)",
								Long: `Update Asset resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deviceconfiguration",

										Short: "Update a 'asset.DeviceConfiguration' resource.",
										Long: `Update a 'asset.DeviceConfiguration' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AssetApi.PatchAssetDeviceConfiguration(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "devicecontractinformation",

										Short: "Update a 'asset.DeviceContractInformation' resource.",
										Long: `Update a 'asset.DeviceContractInformation' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AssetApi.PatchAssetDeviceContractInformation(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deviceregistration",

										Short: "Updates the resource representing the device connector. For example, this can be used to annotate the device connector resource with user-specified tags.",
										Long: `Updates the resource representing the device connector. For example, this can be used to annotate the device connector resource with user-specified tags.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AssetApi.PatchAssetDeviceRegistration(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "manageddevice",

										Short: "Update a 'asset.ManagedDevice' resource.",
										Long: `Update a 'asset.ManagedDevice' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.AssetApi.PatchAssetManagedDevice(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "bios",

								Short: "Update Bios resource(s)",
								Long: `Update Bios resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "bootmode",

										Short: "Update a 'bios.BootMode' resource.",
										Long: `Update a 'bios.BootMode' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.BiosApi.PatchBiosBootMode(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'bios.Policy' resource.",
										Long: `Update a 'bios.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.BiosApi.PatchBiosPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "unit",

										Short: "Update a 'bios.Unit' resource.",
										Long: `Update a 'bios.Unit' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.BiosApi.PatchBiosUnit(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "boot",

								Short: "Update Boot resource(s)",
								Long: `Update Boot resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "devicebootmode",

										Short: "Update a 'boot.DeviceBootMode' resource.",
										Long: `Update a 'boot.DeviceBootMode' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.BootApi.PatchBootDeviceBootMode(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "precisionpolicy",

										Short: "Update a 'boot.PrecisionPolicy' resource.",
										Long: `Update a 'boot.PrecisionPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.BootApi.PatchBootPrecisionPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "compute",

								Short: "Update Compute resource(s)",
								Long: `Update Compute resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "blade",

										Short: "Update a 'compute.Blade' resource.",
										Long: `Update a 'compute.Blade' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ComputeApi.PatchComputeBlade(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "board",

										Short: "Update a 'compute.Board' resource.",
										Long: `Update a 'compute.Board' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ComputeApi.PatchComputeBoard(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "rackunit",

										Short: "Update a 'compute.RackUnit' resource.",
										Long: `Update a 'compute.RackUnit' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ComputeApi.PatchComputeRackUnit(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "serversetting",

										Short: "Update a 'compute.ServerSetting' resource.",
										Long: `Update a 'compute.ServerSetting' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ComputeApi.PatchComputeServerSetting(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "deviceconnector",

								Short: "Update Deviceconnector resource(s)",
								Long: `Update Deviceconnector resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'deviceconnector.Policy' resource.",
										Long: `Update a 'deviceconnector.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.DeviceconnectorApi.PatchDeviceconnectorPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "equipment",

								Short: "Update Equipment resource(s)",
								Long: `Update Equipment resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "chassis",

										Short: "Update a 'equipment.Chassis' resource.",
										Long: `Update a 'equipment.Chassis' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.PatchEquipmentChassis(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fan",

										Short: "Update a 'equipment.Fan' resource.",
										Long: `Update a 'equipment.Fan' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.PatchEquipmentFan(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fanmodule",

										Short: "Update a 'equipment.FanModule' resource.",
										Long: `Update a 'equipment.FanModule' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.PatchEquipmentFanModule(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fex",

										Short: "Update a 'equipment.Fex' resource.",
										Long: `Update a 'equipment.Fex' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.PatchEquipmentFex(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "iocard",

										Short: "Update a 'equipment.IoCard' resource.",
										Long: `Update a 'equipment.IoCard' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.PatchEquipmentIoCard(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ioexpander",

										Short: "Update a 'equipment.IoExpander' resource.",
										Long: `Update a 'equipment.IoExpander' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.PatchEquipmentIoExpander(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "locatorled",

										Short: "Update a 'equipment.LocatorLed' resource.",
										Long: `Update a 'equipment.LocatorLed' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.PatchEquipmentLocatorLed(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "psu",

										Short: "Update a 'equipment.Psu' resource.",
										Long: `Update a 'equipment.Psu' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.PatchEquipmentPsu(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "rackenclosure",

										Short: "Update a 'equipment.RackEnclosure' resource.",
										Long: `Update a 'equipment.RackEnclosure' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.PatchEquipmentRackEnclosure(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "rackenclosureslot",

										Short: "Update a 'equipment.RackEnclosureSlot' resource.",
										Long: `Update a 'equipment.RackEnclosureSlot' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.PatchEquipmentRackEnclosureSlot(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "sharediomodule",

										Short: "Update a 'equipment.SharedIoModule' resource.",
										Long: `Update a 'equipment.SharedIoModule' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.PatchEquipmentSharedIoModule(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "switchcard",

										Short: "Update a 'equipment.SwitchCard' resource.",
										Long: `Update a 'equipment.SwitchCard' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.PatchEquipmentSwitchCard(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "systemiocontroller",

										Short: "Update a 'equipment.SystemIoController' resource.",
										Long: `Update a 'equipment.SystemIoController' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.PatchEquipmentSystemIoController(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "tpm",

										Short: "Update a 'equipment.Tpm' resource.",
										Long: `Update a 'equipment.Tpm' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EquipmentApi.PatchEquipmentTpm(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ether",

								Short: "Update Ether resource(s)",
								Long: `Update Ether resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "physicalport",

										Short: "Update a 'ether.PhysicalPort' resource.",
										Long: `Update a 'ether.PhysicalPort' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.EtherApi.PatchEtherPhysicalPort(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "externalsite",

								Short: "Update Externalsite resource(s)",
								Long: `Update Externalsite resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "authorization",

										Short: "Update a 'externalsite.Authorization' resource.",
										Long: `Update a 'externalsite.Authorization' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ExternalsiteApi.PatchExternalsiteAuthorization(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "fault",

								Short: "Update Fault resource(s)",
								Long: `Update Fault resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "instance",

										Short: "Update a 'fault.Instance' resource.",
										Long: `Update a 'fault.Instance' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FaultApi.PatchFaultInstance(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "fc",

								Short: "Update Fc resource(s)",
								Long: `Update Fc resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "physicalport",

										Short: "Update a 'fc.PhysicalPort' resource.",
										Long: `Update a 'fc.PhysicalPort' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FcApi.PatchFcPhysicalPort(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "firmware",

								Short: "Update Firmware resource(s)",
								Long: `Update Firmware resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "distributable",

										Short: "Update a 'firmware.Distributable' resource.",
										Long: `Update a 'firmware.Distributable' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.PatchFirmwareDistributable(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "driverdistributable",

										Short: "Update a 'firmware.DriverDistributable' resource.",
										Long: `Update a 'firmware.DriverDistributable' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.PatchFirmwareDriverDistributable(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "runningfirmware",

										Short: "Update a 'firmware.RunningFirmware' resource.",
										Long: `Update a 'firmware.RunningFirmware' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.PatchFirmwareRunningFirmware(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "serverconfigurationutilitydistributable",

										Short: "Update a 'firmware.ServerConfigurationUtilityDistributable' resource.",
										Long: `Update a 'firmware.ServerConfigurationUtilityDistributable' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.FirmwareApi.PatchFirmwareServerConfigurationUtilityDistributable(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "graphics",

								Short: "Update Graphics resource(s)",
								Long: `Update Graphics resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "card",

										Short: "Update a 'graphics.Card' resource.",
										Long: `Update a 'graphics.Card' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.GraphicsApi.PatchGraphicsCard(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "controller",

										Short: "Update a 'graphics.Controller' resource.",
										Long: `Update a 'graphics.Controller' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.GraphicsApi.PatchGraphicsController(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "hcl",

								Short: "Update Hcl resource(s)",
								Long: `Update Hcl resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hyperflexsoftwarecompatibilityinfo",

										Short: "Update a 'hcl.HyperflexSoftwareCompatibilityInfo' resource.",
										Long: `Update a 'hcl.HyperflexSoftwareCompatibilityInfo' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HclApi.PatchHclHyperflexSoftwareCompatibilityInfo(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "hyperflex",

								Short: "Update Hyperflex resource(s)",
								Long: `Update Hyperflex resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "appcatalog",

										Short: "Update a 'hyperflex.AppCatalog' resource.",
										Long: `Update a 'hyperflex.AppCatalog' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexAppCatalog(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "autosupportpolicy",

										Short: "Update a 'hyperflex.AutoSupportPolicy' resource.",
										Long: `Update a 'hyperflex.AutoSupportPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexAutoSupportPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "capabilityinfo",

										Short: "Update a 'hyperflex.CapabilityInfo' resource.",
										Long: `Update a 'hyperflex.CapabilityInfo' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexCapabilityInfo(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "cluster",

										Short: "Update a 'hyperflex.Cluster' resource.",
										Long: `Update a 'hyperflex.Cluster' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexCluster(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "clusternetworkpolicy",

										Short: "Update a 'hyperflex.ClusterNetworkPolicy' resource.",
										Long: `Update a 'hyperflex.ClusterNetworkPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexClusterNetworkPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "clusterprofile",

										Short: "Update a 'hyperflex.ClusterProfile' resource.",
										Long: `Update a 'hyperflex.ClusterProfile' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexClusterProfile(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "clusterstoragepolicy",

										Short: "Update a 'hyperflex.ClusterStoragePolicy' resource.",
										Long: `Update a 'hyperflex.ClusterStoragePolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexClusterStoragePolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "extfcstoragepolicy",

										Short: "Update a 'hyperflex.ExtFcStoragePolicy' resource.",
										Long: `Update a 'hyperflex.ExtFcStoragePolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexExtFcStoragePolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "extiscsistoragepolicy",

										Short: "Update a 'hyperflex.ExtIscsiStoragePolicy' resource.",
										Long: `Update a 'hyperflex.ExtIscsiStoragePolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexExtIscsiStoragePolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "featurelimitexternal",

										Short: "Update a 'hyperflex.FeatureLimitExternal' resource.",
										Long: `Update a 'hyperflex.FeatureLimitExternal' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexFeatureLimitExternal(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "featurelimitinternal",

										Short: "Update a 'hyperflex.FeatureLimitInternal' resource.",
										Long: `Update a 'hyperflex.FeatureLimitInternal' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexFeatureLimitInternal(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hxdpversion",

										Short: "Update a 'hyperflex.HxdpVersion' resource.",
										Long: `Update a 'hyperflex.HxdpVersion' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexHxdpVersion(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "localcredentialpolicy",

										Short: "Update a 'hyperflex.LocalCredentialPolicy' resource.",
										Long: `Update a 'hyperflex.LocalCredentialPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexLocalCredentialPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "nodeconfigpolicy",

										Short: "Update a 'hyperflex.NodeConfigPolicy' resource.",
										Long: `Update a 'hyperflex.NodeConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexNodeConfigPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "nodeprofile",

										Short: "Update a 'hyperflex.NodeProfile' resource.",
										Long: `Update a 'hyperflex.NodeProfile' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexNodeProfile(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "proxysettingpolicy",

										Short: "Update a 'hyperflex.ProxySettingPolicy' resource.",
										Long: `Update a 'hyperflex.ProxySettingPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexProxySettingPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "serverfirmwareversion",

										Short: "Update a 'hyperflex.ServerFirmwareVersion' resource.",
										Long: `Update a 'hyperflex.ServerFirmwareVersion' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexServerFirmwareVersion(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "servermodel",

										Short: "Update a 'hyperflex.ServerModel' resource.",
										Long: `Update a 'hyperflex.ServerModel' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexServerModel(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "softwareversionpolicy",

										Short: "Update a 'hyperflex.SoftwareVersionPolicy' resource.",
										Long: `Update a 'hyperflex.SoftwareVersionPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexSoftwareVersionPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "sysconfigpolicy",

										Short: "Update a 'hyperflex.SysConfigPolicy' resource.",
										Long: `Update a 'hyperflex.SysConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexSysConfigPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ucsmconfigpolicy",

										Short: "Update a 'hyperflex.UcsmConfigPolicy' resource.",
										Long: `Update a 'hyperflex.UcsmConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexUcsmConfigPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vcenterconfigpolicy",

										Short: "Update a 'hyperflex.VcenterConfigPolicy' resource.",
										Long: `Update a 'hyperflex.VcenterConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.HyperflexApi.PatchHyperflexVcenterConfigPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "iaas",

								Short: "Update Iaas resource(s)",
								Long: `Update Iaas resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ucsdinfo",

										Short: "Update a 'iaas.UcsdInfo' resource.",
										Long: `Update a 'iaas.UcsdInfo' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IaasApi.PatchIaasUcsdInfo(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "iam",

								Short: "Update Iam resource(s)",
								Long: `Update Iam resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "account",

										Short: "Update a 'iam.Account' resource.",
										Long: `Update a 'iam.Account' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamAccount(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "apikey",

										Short: "Update a 'iam.ApiKey' resource.",
										Long: `Update a 'iam.ApiKey' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamApiKey(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "appregistration",

										Short: "Update a 'iam.AppRegistration' resource.",
										Long: `Update a 'iam.AppRegistration' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamAppRegistration(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "certificate",

										Short: "Update a 'iam.Certificate' resource.",
										Long: `Update a 'iam.Certificate' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamCertificate(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "certificaterequest",

										Short: "Update a 'iam.CertificateRequest' resource.",
										Long: `Update a 'iam.CertificateRequest' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamCertificateRequest(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "endpointuser",

										Short: "Update a 'iam.EndPointUser' resource.",
										Long: `Update a 'iam.EndPointUser' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamEndPointUser(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "endpointuserpolicy",

										Short: "Update a 'iam.EndPointUserPolicy' resource.",
										Long: `Update a 'iam.EndPointUserPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamEndPointUserPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "endpointuserrole",

										Short: "Update a 'iam.EndPointUserRole' resource.",
										Long: `Update a 'iam.EndPointUserRole' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamEndPointUserRole(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "idp",

										Short: "Update a 'iam.Idp' resource.",
										Long: `Update a 'iam.Idp' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamIdp(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "idpreference",

										Short: "Update a 'iam.IdpReference' resource.",
										Long: `Update a 'iam.IdpReference' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamIdpReference(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ldapgroup",

										Short: "Update a 'iam.LdapGroup' resource.",
										Long: `Update a 'iam.LdapGroup' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamLdapGroup(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ldappolicy",

										Short: "Update a 'iam.LdapPolicy' resource.",
										Long: `Update a 'iam.LdapPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamLdapPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ldapprovider",

										Short: "Update a 'iam.LdapProvider' resource.",
										Long: `Update a 'iam.LdapProvider' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamLdapProvider(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "localuserpassword",

										Short: "Update a 'iam.LocalUserPassword' resource.",
										Long: `Update a 'iam.LocalUserPassword' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamLocalUserPassword(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "permission",

										Short: "Update a 'iam.Permission' resource.",
										Long: `Update a 'iam.Permission' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamPermission(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "privatekeyspec",

										Short: "Update a 'iam.PrivateKeySpec' resource.",
										Long: `Update a 'iam.PrivateKeySpec' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamPrivateKeySpec(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "qualifier",

										Short: "Update a 'iam.Qualifier' resource.",
										Long: `Update a 'iam.Qualifier' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamQualifier(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "resourceroles",

										Short: "Update a 'iam.ResourceRoles' resource.",
										Long: `Update a 'iam.ResourceRoles' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamResourceRoles(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "sessionlimits",

										Short: "Update a 'iam.SessionLimits' resource.",
										Long: `Update a 'iam.SessionLimits' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamSessionLimits(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "user",

										Short: "Update a 'iam.User' resource.",
										Long: `Update a 'iam.User' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamUser(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "usergroup",

										Short: "Update a 'iam.UserGroup' resource.",
										Long: `Update a 'iam.UserGroup' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamUserGroup(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "userpreference",

										Short: "Update a 'iam.UserPreference' resource.",
										Long: `Update a 'iam.UserPreference' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IamApi.PatchIamUserPreference(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "infra",

								Short: "Update Infra resource(s)",
								Long: `Update Infra resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "accountexperience",

										Short: "Update a 'infra.AccountExperience' resource.",
										Long: `Update a 'infra.AccountExperience' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.InfraApi.PatchInfraAccountExperience(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "inventory",

								Short: "Update Inventory resource(s)",
								Long: `Update Inventory resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "genericinventory",

										Short: "Update a 'inventory.GenericInventory' resource.",
										Long: `Update a 'inventory.GenericInventory' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.InventoryApi.PatchInventoryGenericInventory(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "genericinventoryholder",

										Short: "Update a 'inventory.GenericInventoryHolder' resource.",
										Long: `Update a 'inventory.GenericInventoryHolder' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.InventoryApi.PatchInventoryGenericInventoryHolder(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ipmioverlan",

								Short: "Update Ipmioverlan resource(s)",
								Long: `Update Ipmioverlan resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'ipmioverlan.Policy' resource.",
										Long: `Update a 'ipmioverlan.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.IpmioverlanApi.PatchIpmioverlanPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "kvm",

								Short: "Update Kvm resource(s)",
								Long: `Update Kvm resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'kvm.Policy' resource.",
										Long: `Update a 'kvm.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.KvmApi.PatchKvmPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "license",

								Short: "Update License resource(s)",
								Long: `Update License resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "accountlicensedata",

										Short: "Update a 'license.AccountLicenseData' resource.",
										Long: `Update a 'license.AccountLicenseData' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.LicenseApi.PatchLicenseAccountLicenseData(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "customerop",

										Short: "Update a 'license.CustomerOp' resource.",
										Long: `Update a 'license.CustomerOp' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.LicenseApi.PatchLicenseCustomerOp(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "licenseinfo",

										Short: "Update a 'license.LicenseInfo' resource.",
										Long: `Update a 'license.LicenseInfo' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.LicenseApi.PatchLicenseLicenseInfo(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "smartlicensetoken",

										Short: "Update a 'license.SmartlicenseToken' resource.",
										Long: `Update a 'license.SmartlicenseToken' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.LicenseApi.PatchLicenseSmartlicenseToken(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ls",

								Short: "Update Ls resource(s)",
								Long: `Update Ls resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "serviceprofile",

										Short: "Update a 'ls.ServiceProfile' resource.",
										Long: `Update a 'ls.ServiceProfile' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.LsApi.PatchLsServiceProfile(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "management",

								Short: "Update Management resource(s)",
								Long: `Update Management resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "controller",

										Short: "Update a 'management.Controller' resource.",
										Long: `Update a 'management.Controller' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ManagementApi.PatchManagementController(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "entity",

										Short: "Update a 'management.Entity' resource.",
										Long: `Update a 'management.Entity' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ManagementApi.PatchManagementEntity(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "interface",

										Short: "Update a 'management.Interface' resource.",
										Long: `Update a 'management.Interface' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ManagementApi.PatchManagementInterface(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "memory",

								Short: "Update Memory resource(s)",
								Long: `Update Memory resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "array",

										Short: "Update a 'memory.Array' resource.",
										Long: `Update a 'memory.Array' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.PatchMemoryArray(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "persistentmemoryconfigresult",

										Short: "Update a 'memory.PersistentMemoryConfigResult' resource.",
										Long: `Update a 'memory.PersistentMemoryConfigResult' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.PatchMemoryPersistentMemoryConfigResult(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "persistentmemoryconfiguration",

										Short: "Update a 'memory.PersistentMemoryConfiguration' resource.",
										Long: `Update a 'memory.PersistentMemoryConfiguration' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.PatchMemoryPersistentMemoryConfiguration(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "persistentmemorynamespace",

										Short: "Update a 'memory.PersistentMemoryNamespace' resource.",
										Long: `Update a 'memory.PersistentMemoryNamespace' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.PatchMemoryPersistentMemoryNamespace(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "persistentmemorynamespaceconfigresult",

										Short: "Update a 'memory.PersistentMemoryNamespaceConfigResult' resource.",
										Long: `Update a 'memory.PersistentMemoryNamespaceConfigResult' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.PatchMemoryPersistentMemoryNamespaceConfigResult(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "persistentmemorypolicy",

										Short: "Update a 'memory.PersistentMemoryPolicy' resource.",
										Long: `Update a 'memory.PersistentMemoryPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.PatchMemoryPersistentMemoryPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "persistentmemoryregion",

										Short: "Update a 'memory.PersistentMemoryRegion' resource.",
										Long: `Update a 'memory.PersistentMemoryRegion' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.PatchMemoryPersistentMemoryRegion(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "persistentmemoryunit",

										Short: "Update a 'memory.PersistentMemoryUnit' resource.",
										Long: `Update a 'memory.PersistentMemoryUnit' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.PatchMemoryPersistentMemoryUnit(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "unit",

										Short: "Update a 'memory.Unit' resource.",
										Long: `Update a 'memory.Unit' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.MemoryApi.PatchMemoryUnit(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "network",

								Short: "Update Network resource(s)",
								Long: `Update Network resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "element",

										Short: "Update a 'network.Element' resource.",
										Long: `Update a 'network.Element' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NetworkApi.PatchNetworkElement(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "networkconfig",

								Short: "Update Networkconfig resource(s)",
								Long: `Update Networkconfig resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'networkconfig.Policy' resource.",
										Long: `Update a 'networkconfig.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NetworkconfigApi.PatchNetworkconfigPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ntp",

								Short: "Update Ntp resource(s)",
								Long: `Update Ntp resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'ntp.Policy' resource.",
										Long: `Update a 'ntp.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.NtpApi.PatchNtpPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "organization",

								Short: "Update Organization resource(s)",
								Long: `Update Organization resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "organization",

										Short: "Update a 'organization.Organization' resource.",
										Long: `Update a 'organization.Organization' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.OrganizationApi.PatchOrganizationOrganization(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "pci",

								Short: "Update Pci resource(s)",
								Long: `Update Pci resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "device",

										Short: "Update a 'pci.Device' resource.",
										Long: `Update a 'pci.Device' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.PciApi.PatchPciDevice(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "link",

										Short: "Update a 'pci.Link' resource.",
										Long: `Update a 'pci.Link' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.PciApi.PatchPciLink(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "switch",

										Short: "Update a 'pci.Switch' resource.",
										Long: `Update a 'pci.Switch' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.PciApi.PatchPciSwitch(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "port",

								Short: "Update Port resource(s)",
								Long: `Update Port resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "group",

										Short: "Update a 'port.Group' resource.",
										Long: `Update a 'port.Group' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.PortApi.PatchPortGroup(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "subgroup",

										Short: "Update a 'port.SubGroup' resource.",
										Long: `Update a 'port.SubGroup' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.PortApi.PatchPortSubGroup(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "processor",

								Short: "Update Processor resource(s)",
								Long: `Update Processor resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "unit",

										Short: "Update a 'processor.Unit' resource.",
										Long: `Update a 'processor.Unit' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ProcessorApi.PatchProcessorUnit(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "recovery",

								Short: "Update Recovery resource(s)",
								Long: `Update Recovery resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backupconfigpolicy",

										Short: "Update a 'recovery.BackupConfigPolicy' resource.",
										Long: `Update a 'recovery.BackupConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.PatchRecoveryBackupConfigPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backupprofile",

										Short: "Update a 'recovery.BackupProfile' resource.",
										Long: `Update a 'recovery.BackupProfile' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.PatchRecoveryBackupProfile(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ondemandbackup",

										Short: "Update a 'recovery.OnDemandBackup' resource.",
										Long: `Update a 'recovery.OnDemandBackup' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.PatchRecoveryOnDemandBackup(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "scheduleconfigpolicy",

										Short: "Update a 'recovery.ScheduleConfigPolicy' resource.",
										Long: `Update a 'recovery.ScheduleConfigPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.RecoveryApi.PatchRecoveryScheduleConfigPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "resource",

								Short: "Update Resource resource(s)",
								Long: `Update Resource resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "group",

										Short: "Update a 'resource.Group' resource.",
										Long: `Update a 'resource.Group' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ResourceApi.PatchResourceGroup(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "sdcard",

								Short: "Update Sdcard resource(s)",
								Long: `Update Sdcard resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'sdcard.Policy' resource.",
										Long: `Update a 'sdcard.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdcardApi.PatchSdcardPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "sdwan",

								Short: "Update Sdwan resource(s)",
								Long: `Update Sdwan resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "profile",

										Short: "Update a 'sdwan.Profile' resource.",
										Long: `Update a 'sdwan.Profile' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdwanApi.PatchSdwanProfile(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "routernode",

										Short: "Update a 'sdwan.RouterNode' resource.",
										Long: `Update a 'sdwan.RouterNode' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdwanApi.PatchSdwanRouterNode(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "routerpolicy",

										Short: "Update a 'sdwan.RouterPolicy' resource.",
										Long: `Update a 'sdwan.RouterPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdwanApi.PatchSdwanRouterPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vmanageaccountpolicy",

										Short: "Update a 'sdwan.VmanageAccountPolicy' resource.",
										Long: `Update a 'sdwan.VmanageAccountPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SdwanApi.PatchSdwanVmanageAccountPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "security",

								Short: "Update Security resource(s)",
								Long: `Update Security resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "unit",

										Short: "Update a 'security.Unit' resource.",
										Long: `Update a 'security.Unit' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SecurityApi.PatchSecurityUnit(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "server",

								Short: "Update Server resource(s)",
								Long: `Update Server resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "profile",

										Short: "Update a 'server.Profile' resource.",
										Long: `Update a 'server.Profile' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.ServerApi.PatchServerProfile(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "smtp",

								Short: "Update Smtp resource(s)",
								Long: `Update Smtp resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'smtp.Policy' resource.",
										Long: `Update a 'smtp.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SmtpApi.PatchSmtpPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "snmp",

								Short: "Update Snmp resource(s)",
								Long: `Update Snmp resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'snmp.Policy' resource.",
										Long: `Update a 'snmp.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SnmpApi.PatchSnmpPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "software",

								Short: "Update Software resource(s)",
								Long: `Update Software resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hclmeta",

										Short: "Update a 'software.HclMeta' resource.",
										Long: `Update a 'software.HclMeta' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwareApi.PatchSoftwareHclMeta(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hyperflexdistributable",

										Short: "Update a 'software.HyperflexDistributable' resource.",
										Long: `Update a 'software.HyperflexDistributable' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwareApi.PatchSoftwareHyperflexDistributable(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "solutiondistributable",

										Short: "Update a 'software.SolutionDistributable' resource.",
										Long: `Update a 'software.SolutionDistributable' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwareApi.PatchSoftwareSolutionDistributable(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "softwarerepository",

								Short: "Update Softwarerepository resource(s)",
								Long: `Update Softwarerepository resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "authorization",

										Short: "Update a 'softwarerepository.Authorization' resource.",
										Long: `Update a 'softwarerepository.Authorization' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwarerepositoryApi.PatchSoftwarerepositoryAuthorization(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "operatingsystemfile",

										Short: "Update a 'softwarerepository.OperatingSystemFile' resource.",
										Long: `Update a 'softwarerepository.OperatingSystemFile' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SoftwarerepositoryApi.PatchSoftwarerepositoryOperatingSystemFile(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "sol",

								Short: "Update Sol resource(s)",
								Long: `Update Sol resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'sol.Policy' resource.",
										Long: `Update a 'sol.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SolApi.PatchSolPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ssh",

								Short: "Update Ssh resource(s)",
								Long: `Update Ssh resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'ssh.Policy' resource.",
										Long: `Update a 'ssh.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SshApi.PatchSshPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "storage",

								Short: "Update Storage resource(s)",
								Long: `Update Storage resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "controller",

										Short: "Update a 'storage.Controller' resource.",
										Long: `Update a 'storage.Controller' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStorageController(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "diskgrouppolicy",

										Short: "Update a 'storage.DiskGroupPolicy' resource.",
										Long: `Update a 'storage.DiskGroupPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStorageDiskGroupPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "enclosure",

										Short: "Update a 'storage.Enclosure' resource.",
										Long: `Update a 'storage.Enclosure' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStorageEnclosure(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "enclosuredisk",

										Short: "Update a 'storage.EnclosureDisk' resource.",
										Long: `Update a 'storage.EnclosureDisk' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStorageEnclosureDisk(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "enclosurediskslotep",

										Short: "Update a 'storage.EnclosureDiskSlotEp' resource.",
										Long: `Update a 'storage.EnclosureDiskSlotEp' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStorageEnclosureDiskSlotEp(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "flexflashcontroller",

										Short: "Update a 'storage.FlexFlashController' resource.",
										Long: `Update a 'storage.FlexFlashController' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStorageFlexFlashController(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "flexflashcontrollerprops",

										Short: "Update a 'storage.FlexFlashControllerProps' resource.",
										Long: `Update a 'storage.FlexFlashControllerProps' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStorageFlexFlashControllerProps(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "flexflashphysicaldrive",

										Short: "Update a 'storage.FlexFlashPhysicalDrive' resource.",
										Long: `Update a 'storage.FlexFlashPhysicalDrive' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStorageFlexFlashPhysicalDrive(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "flexflashvirtualdrive",

										Short: "Update a 'storage.FlexFlashVirtualDrive' resource.",
										Long: `Update a 'storage.FlexFlashVirtualDrive' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStorageFlexFlashVirtualDrive(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "flexutilcontroller",

										Short: "Update a 'storage.FlexUtilController' resource.",
										Long: `Update a 'storage.FlexUtilController' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStorageFlexUtilController(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "flexutilphysicaldrive",

										Short: "Update a 'storage.FlexUtilPhysicalDrive' resource.",
										Long: `Update a 'storage.FlexUtilPhysicalDrive' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStorageFlexUtilPhysicalDrive(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "flexutilvirtualdrive",

										Short: "Update a 'storage.FlexUtilVirtualDrive' resource.",
										Long: `Update a 'storage.FlexUtilVirtualDrive' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStorageFlexUtilVirtualDrive(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "physicaldisk",

										Short: "Update a 'storage.PhysicalDisk' resource.",
										Long: `Update a 'storage.PhysicalDisk' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStoragePhysicalDisk(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "physicaldiskextension",

										Short: "Update a 'storage.PhysicalDiskExtension' resource.",
										Long: `Update a 'storage.PhysicalDiskExtension' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStoragePhysicalDiskExtension(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "physicaldiskusage",

										Short: "Update a 'storage.PhysicalDiskUsage' resource.",
										Long: `Update a 'storage.PhysicalDiskUsage' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStoragePhysicalDiskUsage(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "purearray",

										Short: "Update a 'storage.PureArray' resource.",
										Long: `Update a 'storage.PureArray' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStoragePureArray(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "sasexpander",

										Short: "Update a 'storage.SasExpander' resource.",
										Long: `Update a 'storage.SasExpander' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStorageSasExpander(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "sasport",

										Short: "Update a 'storage.SasPort' resource.",
										Long: `Update a 'storage.SasPort' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStorageSasPort(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "storagepolicy",

										Short: "Update a 'storage.StoragePolicy' resource.",
										Long: `Update a 'storage.StoragePolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStorageStoragePolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vdmemberep",

										Short: "Update a 'storage.VdMemberEp' resource.",
										Long: `Update a 'storage.VdMemberEp' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStorageVdMemberEp(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "virtualdrive",

										Short: "Update a 'storage.VirtualDrive' resource.",
										Long: `Update a 'storage.VirtualDrive' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStorageVirtualDrive(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "virtualdriveextension",

										Short: "Update a 'storage.VirtualDriveExtension' resource.",
										Long: `Update a 'storage.VirtualDriveExtension' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.StorageApi.PatchStorageVirtualDriveExtension(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "syslog",

								Short: "Update Syslog resource(s)",
								Long: `Update Syslog resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'syslog.Policy' resource.",
										Long: `Update a 'syslog.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.SyslogApi.PatchSyslogPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "tam",

								Short: "Update Tam resource(s)",
								Long: `Update Tam resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "advisorycount",

										Short: "Update a 'tam.AdvisoryCount' resource.",
										Long: `Update a 'tam.AdvisoryCount' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TamApi.PatchTamAdvisoryCount(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "advisoryinfo",

										Short: "Update a 'tam.AdvisoryInfo' resource.",
										Long: `Update a 'tam.AdvisoryInfo' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TamApi.PatchTamAdvisoryInfo(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "advisoryinstance",

										Short: "Update a 'tam.AdvisoryInstance' resource.",
										Long: `Update a 'tam.AdvisoryInstance' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TamApi.PatchTamAdvisoryInstance(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "securityadvisory",

										Short: "Update a 'tam.SecurityAdvisory' resource.",
										Long: `Update a 'tam.SecurityAdvisory' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TamApi.PatchTamSecurityAdvisory(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "top",

								Short: "Update Top resource(s)",
								Long: `Update Top resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "system",

										Short: "Update a 'top.System' resource.",
										Long: `Update a 'top.System' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.TopApi.PatchTopSystem(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "virtualization",

								Short: "Update Virtualization resource(s)",
								Long: `Update Virtualization resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vmwarecluster",

										Short: "Update a 'virtualization.VmwareCluster' resource.",
										Long: `Update a 'virtualization.VmwareCluster' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VirtualizationApi.PatchVirtualizationVmwareCluster(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vmwaredatacenter",

										Short: "Update a 'virtualization.VmwareDatacenter' resource.",
										Long: `Update a 'virtualization.VmwareDatacenter' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VirtualizationApi.PatchVirtualizationVmwareDatacenter(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vmwaredatastore",

										Short: "Update a 'virtualization.VmwareDatastore' resource.",
										Long: `Update a 'virtualization.VmwareDatastore' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VirtualizationApi.PatchVirtualizationVmwareDatastore(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vmwarehost",

										Short: "Update a 'virtualization.VmwareHost' resource.",
										Long: `Update a 'virtualization.VmwareHost' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VirtualizationApi.PatchVirtualizationVmwareHost(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vmwarevirtualmachine",

										Short: "Update a 'virtualization.VmwareVirtualMachine' resource.",
										Long: `Update a 'virtualization.VmwareVirtualMachine' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VirtualizationApi.PatchVirtualizationVmwareVirtualMachine(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "vmedia",

								Short: "Update Vmedia resource(s)",
								Long: `Update Vmedia resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'vmedia.Policy' resource.",
										Long: `Update a 'vmedia.Policy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VmediaApi.PatchVmediaPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "vnic",

								Short: "Update Vnic resource(s)",
								Long: `Update Vnic resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethadapterpolicy",

										Short: "Update a 'vnic.EthAdapterPolicy' resource.",
										Long: `Update a 'vnic.EthAdapterPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.PatchVnicEthAdapterPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethif",

										Short: "Update a 'vnic.EthIf' resource.",
										Long: `Update a 'vnic.EthIf' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.PatchVnicEthIf(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethnetworkpolicy",

										Short: "Update a 'vnic.EthNetworkPolicy' resource.",
										Long: `Update a 'vnic.EthNetworkPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.PatchVnicEthNetworkPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethqospolicy",

										Short: "Update a 'vnic.EthQosPolicy' resource.",
										Long: `Update a 'vnic.EthQosPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.PatchVnicEthQosPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fcadapterpolicy",

										Short: "Update a 'vnic.FcAdapterPolicy' resource.",
										Long: `Update a 'vnic.FcAdapterPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.PatchVnicFcAdapterPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fcif",

										Short: "Update a 'vnic.FcIf' resource.",
										Long: `Update a 'vnic.FcIf' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.PatchVnicFcIf(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fcnetworkpolicy",

										Short: "Update a 'vnic.FcNetworkPolicy' resource.",
										Long: `Update a 'vnic.FcNetworkPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.PatchVnicFcNetworkPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "fcqospolicy",

										Short: "Update a 'vnic.FcQosPolicy' resource.",
										Long: `Update a 'vnic.FcQosPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.PatchVnicFcQosPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "lanconnectivitypolicy",

										Short: "Update a 'vnic.LanConnectivityPolicy' resource.",
										Long: `Update a 'vnic.LanConnectivityPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.PatchVnicLanConnectivityPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "sanconnectivitypolicy",

										Short: "Update a 'vnic.SanConnectivityPolicy' resource.",
										Long: `Update a 'vnic.SanConnectivityPolicy' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.VnicApi.PatchVnicSanConnectivityPolicy(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "workflow",

								Short: "Update Workflow resource(s)",
								Long: `Update Workflow resource(s)
`,
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "batchapiexecutor",

										Short: "Update a 'workflow.BatchApiExecutor' resource.",
										Long: `Update a 'workflow.BatchApiExecutor' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.PatchWorkflowBatchApiExecutor(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "customdatatypedefinition",

										Short: "Update a 'workflow.CustomDataTypeDefinition' resource.",
										Long: `Update a 'workflow.CustomDataTypeDefinition' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.PatchWorkflowCustomDataTypeDefinition(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "taskdefinition",

										Short: "Update a 'workflow.TaskDefinition' resource.",
										Long: `Update a 'workflow.TaskDefinition' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.PatchWorkflowTaskDefinition(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "taskinfo",

										Short: "Update a 'workflow.TaskInfo' resource.",
										Long: `Update a 'workflow.TaskInfo' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.PatchWorkflowTaskInfo(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "workflowdefinition",

										Short: "Update a 'workflow.WorkflowDefinition' resource.",
										Long: `Update a 'workflow.WorkflowDefinition' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.PatchWorkflowWorkflowDefinition(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "workflowinfo",

										Short: "Update a 'workflow.WorkflowInfo' resource.",
										Long: `Update a 'workflow.WorkflowInfo' resource.

Provide resource body as JSON on standard input`,
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {
													client.GetConfig().Debug = verbose

													req := client.WorkflowApi.PatchWorkflowWorkflowInfo(authCtx, args[0])

													res, httpResponse, err := req.Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
												Long: `
`,
											}

											return cmd
										}())

									return cmd
								}())

							return cmd
						}())

					return cmd
				}())

			return cmd
		}()
	return rootCmd
}

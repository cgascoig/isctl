package main

import (
	"log"
	"net/http"

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
			}

			cmd.AddCommand(
				func() *cobra.Command {
					cmd := &cobra.Command{
						Use: "create",

						Short: "Create resouce(s)",
					}

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "adapter",

								Short: "Create Adapter resource(s)",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createadapterconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AdapterApi.CreateAdapterConfigPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'adapter.ConfigPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateadapterconfigpolicy",

										Short: "Update a 'adapter.ConfigPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AdapterApi.UpdateAdapterConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createappliancebackup",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.CreateApplianceBackup(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'appliance.Backup' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createappliancebackuppolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.CreateApplianceBackupPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'appliance.BackupPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createappliancedataexportpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.CreateApplianceDataExportPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'appliance.DataExportPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createappliancedeviceclaim",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.CreateApplianceDeviceClaim(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'appliance.DeviceClaim' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createappliancediagsetting",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.CreateApplianceDiagSetting(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'appliance.DiagSetting' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createappliancerestore",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.CreateApplianceRestore(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'appliance.Restore' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateappliancebackuppolicy",

										Short: "Update a 'appliance.BackupPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.UpdateApplianceBackupPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.UpdateApplianceCertificateSetting(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.UpdateApplianceDataExportPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.UpdateApplianceDiagSetting(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.UpdateApplianceSetupInfo(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.UpdateApplianceUpgrade(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.UpdateApplianceUpgradePolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createassetdeviceclaim",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AssetApi.CreateAssetDeviceClaim(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'asset.DeviceClaim' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createassetmanageddevice",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AssetApi.CreateAssetManagedDevice(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'asset.ManagedDevice' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateassetdeviceconfiguration",

										Short: "Update a 'asset.DeviceConfiguration' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AssetApi.UpdateAssetDeviceConfiguration(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AssetApi.UpdateAssetDeviceContractInformation(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AssetApi.UpdateAssetDeviceRegistration(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AssetApi.UpdateAssetManagedDevice(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createbiospolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.BiosApi.CreateBiosPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'bios.Policy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatebiosbootmode",

										Short: "Update a 'bios.BootMode' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.BiosApi.UpdateBiosBootMode(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.BiosApi.UpdateBiosPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.BiosApi.UpdateBiosUnit(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createbootprecisionpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.BootApi.CreateBootPrecisionPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'boot.PrecisionPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatebootdevicebootmode",

										Short: "Update a 'boot.DeviceBootMode' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.BootApi.UpdateBootDeviceBootMode(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.BootApi.UpdateBootPrecisionPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatecomputeblade",

										Short: "Update a 'compute.Blade' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ComputeApi.UpdateComputeBlade(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ComputeApi.UpdateComputeBoard(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ComputeApi.UpdateComputeRackUnit(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ComputeApi.UpdateComputeServerSetting(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createdeviceconnectorpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.DeviceconnectorApi.CreateDeviceconnectorPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'deviceconnector.Policy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatedeviceconnectorpolicy",

										Short: "Update a 'deviceconnector.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.DeviceconnectorApi.UpdateDeviceconnectorPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentchassis",

										Short: "Update a 'equipment.Chassis' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.UpdateEquipmentChassis(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.UpdateEquipmentFan(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.UpdateEquipmentFanModule(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.UpdateEquipmentFex(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.UpdateEquipmentIoCard(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.UpdateEquipmentIoExpander(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.UpdateEquipmentLocatorLed(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.UpdateEquipmentPsu(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.UpdateEquipmentRackEnclosure(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.UpdateEquipmentRackEnclosureSlot(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.UpdateEquipmentSharedIoModule(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.UpdateEquipmentSwitchCard(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.UpdateEquipmentSystemIoController(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.UpdateEquipmentTpm(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateetherphysicalport",

										Short: "Update a 'ether.PhysicalPort' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EtherApi.UpdateEtherPhysicalPort(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createexternalsiteauthorization",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ExternalsiteApi.CreateExternalsiteAuthorization(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'externalsite.Authorization' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateexternalsiteauthorization",

										Short: "Update a 'externalsite.Authorization' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ExternalsiteApi.UpdateExternalsiteAuthorization(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatefaultinstance",

										Short: "Update a 'fault.Instance' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FaultApi.UpdateFaultInstance(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatefcphysicalport",

										Short: "Update a 'fc.PhysicalPort' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FcApi.UpdateFcPhysicalPort(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createfeedbackfeedbackpost",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FeedbackApi.CreateFeedbackFeedbackPost(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'feedback.FeedbackPost' resource.",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createfirmwaredistributable",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FirmwareApi.CreateFirmwareDistributable(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'firmware.Distributable' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createfirmwaredriverdistributable",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FirmwareApi.CreateFirmwareDriverDistributable(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'firmware.DriverDistributable' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createfirmwareeula",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FirmwareApi.CreateFirmwareEula(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'firmware.Eula' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createfirmwareserverconfigurationutilitydistributable",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FirmwareApi.CreateFirmwareServerConfigurationUtilityDistributable(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'firmware.ServerConfigurationUtilityDistributable' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createfirmwareupgrade",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FirmwareApi.CreateFirmwareUpgrade(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'firmware.Upgrade' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatefirmwaredistributable",

										Short: "Update a 'firmware.Distributable' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FirmwareApi.UpdateFirmwareDistributable(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FirmwareApi.UpdateFirmwareDriverDistributable(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FirmwareApi.UpdateFirmwareRunningFirmware(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FirmwareApi.UpdateFirmwareServerConfigurationUtilityDistributable(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

								Short: "Create Graphics resource(s)",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updategraphicscard",

										Short: "Update a 'graphics.Card' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.GraphicsApi.UpdateGraphicsCard(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.GraphicsApi.UpdateGraphicsController(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhclcompatibilitystatus",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HclApi.CreateHclCompatibilityStatus(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hcl.CompatibilityStatus' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhclhyperflexsoftwarecompatibilityinfo",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HclApi.CreateHclHyperflexSoftwareCompatibilityInfo(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hcl.HyperflexSoftwareCompatibilityInfo' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhclsupporteddrivername",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HclApi.CreateHclSupportedDriverName(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hcl.SupportedDriverName' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehclhyperflexsoftwarecompatibilityinfo",

										Short: "Update a 'hcl.HyperflexSoftwareCompatibilityInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HclApi.UpdateHclHyperflexSoftwareCompatibilityInfo(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexappcatalog",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexAppCatalog(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.AppCatalog' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexautosupportpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexAutoSupportPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.AutoSupportPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexcapabilityinfo",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexCapabilityInfo(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.CapabilityInfo' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexclusternetworkpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexClusterNetworkPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.ClusterNetworkPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexclusterprofile",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexClusterProfile(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.ClusterProfile' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexclusterstoragepolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexClusterStoragePolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.ClusterStoragePolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexextfcstoragepolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexExtFcStoragePolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.ExtFcStoragePolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexextiscsistoragepolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexExtIscsiStoragePolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.ExtIscsiStoragePolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexfeaturelimitexternal",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexFeatureLimitExternal(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.FeatureLimitExternal' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexfeaturelimitinternal",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexFeatureLimitInternal(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.FeatureLimitInternal' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexhxdpversion",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexHxdpVersion(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.HxdpVersion' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexlocalcredentialpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexLocalCredentialPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.LocalCredentialPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexnodeconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexNodeConfigPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.NodeConfigPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexnodeprofile",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexNodeProfile(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.NodeProfile' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexproxysettingpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexProxySettingPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.ProxySettingPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexserverfirmwareversion",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexServerFirmwareVersion(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.ServerFirmwareVersion' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexservermodel",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexServerModel(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.ServerModel' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexsoftwareversionpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexSoftwareVersionPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.SoftwareVersionPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexsysconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexSysConfigPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.SysConfigPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexucsmconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexUcsmConfigPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.UcsmConfigPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexvcenterconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexVcenterConfigPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'hyperflex.VcenterConfigPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexappcatalog",

										Short: "Update a 'hyperflex.AppCatalog' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexAppCatalog(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexAutoSupportPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexCapabilityInfo(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexCluster(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexClusterNetworkPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexClusterProfile(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexClusterStoragePolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexExtFcStoragePolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexExtIscsiStoragePolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexFeatureLimitExternal(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexFeatureLimitInternal(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexHxdpVersion(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexLocalCredentialPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexNodeConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexNodeProfile(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexProxySettingPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexServerFirmwareVersion(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexServerModel(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexSoftwareVersionPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexSysConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexUcsmConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.UpdateHyperflexVcenterConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

								Short: "Create Iaas resource(s)",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiaasucsdinfo",

										Short: "Update a 'iaas.UcsdInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IaasApi.UpdateIaasUcsdInfo(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamaccount",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamAccount(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.Account' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamapikey",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamApiKey(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.ApiKey' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamappregistration",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamAppRegistration(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.AppRegistration' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamcertificate",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamCertificate(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.Certificate' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamcertificaterequest",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamCertificateRequest(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.CertificateRequest' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamendpointuser",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamEndPointUser(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.EndPointUser' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamendpointuserpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamEndPointUserPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.EndPointUserPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamendpointuserrole",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamEndPointUserRole(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.EndPointUserRole' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamidp",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamIdp(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.Idp' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamldapgroup",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamLdapGroup(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.LdapGroup' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamldappolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamLdapPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.LdapPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamldapprovider",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamLdapProvider(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.LdapProvider' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiampermission",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamPermission(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.Permission' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamprivatekeyspec",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamPrivateKeySpec(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.PrivateKeySpec' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamqualifier",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamQualifier(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.Qualifier' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamresourceroles",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamResourceRoles(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.ResourceRoles' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamsessionlimits",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamSessionLimits(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.SessionLimits' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamtrustpoint",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamTrustPoint(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.TrustPoint' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamuser",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamUser(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.User' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamusergroup",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamUserGroup(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'iam.UserGroup' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamaccount",

										Short: "Update a 'iam.Account' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamAccount(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamApiKey(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamAppRegistration(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamCertificate(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamCertificateRequest(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamEndPointUser(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamEndPointUserPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamEndPointUserRole(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamIdp(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamIdpReference(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamLdapGroup(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamLdapPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamLdapProvider(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamLocalUserPassword(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamPermission(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamPrivateKeySpec(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamQualifier(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamResourceRoles(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamSessionLimits(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamUser(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamUserGroup(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.UpdateIamUserPreference(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

								Short: "Create Infra resource(s)",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createinfraaccountexperience",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.InfraApi.CreateInfraAccountExperience(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'infra.AccountExperience' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateinfraaccountexperience",

										Short: "Update a 'infra.AccountExperience' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.InfraApi.UpdateInfraAccountExperience(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createinventoryrequest",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.InventoryApi.CreateInventoryRequest(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'inventory.Request' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateinventorygenericinventory",

										Short: "Update a 'inventory.GenericInventory' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.InventoryApi.UpdateInventoryGenericInventory(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.InventoryApi.UpdateInventoryGenericInventoryHolder(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createipmioverlanpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IpmioverlanApi.CreateIpmioverlanPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'ipmioverlan.Policy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateipmioverlanpolicy",

										Short: "Update a 'ipmioverlan.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IpmioverlanApi.UpdateIpmioverlanPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createkvmpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.KvmApi.CreateKvmPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'kvm.Policy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatekvmpolicy",

										Short: "Update a 'kvm.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.KvmApi.UpdateKvmPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createlicenselicenseinfo",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.LicenseApi.CreateLicenseLicenseInfo(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'license.LicenseInfo' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatelicenseaccountlicensedata",

										Short: "Update a 'license.AccountLicenseData' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.LicenseApi.UpdateLicenseAccountLicenseData(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.LicenseApi.UpdateLicenseCustomerOp(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.LicenseApi.UpdateLicenseLicenseInfo(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.LicenseApi.UpdateLicenseSmartlicenseToken(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatelsserviceprofile",

										Short: "Update a 'ls.ServiceProfile' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.LsApi.UpdateLsServiceProfile(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatemanagementcontroller",

										Short: "Update a 'management.Controller' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ManagementApi.UpdateManagementController(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ManagementApi.UpdateManagementEntity(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ManagementApi.UpdateManagementInterface(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "creatememorypersistentmemorypolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.MemoryApi.CreateMemoryPersistentMemoryPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'memory.PersistentMemoryPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememoryarray",

										Short: "Update a 'memory.Array' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.UpdateMemoryArray(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.UpdateMemoryPersistentMemoryConfigResult(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.UpdateMemoryPersistentMemoryConfiguration(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.UpdateMemoryPersistentMemoryNamespace(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.UpdateMemoryPersistentMemoryNamespaceConfigResult(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.UpdateMemoryPersistentMemoryPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.UpdateMemoryPersistentMemoryRegion(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.UpdateMemoryPersistentMemoryUnit(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.UpdateMemoryUnit(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatenetworkelement",

										Short: "Update a 'network.Element' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NetworkApi.UpdateNetworkElement(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createnetworkconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NetworkconfigApi.CreateNetworkconfigPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'networkconfig.Policy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatenetworkconfigpolicy",

										Short: "Update a 'networkconfig.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NetworkconfigApi.UpdateNetworkconfigPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createntppolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NtpApi.CreateNtpPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'ntp.Policy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatentppolicy",

										Short: "Update a 'ntp.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NtpApi.UpdateNtpPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createorganizationorganization",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.OrganizationApi.CreateOrganizationOrganization(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'organization.Organization' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateorganizationorganization",

										Short: "Update a 'organization.Organization' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.OrganizationApi.UpdateOrganizationOrganization(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createosconfigurationfile",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.OsApi.CreateOsConfigurationFile(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'os.ConfigurationFile' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createosinstall",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.OsApi.CreateOsInstall(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'os.Install' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createosossupport",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.OsApi.CreateOsOsSupport(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'os.OsSupport' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createostemplatefile",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.OsApi.CreateOsTemplateFile(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'os.TemplateFile' resource.",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatepcidevice",

										Short: "Update a 'pci.Device' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.PciApi.UpdatePciDevice(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.PciApi.UpdatePciLink(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.PciApi.UpdatePciSwitch(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateportgroup",

										Short: "Update a 'port.Group' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.PortApi.UpdatePortGroup(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.PortApi.UpdatePortSubGroup(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateprocessorunit",

										Short: "Update a 'processor.Unit' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ProcessorApi.UpdateProcessorUnit(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createrecoverybackupconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.RecoveryApi.CreateRecoveryBackupConfigPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'recovery.BackupConfigPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createrecoverybackupprofile",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.RecoveryApi.CreateRecoveryBackupProfile(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'recovery.BackupProfile' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createrecoveryondemandbackup",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.RecoveryApi.CreateRecoveryOnDemandBackup(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'recovery.OnDemandBackup' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createrecoveryrestore",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.RecoveryApi.CreateRecoveryRestore(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'recovery.Restore' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createrecoveryscheduleconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.RecoveryApi.CreateRecoveryScheduleConfigPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'recovery.ScheduleConfigPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updaterecoverybackupconfigpolicy",

										Short: "Update a 'recovery.BackupConfigPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.RecoveryApi.UpdateRecoveryBackupConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.RecoveryApi.UpdateRecoveryBackupProfile(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.RecoveryApi.UpdateRecoveryOnDemandBackup(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.RecoveryApi.UpdateRecoveryScheduleConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createresourcegroup",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ResourceApi.CreateResourceGroup(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'resource.Group' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateresourcegroup",

										Short: "Update a 'resource.Group' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ResourceApi.UpdateResourceGroup(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsdcardpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SdcardApi.CreateSdcardPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'sdcard.Policy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesdcardpolicy",

										Short: "Update a 'sdcard.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SdcardApi.UpdateSdcardPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsdwanprofile",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SdwanApi.CreateSdwanProfile(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'sdwan.Profile' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsdwanrouternode",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SdwanApi.CreateSdwanRouterNode(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'sdwan.RouterNode' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsdwanrouterpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SdwanApi.CreateSdwanRouterPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'sdwan.RouterPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsdwanvmanageaccountpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SdwanApi.CreateSdwanVmanageAccountPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'sdwan.VmanageAccountPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesdwanprofile",

										Short: "Update a 'sdwan.Profile' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SdwanApi.UpdateSdwanProfile(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SdwanApi.UpdateSdwanRouterNode(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SdwanApi.UpdateSdwanRouterPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SdwanApi.UpdateSdwanVmanageAccountPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

								Short: "Create Search resource(s)",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsearchsuggestitem",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SearchApi.CreateSearchSuggestItem(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'search.SuggestItem' resource.",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesecurityunit",

										Short: "Update a 'security.Unit' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SecurityApi.UpdateSecurityUnit(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createserverconfigimport",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ServerApi.CreateServerConfigImport(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'server.ConfigImport' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createserverprofile",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ServerApi.CreateServerProfile(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'server.Profile' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateserverprofile",

										Short: "Update a 'server.Profile' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ServerApi.UpdateServerProfile(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsmtppolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SmtpApi.CreateSmtpPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'smtp.Policy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesmtppolicy",

										Short: "Update a 'smtp.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SmtpApi.UpdateSmtpPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsnmppolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SnmpApi.CreateSnmpPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'snmp.Policy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesnmppolicy",

										Short: "Update a 'snmp.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SnmpApi.UpdateSnmpPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsoftwarehclmeta",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SoftwareApi.CreateSoftwareHclMeta(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'software.HclMeta' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsoftwarehyperflexdistributable",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SoftwareApi.CreateSoftwareHyperflexDistributable(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'software.HyperflexDistributable' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsoftwaresolutiondistributable",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SoftwareApi.CreateSoftwareSolutionDistributable(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'software.SolutionDistributable' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesoftwarehclmeta",

										Short: "Update a 'software.HclMeta' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SoftwareApi.UpdateSoftwareHclMeta(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SoftwareApi.UpdateSoftwareHyperflexDistributable(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SoftwareApi.UpdateSoftwareSolutionDistributable(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsoftwarerepositoryauthorization",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SoftwarerepositoryApi.CreateSoftwarerepositoryAuthorization(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'softwarerepository.Authorization' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsoftwarerepositoryoperatingsystemfile",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SoftwarerepositoryApi.CreateSoftwarerepositoryOperatingSystemFile(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'softwarerepository.OperatingSystemFile' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesoftwarerepositoryauthorization",

										Short: "Update a 'softwarerepository.Authorization' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SoftwarerepositoryApi.UpdateSoftwarerepositoryAuthorization(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SoftwarerepositoryApi.UpdateSoftwarerepositoryOperatingSystemFile(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsolpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SolApi.CreateSolPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'sol.Policy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesolpolicy",

										Short: "Update a 'sol.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SolApi.UpdateSolPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsshpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SshApi.CreateSshPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'ssh.Policy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesshpolicy",

										Short: "Update a 'ssh.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SshApi.UpdateSshPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createstoragediskgrouppolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.CreateStorageDiskGroupPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'storage.DiskGroupPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createstoragestoragepolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.CreateStorageStoragePolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'storage.StoragePolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragecontroller",

										Short: "Update a 'storage.Controller' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStorageController(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStorageDiskGroupPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStorageEnclosure(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStorageEnclosureDisk(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStorageEnclosureDiskSlotEp(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStorageFlexFlashController(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStorageFlexFlashControllerProps(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStorageFlexFlashPhysicalDrive(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStorageFlexFlashVirtualDrive(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStorageFlexUtilController(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStorageFlexUtilPhysicalDrive(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStorageFlexUtilVirtualDrive(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStoragePhysicalDisk(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStoragePhysicalDiskExtension(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStoragePhysicalDiskUsage(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStoragePureArray(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStorageSasExpander(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStorageSasPort(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStorageStoragePolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStorageVdMemberEp(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStorageVirtualDrive(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.UpdateStorageVirtualDriveExtension(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsyslogpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SyslogApi.CreateSyslogPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'syslog.Policy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesyslogpolicy",

										Short: "Update a 'syslog.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SyslogApi.UpdateSyslogPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createtamadvisorycount",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.TamApi.CreateTamAdvisoryCount(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'tam.AdvisoryCount' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createtamadvisoryinfo",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.TamApi.CreateTamAdvisoryInfo(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'tam.AdvisoryInfo' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createtamadvisoryinstance",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.TamApi.CreateTamAdvisoryInstance(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'tam.AdvisoryInstance' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createtamsecurityadvisory",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.TamApi.CreateTamSecurityAdvisory(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'tam.SecurityAdvisory' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatetamadvisorycount",

										Short: "Update a 'tam.AdvisoryCount' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.TamApi.UpdateTamAdvisoryCount(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.TamApi.UpdateTamAdvisoryInfo(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.TamApi.UpdateTamAdvisoryInstance(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.TamApi.UpdateTamSecurityAdvisory(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createtaskpurescopedinventory",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.TaskApi.CreateTaskPureScopedInventory(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'task.PureScopedInventory' resource.",
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "telemetry",

								Short: "Create Telemetry resource(s)",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "querytelemetrytimeseries",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.TelemetryApi.QueryTelemetryTimeSeries(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Perform a Druid time series aggregation request.",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatetopsystem",

										Short: "Update a 'top.System' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.TopApi.UpdateTopSystem(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevirtualizationvmwarecluster",

										Short: "Update a 'virtualization.VmwareCluster' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VirtualizationApi.UpdateVirtualizationVmwareCluster(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VirtualizationApi.UpdateVirtualizationVmwareDatacenter(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VirtualizationApi.UpdateVirtualizationVmwareDatastore(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VirtualizationApi.UpdateVirtualizationVmwareHost(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VirtualizationApi.UpdateVirtualizationVmwareVirtualMachine(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createvmediapolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VmediaApi.CreateVmediaPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vmedia.Policy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevmediapolicy",

										Short: "Update a 'vmedia.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VmediaApi.UpdateVmediaPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createvnicethadapterpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.CreateVnicEthAdapterPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.EthAdapterPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createvnicethif",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.CreateVnicEthIf(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.EthIf' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createvnicethnetworkpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.CreateVnicEthNetworkPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.EthNetworkPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createvnicethqospolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.CreateVnicEthQosPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.EthQosPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createvnicfcadapterpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.CreateVnicFcAdapterPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.FcAdapterPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createvnicfcif",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.CreateVnicFcIf(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.FcIf' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createvnicfcnetworkpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.CreateVnicFcNetworkPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.FcNetworkPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createvnicfcqospolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.CreateVnicFcQosPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.FcQosPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createvniclanconnectivitypolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.CreateVnicLanConnectivityPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.LanConnectivityPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createvnicsanconnectivitypolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.CreateVnicSanConnectivityPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'vnic.SanConnectivityPolicy' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicethadapterpolicy",

										Short: "Update a 'vnic.EthAdapterPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.UpdateVnicEthAdapterPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.UpdateVnicEthIf(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.UpdateVnicEthNetworkPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.UpdateVnicEthQosPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.UpdateVnicFcAdapterPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.UpdateVnicFcIf(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.UpdateVnicFcNetworkPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.UpdateVnicFcQosPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.UpdateVnicLanConnectivityPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.UpdateVnicSanConnectivityPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createworkflowbatchapiexecutor",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.CreateWorkflowBatchApiExecutor(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'workflow.BatchApiExecutor' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createworkflowcustomdatatypedefinition",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.CreateWorkflowCustomDataTypeDefinition(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'workflow.CustomDataTypeDefinition' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createworkflowtaskdefinition",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.CreateWorkflowTaskDefinition(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'workflow.TaskDefinition' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createworkflowworkflowdefinition",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.CreateWorkflowWorkflowDefinition(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'workflow.WorkflowDefinition' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createworkflowworkflowinfo",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.CreateWorkflowWorkflowInfo(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Create a 'workflow.WorkflowInfo' resource.",
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateworkflowbatchapiexecutor",

										Short: "Update a 'workflow.BatchApiExecutor' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.UpdateWorkflowBatchApiExecutor(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.UpdateWorkflowCustomDataTypeDefinition(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.UpdateWorkflowTaskDefinition(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.UpdateWorkflowTaskInfo(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.UpdateWorkflowWorkflowDefinition(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.UpdateWorkflowWorkflowInfo(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
						Use: "delete",

						Short: "Delete resouce(s)",
					}

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "adapter",

								Short: "Delete Adapter resource(s)",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configpolicy",

										Short: "Delete a 'adapter.ConfigPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.AdapterApi.DeleteAdapterConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backup",

										Short: "Delete a 'appliance.Backup' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.ApplianceApi.DeleteApplianceBackup(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.ApplianceApi.DeleteApplianceRestore(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deviceclaim",

										Short: "Delete a 'asset.DeviceClaim' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.AssetApi.DeleteAssetDeviceClaim(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.AssetApi.DeleteAssetDeviceRegistration(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.AssetApi.DeleteAssetManagedDevice(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'bios.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.BiosApi.DeleteBiosPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "precisionpolicy",

										Short: "Delete a 'boot.PrecisionPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.BootApi.DeleteBootPrecisionPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'deviceconnector.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.DeviceconnectorApi.DeleteDeviceconnectorPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "distributable",

										Short: "Delete a 'firmware.Distributable' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.FirmwareApi.DeleteFirmwareDistributable(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.FirmwareApi.DeleteFirmwareDriverDistributable(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.FirmwareApi.DeleteFirmwareServerConfigurationUtilityDistributable(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.FirmwareApi.DeleteFirmwareUpgrade(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hyperflexsoftwarecompatibilityinfo",

										Short: "Delete a 'hcl.HyperflexSoftwareCompatibilityInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HclApi.DeleteHclHyperflexSoftwareCompatibilityInfo(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "appcatalog",

										Short: "Delete a 'hyperflex.AppCatalog' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexAppCatalog(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexAutoSupportPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexCapabilityInfo(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexClusterNetworkPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexClusterProfile(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexClusterStoragePolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexExtFcStoragePolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexExtIscsiStoragePolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexFeatureLimitExternal(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexFeatureLimitInternal(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexHxdpVersion(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexLocalCredentialPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexNodeConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexNodeProfile(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexProxySettingPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexServerFirmwareVersion(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexServerModel(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexSoftwareVersionPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexSysConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexUcsmConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.HyperflexApi.DeleteHyperflexVcenterConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ucsdinfo",

										Short: "Delete a 'iaas.UcsdInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IaasApi.DeleteIaasUcsdInfo(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "account",

										Short: "Delete a 'iam.Account' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamAccount(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamApiKey(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamAppRegistration(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamCertificate(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamCertificateRequest(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamEndPointUser(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamEndPointUserPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamEndPointUserRole(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamIdp(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamLdapGroup(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamLdapPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamLdapProvider(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamOAuthToken(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamPermission(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamPrivateKeySpec(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamQualifier(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamResourceRoles(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamSession(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamSessionLimits(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamTrustPoint(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamUser(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IamApi.DeleteIamUserGroup(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "accountexperience",

										Short: "Delete a 'infra.AccountExperience' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.InfraApi.DeleteInfraAccountExperience(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'ipmioverlan.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.IpmioverlanApi.DeleteIpmioverlanPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'kvm.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.KvmApi.DeleteKvmPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "persistentmemorypolicy",

										Short: "Delete a 'memory.PersistentMemoryPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.MemoryApi.DeleteMemoryPersistentMemoryPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "definition",

										Short: "Delete a 'meta.Definition' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.MetaApi.DeleteMetaDefinition(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'networkconfig.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.NetworkconfigApi.DeleteNetworkconfigPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'ntp.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.NtpApi.DeleteNtpPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "organization",

										Short: "Delete a 'organization.Organization' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.OrganizationApi.DeleteOrganizationOrganization(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configurationfile",

										Short: "Delete a 'os.ConfigurationFile' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.OsApi.DeleteOsConfigurationFile(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backupconfigpolicy",

										Short: "Delete a 'recovery.BackupConfigPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.RecoveryApi.DeleteRecoveryBackupConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.RecoveryApi.DeleteRecoveryBackupProfile(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.RecoveryApi.DeleteRecoveryOnDemandBackup(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.RecoveryApi.DeleteRecoveryRestore(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.RecoveryApi.DeleteRecoveryScheduleConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "group",

										Short: "Delete a 'resource.Group' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.ResourceApi.DeleteResourceGroup(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'sdcard.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.SdcardApi.DeleteSdcardPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "profile",

										Short: "Delete a 'sdwan.Profile' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.SdwanApi.DeleteSdwanProfile(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.SdwanApi.DeleteSdwanRouterNode(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.SdwanApi.DeleteSdwanRouterPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.SdwanApi.DeleteSdwanVmanageAccountPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "profile",

										Short: "Delete a 'server.Profile' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.ServerApi.DeleteServerProfile(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'smtp.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.SmtpApi.DeleteSmtpPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'snmp.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.SnmpApi.DeleteSnmpPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hclmeta",

										Short: "Delete a 'software.HclMeta' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.SoftwareApi.DeleteSoftwareHclMeta(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.SoftwareApi.DeleteSoftwareHyperflexDistributable(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.SoftwareApi.DeleteSoftwareSolutionDistributable(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "operatingsystemfile",

										Short: "Delete a 'softwarerepository.OperatingSystemFile' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.SoftwarerepositoryApi.DeleteSoftwarerepositoryOperatingSystemFile(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'sol.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.SolApi.DeleteSolPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'ssh.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.SshApi.DeleteSshPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "diskgrouppolicy",

										Short: "Delete a 'storage.DiskGroupPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.StorageApi.DeleteStorageDiskGroupPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.StorageApi.DeleteStorageStoragePolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'syslog.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.SyslogApi.DeleteSyslogPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "advisorycount",

										Short: "Delete a 'tam.AdvisoryCount' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.TamApi.DeleteTamAdvisoryCount(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.TamApi.DeleteTamAdvisoryInfo(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.TamApi.DeleteTamAdvisoryInstance(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.TamApi.DeleteTamSecurityAdvisory(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backupinfo",

										Short: "Delete a 'ucsd.BackupInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.UcsdApi.DeleteUcsdBackupInfo(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Delete a 'vmedia.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.VmediaApi.DeleteVmediaPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethadapterpolicy",

										Short: "Delete a 'vnic.EthAdapterPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.VnicApi.DeleteVnicEthAdapterPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.VnicApi.DeleteVnicEthIf(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.VnicApi.DeleteVnicEthNetworkPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.VnicApi.DeleteVnicEthQosPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.VnicApi.DeleteVnicFcAdapterPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.VnicApi.DeleteVnicFcIf(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.VnicApi.DeleteVnicFcNetworkPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.VnicApi.DeleteVnicFcQosPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.VnicApi.DeleteVnicLanConnectivityPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.VnicApi.DeleteVnicSanConnectivityPolicy(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "batchapiexecutor",

										Short: "Delete a 'workflow.BatchApiExecutor' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.WorkflowApi.DeleteWorkflowBatchApiExecutor(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.WorkflowApi.DeleteWorkflowCustomDataTypeDefinition(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.WorkflowApi.DeleteWorkflowTaskDefinition(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.WorkflowApi.DeleteWorkflowWorkflowDefinition(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													httpResponse, err := client.WorkflowApi.DeleteWorkflowWorkflowInfo(authCtx, args[0]).Execute()
													resultHandler(nil, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
					}

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "aaa",

								Short: "Get or list Aaa resource(s)",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "auditrecord",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AaaApi.GetAaaAuditRecordList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'aaa.AuditRecord' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AaaApi.GetAaaAuditRecordByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AdapterApi.GetAdapterConfigPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'adapter.ConfigPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AdapterApi.GetAdapterConfigPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.AdapterApi.GetAdapterExtEthInterfaceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'adapter.ExtEthInterface' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AdapterApi.GetAdapterExtEthInterfaceByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.AdapterApi.GetAdapterHostEthInterfaceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'adapter.HostEthInterface' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AdapterApi.GetAdapterHostEthInterfaceByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.AdapterApi.GetAdapterHostFcInterfaceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'adapter.HostFcInterface' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AdapterApi.GetAdapterHostFcInterfaceByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.AdapterApi.GetAdapterHostIscsiInterfaceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'adapter.HostIscsiInterface' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AdapterApi.GetAdapterHostIscsiInterfaceByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.AdapterApi.GetAdapterUnitList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'adapter.Unit' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AdapterApi.GetAdapterUnitByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backup",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.GetApplianceBackupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.Backup' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.GetApplianceBackupByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ApplianceApi.GetApplianceBackupPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.BackupPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.GetApplianceBackupPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ApplianceApi.GetApplianceCertificateSettingList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.CertificateSetting' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.GetApplianceCertificateSettingByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ApplianceApi.GetApplianceDataExportPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.DataExportPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.GetApplianceDataExportPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ApplianceApi.GetApplianceDeviceClaimList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.DeviceClaim' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.GetApplianceDeviceClaimByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ApplianceApi.GetApplianceDiagSettingList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.DiagSetting' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.GetApplianceDiagSettingByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ApplianceApi.GetApplianceImageBundleList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.ImageBundle' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.GetApplianceImageBundleByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ApplianceApi.GetApplianceNodeInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.NodeInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.GetApplianceNodeInfoByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ApplianceApi.GetApplianceReleaseNoteList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.ReleaseNote' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.GetApplianceReleaseNoteByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ApplianceApi.GetApplianceRestoreList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.Restore' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.GetApplianceRestoreByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ApplianceApi.GetApplianceSetupInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.SetupInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.GetApplianceSetupInfoByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ApplianceApi.GetApplianceSystemInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.SystemInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.GetApplianceSystemInfoByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ApplianceApi.GetApplianceUpgradeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.Upgrade' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.GetApplianceUpgradeByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ApplianceApi.GetApplianceUpgradePolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'appliance.UpgradePolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.GetApplianceUpgradePolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "clustermember",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AssetApi.GetAssetClusterMemberList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'asset.ClusterMember' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AssetApi.GetAssetClusterMemberByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.AssetApi.GetAssetDeviceConfigurationList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'asset.DeviceConfiguration' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AssetApi.GetAssetDeviceConfigurationByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.AssetApi.GetAssetDeviceConnectorManagerList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'asset.DeviceConnectorManager' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AssetApi.GetAssetDeviceConnectorManagerByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.AssetApi.GetAssetDeviceContractInformationList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'asset.DeviceContractInformation' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AssetApi.GetAssetDeviceContractInformationByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.AssetApi.GetAssetDeviceRegistrationList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'asset.DeviceRegistration' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AssetApi.GetAssetDeviceRegistrationByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.AssetApi.GetAssetManagedDeviceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'asset.ManagedDevice' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AssetApi.GetAssetManagedDeviceByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "bootmode",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.BiosApi.GetBiosBootModeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'bios.BootMode' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.BiosApi.GetBiosBootModeByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.BiosApi.GetBiosPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'bios.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.BiosApi.GetBiosPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.BiosApi.GetBiosUnitList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'bios.Unit' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.BiosApi.GetBiosUnitByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "devicebootmode",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.BootApi.GetBootDeviceBootModeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'boot.DeviceBootMode' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.BootApi.GetBootDeviceBootModeByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.BootApi.GetBootPrecisionPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'boot.PrecisionPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.BootApi.GetBootPrecisionPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "blade",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ComputeApi.GetComputeBladeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'compute.Blade' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ComputeApi.GetComputeBladeByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ComputeApi.GetComputeBoardList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'compute.Board' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ComputeApi.GetComputeBoardByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ComputeApi.GetComputePhysicalSummaryList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'compute.PhysicalSummary' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ComputeApi.GetComputePhysicalSummaryByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ComputeApi.GetComputeRackUnitList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'compute.RackUnit' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ComputeApi.GetComputeRackUnitByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ComputeApi.GetComputeServerSettingList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'compute.ServerSetting' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ComputeApi.GetComputeServerSettingByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "alarm",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.CondApi.GetCondAlarmList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'cond.Alarm' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.CondApi.GetCondAlarmByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.CondApi.GetCondHclStatusList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'cond.HclStatus' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.CondApi.GetCondHclStatusByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.CondApi.GetCondHclStatusDetailList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'cond.HclStatusDetail' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.CondApi.GetCondHclStatusDetailByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.CondApi.GetCondHclStatusJobList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'cond.HclStatusJob' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.CondApi.GetCondHclStatusJobByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.DeviceconnectorApi.GetDeviceconnectorPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'deviceconnector.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.DeviceconnectorApi.GetDeviceconnectorPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "chassis",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.EquipmentApi.GetEquipmentChassisList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.Chassis' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.GetEquipmentChassisByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.EquipmentApi.GetEquipmentDeviceSummaryList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.DeviceSummary' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.GetEquipmentDeviceSummaryByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.EquipmentApi.GetEquipmentFanList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.Fan' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.GetEquipmentFanByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.EquipmentApi.GetEquipmentFanModuleList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.FanModule' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.GetEquipmentFanModuleByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.EquipmentApi.GetEquipmentFexList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.Fex' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.GetEquipmentFexByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.EquipmentApi.GetEquipmentIoCardList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.IoCard' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.GetEquipmentIoCardByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.EquipmentApi.GetEquipmentIoExpanderList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.IoExpander' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.GetEquipmentIoExpanderByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.EquipmentApi.GetEquipmentLocatorLedList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.LocatorLed' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.GetEquipmentLocatorLedByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.EquipmentApi.GetEquipmentPsuList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.Psu' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.GetEquipmentPsuByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.EquipmentApi.GetEquipmentRackEnclosureList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.RackEnclosure' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.GetEquipmentRackEnclosureByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.EquipmentApi.GetEquipmentRackEnclosureSlotList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.RackEnclosureSlot' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.GetEquipmentRackEnclosureSlotByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.EquipmentApi.GetEquipmentSharedIoModuleList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.SharedIoModule' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.GetEquipmentSharedIoModuleByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.EquipmentApi.GetEquipmentSwitchCardList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.SwitchCard' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.GetEquipmentSwitchCardByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.EquipmentApi.GetEquipmentSystemIoControllerList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.SystemIoController' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.GetEquipmentSystemIoControllerByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.EquipmentApi.GetEquipmentTpmList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'equipment.Tpm' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.GetEquipmentTpmByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "physicalport",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.EtherApi.GetEtherPhysicalPortList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'ether.PhysicalPort' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EtherApi.GetEtherPhysicalPortByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "authorization",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ExternalsiteApi.GetExternalsiteAuthorizationList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'externalsite.Authorization' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ExternalsiteApi.GetExternalsiteAuthorizationByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "instance",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FaultApi.GetFaultInstanceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'fault.Instance' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FaultApi.GetFaultInstanceByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "physicalport",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FcApi.GetFcPhysicalPortList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'fc.PhysicalPort' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FcApi.GetFcPhysicalPortByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "distributable",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FirmwareApi.GetFirmwareDistributableList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'firmware.Distributable' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FirmwareApi.GetFirmwareDistributableByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.FirmwareApi.GetFirmwareDriverDistributableList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'firmware.DriverDistributable' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FirmwareApi.GetFirmwareDriverDistributableByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.FirmwareApi.GetFirmwareEulaList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'firmware.Eula' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FirmwareApi.GetFirmwareEulaByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.FirmwareApi.GetFirmwareRunningFirmwareList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'firmware.RunningFirmware' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FirmwareApi.GetFirmwareRunningFirmwareByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.FirmwareApi.GetFirmwareServerConfigurationUtilityDistributableList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'firmware.ServerConfigurationUtilityDistributable' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FirmwareApi.GetFirmwareServerConfigurationUtilityDistributableByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.FirmwareApi.GetFirmwareUpgradeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'firmware.Upgrade' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FirmwareApi.GetFirmwareUpgradeByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.FirmwareApi.GetFirmwareUpgradeStatusList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'firmware.UpgradeStatus' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FirmwareApi.GetFirmwareUpgradeStatusByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "catalog",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ForecastApi.GetForecastCatalogList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'forecast.Catalog' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ForecastApi.GetForecastCatalogByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ForecastApi.GetForecastDefinitionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'forecast.Definition' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ForecastApi.GetForecastDefinitionByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ForecastApi.GetForecastInstanceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'forecast.Instance' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ForecastApi.GetForecastInstanceByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "card",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.GraphicsApi.GetGraphicsCardList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'graphics.Card' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.GraphicsApi.GetGraphicsCardByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.GraphicsApi.GetGraphicsControllerList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'graphics.Controller' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.GraphicsApi.GetGraphicsControllerByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "driverimage",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HclApi.GetHclDriverImageList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hcl.DriverImage' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HclApi.GetHclDriverImageByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HclApi.GetHclExemptedCatalogList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hcl.ExemptedCatalog' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HclApi.GetHclExemptedCatalogByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HclApi.GetHclHyperflexSoftwareCompatibilityInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hcl.HyperflexSoftwareCompatibilityInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HclApi.GetHclHyperflexSoftwareCompatibilityInfoByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HclApi.GetHclOperatingSystemList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hcl.OperatingSystem' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HclApi.GetHclOperatingSystemByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HclApi.GetHclOperatingSystemVendorList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hcl.OperatingSystemVendor' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HclApi.GetHclOperatingSystemVendorByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HclApi.GetHclServiceStatusList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hcl.ServiceStatus' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HclApi.GetHclServiceStatusByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "alarm",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexAlarmList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.Alarm' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexAlarmByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexAppCatalogList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.AppCatalog' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexAppCatalogByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexAutoSupportPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.AutoSupportPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexAutoSupportPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexCapabilityInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.CapabilityInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexCapabilityInfoByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexClusterList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.Cluster' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexClusterByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexClusterNetworkPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ClusterNetworkPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexClusterNetworkPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexClusterProfileList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ClusterProfile' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexClusterProfileByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexClusterStoragePolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ClusterStoragePolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexClusterStoragePolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexConfigResultList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ConfigResult' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexConfigResultByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexConfigResultEntryList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ConfigResultEntry' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexConfigResultEntryByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexExtFcStoragePolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ExtFcStoragePolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexExtFcStoragePolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexExtIscsiStoragePolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ExtIscsiStoragePolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexExtIscsiStoragePolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexFeatureLimitExternalList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.FeatureLimitExternal' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexFeatureLimitExternalByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexFeatureLimitInternalList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.FeatureLimitInternal' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexFeatureLimitInternalByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexHealthList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.Health' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexHealthByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexHxdpVersionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.HxdpVersion' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexHxdpVersionByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexLocalCredentialPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.LocalCredentialPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexLocalCredentialPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexNodeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.Node' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexNodeByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexNodeConfigPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.NodeConfigPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexNodeConfigPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexNodeProfileList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.NodeProfile' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexNodeProfileByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexProxySettingPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ProxySettingPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexProxySettingPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexServerFirmwareVersionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ServerFirmwareVersion' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexServerFirmwareVersionByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexServerModelList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.ServerModel' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexServerModelByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexSoftwareVersionPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.SoftwareVersionPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexSoftwareVersionPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexSysConfigPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.SysConfigPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexSysConfigPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexUcsmConfigPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.UcsmConfigPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexUcsmConfigPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.HyperflexApi.GetHyperflexVcenterConfigPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'hyperflex.VcenterConfigPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.GetHyperflexVcenterConfigPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "connectorpack",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IaasApi.GetIaasConnectorPackList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iaas.ConnectorPack' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IaasApi.GetIaasConnectorPackByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IaasApi.GetIaasDeviceStatusList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iaas.DeviceStatus' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IaasApi.GetIaasDeviceStatusByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IaasApi.GetIaasLicenseInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iaas.LicenseInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IaasApi.GetIaasLicenseInfoByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IaasApi.GetIaasMostRunTasksList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iaas.MostRunTasks' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IaasApi.GetIaasMostRunTasksByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IaasApi.GetIaasUcsdInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iaas.UcsdInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IaasApi.GetIaasUcsdInfoByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IaasApi.GetIaasUcsdManagedInfraList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iaas.UcsdManagedInfra' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IaasApi.GetIaasUcsdManagedInfraByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "account",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamAccountList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.Account' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamAccountByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamApiKeyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.ApiKey' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamApiKeyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamAppRegistrationList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.AppRegistration' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamAppRegistrationByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamCertificateList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.Certificate' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamCertificateByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamCertificateRequestList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.CertificateRequest' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamCertificateRequestByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamDomainGroupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.DomainGroup' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamDomainGroupByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamEndPointPrivilegeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.EndPointPrivilege' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamEndPointPrivilegeByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamEndPointRoleList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.EndPointRole' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamEndPointRoleByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamEndPointUserList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.EndPointUser' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamEndPointUserByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamEndPointUserPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.EndPointUserPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamEndPointUserPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamEndPointUserRoleList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.EndPointUserRole' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamEndPointUserRoleByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamIdpList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.Idp' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamIdpByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamIdpReferenceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.IdpReference' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamIdpReferenceByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamLdapGroupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.LdapGroup' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamLdapGroupByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamLdapPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.LdapPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamLdapPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamLdapProviderList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.LdapProvider' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamLdapProviderByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamOAuthTokenList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.OAuthToken' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamOAuthTokenByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamPermissionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.Permission' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamPermissionByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamPrivateKeySpecList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.PrivateKeySpec' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamPrivateKeySpecByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamPrivilegeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.Privilege' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamPrivilegeByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamPrivilegeSetList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.PrivilegeSet' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamPrivilegeSetByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamQualifierList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.Qualifier' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamQualifierByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamResourceLimitsList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.ResourceLimits' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamResourceLimitsByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamResourcePermissionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.ResourcePermission' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamResourcePermissionByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamResourceRolesList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.ResourceRoles' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamResourceRolesByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamRoleList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.Role' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamRoleByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamSecurityHolderList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.SecurityHolder' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamSecurityHolderByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamServiceProviderList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.ServiceProvider' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamServiceProviderByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamSessionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.Session' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamSessionByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamSessionLimitsList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.SessionLimits' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamSessionLimitsByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamSystemList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.System' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamSystemByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamTrustPointList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.TrustPoint' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamTrustPointByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamUserList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.User' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamUserByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamUserGroupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.UserGroup' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamUserGroupByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.IamApi.GetIamUserPreferenceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'iam.UserPreference' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.GetIamUserPreferenceByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "accountexperience",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.InfraApi.GetInfraAccountExperienceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'infra.AccountExperience' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.InfraApi.GetInfraAccountExperienceByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deviceinfo",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.InventoryApi.GetInventoryDeviceInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'inventory.DeviceInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.InventoryApi.GetInventoryDeviceInfoByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.InventoryApi.GetInventoryDnMoBindingList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'inventory.DnMoBinding' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.InventoryApi.GetInventoryDnMoBindingByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.InventoryApi.GetInventoryGenericInventoryList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'inventory.GenericInventory' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.InventoryApi.GetInventoryGenericInventoryByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.InventoryApi.GetInventoryGenericInventoryHolderList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'inventory.GenericInventoryHolder' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.InventoryApi.GetInventoryGenericInventoryHolderByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IpmioverlanApi.GetIpmioverlanPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'ipmioverlan.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IpmioverlanApi.GetIpmioverlanPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "kvmsession",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.KvmApi.GetKvmKvmSessionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'kvm.KvmSession' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.KvmApi.GetKvmKvmSessionByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.KvmApi.GetKvmPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'kvm.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.KvmApi.GetKvmPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "accountlicensedata",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.LicenseApi.GetLicenseAccountLicenseDataList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'license.AccountLicenseData' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.LicenseApi.GetLicenseAccountLicenseDataByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.LicenseApi.GetLicenseCustomerOpList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'license.CustomerOp' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.LicenseApi.GetLicenseCustomerOpByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.LicenseApi.GetLicenseLicenseInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'license.LicenseInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.LicenseApi.GetLicenseLicenseInfoByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.LicenseApi.GetLicenseSmartlicenseTokenList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'license.SmartlicenseToken' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.LicenseApi.GetLicenseSmartlicenseTokenByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "serviceprofile",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.LsApi.GetLsServiceProfileList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'ls.ServiceProfile' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.LsApi.GetLsServiceProfileByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "controller",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ManagementApi.GetManagementControllerList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'management.Controller' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ManagementApi.GetManagementControllerByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ManagementApi.GetManagementEntityList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'management.Entity' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ManagementApi.GetManagementEntityByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ManagementApi.GetManagementInterfaceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'management.Interface' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ManagementApi.GetManagementInterfaceByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "array",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.MemoryApi.GetMemoryArrayList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'memory.Array' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.GetMemoryArrayByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryConfigResultList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'memory.PersistentMemoryConfigResult' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryConfigResultByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryConfigurationList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'memory.PersistentMemoryConfiguration' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryConfigurationByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryNamespaceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'memory.PersistentMemoryNamespace' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryNamespaceByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryNamespaceConfigResultList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'memory.PersistentMemoryNamespaceConfigResult' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryNamespaceConfigResultByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'memory.PersistentMemoryPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryRegionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'memory.PersistentMemoryRegion' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryRegionByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryUnitList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'memory.PersistentMemoryUnit' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryUnitByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.MemoryApi.GetMemoryUnitList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'memory.Unit' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.GetMemoryUnitByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "definition",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.MetaApi.GetMetaDefinitionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'meta.Definition' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MetaApi.GetMetaDefinitionByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "element",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NetworkApi.GetNetworkElementList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'network.Element' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NetworkApi.GetNetworkElementByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.NetworkApi.GetNetworkElementSummaryList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'network.ElementSummary' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NetworkApi.GetNetworkElementSummaryByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NetworkconfigApi.GetNetworkconfigPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'networkconfig.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NetworkconfigApi.GetNetworkconfigPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "apicccopost",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiaapiApi.GetNiaapiApicCcoPostList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.ApicCcoPost' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NiaapiApi.GetNiaapiApicCcoPostByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.NiaapiApi.GetNiaapiApicFieldNoticeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.ApicFieldNotice' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NiaapiApi.GetNiaapiApicFieldNoticeByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.NiaapiApi.GetNiaapiApicHweolList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.ApicHweol' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NiaapiApi.GetNiaapiApicHweolByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.NiaapiApi.GetNiaapiApicLatestMaintainedReleaseList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.ApicLatestMaintainedRelease' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NiaapiApi.GetNiaapiApicLatestMaintainedReleaseByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.NiaapiApi.GetNiaapiApicReleaseRecommendList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.ApicReleaseRecommend' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NiaapiApi.GetNiaapiApicReleaseRecommendByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.NiaapiApi.GetNiaapiApicSweolList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.ApicSweol' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NiaapiApi.GetNiaapiApicSweolByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.NiaapiApi.GetNiaapiDcnmCcoPostList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.DcnmCcoPost' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NiaapiApi.GetNiaapiDcnmCcoPostByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.NiaapiApi.GetNiaapiDcnmFieldNoticeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.DcnmFieldNotice' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NiaapiApi.GetNiaapiDcnmFieldNoticeByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.NiaapiApi.GetNiaapiDcnmHweolList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.DcnmHweol' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NiaapiApi.GetNiaapiDcnmHweolByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.NiaapiApi.GetNiaapiDcnmLatestMaintainedReleaseList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.DcnmLatestMaintainedRelease' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NiaapiApi.GetNiaapiDcnmLatestMaintainedReleaseByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.NiaapiApi.GetNiaapiDcnmReleaseRecommendList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.DcnmReleaseRecommend' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NiaapiApi.GetNiaapiDcnmReleaseRecommendByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.NiaapiApi.GetNiaapiDcnmSweolList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.DcnmSweol' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NiaapiApi.GetNiaapiDcnmSweolByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.NiaapiApi.GetNiaapiFileDownloaderList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.FileDownloader' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NiaapiApi.GetNiaapiFileDownloaderByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.NiaapiApi.GetNiaapiNiaMetadataList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.NiaMetadata' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NiaapiApi.GetNiaapiNiaMetadataByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.NiaapiApi.GetNiaapiVersionRegexList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niaapi.VersionRegex' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NiaapiApi.GetNiaapiVersionRegexByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "niainventory",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiatelemetryApi.GetNiatelemetryNiaInventoryList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niatelemetry.NiaInventory' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NiatelemetryApi.GetNiatelemetryNiaInventoryByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.NiatelemetryApi.GetNiatelemetryNiaLicenseStateList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'niatelemetry.NiaLicenseState' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NiatelemetryApi.GetNiatelemetryNiaLicenseStateByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NtpApi.GetNtpPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'ntp.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NtpApi.GetNtpPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "organization",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.OrganizationApi.GetOrganizationOrganizationList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'organization.Organization' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.OrganizationApi.GetOrganizationOrganizationByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "catalog",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.OsApi.GetOsCatalogList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'os.Catalog' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.OsApi.GetOsCatalogByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.OsApi.GetOsConfigurationFileList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'os.ConfigurationFile' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.OsApi.GetOsConfigurationFileByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.OsApi.GetOsInstallList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'os.Install' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.OsApi.GetOsInstallByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "coprocessorcard",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.PciApi.GetPciCoprocessorCardList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'pci.CoprocessorCard' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.PciApi.GetPciCoprocessorCardByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.PciApi.GetPciDeviceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'pci.Device' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.PciApi.GetPciDeviceByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.PciApi.GetPciLinkList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'pci.Link' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.PciApi.GetPciLinkByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.PciApi.GetPciSwitchList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'pci.Switch' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.PciApi.GetPciSwitchByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "group",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.PortApi.GetPortGroupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'port.Group' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.PortApi.GetPortGroupByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.PortApi.GetPortSubGroupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'port.SubGroup' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.PortApi.GetPortSubGroupByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "unit",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ProcessorApi.GetProcessorUnitList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'processor.Unit' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ProcessorApi.GetProcessorUnitByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backupconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.RecoveryApi.GetRecoveryBackupConfigPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'recovery.BackupConfigPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.RecoveryApi.GetRecoveryBackupConfigPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.RecoveryApi.GetRecoveryBackupProfileList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'recovery.BackupProfile' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.RecoveryApi.GetRecoveryBackupProfileByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.RecoveryApi.GetRecoveryConfigResultList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'recovery.ConfigResult' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.RecoveryApi.GetRecoveryConfigResultByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.RecoveryApi.GetRecoveryConfigResultEntryList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'recovery.ConfigResultEntry' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.RecoveryApi.GetRecoveryConfigResultEntryByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.RecoveryApi.GetRecoveryOnDemandBackupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'recovery.OnDemandBackup' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.RecoveryApi.GetRecoveryOnDemandBackupByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.RecoveryApi.GetRecoveryRestoreList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'recovery.Restore' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.RecoveryApi.GetRecoveryRestoreByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.RecoveryApi.GetRecoveryScheduleConfigPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'recovery.ScheduleConfigPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.RecoveryApi.GetRecoveryScheduleConfigPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "group",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ResourceApi.GetResourceGroupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'resource.Group' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ResourceApi.GetResourceGroupByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ResourceApi.GetResourceGroupMemberList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'resource.GroupMember' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ResourceApi.GetResourceGroupMemberByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ResourceApi.GetResourceLicenseResourceCountList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'resource.LicenseResourceCount' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ResourceApi.GetResourceLicenseResourceCountByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ResourceApi.GetResourceMembershipList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'resource.Membership' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ResourceApi.GetResourceMembershipByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ResourceApi.GetResourceMembershipHolderList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'resource.MembershipHolder' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ResourceApi.GetResourceMembershipHolderByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SdcardApi.GetSdcardPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'sdcard.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SdcardApi.GetSdcardPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "profile",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SdwanApi.GetSdwanProfileList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'sdwan.Profile' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SdwanApi.GetSdwanProfileByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.SdwanApi.GetSdwanRouterNodeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'sdwan.RouterNode' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SdwanApi.GetSdwanRouterNodeByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.SdwanApi.GetSdwanRouterPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'sdwan.RouterPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SdwanApi.GetSdwanRouterPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.SdwanApi.GetSdwanVmanageAccountPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'sdwan.VmanageAccountPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SdwanApi.GetSdwanVmanageAccountPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "searchitem",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SearchApi.GetSearchSearchItemList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'search.SearchItem' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SearchApi.GetSearchSearchItemByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.SearchApi.GetSearchTagItemList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'search.TagItem' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SearchApi.GetSearchTagItemByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "unit",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SecurityApi.GetSecurityUnitList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'security.Unit' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SecurityApi.GetSecurityUnitByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configchangedetail",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ServerApi.GetServerConfigChangeDetailList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'server.ConfigChangeDetail' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ServerApi.GetServerConfigChangeDetailByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ServerApi.GetServerConfigImportList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'server.ConfigImport' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ServerApi.GetServerConfigImportByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ServerApi.GetServerConfigResultList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'server.ConfigResult' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ServerApi.GetServerConfigResultByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ServerApi.GetServerConfigResultEntryList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'server.ConfigResultEntry' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ServerApi.GetServerConfigResultEntryByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.ServerApi.GetServerProfileList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'server.Profile' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ServerApi.GetServerProfileByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SmtpApi.GetSmtpPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'smtp.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SmtpApi.GetSmtpPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SnmpApi.GetSnmpPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'snmp.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SnmpApi.GetSnmpPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hclmeta",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SoftwareApi.GetSoftwareHclMetaList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'software.HclMeta' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SoftwareApi.GetSoftwareHclMetaByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.SoftwareApi.GetSoftwareHyperflexDistributableList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'software.HyperflexDistributable' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SoftwareApi.GetSoftwareHyperflexDistributableByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.SoftwareApi.GetSoftwareSolutionDistributableList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'software.SolutionDistributable' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SoftwareApi.GetSoftwareSolutionDistributableByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "authorization",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SoftwarerepositoryApi.GetSoftwarerepositoryAuthorizationList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'softwarerepository.Authorization' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SoftwarerepositoryApi.GetSoftwarerepositoryAuthorizationByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.SoftwarerepositoryApi.GetSoftwarerepositoryCatalogList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'softwarerepository.Catalog' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SoftwarerepositoryApi.GetSoftwarerepositoryCatalogByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.SoftwarerepositoryApi.GetSoftwarerepositoryOperatingSystemFileList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'softwarerepository.OperatingSystemFile' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SoftwarerepositoryApi.GetSoftwarerepositoryOperatingSystemFileByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SolApi.GetSolPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'sol.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SolApi.GetSolPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SshApi.GetSshPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'ssh.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SshApi.GetSshPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "controller",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageControllerList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.Controller' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStorageControllerByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStorageDiskGroupPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.DiskGroupPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStorageDiskGroupPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStorageEnclosureList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.Enclosure' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStorageEnclosureByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStorageEnclosureDiskList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.EnclosureDisk' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStorageEnclosureDiskByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStorageEnclosureDiskSlotEpList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.EnclosureDiskSlotEp' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStorageEnclosureDiskSlotEpByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStorageFlexFlashControllerList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.FlexFlashController' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStorageFlexFlashControllerByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStorageFlexFlashControllerPropsList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.FlexFlashControllerProps' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStorageFlexFlashControllerPropsByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStorageFlexFlashPhysicalDriveList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.FlexFlashPhysicalDrive' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStorageFlexFlashPhysicalDriveByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStorageFlexFlashVirtualDriveList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.FlexFlashVirtualDrive' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStorageFlexFlashVirtualDriveByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStorageFlexUtilControllerList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.FlexUtilController' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStorageFlexUtilControllerByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStorageFlexUtilPhysicalDriveList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.FlexUtilPhysicalDrive' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStorageFlexUtilPhysicalDriveByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStorageFlexUtilVirtualDriveList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.FlexUtilVirtualDrive' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStorageFlexUtilVirtualDriveByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStoragePhysicalDiskList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PhysicalDisk' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStoragePhysicalDiskByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStoragePhysicalDiskExtensionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PhysicalDiskExtension' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStoragePhysicalDiskExtensionByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStoragePhysicalDiskUsageList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PhysicalDiskUsage' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStoragePhysicalDiskUsageByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStoragePureArrayList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureArray' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStoragePureArrayByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStoragePureControllerList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureController' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStoragePureControllerByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStoragePureDiskList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureDisk' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStoragePureDiskByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStoragePureHostList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureHost' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStoragePureHostByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStoragePureHostGroupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureHostGroup' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStoragePureHostGroupByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStoragePureHostLunList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureHostLun' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStoragePureHostLunByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStoragePurePortList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PurePort' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStoragePurePortByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStoragePureProtectionGroupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureProtectionGroup' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStoragePureProtectionGroupByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStoragePureProtectionGroupSnapshotList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureProtectionGroupSnapshot' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStoragePureProtectionGroupSnapshotByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStoragePureReplicationScheduleList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureReplicationSchedule' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStoragePureReplicationScheduleByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStoragePureSnapshotScheduleList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureSnapshotSchedule' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStoragePureSnapshotScheduleByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStoragePureVolumeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureVolume' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStoragePureVolumeByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStoragePureVolumeSnapshotList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.PureVolumeSnapshot' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStoragePureVolumeSnapshotByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStorageSasExpanderList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.SasExpander' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStorageSasExpanderByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStorageSasPortList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.SasPort' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStorageSasPortByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStorageStoragePolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.StoragePolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStorageStoragePolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStorageVdMemberEpList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.VdMemberEp' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStorageVdMemberEpByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStorageVirtualDriveList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.VirtualDrive' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStorageVirtualDriveByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.StorageApi.GetStorageVirtualDriveExtensionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'storage.VirtualDriveExtension' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.GetStorageVirtualDriveExtensionByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SyslogApi.GetSyslogPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'syslog.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SyslogApi.GetSyslogPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "advisorycount",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.TamApi.GetTamAdvisoryCountList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'tam.AdvisoryCount' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.TamApi.GetTamAdvisoryCountByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.TamApi.GetTamAdvisoryInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'tam.AdvisoryInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.TamApi.GetTamAdvisoryInfoByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.TamApi.GetTamAdvisoryInstanceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'tam.AdvisoryInstance' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.TamApi.GetTamAdvisoryInstanceByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.TamApi.GetTamSecurityAdvisoryList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'tam.SecurityAdvisory' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.TamApi.GetTamSecurityAdvisoryByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "auditlog",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.TerminalApi.GetTerminalAuditLogList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'terminal.AuditLog' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.TerminalApi.GetTerminalAuditLogByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "system",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.TopApi.GetTopSystemList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'top.System' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.TopApi.GetTopSystemByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backupinfo",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.UcsdApi.GetUcsdBackupInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'ucsd.BackupInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.UcsdApi.GetUcsdBackupInfoByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vmwarecluster",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VirtualizationApi.GetVirtualizationVmwareClusterList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'virtualization.VmwareCluster' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VirtualizationApi.GetVirtualizationVmwareClusterByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.VirtualizationApi.GetVirtualizationVmwareDatacenterList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'virtualization.VmwareDatacenter' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VirtualizationApi.GetVirtualizationVmwareDatacenterByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.VirtualizationApi.GetVirtualizationVmwareDatastoreList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'virtualization.VmwareDatastore' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VirtualizationApi.GetVirtualizationVmwareDatastoreByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.VirtualizationApi.GetVirtualizationVmwareHostList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'virtualization.VmwareHost' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VirtualizationApi.GetVirtualizationVmwareHostByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.VirtualizationApi.GetVirtualizationVmwareVcenterList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'virtualization.VmwareVcenter' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VirtualizationApi.GetVirtualizationVmwareVcenterByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.VirtualizationApi.GetVirtualizationVmwareVirtualMachineList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'virtualization.VmwareVirtualMachine' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VirtualizationApi.GetVirtualizationVmwareVirtualMachineByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VmediaApi.GetVmediaPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vmedia.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VmediaApi.GetVmediaPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethadapterpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.GetVnicEthAdapterPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.EthAdapterPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.GetVnicEthAdapterPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.VnicApi.GetVnicEthIfList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.EthIf' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.GetVnicEthIfByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.VnicApi.GetVnicEthNetworkPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.EthNetworkPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.GetVnicEthNetworkPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.VnicApi.GetVnicEthQosPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.EthQosPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.GetVnicEthQosPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.VnicApi.GetVnicFcAdapterPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.FcAdapterPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.GetVnicFcAdapterPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.VnicApi.GetVnicFcIfList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.FcIf' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.GetVnicFcIfByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.VnicApi.GetVnicFcNetworkPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.FcNetworkPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.GetVnicFcNetworkPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.VnicApi.GetVnicFcQosPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.FcQosPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.GetVnicFcQosPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.VnicApi.GetVnicLanConnectivityPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.LanConnectivityPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.GetVnicLanConnectivityPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.VnicApi.GetVnicSanConnectivityPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'vnic.SanConnectivityPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.GetVnicSanConnectivityPolicyByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "batchapiexecutor",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.GetWorkflowBatchApiExecutorList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.BatchApiExecutor' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.GetWorkflowBatchApiExecutorByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.WorkflowApi.GetWorkflowBuildTaskMetaList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.BuildTaskMeta' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.GetWorkflowBuildTaskMetaByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.WorkflowApi.GetWorkflowBuildTaskMetaOwnerList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.BuildTaskMetaOwner' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.GetWorkflowBuildTaskMetaOwnerByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.WorkflowApi.GetWorkflowCatalogList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.Catalog' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.GetWorkflowCatalogByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.WorkflowApi.GetWorkflowCustomDataTypeDefinitionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.CustomDataTypeDefinition' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.GetWorkflowCustomDataTypeDefinitionByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.WorkflowApi.GetWorkflowPendingDynamicWorkflowInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.PendingDynamicWorkflowInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.GetWorkflowPendingDynamicWorkflowInfoByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.WorkflowApi.GetWorkflowTaskDefinitionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.TaskDefinition' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.GetWorkflowTaskDefinitionByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.WorkflowApi.GetWorkflowTaskInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.TaskInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.GetWorkflowTaskInfoByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.WorkflowApi.GetWorkflowTaskMetaList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.TaskMeta' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.GetWorkflowTaskMetaByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.WorkflowApi.GetWorkflowWorkflowDefinitionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.WorkflowDefinition' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.GetWorkflowWorkflowDefinitionByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.WorkflowApi.GetWorkflowWorkflowInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.WorkflowInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.GetWorkflowWorkflowInfoByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

											res, httpResponse, err := client.WorkflowApi.GetWorkflowWorkflowMetaList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},

										Short: "Read a 'workflow.WorkflowMeta' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.GetWorkflowWorkflowMetaByMoid(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
					}

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "adapter",

								Short: "Update Adapter resource(s)",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "configpolicy",

										Short: "Update a 'adapter.ConfigPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AdapterApi.PatchAdapterConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backuppolicy",

										Short: "Update a 'appliance.BackupPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.PatchApplianceBackupPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.PatchApplianceCertificateSetting(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.PatchApplianceDataExportPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.PatchApplianceDiagSetting(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.PatchApplianceSetupInfo(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.PatchApplianceUpgrade(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ApplianceApi.PatchApplianceUpgradePolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deviceconfiguration",

										Short: "Update a 'asset.DeviceConfiguration' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AssetApi.PatchAssetDeviceConfiguration(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AssetApi.PatchAssetDeviceContractInformation(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AssetApi.PatchAssetDeviceRegistration(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.AssetApi.PatchAssetManagedDevice(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "bootmode",

										Short: "Update a 'bios.BootMode' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.BiosApi.PatchBiosBootMode(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.BiosApi.PatchBiosPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.BiosApi.PatchBiosUnit(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "devicebootmode",

										Short: "Update a 'boot.DeviceBootMode' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.BootApi.PatchBootDeviceBootMode(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.BootApi.PatchBootPrecisionPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "blade",

										Short: "Update a 'compute.Blade' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ComputeApi.PatchComputeBlade(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ComputeApi.PatchComputeBoard(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ComputeApi.PatchComputeRackUnit(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ComputeApi.PatchComputeServerSetting(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'deviceconnector.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.DeviceconnectorApi.PatchDeviceconnectorPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "chassis",

										Short: "Update a 'equipment.Chassis' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.PatchEquipmentChassis(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.PatchEquipmentFan(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.PatchEquipmentFanModule(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.PatchEquipmentFex(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.PatchEquipmentIoCard(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.PatchEquipmentIoExpander(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.PatchEquipmentLocatorLed(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.PatchEquipmentPsu(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.PatchEquipmentRackEnclosure(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.PatchEquipmentRackEnclosureSlot(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.PatchEquipmentSharedIoModule(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.PatchEquipmentSwitchCard(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.PatchEquipmentSystemIoController(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EquipmentApi.PatchEquipmentTpm(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "physicalport",

										Short: "Update a 'ether.PhysicalPort' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.EtherApi.PatchEtherPhysicalPort(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "authorization",

										Short: "Update a 'externalsite.Authorization' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ExternalsiteApi.PatchExternalsiteAuthorization(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "instance",

										Short: "Update a 'fault.Instance' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FaultApi.PatchFaultInstance(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "physicalport",

										Short: "Update a 'fc.PhysicalPort' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FcApi.PatchFcPhysicalPort(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "distributable",

										Short: "Update a 'firmware.Distributable' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FirmwareApi.PatchFirmwareDistributable(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FirmwareApi.PatchFirmwareDriverDistributable(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FirmwareApi.PatchFirmwareRunningFirmware(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.FirmwareApi.PatchFirmwareServerConfigurationUtilityDistributable(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "card",

										Short: "Update a 'graphics.Card' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.GraphicsApi.PatchGraphicsCard(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.GraphicsApi.PatchGraphicsController(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hyperflexsoftwarecompatibilityinfo",

										Short: "Update a 'hcl.HyperflexSoftwareCompatibilityInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HclApi.PatchHclHyperflexSoftwareCompatibilityInfo(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "appcatalog",

										Short: "Update a 'hyperflex.AppCatalog' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexAppCatalog(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexAutoSupportPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexCapabilityInfo(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexCluster(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexClusterNetworkPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexClusterProfile(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexClusterStoragePolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexExtFcStoragePolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexExtIscsiStoragePolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexFeatureLimitExternal(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexFeatureLimitInternal(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexHxdpVersion(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexLocalCredentialPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexNodeConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexNodeProfile(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexProxySettingPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexServerFirmwareVersion(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexServerModel(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexSoftwareVersionPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexSysConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexUcsmConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.HyperflexApi.PatchHyperflexVcenterConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ucsdinfo",

										Short: "Update a 'iaas.UcsdInfo' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IaasApi.PatchIaasUcsdInfo(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "account",

										Short: "Update a 'iam.Account' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamAccount(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamApiKey(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamAppRegistration(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamCertificate(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamCertificateRequest(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamEndPointUser(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamEndPointUserPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamEndPointUserRole(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamIdp(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamIdpReference(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamLdapGroup(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamLdapPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamLdapProvider(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamLocalUserPassword(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamPermission(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamPrivateKeySpec(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamQualifier(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamResourceRoles(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamSessionLimits(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamUser(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamUserGroup(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IamApi.PatchIamUserPreference(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "accountexperience",

										Short: "Update a 'infra.AccountExperience' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.InfraApi.PatchInfraAccountExperience(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "genericinventory",

										Short: "Update a 'inventory.GenericInventory' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.InventoryApi.PatchInventoryGenericInventory(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.InventoryApi.PatchInventoryGenericInventoryHolder(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'ipmioverlan.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.IpmioverlanApi.PatchIpmioverlanPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'kvm.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.KvmApi.PatchKvmPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "accountlicensedata",

										Short: "Update a 'license.AccountLicenseData' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.LicenseApi.PatchLicenseAccountLicenseData(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.LicenseApi.PatchLicenseCustomerOp(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.LicenseApi.PatchLicenseLicenseInfo(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.LicenseApi.PatchLicenseSmartlicenseToken(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "serviceprofile",

										Short: "Update a 'ls.ServiceProfile' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.LsApi.PatchLsServiceProfile(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "controller",

										Short: "Update a 'management.Controller' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ManagementApi.PatchManagementController(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ManagementApi.PatchManagementEntity(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ManagementApi.PatchManagementInterface(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "array",

										Short: "Update a 'memory.Array' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.PatchMemoryArray(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.PatchMemoryPersistentMemoryConfigResult(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.PatchMemoryPersistentMemoryConfiguration(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.PatchMemoryPersistentMemoryNamespace(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.PatchMemoryPersistentMemoryNamespaceConfigResult(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.PatchMemoryPersistentMemoryPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.PatchMemoryPersistentMemoryRegion(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.PatchMemoryPersistentMemoryUnit(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.MemoryApi.PatchMemoryUnit(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "element",

										Short: "Update a 'network.Element' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NetworkApi.PatchNetworkElement(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'networkconfig.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NetworkconfigApi.PatchNetworkconfigPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'ntp.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.NtpApi.PatchNtpPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "organization",

										Short: "Update a 'organization.Organization' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.OrganizationApi.PatchOrganizationOrganization(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "device",

										Short: "Update a 'pci.Device' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.PciApi.PatchPciDevice(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.PciApi.PatchPciLink(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.PciApi.PatchPciSwitch(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "group",

										Short: "Update a 'port.Group' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.PortApi.PatchPortGroup(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.PortApi.PatchPortSubGroup(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "unit",

										Short: "Update a 'processor.Unit' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ProcessorApi.PatchProcessorUnit(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "backupconfigpolicy",

										Short: "Update a 'recovery.BackupConfigPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.RecoveryApi.PatchRecoveryBackupConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.RecoveryApi.PatchRecoveryBackupProfile(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.RecoveryApi.PatchRecoveryOnDemandBackup(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.RecoveryApi.PatchRecoveryScheduleConfigPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "group",

										Short: "Update a 'resource.Group' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ResourceApi.PatchResourceGroup(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'sdcard.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SdcardApi.PatchSdcardPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "profile",

										Short: "Update a 'sdwan.Profile' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SdwanApi.PatchSdwanProfile(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SdwanApi.PatchSdwanRouterNode(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SdwanApi.PatchSdwanRouterPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SdwanApi.PatchSdwanVmanageAccountPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "unit",

										Short: "Update a 'security.Unit' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SecurityApi.PatchSecurityUnit(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "profile",

										Short: "Update a 'server.Profile' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.ServerApi.PatchServerProfile(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'smtp.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SmtpApi.PatchSmtpPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'snmp.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SnmpApi.PatchSnmpPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "hclmeta",

										Short: "Update a 'software.HclMeta' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SoftwareApi.PatchSoftwareHclMeta(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SoftwareApi.PatchSoftwareHyperflexDistributable(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SoftwareApi.PatchSoftwareSolutionDistributable(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "authorization",

										Short: "Update a 'softwarerepository.Authorization' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SoftwarerepositoryApi.PatchSoftwarerepositoryAuthorization(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SoftwarerepositoryApi.PatchSoftwarerepositoryOperatingSystemFile(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'sol.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SolApi.PatchSolPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'ssh.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SshApi.PatchSshPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "controller",

										Short: "Update a 'storage.Controller' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStorageController(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStorageDiskGroupPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStorageEnclosure(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStorageEnclosureDisk(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStorageEnclosureDiskSlotEp(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStorageFlexFlashController(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStorageFlexFlashControllerProps(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStorageFlexFlashPhysicalDrive(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStorageFlexFlashVirtualDrive(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStorageFlexUtilController(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStorageFlexUtilPhysicalDrive(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStorageFlexUtilVirtualDrive(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStoragePhysicalDisk(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStoragePhysicalDiskExtension(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStoragePhysicalDiskUsage(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStoragePureArray(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStorageSasExpander(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStorageSasPort(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStorageStoragePolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStorageVdMemberEp(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStorageVirtualDrive(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.StorageApi.PatchStorageVirtualDriveExtension(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'syslog.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.SyslogApi.PatchSyslogPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "advisorycount",

										Short: "Update a 'tam.AdvisoryCount' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.TamApi.PatchTamAdvisoryCount(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.TamApi.PatchTamAdvisoryInfo(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.TamApi.PatchTamAdvisoryInstance(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.TamApi.PatchTamSecurityAdvisory(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "system",

										Short: "Update a 'top.System' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.TopApi.PatchTopSystem(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "vmwarecluster",

										Short: "Update a 'virtualization.VmwareCluster' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VirtualizationApi.PatchVirtualizationVmwareCluster(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VirtualizationApi.PatchVirtualizationVmwareDatacenter(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VirtualizationApi.PatchVirtualizationVmwareDatastore(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VirtualizationApi.PatchVirtualizationVmwareHost(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VirtualizationApi.PatchVirtualizationVmwareVirtualMachine(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "policy",

										Short: "Update a 'vmedia.Policy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VmediaApi.PatchVmediaPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "ethadapterpolicy",

										Short: "Update a 'vnic.EthAdapterPolicy' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.PatchVnicEthAdapterPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.PatchVnicEthIf(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.PatchVnicEthNetworkPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.PatchVnicEthQosPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.PatchVnicFcAdapterPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.PatchVnicFcIf(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.PatchVnicFcNetworkPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.PatchVnicFcQosPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.PatchVnicLanConnectivityPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.VnicApi.PatchVnicSanConnectivityPolicy(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "batchapiexecutor",

										Short: "Update a 'workflow.BatchApiExecutor' resource.",
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.PatchWorkflowBatchApiExecutor(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.PatchWorkflowCustomDataTypeDefinition(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.PatchWorkflowTaskDefinition(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.PatchWorkflowTaskInfo(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.PatchWorkflowWorkflowDefinition(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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
									}

									cmd.AddCommand(
										func() *cobra.Command {
											cmd := &cobra.Command{
												Use: "moid",

												Run: func(cmd *cobra.Command, args []string) {

													res, httpResponse, err := client.WorkflowApi.PatchWorkflowWorkflowInfo(authCtx, args[0]).Execute()
													resultHandler(res, httpResponse, err)

												},

												Args: cobra.ExactArgs(1),

												Short: "",
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

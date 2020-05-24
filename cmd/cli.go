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

type ResultHandler = func(result interface{}, httpResponse *http.Response, err error)

func GetCommands(client *openapi.APIClient, resultHandler ResultHandler) *cobra.Command {
	rootCmd :=
		func() *cobra.Command {
			cmd := &cobra.Command{
				Use: "",
			}

			cmd.AddCommand(
				func() *cobra.Command {
					cmd := &cobra.Command{
						Use: "delete",
					}

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "adapter",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteadapterconfigpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteappliancebackup",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteappliancerestore",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteassetdeviceclaim",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteassetdeviceregistration",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteassetmanageddevice",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletebiospolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletebootprecisionpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletedeviceconnectorpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletefirmwaredistributable",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletefirmwaredriverdistributable",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletefirmwareserverconfigurationutilitydistributable",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletefirmwareupgrade",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehclhyperflexsoftwarecompatibilityinfo",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexappcatalog",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexautosupportpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexcapabilityinfo",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexclusternetworkpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexclusterprofile",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexclusterstoragepolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexextfcstoragepolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexextiscsistoragepolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexfeaturelimitexternal",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexfeaturelimitinternal",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexhxdpversion",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexlocalcredentialpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexnodeconfigpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexnodeprofile",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexproxysettingpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexserverfirmwareversion",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexservermodel",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexsoftwareversionpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexsysconfigpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexucsmconfigpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletehyperflexvcenterconfigpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiaasucsdinfo",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamaccount",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamapikey",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamappregistration",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamcertificate",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamcertificaterequest",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamendpointuser",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamendpointuserpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamendpointuserrole",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamidp",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamldapgroup",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamldappolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamldapprovider",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamoauthtoken",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiampermission",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamprivatekeyspec",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamqualifier",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamresourceroles",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamsession",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamsessionlimits",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamtrustpoint",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamuser",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteiamusergroup",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteinfraaccountexperience",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteipmioverlanpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletekvmpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletememorypersistentmemorypolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletemetadefinition",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletenetworkconfigpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletentppolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteorganizationorganization",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteosconfigurationfile",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleterecoverybackupconfigpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleterecoverybackupprofile",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleterecoveryondemandbackup",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleterecoveryrestore",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleterecoveryscheduleconfigpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteresourcegroup",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletesdcardpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletesdwanprofile",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletesdwanrouternode",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletesdwanrouterpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletesdwanvmanageaccountpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteserverprofile",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletesmtppolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletesnmppolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletesoftwarehclmeta",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletesoftwarehyperflexdistributable",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletesoftwaresolutiondistributable",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletesoftwarerepositoryoperatingsystemfile",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletesolpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletesshpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletestoragediskgrouppolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletestoragestoragepolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletesyslogpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletetamadvisorycount",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletetamadvisoryinfo",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletetamadvisoryinstance",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletetamsecurityadvisory",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteucsdbackupinfo",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletevmediapolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletevnicethadapterpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletevnicethif",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletevnicethnetworkpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletevnicethqospolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletevnicfcadapterpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletevnicfcif",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletevnicfcnetworkpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletevnicfcqospolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletevniclanconnectivitypolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deletevnicsanconnectivitypolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteworkflowbatchapiexecutor",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteworkflowcustomdatatypedefinition",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteworkflowtaskdefinition",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteworkflowworkflowdefinition",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "deleteworkflowworkflowinfo",
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
					}

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "aaa",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getaaaauditrecordbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getaaaauditrecordlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AaaApi.GetAaaAuditRecordList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "adapter",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getadapterconfigpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getadapterconfigpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AdapterApi.GetAdapterConfigPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getadapterextethinterfacebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getadapterextethinterfacelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AdapterApi.GetAdapterExtEthInterfaceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getadapterhostethinterfacebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getadapterhostethinterfacelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AdapterApi.GetAdapterHostEthInterfaceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getadapterhostfcinterfacebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getadapterhostfcinterfacelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AdapterApi.GetAdapterHostFcInterfaceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getadapterhostiscsiinterfacebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getadapterhostiscsiinterfacelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AdapterApi.GetAdapterHostIscsiInterfaceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getadapterunitbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getadapterunitlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AdapterApi.GetAdapterUnitList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "appliance",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancebackupbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancebackuplist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.GetApplianceBackupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancebackuppolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancebackuppolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.GetApplianceBackupPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancecertificatesettingbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancecertificatesettinglist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.GetApplianceCertificateSettingList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancedataexportpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancedataexportpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.GetApplianceDataExportPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancedeviceclaimbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancedeviceclaimlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.GetApplianceDeviceClaimList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancediagsettingbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancediagsettinglist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.GetApplianceDiagSettingList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getapplianceimagebundlebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getapplianceimagebundlelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.GetApplianceImageBundleList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancenodeinfobymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancenodeinfolist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.GetApplianceNodeInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancereleasenotebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancereleasenotelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.GetApplianceReleaseNoteList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancerestorebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancerestorelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.GetApplianceRestoreList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancesetupinfobymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancesetupinfolist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.GetApplianceSetupInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancesysteminfobymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getappliancesysteminfolist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.GetApplianceSystemInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getapplianceupgradebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getapplianceupgradelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.GetApplianceUpgradeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getapplianceupgradepolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getapplianceupgradepolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.GetApplianceUpgradePolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "asset",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getassetclustermemberbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getassetclustermemberlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AssetApi.GetAssetClusterMemberList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getassetdeviceconfigurationbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getassetdeviceconfigurationlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AssetApi.GetAssetDeviceConfigurationList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getassetdeviceconnectormanagerbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getassetdeviceconnectormanagerlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AssetApi.GetAssetDeviceConnectorManagerList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getassetdevicecontractinformationbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getassetdevicecontractinformationlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AssetApi.GetAssetDeviceContractInformationList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getassetdeviceregistrationbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getassetdeviceregistrationlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AssetApi.GetAssetDeviceRegistrationList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getassetmanageddevicebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getassetmanageddevicelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AssetApi.GetAssetManagedDeviceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "bios",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getbiosbootmodebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getbiosbootmodelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.BiosApi.GetBiosBootModeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getbiospolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getbiospolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.BiosApi.GetBiosPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getbiosunitbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getbiosunitlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.BiosApi.GetBiosUnitList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "boot",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getbootdevicebootmodebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getbootdevicebootmodelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.BootApi.GetBootDeviceBootModeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getbootprecisionpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getbootprecisionpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.BootApi.GetBootPrecisionPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "compute",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getcomputebladebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getcomputebladelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ComputeApi.GetComputeBladeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getcomputeboardbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getcomputeboardlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ComputeApi.GetComputeBoardList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getcomputephysicalsummarybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getcomputephysicalsummarylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ComputeApi.GetComputePhysicalSummaryList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getcomputerackunitbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getcomputerackunitlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ComputeApi.GetComputeRackUnitList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getcomputeserversettingbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getcomputeserversettinglist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ComputeApi.GetComputeServerSettingList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "cond",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getcondalarmbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getcondalarmlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.CondApi.GetCondAlarmList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getcondhclstatusbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getcondhclstatusdetailbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getcondhclstatusdetaillist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.CondApi.GetCondHclStatusDetailList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getcondhclstatusjobbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getcondhclstatusjoblist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.CondApi.GetCondHclStatusJobList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getcondhclstatuslist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.CondApi.GetCondHclStatusList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "deviceconnector",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getdeviceconnectorpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getdeviceconnectorpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.DeviceconnectorApi.GetDeviceconnectorPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "equipment",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentchassisbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentchassislist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.EquipmentApi.GetEquipmentChassisList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentdevicesummarybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentdevicesummarylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.EquipmentApi.GetEquipmentDeviceSummaryList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentfanbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentfanlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.EquipmentApi.GetEquipmentFanList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentfanmodulebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentfanmodulelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.EquipmentApi.GetEquipmentFanModuleList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentfexbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentfexlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.EquipmentApi.GetEquipmentFexList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentiocardbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentiocardlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.EquipmentApi.GetEquipmentIoCardList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentioexpanderbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentioexpanderlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.EquipmentApi.GetEquipmentIoExpanderList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentlocatorledbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentlocatorledlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.EquipmentApi.GetEquipmentLocatorLedList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentpsubymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentpsulist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.EquipmentApi.GetEquipmentPsuList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentrackenclosurebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentrackenclosurelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.EquipmentApi.GetEquipmentRackEnclosureList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentrackenclosureslotbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentrackenclosureslotlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.EquipmentApi.GetEquipmentRackEnclosureSlotList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentsharediomodulebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentsharediomodulelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.EquipmentApi.GetEquipmentSharedIoModuleList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentswitchcardbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentswitchcardlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.EquipmentApi.GetEquipmentSwitchCardList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentsystemiocontrollerbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmentsystemiocontrollerlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.EquipmentApi.GetEquipmentSystemIoControllerList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmenttpmbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getequipmenttpmlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.EquipmentApi.GetEquipmentTpmList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ether",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getetherphysicalportbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getetherphysicalportlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.EtherApi.GetEtherPhysicalPortList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "externalsite",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getexternalsiteauthorizationbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getexternalsiteauthorizationlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ExternalsiteApi.GetExternalsiteAuthorizationList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "fault",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getfaultinstancebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getfaultinstancelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FaultApi.GetFaultInstanceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "fc",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getfcphysicalportbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getfcphysicalportlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FcApi.GetFcPhysicalPortList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "firmware",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getfirmwaredistributablebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getfirmwaredistributablelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FirmwareApi.GetFirmwareDistributableList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getfirmwaredriverdistributablebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getfirmwaredriverdistributablelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FirmwareApi.GetFirmwareDriverDistributableList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getfirmwareeulabymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getfirmwareeulalist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FirmwareApi.GetFirmwareEulaList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getfirmwarerunningfirmwarebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getfirmwarerunningfirmwarelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FirmwareApi.GetFirmwareRunningFirmwareList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getfirmwareserverconfigurationutilitydistributablebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getfirmwareserverconfigurationutilitydistributablelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FirmwareApi.GetFirmwareServerConfigurationUtilityDistributableList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getfirmwareupgradebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getfirmwareupgradelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FirmwareApi.GetFirmwareUpgradeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getfirmwareupgradestatusbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getfirmwareupgradestatuslist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FirmwareApi.GetFirmwareUpgradeStatusList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "forecast",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getforecastcatalogbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getforecastcataloglist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ForecastApi.GetForecastCatalogList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getforecastdefinitionbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getforecastdefinitionlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ForecastApi.GetForecastDefinitionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getforecastinstancebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getforecastinstancelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ForecastApi.GetForecastInstanceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "graphics",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getgraphicscardbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getgraphicscardlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.GraphicsApi.GetGraphicsCardList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getgraphicscontrollerbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getgraphicscontrollerlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.GraphicsApi.GetGraphicsControllerList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "hcl",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethcldriverimagebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethcldriverimagelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HclApi.GetHclDriverImageList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethclexemptedcatalogbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethclexemptedcataloglist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HclApi.GetHclExemptedCatalogList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethclhyperflexsoftwarecompatibilityinfobymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethclhyperflexsoftwarecompatibilityinfolist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HclApi.GetHclHyperflexSoftwareCompatibilityInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethcloperatingsystembymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethcloperatingsystemlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HclApi.GetHclOperatingSystemList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethcloperatingsystemvendorbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethcloperatingsystemvendorlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HclApi.GetHclOperatingSystemVendorList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethclservicestatusbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethclservicestatuslist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HclApi.GetHclServiceStatusList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "hyperflex",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexalarmbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexalarmlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexAlarmList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexappcatalogbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexappcataloglist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexAppCatalogList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexautosupportpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexautosupportpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexAutoSupportPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexcapabilityinfobymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexcapabilityinfolist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexCapabilityInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexclusterbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexclusterlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexClusterList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexclusternetworkpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexclusternetworkpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexClusterNetworkPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexclusterprofilebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexclusterprofilelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexClusterProfileList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexclusterstoragepolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexclusterstoragepolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexClusterStoragePolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexconfigresultbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexconfigresultentrybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexconfigresultentrylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexConfigResultEntryList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexconfigresultlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexConfigResultList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexextfcstoragepolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexextfcstoragepolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexExtFcStoragePolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexextiscsistoragepolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexextiscsistoragepolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexExtIscsiStoragePolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexfeaturelimitexternalbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexfeaturelimitexternallist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexFeatureLimitExternalList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexfeaturelimitinternalbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexfeaturelimitinternallist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexFeatureLimitInternalList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexhealthbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexhealthlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexHealthList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexhxdpversionbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexhxdpversionlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexHxdpVersionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexlocalcredentialpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexlocalcredentialpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexLocalCredentialPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexnodebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexnodeconfigpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexnodeconfigpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexNodeConfigPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexnodelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexNodeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexnodeprofilebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexnodeprofilelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexNodeProfileList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexproxysettingpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexproxysettingpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexProxySettingPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexserverfirmwareversionbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexserverfirmwareversionlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexServerFirmwareVersionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexservermodelbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexservermodellist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexServerModelList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexsoftwareversionpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexsoftwareversionpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexSoftwareVersionPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexsysconfigpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexsysconfigpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexSysConfigPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexucsmconfigpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexucsmconfigpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexUcsmConfigPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexvcenterconfigpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gethyperflexvcenterconfigpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.GetHyperflexVcenterConfigPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "iaas",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiaasconnectorpackbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiaasconnectorpacklist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IaasApi.GetIaasConnectorPackList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiaasdevicestatusbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiaasdevicestatuslist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IaasApi.GetIaasDeviceStatusList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiaaslicenseinfobymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiaaslicenseinfolist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IaasApi.GetIaasLicenseInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiaasmostruntasksbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiaasmostruntaskslist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IaasApi.GetIaasMostRunTasksList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiaasucsdinfobymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiaasucsdinfolist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IaasApi.GetIaasUcsdInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiaasucsdmanagedinfrabymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiaasucsdmanagedinfralist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IaasApi.GetIaasUcsdManagedInfraList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "iam",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamaccountbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamaccountlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamAccountList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamapikeybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamapikeylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamApiKeyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamappregistrationbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamappregistrationlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamAppRegistrationList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamcertificatebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamcertificatelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamCertificateList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamcertificaterequestbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamcertificaterequestlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamCertificateRequestList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamdomaingroupbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamdomaingrouplist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamDomainGroupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamendpointprivilegebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamendpointprivilegelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamEndPointPrivilegeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamendpointrolebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamendpointrolelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamEndPointRoleList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamendpointuserbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamendpointuserlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamEndPointUserList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamendpointuserpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamendpointuserpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamEndPointUserPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamendpointuserrolebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamendpointuserrolelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamEndPointUserRoleList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamidpbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamidplist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamIdpList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamidpreferencebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamidpreferencelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamIdpReferenceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamldapgroupbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamldapgrouplist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamLdapGroupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamldappolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamldappolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamLdapPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamldapproviderbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamldapproviderlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamLdapProviderList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamoauthtokenbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamoauthtokenlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamOAuthTokenList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiampermissionbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiampermissionlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamPermissionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamprivatekeyspecbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamprivatekeyspeclist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamPrivateKeySpecList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamprivilegebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamprivilegelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamPrivilegeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamprivilegesetbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamprivilegesetlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamPrivilegeSetList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamqualifierbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamqualifierlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamQualifierList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamresourcelimitsbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamresourcelimitslist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamResourceLimitsList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamresourcepermissionbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamresourcepermissionlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamResourcePermissionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamresourcerolesbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamresourceroleslist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamResourceRolesList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamrolebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamrolelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamRoleList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamsecurityholderbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamsecurityholderlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamSecurityHolderList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamserviceproviderbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamserviceproviderlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamServiceProviderList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamsessionbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamsessionlimitsbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamsessionlimitslist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamSessionLimitsList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamsessionlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamSessionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamsystembymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamsystemlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamSystemList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamtrustpointbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamtrustpointlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamTrustPointList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamuserbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamusergroupbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamusergrouplist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamUserGroupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamuserlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamUserList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamuserpreferencebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getiamuserpreferencelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.GetIamUserPreferenceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "infra",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getinfraaccountexperiencebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getinfraaccountexperiencelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.InfraApi.GetInfraAccountExperienceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "inventory",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getinventorydeviceinfobymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getinventorydeviceinfolist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.InventoryApi.GetInventoryDeviceInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getinventorydnmobindingbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getinventorydnmobindinglist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.InventoryApi.GetInventoryDnMoBindingList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getinventorygenericinventorybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getinventorygenericinventoryholderbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getinventorygenericinventoryholderlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.InventoryApi.GetInventoryGenericInventoryHolderList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getinventorygenericinventorylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.InventoryApi.GetInventoryGenericInventoryList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ipmioverlan",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getipmioverlanpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getipmioverlanpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IpmioverlanApi.GetIpmioverlanPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "kvm",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getkvmkvmsessionbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getkvmkvmsessionlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.KvmApi.GetKvmKvmSessionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getkvmpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getkvmpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.KvmApi.GetKvmPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "license",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getlicenseaccountlicensedatabymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getlicenseaccountlicensedatalist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.LicenseApi.GetLicenseAccountLicenseDataList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getlicensecustomeropbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getlicensecustomeroplist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.LicenseApi.GetLicenseCustomerOpList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getlicenselicenseinfobymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getlicenselicenseinfolist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.LicenseApi.GetLicenseLicenseInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getlicensesmartlicensetokenbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getlicensesmartlicensetokenlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.LicenseApi.GetLicenseSmartlicenseTokenList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ls",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getlsserviceprofilebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getlsserviceprofilelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.LsApi.GetLsServiceProfileList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "management",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmanagementcontrollerbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmanagementcontrollerlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ManagementApi.GetManagementControllerList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmanagemententitybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmanagemententitylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ManagementApi.GetManagementEntityList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmanagementinterfacebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmanagementinterfacelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ManagementApi.GetManagementInterfaceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "memory",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmemoryarraybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmemoryarraylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.MemoryApi.GetMemoryArrayList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmemorypersistentmemoryconfigresultbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmemorypersistentmemoryconfigresultlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryConfigResultList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmemorypersistentmemoryconfigurationbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmemorypersistentmemoryconfigurationlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryConfigurationList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmemorypersistentmemorynamespacebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmemorypersistentmemorynamespaceconfigresultbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmemorypersistentmemorynamespaceconfigresultlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryNamespaceConfigResultList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmemorypersistentmemorynamespacelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryNamespaceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmemorypersistentmemorypolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmemorypersistentmemorypolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmemorypersistentmemoryregionbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmemorypersistentmemoryregionlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryRegionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmemorypersistentmemoryunitbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmemorypersistentmemoryunitlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.MemoryApi.GetMemoryPersistentMemoryUnitList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmemoryunitbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmemoryunitlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.MemoryApi.GetMemoryUnitList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "meta",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmetadefinitionbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getmetadefinitionlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.MetaApi.GetMetaDefinitionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "network",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getnetworkelementbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getnetworkelementlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NetworkApi.GetNetworkElementList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getnetworkelementsummarybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getnetworkelementsummarylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NetworkApi.GetNetworkElementSummaryList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "networkconfig",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getnetworkconfigpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getnetworkconfigpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NetworkconfigApi.GetNetworkconfigPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "niaapi",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapiapicccopostbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapiapicccopostlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiaapiApi.GetNiaapiApicCcoPostList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapiapicfieldnoticebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapiapicfieldnoticelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiaapiApi.GetNiaapiApicFieldNoticeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapiapichweolbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapiapichweollist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiaapiApi.GetNiaapiApicHweolList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapiapiclatestmaintainedreleasebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapiapiclatestmaintainedreleaselist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiaapiApi.GetNiaapiApicLatestMaintainedReleaseList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapiapicreleaserecommendbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapiapicreleaserecommendlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiaapiApi.GetNiaapiApicReleaseRecommendList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapiapicsweolbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapiapicsweollist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiaapiApi.GetNiaapiApicSweolList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapidcnmccopostbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapidcnmccopostlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiaapiApi.GetNiaapiDcnmCcoPostList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapidcnmfieldnoticebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapidcnmfieldnoticelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiaapiApi.GetNiaapiDcnmFieldNoticeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapidcnmhweolbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapidcnmhweollist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiaapiApi.GetNiaapiDcnmHweolList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapidcnmlatestmaintainedreleasebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapidcnmlatestmaintainedreleaselist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiaapiApi.GetNiaapiDcnmLatestMaintainedReleaseList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapidcnmreleaserecommendbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapidcnmreleaserecommendlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiaapiApi.GetNiaapiDcnmReleaseRecommendList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapidcnmsweolbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapidcnmsweollist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiaapiApi.GetNiaapiDcnmSweolList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapifiledownloaderbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapifiledownloaderlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiaapiApi.GetNiaapiFileDownloaderList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapiniametadatabymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapiniametadatalist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiaapiApi.GetNiaapiNiaMetadataList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapiversionregexbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniaapiversionregexlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiaapiApi.GetNiaapiVersionRegexList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "niatelemetry",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniatelemetryniainventorybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniatelemetryniainventorylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiatelemetryApi.GetNiatelemetryNiaInventoryList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniatelemetrynialicensestatebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getniatelemetrynialicensestatelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NiatelemetryApi.GetNiatelemetryNiaLicenseStateList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ntp",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getntppolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getntppolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NtpApi.GetNtpPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "organization",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getorganizationorganizationbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getorganizationorganizationlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.OrganizationApi.GetOrganizationOrganizationList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "os",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getoscatalogbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getoscataloglist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.OsApi.GetOsCatalogList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getosconfigurationfilebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getosconfigurationfilelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.OsApi.GetOsConfigurationFileList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getosinstallbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getosinstalllist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.OsApi.GetOsInstallList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "pci",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getpcicoprocessorcardbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getpcicoprocessorcardlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.PciApi.GetPciCoprocessorCardList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getpcidevicebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getpcidevicelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.PciApi.GetPciDeviceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getpcilinkbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getpcilinklist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.PciApi.GetPciLinkList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getpciswitchbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getpciswitchlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.PciApi.GetPciSwitchList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "port",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getportgroupbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getportgrouplist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.PortApi.GetPortGroupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getportsubgroupbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getportsubgrouplist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.PortApi.GetPortSubGroupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "processor",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getprocessorunitbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getprocessorunitlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ProcessorApi.GetProcessorUnitList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "recovery",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getrecoverybackupconfigpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getrecoverybackupconfigpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.RecoveryApi.GetRecoveryBackupConfigPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getrecoverybackupprofilebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getrecoverybackupprofilelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.RecoveryApi.GetRecoveryBackupProfileList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getrecoveryconfigresultbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getrecoveryconfigresultentrybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getrecoveryconfigresultentrylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.RecoveryApi.GetRecoveryConfigResultEntryList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getrecoveryconfigresultlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.RecoveryApi.GetRecoveryConfigResultList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getrecoveryondemandbackupbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getrecoveryondemandbackuplist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.RecoveryApi.GetRecoveryOnDemandBackupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getrecoveryrestorebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getrecoveryrestorelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.RecoveryApi.GetRecoveryRestoreList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getrecoveryscheduleconfigpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getrecoveryscheduleconfigpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.RecoveryApi.GetRecoveryScheduleConfigPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "resource",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getresourcegroupbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getresourcegrouplist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ResourceApi.GetResourceGroupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getresourcegroupmemberbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getresourcegroupmemberlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ResourceApi.GetResourceGroupMemberList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getresourcelicenseresourcecountbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getresourcelicenseresourcecountlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ResourceApi.GetResourceLicenseResourceCountList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getresourcemembershipbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getresourcemembershipholderbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getresourcemembershipholderlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ResourceApi.GetResourceMembershipHolderList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getresourcemembershiplist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ResourceApi.GetResourceMembershipList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "sdcard",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsdcardpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsdcardpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SdcardApi.GetSdcardPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "sdwan",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsdwanprofilebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsdwanprofilelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SdwanApi.GetSdwanProfileList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsdwanrouternodebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsdwanrouternodelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SdwanApi.GetSdwanRouterNodeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsdwanrouterpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsdwanrouterpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SdwanApi.GetSdwanRouterPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsdwanvmanageaccountpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsdwanvmanageaccountpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SdwanApi.GetSdwanVmanageAccountPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "search",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsearchsearchitembymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsearchsearchitemlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SearchApi.GetSearchSearchItemList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsearchtagitembymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsearchtagitemlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SearchApi.GetSearchTagItemList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "security",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsecurityunitbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsecurityunitlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SecurityApi.GetSecurityUnitList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "server",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getserverconfigchangedetailbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getserverconfigchangedetaillist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ServerApi.GetServerConfigChangeDetailList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getserverconfigimportbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getserverconfigimportlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ServerApi.GetServerConfigImportList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getserverconfigresultbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getserverconfigresultentrybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getserverconfigresultentrylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ServerApi.GetServerConfigResultEntryList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getserverconfigresultlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ServerApi.GetServerConfigResultList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getserverprofilebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getserverprofilelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ServerApi.GetServerProfileList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "smtp",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsmtppolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsmtppolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SmtpApi.GetSmtpPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "snmp",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsnmppolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsnmppolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SnmpApi.GetSnmpPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "software",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsoftwarehclmetabymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsoftwarehclmetalist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SoftwareApi.GetSoftwareHclMetaList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsoftwarehyperflexdistributablebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsoftwarehyperflexdistributablelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SoftwareApi.GetSoftwareHyperflexDistributableList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsoftwaresolutiondistributablebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsoftwaresolutiondistributablelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SoftwareApi.GetSoftwareSolutionDistributableList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "softwarerepository",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsoftwarerepositoryauthorizationbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsoftwarerepositoryauthorizationlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SoftwarerepositoryApi.GetSoftwarerepositoryAuthorizationList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsoftwarerepositorycatalogbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsoftwarerepositorycataloglist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SoftwarerepositoryApi.GetSoftwarerepositoryCatalogList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsoftwarerepositoryoperatingsystemfilebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsoftwarerepositoryoperatingsystemfilelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SoftwarerepositoryApi.GetSoftwarerepositoryOperatingSystemFileList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "sol",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsolpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsolpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SolApi.GetSolPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ssh",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsshpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsshpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SshApi.GetSshPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "storage",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragecontrollerbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragecontrollerlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageControllerList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragediskgrouppolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragediskgrouppolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageDiskGroupPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageenclosurebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageenclosurediskbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageenclosuredisklist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageEnclosureDiskList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageenclosurediskslotepbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageenclosuredisksloteplist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageEnclosureDiskSlotEpList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageenclosurelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageEnclosureList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageflexflashcontrollerbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageflexflashcontrollerlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageFlexFlashControllerList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageflexflashcontrollerpropsbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageflexflashcontrollerpropslist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageFlexFlashControllerPropsList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageflexflashphysicaldrivebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageflexflashphysicaldrivelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageFlexFlashPhysicalDriveList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageflexflashvirtualdrivebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageflexflashvirtualdrivelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageFlexFlashVirtualDriveList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageflexutilcontrollerbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageflexutilcontrollerlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageFlexUtilControllerList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageflexutilphysicaldrivebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageflexutilphysicaldrivelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageFlexUtilPhysicalDriveList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageflexutilvirtualdrivebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstorageflexutilvirtualdrivelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageFlexUtilVirtualDriveList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragephysicaldiskbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragephysicaldiskextensionbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragephysicaldiskextensionlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStoragePhysicalDiskExtensionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragephysicaldisklist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStoragePhysicalDiskList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragephysicaldiskusagebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragephysicaldiskusagelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStoragePhysicalDiskUsageList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepurearraybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepurearraylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStoragePureArrayList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepurecontrollerbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepurecontrollerlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStoragePureControllerList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepurediskbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepuredisklist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStoragePureDiskList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepurehostbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepurehostgroupbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepurehostgrouplist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStoragePureHostGroupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepurehostlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStoragePureHostList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepurehostlunbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepurehostlunlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStoragePureHostLunList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepureportbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepureportlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStoragePurePortList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepureprotectiongroupbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepureprotectiongrouplist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStoragePureProtectionGroupList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepureprotectiongroupsnapshotbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepureprotectiongroupsnapshotlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStoragePureProtectionGroupSnapshotList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepurereplicationschedulebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepurereplicationschedulelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStoragePureReplicationScheduleList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepuresnapshotschedulebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepuresnapshotschedulelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStoragePureSnapshotScheduleList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepurevolumebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepurevolumelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStoragePureVolumeList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepurevolumesnapshotbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragepurevolumesnapshotlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStoragePureVolumeSnapshotList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragesasexpanderbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragesasexpanderlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageSasExpanderList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragesasportbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragesasportlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageSasPortList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragestoragepolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragestoragepolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageStoragePolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragevdmemberepbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragevdmembereplist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageVdMemberEpList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragevirtualdrivebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragevirtualdriveextensionbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragevirtualdriveextensionlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageVirtualDriveExtensionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getstoragevirtualdrivelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.GetStorageVirtualDriveList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "syslog",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsyslogpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getsyslogpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SyslogApi.GetSyslogPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "tam",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gettamadvisorycountbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gettamadvisorycountlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.TamApi.GetTamAdvisoryCountList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gettamadvisoryinfobymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gettamadvisoryinfolist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.TamApi.GetTamAdvisoryInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gettamadvisoryinstancebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gettamadvisoryinstancelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.TamApi.GetTamAdvisoryInstanceList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gettamsecurityadvisorybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gettamsecurityadvisorylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.TamApi.GetTamSecurityAdvisoryList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "terminal",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getterminalauditlogbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getterminalauditloglist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.TerminalApi.GetTerminalAuditLogList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "top",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gettopsystembymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "gettopsystemlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.TopApi.GetTopSystemList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "ucsd",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getucsdbackupinfobymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getucsdbackupinfolist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.UcsdApi.GetUcsdBackupInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "virtualization",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvirtualizationvmwareclusterbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvirtualizationvmwareclusterlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VirtualizationApi.GetVirtualizationVmwareClusterList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvirtualizationvmwaredatacenterbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvirtualizationvmwaredatacenterlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VirtualizationApi.GetVirtualizationVmwareDatacenterList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvirtualizationvmwaredatastorebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvirtualizationvmwaredatastorelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VirtualizationApi.GetVirtualizationVmwareDatastoreList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvirtualizationvmwarehostbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvirtualizationvmwarehostlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VirtualizationApi.GetVirtualizationVmwareHostList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvirtualizationvmwarevcenterbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvirtualizationvmwarevcenterlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VirtualizationApi.GetVirtualizationVmwareVcenterList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvirtualizationvmwarevirtualmachinebymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvirtualizationvmwarevirtualmachinelist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VirtualizationApi.GetVirtualizationVmwareVirtualMachineList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "vmedia",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvmediapolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvmediapolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VmediaApi.GetVmediaPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "vnic",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvnicethadapterpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvnicethadapterpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.GetVnicEthAdapterPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvnicethifbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvnicethiflist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.GetVnicEthIfList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvnicethnetworkpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvnicethnetworkpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.GetVnicEthNetworkPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvnicethqospolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvnicethqospolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.GetVnicEthQosPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvnicfcadapterpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvnicfcadapterpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.GetVnicFcAdapterPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvnicfcifbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvnicfciflist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.GetVnicFcIfList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvnicfcnetworkpolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvnicfcnetworkpolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.GetVnicFcNetworkPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvnicfcqospolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvnicfcqospolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.GetVnicFcQosPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvniclanconnectivitypolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvniclanconnectivitypolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.GetVnicLanConnectivityPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvnicsanconnectivitypolicybymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getvnicsanconnectivitypolicylist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.GetVnicSanConnectivityPolicyList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "workflow",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowbatchapiexecutorbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowbatchapiexecutorlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.GetWorkflowBatchApiExecutorList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowbuildtaskmetabymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowbuildtaskmetalist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.GetWorkflowBuildTaskMetaList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowbuildtaskmetaownerbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowbuildtaskmetaownerlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.GetWorkflowBuildTaskMetaOwnerList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowcatalogbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowcataloglist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.GetWorkflowCatalogList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowcustomdatatypedefinitionbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowcustomdatatypedefinitionlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.GetWorkflowCustomDataTypeDefinitionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowpendingdynamicworkflowinfobymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowpendingdynamicworkflowinfolist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.GetWorkflowPendingDynamicWorkflowInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowtaskdefinitionbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowtaskdefinitionlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.GetWorkflowTaskDefinitionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowtaskinfobymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowtaskinfolist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.GetWorkflowTaskInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowtaskmetabymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowtaskmetalist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.GetWorkflowTaskMetaList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowworkflowdefinitionbymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowworkflowdefinitionlist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.GetWorkflowWorkflowDefinitionList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowworkflowinfobymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowworkflowinfolist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.GetWorkflowWorkflowInfoList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowworkflowmetabymoid",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "getworkflowworkflowmetalist",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.GetWorkflowWorkflowMetaList(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					return cmd
				}())

			cmd.AddCommand(
				func() *cobra.Command {
					cmd := &cobra.Command{
						Use: "patch",
					}

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "adapter",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchadapterconfigpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchappliancebackuppolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchappliancecertificatesetting",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchappliancedataexportpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchappliancediagsetting",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchappliancesetupinfo",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchapplianceupgrade",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchapplianceupgradepolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchassetdeviceconfiguration",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchassetdevicecontractinformation",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchassetdeviceregistration",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchassetmanageddevice",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchbiosbootmode",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchbiospolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchbiosunit",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchbootdevicebootmode",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchbootprecisionpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchcomputeblade",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchcomputeboard",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchcomputerackunit",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchcomputeserversetting",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchdeviceconnectorpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchequipmentchassis",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchequipmentfan",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchequipmentfanmodule",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchequipmentfex",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchequipmentiocard",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchequipmentioexpander",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchequipmentlocatorled",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchequipmentpsu",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchequipmentrackenclosure",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchequipmentrackenclosureslot",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchequipmentsharediomodule",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchequipmentswitchcard",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchequipmentsystemiocontroller",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchequipmenttpm",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchetherphysicalport",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchexternalsiteauthorization",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchfaultinstance",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchfcphysicalport",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchfirmwaredistributable",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchfirmwaredriverdistributable",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchfirmwarerunningfirmware",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchfirmwareserverconfigurationutilitydistributable",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchgraphicscard",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchgraphicscontroller",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhclhyperflexsoftwarecompatibilityinfo",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexappcatalog",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexautosupportpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexcapabilityinfo",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexcluster",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexclusternetworkpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexclusterprofile",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexclusterstoragepolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexextfcstoragepolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexextiscsistoragepolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexfeaturelimitexternal",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexfeaturelimitinternal",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexhxdpversion",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexlocalcredentialpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexnodeconfigpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexnodeprofile",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexproxysettingpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexserverfirmwareversion",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexservermodel",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexsoftwareversionpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexsysconfigpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexucsmconfigpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchhyperflexvcenterconfigpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiaasucsdinfo",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamaccount",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamapikey",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamappregistration",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamcertificate",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamcertificaterequest",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamendpointuser",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamendpointuserpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamendpointuserrole",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamidp",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamidpreference",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamldapgroup",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamldappolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamldapprovider",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamlocaluserpassword",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiampermission",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamprivatekeyspec",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamqualifier",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamresourceroles",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamsessionlimits",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamuser",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamusergroup",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchiamuserpreference",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchinfraaccountexperience",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchinventorygenericinventory",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchinventorygenericinventoryholder",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchipmioverlanpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchkvmpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchlicenseaccountlicensedata",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchlicensecustomerop",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchlicenselicenseinfo",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchlicensesmartlicensetoken",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchlsserviceprofile",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchmanagementcontroller",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchmanagemententity",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchmanagementinterface",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchmemoryarray",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchmemorypersistentmemoryconfigresult",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchmemorypersistentmemoryconfiguration",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchmemorypersistentmemorynamespace",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchmemorypersistentmemorynamespaceconfigresult",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchmemorypersistentmemorypolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchmemorypersistentmemoryregion",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchmemorypersistentmemoryunit",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchmemoryunit",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchnetworkelement",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchnetworkconfigpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchntppolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchorganizationorganization",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchpcidevice",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchpcilink",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchpciswitch",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchportgroup",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchportsubgroup",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchprocessorunit",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchrecoverybackupconfigpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchrecoverybackupprofile",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchrecoveryondemandbackup",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchrecoveryscheduleconfigpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchresourcegroup",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchsdcardpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchsdwanprofile",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchsdwanrouternode",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchsdwanrouterpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchsdwanvmanageaccountpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchsecurityunit",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchserverprofile",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchsmtppolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchsnmppolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchsoftwarehclmeta",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchsoftwarehyperflexdistributable",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchsoftwaresolutiondistributable",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchsoftwarerepositoryauthorization",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchsoftwarerepositoryoperatingsystemfile",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchsolpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchsshpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstoragecontroller",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstoragediskgrouppolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstorageenclosure",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstorageenclosuredisk",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstorageenclosurediskslotep",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstorageflexflashcontroller",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstorageflexflashcontrollerprops",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstorageflexflashphysicaldrive",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstorageflexflashvirtualdrive",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstorageflexutilcontroller",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstorageflexutilphysicaldrive",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstorageflexutilvirtualdrive",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstoragephysicaldisk",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstoragephysicaldiskextension",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstoragephysicaldiskusage",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstoragepurearray",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstoragesasexpander",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstoragesasport",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstoragestoragepolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstoragevdmemberep",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstoragevirtualdrive",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchstoragevirtualdriveextension",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchsyslogpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchtamadvisorycount",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchtamadvisoryinfo",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchtamadvisoryinstance",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchtamsecurityadvisory",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchtopsystem",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchvirtualizationvmwarecluster",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchvirtualizationvmwaredatacenter",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchvirtualizationvmwaredatastore",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchvirtualizationvmwarehost",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchvirtualizationvmwarevirtualmachine",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchvmediapolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchvnicethadapterpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchvnicethif",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchvnicethnetworkpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchvnicethqospolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchvnicfcadapterpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchvnicfcif",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchvnicfcnetworkpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchvnicfcqospolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchvniclanconnectivitypolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchvnicsanconnectivitypolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchworkflowbatchapiexecutor",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchworkflowcustomdatatypedefinition",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchworkflowtaskdefinition",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchworkflowtaskinfo",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchworkflowworkflowdefinition",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "patchworkflowworkflowinfo",
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
						Use: "post",
					}

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "adapter",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createadapterconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AdapterApi.CreateAdapterConfigPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateadapterconfigpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createappliancebackup",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ApplianceApi.CreateApplianceBackup(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
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
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateappliancebackuppolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateappliancecertificatesetting",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateappliancedataexportpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateappliancediagsetting",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateappliancesetupinfo",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateapplianceupgrade",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateapplianceupgradepolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createassetdeviceclaim",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.AssetApi.CreateAssetDeviceClaim(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
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
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateassetdeviceconfiguration",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateassetdevicecontractinformation",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateassetdeviceregistration",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateassetmanageddevice",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createbiospolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.BiosApi.CreateBiosPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatebiosbootmode",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatebiospolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatebiosunit",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createbootprecisionpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.BootApi.CreateBootPrecisionPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatebootdevicebootmode",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatebootprecisionpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatecomputeblade",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatecomputeboard",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatecomputerackunit",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatecomputeserversetting",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createdeviceconnectorpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.DeviceconnectorApi.CreateDeviceconnectorPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatedeviceconnectorpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentchassis",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentfan",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentfanmodule",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentfex",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentiocard",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentioexpander",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentlocatorled",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentpsu",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentrackenclosure",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentrackenclosureslot",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentsharediomodule",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentswitchcard",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmentsystemiocontroller",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateequipmenttpm",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateetherphysicalport",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createexternalsiteauthorization",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ExternalsiteApi.CreateExternalsiteAuthorization(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateexternalsiteauthorization",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatefaultinstance",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatefcphysicalport",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createfeedbackfeedbackpost",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FeedbackApi.CreateFeedbackFeedbackPost(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "firmware",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createfirmwaredistributable",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.FirmwareApi.CreateFirmwareDistributable(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
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
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatefirmwaredistributable",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatefirmwaredriverdistributable",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatefirmwarerunningfirmware",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatefirmwareserverconfigurationutilitydistributable",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updategraphicscard",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updategraphicscontroller",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhclcompatibilitystatus",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HclApi.CreateHclCompatibilityStatus(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
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
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehclhyperflexsoftwarecompatibilityinfo",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createhyperflexappcatalog",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.HyperflexApi.CreateHyperflexAppCatalog(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
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
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexappcatalog",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexautosupportpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexcapabilityinfo",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexcluster",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexclusternetworkpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexclusterprofile",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexclusterstoragepolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexextfcstoragepolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexextiscsistoragepolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexfeaturelimitexternal",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexfeaturelimitinternal",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexhxdpversion",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexlocalcredentialpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexnodeconfigpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexnodeprofile",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexproxysettingpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexserverfirmwareversion",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexservermodel",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexsoftwareversionpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexsysconfigpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexucsmconfigpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatehyperflexvcenterconfigpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiaasucsdinfo",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createiamaccount",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IamApi.CreateIamAccount(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
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
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamaccount",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamapikey",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamappregistration",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamcertificate",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamcertificaterequest",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamendpointuser",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamendpointuserpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamendpointuserrole",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamidp",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamidpreference",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamldapgroup",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamldappolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamldapprovider",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamlocaluserpassword",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiampermission",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamprivatekeyspec",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamqualifier",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamresourceroles",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamsessionlimits",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamuser",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamusergroup",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateiamuserpreference",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createinfraaccountexperience",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.InfraApi.CreateInfraAccountExperience(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateinfraaccountexperience",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createinventoryrequest",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.InventoryApi.CreateInventoryRequest(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateinventorygenericinventory",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateinventorygenericinventoryholder",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createipmioverlanpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.IpmioverlanApi.CreateIpmioverlanPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateipmioverlanpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createkvmpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.KvmApi.CreateKvmPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatekvmpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createlicenselicenseinfo",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.LicenseApi.CreateLicenseLicenseInfo(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatelicenseaccountlicensedata",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatelicensecustomerop",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatelicenselicenseinfo",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatelicensesmartlicensetoken",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatelsserviceprofile",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatemanagementcontroller",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatemanagemententity",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatemanagementinterface",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "creatememorypersistentmemorypolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.MemoryApi.CreateMemoryPersistentMemoryPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememoryarray",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememorypersistentmemoryconfigresult",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememorypersistentmemoryconfiguration",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememorypersistentmemorynamespace",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememorypersistentmemorynamespaceconfigresult",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememorypersistentmemorypolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememorypersistentmemoryregion",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememorypersistentmemoryunit",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatememoryunit",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatenetworkelement",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createnetworkconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NetworkconfigApi.CreateNetworkconfigPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatenetworkconfigpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createntppolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.NtpApi.CreateNtpPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatentppolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createorganizationorganization",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.OrganizationApi.CreateOrganizationOrganization(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateorganizationorganization",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createosconfigurationfile",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.OsApi.CreateOsConfigurationFile(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
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
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "pci",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatepcidevice",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatepcilink",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatepciswitch",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateportgroup",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateportsubgroup",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateprocessorunit",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createrecoverybackupconfigpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.RecoveryApi.CreateRecoveryBackupConfigPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
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
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updaterecoverybackupconfigpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updaterecoverybackupprofile",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updaterecoveryondemandbackup",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updaterecoveryscheduleconfigpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createresourcegroup",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ResourceApi.CreateResourceGroup(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateresourcegroup",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsdcardpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SdcardApi.CreateSdcardPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesdcardpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsdwanprofile",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SdwanApi.CreateSdwanProfile(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
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
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesdwanprofile",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesdwanrouternode",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesdwanrouterpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesdwanvmanageaccountpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsearchsuggestitem",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SearchApi.CreateSearchSuggestItem(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "security",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesecurityunit",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createserverconfigimport",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.ServerApi.CreateServerConfigImport(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
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
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateserverprofile",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsmtppolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SmtpApi.CreateSmtpPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesmtppolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsnmppolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SnmpApi.CreateSnmpPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesnmppolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsoftwarehclmeta",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SoftwareApi.CreateSoftwareHclMeta(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
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
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesoftwarehclmeta",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesoftwarehyperflexdistributable",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesoftwaresolutiondistributable",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsoftwarerepositoryauthorization",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SoftwarerepositoryApi.CreateSoftwarerepositoryAuthorization(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
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
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesoftwarerepositoryauthorization",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesoftwarerepositoryoperatingsystemfile",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsolpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SolApi.CreateSolPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesolpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsshpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SshApi.CreateSshPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesshpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createstoragediskgrouppolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.StorageApi.CreateStorageDiskGroupPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
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
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragecontroller",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragediskgrouppolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageenclosure",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageenclosuredisk",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageenclosurediskslotep",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageflexflashcontroller",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageflexflashcontrollerprops",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageflexflashphysicaldrive",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageflexflashvirtualdrive",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageflexutilcontroller",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageflexutilphysicaldrive",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestorageflexutilvirtualdrive",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragephysicaldisk",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragephysicaldiskextension",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragephysicaldiskusage",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragepurearray",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragesasexpander",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragesasport",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragestoragepolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragevdmemberep",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragevirtualdrive",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatestoragevirtualdriveextension",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createsyslogpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.SyslogApi.CreateSyslogPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatesyslogpolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createtamadvisorycount",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.TamApi.CreateTamAdvisoryCount(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
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
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatetamadvisorycount",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatetamadvisoryinfo",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatetamadvisoryinstance",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatetamsecurityadvisory",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createtaskpurescopedinventory",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.TaskApi.CreateTaskPureScopedInventory(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "telemetry",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "querytelemetrytimeseries",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.TelemetryApi.QueryTelemetryTimeSeries(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							return cmd
						}())

					cmd.AddCommand(
						func() *cobra.Command {
							cmd := &cobra.Command{
								Use: "top",
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatetopsystem",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevirtualizationvmwarecluster",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevirtualizationvmwaredatacenter",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevirtualizationvmwaredatastore",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevirtualizationvmwarehost",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevirtualizationvmwarevirtualmachine",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createvmediapolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VmediaApi.CreateVmediaPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevmediapolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createvnicethadapterpolicy",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.VnicApi.CreateVnicEthAdapterPolicy(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
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
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicethadapterpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicethif",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicethnetworkpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicethqospolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicfcadapterpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicfcif",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicfcnetworkpolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicfcqospolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevniclanconnectivitypolicy",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updatevnicsanconnectivitypolicy",
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
											}

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
							}

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "createworkflowbatchapiexecutor",

										Run: func(cmd *cobra.Command, args []string) {

											res, httpResponse, err := client.WorkflowApi.CreateWorkflowBatchApiExecutor(authCtx).Execute()
											resultHandler(res, httpResponse, err)

										},
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
									}

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateworkflowbatchapiexecutor",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateworkflowcustomdatatypedefinition",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateworkflowtaskdefinition",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateworkflowtaskinfo",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateworkflowworkflowdefinition",
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
											}

											return cmd
										}())

									return cmd
								}())

							cmd.AddCommand(
								func() *cobra.Command {
									cmd := &cobra.Command{
										Use: "updateworkflowworkflowinfo",
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

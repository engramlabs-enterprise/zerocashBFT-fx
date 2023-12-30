// Copyright 2015 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

// Contains the geth command usage template and generator.

package main

import (
	"io"
	"sort"

	"github.com/SMEs-Blockchain/zerocash/cmd/utils"
	"github.com/SMEs-Blockchain/zerocash/internal/debug"
	"github.com/SMEs-Blockchain/zerocash/internal/flags"
	"gopkg.in/urfave/cli.v1"
)

// Quorum
var quorumAccountFlagGroup = "QUORUM ACCOUNT"

// End Quorum

// AppHelpFlagGroups is the application flags, grouped by functionality.
var AppHelpFlagGroups = []flags.FlagGroup{
	{
		Name: "Ethereum - Zerocash",
		Flags: []cli.Flag{
			configFileFlag,
			utils.DataDirFlag,
			utils.AncientFlag,
			utils.MinFreeDiskSpaceFlag,
			utils.KeyStoreDirFlag,
			utils.USBFlag,
			utils.SmartCardDaemonPathFlag,
			utils.NetworkIdFlag,
			utils.MainnetFlag,
			utils.GoerliFlag,
			utils.RinkebyFlag,
			utils.YoloV3Flag,
			utils.RopstenFlag,
			utils.SyncModeFlag,
			utils.ExitWhenSyncedFlag,
			utils.GCModeFlag,
			utils.TxLookupLimitFlag,
			utils.EthStatsURLFlag,
			utils.IdentityFlag,
			utils.LightKDFFlag,
			utils.AuthorizationListFlag,
		},
	},
	{
		Name: "Light Client",
		Flags: []cli.Flag{
			utils.LightServeFlag,
			utils.LightIngressFlag,
			utils.LightEgressFlag,
			utils.LightMaxPeersFlag,
			utils.UltraLightServersFlag,
			utils.UltraLightFractionFlag,
			utils.UltraLightOnlyAnnounceFlag,
			utils.LightNoPruneFlag,
			utils.LightNoSyncServeFlag,
		},
	},
	{
		Name: "Developer Chain",
		Flags: []cli.Flag{
			utils.DeveloperFlag,
			utils.DeveloperPeriodFlag,
		},
	},
	{
		Name: "Ethash",
		Flags: []cli.Flag{
			utils.EthashCacheDirFlag,
			utils.EthashCachesInMemoryFlag,
			utils.EthashCachesOnDiskFlag,
			utils.EthashCachesLockMmapFlag,
			utils.EthashDatasetDirFlag,
			utils.EthashDatasetsInMemoryFlag,
			utils.EthashDatasetsOnDiskFlag,
			utils.EthashDatasetsLockMmapFlag,
		},
	},
	{
		Name: "Transaction Pool",
		Flags: []cli.Flag{
			utils.TxPoolLocalsFlag,
			utils.TxPoolNoLocalsFlag,
			utils.TxPoolJournalFlag,
			utils.TxPoolRejournalFlag,
			utils.TxPoolPriceLimitFlag,
			utils.TxPoolPriceBumpFlag,
			utils.TxPoolAccountSlotsFlag,
			utils.TxPoolGlobalSlotsFlag,
			utils.TxPoolAccountQueueFlag,
			utils.TxPoolGlobalQueueFlag,
			utils.TxPoolLifetimeFlag,
		},
	},
	{
		Name: "Performance Tuning",
		Flags: []cli.Flag{
			utils.CacheFlag,
			utils.CacheDatabaseFlag,
			utils.CacheTrieFlag,
			utils.CacheTrieJournalFlag,
			utils.CacheTrieRejournalFlag,
			utils.CacheGCFlag,
			utils.CacheSnapshotFlag,
			utils.CacheNoPrefetchFlag,
			utils.CachePreimagesFlag,
		},
	},
	{
		Name: "Account",
		Flags: []cli.Flag{
			utils.UnlockedAccountFlag,
			utils.PasswordFileFlag,
			utils.ExternalSignerFlag,
			utils.InsecureUnlockAllowedFlag,
		},
	},
	{
		Name: "Api & Console",
		Flags: []cli.Flag{
			utils.IPCDisabledFlag,
			utils.IPCPathFlag,
			utils.HTTPEnabledFlag,
			utils.HTTPListenAddrFlag,
			utils.HTTPPortFlag,
			utils.HTTPApiFlag,
			utils.HTTPPathPrefixFlag,
			utils.HTTPCORSDomainFlag,
			utils.HTTPVirtualHostsFlag,
			utils.WSEnabledFlag,
			utils.WSListenAddrFlag,
			utils.WSPortFlag,
			utils.WSApiFlag,
			utils.WSPathPrefixFlag,
			utils.WSAllowedOriginsFlag,
			utils.GraphQLEnabledFlag,
			utils.GraphQLCORSDomainFlag,
			utils.GraphQLVirtualHostsFlag,
			utils.RPCGlobalGasCapFlag,
			utils.RPCGlobalTxFeeCapFlag,
			utils.AllowUnprotectedTxs,
			utils.JSpathFlag,
			utils.ExecFlag,
			utils.PreloadJSFlag,
			// Quorum
			utils.RPCClientToken,
			utils.RPCClientTLSInsecureSkipVerify,
			utils.RPCClientTLSCert,
			utils.RPCClientTLSCaCert,
			utils.RPCClientTLSCipherSuites,
		},
	},
	{
		Name: "Networking",
		Flags: []cli.Flag{
			utils.BootnodesFlag,
			utils.DNSDiscoveryFlag,
			utils.ListenPortFlag,
			utils.MaxPeersFlag,
			utils.MaxPendingPeersFlag,
			utils.NATFlag,
			utils.NoDiscoverFlag,
			utils.DiscoveryV5Flag,
			utils.NetrestrictFlag,
			utils.NodeKeyFileFlag,
			utils.NodeKeyHexFlag,
		},
	},
	{
		Name: "Miner",
		Flags: []cli.Flag{
			utils.MiningEnabledFlag,
			utils.MinerThreadsFlag,
			utils.MinerNotifyFlag,
			utils.MinerNotifyFullFlag,
			utils.MinerGasPriceFlag,
			utils.MinerGasTargetFlag,
			utils.MinerGasLimitFlag,
			utils.MinerEtherbaseFlag,
			utils.MinerExtraDataFlag,
			utils.MinerRecommitIntervalFlag,
			utils.MinerNoVerfiyFlag,
		},
	},
	{
		Name: "Gas Price Oracle",
		Flags: []cli.Flag{
			utils.GpoBlocksFlag,
			utils.GpoPercentileFlag,
			utils.GpoMaxGasPriceFlag,
		},
	},
	{
		Name: "Virtual Machine",
		Flags: []cli.Flag{
			utils.VMEnableDebugFlag,
			utils.EVMInterpreterFlag,
			utils.EWASMInterpreterFlag,
			// Quorum - timout for calls
			utils.EVMCallTimeOutFlag,
		},
	},
	{
		Name: "Logging & Debugging",
		Flags: append([]cli.Flag{
			utils.FakePoWFlag,
			utils.NoCompactionFlag,
		}, debug.Flags...),
	},
	{
		Name:  "Metrics & Stats",
		Flags: metricsFlags,
	},
	{
		Name: "Aliased (deprecated)",
		Flags: []cli.Flag{
			utils.NoUSBFlag,
			utils.LegacyRPCEnabledFlag,
			utils.LegacyRPCListenAddrFlag,
			utils.LegacyRPCPortFlag,
			utils.LegacyRPCCORSDomainFlag,
			utils.LegacyRPCVirtualHostsFlag,
			utils.LegacyRPCApiFlag,
		},
	},
	// QUORUM
	{
		Name: "ZerocashBFT",
		Flags: []cli.Flag{
			utils.QuorumImmutabilityThreshold,
			utils.EnableNodePermissionFlag,
			utils.PluginSettingsFlag,
			utils.PluginSkipVerifyFlag,
			utils.PluginLocalVerifyFlag,
			utils.PluginPublicKeyFlag,
			utils.AllowedFutureBlockTimeFlag,
			utils.MultitenancyFlag,
			utils.RevertReasonFlag,
			utils.QuorumEnablePrivateTrieCache,
			utils.QuorumEnablePrivacyMarker,
		},
	},
	{
		Name: "Zerocash Client/Server",
		Flags: []cli.Flag{
			utils.QuorumLightServerFlag,
			utils.QuorumLightServerP2PListenPortFlag,
			utils.QuorumLightServerP2PMaxPeersFlag,
			utils.QuorumLightServerP2PNetrestrictFlag,
			utils.QuorumLightServerP2PPermissioningFlag,
			utils.QuorumLightServerP2PPermissioningPrefixFlag,
			utils.QuorumLightClientFlag,
			utils.QuorumLightClientPSIFlag,
			utils.QuorumLightClientTokenEnabledFlag,
			utils.QuorumLightClientTokenValueFlag,
			utils.QuorumLightClientTokenManagementFlag,
			utils.QuorumLightClientRPCTLSFlag,
			utils.QuorumLightClientRPCTLSInsecureSkipVerifyFlag,
			utils.QuorumLightClientRPCTLSCACertFlag,
			utils.QuorumLightClientRPCTLSCertFlag,
			utils.QuorumLightClientRPCTLSKeyFlag,
			utils.QuorumLightClientServerNodeFlag,
			utils.QuorumLightClientServerNodeRPCFlag,
			utils.QuorumLightTLSFlag,
			utils.QuorumLightTLSCertFlag,
			utils.QuorumLightTLSKeyFlag,
			utils.QuorumLightTLSCACertsFlag,
			utils.QuorumLightTLSClientAuthFlag,
			utils.QuorumLightTLSCipherSuitesFlag,
		},
	},
	{
		Name: "Zerocash Private Transcation Manager",
		Flags: []cli.Flag{
			utils.QuorumPTMUnixSocketFlag,
			utils.QuorumPTMUrlFlag,
			utils.QuorumPTMTimeoutFlag,
			utils.QuorumPTMDialTimeoutFlag,
			utils.QuorumPTMHttpIdleTimeoutFlag,
			utils.QuorumPTMHttpWriteBufferSizeFlag,
			utils.QuorumPTMHttpReadBufferSizeFlag,
			utils.QuorumPTMTlsModeFlag,
			utils.QuorumPTMTlsRootCaFlag,
			utils.QuorumPTMTlsClientCertFlag,
			utils.QuorumPTMTlsClientKeyFlag,
			utils.QuorumPTMTlsInsecureSkipVerify,
		},
	},
	{
		Name: quorumAccountFlagGroup,
		Flags: []cli.Flag{
			utils.AccountPluginNewAccountConfigFlag,
		},
	},
	{
		Name: "Raft",
		Flags: []cli.Flag{
			utils.RaftModeFlag,
			utils.RaftBlockTimeFlag,
			utils.RaftJoinExistingFlag,
			utils.RaftPortFlag,
			utils.RaftDNSEnabledFlag,
		},
	},
	{
		Name: "Istanbul",
		Flags: []cli.Flag{
			utils.IstanbulRequestTimeoutFlag,
			utils.IstanbulBlockPeriodFlag,
		},
	},
	// END QUORUM
	{
		Name: "Misc",
		Flags: []cli.Flag{
			utils.SnapshotFlag,
			utils.BloomFilterSizeFlag,
			cli.HelpFlag,
			utils.CatalystFlag,
		},
	},
}

func init() {
	// Override the default app help template
	cli.AppHelpTemplate = flags.AppHelpTemplate

	// Override the default app help printer, but only for the global app help
	originalHelpPrinter := cli.HelpPrinter
	cli.HelpPrinter = func(w io.Writer, tmpl string, data interface{}) {
		if tmpl == flags.AppHelpTemplate {
			// Iterate over all the flags and add any uncategorized ones
			categorized := make(map[string]struct{})
			for _, group := range AppHelpFlagGroups {
				for _, flag := range group.Flags {
					categorized[flag.String()] = struct{}{}
				}
			}
			deprecated := make(map[string]struct{})
			for _, flag := range utils.DeprecatedFlags {
				deprecated[flag.String()] = struct{}{}
			}
			// Only add uncategorized flags if they are not deprecated
			var uncategorized []cli.Flag
			for _, flag := range data.(*cli.App).Flags {
				if _, ok := categorized[flag.String()]; !ok {
					if _, ok := deprecated[flag.String()]; !ok {
						uncategorized = append(uncategorized, flag)
					}
				}
			}
			if len(uncategorized) > 0 {
				// Append all ungategorized options to the misc group
				miscs := len(AppHelpFlagGroups[len(AppHelpFlagGroups)-1].Flags)
				AppHelpFlagGroups[len(AppHelpFlagGroups)-1].Flags = append(AppHelpFlagGroups[len(AppHelpFlagGroups)-1].Flags, uncategorized...)

				// Make sure they are removed afterwards
				defer func() {
					AppHelpFlagGroups[len(AppHelpFlagGroups)-1].Flags = AppHelpFlagGroups[len(AppHelpFlagGroups)-1].Flags[:miscs]
				}()
			}

			// remove the Quorum account options from the main app usage as these should only be used by the geth account sub commands
			for i, group := range AppHelpFlagGroups {
				if group.Name == quorumAccountFlagGroup {
					AppHelpFlagGroups = append(AppHelpFlagGroups[:i], AppHelpFlagGroups[i+1:]...)
				}
			}

			// Render out custom usage screen
			originalHelpPrinter(w, tmpl, flags.HelpData{App: data, FlagGroups: AppHelpFlagGroups})
		} else if tmpl == flags.CommandHelpTemplate {
			// Iterate over all command specific flags and categorize them
			categorized := make(map[string][]cli.Flag)
			for _, flag := range data.(cli.Command).Flags {
				if _, ok := categorized[flag.String()]; !ok {
					categorized[flags.FlagCategory(flag, AppHelpFlagGroups)] = append(categorized[flags.FlagCategory(flag, AppHelpFlagGroups)], flag)
				}
			}

			// sort to get a stable ordering
			sorted := make([]flags.FlagGroup, 0, len(categorized))
			for cat, flgs := range categorized {
				sorted = append(sorted, flags.FlagGroup{Name: cat, Flags: flgs})
			}
			sort.Sort(flags.ByCategory(sorted))

			// add sorted array to data and render with default printer
			originalHelpPrinter(w, tmpl, map[string]interface{}{
				"cmd":              data,
				"categorizedFlags": sorted,
			})
		} else {
			originalHelpPrinter(w, tmpl, data)
		}
	}
}

package cmd

import (
	"context"

	"codis/pkg/p2p"

	"github.com/spf13/cobra"
)

// startPeerCmd starts a "regular" node. This node reaches out to the given bootstrap nodes to learn about the network.
// After discovering the peers on the network, the node can participate in the cryptographic protocols.
func startPeerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "peer",
		Short: "Starts a new peer",
		Long:  `Creates a new peer that connects to bootstrap nodes.`,
		Run: func(c *cobra.Command, args []string) {
			ctx := context.Background()
			peer := p2p.NewPeer(ctx, cfg.Peer.Bootstraps, cfg.Network.PSK, cfg.Peer.KeyID)
			rendezvous := "1ae697d4da3b45469e81ef80dba7ad40"

			if err := peer.AdvertiseConnect(ctx, rendezvous); err != nil {
				logger.Error(err)
			}
			logger.Debug("Peer advertised itself at the %s rendezvous point.", rendezvous)

			if err := peer.StartRPCServer(); err != nil {
				logger.Error(err)
			}
			logger.Debug("Started RPC server.")

			logger.Info("Peer is running! Listening at %s.", peer.ListenAddrs)
			peer.RunUntilCancel()
		},
	}

	// We differentiate a bootstrap node and a peer by whether it has any knowledge of a network outside itself. If the
	// node is to be considered a peer, we must give a bootstrap address as the entrypoint to larger network.
	cmd.PersistentFlags().StringSliceVar(&cfg.Peer.Bootstraps, "bootstraps", []string{}, "bootstrap addrs")
	if err := cmd.MarkPersistentFlagRequired("bootstraps"); err != nil {
		logger.Fatal(err)
	}

	return cmd
}

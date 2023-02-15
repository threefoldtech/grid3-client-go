// Package integration for integration tests
package integration

import (
	"context"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/threefoldtech/grid3-go/deployer"
	"github.com/threefoldtech/grid3-go/workloads"
	"github.com/threefoldtech/zos/pkg/gridtypes"
)

func TestTwoVMsSameNetwork(t *testing.T) {
	tfPluginClient, err := setup()
	assert.NoError(t, err)

	publicKey, privateKey, err := GenerateSSHKeyPair()
	assert.NoError(t, err)

	filter := deployer.NodeFilter{
		CRU:       2,
		SRU:       2,
		MRU:       1,
		Status:    "up",
		PublicIPs: true,
	}
	nodeIDs, err := deployer.FilterNodes(filter, deployer.RMBProxyURLs[tfPluginClient.Network])
	assert.NoError(t, err)
	nodeIDs, err = deployer.FilterNodesWithPublicConfigs(tfPluginClient.SubstrateConn, tfPluginClient.NcPool, nodeIDs)
	assert.NoError(t, err)

	nodeID := nodeIDs[0]

	network := workloads.ZNet{
		Name:        "vmsTestingNetwork",
		Description: "network for testing",
		Nodes:       []uint32{nodeID},
		IPRange: gridtypes.NewIPNet(net.IPNet{
			IP:   net.IPv4(10, 20, 0, 0),
			Mask: net.CIDRMask(16, 32),
		}),
		AddWGAccess: false,
	}

	vm1 := workloads.VM{
		Name:       "vm1",
		Flist:      "https://hub.grid.tf/tf-official-apps/threefoldtech-ubuntu-22.04.flist",
		CPU:        2,
		PublicIP6:  true,
		Planetary:  true,
		Memory:     1024,
		Entrypoint: "/sbin/zinit init",
		EnvVars: map[string]string{
			"SSH_KEY": publicKey,
		},
		IP:          "10.20.2.5",
		NetworkName: network.Name,
	}

	vm2 := workloads.VM{
		Name:       "vm2",
		Flist:      "https://hub.grid.tf/tf-official-apps/threefoldtech-ubuntu-22.04.flist",
		CPU:        2,
		PublicIP6:  true,
		Planetary:  true,
		Memory:     1024,
		Entrypoint: "/sbin/zinit init",
		EnvVars: map[string]string{
			"SSH_KEY": publicKey,
		},
		IP:          "10.20.2.6",
		NetworkName: network.Name,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	err = tfPluginClient.NetworkDeployer.Deploy(ctx, &network)
	assert.NoError(t, err)

	t.Run("public ipv6, yggdrasil and public IPv4", func(t *testing.T) {
		dl := workloads.NewDeployment("vm", nodeID, "", nil, network.Name, nil, nil, []workloads.VM{vm1, vm2}, nil)
		err = tfPluginClient.DeploymentDeployer.Deploy(ctx, &dl)
		assert.NoError(t, err)

		v1, err := tfPluginClient.State.LoadVMFromGrid(nodeID, vm1.Name, dl.Name)
		assert.NoError(t, err)

		v2, err := tfPluginClient.State.LoadVMFromGrid(nodeID, vm2.Name, dl.Name)
		assert.NoError(t, err)

		yggIP1 := v1.YggIP
		yggIP2 := v2.YggIP

		assert.NotEmpty(t, yggIP1)
		assert.NotEmpty(t, yggIP1)

		privateIP1 := v1.IP
		privateIP2 := v2.IP

		publicIP6_1 := strings.Split(v1.ComputedIP6, "/")[0]
		publicIP6_2 := strings.Split(v2.ComputedIP6, "/")[0]

		_, err = RemoteRun("root", yggIP1, "apt install -y netcat", privateKey)
		assert.NoError(t, err)

		_, err = RemoteRun("root", yggIP2, "apt install -y netcat", privateKey)
		assert.NoError(t, err)

		// check privateIP2 from vm1
		_, err = RemoteRun("root", yggIP1, "nc -z "+privateIP2+" 22", privateKey)
		assert.NoError(t, err)

		// check privateIP1 from vm2
		_, err = RemoteRun("root", yggIP2, "nc -z "+privateIP1+" 22", privateKey)
		assert.NoError(t, err)

		// check yggIP2 from vm1
		_, err = RemoteRun("root", yggIP1, "nc -z "+yggIP2+" 22", privateKey)
		assert.NoError(t, err)

		// check yggIP1 from vm2
		_, err = RemoteRun("root", yggIP2, "nc -z "+yggIP1+" 22", privateKey)
		assert.NoError(t, err)

		// check publicIP62 from vm1
		_, err = RemoteRun("root", yggIP1, "nc -z "+publicIP6_2+" 22", privateKey)
		assert.NoError(t, err)

		// check publicIP61 from vm2
		_, err = RemoteRun("root", yggIP2, "nc -z "+publicIP6_1+" 22", privateKey)
		assert.NoError(t, err)

		// cancel all
		err = tfPluginClient.DeploymentDeployer.Cancel(ctx, &dl)
		assert.NoError(t, err)

		err = tfPluginClient.NetworkDeployer.Cancel(ctx, &network)
		assert.NoError(t, err)
	})
}

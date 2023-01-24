// Package workloads includes workloads types (vm, zdb, qsfs, public IP, gateway name, gateway fqdn, disk)
package workloads

import (
	"github.com/threefoldtech/substrate-client"
	"github.com/threefoldtech/zos/pkg/gridtypes"
)

// Deployment struct
type Deployment struct {
	Name             string
	NodeID           uint32
	SolutionType     string
	SolutionProvider *uint64
	NetworkName      string
	Disks            []Disk
	Zdbs             []ZDB
	Vms              []VM
	Qsfs             []QSFS
}

// NewDeployment generates a new deployment
func NewDeployment(twin uint32, workloads []gridtypes.Workload) gridtypes.Deployment {
	return gridtypes.Deployment{
		Version: 0,
		TwinID:  twin, //LocalTwin,
		// this contract id must match the one on substrate
		Workloads: workloads,
		SignatureRequirement: gridtypes.SignatureRequirement{
			WeightRequired: 1,
			Requests: []gridtypes.SignatureRequest{
				{
					TwinID: twin,
					Weight: 1,
				},
			},
		},
	}
}

// GatewayWorkloadGenerator is an interface for a gateway workload generator
type GatewayWorkloadGenerator interface {
	ZosWorkload() gridtypes.Workload
}

// NewDeploymentWithGateway generates a new deployment with a gateway workload
func NewDeploymentWithGateway(identity substrate.Identity, twinID uint32, version uint32, gw GatewayWorkloadGenerator) (gridtypes.Deployment, error) {
	dl := NewDeployment(twinID, []gridtypes.Workload{})
	dl.Version = version

	dl.Workloads = append(dl.Workloads, gw.ZosWorkload())
	dl.Workloads[0].Version = version

	err := dl.Sign(twinID, identity)
	if err != nil {
		return gridtypes.Deployment{}, err
	}

	return dl, nil
}

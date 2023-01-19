// Package workloads includes workloads types (vm, zdb, qsfs, public IP, gateway name, gateway fqdn, disk)
package workloads

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/threefoldtech/grid3-go/deployer"
	"github.com/threefoldtech/zos/pkg/gridtypes"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
)

// ZDB workload struct
type ZDB struct {
	Name        string
	Password    string
	Public      bool
	Size        int
	Description string
	Mode        string
	IPs         []string
	Port        uint32
	Namespace   string
}

// NewZDBFromSchema converts a map including zdb data to a zdb struct
func NewZDBFromSchema(zdb map[string]interface{}) ZDB {
	ips := zdb["ips"].([]string)

	return ZDB{
		Name:        zdb["name"].(string),
		Size:        zdb["size"].(int),
		Description: zdb["description"].(string),
		Password:    zdb["password"].(string),
		Public:      zdb["public"].(bool),
		Mode:        zdb["mode"].(string),
		IPs:         ips,
		Port:        uint32(zdb["port"].(int)),
		Namespace:   zdb["namespace"].(string),
	}
}

// NewZDBFromWorkload generates a new zdb from a workload
func NewZDBFromWorkload(wl *gridtypes.Workload) (ZDB, error) {
	dataI, err := wl.WorkloadData()
	if err != nil {
		return ZDB{}, errors.Wrap(err, "failed to get workload data")
	}

	data, ok := dataI.(*zos.ZDB)
	if !ok {
		return ZDB{}, errors.New("couldn't cast workload data")
	}

	var result zos.ZDBResult

	if len(wl.Result.Data) > 0 {
		if err := json.Unmarshal(wl.Result.Data, &result); err != nil {
			return ZDB{}, errors.Wrap(err, "failed to get zdb result")
		}
	}

	return ZDB{
		Name:        wl.Name.String(),
		Description: wl.Description,
		Password:    data.Password,
		Public:      data.Public,
		Size:        int(data.Size / gridtypes.Gigabyte),
		Mode:        data.Mode.String(),
		IPs:         result.IPs,
		Port:        uint32(result.Port),
		Namespace:   result.Namespace,
	}, nil
}

// GetName returns zdb name
func (z *ZDB) GetName() string {
	return z.Name
}

// GenerateZDBWorkload generates a zdb workload
func (z *ZDB) GenerateZDBWorkload() gridtypes.Workload {
	workload := gridtypes.Workload{
		Name:        gridtypes.Name(z.Name),
		Type:        zos.ZDBType,
		Description: z.Description,
		Version:     0,
		Data: gridtypes.MustMarshal(zos.ZDB{
			Size:     gridtypes.Unit(z.Size) * gridtypes.Gigabyte,
			Mode:     zos.ZDBMode(z.Mode),
			Password: z.Password,
			Public:   z.Public,
		}),
	}
	return workload
}

// Dictify converts a zdb to a map(dict) object
func (z *ZDB) Dictify() map[string]interface{} {
	res := make(map[string]interface{})
	res["name"] = z.Name
	res["description"] = z.Description
	res["size"] = z.Size
	res["mode"] = z.Mode
	res["ips"] = z.IPs
	res["namespace"] = z.Namespace
	res["port"] = int(z.Port)
	res["password"] = z.Password
	res["public"] = z.Public
	return res
}

// GenerateWorkloadFromZDB generates a workload from a zdb
func (z *ZDB) GenerateWorkloadFromZDB() (gridtypes.Workload, error) {
	return gridtypes.Workload{
		Name:        gridtypes.Name(z.Name),
		Type:        zos.ZDBType,
		Description: z.Description,
		Version:     0,
		Data: gridtypes.MustMarshal(zos.ZDB{
			Size:     gridtypes.Unit(z.Size) * gridtypes.Gigabyte,
			Mode:     zos.ZDBMode(z.Mode),
			Password: z.Password,
			Public:   z.Public,
		}),
	}, nil
}

// Stage for staging workloads
func (z *ZDB) Stage(manager deployer.DeploymentManager, NodeId uint32) error {
	workloadsMap := map[uint32][]gridtypes.Workload{}
	workloads := make([]gridtypes.Workload, 0)

	workload, err := z.GenerateWorkloadFromZDB()
	if err != nil {
		return err
	}

	workloads = append(workloads, workload)
	workloadsMap[NodeId] = workloads

	err = manager.SetWorkloads(workloadsMap)
	return err
}

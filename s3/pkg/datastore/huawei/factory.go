package hws

import (
	"github.com/opensds/crsytal/api/pkg/utils/obs"
	"github.com/opensds/crsytal/backend/pkg/utils/constants"
	backendpb "github.com/opensds/crsytal/backend/proto"
	"github.com/opensds/crsytal/s3/pkg/datastore/driver"
)

type HWObsDriverFactory struct {
}

func (cdf *HWObsDriverFactory) CreateDriver(backend *backendpb.BackendDetail) (driver.StorageDriver, error) {
	endpoint := backend.Endpoint
	AccessKeyID := backend.Access
	AccessKeySecret := backend.Security

	client, err := obs.New(AccessKeyID, AccessKeySecret, endpoint)
	if err != nil {
		return nil, err
	}

	adap := &OBSAdapter{backend: backend, client: client}

	return adap, nil
}

func init() {
	driver.RegisterDriverFactory(constants.BackendTypeObs, &HWObsDriverFactory{})
	driver.RegisterDriverFactory(constants.BackendFusionStorage, &HWObsDriverFactory{})
}

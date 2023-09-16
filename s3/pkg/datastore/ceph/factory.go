package ceph

import (
	"github.com/webrtcn/s3client"

	"github.com/opensds/crystal/backend/pkg/utils/constants"
	backendpb "github.com/opensds/crystal/backend/proto"
	"github.com/opensds/crystal/s3/pkg/datastore/driver"
)

type CephS3DriverFactory struct {
}

func (cdf *CephS3DriverFactory) CreateDriver(backend *backendpb.BackendDetail) (driver.StorageDriver, error) {
	endpoint := backend.Endpoint
	AccessKeyID := backend.Access
	AccessKeySecret := backend.Security
	sess := s3client.NewClient(endpoint, AccessKeyID, AccessKeySecret)
	adap := &CephAdapter{backend: backend, session: sess}

	return adap, nil
}

func init() {
	driver.RegisterDriverFactory(constants.BackendTypeCeph, &CephS3DriverFactory{})
}

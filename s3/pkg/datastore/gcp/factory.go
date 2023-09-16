package gcp

import (
	"github.com/webrtcn/s3client"

	"github.com/opensds/crystal/backend/pkg/utils/constants"
	backendpb "github.com/opensds/crystal/backend/proto"
	"github.com/opensds/crystal/s3/pkg/datastore/driver"
)

type GcsDriverFactory struct {
}

func (cdf *GcsDriverFactory) CreateDriver(backend *backendpb.BackendDetail) (driver.StorageDriver, error) {
	endpoint := backend.Endpoint
	AccessKeyID := backend.Access
	AccessKeySecret := backend.Security
	sess := s3client.NewClient(endpoint, AccessKeyID, AccessKeySecret)
	adap := &GcsAdapter{backend: backend, session: sess}

	return adap, nil
}

func init() {
	driver.RegisterDriverFactory(constants.BackendTypeGcs, &GcsDriverFactory{})
}

package ibmcos

import (
	"github.com/opensds/crystal/backend/pkg/utils/constants"
	backendpb "github.com/opensds/crystal/backend/proto"
	"github.com/opensds/crystal/s3/pkg/datastore/aws"
	"github.com/opensds/crystal/s3/pkg/datastore/driver"
)

type IBMCOSDriverFactory struct {
}

func (factory *IBMCOSDriverFactory) CreateDriver(backend *backendpb.BackendDetail) (driver.StorageDriver, error) {
	awss3Fac := &aws.AwsS3DriverFactory{}
	return awss3Fac.CreateDriver(backend)
}

func init() {
	driver.RegisterDriverFactory(constants.BackendTypeIBMCos, &IBMCOSDriverFactory{})
}

package ibmcos

import (
	"github.com/opensds/crsytal/backend/pkg/utils/constants"
	backendpb "github.com/opensds/crsytal/backend/proto"
	"github.com/opensds/crsytal/s3/pkg/datastore/aws"
	"github.com/opensds/crsytal/s3/pkg/datastore/driver"
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

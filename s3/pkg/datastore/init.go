package datastore

import (
	_ "github.com/opensds/crsytal/s3/pkg/datastore/alibaba"
	_ "github.com/opensds/crsytal/s3/pkg/datastore/aws"
	_ "github.com/opensds/crsytal/s3/pkg/datastore/azure"
	_ "github.com/opensds/crsytal/s3/pkg/datastore/ceph"
	_ "github.com/opensds/crsytal/s3/pkg/datastore/gcp"
	_ "github.com/opensds/crsytal/s3/pkg/datastore/huawei"
	_ "github.com/opensds/crsytal/s3/pkg/datastore/ibm"
	_ "github.com/opensds/crsytal/s3/pkg/datastore/sony"
	_ "github.com/opensds/crsytal/s3/pkg/datastore/yig"
)

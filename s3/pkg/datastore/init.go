package datastore

import (
	_ "github.com/opensds/crystal/s3/pkg/datastore/aws"
	_ "github.com/opensds/crystal/s3/pkg/datastore/azure"
	_ "github.com/opensds/crystal/s3/pkg/datastore/gcp"
	_ "github.com/opensds/crystal/s3/pkg/datastore/ibm"
	_ "github.com/opensds/crystal/s3/pkg/datastore/yig"
)

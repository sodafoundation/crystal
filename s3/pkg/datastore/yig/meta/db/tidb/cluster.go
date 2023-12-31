package tidb

import (
	"github.com/opensds/crystal/s3/pkg/datastore/yig/meta/types"
)

//cluster
func (t *Tidb) GetCluster(fsid, pool string) (cluster types.Cluster, err error) {
	sqltext := "select fsid,pool,weight from cluster where fsid=? and pool=?"
	err = t.DB.QueryRow(sqltext, fsid, pool).Scan(
		&cluster.Fsid,
		&cluster.Pool,
		&cluster.Weight,
	)
	return
}

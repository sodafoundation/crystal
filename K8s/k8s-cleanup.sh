#Delete the services
kubectl delete service  zookeeper -n soda-crystal
kubectl delete service  tidb -n soda-crystal
kubectl delete service  redis -n soda-crystal
kubectl delete service  kafka -n soda-crystal
kubectl delete service  api -n soda-crystal
kubectl delete service  s3api -n soda-crystal
kubectl delete service  s3 -n soda-crystal
kubectl delete service  backend -n soda-crystal
kubectl delete service  block -n soda-crystal
kubectl delete service  file -n soda-crystal
kubectl delete service  datamover -n soda-crystal
kubectl delete service  dataflow -n soda-crystal

# Delete the deployments
kubectl delete deployment zookeeper -n soda-crystal
kubectl delete deployment redis -n soda-crystal
kubectl delete deployment tidb -n soda-crystal
kubectl delete deployment kafka -n soda-crystal
kubectl delete deployment block -n soda-crystal
kubectl delete deployment file -n soda-crystal
kubectl delete deployment backend -n soda-crystal
kubectl delete deployment s3api -n soda-crystal
kubectl delete deployment s3 -n soda-crystal
kubectl delete deployment api -n soda-crystal
kubectl delete deployment datamover -n soda-crystal
kubectl delete deployment dataflow -n soda-crystal

# Delete config maps
kubectl delete configmap multicloud-config -n soda-crystal
kubectl delete configmap tidb-config -n soda-crystal
kubectl delete configmap s3-config -n soda-crystal
kubectl delete configmap tidb-sql -n soda-crystal
kubectl delete configmap s3-sql -n soda-crystal
kubectl delete configmap yig-sql -n soda-crystal




kubectl delete svc mongodb-0-service -n soda-crystal
kubectl delete svc mongodb-1-service -n soda-crystal
kubectl delete svc mongodb-2-service -n soda-crystal

kubectl delete deployment  mongodb-0 -n soda-crystal
kubectl delete deployment  mongodb-1 -n soda-crystal
kubectl delete deployment  mongodb-2 -n soda-crystal

# Delete pv, pvc
kubectl patch pvc mongo-0-pv-claim -p '{"metadata":{"finalizers": []}}' --type=merge -n soda-crystal
kubectl patch pvc mongo-1-pv-claim -p '{"metadata":{"finalizers": []}}' --type=merge -n soda-crystal
kubectl patch pvc mongo-2-pv-claim -p '{"metadata":{"finalizers": []}}' --type=merge -n soda-crystal

kubectl delete pvc mongo-0-pv-claim -n soda-crystal
kubectl delete pvc mongo-1-pv-claim -n soda-crystal
kubectl delete pvc mongo-2-pv-claim -n soda-crystal

kubectl delete pv mongo-0-pv-volume -n soda-crystal
kubectl delete pv mongo-1-pv-volume -n soda-crystal
kubectl delete pv mongo-2-pv-volume -n soda-crystal

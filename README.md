This is a tool to generate dummy pods in order to overload our system!

# TO CREATE 50 PODS
go run create-dummy-containers.go

# TO REMOVE ALL PODS
kubectl delete --selector app=demo

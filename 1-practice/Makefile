clustername="1-practice"

create-cluster:
	@sh ./create-cluster.sh 

delete-cluster:
	@kind delete cluster --name $(clustername)

build-webapp:
	@docker build -t txuna/kind-webapp:latest -f cmd/webapp/Dockerfile .
	@docker push txuna/kind-webapp

build-monitor:
	@docker build -t txuna/kind-monitor:latest -f cmd/monitor/Dockerfile .
	@docker push txuna/kind-monitor

deploy-webapp:
	@helm upgrade --install -f values/webapp/local.yaml webapp install/helm/webapp

deploy-monitor:
	@helm upgrade --install -f values/monitor/local.yaml monitor install/helm/monitor

remove-webapp:
	@helm uninstall webapp

remove-monitor:
	@helm uninstall monitor
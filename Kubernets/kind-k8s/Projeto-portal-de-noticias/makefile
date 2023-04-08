up:
#	=> Sistema de noticias
### Service
	@kubectl apply -f ./sistema/svc-sistema-noticias.yaml
### Pod
# @kubectl apply -f ./sistema/sistema-noticias.yaml
### Deploy
	@kubectl apply -f ./sistema/deployment.yaml

# => Portal de noticias
### Service
	@kubectl apply -f ./portal/svc-portal-noticias.yaml
### Pod
# @kubectl apply -f ./portal/portal-noticias.yaml
### Deploy
	@kubectl apply -f ./portal/deployment.yaml

# => Database portal de noticias
### ConfigMap
	@kubectl apply -f ./database/db-configmap.yaml
### Service
	@kubectl apply -f ./database/svc-db-noticias.yaml
### Pod
#	@kubectl apply -f ./database/db-noticias.yaml
### Deploy
	@kubectl apply -f ./database/deployment.yaml


delete-all:
# Deletando todos os serviços
	@kubectl delete services --all

# Deletando todos os deployments
	@kubectl delete deployments --all
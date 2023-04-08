## Criaremos um cluster kubernetes com:
  - 1 pods
  - 1 service para o pod
  - mapeando configurando as portas do cluster para usar o localhost

### Config
  Nesse configuramos o mapeamento de portas entre o NODE
  e o host local.
  [localhos:8080] -> [nodeIP:30000]

  ```yaml
  kind: Cluster
  apiVersion: kind.x-k8s.io/v1alpha4
  nodes:
  - role: control-plane
    extraPortMappings:
    - containerPort: 30000
      hostPort: 8080

  ```
### Service
  Utilizamo service na frente dos nossos pods
  para que possamos encaminhas as requisições que chegam.
  Como os pods são efêmeros, o service mapeia os pods pela
  label, e encaminha a requisição.

  Como local não temos um loadbalace vamo acesso o serviço
  através no NODE que tem o ip exposto fora do cluste.
  Utilizamos assim um service do tipo NodePort.
  [nodeIP:30000] -> [service:8080] -> [pod:80]

  ```yaml
  apiVersion: v1
  kind: Service
  metadata:
    name: svc-nginx
  spec:
    type: NodePort
    selector:
      app: nginx-pod
    ports:
    - port: 3001
      targetPort: 80
      nodePort: 30000
  ```

### Pod
  O pod vai ser onde o container da nossa aplicação
  ira rodar de fato.
  Nele definimos qual será imagem, recursos e a porta que será exposta.
  > Atenção a chave da label para que o service faça o mapeamente corretamente. [app:nome] e não [name:nome]

  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: nginx-pod
    labels:
      app: nginx-pod
  spec:
    containers:
    - name: nginx-pod
      image: nginx:latest
      resources:
        limits:
          memory: "100Mi"
          cpu: "100m"
      ports:
        - containerPort: 80
  ```

### Execução
  Criação do cluster com a config.
  ```bash
  $ kind create cluster --name my-cluster --config=cluster-config.yaml
  ```

  Criação do pod e do service.
  ```bash
  $ kubectl apply -f pod.yaml
  $ kubectl apply -f service.yaml
  ```

  Como o roteamento de portas ficou.
  ```md
   [localhos:8080] -> [nodeIP:30000] -> [service:8080] -> [pod:80]
  ```

  Acessando a recurso.
  ```bash
  $ curl localhost:8080
  ```
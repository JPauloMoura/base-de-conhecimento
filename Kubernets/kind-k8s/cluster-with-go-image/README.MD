# buidar um arquivo GO

# buldar uma imagem docker

# Carregar uma imagem docker no k8s
  $ kind load docker-image api --name cluste-go

# Para que não de erro no pull da imagem use isso no yaml do pod:
spec:
  containers:
  - name: api-go-pod
    image: api:latest
    # esse é o atributo
    imagePullPolicy: Never

# outra forma é criando a imagem com um formato corretor ex:
$ docker build -t api:0.1.0 .
$ kind load docker-image api:0.1.0

spec:
  containers:
  - name: sleepy
    image: sleepy:0.1.0

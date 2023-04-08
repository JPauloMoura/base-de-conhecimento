# =========Docker e Portainer do zero Containers em desenvolvimento================

> Docker
 - conteines são processos que vão rodar nossas aplicaçoes de forma isolada
 - imagens são um descritivo de tudo que vai estar em um container
 - para um conteiner ficar de pé ele precisa deixar um processo em execução

 - entrando no terminal de um container quer esteja rodando
   $ docker exec -it <nome_container> bash
 
 - toda vez que executamos um docker run é criada uma NOVA nome_imagem
 - as imagens são imutaveis

 Imagem vs Container
 - imagine o conceito de POO 
 - quando rodamos um "docker run Image" estamas instaciando uma "classe"
 - subido assim um "container" / criado um "objeto"

> criando sua propria imagen docker
 - crie um arquivo "Dockerfile" 
 ex: criando uma imagem modificada do nginx
 ````
 FROM nginx:latest

 COPY html/* /usr/share/nginx/html
 ```
 $ docker build -t nome/image ./diretorio/Dockerfile

> subir imagem no dockerhub
  $ docker push nome/imagem

> Docker Compose
  - crie um arquivo "docker-compose.yaml" com as configuraçoes necessarias
  - e execute na pastar do arquivo
  $ docker-compose up -d

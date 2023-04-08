### Docker é uma coleção de tecnologias que facilita o deploy e execução de aplicações:
  Docker Engine:| Faz intermédio entre containes e SO.|
  --- | -
  Docker Compose:| Orquestração de multiplos containers.|
  Docker Swarm:| Colocar multiplos docker engines para tarbalhar juntos em um cluster.|
  Docker Hub:| Repositorio de imagens de containers.|
  Docker Machine:| Permite gerenciar o docker em um host virtual.|

### Hello Docker!
Quando executamos o comando `docker run hello-world`, estamos dizendo para o 
docker criar um container com a imagem do hello-world, caso não tenha essa imagem na sua máquina
ele ira buscar no [Docker Hub](https://hub.docker.com/search?type=image):</br>
  `$ docker run  hello-world`

> Dica para brincar com o docker -> https://labs.play-with-docker.com/






















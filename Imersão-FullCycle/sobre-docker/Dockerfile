# imagem base
FROM nginx:latest
# instalações podem ser executadas usando RUN
RUN ap-get update && apt-get install -y vim
# copia nosso local para o caminho no container
COPY index.html /usr/share/nginx/html/
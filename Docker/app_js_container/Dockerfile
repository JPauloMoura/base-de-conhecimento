FROM node
# cria diretorio de trabalho
WORKDIR /app_js_container
# realizamos a copia do package separada, para que possamos utilizar cache
COPY package.json .
RUN npm install
COPY . .
EXPOSE 3000
# comando usado para execução do container
CMD [ "npm", "start" ]
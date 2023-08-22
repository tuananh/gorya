# pull official base image
FROM node:current-alpine3.18

ENV NODE_OPTIONS=--openssl-legacy-provider
# set working directory
WORKDIR /app

# install app dependencies
#copies package.json and package-lock.json to Docker environment
COPY client/package.json ./
RUN npm install

# Installs all node packages

# Copies everything over to Docker environment
COPY ./client ./

CMD yarn start

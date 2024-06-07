FROM node:20.11.1-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install

RUN npm build

ENTRYPOINT ["npm", "start"]


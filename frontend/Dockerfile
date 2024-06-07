FROM node:20.11.1-alpine AS builder

WORKDIR /app

COPY package*.json .

RUN npm install

COPY . .

ARG API_URL

RUN API_URL=$API_URL npm run build

FROM node:20.11.1-alpine

WORKDIR /app

COPY --from=builder /app .

ENTRYPOINT ["npm", "run", "start"]


FROM node:20.11.1-alpine AS builder

WORKDIR /app

COPY package*.json .

RUN npm install

COPY . .

ARG API_URL
ENV API_URL=$API_URL

RUN npm run build

FROM node:20.11.1-alpine

WORKDIR /app

COPY --from=builder /app .

ENTRYPOINT ["npm", "run", "start"]


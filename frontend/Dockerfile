FROM node:20.11.1-alpine AS builder

WORKDIR /app

COPY package*.json .

RUN npm install

COPY . .

ENTRYPOINT ["npm", "run", "build"]

FROM node:20.11.1-alpine

WORKDIR /app

COPY --from=builder /app .

ENTRYPOINT ["npm", "run", "start"]


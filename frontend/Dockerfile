FROM node:20.11.1-alpine AS builder

WORKDIR /app

COPY package*.json .

RUN npm install

COPY . .

ENTRYPOINT ["npm run build"]

FROM node:20.11.1-alpine

COPY --from=builder /app /app

ENTRYPOINT ["npm run start"]


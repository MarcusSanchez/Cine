FROM node:20.11.1-alpine AS builder

WORKDIR /app

COPY package*.json .

RUN npm install

COPY . .

CMD ["npm", "run", "build"]

FROM node:20.11.1-alpine

WORKDIR /app

COPY --from=builder /app .

ENTRYPOINT ["npm", "run", "start"]


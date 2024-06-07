FROM node:20.11.1-alpine AS builder

WORKDIR /app

COPY package*.json .

RUN npm install

COPY . .

CMD ["npm run build"]

FROM node:20.11.1-alpine

COPY --from=builder /app /app

CMD ["npm run start"]


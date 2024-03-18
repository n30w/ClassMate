FROM node:latest

WORKDIR /app
COPY package*.json ./
RUN corepack enable
RUN yarn
COPY . .
EXPOSE 3000
CMD yarn dev

FROM node:13.12.0-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm config set registry http://registry.npmjs.org/
RUN npm install
COPY . .

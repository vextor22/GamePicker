FROM node:latest as build-stage
WORKDIR /app
COPY ./gamepicker/package*.json ./
RUN npm install
COPY ./gamepicker .
RUN npm run build

FROM nginx as production-stage
RUN mkdir /app
COPY --from=build-stage /app/dist /app
COPY ./proxy/nginx.conf /etc/nginx/nginx.conf
RUN chown nginx -R /app

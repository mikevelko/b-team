FROM node:10-alpine as build-step

ARG PUBLIC_URL
ENV PUBLIC_URL=$PUBLIC_URL

RUN mkdir /app
WORKDIR /app
COPY package.json /app
RUN npm install
COPY . /app
RUN npm run build

FROM nginx:1.17.1-alpine

COPY --from=build-step /app/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

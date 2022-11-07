FROM node:latest as build

WORKDIR /app
COPY . /app
RUN npm install && npm run build

FROM nginx

COPY --from=build /app/dist/client /usr/share/nginx/html


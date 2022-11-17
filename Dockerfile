FROM ghcr.io/greboid/dockerfiles/alpine:latest as build
RUN apk add npm
WORKDIR /app
COPY . /app
RUN npm install && npm run build

FROM ghcr.io/greboid/dockerfiles/alpine:latest
RUN apk add nodejs-current
COPY --from=build /app/dist /app
COPY --from=build /app/package.json /app/package.json
CMD ["node", "/app/index.js"]

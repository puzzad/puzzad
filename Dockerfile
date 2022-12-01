FROM ghcr.io/greboid/dockerfiles/alpine:latest as build
RUN apk add npm curl
WORKDIR /app
COPY . /app
RUN npm install && npm run build
WORKDIR /sws
RUN curl -Ss -o sws.tgz -L $(curl https://api.github.com/repos/static-web-server/static-web-server/releases/latest | grep browser_download_url | grep x86_64-unknown-linux-musl | cut -d '"' -f4); \
    tar xf sws.tgz --strip-components 1;

FROM ghcr.io/greboid/dockerfiles/alpine:latest
COPY --from=build /sws/static-web-server /sws
COPY --from=build /app/build /app
COPY --from=build /app/package.json /app/package.json
WORKDIR /app
CMD ["/sws", "--port", "8080", "--root", "/app", "--page-fallback", "/app/index.html"]

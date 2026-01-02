FROM oven/bun:1.3.5 AS build
WORKDIR /app
COPY . /app
RUN bun install && bun run build

FROM ghcr.io/static-web-server/static-web-server:2.40.1
COPY --from=build /app/build /app
COPY --from=build /app/package.json /app/package.json
WORKDIR /app
EXPOSE 3000
CMD ["--port", "3000", "--root", "/app", "--page-fallback", "/app/200.html"]

FROM ghcr.io/greboid/dockerfiles/alpine:latest as build
RUN apk add npm
WORKDIR /app
COPY . /app
RUN npm install && npm run build

FROM ghcr.io/greboid/dockerfiles/alpine:latest
RUN apk add nodejs-current
COPY --from=build /app/dist /app
COPY --from=build /app/package.json /app/package.json
RUN echo "" >> /app/index.ts
RUN echo "function handleExit(signal) {" >> /app/index.js
RUN echo "console.log('Terminating.')" >> /app/index.js
RUN echo "process.exit(0);" >> /app/index.js
RUN echo "}" >> /app/index.js
RUN echo "process.on('SIGINT', handleExit);" >> /app/index.js
RUN echo "process.on('SIGQUIT', handleExit);" >> /app/index.js
RUN echo "process.on('SIGTERM', handleExit);" >> /app/index.js
WORKDIR /app
CMD ["node", "index.js"]

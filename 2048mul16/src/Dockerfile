FROM node:18 as build-stage

WORKDIR /app

COPY src /app/

RUN rm -rf node_modules && \ 
    npm install -g pnpm && \
    pnpm install && \ 
    pnpm run build

FROM nginx:1.25.3-alpine

COPY --from=build-stage /app/dist /usr/share/nginx/html

EXPOSE 80
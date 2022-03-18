FROM node:lts-alpine3.15 AS builder

COPY ./ui /ui

WORKDIR /ui

RUN npm install && npm run build

FROM caddy:alpine

COPY --from=builder /ui/dist /srv

COPY ./caddy/Caddyfile /etc/caddy/Caddyfile
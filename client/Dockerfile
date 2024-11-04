FROM node:20.16.0-alpine as build
WORKDIR /app
COPY . .
RUN npm install && npm run build

FROM node:18.17.0-alpine as prod
ENV Certificate_Path=/run/secrets/certificate.pfx
ENV Certificate_Password=/run/secrets/certificate_password
ENV NODE_OPTIONS=--openssl-legacy-provider
RUN npm install -g serve
COPY --from=build /app/dist /app/dist
WORKDIR /app/dist
ENTRYPOINT serve --cors -l 443 --ssl-cert $Certificate_Path --ssl-pass $Certificate_Password -s -n
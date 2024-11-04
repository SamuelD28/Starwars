FROM node:20.16.0 as build
WORKDIR /app
COPY . .
RUN npm install && npm run build

FROM node:20.16.0-alpine as prod
RUN npm install -g serve
COPY --from=build /app/dist /app/dist
WORKDIR /app/dist
ENTRYPOINT serve --cors -l 5173
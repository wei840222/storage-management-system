FROM node AS build
WORKDIR /app
ADD . /app
RUN yarn && yarn build

FROM nginx
COPY --from=build /app/dist /usr/share/nginx/html
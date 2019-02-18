FROM node:10.15.1 as build-deps
WORKDIR /usr/src/app
COPY package.json ./
RUN yarn
COPY . ./

ENV REACT_BACKEND_URL "http://opendata.com/v1/graphql"
RUN yarn build

FROM nginx:1.12-alpine
COPY --from=build-deps /usr/src/app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
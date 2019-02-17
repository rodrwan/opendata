FROM alpine

RUN apk add --update ca-certificates

COPY bin/graphql /usr/bin/graphql

EXPOSE 3001

ENTRYPOINT graphql
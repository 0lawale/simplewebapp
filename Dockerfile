# build stage
FROM golang:1.20-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /bin/webapp

# final image
FROM alpine:3.18
RUN addgroup -S app && adduser -S app -G app
COPY --from=build /bin/webapp /usr/local/bin/webapp
USER app
EXPOSE 8080
HEALTHCHECK --interval=10s --timeout=3s --start-period=5s CMD wget -qO- http://localhost:8080/health || exit 1
ENTRYPOINT ["/usr/local/bin/webapp"]

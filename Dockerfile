# build the server
FROM golang as build
ADD tp.go /app/tp.go
WORKDIR /app/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/tp tp.go

# add it into a scratch image
FROM scratch
WORKDIR /
COPY --from=build /app/tp /tp

# and set the entry command
EXPOSE 80
CMD ["/tr", "0.0.0.0:80"]
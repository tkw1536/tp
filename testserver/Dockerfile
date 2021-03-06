# build the user permission server
FROM alpine as permission

# Create www-data
ENV USER=www-data
ENV UID=82
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

# build the server
FROM golang as build

# build the app
ADD echo.go /app/echo.go
WORKDIR /app/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/echo echo.go

# add it into a scratch image
FROM scratch
WORKDIR /

# add the user
COPY --from=permission /etc/passwd /etc/passwd
COPY --from=permission /etc/group /etc/group

# add the app
COPY --from=build /app/echo /echo

# and set the entry command
EXPOSE 8080
USER www-data:www-data
CMD ["/echo", "0.0.0.0:8080"]

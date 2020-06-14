FROM golang:alpine AS builder

WORKDIR /src/sqrt

COPY ./src/sqrt .

RUN CGO_ENABLED=0 GOOS=linux go build -o sqrt -a -installsuffix cgo -ldflags '-w -extldflags "-static"'

WORKDIR /bin
RUN cp /src/sqrt/sqrt ./sqrt

FROM scratch

COPY --from=builder /bin .

ENTRYPOINT ["./sqrt"]
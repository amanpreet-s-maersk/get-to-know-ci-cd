FROM golang:alpine3.18 as builder
WORKDIR /
COPY . .
RUN go build -o main

FROM scratch
WORKDIR /app
COPY --from=builder . .
CMD [ "./main" ]
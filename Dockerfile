FROM golang as build

WORKDIR /app

ENV CGO_ENABLED=0

COPY . .

RUN go build -o api

FROM alpine 

COPY --from=build /app/api api

CMD [ "./api" ]
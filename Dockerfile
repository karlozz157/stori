FROM golang:1.18-alpine

WORKDIR /app

ENV CSV_PATH="/app/csv/example.csv"
ENV EMAIL_TO="karlozz157@gmail.com"
ENV EMAIL_SUBJECT="STORI | Estado de Cuenta"
ENV EMAIL_TEMPLATE="/app/templates/summary.html"
ENV EMAIL_HOST="sandbox.smtp.mailtrap.io"
ENV EMAIL_PORT="2525"
ENV EMAIL_USERNAME="98bed2986d7e75"
ENV EMAIL_PASSWORD="85b8ea42eee043"
ENV SERVER_PORT=":6969"

COPY . .

RUN apk add gcc musl-dev

RUN go get -d -v ./...

RUN go build -o stori .

EXPOSE 6969

CMD ["./stori"]

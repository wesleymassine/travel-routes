FROM golang:1.15
ARG file
ENV APP_NAME travel-routes
ENV PORT 5000

COPY . /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

RUN go get ./
RUN go build -o ${APP_NAME}

CMD ./${APP_NAME} ${file}

EXPOSE ${PORT}
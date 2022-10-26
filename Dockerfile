FROM go AS builder

WORKDIR /app
COPY . .

ENV CGO_ENABLED=0
RUN go build

###############################

FROM alpine

RUN apk add ruby

COPY --from=builder /app/overseer/overseer /bin
RUN mkdir -p /var/www/monolith/
COPY static  /var/www/monolith/
COPY cgi-bin /var/www/monolith/

VOLUME /etc/config.json
EXPOSE 9900-9999

CMD ["/bin/overseer", "/etc/config.json"]
FROM go AS builder

WORKDIR /app
COPY . .

ENV CGO_ENABLED=0
RUN go build

###############################

FROM alpine

COPY --from=builder /app/overseer/overseer /bin

VOLUME /etc/config.json
EXPOSE 9900-9999

CMD ["/bin/overseer", "/etc/config.json"]
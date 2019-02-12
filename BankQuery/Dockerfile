FROM alpine
RUN mkdir -p app/{static,run}
COPY config.json /app/run/config.json
ADD bankQuery /app/run/bankQuery

EXPOSE 7081

ENTRYPOINT ["/app/run/bankQuery"]
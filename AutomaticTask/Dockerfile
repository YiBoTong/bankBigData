FROM alpine
RUN mkdir -p app/{ftp,log,run}
COPY config.json /app/run/config.json
ADD automatic_task /app/run/automatic_task

ENTRYPOINT ["/app/run/automatic_task"]
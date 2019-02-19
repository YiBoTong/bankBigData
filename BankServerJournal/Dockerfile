FROM alpine
RUN mkdir -p app/{ftp,log,run,excelTemp,excelFile}
COPY _excelTemp /app/excelTemp
COPY config.json /app/run/config.json
ADD bank_server_journal /app/run/bank_server_journal

ENTRYPOINT ["/app/run/bank_server_journal"]
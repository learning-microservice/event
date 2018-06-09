FROM scratch
COPY bin/event /event
ENTRYPOINT ["/event"]
EXPOSE 19000

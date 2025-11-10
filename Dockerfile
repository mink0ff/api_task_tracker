FROM ubuntu:latest
LABEL authors="miko"

ENTRYPOINT ["top", "-b"]
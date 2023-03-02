FROM golang:1.19
RUN mkdir -p src/app
WORKDIR /src/app/
COPY . /src/app
EXPOSE 8088
CMD make all
FROM golang:1.18
COPY main.go src
RUN cd src && go mod init michelececcacci/db-proj && go mod tidy && go build . 
ENTRYPOINT [ "./src/db-proj" ]
#Baseimage
FROM golang:1.16.4-alpine3.13

#Copying codebase to container
COPY /codebase/ /codebase/

#Building the binary
RUN cd /codebase && go build -v -o /codebase/bin/server /codebase/src/main.go

#Setting env
ENV PORT=80

#Run command for binary
CMD ["sh", "-c", "/codebase/bin/server"]
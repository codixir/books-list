FROM golang:1.11

WORKDIR /go/src/go-books-list  
COPY . . 
RUN go get -d -v ./...
RUN go install -v ./...

# ADD . /go-books-list  

RUN  go build -o go-books-list .

EXPOSE 8000:8000

CMD ["./go-books-list"]
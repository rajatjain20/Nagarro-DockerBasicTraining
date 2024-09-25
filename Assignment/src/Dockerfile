FROM golang:1.23.1

# Set Work Directory
WORKDIR /go/app

# Copy my golang code
COPY ./webdemo.go /go/app/
COPY ./dbFunc.go /go/app/

# copy mod file of the project
COPY ./go.mod /go/app/
COPY ./go.sum /go/app/

# build the code
RUN go build -v -o ./ .

# Set Environment Variables
ENV TZ="Asia/Kolkata"

# expose the port used in listener
EXPOSE 3333
CMD [ "./webdemo" ]
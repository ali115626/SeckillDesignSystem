FROM golang:latest
#这个是最底层的image

#MAINTAINER yuy "test@163.com"

RUN mkdir /webapp


WORKDIR /webapp

COPY . /webapp
RUN go run main.go

EXPOSE 8081

#RUN #chmod +x router  哟这个是改变文件的？看来有命令出错了 他就没办法建立images了
#ENTRYPOINT ["./router"]
FROM alpine:latest
RUN mkdir /app
WORKDIR /app
ADD gin-blog /app/gin-blog
ADD ./conf /app/conf
CMD ["./gin-blog"]
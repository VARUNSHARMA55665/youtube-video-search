# temp container 
FROM golang:alpine as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o video_search .
# Final build with minimal FS
FROM golang:alpine as finalBuild
WORKDIR /app
COPY --from=builder /app/video_search /app/video_search
CMD ["/app/video_search"]

FROM alpine
ADD menger-service /menger-service
ENTRYPOINT [ "/menger-service" ]

FROM debian

COPY ./dist/kubernetes_webhook /app/kubernetes_webhook
COPY ./certs /app/certs
WORKDIR /app
ENV IT_DEVOPS_ENV="prodution"
CMD ["/app/kubernetes_webhook"]

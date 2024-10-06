FROM rancher/rancher:latest

EXPOSE 80 443

CMD ["--no-cacerts"]
FROM mhart/alpine-node:7.1.0
RUN mkdir -p /opt/app
WORKDIR /opt/app/
COPY main . 

EXPOSE 3000 
CMD ["sh","run.sh"]
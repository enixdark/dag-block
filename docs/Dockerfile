#--------------------------------------
# tage: Compile Apps
#--------------------------------------
FROM node:8.9.0

WORKDIR /app

COPY ./package.json /app

RUN npm install
RUN npm install swagger-cli@2.1.1 --save
COPY . /app/

RUN make build


#--------------------------------------
# Stage: Packaging Apps
#--------------------------------------
FROM nginx:1.13-alpine

VOLUME /app
COPY --from=0 /app/docs/html /app

EXPOSE 80

COPY docker-entrypoint.sh /
RUN chmod +x /docker-entrypoint.sh
CMD ["/docker-entrypoint.sh"]
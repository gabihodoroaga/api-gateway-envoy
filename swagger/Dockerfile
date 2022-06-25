FROM nginx:alpine

ENV SWAGGER_HOST_NAME=""

COPY nginx.conf /etc/nginx/nginx.conf
COPY public /usr/share/nginx/html
COPY swagger.json /usr/share/nginx/html/swagger/swagger.json
COPY swagger-init.sh /docker-entrypoint.d/swagger-init.sh

EXPOSE 8080

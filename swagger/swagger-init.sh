#!/bin/sh

sed -i "s/SWAGGER_HOST_NAME/$SWAGGER_HOST_NAME/g" /usr/share/nginx/html/swagger/swagger.json

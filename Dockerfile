FROM nginx:stable
MAINTAINER Allen Zou <zyl_work@163.com>

RUN echo "You must build & install the pages first"

RUN rm -rf /etc/nginx/conf.d
COPY ./etc/nginx /etc/nginx
COPY ./www /data/www


#EXPOSE 80 443
#
#CMD ["nginx", "-g", "daemon off;"]
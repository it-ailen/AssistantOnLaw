FROM nginx:stable
MAINTAINER Allen Zou <zyl_work@163.com>

RUN rm -rf /etc/nginx/conf.d
COPY ./etc/nginx /etc/nginx
COPY ./fake-cdn /data/cdn
COPY ./www /data/www


#EXPOSE 80 443
#
#CMD ["nginx", "-g", "daemon off;"]
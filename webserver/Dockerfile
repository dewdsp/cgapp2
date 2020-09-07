FROM nginx:alpine

LABEL maintainer="Vic Shóstak <truewebartisans@gmail.com>"

# Copy Nginx config.
COPY ./configs/nginx.conf /etc/nginx/nginx.conf

# Copy default website config.
COPY ./configs/default.conf /etc/nginx/conf.d/default.conf

# Export necessary port.
EXPOSE 80
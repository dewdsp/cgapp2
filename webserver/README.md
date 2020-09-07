# Configured Docker container with Nginx web/proxy server for [Create Go App CLI](https://github.com/create-go-app/cli)

[Nginx](https://nginx.org/) [engine x] is an HTTP and reverse proxy server, a mail proxy server, and a generic TCP/UDP proxy server, originally written by Igor Sysoev. For a long time, it has been running on many heavily loaded Russian sites including Yandex, Mail.Ru, VK, and Rambler.

## ⚙️ Configuration

```nginx
# ./configs/nginx.conf

# Start NGINX configuration
worker_processes  auto;

events {
    use                 epoll;
    multi_accept        on;
    worker_connections  1024;
}

http {
    charset       utf-8;
    sendfile      on;
    tcp_nopush    on;
    tcp_nodelay   on;
    server_tokens off;
    log_not_found off;

    # Bucket size for server names hash
    server_names_hash_bucket_size 64;

    # MIME
    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;

    # Logging
    error_log  /var/log/nginx/error.log crit;
    access_log off;

    # Timeouts
    send_timeout              2;
    keepalive_timeout         30;
    keepalive_requests        100;
    client_body_timeout       10;
    reset_timedout_connection on;

    # Max body size
    client_max_body_size 4m;

    # Cache
    open_file_cache          max=200000 inactive=20s;
    open_file_cache_valid    30s;
    open_file_cache_errors   on;
    open_file_cache_min_uses 2;

    # Gzip
    gzip            on;
    gzip_vary       on;
    gzip_disable    "msie6";
    gzip_proxied    any;
    gzip_comp_level 6;

    # File types for compress via gzip
    gzip_types  text/plain text/css application/json application/x-javascript text/xml
                application/xml application/xml+rss text/javascript application/javascript
                image/svg+xml image/gif image/png image/jpeg image/x-icon image/webp;

    # SSL
    ssl_ciphers               ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384;
    ssl_protocols             TLSv1.2 TLSv1.3;
    ssl_session_cache         shared:SSL:10m;
    ssl_session_timeout       1d;
    ssl_session_tickets       off;
    ssl_prefer_server_ciphers off;

    # SSL root cert
    ssl_dhparam /etc/ssl/certs/dhparam.pem;

    # Security headers
    add_header X-Frame-Options           "SAMEORIGIN" always;
    add_header Referrer-Policy           "no-referrer-when-downgrade" always;
    add_header X-XSS-Protection          "1; mode=block" always;
    add_header X-Content-Type-Options    "nosniff" always;
    add_header Content-Security-Policy   "default-src 'self' http: https: data: blob: 'unsafe-inline'" always;
    add_header Strict-Transport-Security "max-age=31536000; includeSubDomains; preload" always;

    # Load configs
    include /usr/share/nginx/modules/*.conf;
    include /etc/nginx/conf.d/*.conf;
    include /etc/nginx/sites-enabled/*;
}
```

## ⚠️ License

MIT &copy; [Vic Shóstak](https://github.com/koddr) & [True web artisans](https://1wa.co/).
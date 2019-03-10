FROM redocly/redoc
COPY openapi.yaml /usr/share/nginx/html/openapi.yaml
ENV SPEC_URL openapi.yaml
EXPOSE 80
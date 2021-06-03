FROM scratch

LABEL maintainer="estafette.io" \
    description="The ${ESTAFETTE_GIT_NAME} component is an Estafette extension to update build status in Github for builds handled by Estafette CI"

COPY ca-certificates.crt /etc/ssl/certs/
COPY ${ESTAFETTE_GIT_NAME} /

ENV ESTAFETTE_LOG_FORMAT="console"

ENTRYPOINT ["/${ESTAFETTE_GIT_NAME}"]
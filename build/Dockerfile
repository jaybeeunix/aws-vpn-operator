FROM registry.access.redhat.com/ubi7/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/aws-vpn-operator \
    USER_UID=1001 \
    USER_NAME=aws-vpn-operator

# install operator binary
COPY build/_output/bin/aws-vpn-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}

FROM golang:1.19-buster as builder

# Create and change to the app directory.
WORKDIR /app

# ===================== Prepare bitbucket credentials ===========================
# For prod
# ARG BITBUCKET_LOGIN
# ARG BITBUCKET_PASS
# RUN echo "machine bitbucket.org login ${BITBUCKET_LOGIN} password ${BITBUCKET_PASS}" > ~/.netrc

# For local
#RUN apt-get update && apt-get install -y ca-certificates git-core ssh
#ADD ssh/id_rsa /root/.ssh/id_rsa
#ADD ssh/id_rsa.pub /root/.ssh/id_rsa.pub
#RUN chmod 600 /root/.ssh/id_rsa
#RUN chmod 600 /root/.ssh/id_rsa.pub
#RUN ssh-keyscan -t rsa bitbucket.org 2>&1 >> /root/.ssh/known_hosts
#
#RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config
#RUN git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"
# =================================================================================

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./

RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN CGO_ENABLE=0 go build -v -o service ./cmd/main.go
RUN ls

FROM debian:buster-slim

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/service /app/service


RUN chmod +x /app/service

EXPOSE 3000

CMD ["/app/service", "app", "start"]
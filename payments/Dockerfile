FROM rust:1.36.0 as cargo-build

RUN apt-get update && \
    apt-get install musl-tools -y && \
    rustup target add x86_64-unknown-linux-musl

WORKDIR /usr/src/payments

COPY Cargo.toml Cargo.toml

RUN mkdir src/ && \
    echo "fn main() {println!(\"Updating the build!\")}" > src/main.rs && \
    RUSTFLAGS=-Clinker=musl-gcc cargo build --release --target=x86_64-unknown-linux-musl && \
    rm -f target/x86_64-unknown-linux-musl/release/deps/payments*

COPY . .

RUN RUSTFLAGS=-Clinker=musl-gcc cargo build --release --target=x86_64-unknown-linux-musl

FROM alpine:latest

RUN addgroup -g 1000 payments && adduser -D -s /bin/sh -u 1000 -G payments payments

WORKDIR /home/payments/bin/

COPY --from=cargo-build /usr/src/payments/target/x86_64-unknown-linux-musl/release/payments .

RUN chown payments:payments payments

USER payments

CMD ["./payments"]
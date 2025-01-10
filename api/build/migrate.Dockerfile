FROM migrate/migrate

WORKDIR /migrations

COPY ./cmd/migrations /migrations

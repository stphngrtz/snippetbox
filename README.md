# Snippetbox
This projects covers the topics from the book [Let's go](https://lets-go.alexedwards.net/) by [Alex Edwards](https://github.com/alexedwards).

Run the MySQL database instance with Docker.
```bash
docker run --name mysql -p 3306:3306 -e MYSQL_DATABASE=snippetbox -e MYSQL_USER=web -e MYSQL_PASSWORD=pass -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:latest
```

Connect to the MySQL database instance and login as user.
```bash
docker exec -it mysql bash
mysql -D snippetbox -u web -p
```

Run the application.
```bash
go run ./cmd/web
```

Build and run the application.
```bash
go build -o ./tmp/web ./cmd/web/
cp -r ./tls /tmp/
cd /tmp/
./web 
```

Run tests.
```bash
go test -v ./cmd/web
go test -v ./internal/models
go test -v -short ./...

```

Run tests with coverage.
```bash
go test -cover ./...

go test -covermode=count -coverprofile=./tmp/profile.out ./...
go tool cover -html=./tmp/profile.out
```

Set up the database (execute as user).
```sql
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'An old silent pond',
    'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'Over the wintry forest',
    'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'First autumn morning',
    'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);

SELECT id, title, expires FROM snippets;

CREATE TABLE sessions (
    token CHAR(43) PRIMARY KEY,
    data BLOB NOT NULL,
    expiry TIMESTAMP(6) NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);

CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
```

Connect to the MySQL database instance and login as root.
```bash
docker exec -it mysql bash
mysql -u root -p
```

Create and set up the test database (execute as root).
```sql
CREATE DATABASE test_snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE USER 'test_web';
GRANT CREATE, DROP, ALTER, INDEX, SELECT, INSERT, UPDATE, DELETE ON test_snippetbox.* TO 'test_web';
ALTER USER 'test_web' IDENTIFIED BY 'pass';
```

Create a TLS certificate.
```bash
mkdir tls
cd tls
go run "C:\Program Files\Go\src\crypto\tls\generate_cert.go" --rsa-bits=2048 --host=localhost
```

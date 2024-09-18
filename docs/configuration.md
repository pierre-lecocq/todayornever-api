# Configuration

The application reads its configuration from a `.env` file.

Create one with the desired values at the same directory level than the binary.

Here is an example:

```toml
SERVICE_HOST=localhost
SERVICE_PORT=8080

LOGGER_LEVEL=3

AUTH_ISSUER=todayornever-api
AUTH_SECRET=<REDACTED>
AUTH_EXPIRES=1

DATABASE_ENGINE=sqlite3
DATABASE_DSN=./todayorneverd.db
```

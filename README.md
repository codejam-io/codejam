# CodeJam Site

## docker compose
### Deploying

1. Pull all images

```bash
docker compose pull database redis
```

2. Build the docker image: `docker compose build server`

```bash
docker compose build server
```

3. Copy the example configuration file from `./backend/app/config.example.toml` to `./backend/app/config.toml`
4. Edit the configuration file
    1. The database user will be "postgres"
    2. The database host will be "database" (as the name of the service in the `docker-compose.yaml` file)
    3. The redis host will be "redis" (as the name of the service in the `docker-compose.yaml` file)
5. Start the PostgreSQL and Redis containers

```bash
docker compose up -d database redis
```

6. Start the docker container through docker compose

```bash
docker compose up -d server
```

7. Visit website at [http://localhost:8080](http://localhost:8080)

### Providing more overrides

1. Unignore the `docker-compose.override.yaml` in the `.gitignore` file
2. Make changes
3. Stage the updated `docker-compose.override.yaml` file
4. Ignore the `docker-compose.override.yaml` file in the `.gitignore` file

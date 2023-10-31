## Arv Toolkit

### Configuration
Configure via environment variable

| Environment variable   | Description                                           | Required | Default value | Example                                 |
| ---------------------- | ----------------------------------------------------- | -------- | ------------- | --------------------------------------- |
| SPLYR_SP_DC            | Spotify cookie sp_dc value                            | false    | -             | -                                       |
| SPLYR_TOKEN_CACHE_PATH | Spotify token response cache path                     | false    | "."           | /data/                                  |
| AUTH_TOKEN_SECRET      | Token secret use for signing token                    | true     | -             | -                                       |
| AUTH_VALID_CREDENTIALS | List of pair username and password for authentication | true     | -             | username1;password1,username2;password2 |
| REST_SERVER_PORT       | Port for server to listen                             | false    | 8080          | 80                                      |
| REST_SERVER_IP         | Netwrok interface for server to listen                | false    | 127.0.0.1     | 10.10.10.1                              |
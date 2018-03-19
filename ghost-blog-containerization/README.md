## Dockerization

The Ghost blog installation is run with root user and the configuration as admin user. It's always a good container practice to have a standard user for running application.

```
docker build -t mngaonkar/ghost .
```

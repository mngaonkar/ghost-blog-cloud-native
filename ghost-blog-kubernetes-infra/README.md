### Prerequisite
Set DigitalOcean token environment variable
```
export DO_TOKEN=<your token here>
```

Make sure public and private keys are generated for local machine. The kyes are located at ~/.ssh/id_rsa as private and ~/.ssh/id_rsa.pub as public keys. If the keys are not found, generate the key pair.
```
ssh-keygen -t rsa
```


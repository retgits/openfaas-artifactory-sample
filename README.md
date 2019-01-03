# OpenFaaS & Artifactory Sample

A small [OpenFaaS](https://openfaas.com) app built in Go, that uses JFrog Artifactory as Docker registry

## Run the app

To run the app, assuming you have [OpenFaaS](https://openfaas.com) and the [faas-cli](https://docs.openfaas.com/cli/install/) installed, you'll need to update the `hello-openfaas.yml` on:

* Line 3: Make sure this matches your OpenFaaS gateway URL (port 31112 is the default gateway port for OpenFaaS)
* Line 8: Make sure this matches `<artifactory host>:<port>/<docker repository>/hello-openfaas:latest` (`docker` is the default name of an Artifactory repository for Docker images)

To run you'll need two commands

```bash
docker login <artifactory host>:<port>
faas-cli up -f hello-openfaas.yml
```

## Invoke the function

To invoke the deployed function run

```bash
echo "Hello World!" | faas-cli invoke hello-openfaas
```

## License

See the [LICENSE](./LICENSE) file in the repository
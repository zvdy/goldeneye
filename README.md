# GoldenEye

[![Go Report Card](https://goreportcard.com/badge/github.com/zvdy/goldeneye)](https://goreportcard.com/report/github.com/zvdy/goldeneye)

<p align="center">
  <img src="https://i.ibb.co/WV14jNj/logo.png" alt="GoldenEye Logo" width="170" height="170">
</p>

GoldenEye is a powerful HTTP DoS (Denial of Service) attack tool designed to stress test web servers. It allows you to simulate multiple concurrent connections to a target URL, using various HTTP methods and user agents. This tool is intended for educational purposes and to test the robustness of your own web servers.

## Features

- Simulate multiple concurrent connections
- Support for custom HTTP methods (GET, POST, RANDOM)
- Option to disable SSL certificate verification
- Customizable user agents
- Debug mode for verbose output

## Installation

To install GoldenEye, you need to have Go installed on your system. Clone the repository and build the project:

```sh
git clone https://github.com/zvdy/goldeneye.git
cd goldeneye
go build -o goldeneye ./cmd/goldeneye/main.go
```

## Usage

To run GoldenEye, use the following command:

```sh
./goldeneye <url> [OPTIONS]
```

### Options

| Flag           | Description                                         | Default                  |
|----------------|-----------------------------------------------------|--------------------------|
| `-u, --useragents` | File with user-agents to use                         | randomly generated       |
| `-w, --workers`    | Number of concurrent workers                         | 50                       |
| `-s, --sockets`    | Number of concurrent sockets                         | 30                       |
| `-m, --method`     | HTTP Method to use ('get', 'post', 'random')         | get                      |
| `-d, --debug`      | Enable Debug Mode (more verbose output)              | False                    |
| `-n, --nosslcheck` | Do not verify SSL Certificate                        | True                     |
| `-h, --help`       | Shows this help                                      |                          |

### Example

```sh
./goldeneye http://example.com -w 100 -s 50 -m post -d
```

This command will start a DoS attack on `http://example.com` with `100 workers`, each opening `50 sockets`, using the `POST` method, and enabling debug mode.

### Docker

To build and run the application using Docker, follow these steps:

1. **Build the Docker image**

   ```sh
   docker build -t goldeneye .
   ```

2. **Run the Docker container**

   ```sh
   docker run -it --rm goldeneye <url> [OPTIONS]
   ```

   Replace `<url>` and `[OPTIONS]` with the appropriate values for your application.

### Docker Hub Repository

You can also pull the pre-built Docker image from Docker Hub:

```sh
docker pull zvdy/goldeneye
```

Docker Hub Repository: [zvdy/goldeneye](https://hub.docker.com/repository/docker/zvdy/goldeneye/general)

### Example

To run the application with a specific URL and options:

```sh
docker run -it --rm zvdy/goldeneye http://example.com 
```

### Dockerfile

Ensure you have the [Dockerfile](Dockerfile) in your cloned repository.

### Testing

You can run the tests in order to check the functionality too:

```sh
go test -v ./...
```


## Use Cases

- **Stress Testing**: Test the robustness and performance of your web servers under heavy load.
- **Educational Purposes**: Learn about HTTP DoS attacks and how to mitigate them.
- **Security Testing**: Identify potential vulnerabilities in your web infrastructure.

## Disclaimer

GoldenEye is intended for educational purposes and testing your own web servers. Do not use this tool to attack websites without permission. Unauthorized use of this tool may violate local, state, and federal laws.

Read [DISCLAIMER](DISCLAIMER.md) for more details.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## Acknowledgements
 
This project is inspired by and based on the original [goldeneye](https://github.com/jseidl/GoldenEye) by [@jseidl](https://github.com/jseidl/). Special thanks to the original author for their work and inspiration.

## Contact

For any questions or support, please open an issue on the GitHub repository.



## Overview 
This project is a Docker-based environment that leverages LibreOffice for converting various types of documents to PDF. The supported document formats include but are not limited to Word, OpenDocument, and more.

An instance of LibreOffice runs within the Docker container and is controlled through a local socket using the UNO protocol. The container is designed to be scalable and efficient, allowing for high-performance conversion of documents to PDF format.

## Getting started
To get started, ensure that you have Docker installed on your machine. If not, you can download Docker from the [https://www.docker.com/get-started/](official Docker website).

This micro-service is intended for use exclusively as a docker container. While it may be possible to run this standalone is not an expected by the maintainer and you should rely on your own in this case.
After ensuring Docker is installed and running, follow these steps:
1. Clone this repository to your local machine.
2. Navigate to the project directory.
3. Build the Docker image using the provided Dockerfile.

```shell
docker build -t libreoffice_converter . 
docker run -d -p 6000:6000 libreoffice_converter
```

## Usage
Example:

```shell
curl -X POST -F "file=@path_to/file.doc" http://localhost:6000/docx2pdf > your_converted_file.pdf
```

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please submit a pull request with your proposed changes.

## License

MIT, see `LICENSE` for details.

## Acknowledgments

This project utilizes `LibreOffice`, a powerful open-source office suite, to perform document conversion. `Docker` a fantastic tool that is used to automate the deployment of applications in lightweight containers so that applications can work efficiently in different environments in isolation. And `Golang` a statically typed, mainstream language designed at Google.

## Contact

If you have any questions, issues, or suggestions, feel free to open an issue or submit a pull request.

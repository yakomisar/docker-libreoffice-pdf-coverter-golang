# LibreOffice PDF Converter Docker Environment

This repository provides a Docker-based environment that utilizes LibreOffice for converting various types of documents into PDF format. It supports a multitude of document formats, including, but not limited to, Word, Excel, PowerPoint and OpenDocument.

LibreOffice can convert a wide variety of document formats into PDF. These include:
1. Microsoft Office formats: Word (.doc, .docx), Excel (.xls, .xlsx), PowerPoint (.ppt, .pptx)
2. OpenDocument formats: Text (.odt), Spreadsheet (.ods), Presentation (.odp)
3. Other text formats: Rich Text Format (.rtf), Text (.txt), HTML (.html), XML (.xml)
4. Publisher files (.pub)
5. WordPerfect files (.wpd)
6. StarWriter files (.sdw)
7. Visio files (.vsd)
8. Quattro Pro files (.wb2)
9. Lotus 1-2-3 files (.123)

## System Design

A LibreOffice instance operates within the Docker container and available from your local machine (e.g. `localhost:6000`). This design allows the system to be both scalable and efficient, ensuring high-performance document conversions to PDF format.

## Prerequisites
This service is exclusively designed for operation as a Docker container. Although standalone operation may be possible, it is neither expected nor officially supported by the maintainers of this project.

Before you start, ensure Docker is installed and running on your machine. If not, download Docker from the [official Docker website](https://www.docker.com/get-started).

## Installation

Follow these steps to set up the environment:
1. Clone this repository to your local machine.
2. Navigate to the project directory.
3. Build the Docker image using the provided Dockerfile.

```shell
docker build -t libreoffice_converter . 
docker run -d -p 6000:6000 libreoffice_converter
```

## Usage
Use the following command to convert a document to PDF:

```shell
curl -X POST -F "file=@path_to/file.doc" http://localhost:6000/doc2pdf > your_converted_file.pdf
```

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please submit a pull request with your proposed changes.

## License

This project is licensed under the MIT License. For more details, please see the `LICENSE` file.

## Acknowledgments

This project leverages several significant open-source technologies:

- `LibreOffice` - A robust open-source office suite used for document conversion.
- `Docker` - An excellent tool for automating the deployment of applications within lightweight containers, enabling apps to work efficiently across various environments.
- `Golang` - A statically-typed, compiled language designed at Google.

## Contact

Should you have any questions, issues, or suggestions, please feel free to open an issue or submit a pull request.

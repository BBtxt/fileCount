# File Extension Counter (fileCount)

fileCount is a command-line tool written in Go that scans a specified directory and counts the occurrences of each file extension within that directory. It is useful for quickly understanding the distribution of different file types in a directory.

## Getting Started

These instructions will guide you through the process of building and using fileCount.
Prerequisites

To use fileCount, you need to have Go installed on your system. You can download and install Go from https://golang.org/dl/.

## Installation

    Clone this repository to your local machine. Alternatively, you can just copy the main.go file.
```sh

git clone https://github.com/yourusername/fileext.git
cd fileext
```

Build the Executable: Run the following command in the directory where main.go is located to build the executable:

```sh

    go build -o fileext main.go

    This command creates an executable named fileext.
```

## Usage

To use fileCount, simply run the executable from the command line, followed by the path of the directory you want to scan. For example:

```sh

./fileCount /path/to/directory
```
Replace /path/to/directory with the actual path of the directory you wish to scan.

The tool will output a list of file extensions found in the directory along with the count of files for each extension.

## Example

Running fileCount on a directory might produce output like this:

```sh

./fileCount /path/to/directory

Extension: .jpg, Count: 374
Extension: .DNG, Count: 9014
Extension: .tif, Count: 265
Extension: .JPG, Count: 476
```
This output shows that in the specified directory, there are 374 ".jpg" files, 9014 ".DNG" files, 265 ".tif" files, and 476 ".JPG" files. 

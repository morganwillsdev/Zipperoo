# Zipperoo

Zipperoo is an open source CLI tool for creating and decompressing zip files, written in Go. Currently the only supported filetype is .zip but more will be added soon.

## How to use

First, clone the repository and build the files using:
```
~$ go build
```
then run the executable file with the correct commands as listed below.

## Commands

### -zip

Compresses the files specified by [files_to_zip] to the zip file named in [output.zip] e.g.
```
~$ ./Zipperoo -zip [output.zip] [files_to_zip]
```

### -unzip

Decompresses the file [zip_file.zip] to the directory [output_dir] defaults to zipperoo-output if no [output_dir] specified e.g.
```
~$ ./Zipperoo -unzip [zip_file.zip] [output_dir]
```

### -help

Displays a help message with instructions on commands and how to use them e.g.
```
~$ ./Zipperoo -help
```
## GO-RELOADED
This project involves a Go application designed to process and format text from an input file and write the modified text to an output file. The application performs various text manipulations including punctuation formatting, quote adjustments, and specific text transformations based on predefined commands.

## Getting Started

### Prerequisites

- Go (Golang) installed on your machine.
- Input file containing text to be processed.

### Running the Program

To run this program, you'll need to compile it with Go and then execute it with two command-line arguments: the input file and the output file.

```bash
go run main.go input.txt output.txt
```

## Code Description

### Functions

#### `main`

The `main` function orchestrates the file opening, reading, processing, and writing operations. It handles the initialization of file scanners and writers, iterates over each line of the input file, applies formatting functions, and writes the result to the output file.

#### `formatPunctuation`

This function adjusts punctuation spacing in the text. It ensures that punctuation marks are followed by a space if appropriate, handles ellipses correctly by preventing unwanted spaces within sequences, and removes unnecessary spaces before punctuation marks.

#### `formatSingleQuotes`

This function modifies the text to correctly format content within single quotes. It removes any spaces immediately inside the single quotes, ensuring the quoted text is cleanly formatted without leading or trailing spaces.

### Standard Library Packages

#### `bufio`

- Used for buffered reading and writing, which improves I/O efficiency. In this program, `bufio.NewScanner` is used to read input from a file line by line, and `bufio.NewWriter` could be used for buffered output (though not explicitly shown here).

#### `fmt`

- Provides I/O formatting capabilities. The program uses `fmt.Println` to print error messages and notifications to the console, aiding in debugging and user interaction.

#### `os`

- Facilitates operating system functionality like file handling and command-line arguments. The program uses `os.Open` to open the input file for reading, `os.Create` to create the output file for writing, and `os.Args` to access command-line arguments.

#### `strconv`

- Contains functions to convert basic data types to and from string representations. In this program, `strconv.ParseInt` is used to convert strings in different numeric bases (binary and hexadecimal) into integer values.

#### `strings`

- Provides functions to manipulate UTF-8 encoded strings. The program uses multiple functions from this package:
  - `strings.Fields` to split lines into words.
  - `strings.ToLower` to convert strings to lowercase for case-insensitive comparison.
  - `strings.TrimSpace` to trim unwanted spaces around strings.
  - `strings.Join` to concatenate slices of words back into a single string.

## Conclusion

This Go application efficiently processes text files to apply specific formatting rules and transformations. It is designed to be easy to use and modify for various text processing needs, leveraging the power of Go's standard library to handle file I/O and string manipulation with high performance.



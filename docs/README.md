// README.md
# Data Parser

A .NET Core library for parsing data from various sources.

## Features

*   Supports CSV, JSON, and XML file formats
*   Ability to parse data from standard input
*   Optional filtering and sorting of parsed data

## Usage

### Installing

Install the Data Parser library using NuGet:
```bash
dotnet add package data-parser
```
### Parsing CSV Data

```csharp
using DataParser;

// Create a new CSV parser
var csvParser = new CsvParser("path/to/csv/file.csv");

// Parse the CSV data into a list of objects
var data = csvParser.Parse();
```
### Parsing JSON Data

```csharp
using DataParser;

// Create a new JSON parser
var jsonParser = new JsonParser("path/to/json/file.json");

// Parse the JSON data into a list of objects
var data = jsonParser.Parse();
```
### Parsing XML Data

```csharp
using DataParser;

// Create a new XML parser
var xmlParser = new XmlParser("path/to/xml/file.xml");

// Parse the XML data into a list of objects
var data = xmlParser.Parse();
```
### Parsing Standard Input

```csharp
using DataParser;

// Create a new parser from standard input
var parser = new Parser(Console.In);

// Parse the data into a list of objects
var data = parser.Parse();
```
### Filtering and Sorting Data

```csharp
using DataParser;

// Create a new CSV parser
var csvParser = new CsvParser("path/to/csv/file.csv");

// Filter and sort the parsed data
var filteredData = csvParser.Parse().Where(x => x.Id > 10).OrderBy(x => x.Name);
```
## Contributing

Contributions are welcome. Please see the [CONTRIBUTING.md](CONTRIBUTING.md) file for more information.

## License

This library is licensed under the [MIT License](LICENSE.md).
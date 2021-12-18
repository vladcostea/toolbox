## Toolbox CLI

A simple webdev utility program that allows developers to quickly validate and format JSON code, convert from UNIX epoch to timestamp and back, validate URLs, spin up a simple server and more.

`Toolbox` is meant to be used with the system clipboard to help developers quickly switch contexts without having to `cat` and pipe files or output streams.

### Example

```shell
# Validate JSON
# toolbox will use whatever you have in your most recent clipboard
toolbox json 

# Format JSON
toolbox json format
```

### TODO

- [] File as input (--file)
- [] XML validation and formatting
- [] YAML linting.
- [] JSON minification
- [] UNIX epoch / timestamp conversions
- [] CSS minification
- [] URL validation
- [] URL parser
- [] Simple dev server

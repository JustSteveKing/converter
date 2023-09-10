# Converter

A CLI tool to turn JSON and YAML into PHP Data Objects quickly

## Usage

### JSON

Creating a PHP Data Object from JSON

```bash
converter json -c Framework -j '{ "name": "Laravel", "version": 10.0, "active": true }'
```

```php
final readonly class Framework
{
    public function __construct(
        private string $name,
        private int $version,
        private bool $active,
    ) {}
}
```

### YAML

Currently in progress


### JSON Schema

Convert a JSON Payload into a JSON Schema object

```bash
converter schema --json='{"name": "John", "age": 25, "isStudent": true}'
```

```json
{
  "properties": {
    "age": {
      "type": "integer"
    },
    "isStudent": {
      "type": "boolean"
    },
    "name": {
      "type": "string"
    }
  },
  "type": "object"
}
```
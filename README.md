# Converter

A CLI tool to turn JSON and YAML into PHP Data Objects quickly

## Usage

Creating a PHP Data Object from JSON

### JSON

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
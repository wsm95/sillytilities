# S I L L Y T I L I T I E S

These are utilities, that are silly. sillytilities

## guidpls

GUID PLEEEEEASE! Generates a random guid and copies it to the clipboard.

```
usage: guidpls [-h] [-n] [count]

Generate and copy GUIDs to clipboard.

positional arguments:
  count Number of GUIDs to generate (default is 1)

options:
  -h, --help show this help message and exit
  -n, --nocopy Do not copy GUIDs to clipboard
```

## jottin

Hate when random website or applications steal your precious JWT token? Well fear no more. This will read the JWT token either from your clipboard, or pass it in as the first argument. Then it will decode and display the contents to the console.

```
usage: jottin [-h] [token]

Decode JWT tokens.

positional arguments:
  token       JWT token to decode

options:
  -h, --help  show this help message and exit
```

## How to run

Add the `exe` directory to your PATH variables. This will allow you to call all sillytilities from anywhere, just follow the usage guides.

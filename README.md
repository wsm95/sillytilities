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
  -h, --help
      Show this help message and exit
  -n, --nocopy
      Do not copy GUIDs to clipboard
```

## jottin

Hate when random website or applications steal your precious JWT token? Well fear no more. This will read the JWT token either from your clipboard, or pass it in as the first argument. Then it will decode and display the contents to the console.

```
usage: jottin [-h] [token]

Decode JWT tokens.

positional arguments:
  token       JWT token to decode

options:
  -h, --help
    Show this help message and exit
```

## bing

Sometime you need to bong. So lets bing.

```
usage: bing [-h] [-f] [-t]

Lets bong.

options:
  -h, --help
      Show this help message and exit
  -f, --frequency
      Frequency of the tone in Hz (default 116 [Bb])
  -t, --time
      Duration of the tone in seconds (default 3)
```

## iponcmd

Use this when you need your IP on command. This will print your current IP to the console and copy to your clipboard.

```
usage: iponcmd [-n] [-h]

Print your current IP address to the screen.

options:
  -n, --nocopy
      Do not copy GUIDs to clipboard
  -h, --help
      Show this help message and exit
```

## hawthawthawt

It's get-ting HAWT HAWT HAWT. Computer feeling a little hot under the collar? Use this to print the current temperature to the screen.

```
usage: hawthawthawt [-c] [-f] [-k]

Print the CPU temperature to the screen.

options:
  -c, --celsius
    Display temperature in Celsius
  -f, --fahrenheit
    Display temperature in Fahrenheit (default)
  -k, --kelvin
    Display temperature in Kelvin
```

## How to run

Add the `exe` directory to your PATH variables. This will allow you to call all sillytilities from anywhere, just follow the usage guides.

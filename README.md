# ende (encoder/decoder)

Simple cli tool for encoding and decoding to and from common formats.

## Examples
```
# encode, base64 as default
echo Hello | ende
# outputs SGVsbG8K

# decode, base64 as default
echo SGVsbG8K | ende -d
```


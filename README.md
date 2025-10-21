![alt text](AutoDeGoder.png)

# AutoDeGOder

This is a Go utility script that performs XOR + NOT decryption on hex-encoded strings using a given key. (It's a bit of a misnomer, but I like the name.)

## Features

- Decrypts multiple hex-encoded strings in one run
- Uses XOR + NOT operation for decryption
- Outputs results in a structured JSON format
- Handles multiple blobs in a single execution

## Usage

1. Open `main.go`
2. Modify the `blobs` array to include your hex-encoded strings:
```go
blobs := []string{
    "ENCODEDSTRING1",
    "ENCODEDSTRING2",
    "ENCODEDSTRING3",
}
```
3. Set your decryption key:
```go
key := "DECRYPTIONKEY"
```
4. Run the program:
```bash
go run main.go
```

## Output

The program creates a file named `decrypted_blobs.json` containing the decryption results in the following format saved in the same directory main.go runs. 
# lseg-exam

## How to work with this project

1. Copy .env.example to .env
```
cp .env.example .env
```
2. Update values in .env prioritize AWS_ACCESS_KEY_ID + AWS_SECRET_ACCESS_KEY over AWS_PROFILE
3. Install golang in your local computer. [Installation guide](https://go.dev/doc/install)
4. run go mod tidy to ensuring your project only uses necessary packages and dependencies in go.mod
5. Build executable file

for mac m1
```
GOOS=darwin GOARCH=arm64 go build -o lseg-exam
```

for mac intel (64-bit)
```
GOOS=darwin GOARCH=amd64 go build -o lseg-exam
```

for windows (64-bit)
```
GOOS=windows GOARCH=amd64 go build -o lseg-exam
```

6. grant permission (if needed)
```
chmod +x ./lseg-exam
```
7. Run file
```
./lseg-exam
```

## No license
# lseg-exam

## How to work with this project

1. Install golang in your local computer. [Installation guide](https://go.dev/doc/install)
2. run go mod tidy to ensuring your project only uses necessary packages and dependencies in go.mod
3. Build executable file

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

for linux
```
GOOS=linux GOARCH=amd64 go build -o lseg-exam
```

4. grant permission (if needed)
```
chmod +x ./lseg-exam
```

5. upload file to server use rsync (ensure your ~/.ssh/config is set)
```
rsync -a /path/to/your/file/lseg-exam remote_user@remote_host_or_ip:/path/on/target/dir/you/want/
```
6. Run file
```
./lseg-exam
```
7. got some result
```
Metadata: &{AccountID:xxxxxxxxxxxxxx Architecture:x86_64 AvailabilityZone:ap-southeast-1a BillingProducts:[] DevpayProductCodes:[] MarketplaceProductCodes:[] ImageID:ami-xxxxxxxxxxxxxxxxx InstanceID:i-xxxxxxxxxxxxxxxxx InstanceType:t3a.medium KernelID:null PendingTime:2025-03-26T00:50:09Z PrivateIP:xx.xx.xx.xx RamdiskID:null Region:ap-southeast-1 Version:2017-09-30}
```
8. or quick check in instnce
```
curl http://169.254.169.254/latest/dynamic/instance-identity/document
```

## No license
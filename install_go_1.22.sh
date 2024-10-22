FILE=go1.22.1.linux-amd64.tar.gz
GODIR=go-install-dir # Changed from GOBIN for clarity
apt-get update
apt-get install -y ca-certificates
if [ ! -f "$FILE" ]; then 
    apt-get install -y wget 
    wget https://go.dev/dl/$FILE
fi
if [ ! -d "$GODIR" ]; then 
    mkdir "$GODIR"
    tar -C "./$GODIR" -xzf $FILE
fi
export PATH=$PATH:${PWD}/${GODIR}/go/bin 
echo $PATH

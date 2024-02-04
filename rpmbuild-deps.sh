#!/usr/bin/env bash

echo "Installing BuildRequires dependencies";echo
grep ^BuildRequires "nxrmuploader.spec" |awk -F\: '{print "sudo dnf install -y"$2}'|sed -e 's/,/ /g' | sh
echo;echo;echo "Done. Now installing the Go binaries"

echo "Fetching archive..."
sudo wget -q https://go.dev/dl/go1.21.6.linux-amd64.tar.gz -O /opt/go.tar.gz

echo "Unarchiving..."
cd /opt ; sudo rm -rf go;sudo tar zxf /tmp/go.tar.gz; sudo rm -f /tmp/go.tar.gz

echo "Completed."


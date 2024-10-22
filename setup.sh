./install_go_1.22.sh
sudo chmod 777 keys
export PATH=${PWD}/gobin/go/bin:$PATH
./.github/scripts/init-temp-keys.sh
sudo chmod 777 keys/*
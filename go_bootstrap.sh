bash install_go_1.22.sh
/app/wait-for-it.sh keycloak:8888 --timeout=60 --strict -- \
sh -c "
  ./gobin/go/bin/go run ./service provision keycloak -e http://keycloak:8888/auth && \
  ./gobin/go/bin/go run ./service start
"

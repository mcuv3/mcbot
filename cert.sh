#! /usr/bin/env bash

echo | openssl s_client -servername lykke.com -connect hft-apiv2-grpc.lykke.com:443 |\
  sed -ne '/-BEGIN CERTIFICATE-/,/-END CERTIFICATE-/p' > cert.crt
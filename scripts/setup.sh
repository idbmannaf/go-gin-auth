#!/bin/bash
../
echo "\033[1;32m Generating JWT EC keys\033[0m"
openssl ecparam -name prime256v1 -genkey -noout -out jwtEC256.key
openssl ec -in jwtEC256.key -pubout -out jwtEC256.key.pub
JWT_PUBLIC_KEY_BASE64=$(base64 -i jwtEC256.key.pub)
JWT_PRIVATE_KEY_BASE64=$(base64 -i jwtEC256.key)
rm jwtEC256.key jwtEC256.key.pub

echo -e "\033[1;32m Updating .env file with the base64 encoded JWT keys\033[0m"
echo "JWT_PUBLIC_KEY_BASE64=\"$JWT_PUBLIC_KEY_BASE64\"" >> .env
echo "JWT_PRIVATE_KEY_BASE64=\"$JWT_PRIVATE_KEY_BASE64\"" >> .env

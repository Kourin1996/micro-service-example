mkdir -p ./cert

rm -f cert/*.pem

# CAの秘密鍵と証明書生成
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ./cert/ca-key.pem -out ./cert/ca-cert.pem -subj "/C=JP/ST=Tokyo/L=Tokyo/O=Kourin/OU=Development/CN=*.kourin.com/emailAddress=example@gmail.com"

echo "CA's self-signed certificate"
openssl x509 -in ./cert/ca-cert.pem -noout -text

# サーバーサイドの秘密鍵とCSRを生成
openssl req -newkey rsa:4096 -nodes -keyout ./cert/server-key.pem -out ./cert/server-req.pem -subj "/C=JP/ST=Tokyo/L=Tokyo/O=Kourin/OU=Development/CN=*.kourin.com/emailAddress=example@gmail.com"

# サーバーサイドのCSRをCAの秘密鍵を使って署名して証明書を生成
openssl x509 -req -in ./cert/server-req.pem -days 60 -CA ./cert/ca-cert.pem -CAkey ./cert/ca-key.pem -CAcreateserial -out ./cert/server-cert.pem -extfile ./ext.cnf

echo "Server's signed certificate"
openssl x509 -in ./cert/server-cert.pem -noout -text

# クライアントサイドの秘密鍵とCSRを生成
openssl req -newkey rsa:4096 -nodes -keyout ./cert/client-key.pem -out ./cert/client-req.pem -subj "/C=FR/ST=Alsace/L=Strasbourg/O=PC Client/OU=Computer/CN=*.pcclient.com/emailAddress=pcclient@gmail.com"

# クライアントサイドのCSRをCAの秘密鍵を使って署名して証明書を生成
openssl x509 -req -in ./cert/client-req.pem -days 60 -CA ./cert/ca-cert.pem -CAkey ./cert/ca-key.pem -CAcreateserial -out ./cert/client-cert.pem -extfile ./ext.cnf

echo "Client's signed certificate"
openssl x509 -in ./cert/client-cert.pem -noout -text

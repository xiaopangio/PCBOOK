del *.pem
# 1.generate CA' private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=CH/ST=GANSU/L=LANZHOU/O=LZU/OU=HIGHSCHOOL/CN=*.tech.xiaobai/emailAddress=techschool@gmail.com"
echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text
# 2.generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes  -keyout server-key.pem  -out server-req.pem -subj "/C=CH/ST=GANSU/L=LANZHOU/O=PC BOOK/OU=computer/CN=*.pcbook.com/emailAddress=pcbook@gmail.com"
# 3.use CA's private key sign web server's CSR and get back the signed certificate
openssl x509 -req -in server-req.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extensions v3_ca -extfile server-ext.cnf
echo "server's certificate"
openssl x509 -in server-cert.pem -noout -text
# 4.generate web client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes  -keyout client-key.pem  -out client-req.pem -subj "/C=CH/ST=GANSU/L=LANZHOU/O=PC Client/OU=computer/CN=*.pcclient.com/emailAddress=pcclient@gmail.com"
# 3.use CA's private key sign web client's CSR and get back the signed certificate
openssl x509 -req -in client-req.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extensions v3_ca -extfile client-ext.cnf
echo "client's certificate"
openssl x509 -in client-cert.pem -noout -text
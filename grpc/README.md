###单向验证:
genrsa -des3 -out server.key 2048   => 服务器私钥
req -new -key server.key -out server.csr    => 请求生成证书文件
common name -> localhost    => 客户端必须填相同的域名 
rsa -in server.key -out server_no_passwd.key    => 不用密码验证
x509 -req -days 365 -in server.csr -signkey server_no_passwd.key -out server.crt    =>生成证书


###双向验证
//生成CA, key和pem
genrsa -out ca.key 2048
req -new -x509 -days 365 -key ca.key -out ca.pem
common name -> localhost

//生成Server, key和pem
genrsa -out server.key 2048
req -new -key server.key -out server.csr
x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 365 -in server.csr -out server.pem

//生成client, key和pem
ecparam -genkey -name secp384r1 -out client.key
req -new -key client.key -out client.csr
x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 365 -in client.csr -out client.pem

###gateway
将grpc-ecosystem\grpc-gateway@v1.14.5\third_party\googleapis中的google目录拷贝到pbfiles目录, 便于引用
protoc --grpc-gateway_out=logtostderr=true:../services .\prod.proto
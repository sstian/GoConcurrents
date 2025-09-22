
## 认证和安全传输，使用openssl生成

### 生成证书
```cmd
# 1. 生成私钥 | server.key
openssl genrsa -out server.key 2048

# 2. 生成证书。全部回车即可，可以不填 | server.key -> server.crt
openssl req -new -x509 -key server.key -out server.crt -days 36500
// 国家名称 - Country Name (2 letter code) [AU]:
// 省会名称 - State or Province Name (full name) [Some-State]:
// 城市名称 - Locality Name (eg, city) []:
// 组织名称 - Organization Name (eg, company) [Internet Widgits Pty Ltd]:
// 部门名称 - Organizational Unit Name (eg, section) []:
// 服务器或网站名称 - Common Name (e.g. server FQDN or YOUR name) []:
// 邮件地址 - Email Address []:

# 3. 生成csr | server.key -> server.csr
openssl req -new -key server.key -out server.csr
// Country Name (2 letter code) [AU]:
// State or Province Name (full name) [Some-State]:
// Locality Name (eg, city) []:
// Organization Name (eg, company) [Internet Widgits Pty Ltd]:
// Organizational Unit Name (eg, section) []:
// Common Name (e.g. server FQDN or YOUR name) []:
// Email Address []:
// 
// Please enter the following 'extra' attributes
// to be sent with your certificate request
// A challenge password []:
// An optional company name []:

```

### 更改openssl文件
```
#1)复制一份你安装的openss1的bin目录里面的 openss1.cfg 文件到你项目所在的目录
#2)找到 [ CA_default ] ，打开 copy_extensions = copy (就是把前面的#去掉)
#3)找到 [ req ] ,打开 req_extensions = v3_reg # The extensions to add to a certificate request
#4)找到 [ v3_req ] ，添加 subjectAltName = @alt_names
#5)添加新的标签 [ alt_names ] ，和标签字段 DNS.1 =*.kuangstudy.com
```

### 生成私钥信息
```cmd
#生成证书私钥 | test.key
openssl genpkey -algorithm RSA -out test.key
// .....+......+...+...+.......+........+...................+.....+....+.....+.......+.....+...+.+...............+.....+...+.+..+++++++++++++++++++++++++++++++++++++++*......+.+..+.+.....+......+.+..+.............+........+....+..+.+..+...+.............+..+...+++++++++++++++++++++++++++++++++++++++*....+....+...............+.....+......+.......+..+.............+.........+...........+...+.........+...+.......+.....+...+..........++++++
.......+.........+.............+..+...+++++++++++++++++++++++++++++++++++++++*.......+++++++++++++++++++++++++++++++++++++++*.....+.......+...+..+.+...+.....+....+...+.....+.......+......+..+...+.......+.....+...+.+..............+.+......+..+..........++++++

#通过私钥test.key生成证书请求文件test.csr(注意cfg和cnf) | test.key + openssl.cfg -> test.csr
#test.csr是生成的证书请求文件。server.crt和server.key是CA证书文件和key,用来对test.csr进行签名认证，这两个文件在第一部分生成。
openssl req -new -nodes -key test.key -out test.csr -days 3650 -subj "/C=cn/OU=myorg/O=mycomp/CN=myname" -config ./openssl.cfg -extensions v3_req

#上述一般会有警告信息：Ignoring -days without -x509; not generating a certificate
#这个警告表明，在没有指定 -x509 选项的情况下，-days 选项将被忽略。-days 选项通常用于指定自签名证书（self-signed certificate）的有效期限，而不是 CSR。因此，如果你想生成一个自签名证书，你需要添加 -x509 选项，但是这样做下面的生成SAN证书会出现错误！！
#openssl req -new -nodes -key test.key -out test.csr -x509 -days 3650 -subj "/C=cn/OU=myorg/O=mycomp/CN=myname" -config ./openssl.cfg -extensions v3_req

#生成SAN证书pem | test.csr + openssl.cfg + server.key + server.crt -> test.pem
openssl x509 -req -days 365 -in test.csr -out test.pem -CA server.crt -CAkey server.key -CAcreateserial -extfile ./openssl.cfg -extensions v3_req
//Certificate request self-signature ok
//subject=C=cn, OU=myorg, O=mycomp, CN=myname
```
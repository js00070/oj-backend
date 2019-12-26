echo '\n测试登录接口'
curl -X POST --cookie-jar cookie.txt $HOST:8080/api/v1/user/login -d \
'{"user_name":"admin","password":"12345678"}'

echo '\n测试登录账号为空'
curl -X POST --cookie-jar cookie.txt $HOST:8080/api/v1/user/login -d \
'{"user_name":"  ","password":"12345678"}'

echo '\n测试登录密码为空'
curl -X POST --cookie-jar cookie.txt $HOST:8080/api/v1/user/login -d \
'{"user_name":"admin","password":"  "}'

echo '\n测试登录账号不存在'
curl -X POST --cookie-jar cookie.txt $HOST:8080/api/v1/user/login -d \
'{"user_name":"beijingdaxueuniverisity","password":"12345678"}'

echo '\n测试登录密码错误'
curl -X POST --cookie-jar cookie.txt $HOST:8080/api/v1/user/login -d \
'{"user_name":"admin","password":"univerisityofbeijingdaxue"}'

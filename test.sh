echo '接口测试：\n'

echo '\n测试注册接口'
curl -X POST $HOST:8080/api/v1/user/register -d \
'{"nickname":"admin","user_name":"admin","password":"12345678","password_confirm":"12345678"}'

echo '\n测试登陆接口'
curl -X POST --cookie-jar cookie.txt $HOST:8080/api/v1/user/login -d \
'{"user_name":"admin","password":"12345678"}'

echo '\n测试提交代码接口'
curl -X POST --cookie cookie.txt $HOST:8080/api/v1/commit -d \
'{"code":"codecodecode11111"}'

echo '\n测试提交代码接口'
curl -X POST --cookie cookie.txt $HOST:8080/api/v1/commit -d \
'{"code":"codecodecode22222"}'

echo '\n测试获取commit信息接口'
curl --cookie cookie.txt $HOST:8080/api/v1/commitlist

rm cookie.txt
echo '\n测试注册用户名空'
curl -X POST $HOST:8080/api/v1/user/register -d \
'{"nickname":"","user_name":"admin","password":"12345678","password_confirm":"12345678"}'

echo '\n测试注册密码password为空'
curl -X POST $HOST:8080/api/v1/user/register -d \
'{"nickname":"admin","user_name":"admin","password":"","password_confirm":"12345678"}'

echo '\n测试注册密码password_confirm为空'
curl -X POST $HOST:8080/api/v1/user/register -d \
'{"nickname":"admin","user_name":"admin","password":"12345678","password_confirm":""}'

echo '\n测试注册密码password和password_confirm不一致'
curl -X POST $HOST:8080/api/v1/user/register -d \
'{"nickname":"admin","user_name":"admin","password":"12345678","password_confirm":"1234567"}'

echo '\n测试注册用户名昵称不存在'
curl -X POST $HOST:8080/api/v1/user/register -d \
'{"nickname":"admin","user_name":"","password":"12345678","password_confirm":"12345678"}'

echo '\n测试注册用户名昵称已存在'
curl -X POST $HOST:8080/api/v1/user/register -d \
'{"nickname":"admin","user_name":"admin","password":"12345678","password_confirm":"12345678"}'

echo '\n测试注册用户名已存在'
curl -X POST $HOST:8080/api/v1/user/register -d \
'{"nickname":"admin","user_name":"admin1","password":"12345678","password_confirm":"12345678"}'
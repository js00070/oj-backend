echo '接口测试：\n'

echo '\n测试注册接口'
curl -X POST $HOST:8080/api/v1/user/register -d \
'{"nickname":"admin","user_name":"admin","password":"12345678","password_confirm":"12345678"}'

echo '\n测试登陆接口'
curl -X POST --cookie-jar cookie.txt $HOST:8080/api/v1/user/login -d \
'{"user_name":"admin","password":"12345678"}'

echo '\n获取题目列表'
curl --cookie cookie.txt $HOST:8080/api/v1/problemlist

echo '\n提交第一题正确的代码'
curl -X POST --cookie cookie.txt $HOST:8080/api/v1/commit -d \
'{"code":"a = int(input());print(a+a)","lang":"python","pid":1}'

echo '\n提交第二题正确的代码'
curl -X POST --cookie cookie.txt $HOST:8080/api/v1/commit -d \
'{"code":"a = int(input());print(a*a)","lang":"python","pid":2}'

echo '\n测试获取commit信息接口'
curl --cookie cookie.txt $HOST:8080/api/v1/commitlist

echo '\n等待三秒'
sleep 3

echo '\n测试获取commit信息接口'
curl --cookie cookie.txt $HOST:8080/api/v1/commitlist

echo '\n'

echo '\n提交第一题错误的代码'
curl -X POST --cookie cookie.txt $HOST:8080/api/v1/commit -d \
'{"code":"a = int(input());print(a+a+a)","lang":"python","pid":1}'

echo '\n提交第二题错误的代码'
curl -X POST --cookie cookie.txt $HOST:8080/api/v1/commit -d \
'{"code":"a = int(input());print(a*a*a)","lang":"python","pid":2}'

echo '\n等待三秒'
sleep 3

echo '\n测试获取commit信息接口'
curl --cookie cookie.txt $HOST:8080/api/v1/commitlist

echo '\n'

rm cookie.txt
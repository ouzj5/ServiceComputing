##Agenda CMD 设计文档

### 命令说明

#### 用户注册

`agenda register -u username -p password - e email -t telphone`

用户输入该命令会进行注册操作，用户必须给出下面四个参数的值，否则就会创建不成功。注册成功后用户的信息将会保存到电脑中。

参数：

- --username/-u：用户名
- --password/-p：密码
- --email/-e：邮箱
- --telphone/-t：电话

#### 用户登陆

`agenda login -u username -p password`

用户输入该命令会进行登陆操作，用户必须给出下面两个参数的值，然后程序会根据电脑存放的账号密码对进行匹配，匹配成功后将状态设置为已登陆对应的用户，不成功则返回错误信息。

参数：

- --username/-u：用户名
- --password/-p：密码

#### 用户登出

`agenda logout`

用户输入该命令会进行登出操作，用户不需要给出参数，然后程序会根据电脑存放的临时文件判断登陆状态，如果已经登陆则转到登出状态并返回登出信息，如果没有登陆则返回错误信息。
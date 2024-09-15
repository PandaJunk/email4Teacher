## 套磁小脚本

只需知道老师的姓名和邮箱还有学校，自备邮件模板和简历



### 1. 更改config.yaml文件

~~~yaml
myEmail: "自己的邮箱"
student: "姓名"
mySchool: "自己的学校"
major: "自己的专业"
txtPath: "邮件模板txt的路径" # 例如：resource/email.txt
auth: "邮箱的授权码"
smtpHost: "smtp地址" # 例如："smtp.163.com“
smtpPort: 465
subject: "邮件主题"
pdfPath: "简历路径" # 例如：resource/简历.pdf
teachers: # 老师列表，可以将一个学校的老师写在一个yaml文件中，然后在config.yaml中引用yaml文件路径
    - "resource/list/example.yaml"
~~~

修改自己的config配置



### 2. 添加要套的学校的yaml

例如我要套梦校，其中所有老师都写在example.yaml中，example.yaml中的内容如下：

~~~yaml
school: "梦校名字"
teachers:
  - name: "老师1"
    email: "老师1的邮箱"
  - name: "老师2"
    email: "老师2的邮箱"
~~~



### 3. 准备邮件模板和简历

简历就不多说了，转成pdf引用路径就可以，具体说一下邮件模板，给出的例子如下：

```txt
尊敬的Teacher老师：
    您好！我叫Student，来自MySchool的Major专业。。。
    School学术氛围浓厚，能够进入这里学习也一直是我的梦想。。。

学生：Student
Year年Month月Day日
```

其中Teacher、Student、MySchool、Major、Year、Month、Day为占位符

会根据config.yaml和老师列表中提供的信息进行替换

Year、Month、Day会自动获取

在需要老师姓名、自己姓名、自己学校名字、自己专业、年、月、日的地方使用如上占位符即可



### 4.发送邮件

运行main.go或者main.exe，看到全部邮件发送完毕则发送完成（***使用前记得先试一试拿自己的邮箱，避免产生问题***）


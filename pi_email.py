import imaplib
import email
import re
from email.header import decode_header

import pandas as pd
import paramiko

# 邮件服务器信息
IMAP_SERVER = 'imap.qq.com'  # 替换为你的IMAP服务器地址
USERNAME = '1047871481@qq.com'        # 替换为你的邮箱地址
PASSWORD = 'vnxdergxgcmkbccf'        # 替换为你的邮箱密码

# 连接到IMAP服务器
mail = imaplib.IMAP4_SSL(IMAP_SERVER)

# 登录
mail.login(USERNAME, PASSWORD)

# 选择INBOX文件夹
mail.select('INBOX')

# 搜索主题为 "PiNetWorkNode" 的邮件
search_criteria = 'SUBJECT "PiNetWorkNode"'
# 搜索所有邮件
status, messages = mail.search(None, search_criteria)  # 搜索所有邮件
if status != 'OK':
    print("无法搜索邮件")
    exit()

# 获取邮件ID列表
mail_ids = messages[0].split()

result_list = []
# 遍历每封邮件
for mail_id in mail_ids:
    # 获取邮件内容
    status, msg_data = mail.fetch(mail_id, '(RFC822)')  # 获取完整的邮件内容
    if status != 'OK':
        print(f"无法获取邮件 {mail_id}")
        continue

    result = {}
    # 解析邮件内容
    for response_part in msg_data:
        if isinstance(response_part, tuple):
            msg = email.message_from_bytes(response_part[1])  # 将邮件内容解析为Message对象

            # 解析邮件头
            subject, encoding = decode_header(msg['Subject'])[0]
            if isinstance(subject, bytes):
                subject = subject.decode(encoding if encoding else 'utf-8')

            content_type = msg.get_content_type()
            body = msg.get_payload(decode=True).decode()
            if content_type == "text/plain":
                for line in body.split('\n'):
                    if ':' in line:  # 检查是否包含冒号
                        key, value = re.split(r':\s*', line, maxsplit=1)
                        result[key.strip()] = value.strip()
                print(result)
                # 将字典追加到列表
                result_list.append(result)

df = pd.DataFrame(result_list)
df = df.drop_duplicates(subset=['服务器'])
df.to_excel('PiNetWorkNode_emails.xlsx', index=False)


# 登出
mail.logout()


df = pd.read_excel('PiNetWorkNode_emails.xlsx')

# 检查表头是否正确
required_columns = ['服务器', '端口', '账号', '密码']
if not all(column in df.columns for column in required_columns):
    print("Excel 文件缺少必要的列！")
    exit()

# 遍历每一行数据
for index, row in df.iterrows():
    hostname = row['服务器']
    port = row['端口']
    username = row['账号']
    password = row['密码']

    # 创建 SSH 客户端
    ssh = paramiko.SSHClient()
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())

    try:
        # 连接到服务器
        ssh.connect(hostname, port, username, password)
        print(f"{hostname} 连接成功！")

        # 执行命令（示例：获取服务器主机名）
        stdin, stdout, stderr = ssh.exec_command('hostname')
        hostname_output = stdout.read().decode('utf-8').strip()
        print(f"服务器主机名: {hostname_output}")

    except Exception:
        pass
    finally:
        # 关闭连接
        ssh.close()
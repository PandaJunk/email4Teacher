package main

import (
	"bufio"
	"crypto/tls"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func readText() string {
	// 打开txt文件
	file, err := os.Open(globalInfo.TxtPath)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer file.Close()

	// 使用bufio扫描器按行读取文件
	text := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 输出每一行的内容
		text += scanner.Text() + "\n"
	}

	// 检查扫描过程中是否出现错误
	if err := scanner.Err(); err != nil {
		log.Fatalf("读取文件时出错: %v", err)
	}

	return text
}

func sendEmail(from, rec, subject, text, attach string) error {
	// 初始化邮件内容
	e := email.NewEmail()
	e.From = from
	e.To = []string{rec}
	e.Subject = subject
	e.Text = []byte(text)
	_, err := e.AttachFile(attach)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}

	// 设置SMTP服务器信息
	smtpHost := globalInfo.SmtpHost
	smtpPort := globalInfo.SmtpPort

	// 建立TCP连接并启动TLS
	addr := smtpHost + ":" + smtpPort
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	// 连接到SMTP服务器
	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return err
	}

	// 创建SMTP客户端
	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return err
	}

	// 设置身份认证
	auth := smtp.PlainAuth("", globalInfo.MyEmail, globalInfo.Auth, smtpHost)
	if err = client.Auth(auth); err != nil {
		return err
	}

	// 发送邮件
	if err = e.SendWithTLS(addr, auth, tlsconfig); err != nil {
		return err
	}

	// 关闭客户端
	if err = client.Quit(); err != nil {
		return err
	}

	return nil
}

func textFmt(text, mySchool, major, teacher, school, year, month, day string) string {
	text = strings.Replace(text, "Teacher", teacher, -1)
	text = strings.Replace(text, "MySchool", mySchool, -1)
	text = strings.Replace(text, "Major", major, -1)
	text = strings.Replace(text, "School", school, -1)
	text = strings.Replace(text, "Student", globalInfo.Student, -1)
	text = strings.Replace(text, "Year", year, -1)
	text = strings.Replace(text, "Month", month, -1)
	text = strings.Replace(text, "Day", day, -1)
	return text
}

var globalInfo Info

func main() {

	config, list := getTeacherInfo("config.yaml")
	globalInfo = config

	mod := readText()

	today := time.Now()
	year := today.Year()
	month := today.Month()
	day := today.Day()

	var wg sync.WaitGroup
	for i := 0; i < len(list); i++ {
		school := list[i].School
		for j := 0; j < len(list[i].Teachers); j++ {
			wg.Add(1)
			go func(t Teacher, school string) {
				defer wg.Done()
				text := mod
				text = textFmt(text, globalInfo.MySchool, globalInfo.Major, t.Name, school, strconv.Itoa(year), strconv.Itoa(int(month)), strconv.Itoa(day))
				err := sendEmail(globalInfo.MyEmail, t.Email, globalInfo.Subject, text, globalInfo.PdfPath)
				if err != nil {
					log.Fatalf("发送给 %s 的邮件失败！\n", t.Name)
					return
				}
				log.Printf("邮件已发送给 %s 的 %s！\n", school, t.Name)
			}(list[i].Teachers[j], school)
		}

	}
	wg.Wait()
	log.Println("全部邮件发送成功！")
}

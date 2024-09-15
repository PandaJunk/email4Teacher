package main

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

// Teacher 结构体
type Teacher struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

// Config 结构体，包含服务器列表和数据库配置
type Info struct {
	MyEmail     string   `yaml:"myEmail"`
	MySchool    string   `yaml:"mySchool"`
	Major       string   `yaml:"major"`
	Auth        string   `yaml:"auth"`
	TxtPath     string   `yaml:"txtPath"`
	Student     string   `yaml:"student"`
	SmtpHost    string   `yaml:"smtpHost"`
	SmtpPort    string   `yaml:"smtpPort"`
	Subject     string   `yaml:"subject"`
	PdfPath     string   `yaml:"pdfPath"`
	TeacherList []string `yaml:"teachers"` // 结构体数组
}

type TeacherList struct {
	School   string    `yaml:"school"`
	Teachers []Teacher `yaml:"teachers"`
}

func getTeacherInfo(file string) (Info, []TeacherList) {
	// 读取YAML文件
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("无法读取文件: %v", err)
	}

	// 初始化配置结构体
	var info Info

	// 解析YAML文件内容
	err = yaml.Unmarshal(data, &info)
	if err != nil {
		log.Fatalf("解析YAML文件时出错: %v", err)
	}
	var teachers = make([]TeacherList, 0)
	for i := 0; i < len(info.TeacherList); i++ {
		listData, err := os.ReadFile(info.TeacherList[i])
		if err != nil {
			log.Fatalf("无法读取文件: %v", err)
		}
		// 初始化配置结构体
		var list TeacherList

		// 解析YAML文件内容
		err = yaml.Unmarshal(listData, &list)
		if err != nil {
			log.Fatalf("解析YAML文件时出错: %v", err)
		}

		teachers = append(teachers, list)
	}
	return info, teachers
}

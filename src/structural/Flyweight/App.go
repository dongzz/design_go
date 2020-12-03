package main

import "fmt"

type Website interface {
	Use(username string)
}

type ConcreteWebsite struct {
	name string
}

func (website *ConcreteWebsite) Use(username string) {
	fmt.Printf("用户：%s , 使用了【%s】网站 \n", username, website.name)
}

type WebsiteFactory struct {
	websiteMap map[string]Website
}

func (factory *WebsiteFactory) GetWebsite(name string) Website {
	if w, ok := factory.websiteMap[name]; ok {
		return w
	} else {
		c := &ConcreteWebsite{name: name}
		factory.websiteMap[name] = c
		return c
	}
}

func main() {
	factory := WebsiteFactory{websiteMap: make(map[string]Website)}

	b := factory.GetWebsite("博客")
	b.Use("dongz")

	c := factory.GetWebsite("bilibili")
	c.Use("caicai")

	d := factory.GetWebsite("bilibili")
	d.Use("pangpang")
	fmt.Println(c == d)
}

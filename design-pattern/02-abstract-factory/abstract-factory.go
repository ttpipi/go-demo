package abstract_factory

import "fmt"

type Item interface {
	Use()
}

//创建多种元素, 这些元素组成一个产品--抽象工厂
type ThemeFactory interface {
	//元素必须固定, 添加一种元素所有工厂子类都要修改
	//"开闭原则的倾斜性": 只对产品开放, 对元素不开放
	CreateDesktop() Item
	CreateIcon() Item
	CreateFont() Item
}

//RedTheme
type RedDesktop struct{}

func (r RedDesktop) Use() {
	fmt.Println("red desktop")
}

type RedIcon struct{}

func (r RedIcon) Use() {
	fmt.Println("red icon")
}

type Font struct{}

func (f Font) Use() {
	fmt.Println("red font")
}

//RedThemeFactory
type RedThemeFactory struct{}

func (r RedThemeFactory) CreateFont() Item {
	return &Font{}
}

func (r RedThemeFactory) CreateDesktop() Item {
	return &RedDesktop{}
}

func (r RedThemeFactory) CreateIcon() Item {
	return &RedIcon{}
}

//BlackThem
type BlackDesktop struct{}

func (b BlackDesktop) Use() {
	fmt.Println("black desktop")
}

type BlackIcon struct{}

func (b BlackIcon) Use() {
	fmt.Println("black icon")
}

type BlackFont struct{}

func (b BlackFont) Use() {
	fmt.Println("black font")
}

//BlackThemeFactory
type BlackThemFactory struct{}

func (b BlackThemFactory) CreateDesktop() Item {
	return &BlackDesktop{}
}

func (b BlackThemFactory) CreateIcon() Item {
	return &BlackIcon{}
}

func (b BlackThemFactory) CreateFont() Item {
	return &BlackFont{}
}

package abstract_factory

import "testing"

//同个工厂生产的元素必须在一起使用(也就是不同产品不能一起使用)
func useTheme(factory ThemeFactory) {
	factory.CreateDesktop().Use()
	factory.CreateIcon().Use()
	factory.CreateFont().Use()
}

func TestRedTheme(t *testing.T) {
	var factory ThemeFactory
	factory = &RedThemeFactory{} //使用者只需要知道哪些主题工厂, 不需要知道具体产品
	useTheme(factory)

	factory = &BlackThemFactory{}
	useTheme(factory)
}

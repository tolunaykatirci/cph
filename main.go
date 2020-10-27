package main

import (
	//"github.com/bilginyuksel/cordova-plugin-helper/writer"

	"github.com/bilginyuksel/cordova-plugin-helper/parser"
	"github.com/bilginyuksel/cordova-plugin-helper/reader"
)

func main() {
	plg, _ := parser.ParseXML("parser/plugin.xml")
	javaFiles, _ := reader.FilePathWalkDir("src")
	plg.Platform.NewSource(javaFiles)
	jsModules, _ := reader.FilePathWalkDir("www")
	plg.NewJsModules(jsModules)
	parser.CreateXML(plg, "plg.xml")
	//ctx := kong.Parse(&cli)
	// Call the Run() method of the selected parsed command.
	//err := ctx.Run(&Context{Debug: cli.Debug})
	//ctx.FatalIfErrorf(err)
}


// Context ...
type Context struct {
	Debug bool
}

// PluginXMLCmd ...
type PluginXMLCmd struct {
	Path string `help:"String to path"`
	Sync bool   `help:"Sync plugin.xml file. This command will search the related directories and if it finds any missing or unnecessary field it will add or delete automatically."`
}

// AddLicenseCmd ...
type AddLicenseCmd struct {
	Paths          []string `arg optional name:"path" help:"Paths to list." type:"path"`
	FileExtensions []string `arg optional name:"extension" help:"Instead of giving every files path, just give file extensions here."`
	License        string   `help:"License file path to use."`
}

// Run ...
// func (pl *PluginXMLCmd) Run(ctx *Context) error {
// 	fmt.Println(pl)
// 	return nil
// }

// Run ...
// func (l *AddLicenseCmd) Run(ctx *Context) error {
// 	fmt.Println("ls", l.Paths)
// 	return nil
// }

var cli struct {
	Debug bool `help:"Enable debug mode."`

	PluginXML  PluginXMLCmd  `cmd help:"You can use the functions in that command to manipulate plugin.xml file under cordova plugin root project directory."`
	AddLicense AddLicenseCmd `cmd help:"Add license to files."`
}

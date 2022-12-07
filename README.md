# clientDemo
一个用walk框架写的windowsGUI示例(golang) / A windows GUI demo written by walk framework(golang)

# 项目结构 / Project Structure
1:主目录(clientDemo)
master directory(clientDemo)

2:资源目录(clientDemo/resource)[包含图片和manifest声明打包文件]
resource directory(clientDemo/resource)[Include imgs and manifest files]

3:服务目录(clientDemo/service)[包含所有非ui相关的功能]
service directory(clientDemo/service)[Include all functions except ui]

4:界面目录(clientDemo/ui)[包含所有ui功能]
ui directory(clientDemo/ui)[Include all ui functions]

5:main.go[程序入口]
main.go[Program entry]

6:clientDemo.syso[由rsrc工具执行生成]
clientDemo.syso[Generate by rsrc tool]

7:go.mod[golang的包管理文件]
go.mod[Package management file of golang]

# 步骤 / Steps
1:安装golang环境,需要Go 1.11.x或更新的版本(https://go.dev/doc/install)
Install the golang environment,require Go 1.11.x or later(https://go.dev/doc/install)

2:执行"go mod tidy"命令下载go.mod文件中的依赖
Run "go mod tidy" command to download the dependencies in the go.mod file

3:如果想完整掌握项目编译过程，则需要安装rsrc工具
If you want to fully master the project compilation process, you need to install the rsrc tool

(1):执行"go get github.com/akavel/rsrc" 命令下载rsrc工具
Execute the "go get github.com/akavel/rsrc" command to download the rsrc tool

(2):进入GOPATH并找到rsrc目录
Enter GOPATH and find the rsrc directory

如果配置了GOPATH环境变量则可以直接通过"%GOPATH%\pkg\mod\github.com\akavel"进入
If the GOPATH environment variable is configured, you can directly enter through "%GOPATH%\pkg\mod\github.com\akavel"

如果没有配置且golang是默认安装的话，GOPATH通常会在"C:\Users\yourComputerName\go"目录下
If there is no configuration and golang is installed by default, GOPATH is usually in the "C: Users yourComputerName go" directory

(3):进入rsrc目录并将rsrc.exe文件拷贝到%GOROOT%\bin目录下,这时你就可以使用rsrc命令了(%GOROOT%是golang的安装目录)
Enter the rsrc directory and set the rsrc Copy the exe file to the% GOROOT% bin directory, and then you can use the rsrc command (% GOROOT% is the installation directory of golang)

(4):使用"rsrc -manifest resource/clientDemo.exe.manifest -o clientDemo.syso -ico resource/icon.ico"命令进行打包,"-manifest resource/clientDemo.exe.manifest"表示将"resource/clientDemo.exe.manifest"文件作为打包时的声明文件,"-o clientDemo.syso"设置生成的文件名称,"-ico resource/icon.ico"声明最终生成.exe文件的图标。PS:icon.ico文件必须是标准的.ico文件，不能是.jpeg文件直接修改文件后缀名形成。
Use the "rsrc -manifest resource/clientDemo.exe.manifest -o clientDemo.syso -ico resource/icon.ico" command to package. "-manifest resource/clientDemo.exe.manifest" means to use the "resource/clientDemo.exe.manifest" file as the declaration file when packaging. "-o clientDemo.syso" sets the name of the generated file. "-ico resource/icon.ico" declares the icon of the final generated.exe file. PS:The icon.ico file must be a standard. ico file, not a .jpeg file.can't get by Modifing the file suffix!

(5):执行" go build -ldflags="-H windowsgui" "命令进行编译，并形成.exe文件，" -ldflags="-H windowsgui" "参数用于消除cmd命令框,如果需要则可以直接使用"go build"命令进行编译
Execute the " go build -ldflags="-H windowsgui" " command to compile and form Exe file, " -ldflags="-H windowsgui" " parameter is used to eliminate the cmd command box. If necessary, you can directly use the "go build" command to compile

(6):运行clientDemo.exe文件
run clientDemo.exe file


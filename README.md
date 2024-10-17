# resume

个人简历API接口

go version go1.22.2 windows/amd64

在Windows平台上为Go程序添加图标可以按照以下步骤进行：

1. **准备图标文件**：确保你有一个图标文件，例如`icon.ico`。

2. **创建资源文件**：创建一个名为`resource.rc`的资源文件，内容如下：
    ```rc
    id ICON "icon.ico"
    ```

3. **安装MinGW工具链**：如果还没有安装MinGW，可以从[MinGW官网](http://www.mingw.org/)下载并安装。

4. **编译资源文件**：使用`windres`工具将资源文件编译为对象文件：
    ```sh
    windres resource.rc -O coff -o resource.syso
    ```

5. **编译Go程序并链接资源文件**：
   
   ```sh
   // 带控制台 Go 提供了 -trimpath 选项，可以去除编译时的路径信息
   go build -ldflags "-s -w" -trimpath -o resume-v1.0.exe
   
   // 不带控制台
   go build -ldflags="-H windowsgui" -o resume-v1.0.exe
   
   // -s：去除符号表
   // -w：去除调试信息
   // -trimpath：消除路径信息，避免暴露源码路径
   ```

web页面 [https://github.com/clovme/resume-web.git](https://github.com/clovme/resume-web.git)
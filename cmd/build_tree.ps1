# 改变当前PowerShell会话的位置到上一级目录，即项目根目录
Set-Location -Path C:\Users\nbdhc\GolandProjects\Tree_plantuml
Get-Location
Write-Host "Current directory: $(Get-Location)"
# 编译程序，输出到上级目录的output文件夹
go build -o ../output/TreeThisFolder.exe ./cmd/tree



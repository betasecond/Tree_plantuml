@echo off
Net session >nul 2>&1 || powershell start-process remove.bat -verb runas
SETLOCAL EnableExtensions

:: ������ǰ��ӵĽű�����
set SCRIPT_NAME=�����ļ��нṹ

:: ɾ��֮ǰ���ļ��б�������ӵ�ע�����
REG DELETE "HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\CommandStore\shell\Item0" /f
REG DELETE "HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\CommandStore\shell\Item1" /f
REG DELETE "HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\CommandStore\shell\Item2" /f
REG DELETE "HKEY_CLASSES_ROOT\Directory\Background\shell\Item0" /f
echo �Ҽ��˵�����ɾ����
rd /S /Q "C:\Program Files\Tree This Folder"
pause
ENDLOCAL

@echo off
setlocal

set APP=%cd%\main.exe

:: 使用说明
:menu
cls
echo ===========================
echo 请选择操作:
echo 1. 启动程序 (start)
echo 2. 停止程序 (stop)
echo 3. 查看状态 (status)
echo 4. 重启程序 (restart)
echo 5. 退出
echo ===========================
set /p choice="请输入选项 [1-5]: "

if "%choice%"=="1" (
    call :start
    pause
    goto menu
) else if "%choice%"=="2" (
    call :stop
    pause
    goto menu
) else if "%choice%"=="3" (
    call :status
    pause
    goto menu
) else if "%choice%"=="4" (
    call :restart
    pause
    goto menu
) else if "%choice%"=="5" (
    exit /b
) else (
    echo 无效的选项，请重试。
    pause
    goto menu
)

:: 程序状态
:status
tasklist /FI "IMAGENAME eq main.exe" | find /I "main.exe" >nul
if %errorlevel%==0 (
    for /f "tokens=2" %%i in ('tasklist /FI "IMAGENAME eq main.exe"') do (
        echo %APP% is running, pid is %%i
    )
) else (
    echo %APP% is not running
)
exit /b

:: 启动程序
:start
tasklist /FI "IMAGENAME eq main.exe" | find /I "main.exe" >nul
if %errorlevel%==0 (
    echo %APP% is already running
) else (
    start "" "%APP%" > console.log 2>&1
    echo %APP% start success
)
exit /b

:: 停止程序
:stop
tasklist /FI "IMAGENAME eq main.exe" | find /I "main.exe" >nul
if %errorlevel%==0 (
    for /f "tokens=2" %%i in ('tasklist /FI "IMAGENAME eq main.exe"') do (
        taskkill /PID %%i /F
        echo %APP% stop success
    )
) else (
    echo %APP% is not running
)
exit /b

:: 重启程序
:restart
call :stop
call :start
exit /b

endlocal
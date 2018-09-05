package winsvc

/*
#include <windows.h>
#include <stdio.h>

SERVICE_STATUS serviceStatus;
SERVICE_STATUS_HANDLE serviceStatusHandle;

void OurServiceControlHandler(DWORD request) {
	switch (request) {
		// Comment it for disable stop service //
		case SERVICE_CONTROL_STOP:
			serviceStatus.dwWin32ExitCode = 0;
			serviceStatus.dwCurrentState = SERVICE_STOPPED;
			SetServiceStatus(serviceStatusHandle, &serviceStatus);
		case SERVICE_CONTROL_SHUTDOWN:
			serviceStatus.dwWin32ExitCode = 0;
			serviceStatus.dwCurrentState = SERVICE_STOPPED;
			SetServiceStatus(serviceStatusHandle, &serviceStatus);
		//  END ///
		default:
			break;
	}
	SetServiceStatus(serviceStatusHandle, &serviceStatus);
	return;
}

BOOL OurServiceInit(LPCWSTR svc_name) {
	serviceStatus.dwServiceType       = SERVICE_WIN32_OWN_PROCESS;
	serviceStatus.dwCurrentState      = SERVICE_START_PENDING;
	serviceStatus.dwControlsAccepted  = SERVICE_ACCEPT_STOP | SERVICE_ACCEPT_SHUTDOWN;
	serviceStatus.dwWin32ExitCode     = 0;
	serviceStatus.dwServiceSpecificExitCode = 0;
	serviceStatus.dwCheckPoint        = 0;
	serviceStatus.dwWaitHint          = 0;
	serviceStatusHandle = RegisterServiceCtrlHandlerW(svc_name, (LPHANDLER_FUNCTION)OurServiceControlHandler);
	if (serviceStatusHandle == (SERVICE_STATUS_HANDLE)0) {
    	return 1;
  	}
	serviceStatus.dwCurrentState = SERVICE_RUNNING;
 	SetServiceStatus(serviceStatusHandle, &serviceStatus);
	return 0;
}

void OurServiceSetState(DWORD state) {
	serviceStatus.dwCurrentState = state;
	SetServiceStatus(serviceStatusHandle, &serviceStatus);
}

*/
import (
	"C"
	"syscall"
)

func OurServiceInit(service_name string) int {
	svcNamePtr, _ := syscall.UTF16PtrFromString(service_name)
	return int(C.OurServiceInit(svcNamePtr))
}

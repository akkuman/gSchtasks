package win32

import (
	"syscall"
	"unsafe"
)

const NULL = 0
const TRUE = 1
const CLSCTX_INPROC_SERVER = 0x1

var FAILED = func(hr HRESULT) bool { return hr < 0 }

type (
	LPVOID  uintptr
	LMSTR   *uint16
	DWORD   uint32
	LPBYTE  *byte
	LPDWORD *uint32
	BOOL    int
	// LPWSTR syscall.UTF16PtrFromString
	LPWSTR *uint16
	// LPCWSTR syscall.UTF16PtrFromString
	LPCWSTR *uint16
	// LPCOLESTR syscall.UTF16PtrFromString
	LPCOLESTR *uint16
	// BSTR syscall.UTF16PtrFromString
	BSTR    *uint16
	HRESULT int32
	// REFCLSID *GUID
	REFCLSID uintptr
	// REFIID *GUID
	REFIID uintptr
	// LPUNKNOWN *IUnknown
	LPUNKNOWN uintptr
)

type ITaskService struct {
	vtbl *iTaskServiceVtbl
}

type iTaskServiceVtbl struct {
	// IUnknown
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// IDispatch
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr

	// ITaskService
	GetFolder           uintptr
	GetRunningTasks     uintptr
	NewTask             uintptr
	Connect             uintptr
	get_Connected       uintptr
	get_TargetServer    uintptr
	get_ConnectedUser   uintptr
	get_ConnectedDomain uintptr
	get_HighestVersion  uintptr
}

type ITaskFolder struct {
	vtbl *iTaskFolderVtbl
}

type iTaskFolderVtbl struct {
	// IUnknown
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// IDispatch
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr

	// ITaskFolder
	get_Name               uintptr
	get_Path               uintptr
	GetFolder              uintptr
	GetFolders             uintptr
	CreateFolder           uintptr
	DeleteFolder           uintptr
	GetTask                uintptr
	GetTasks               uintptr
	DeleteTask             uintptr
	RegisterTask           uintptr
	RegisterTaskDefinition uintptr
	GetSecurityDescriptor  uintptr
	SetSecurityDescriptor  uintptr
}

type ITaskDefinition struct {
	vtbl *iTaskDefinitionVtbl
}

type iTaskDefinitionVtbl struct {
	// IUnknown
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// IDispatch
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr

	// ITaskDefinition
	get_RegistrationInfo uintptr
	put_RegistrationInfo uintptr
	get_Triggers         uintptr
	put_Triggers         uintptr
	get_Settings         uintptr
	put_Settings         uintptr
	get_Data             uintptr
	put_Data             uintptr
	get_Principal        uintptr
	put_Principal        uintptr
	get_Actions          uintptr
	put_Actions          uintptr
	get_XmlText          uintptr
	put_XmlText          uintptr
}

type ITaskScheduler struct {
	vtbl *iTaskSchedulerVtbl
}

type iTaskSchedulerVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	SetTargetComputer uintptr
	GetTargetComputer uintptr
	Enum              uintptr
	Activate          uintptr
	Delete            uintptr
	NewWorkItem       uintptr
	AddWorkItem       uintptr
	IsOfType          uintptr
}

type ITask struct {
	vtbl *iTaskVtbl
}

type iTaskVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	CreateTrigger         uintptr
	DeleteTrigger         uintptr
	GetTriggerCount       uintptr
	GetTrigger            uintptr
	GetTriggerString      uintptr
	GetRunTimes           uintptr
	GetNextRunTime        uintptr
	SetIdleWait           uintptr
	GetIdleWait           uintptr
	Run                   uintptr
	Terminate             uintptr
	EditWorkItem          uintptr
	GetMostRecentRunTime  uintptr
	GetStatus             uintptr
	GetExitCode           uintptr
	SetComment            uintptr
	GetComment            uintptr
	SetCreator            uintptr
	GetCreator            uintptr
	SetWorkItemData       uintptr
	GetWorkItemData       uintptr
	SetErrorRetryCount    uintptr
	GetErrorRetryCount    uintptr
	SetErrorRetryInterval uintptr
	GetErrorRetryInterval uintptr
	SetFlags              uintptr
	GetFlags              uintptr
	SetAccountInformation uintptr
	GetAccountInformation uintptr

	SetApplicationName  uintptr
	GetApplicationName  uintptr
	SetParameters       uintptr
	GetParameters       uintptr
	SetWorkingDirectory uintptr
	GetWorkingDirectory uintptr
	SetPriority         uintptr
	GetPriority         uintptr
	SetTaskFlags        uintptr
	GetTaskFlags        uintptr
	SetMaxRunTime       uintptr
	GetMaxRunTime       uintptr
}

type IPersistFile struct {
	vtbl *iPersistFileVtbl
}

type iPersistFileVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	GetClassID uintptr

	IsDirty       uintptr
	Load          uintptr
	Save          uintptr
	SaveCompleted uintptr
	GetCurFile    uintptr
}

func (obj *ITaskService) Release() HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return HRESULT(ret)
}

func (obj *ITaskService) Connect(serverName, user, domain, password VARIANT) HRESULT {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.Connect,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&serverName[0])),
		uintptr(unsafe.Pointer(&user[0])),
		uintptr(unsafe.Pointer(&domain[0])),
		uintptr(unsafe.Pointer(&password[0])),
		0,
	)
	return HRESULT(ret)
}

func (obj *ITaskService) GetFolder(path BSTR, ppFolder **ITaskFolder) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetFolder,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(ppFolder)),
	)
	return HRESULT(ret)
}

// HRESULT STDMETHODCALLTYPE NewWorkItem(
// 	/* [in] */ LPCWSTR pwszTaskName,
// 	/* [in] */ REFCLSID rclsid,
// 	/* [in] */ REFIID riid,
// 	/* [out] */ IUnknown **ppUnk)
func (obj *ITaskScheduler) NewWorkItem(pwszTaskName *uint16, rclsid *GUID, riid *GUID, ppUnk **ITask) HRESULT {
	hr, _, _ := syscall.Syscall6(
		obj.vtbl.NewWorkItem,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(pwszTaskName)),
		uintptr(unsafe.Pointer(rclsid)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppUnk)),
		0,
	)
	return HRESULT(hr)
}

func (obj *ITaskScheduler) Release() HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return HRESULT(ret)
}

func (obj *ITask) QueryInterface(riid *GUID, ppvObject **IPersistFile) HRESULT {
	hr, _, _ := syscall.Syscall(
		obj.vtbl.QueryInterface,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObject)),
	)
	return HRESULT(hr)
}

func (obj *ITask) Release() HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return HRESULT(ret)
}

func (obj *IPersistFile) Save(pszFileName LPCOLESTR, fRemember BOOL) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Save,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(pszFileName)),
		uintptr(fRemember),
	)
	return HRESULT(ret)
}

func (obj *IPersistFile) Release() HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return HRESULT(ret)
}

func (obj *ITaskFolder) DeleteTask(name BSTR, flags int32) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.DeleteTask,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(name)),
		uintptr(flags),
	)
	return HRESULT(ret)
}

//sys CoInitialize(pvReserved LPVOID) (hResult HRESULT) = ole32.CoInitialize
//sys CoInitializeEx(pvReserved LPVOID, dwCoInit DWORD) (hResult HRESULT) = ole32.CoInitializeEx
//sys CoUninitialize() = ole32.CoUninitialize
//sys CoCreateInstance(rclsid REFCLSID, pUnkOuter LPUNKNOWN, dwClsContext DWORD, riid REFIID, ppv LPVOID) (hResult HRESULT) = ole32.CoCreateInstance
//sys VariantInit(pvarg *VARIANT) = oleaut32.VariantInit

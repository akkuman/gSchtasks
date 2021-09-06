package win32

import (
	"syscall"
	"unsafe"
)

const NULL = 0
const TRUE = 1
const CLSCTX_INPROC_SERVER = 0x1
const TASK_CREATE_OR_UPDATE = 0x6

var FAILED = func(hr HRESULT) bool { return hr < 0 }

type (
	VARIANT_BOOL       int16
	LPVOID             uintptr
	LMSTR              *uint16
	DWORD              uint32
	LPBYTE             *byte
	LPDWORD            *uint32
	BOOL               int
	TASK_TRIGGER_TYPE2 int
	TASK_ACTION_TYPE   int
	TASK_LOGON_TYPE    int
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

var (
	VARIANT_TRUE  VARIANT_BOOL = -1
	VARIANT_FALSE VARIANT_BOOL = 0

	TASK_TRIGGER_TIME TASK_TRIGGER_TYPE2 = 1

	TASK_ACTION_EXEC TASK_ACTION_TYPE = 0

	TASK_LOGON_INTERACTIVE_TOKEN TASK_LOGON_TYPE = 3
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

type IRegistrationInfo struct {
	vtbl *iRegistrationInfoVtbl
}

type iRegistrationInfoVtbl struct {
	// IUnknown
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// IDispatch
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr

	// IRegistrationInfo
	get_Description        uintptr
	put_Description        uintptr
	get_Author             uintptr
	put_Author             uintptr
	get_Version            uintptr
	put_Version            uintptr
	get_Date               uintptr
	put_Date               uintptr
	get_Documentation      uintptr
	put_Documentation      uintptr
	get_XmlText            uintptr
	put_XmlText            uintptr
	get_URI                uintptr
	put_URI                uintptr
	get_SecurityDescriptor uintptr
	put_SecurityDescriptor uintptr
	get_Source             uintptr
	put_Source             uintptr
}

type IPrincipal struct {
	vtbl *iPrincipalVtbl
}

type iPrincipalVtbl struct {
	// IUnknown
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// IDispatch
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr

	// IPrincipal
	get_Id          uintptr
	put_Id          uintptr
	get_DisplayName uintptr
	put_DisplayName uintptr
	get_UserId      uintptr
	put_UserId      uintptr
	get_LogonType   uintptr
	put_LogonType   uintptr
	get_GroupId     uintptr
	put_GroupId     uintptr
	get_RunLevel    uintptr
	put_RunLevel    uintptr
}

type ITaskSettings struct {
	vtbl *iTaskSettingsVtbl
}

type iTaskSettingsVtbl struct {
	// IUnknown
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// IDispatch
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr

	// ITaskSettings
	get_AllowDemandStart           uintptr
	put_AllowDemandStart           uintptr
	get_RestartInterval            uintptr
	put_RestartInterval            uintptr
	get_RestartCount               uintptr
	put_RestartCount               uintptr
	get_MultipleInstances          uintptr
	put_MultipleInstances          uintptr
	get_StopIfGoingOnBatteries     uintptr
	put_StopIfGoingOnBatteries     uintptr
	get_DisallowStartIfOnBatteries uintptr
	put_DisallowStartIfOnBatteries uintptr
	get_AllowHardTerminate         uintptr
	put_AllowHardTerminate         uintptr
	get_StartWhenAvailable         uintptr
	put_StartWhenAvailable         uintptr
	get_XmlText                    uintptr
	put_XmlText                    uintptr
	get_RunOnlyIfNetworkAvailable  uintptr
	put_RunOnlyIfNetworkAvailable  uintptr
	get_ExecutionTimeLimit         uintptr
	put_ExecutionTimeLimit         uintptr
	get_Enabled                    uintptr
	put_Enabled                    uintptr
	get_DeleteExpiredTaskAfter     uintptr
	put_DeleteExpiredTaskAfter     uintptr
	get_Priority                   uintptr
	put_Priority                   uintptr
	get_Compatibility              uintptr
	put_Compatibility              uintptr
	get_Hidden                     uintptr
	put_Hidden                     uintptr
	get_IdleSettings               uintptr
	put_IdleSettings               uintptr
	get_RunOnlyIfIdle              uintptr
	put_RunOnlyIfIdle              uintptr
	get_WakeToRun                  uintptr
	put_WakeToRun                  uintptr
	get_NetworkSettings            uintptr
	put_NetworkSettings            uintptr
}

type IIdleSettings struct {
	vtbl *iIdleSettingsVtbl
}

type iIdleSettingsVtbl struct {
	// IUnknown
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// IDispatch
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr

	// IIdleSettings
	get_IdleDuration  uintptr
	put_IdleDuration  uintptr
	get_WaitTimeout   uintptr
	put_WaitTimeout   uintptr
	get_StopOnIdleEnd uintptr
	put_StopOnIdleEnd uintptr
	get_RestartOnIdle uintptr
	put_RestartOnIdle uintptr
}

type ITriggerCollection struct {
	vtbl *iTriggerCollectionVtbl
}

type iTriggerCollectionVtbl struct {
	// IUnknown
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// IDispatch
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr

	// ITriggerCollection
	get_Count    uintptr
	get_Item     uintptr
	get__NewEnum uintptr
	Create       uintptr
	Remove       uintptr
	Clear        uintptr
}

type ITrigger struct {
	vtbl *iTriggerVtbl
}

type iTriggerVtbl struct {
	// IUnknown
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// IDispatch
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr

	// ITrigger
	get_Type               uintptr
	get_Id                 uintptr
	put_Id                 uintptr
	get_Repetition         uintptr
	put_Repetition         uintptr
	get_ExecutionTimeLimit uintptr
	put_ExecutionTimeLimit uintptr
	get_StartBoundary      uintptr
	put_StartBoundary      uintptr
	get_EndBoundary        uintptr
	put_EndBoundary        uintptr
	get_Enabled            uintptr
	put_Enabled            uintptr
}

type ITimeTrigger struct {
	vtbl *iTimeTriggerVtbl
}

type iTimeTriggerVtbl struct {
	// IUnknown
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// IDispatch
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr

	// ITrigger
	get_Type               uintptr
	get_Id                 uintptr
	put_Id                 uintptr
	get_Repetition         uintptr
	put_Repetition         uintptr
	get_ExecutionTimeLimit uintptr
	put_ExecutionTimeLimit uintptr
	get_StartBoundary      uintptr
	put_StartBoundary      uintptr
	get_EndBoundary        uintptr
	put_EndBoundary        uintptr
	get_Enabled            uintptr
	put_Enabled            uintptr

	// ITimeTrigger
	get_RandomDelay uintptr
	put_RandomDelay uintptr
}

type IActionCollection struct {
	vtbl *iActionCollectionVtbl
}

type iActionCollectionVtbl struct {
	// IUnknown
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// IDispatch
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr

	// IActionCollection
	get_Count    uintptr
	get_Item     uintptr
	get__NewEnum uintptr
	get_XmlText  uintptr
	put_XmlText  uintptr
	Create       uintptr
	Remove       uintptr
	Clear        uintptr
	get_Context  uintptr
	put_Context  uintptr
}

type IAction struct {
	vtbl *iActionVtbl
}

type iActionVtbl struct {
	// IUnknown
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// IDispatch
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr

	// IAction
	get_Id   uintptr
	put_Id   uintptr
	get_Type uintptr
}

type IExecAction struct {
	vtbl *iExecActionVtbl
}

type iExecActionVtbl struct {
	// IUnknown
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// IDispatch
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr

	// IAction
	get_Id   uintptr
	put_Id   uintptr
	get_Type uintptr

	// IExecAction
	get_Path             uintptr
	put_Path             uintptr
	get_Arguments        uintptr
	put_Arguments        uintptr
	get_WorkingDirectory uintptr
	put_WorkingDirectory uintptr
}

type IRegisteredTask struct {
	vtbl *iRegisteredTaskVtbl
}

type iRegisteredTaskVtbl struct {
	// IUnknown
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// IDispatch
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr

	// IRegisteredTask
	get_Name               uintptr
	get_Path               uintptr
	get_State              uintptr
	get_Enabled            uintptr
	put_Enabled            uintptr
	Run                    uintptr
	RunEx                  uintptr
	GetInstances           uintptr
	get_LastRunTime        uintptr
	get_LastTaskResult     uintptr
	get_NumberOfMissedRuns uintptr
	get_NextRunTime        uintptr
	get_Definition         uintptr
	get_Xml                uintptr
	GetSecurityDescriptor  uintptr
	SetSecurityDescriptor  uintptr
	Stop                   uintptr
	GetRunTimes            uintptr
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

func (obj *ITaskService) NewTask(flags DWORD, ppDefinition **ITaskDefinition) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.NewTask,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(flags),
		uintptr(unsafe.Pointer(ppDefinition)),
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

func (obj *ITaskFolder) Release() HRESULT {
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

func (obj *ITaskFolder) RegisterTaskDefinition(
	path BSTR,
	pDefinition *ITaskDefinition,
	flags int,
	userId VARIANT,
	password VARIANT,
	logonType TASK_LOGON_TYPE,
	sddl VARIANT,
	ppTask **IRegisteredTask,
) HRESULT {
	ret, _, _ := syscall.Syscall9(
		obj.vtbl.RegisterTaskDefinition,
		9,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(pDefinition)),
		uintptr(flags),
		uintptr(unsafe.Pointer(&userId[0])),
		uintptr(unsafe.Pointer(&password[0])),
		uintptr(logonType),
		uintptr(unsafe.Pointer(&sddl[0])),
		uintptr(unsafe.Pointer(ppTask)),
	)
	return HRESULT(ret)
}

func (obj *IRegistrationInfo) PutAuthor(author BSTR) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.put_Author,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(author)),
		0,
	)
	return HRESULT(ret)
}

func (obj *IRegistrationInfo) Release() HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return HRESULT(ret)
}

func (obj *ITaskDefinition) GetRegistrationInfo(ppRegistrationInfo **IRegistrationInfo) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.get_RegistrationInfo,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(ppRegistrationInfo)),
		0,
	)
	return HRESULT(ret)
}

func (obj *ITaskDefinition) Release() HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return HRESULT(ret)
}

func (obj *ITaskDefinition) GetPrincipal(ppPrincipal **IPrincipal) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.get_Principal,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(ppPrincipal)),
		0,
	)
	return HRESULT(ret)
}

func (obj *ITaskDefinition) GetSettings(ppSettings **ITaskSettings) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.get_Settings,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(ppSettings)),
		0,
	)
	return HRESULT(ret)
}

func (obj *ITaskDefinition) GetTriggers(ppTriggers **ITriggerCollection) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.get_Triggers,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(ppTriggers)),
		0,
	)
	return HRESULT(ret)
}

func (obj *ITaskDefinition) GetActions(ppActions **IActionCollection) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.get_Actions,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(ppActions)),
		0,
	)
	return HRESULT(ret)
}

func (obj *IPrincipal) PutLogonType(logon TASK_LOGON_TYPE) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.put_LogonType,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(logon),
		0,
	)
	return HRESULT(ret)
}

func (obj *IPrincipal) Release() HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return HRESULT(ret)
}

func (obj *ITaskSettings) PutStartWhenAvailable(startWhenAvailable VARIANT_BOOL) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.put_StartWhenAvailable,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(startWhenAvailable),
		0,
	)
	return HRESULT(ret)
}

func (obj *ITaskSettings) Release() HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return HRESULT(ret)
}

func (obj *ITaskSettings) GetIdleSettings(ppIdleSettings **IIdleSettings) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.get_IdleSettings,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(ppIdleSettings)),
		0,
	)
	return HRESULT(ret)
}

func (obj *IIdleSettings) PutWaitTimeout(timeout BSTR) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.put_WaitTimeout,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(timeout)),
		0,
	)
	return HRESULT(ret)
}

func (obj *IIdleSettings) Release() HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return HRESULT(ret)
}

func (obj *ITriggerCollection) Create(_type TASK_TRIGGER_TYPE2, ppTrigger **ITrigger) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Create,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(_type),
		uintptr(unsafe.Pointer(ppTrigger)),
	)
	return HRESULT(ret)
}

func (obj *ITriggerCollection) Release() HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return HRESULT(ret)
}

func (obj *ITrigger) QueryInterface(riid REFIID, ppvObject **ITimeTrigger) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.QueryInterface,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(riid),
		uintptr(unsafe.Pointer(ppvObject)),
	)
	return HRESULT(ret)
}

func (obj *ITrigger) Release() HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return HRESULT(ret)
}

func (obj *ITimeTrigger) PutId(id BSTR) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.put_Id,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(id)),
		0,
	)
	return HRESULT(ret)
}

func (obj *ITimeTrigger) PutEndBoundary(end BSTR) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.put_EndBoundary,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(end)),
		0,
	)
	return HRESULT(ret)
}

func (obj *ITimeTrigger) PutStartBoundary(start BSTR) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.put_StartBoundary,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(start)),
		0,
	)
	return HRESULT(ret)
}

func (obj *IActionCollection) Create(_type TASK_ACTION_TYPE, ppAction **IAction) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Create,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(_type),
		uintptr(unsafe.Pointer(ppAction)),
	)
	return HRESULT(ret)
}

func (obj *IActionCollection) Release() HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return HRESULT(ret)
}

func (obj *IAction) QueryInterface(riid REFIID, ppvObject **IExecAction) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.QueryInterface,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(riid),
		uintptr(unsafe.Pointer(ppvObject)),
	)
	return HRESULT(ret)
}

func (obj *IAction) Release() HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return HRESULT(ret)
}

func (obj *IExecAction) PutPath(path BSTR) HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.put_Path,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(path)),
		0,
	)
	return HRESULT(ret)
}

func (obj *IExecAction) Release() HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return HRESULT(ret)
}

func (obj *IRegisteredTask) Release() HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return HRESULT(ret)
}

//sys CoInitialize(pvReserved LPVOID) (hResult HRESULT) = ole32.CoInitialize
//sys CoInitializeEx(pvReserved LPVOID, dwCoInit DWORD) (hResult HRESULT) = ole32.CoInitializeEx
//sys CoUninitialize() = ole32.CoUninitialize
//sys CoCreateInstance(rclsid REFCLSID, pUnkOuter LPUNKNOWN, dwClsContext DWORD, riid REFIID, ppv LPVOID) (hResult HRESULT) = ole32.CoCreateInstance
//sys VariantInit(pvarg *VARIANT) = oleaut32.VariantInit
//sys SysAllocString(psz uintptr) (ss uintptr) = oleaut32.SysAllocString

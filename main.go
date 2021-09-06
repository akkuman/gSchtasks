package main

import (
	"fmt"
	"gSchtasks/win32"
	"syscall"
	"unsafe"
)

func main() {
	// 初始化 COM
	hr := win32.CoInitialize(win32.NULL)
	if win32.FAILED(hr) {
		fmt.Printf("\nCoInitializeEx failed: %x", hr)
		return
	}
	// 任务名称
	wszTaskName, err := syscall.UTF16PtrFromString("Time Trigger Test Task")
	if err != nil {
		panic(err)
	}
	// 触发任务时执行的可执行文件的路径
	wstrExecutablePath, err := syscall.UTF16PtrFromString("G:\\Windows\\System32\\notepad.exe")
	if err != nil {
		panic(err)
	}

	// 创建一个任务服务的实例
	var pService *win32.ITaskService
	hr = win32.CoCreateInstance(
		win32.REFCLSID(unsafe.Pointer(win32.NewGUID("0f87369f-a4e5-4cfc-bd3e-73e6154572dd"))),
		win32.LPUNKNOWN(win32.NULL),
		win32.DWORD(win32.CLSCTX_INPROC_SERVER),
		win32.REFIID(unsafe.Pointer(win32.NewGUID("2faba4c7-4da9-4013-9697-20cc3fd40f85"))),
		win32.LPVOID(unsafe.Pointer(&pService)),
	)
	if win32.FAILED(hr) {
		fmt.Printf("Failed to create an instance of ITaskService: %x", hr)
		win32.CoUninitialize()
		return
	}

	// 连接到任务服务
	var variants [4]win32.VARIANT
	for i := range variants {
		win32.VariantInit(&variants[i])
	}
	hr = pService.Connect(
		variants[0],
		variants[1],
		variants[2],
		variants[3],
	)
	if win32.FAILED(hr) {
		fmt.Printf("ITaskService::Connect failed: %x", hr)
		pService.Release()
		win32.CoUninitialize()
		return
	}

	// 获取指向根部任务文件夹的指针。此文件夹将保存已注册的新任务
	var pRootFolder *win32.ITaskFolder
	path, err := syscall.UTF16PtrFromString("\\")
	if err != nil {
		panic(err)
	}
	hr = pService.GetFolder(
		win32.BSTR(path),
		&pRootFolder,
	)
	if win32.FAILED(hr) {
		fmt.Printf("Cannot get Root folder pointer: %x", hr)
		pService.Release()
		win32.CoUninitialize()
		return
	}

	// 如果存在相同的任务就先删除
	pRootFolder.DeleteTask(win32.BSTR(wszTaskName), 0)

	// 创建任务定义对象
	var pTask *win32.ITaskDefinition
	hr = pService.NewTask(0, &pTask)

	pService.Release() // 清除 COM，指针不再使用

	if win32.FAILED(hr) {
		fmt.Printf("Failed to CoCreate an instance of the TaskService class: %x", hr)
		pRootFolder.Release()
		win32.CoUninitialize()
		return
	}

	// 获取设置任务定义的注册信息
	var pRegInfo *win32.IRegistrationInfo
	hr = pTask.GetRegistrationInfo(&pRegInfo)
	if win32.FAILED(hr) {
		fmt.Printf("\nCannot get identification pointer: %x", hr)
		pRootFolder.Release()
		pTask.Release()
		win32.CoUninitialize()
		return
	}

	authorName, _ := syscall.UTF16PtrFromString("Author Name")
	hr = pRegInfo.PutAuthor(win32.BSTR(authorName))
	pRegInfo.Release()
	if win32.FAILED(hr) {
		fmt.Printf("\nCannot put identification info: %x", hr)
		pRootFolder.Release()
		pTask.Release()
		win32.CoUninitialize()
		return
	}

	// 为任务创建委托人--这些证书将被传递给RegisterTaskDefinition的证书覆盖
	var pPrincipal *win32.IPrincipal
	hr = pTask.GetPrincipal(&pPrincipal)
	if win32.FAILED(hr) {
		fmt.Printf("\nCannot get principal pointer: %x", hr)
		pRootFolder.Release()
		pTask.Release()
		win32.CoUninitialize()
		return
	}

	// 将主登录类型设置为交互式登录
	hr = pPrincipal.PutLogonType(win32.TASK_LOGON_INTERACTIVE_TOKEN)
	pPrincipal.Release()
	if win32.FAILED(hr) {
		fmt.Printf("\nCannot put principal info: %x", hr)
		pRootFolder.Release()
		pTask.Release()
		win32.CoUninitialize()
		return
	}

	// 给任务创建配置
	var pSettings *win32.ITaskSettings
	hr = pTask.GetSettings(&pSettings)
	if win32.FAILED(hr) {
		fmt.Printf("\nCannot get settings pointer: %x", hr)
		pRootFolder.Release()
		pTask.Release()
		win32.CoUninitialize()
		return
	}

	// 给任务的配置项设置值
	hr = pSettings.PutStartWhenAvailable(win32.VARIANT_TRUE)
	pSettings.Release()
	if win32.FAILED(hr) {
		fmt.Printf("\nCannot put setting information: %x", hr)
		pRootFolder.Release()
		pTask.Release()
		win32.CoUninitialize()
		return
	}

	// 给任务设置 idle 配置
	var pIdleSettings *win32.IIdleSettings
	hr = pSettings.GetIdleSettings(&pIdleSettings)
	if win32.FAILED(hr) {
		fmt.Printf("\nCannot get idle setting information: %x", hr)
		pRootFolder.Release()
		pTask.Release()
		win32.CoUninitialize()
		return
	}

	timeout, _ := syscall.UTF16PtrFromString("PT5M")
	hr = pIdleSettings.PutWaitTimeout(timeout)
	pIdleSettings.Release()
	if win32.FAILED(hr) {
		fmt.Printf("\nCannot put idle setting information: %x", hr)
		pRootFolder.Release()
		pTask.Release()
		win32.CoUninitialize()
		return
	}

	// 获取触发器集合以插入时间触发器
	var pTriggerCollection *win32.ITriggerCollection
	hr = pTask.GetTriggers(&pTriggerCollection)
	if win32.FAILED(hr) {
		fmt.Printf("\nCannot get trigger collection: %x", hr)
		pRootFolder.Release()
		pTask.Release()
		win32.CoUninitialize()
		return
	}

	// 给任务添加时间触发器
	var pTrigger *win32.ITrigger
	hr = pTriggerCollection.Create(win32.TASK_TRIGGER_TIME, &pTrigger)
	pTriggerCollection.Release()
	if win32.FAILED(hr) {
		fmt.Printf("\nCannot create trigger: %x", hr)
		pRootFolder.Release()
		pTask.Release()
		win32.CoUninitialize()
		return
	}

	var pTimeTrigger *win32.ITimeTrigger
	hr = pTrigger.QueryInterface(win32.REFIID(unsafe.Pointer(win32.NewGUID("b45747e0-eba7-4276-9f29-85c5bb300006"))), &pTimeTrigger)
	pTrigger.Release()
	if win32.FAILED(hr) {
		fmt.Printf("\nQueryInterface call failed for ITimeTrigger: %x", hr)
		pRootFolder.Release()
		pTask.Release()
		win32.CoUninitialize()
		return
	}

	id, _ := syscall.UTF16PtrFromString("Trigger1")
	hr = pTimeTrigger.PutId(id)
	if win32.FAILED(hr) {
		fmt.Printf("\nCannot put trigger ID: %x", hr)
	}

	end, _ := syscall.UTF16PtrFromString("2015-05-02T08:00:00")
	hr = pTimeTrigger.PutEndBoundary(end)
	if win32.FAILED(hr) {
		fmt.Printf("\nCannot put end boundary on trigger: %x", hr)
	}

	// 设置任务在某个时间开始
	// 时间格式应该是YYY-MM-DDTHH:MM:SS(+-)(timezone)
	start, _ := syscall.UTF16PtrFromString("2005-01-01T12:05:00")
	hr = pTimeTrigger.PutStartBoundary(start)
	if win32.FAILED(hr) {
		fmt.Printf("\nCannot add start boundary to trigger: %x", hr)
		pRootFolder.Release()
		pTask.Release()
		win32.CoUninitialize()
		return
	}

	// 给任务添加一个 action，这个样例任务将运行 notepad.exe
	var pActionCollection *win32.IActionCollection
	// 获取任务动作集合指针
	hr = pTask.GetActions(&pActionCollection)
	if win32.FAILED(hr) {
		fmt.Printf("\nCannot get Task collection pointer: %x", hr)
		pRootFolder.Release()
		pTask.Release()
		win32.CoUninitialize()
		return
	}

	// 创建动作，指定它是一个可执行的动作
	var pAction *win32.IAction
	hr = pActionCollection.Create(win32.TASK_ACTION_EXEC, &pAction)
	pActionCollection.Release()
	if win32.FAILED(hr) {
		fmt.Printf("\nCannot create the action: %x", hr)
		pRootFolder.Release()
		pTask.Release()
		win32.CoUninitialize()
		return
	}

	// QI for the executable task pointer.
	var pExecAction *win32.IExecAction
	hr = pAction.QueryInterface(win32.REFIID(unsafe.Pointer(win32.NewGUID("4c3d624d-fd6b-49a3-b9b7-09cb3cd3f047"))), &pExecAction)
	pAction.Release()
	if win32.FAILED(hr) {
		fmt.Printf("\nQueryInterface call failed for IExecAction: %x", hr)
		pRootFolder.Release()
		pTask.Release()
		win32.CoUninitialize()
		return
	}

	// 将可执行文件的路径设为notepad.exe
	hr = pExecAction.PutPath(wstrExecutablePath)
	pExecAction.Release()
	if win32.FAILED(hr) {
		fmt.Printf("\nCannot put action path: %x", hr)
		pRootFolder.Release()
		pTask.Release()
		win32.CoUninitialize()
		return
	}

	// 将任务保存在根文件夹中
	var pRegisteredTask *win32.IRegisteredTask
	var userId, password win32.VARIANT
	win32.VariantInit(&userId)
	win32.VariantInit(&password)
	hr = pRootFolder.RegisterTaskDefinition(
		wszTaskName,
		pTask,
		win32.TASK_CREATE_OR_UPDATE,
		userId,
		password,
		win32.TASK_LOGON_INTERACTIVE_TOKEN,
		win32.NewVariantWithStr(""),
		&pRegisteredTask,
	)

	if win32.FAILED(hr) {
		fmt.Printf("\nError saving the Task : %x", hr)
		pRootFolder.Release()
		pTask.Release()
		win32.CoUninitialize()
		return
	}

	fmt.Printf("\n Success! Task successfully registered. ")

	//  Clean up.
	pRootFolder.Release()
	pTask.Release()
	pRegisteredTask.Release()
	win32.CoUninitialize()
	return
}

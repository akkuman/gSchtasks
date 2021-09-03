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
		fmt.Printf("ITaskService::Connect failed: %x", hr)
		pService.Release()
		win32.CoUninitialize()
		return
	}

	// 如果存在相同的任务就先删除
	pRootFolder.DeleteTask(win32.BSTR(wszTaskName), 0)

}

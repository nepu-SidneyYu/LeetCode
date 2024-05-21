package main

import (
	"fmt"
	"log"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func main() {
	ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
	defer ole.CoUninitialize()

	// 创建WMI服务
	unknown, err := oleutil.CreateObject("WbemScripting.SWbemLocator")
	if err != nil {
		log.Fatalf("CreateObject: %v", err)
	}
	defer unknown.Release()

	wmi, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		log.Fatalf("QueryInterface: %v", err)
	}
	defer wmi.Release()

	serviceRaw, err := oleutil.CallMethod(wmi, "ConnectServer")
	if err != nil {
		log.Fatalf("ConnectServer: %v", err)
	}
	service := serviceRaw.ToIDispatch()
	defer serviceRaw.Clear()

	// 查询Win32_PnPEntity类以获取设备信息
	query := "SELECT * FROM Win32_PnPEntity WHERE PNPDeviceID LIKE '%USB%'"
	rawSet, err := oleutil.CallMethod(service, "ExecQuery", query)
	if err != nil {
		log.Fatalf("ExecQuery: %v", err)
	}
	defer rawSet.Clear()
	result := rawSet.ToIDispatch()
	defer result.Release()

	for {
		enumProperty, err := oleutil.GetProperty(result, "_NewEnum")
		if err != nil {
			log.Fatalf("GetProperty: %v", err)
		}
		enum := enumProperty.ToIUnknown()
		defer enumProperty.Clear()

		unknowns, err := oleutil.EnumObjects(enum)
		if err != nil {
			log.Fatalf("EnumObjects: %v", err)
		}

		// 遍历设备列表
		for _, item := range unknowns {
			device := item.ToIDispatch()
			pnpDeviceID, err := oleutil.GetProperty(device, "PNPDeviceID")
			if err == nil {
				fmt.Printf("Device inserted or removed: %v\n", pnpDeviceID.ToString())
			}
			device.Release()
		}

		ole.Sleep(1000) // 休眠1秒钟以减小CPU负载
	}
}

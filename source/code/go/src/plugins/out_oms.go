package main

import (
	"github.com/fluent/fluent-bit-go/output"
)
import (
	"C"
	"unsafe"
)

// FLBPluginRegister registers the plugin
func FLBPluginRegister(ctx unsafe.Pointer) int {
	return output.FLBPluginRegister(ctx, "oms", "Stdout GO!")
}

// FLBPluginInit initializes the plugin
// (fluentbit will call this)
// ctx (context) pointer to fluentbit context (state/ c code)
func FLBPluginInit(ctx unsafe.Pointer) int {
	Log("Initializing out_oms go plugin for fluentbit")
	PluginConfiguration = ReadConfig("/etc/opt/microsoft/docker-cimprov/out_oms.conf")
	CreateHTTPClient()
	go initMaps()
	go updateIgnoreContainerIds()
	return output.FLB_OK
}

// FLBPluginFlush flushes the data in the stream
func FLBPluginFlush(data unsafe.Pointer, length C.int, tag *C.char) int {
	var count int
	var ret int
	var record map[interface{}]interface{}
	var records []map[interface{}]interface{}

	// Create Fluent Bit decoder
	dec := output.NewDecoder(data, int(length))

	// Iterate Records
	count = 0
	for {
		// Extract Record
		ret, _, record = output.GetRecord(dec)
		if ret != 0 {
			break
		}
		records = append(records, record)
		count++
	}
	return PostDataHelper(records)
}

// FLBPluginExit exits the plugin
func FLBPluginExit() int {
	return output.FLB_OK
}

func main() {
}

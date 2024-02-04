package catalog

import (
	"fmt"
	"os"

	mnt "github.com/moby/sys/mountinfo"
	support "github.com/mt1976/admin_me/support"
	mem "github.com/shirou/gopsutil/mem"
	cpu "github.com/shirou/gopsutil/v3/cpu"
	dsk "github.com/shirou/gopsutil/v3/disk"
	hst "github.com/shirou/gopsutil/v3/host"
)

// Catalogs the resources available on a given system.

var debugMode bool = false
var crt support.Crt

type info struct {
	data []string
}

func Run(crtIn support.Crt, debug bool, path string) {

	debugMode = debug
	crt = crtIn
	hostname := support.GetHostname(crt)
	outputFilename := "catalog_" + hostname + "_" + support.GetTimeStamp() + ".info"

	info := NewInfo()
	//X := T

	info.breakData("Cataloging system resources")
	//fmt.Println(crt.PR("Output file = "+outputFilename, T))
	crt.Print("Output file : " + outputFilename)
	crt.Break()
	info.storeData("Hostname", support.GetHostname(crt))
	info.storeData("Machine Name", support.GetSystemInfo(crt))
	info.storeData("Username", support.GetUsername(crt))
	info.storeData("Current Path", path)

	m, _ := mem.VirtualMemory()
	info.storeData("Total Memory", support.Human(m.Total))
	info.storeData("Available Memory", support.Human(m.Available))
	info.storeData("Used Memory", support.Human(m.Used))
	info.storeData("Used Percent", fmt.Sprintf("%f", m.UsedPercent))
	info.storeData("Free Memory", support.Human(m.Free))
	c, _ := cpu.Info()
	//range through each cpu and store the cpu details

	for _, cpu := range c {
		crt.Break()
		info.storeData(fmt.Sprintf("CPU %d", cpu.CPU), support.Human(cpu.CPU))
		info.storeData(fmt.Sprintf("CPU %d Model", cpu.CPU), support.Human(cpu.ModelName))
		info.storeData(fmt.Sprintf("CPU %d Cores", cpu.CPU), support.Human(cpu.Cores))
		info.storeData(fmt.Sprintf("CPU %d Mhz", cpu.CPU), support.Human(cpu.Mhz))
		info.storeData(fmt.Sprintf("CPU %d Cache Size", cpu.CPU), support.Human(cpu.CacheSize))
		info.storeData(fmt.Sprintf("CPU %d Flags", cpu.CPU), support.Human(cpu.Flags))
		info.storeData(fmt.Sprintf("CPU %d Stepping", cpu.CPU), support.Human(cpu.Stepping))
		info.storeData(fmt.Sprintf("CPU %d Vendor ID", cpu.CPU), support.Human(cpu.VendorID))
		info.storeData(fmt.Sprintf("CPU %d Family", cpu.CPU), support.Human(cpu.Family))
		info.storeData(fmt.Sprintf("CPU %d Model", cpu.CPU), support.Human(cpu.Model))
		info.storeData(fmt.Sprintf("CPU %d Physical ID", cpu.CPU), support.Human(cpu.PhysicalID))
		info.storeData(fmt.Sprintf("CPU %d Core ID", cpu.CPU), support.Human(cpu.CoreID))
		info.storeData(fmt.Sprintf("CPU %d Microcode", cpu.CPU), support.Human(cpu.Microcode))
		info.storeData(fmt.Sprintf("CPU %d Model Name", cpu.CPU), support.Human(cpu.ModelName))
	}

	ht, _ := hst.Info()
	//range through each host and print host info
	crt.Break()
	info.storeData(fmt.Sprintf("Host %s", "ID"), support.Human(ht.HostID))
	info.storeData(fmt.Sprintf("Host %s Hostname", ""), support.Human(ht.Hostname))
	info.storeData(fmt.Sprintf("Host %s Uptime", ""), support.Human(ht.Uptime))
	info.storeData(fmt.Sprintf("Host %s Boot Time", ""), support.Human(ht.BootTime))
	info.storeData(fmt.Sprintf("Host %s Procs", ""), support.Human(ht.Procs))
	info.storeData(fmt.Sprintf("Host %s OS", ""), support.Human(ht.OS))
	info.storeData(fmt.Sprintf("Host %s Platform", ""), support.Human(ht.Platform))
	info.storeData(fmt.Sprintf("Host %s Platform Family", ""), support.Human(ht.PlatformFamily))
	info.storeData(fmt.Sprintf("Host %s Platform Version", ""), support.Human(ht.PlatformVersion))
	info.storeData(fmt.Sprintf("Host %s Kernel Version", ""), support.Human(ht.KernelVersion))
	info.storeData(fmt.Sprintf("Host %s Virtualization System", ""), support.Human(ht.VirtualizationSystem))
	info.storeData(fmt.Sprintf("Host %s Virtualization Role", ""), support.Human(ht.VirtualizationRole))
	//info.storeData( T,  "Host Info", T.Human( h))

	v, _ := mnt.GetMounts(nil)
	//zz := 0
	for zz, v := range v {
		//info.storeData( T,  fmt.Sprintf("Mount %d", zz), T.Human( v))
		crt.Break()
		info.storeData(fmt.Sprintf("Mount %d ID", zz), support.Human(v.ID))
		info.storeData(fmt.Sprintf("Mount %d Major", zz), support.Human(v.Major))
		info.storeData(fmt.Sprintf("Mount %d Minor", zz), support.Human(v.Minor))
		info.storeData(fmt.Sprintf("Mount %d Root", zz), support.Human(v.Root))
		info.storeData(fmt.Sprintf("Mount %d Parent", zz), support.Human(v.Parent))
		info.storeData(fmt.Sprintf("Mount %d Mountpoint", zz), support.Human(v.Mountpoint))
		info.storeData(fmt.Sprintf("Mount %d Options", zz), support.Human(v.Options))
		info.storeData(fmt.Sprintf("Mount %d Optional", zz), support.Human(v.Optional))
		info.storeData(fmt.Sprintf("Mount %d FSType", zz), support.Human(v.FSType))
		info.storeData(fmt.Sprintf("Mount %d Source", zz), support.Human(v.Source))
		info.storeData(fmt.Sprintf("Mount %d VFSOptions", zz), support.Human(v.VFSOptions))

		usage, _ := dsk.Usage(v.Mountpoint)
		//info.storeData( T,  fmt.Sprintf("Disk Usage %d", zz), T.Human( usage))
		info.storeData(fmt.Sprintf("Mount %d Total", zz), support.Human(usage.Total))
		info.storeData(fmt.Sprintf("Mount %d Free", zz), support.Human(usage.Free))
		info.storeData(fmt.Sprintf("Mount %d Used", zz), support.Human(usage.Used))
		info.storeData(fmt.Sprintf("Mount %d UsedPercent", zz), support.Human(usage.UsedPercent))
		info.storeData(fmt.Sprintf("Mount %d InodesTotal", zz), support.Human(usage.InodesTotal))
		info.storeData(fmt.Sprintf("Mount %d InodesUsed", zz), support.Human(usage.InodesUsed))
		info.storeData(fmt.Sprintf("Mount %d InodesFree", zz), support.Human(usage.InodesFree))
		info.storeData(fmt.Sprintf("Mount %d InodesUsedPercent", zz), support.Human(usage.InodesUsedPercent))

	}
	if !debugMode {
		// Open output file
		file, err := openFile(outputFilename)
		if err != nil {
			return
		}
		defer file.Close()
		err = writeStringSliceToFile(file, info.data)
		if err != nil {
			return
		}
	}
}

func (i *info) breakData(title string) {
	i.data = append(i.data, crt.Format("", ""))
	i.data = append(i.data, crt.Format(title, ""))
	i.data = append(i.data, crt.Format("", ""))
	crt.Print(crt.Bold(title))
	crt.Break()
}

func NewInfo() info {
	return info{}
}

func (i *info) storeData(title string, data string) {

	// Padd title to 15 characters
	title = fmt.Sprintf("%-30s", title)

	i.data = append(i.data, crt.Format(title+": "+data, ""))
	//fmt.Println(support.PR(title+": "+support.BOLD+data+support.RESET, T))
	crt.Print(title + ": " + crt.Bold(data))
}

func openFile(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//fmt.Printf("%s Error opening file %s: %v\n", crt.CHnormal, filename, err)
		crt.Error("Error opening file : "+crt.Bold(filename), err)
		return nil, err
	}
	return file, nil
}

func writeStringSliceToFile(file *os.File, info []string) error {
	for _, line := range info {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			//fmt.Printf("%s Error writing to file %s: %v\n", crt.CHnormal, file.Name(), err)
			crt.Error("Error writing to file : "+crt.Bold(file.Name()), err)
			return err
		}
	}
	return nil
}

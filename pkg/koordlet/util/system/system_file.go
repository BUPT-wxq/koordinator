/*
Copyright 2022 The Koordinator Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package system

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	utilsysctl "k8s.io/component-helpers/node/util/sysctl"
	"k8s.io/klog/v2"
)

const (
	ProcStatName          = "stat"
	ProcMemInfoName       = "meminfo"
	SysctlSubDir          = "sys"
	ProcCPUInfoName       = "cpuinfo"
	KernelCmdlineFileName = "cmdline"

	KernelSchedGroupIdentityEnable = "kernel/sched_group_identity_enabled"

	SysNUMASubDir = "bus/node/devices"

	SysCPUSMTActiveSubPath       = "devices/system/cpu/smt/active"
	SysIntelPStateNoTurboSubPath = "devices/system/cpu/intel_pstate/no_turbo"
)

var (
	// Jiffies is the duration unit of CPU stats. Normally, it is 10ms.
	Jiffies = float64(10 * time.Millisecond)
)

func init() {
	// $ getconf CLK_TCK > jiffies
	if err := initJiffies(); err != nil {
		klog.Warningf("failed to get Jiffies, use the default %v, err: %v", Jiffies, err)
	}
}

// initJiffies use command "getconf CLK_TCK" to fetch the clock tick on current host,
// if the command doesn't exist, uses the default value 10ms for jiffies
func initJiffies() error {
	getconf, err := exec.LookPath("getconf")
	if err != nil {
		return nil
	}
	cmd := exec.Command(getconf, "CLK_TCK")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err = cmd.Run(); err != nil {
		return err
	}
	ticks, err := strconv.ParseFloat(strings.TrimSpace(out.String()), 64)
	if err != nil {
		return err
	}
	Jiffies = float64(time.Second / time.Duration(ticks))
	return nil
}

func GetPeriodTicks(start, end time.Time) float64 {
	return float64(end.Sub(start)) / Jiffies
}

func GetProcFilePath(procRelativePath string) string {
	return filepath.Join(Conf.ProcRootDir, procRelativePath)
}

func GetProcRootDir() string {
	return Conf.ProcRootDir
}

func GetSysRootDir() string {
	return Conf.SysRootDir
}

func GetSysNUMADir() string {
	return filepath.Join(Conf.SysRootDir, SysNUMASubDir)
}

func GetNUMAMemInfoPath(numaNodeSubDir string) string {
	return filepath.Join(Conf.SysRootDir, SysNUMASubDir, numaNodeSubDir, ProcMemInfoName)
}

func GetCPUInfoPath() string {
	return filepath.Join(Conf.ProcRootDir, ProcCPUInfoName)
}

func GetSysCPUSMTActivePath() string {
	return filepath.Join(Conf.SysRootDir, SysCPUSMTActiveSubPath)
}

func GetSysIntelPStateNoTurboPath() string {
	return filepath.Join(Conf.SysRootDir, SysIntelPStateNoTurboSubPath)
}

func GetProcSysFilePath(file string) string {
	return filepath.Join(Conf.ProcRootDir, SysctlSubDir, file)
}

var _ utilsysctl.Interface = &ProcSysctl{}

// ProcSysctl implements Interface by reading and writing files under /proc/sys
type ProcSysctl struct{}

func NewProcSysctl() utilsysctl.Interface {
	return &ProcSysctl{}
}

func (*ProcSysctl) GetSysctl(sysctl string) (int, error) {
	data, err := os.ReadFile(GetProcSysFilePath(sysctl))
	if err != nil {
		return -1, err
	}
	val, err := strconv.Atoi(strings.Trim(string(data), " \n"))
	if err != nil {
		return -1, err
	}
	return val, nil
}

// SetSysctl modifies the specified sysctl flag to the new value
func (*ProcSysctl) SetSysctl(sysctl string, newVal int) error {
	return os.WriteFile(GetProcSysFilePath(sysctl), []byte(strconv.Itoa(newVal)), 0640)
}

func SetSchedGroupIdentity(enable bool) error {
	s := NewProcSysctl()
	cur, err := s.GetSysctl(KernelSchedGroupIdentityEnable)
	if err != nil {
		return fmt.Errorf("cannot get sysctl group identity, err: %v", err)
	}
	v := 0 // 0: disabled; 1: enabled
	if enable {
		v = 1
	}
	if cur == v {
		klog.V(6).Infof("SetSchedGroupIdentity skips since current sysctl config is already %v", enable)
		return nil
	}

	err = s.SetSysctl(KernelSchedGroupIdentityEnable, v)
	if err != nil {
		return fmt.Errorf("cannot set sysctl group identity, err: %v", err)
	}
	klog.V(4).Infof("SetSchedGroupIdentity set sysctl config successfully, value %v", v)
	return nil
}

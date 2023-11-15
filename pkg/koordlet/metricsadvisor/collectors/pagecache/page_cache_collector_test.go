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

package pagecache

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	gocache "github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metricsadvisor/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
	mock_statesinformer "github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer/mockstatesinformer"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/util/system"
)

func Test_collectPageCache(t *testing.T) {
	testNow := time.Now()
	testContainerID := "containerd://123abc"
	testPodMetaDir := "kubepods.slice/kubepods-podxxxxxxxx.slice"
	testPodParentDir := "/kubepods.slice/kubepods-podxxxxxxxx.slice"
	testContainerParentDir := "/kubepods.slice/kubepods-podxxxxxxxx.slice/cri-containerd-123abc.scope"
	testMemStat := `
	total_cache 104857600
	total_rss 104857600
	total_inactive_anon 104857600
	total_active_anon 0
	total_inactive_file 104857600
	total_active_file 0
	total_unevictable 0
	`
	meminfo := `MemTotal:       1048576 kB
	MemFree:          262144 kB
	MemAvailable:     524288 kB
	Buffers:               0 kB
	Cached:           262144 kB
	SwapCached:            0 kB
	Active:           524288 kB
	Inactive:         262144 kB
	Active(anon):     262144 kB
	Inactive(anon):   262144 kB
	Active(file):          0 kB
	Inactive(file):   262144 kB
	Unevictable:           0 kB
	Mlocked:               0 kB
	SwapTotal:             0 kB
	SwapFree:              0 kB
	Dirty:                 0 kB
	Writeback:             0 kB
	AnonPages:             0 kB
	Mapped:                0 kB
	Shmem:                 0 kB
	Slab:                  0 kB
	SReclaimable:          0 kB
	SUnreclaim:            0 kB
	KernelStack:           0 kB
	PageTables:            0 kB
	NFS_Unstable:          0 kB
	Bounce:                0 kB
	WritebackTmp:          0 kB
	CommitLimit:           0 kB
	Committed_AS:          0 kB
	VmallocTotal:          0 kB
	VmallocUsed:           0 kB
	VmallocChunk:          0 kB
	HardwareCorrupted:     0 kB
	AnonHugePages:         0 kB
	ShmemHugePages:        0 kB
	ShmemPmdMapped:        0 kB
	CmaTotal:              0 kB
	CmaFree:               0 kB
	HugePages_Total:       0
	HugePages_Free:        0
	HugePages_Rsvd:        0
	HugePages_Surp:        0
	Hugepagesize:          0 kB
	DirectMap4k:           0 kB
	DirectMap2M:           0 kB
	DirectMap1G:           0 kB`
	testPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pod",
			Namespace: "test",
			UID:       "xxxxxxxx",
		},
		Status: corev1.PodStatus{
			Phase: corev1.PodRunning,
			ContainerStatuses: []corev1.ContainerStatus{
				{
					Name:        "test-container",
					ContainerID: testContainerID,
					State: corev1.ContainerState{
						Running: &corev1.ContainerStateRunning{},
					},
				},
				{
					Name:        "test-failed-container",
					ContainerID: testContainerID,
					State:       corev1.ContainerState{},
				},
			},
		},
	}
	testFailedPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-failed-pod",
			Namespace: "test",
			UID:       "yyyyyy",
		},
		Status: corev1.PodStatus{
			Phase: corev1.PodFailed,
			ContainerStatuses: []corev1.ContainerStatus{
				{
					Name:        "test-container",
					ContainerID: testContainerID,
				},
			},
		},
	}

	type fields struct {
		podFilterOption       framework.PodFilter
		getPodMetas           []*statesinformer.PodMeta
		initPodLastStat       func(lastState *gocache.Cache)
		initContainerLastStat func(lastState *gocache.Cache)
		SetSysUtil            func(helper *system.FileTestUtil)
	}
	tests := []struct {
		name        string
		fields      fields
		wantEnable  bool
		wantStarted bool
	}{
		{
			name: "collect pagecache info",
			fields: fields{
				podFilterOption: framework.DefaultPodFilter,
				getPodMetas: []*statesinformer.PodMeta{
					{
						CgroupDir: testPodMetaDir,
						Pod:       testPod,
					},
				},
				initPodLastStat: func(lastState *gocache.Cache) {
					lastState.Set(string(testPod.UID), framework.CPUStat{
						CPUUsage:  0,
						Timestamp: testNow.Add(-time.Second),
					}, gocache.DefaultExpiration)
				},
				initContainerLastStat: func(lastState *gocache.Cache) {
					lastState.Set(testContainerID, framework.CPUStat{
						CPUUsage:  0,
						Timestamp: testNow.Add(-time.Second),
					}, gocache.DefaultExpiration)
				},
				SetSysUtil: func(helper *system.FileTestUtil) {
					helper.WriteProcSubFileContents(system.ProcMemInfoName, meminfo)
					helper.WriteCgroupFileContents(testPodParentDir, system.MemoryStat, testMemStat)
					helper.WriteCgroupFileContents(testContainerParentDir, system.MemoryStat, testMemStat)
				},
			},
			wantEnable:  true,
			wantStarted: true,
		},
		{
			name: "test failed pod and failed container",
			fields: fields{
				podFilterOption: framework.DefaultPodFilter,
				getPodMetas: []*statesinformer.PodMeta{
					{
						CgroupDir: testPodMetaDir,
						Pod:       testPod,
					},
					{
						Pod: testFailedPod,
					},
				},
				initPodLastStat: func(lastState *gocache.Cache) {
					lastState.Set(string(testPod.UID), framework.CPUStat{
						CPUUsage:  0,
						Timestamp: testNow.Add(-time.Second),
					}, gocache.DefaultExpiration)
				},
				initContainerLastStat: func(lastState *gocache.Cache) {
					lastState.Set(testContainerID, framework.CPUStat{
						CPUUsage:  0,
						Timestamp: testNow.Add(-time.Second),
					}, gocache.DefaultExpiration)
				},
				SetSysUtil: func(helper *system.FileTestUtil) {
					helper.WriteProcSubFileContents(system.ProcMemInfoName, meminfo)
					helper.WriteCgroupFileContents(testPodParentDir, system.MemoryStat, testMemStat)
				},
			},
			wantEnable:  true,
			wantStarted: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helper := system.NewFileTestUtil(t)
			defer helper.Cleanup()
			if tt.fields.SetSysUtil != nil {
				tt.fields.SetSysUtil(helper)
			}
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			metricCache, err := metriccache.NewMetricCache(&metriccache.Config{
				TSDBPath:              t.TempDir(),
				TSDBEnablePromMetrics: false,
			})
			assert.NoError(t, err)
			defer func() {
				metricCache.Close()
			}()
			statesInformer := mock_statesinformer.NewMockStatesInformer(ctrl)
			statesInformer.EXPECT().HasSynced().Return(true).AnyTimes()
			statesInformer.EXPECT().GetAllPods().Return(tt.fields.getPodMetas).Times(1)
			collector := New(&framework.Options{
				Config: &framework.Config{
					EnablePageCacheCollector: true,
				},
				StatesInformer: statesInformer,
				MetricCache:    metricCache,
				CgroupReader:   resourceexecutor.NewCgroupReader(),
			})
			c := collector.(*pageCacheCollector)
			assert.NotPanics(t, func() {
				c.collectPageCache()
			})
			assert.Equal(t, tt.wantEnable, c.Enabled())
			assert.Equal(t, tt.wantStarted, c.Started())
		})

	}
}

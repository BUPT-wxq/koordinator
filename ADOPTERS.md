# Adopters

This list shows adopters of Koordinator. If you're using Koordinator in some way, even only evaluating, then please feel
free to get in touch.

## Maturity Level

- 💡 Sample (for demonstrating and inspiring purpose)
- 👶 Alpha (used in companies for pilot projects)
- 👦 Beta (used in companies and developed actively)
- 👨 Stable (used in companies for production workloads)

## Adopters list

| Organization                                                        | Contact                                                                              | Maturity | Description of Use                                                                                                                                                                                                                                                                         |
|---------------------------------------------------------------------|--------------------------------------------------------------------------------------|----------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [PITS Global Data Recovery Services](https://pitsdatarecovery.net/) | [@benjx1990](https://github.com/benjx1990)                                           | Stable   | Koordinator is used to manage highly-loaded internal infrastructure providing reliable business solutions for scheduling and ML.                                                                                                                                                           |
| [XiaoHongShu](https://xiaohongshu.com)                              | [@cheimu](https://github.com/cheimu)                                                 | Stable   | Heavily use Koordinator as base building block and inspiration in production for online/offline colocation, including qos-ensurance and fine-grained CPU/GPU scheduling.                                                                                                                   |
| [iQIYI](https://www.iqiyi.com/)                                     | [@Herbert](https://github.com/wangxiaoq) [@sunwuhao](mailTo:sunwuhao001@hotmail.com) | Stable   | Use Koordinatior in production cluster for online/offline colocation to improve k8s cluster resource utilization.                                                                                                                                                                          |
| [Quwan](https://www.52tt.com)                                       | [@zhushaohua](mailTo:zhushaohua@52tt.com)                                            | Stable   | Use Koordinatior in production cluster for online/offline colocation to improve k8s cluster resource utilization.                                                                                                                                                                          |
| [360](https://www.360.com)                                          | [@liuming](https://github.com/lucming)                                               | Stable   | 1. Machines running some middleware procs they can limit the resources they use and we will use the idle resources to run some online/offline pods to increase resource utilization; 2. To improve the resource utilization in production clusters with online/offline colocation.         |
| [meiyapico](https://www.300188.cn/)                                 | [@complone](https://github.com/complone)                                             | Sample   | Koordinatior is used as the elastic scheduler of flink, hudi, and alloxio to maximize the use of resources by using resource quotas and fair scheduling                                                                                                                                    |
| [dewu](https://www.dewu.com/)                                       | [@zhouzijiang](https://github.com/zhouzijiang)                                       | Stable   | Use Koordinatior in production cluster for online/offline/flink colocation to improve k8s cluster resource utilization.                                                                                                                                                                    |
## Process

Send a PR
to [https://github.com/koordinator-sh/koordinator/blob/main/ADOPTERS.md](https://github.com/koordinator-sh/koordinator/edit/main/ADOPTERS.md)
with a brief description of how you're using Koordinator, this should include your use-case.

And please kindly reformat the above table after you added yourself to the list.

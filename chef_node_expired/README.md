# 查詢Chef client 未執行 Node List

用Chef API問Chef server ohai_time 超過6小時的Node，在用ohai_time排序。

Makefile裡有上傳lambda，Go run的task


# output

```
+---+------------------+-----------------------------------+--------+---------------------------------------------------------------------------------------+---------+----+
| # | EXPRIED CHECK IN | NODE                              | OS     | URL                                                                                   | VERSION | IP |
+---+------------------+-----------------------------------+--------+---------------------------------------------------------------------------------------+---------+----+
| 1 | a long while ago | prd-crmapi-i-0a1cdefd5245b56c7    | unknow | https://chef.senao.com.tw/organizations/senao/nodes/prd-crmapi-i-0a1cdefd5245b56c7    |         |    |
| 2 | a long while ago | prd-ecfe-spot-i-0e73afb31c9bca0e1 | unknow | https://chef.senao.com.tw/organizations/senao/nodes/prd-ecfe-spot-i-0e73afb31c9bca0e1 |         |    |
| 3 | a long while ago | prd-crmapi-i-001ddc5387779aec5    | unknow | https://chef.senao.com.tw/organizations/senao/nodes/prd-crmapi-i-001ddc5387779aec5    |         |    |
| 4 | a long while ago | prd-ccasapi-i-04024f51f39ce6b52   | unknow | https://chef.senao.com.tw/organizations/senao/nodes/prd-ccasapi-i-04024f51f39ce6b52   |         |    |
| 5 | a long while ago | prd-ccasapi-i-03861bc7efe39495d   | unknow | https://chef.senao.com.tw/organizations/senao/nodes/prd-ccasapi-i-03861bc7efe39495d   |         |    |
+---+------------------+-----------------------------------+--------+---------------------------------------------------------------------------------------+---------+----+
```

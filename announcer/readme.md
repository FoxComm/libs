## Announcer Package

#### What does it do? 
The announcer package is used to broadcast the initialization (startup) of a parent service.  This is to inform other services in the cluster that the parent service has been started.  

#### What does it depend on?
This package depends on:
 * [etcd](https://github.com/coreos/etcd): A distributed key/value store created by the CoreOS team.
 * etcd_client: Our local etcd client that wraps the GoLang [etcd client](https://github.com/coreos/go-etcd) written by the CoreOS team. 

#### Who uses it? 
Any other service that is reading etcd to look at existing endpoints.  The most relevant use cases for this is routing traffic to the parent service, which is done by the [router package](../router/router.md).

#### What are common issues I might face with this package?
Most issues you might face with this package are related to etcd:
 * etcd might not be running locally.  You can check this by doing `etcdctl ls`.  (Install etcdctl first!)  You could also simply do `ps aux | grep etcd` if you do not have or want etcdctl.
 * perhaps you have un-pruned stale endpoints in your etcd service.  You can check with `etcdctl ls /endpoints --recursive`.  You can prune all the existing endpoints with `etcdctl rm /endpoints --recursive`.  These should be replaced when services using the announcer package are restarted.





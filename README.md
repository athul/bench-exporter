# Bench Exporter
Bench Exporter is a Prometheus exporter made for Frappe Bench(es)

> **Warning**  
> Bench Exporter is in pre-alpha. There are a lot of features unwritten and to be implemented

### Data

Bench Exporter uses the Prometheus client for Go. 



The data which are exported are:  

**Bench**
- Bench Version
- Bench Site(s)
- Bench App versions
- Bench Site apps on each site  

**User**  
- All user count
- Active user count
- System Manager count

### Flags

- `--bench` : Path to bench folder. Defaults to `/home/frappe/frappe-bench`
- `--users` : To export user data(All user count, System Manager count and Active User count)

> **Note**  
>The default port of the exporter is `9101`.  
> The exporter path is `/metrics`

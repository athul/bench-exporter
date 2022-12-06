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

## Usage

Usually promethues exporters are run as a systemd service. This way you can define it once and forget about it even running. Here is a working example of configuring Bench Exporter with Systemd.

```service
# bench_exporter.service

[Unit]
Description=Bench Exporter
Wants=network-online.target
After=network-online.target

[Service]
User=root
Group=root
Type=simple
ExecStart=/usr/local/bin/bench_exporter --bench /home/frappe/frappe-bench --users

[Install]
WantedBy=multi-user.target

```

Here you have to replace the `bench-exporter` path and the path to your bench.

You can enable and start the service with,

```shell
$ systemctl enable bench_exporter.service
$ systemctl start bench_exporter.service
```

Once you've started the bench\_exporter service, You can navigate to `https://localhost:9101/metrics` and you can see the metrics.

## Grafana

Once Prometheus fetches the data, you can connect Prometheus to Grafana server which could use the promethues data to display the fetched data into a more insightful manner. 

You can create a new dashboard in Grafana to fetch the data. You can find a demo dashboard which uses Bench Exporter's data in the `dashboards` folder. You can read more about importing dashboards [here](https://grafana.com/docs/grafana/v9.3/dashboards/manage-dashboards).

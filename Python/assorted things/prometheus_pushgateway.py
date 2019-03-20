import prometheus_client
import os

registry = prometheus_client.CollectorRegistry()

duration = prometheus_client.Gauge('dani_job_duration_seconds', 'Duration of my job in seconds', registry=registry)
current_directory_filenum = prometheus_client.Gauge('dani_current_directory_filenum', 'Amount of files in the script directory', registry=registry)


with duration.time():
    print("Hello world")
    current_directory_filenum.set(10)



prometheus_client.pushadd_to_gateway('10.1.9.75:9091', job='dani_testjob', registry=registry)
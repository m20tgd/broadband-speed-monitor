To build image to work on local and Pi architecture and push to Docker Hub run:

```bash
docker buildx build --platform linux/amd64,linux/arm64 -t m20tgd/broadband-speed-monitor --push .
```

And then run this on the Pi to pull it:

```bash
docker pull m20tgd/broadband-speed-monitor

```

Create .env file with variable `INFLUXDB_TOKEN`



Add below entry to CRONTAB:



```bash
  * *  *   *   *     nohup docker run --rm --env-file /path/to/.env m20tgd/broadband-speed-monitor >> /path/to/log/broadband-speed.txt 2>&1 < /dev/null &
```

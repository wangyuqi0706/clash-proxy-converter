# Converter

Converter your share link begin with `ss://` or `vmess://` to a clash config file.

# How To Use
Download and unarchive the tool

Write your links into a text file, one link per line.
```
ss://your_link
vmess://another_link
```
Run the tool in terminal
```
./converter -i links.txt
```

You will find a `config.yaml` in the same directory.It only contains `proxies` section.

You need to add other sections. See [Clash Document](https://lancellc.gitbook.io/clash/clash-config-file/an-example-configuration-file)
# Popmedia2

## A media server in go-lang

### To test:
```
./test.sh
```

### To run:
```
./run.sh
```

### To build:
```
./build.sh
```

### To deploy (all platforms):
```
./deploy.sh
```
This will create zip files under artifact for all platforms, select your platform and move the zip to the server.  Unzip the file and run the install.sh on the server.

### To install:
```
sudo ./install.sh
```

### To start:
```
sudo service popmedia-server-service start
```

### To stop:
```
sudo service popmedia-server-service stop
```

### To restart:
```
sudo service popmedia-server-service restart
```

### To configure:
Edit file "config.json" with the desired settings.
- Port = port to run on.
- Root = root location of site.
- *MediaExt = currently only support MP4*
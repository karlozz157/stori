# stori

### download the proyect
```bash
git clone git@github.com:karlozz157/stori.git
```

### build 
```bash
cd stori && docker build -t stori .
```

### run
```bash
docker run -p 6969:6969 stori
```

### create the summary
```bash
curl -X POST http://localhost:6969/summary
```

### check the summary
```bash
curl -X GET http://localhost:6969/summary
```

### send by email
```bash
curl -X POST http://localhost:6969/summary/send
```

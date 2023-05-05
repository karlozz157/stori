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

### you can find the csv
```bash
cat csv/example.csv
```

### check the summary 
![image](https://user-images.githubusercontent.com/4811721/236516409-060e7b69-a20c-4ad5-91a9-a53931c32e9f.png)

### check the email
![image](https://user-images.githubusercontent.com/4811721/236516638-bec29a4b-4a48-4f34-8581-612db2d5265a.png)

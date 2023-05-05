# stori

### download the proyect
git clone git@github.com:karlozz157/stori.git

### build 
cd stori && docker build -t stori .

### run 
docker run -p 6969:6969 stori

### create the summary
curl -X POST http://localhost:6969/summary

### check the summary
curl -X GET http://localhost:6969/summary

### send by email
curl -X POST http://localhost:6969/summary/send

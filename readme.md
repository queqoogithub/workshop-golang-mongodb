# **Workshop Golang & MongoDB (2 Hrs)**
&nbsp;<br />
&nbsp;
## Install Requirement
&nbsp;
1. Golang v1.19 หรือสูงกว่า
   - windows https://go.dev/dl/go1.19.5.windows-amd64.msi
   - Apple macOS (ARM64) https://go.dev/dl/go1.19.5.darwin-arm64.pkg
2. Mongodb https://www.mongodb.com/try/download/community
3. Mongodb Compass เอาไว้ต่อ Mongodb https://www.mongodb.com/try/download/compass
4. Visual Studio Code https://code.visualstudio.com/download
5. Docker Desktop (optional) https://www.docker.com/products/docker-desktop/
&nbsp;<br />
&nbsp;
## **============= Agenda =============**
&nbsp;<br />
### 1. Introduction Golang (10 min) 
### 2. Go Technical (30 min) https://www.babelcoder.com/blog/articles/intro-to-golang
 - Variable https://www.golangprograms.com/go-language/variables.html
 - Struct https://www.golangprograms.com/go-language/struct.html
 - Pointer https://dev.to/iporsut/go-pointer-pointer-go-3212
 - Go http
 - Go routine (Asynchronous)
### 3. Break (10 min)
### 4. Go Workshop from sourcecode (50 min)
 - Rest API framework from Gin-Gonic https://github.com/kittisak-kueb/workshop-golang-mongodb-src/tree/main/1.gin-restful
 - CRUD Data with MongoDB https://github.com/kittisak-kueb/workshop-golang-mongodb-src/tree/main/2.crud
 - Go middleware https://github.com/kittisak-kueb/workshop-golang-mongodb-src/tree/main/3.middleware
 - MongoDB usage https://github.com/kittisak-kueb/workshop-golang-mongodb-src/tree/main/4.mongodb
### 5. Go and Docker (10 min)
### 6. End Session / Q&A / HomeWork (10 min)
&nbsp;<br />
&nbsp;
# **Homework**
  - เพิ่มเส้น Company Profile โดยทำแบบ CRUD
  - กำหนดเส้น middleware ส่วน insert/update ให้ทำการเช็ค http header X-Api-Key โดยค่าจะต้องเท่ากับ my-gosoft-password ถ้าไม่ใช่ค่าตัวนี้ ให้ส่งข้อความมาบอกว่า invalid API Key พร้อมทั้ง http status 401
  - กำหนดให้ค่าที่เก็บ company profile ดังนี้
  <br />
  {companyId: "109",companyName: "Gosoft",employeeAmt: 100,companyStatus: 1}
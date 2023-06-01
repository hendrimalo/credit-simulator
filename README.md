# CREDIT-SIMULATOR-WEB-SERVICE

## INSTALL APLIKASI DARI SOURCE CODE
- **download package**: jalankan "go mod tidy" pada terminal lokasi project
- **jalankan program**: jalankan "go run .\main.go"

## CARA PENGGUNAAN APLIKASI
Aplikasi ini berbasis web service untuk bisa menjalankannya diperlukan aplikasi client seperti browser / postman. Dalam case ini saya akan menggunakan postman sebagai contoh. Export postman sudah tersedia di root project dengan nama "credit-simulator.postman_collection"
- **export postman collection**: akan terdapat collection baru dengan nama "Credit Simulator"
- **CREATE | Credit Simulator With Text**: untuk membaca inputan dari file txt dari url param "***v1/credit-simulator/$nama_file" #note: pastikan nama file sesuai dan tersedia
- **CREATE | Credit Simulator With Request**: untuk membaca inputan dari request body json

## MEMBACA INPUT DARI TXT FILE
- **format**: terdiri dari 6 variabel yang dipisahkan dengan tanda ","
- vehicle_type: **string**
- vehicle_condition: **string**
- year: **int**
- down_payment: **float32**
- total: **float32**
- tenor: **int**

- **contoh**: contoh format yang benar
CAR,NEW,2023,60000000,150000000,6

- **lokasi** lokasi default path
./files/input_credit/$nama_file

## MENJALANKAN UNIT TESTING
- **credit_simulator**: unit testing credit simulator
go test .\services\credit_simulator\app.go .\services\credit_simulator\interface.go .\services\credit_simulator\app_test.go -v

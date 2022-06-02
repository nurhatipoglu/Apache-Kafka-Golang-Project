# Apache-Kafka-Golang-Project

### 170419011 Nur Hatipoğlu Açık Kaynak Kodlu Yazılımlar Dönem Proje Ödevi

![alt text](https://www.loginradius.com/blog/static/30bb9b9901e76f3498d68913a0675ea9/03979/index.png 'Logo Title Text 1')

### Apache Kafka Uygulaması Geliştirme Projesi

Go programlama dili kullanılarak Apache Kafka’ya "myTopic" adı altına bir JSON mesajını üreten (procuder) ve bu "myTopic" deki mesajları tüketen (consumer) uygulaması geliştirilmiştir.
Producer bir json mesajı üretir ve post metodu ile çalışır. Consumer okuduğu mesajı consola bastırır.
Consumer ve Producer kodları yazıldıktan sonra ikisi için de Dockerfile oluşturulur.

---

Git üzerinden git clone yaparak proje çalışılacak olan klasöre klonlanır.
Kodlarımız indikten sonra ilk yapmamız gereken Consumer ve Producer a ait Dockerfile ları build etmemiz gerekir. Daha sonra docker compose dosyası ayaklandırılır. Post isteği atılır ve consumer mesajı konsola yazdırır.
Aşağıda detaylı açıklanmıştır.

### Consumer a ait Image Build Etmek

consumer dizininin içine girilir. Sırasıyla aşağıdaki komutlar çalıştırılır;

```javascript
go build .
docker build -f Dockerfile -t hatipoglu-consumer .
```

### Producer a ait Image Build Etmek

rest-api dizininin içine girilir. Sırasıyla aşağıdaki komutlar çalıştırılır;

```javascript
go build .
docker build -f Dockerfile -t hatipoglu-rest-api .
```

### Docker Compose Dosyasını Çalıştırmak

docker-compose.yml dosyasının dizinine yani projenin ana dizine çıkılır. Aşağıdaki komut çalıştırılır;

```javascript
docker-compose up -d
```

---

Postman den aşağıdaki yapıda; http://localhost:3003/ adresine post isteği atıldığında consumer ın konsola atılan mesajı yazdırdığını görürüz.

```javascript
{
"message" : "How Are You"
}
```

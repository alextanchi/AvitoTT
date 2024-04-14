Что не сделано: 
1. "Реализуйте интеграционный или E2E-тест на сценарий получения баннера".
2. Не создан makefile \

Для запуска докера docker-compose up\
docker поднимает базу\
Запуск: go run main.go

Токен для юзера: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwidXNlcklkIjoiM2IxYjE5ZjMtNDU1ZC00NzRkLTkyZmMtNjVhNzY1NTFiMTZmIiwiaWF0IjoxNTE2MjM5MDIyfQ.fW4mz71VDG3_EJAbdprRRslUeHzby4GZmVoceqI0zBM \
Токен для админа: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwidXNlcklkIjoiMGQ5N2UyM2ItMDg5Yi00MjM3LTljMmMtZDFlMTAwNTc2OTIwIiwiaWF0IjoxNTE2MjM5MDIyfQ.uxPICTLdCXY2876Oz7UG7szFXB0NMwsKqEJnCNMxoA4 

**GET /user_banner** 
```
curl --location 'http://localhost:8080/user_banner?tag_id=2&feature_id=1&use_last_revision=true' 
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwidXNlcklkIjoiM2IxYjE5ZjMtNDU1ZC00NzRkLTkyZmMtNjVhNzY1NTFiMTZmIiwiaWF0IjoxNTE2MjM5MDIyfQ.fW4mz71VDG3_EJAbdprRRslUeHzby4GZmVoceqI0zBM'
 ```

**GET /banner** 
```
curl --location 'http://localhost:8080/banner?feature_id=1&tag_id=2&limit=10&offset=0' 
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwidXNlcklkIjoiMGQ5N2UyM2ItMDg5Yi00MjM3LTljMmMtZDFlMTAwNTc2OTIwIiwiaWF0IjoxNTE2MjM5MDIyfQ.uxPICTLdCXY2876Oz7UG7szFXB0NMwsKqEJnCNMxoA4' 
```

**POST /banner** 
```curl --location 'http://localhost:8080/banner' 
--header 'Content-Type: application/json' 
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwidXNlcklkIjoiMGQ5N2UyM2ItMDg5Yi00MjM3LTljMmMtZDFlMTAwNTc2OTIwIiwiaWF0IjoxNTE2MjM5MDIyfQ.uxPICTLdCXY2876Oz7UG7szFXB0NMwsKqEJnCNMxoA4' 
--data '{
  "tag_ids": [
    1
  ],
  "feature_id": 1,
  "content": {
    "title": "Баннер3",
    "text": "Текст3",
    "url": "Url3"
  },
  "is_active": true
}
'
```

**PATCH /banner/{id}** 
```
curl --location --request PATCH 'http://localhost:8080/banner/3' 
--header 'Content-Type: application/json' 
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwidXNlcklkIjoiMGQ5N2UyM2ItMDg5Yi00MjM3LTljMmMtZDFlMTAwNTc2OTIwIiwiaWF0IjoxNTE2MjM5MDIyfQ.uxPICTLdCXY2876Oz7UG7szFXB0NMwsKqEJnCNMxoA4' 
--data '{
  "tag_ids": [
    101
  ],
  "feature_id": 3,
  "content": {
    "title": "БаннерОбновленный",
    "text": "ТекстОбновленный",
    "url": "some_url"
  },
  "is_active": true
}'
```

**DELETE /banner/{id}** 
```curl --location --request DELETE 'http://localhost:8080/banner/1' 
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwidXNlcklkIjoiMGQ5N2UyM2ItMDg5Yi00MjM3LTljMmMtZDFlMTAwNTc2OTIwIiwiaWF0IjoxNTE2MjM5MDIyfQ.uxPICTLdCXY2876Oz7UG7szFXB0NMwsKqEJnCNMxoA4'
```




Что можно улучшить:
1. Добавить логи
2. Добавить транзакции
3. GetBannersByFeatureAndTag возвращает только один тег у каждого баннера
4. Добавить уровни логгирования
5. Методы работают только с существующими фичами





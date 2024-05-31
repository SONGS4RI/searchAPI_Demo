# searchAPI_Demo

gin gonic과 elasticsearch를 활용하여 검색API를 구현해 보는 프로젝트입니다.

go 1.20 버전을 사용하였습니다.

```json
// kibana를 사용해서 elasticsearch에서 스냅샷 복구
// 스냅샷 등록
PUT _snapshot/search_example
{
  "type": "fs",
  "settings": {
    "location": "/es/book_backup/search_example",
    "compress": true
  }
}

// 스냅샷 등록 확인
GET _snapshot/search_example/_all

// 스냅샷 복구
POST _snapshot/search_example/movie-search/_restore
```

```
```
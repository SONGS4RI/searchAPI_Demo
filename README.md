# searchAPI_Demo

gin gonic과 elasticsearch를 활용하여 검색API를 구현해 보는 프로젝트입니다.

go 1.20 버전을 사용하였습니다.

## 이용 방법
1. go 1.20 버전 설치 -> go version 으로 설치 확인
2. elasticsearch와 kibana 컨테이너 생성
    ```BASH
    docker-compose up -d
    ```
3. 키바나를 사용해서 예제 인덱스 스냅샷 복구 -> http://localhost:5601 에서 메뉴바 Management/Dev Tools 에서 콘솔 사용
    ```BASH
    # elasticsearch 노드에 스냅샷 등록
    PUT _snapshot/search_example
    {
      "type": "fs",
      "settings": {
        "location": "/es/book_backup/search_example",
        "compress": true
      }
    }
    # 스냅샷 등록 확인
    GET _snapshot/search_example/_all

    # 스냅샷 복구
    POST _snapshot/search_example/movie-search/_restore
    ```
4. 프로젝트 루트에서
    ```BASH
    go run cmd/main.go
    ```
5. http://localhost:8080 에서 스웨거 사용 가능
package elasticconn

import (
	"log"
	"time"

	"github.com/olivere/elastic/v7"
)

func InitES() {
	_, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),        // elasticsearch 서버 설정 & ','으로 다수 등록 가능
		elastic.SetSniff(false),                        // 클러스터 sniffing 비활성화
		elastic.SetHealthcheckInterval(10*time.Second), // 클러스터 상태 확인 간격 설정
		elastic.SetRetrier(elastic.NewBackoffRetrier(elastic.NewExponentialBackoff(100*time.Millisecond, 10*time.Second)))) // 재시도 전략

	if err != nil {
		log.Fatalln("ELASTIC CLIENT 생성 ERROR:", err)
	}
	log.Println("ELASTIC CLIENT 생성")
}

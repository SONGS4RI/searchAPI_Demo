package elasticconn

import (
	"context"
	"log"
	"time"

	"github.com/olivere/elastic/v7"
)

var (
	EsClient *elastic.Client
)

func InitEs() {
	EsClient, _ = elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),        // elasticsearch 서버 설정 & ','으로 다수 등록 가능
		elastic.SetSniff(false),                        // 클러스터 sniffing 비활성화
		elastic.SetHealthcheckInterval(10*time.Second), // 클러스터 상태 확인 간격 설정
		elastic.SetRetrier(elastic.NewBackoffRetrier(elastic.NewExponentialBackoff(100*time.Millisecond, 10*time.Second)))) // 재시도 전략

	if _, err := EsClient.CatHealth().Do(context.TODO()); err != nil {
		log.Println("ELASTIC CLIENT  연결 실패", err)
	} else {
		log.Println("ELASTIC CLIENT  연결 성공!!!")
	}
}

func GetEsClinet() (client *elastic.Client) {
	return EsClient
}

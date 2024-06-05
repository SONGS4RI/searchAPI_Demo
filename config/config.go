package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

const (
	ELASTIC_SEED_JSON_FILE = "elastic_seed.json"
)

var (
	ElasticSeedMap = make(map[string]map[string]string)
)

// 프로세스 시작시 파일에서 엘라스틱서치의 시드정보 로딩
func LoadElasticSeed() {
	elasticSeedJsonPath := path.Join(getConfigPath(), ELASTIC_SEED_JSON_FILE)
	if _, sErr := os.Stat(elasticSeedJsonPath); sErr == nil {
		if target, readErr := os.ReadFile(elasticSeedJsonPath); readErr == nil {
			// 시드 파일에서 시드정보를 변수에 적재한다.
			_ = json.Unmarshal(target, &ElasticSeedMap)
			fmt.Println("시드파일 조회 완료.")
			fmt.Println(ElasticSeedMap)
		}
	}
}

// 설정 파일 경로
func getConfigPath() (filePath string) {
	if _, err := os.Stat(filePath); err != nil {
		// 디렉토리가 존재하지 않으면 생성한다.
		_ = os.MkdirAll(filePath, 0755)
	}
	return
}

func GetLocalElasticSeed() (url string, username string, password string) {
	url = ElasticSeedMap["LOCAL_ELASTIC_SEED"]["URL"]
	username = ElasticSeedMap["LOCAL_ELASTIC_SEED"]["USERNAME"]
	password = ElasticSeedMap["LOCAL_ELASTIC_SEED"]["PASSWORD"]
	return
}
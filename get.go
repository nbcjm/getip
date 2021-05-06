package func1

import (
	"encoding/json"
	"net/http"
	"time"
)

type Info struct {
	IP string `json:"ip"`
	// https://stackoverflow.com/questions/28228393/json-unmarshal-returning-blank-structure
	// 구조체 필드가 소문자로 시작하면 export 되지 않는다.
	Country string `json:"country"`
	Country_code string `json:"cc"`
}

func GetInfo(url string) (Info) {

	//fmt.Println(reflect.TypeOf(Info{}).NumField()) // this return 3
	client := http.Client{Timeout: 10 * time.Second} //Timeout 10 초의 속성을 가진 Client 객체 생성
	resp, err := client.Get(url) //http.get 메소드 호출
	if err != nil {
		panic(err) // 로그 메세지를 작성 한 후 고 루틴 종료
	}
	defer resp.Body.Close()

	Information := Info{}
	json.NewDecoder(resp.Body).Decode(&Information)

	return Information
}


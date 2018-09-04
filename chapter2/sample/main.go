// 패키지
// 지금 당장은 패키지가 컴파일된 코드를 구별하기 위한 단위를 정의하고
// 여기에 이름을 부여하여 그안에 정의된 식별자들을 구별하기 위한 일종의 네임스페이스 같은것이라고 알아두자
package main

// 외부 코드를 가져오는 코드
// 같은 폴더에 저장된 모든 소스 파일은 같은 패키지 이름을 사용해야 하며,
// 폴더 이름과 동일한 패키지 이름을 사용하는 것이 관례다.
import (
    "fmt"
    _ "github.com/code-j/go-in-action/chapter2/sample/matchers" // _ 빈 식별자 사용
    "github.com/code-j/go-in-action/chapter2/sample/search"
    "log"
    "os"
)

// init 함수는 main 함수보다 먼저 호출된다.
func init() {
    // 표준 출력으로 로그를 출력하도록 변경한다.
    // 2018/09/04 21~ 과 같이 로그 등록이 설정된다.
    log.SetOutput(os.Stdout)
    fmt.Println("##### 표준 출력 로그 설정 #####")
}

// main 함수는 프로그램의 진입점이다.
// main 함수가 main 패키지에 선언되어 있지 않으면 빌드 도구는 실행파일을 만들어내지 못한다.
func main() {
    // 지정된 검색어로 검색을 수행한다. (즉, Sherlock Holmes로 검색하겠다는거지)
    search.Run("Sherlock Holmes")
}

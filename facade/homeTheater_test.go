package facade

import "testing"

/*
퍼사드 패턴은 서브시스템에 있는 일련의 인터페이스를 통합 인터페이스로 묶어 줍니다
또한 고수준 인터페이스도 정의하므로 서브시스템을 더 편리하게 사용할 수 있습니다.
*/
func TestHomeTheaterDrive(t *testing.T) {
	homeTheaterFacade := new(HomeTheaterFacade)
	// ..homeTheaterFacade.Constructor(...)

	homeTheaterFacade.WatchMovie("인디아나 존스:레이더스")
	homeTheaterFacade.EndMovie()
}

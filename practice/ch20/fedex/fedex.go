package fedex

import "fmt"

type FedexSender struct{} // 타입선언시 인터페이스 구현 여부를 명시하지 않아도 된다.

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex sends %s parcel\n", parcel)
}

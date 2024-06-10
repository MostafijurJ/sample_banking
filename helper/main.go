package main

import "sample_banking/helper/serializer"

func main() {
	serializer.TestJson()
	serializer.TestGob()
	serializer.TestProtoBuff()
}

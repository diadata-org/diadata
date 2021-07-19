package datasource

import "testing"

func TestInitExchange(t *testing.T) {
	 s,_ := InitSource()
	exchanges :=  s.GetExchanges()
	got := exchanges["0x"].Contract.String()
	if got != "0x61935CbDd02287B511119DDb11Aeb42F1593b7Ef"{
		t.Errorf("exchanges[\"0x\"].Contract = %s; want 0x61935cbdd02287b511119ddb11aeb42f1593b7ef", got)
	}
}
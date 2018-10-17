package diaApi

import(
	"encoding/json"
)

func (e *SymbolDetails) UnmarshalBinary(data []byte) error {
  if err := json.Unmarshal(data, &e); err != nil {
    return err
  }
  return nil
}

func (e *Coins) UnmarshalBinary(data []byte) error {
  if err := json.Unmarshal(data, &e); err != nil {
    return err
  }
  return nil
}

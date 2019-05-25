package types

import (
	"gxclient-go/core/transaction"
	"github.com/pquerna/ffjson/ffjson"
	"encoding/json"
)

type AccountNameEqLitPredicate struct {
	AccountId   GrapheneID `json:"account_id"`
	Name    	string `json:"name"`
}

func (p AccountNameEqLitPredicate) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(PredicateAccountNameEqLit))
	enc.Encode(p.AccountId)
	enc.Encode(p.Name)
	return nil
}

type AssetSymbolEqLitPredicate struct {
	AssetId   GrapheneID `json:"asset_id"`
	Symbol    string `json:"symbol"`
}

func (p AssetSymbolEqLitPredicate) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(PredicateAssetSymbolEqLit))
	enc.Encode(p.AssetId)
	enc.Encode(p.Symbol)
	return nil
}

type BlockIdPredicate struct {
	ID   []byte `json:"id"`
}

func (p BlockIdPredicate) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(PredicateBlockId))
	size := len(p.ID)
	enc.EncodeUVarint(uint64(size))
	enc.Encode(p.ID)
	return nil
}

type Predicates []Predicate

func (p Predicates) MarshalTransaction(enc *transaction.Encoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return err
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return err
		}
	}

	return nil
}

type Predicate struct {
	Type        PredicateType
	Pre			interface{}
}

func (p Predicate) MarshalTransaction(enc *transaction.Encoder) error {
	switch p.Type {
	case PredicateAccountNameEqLit:
		if err := enc.Encode(p.Pre.(*AccountNameEqLitPredicate)); err != nil {
			return err
		}
	case PredicateAssetSymbolEqLit:
		if err := enc.Encode(p.Pre.(*AssetSymbolEqLitPredicate)); err != nil {
			return err
		}
	case PredicateBlockId:
		if err := enc.Encode(p.Pre.(*BlockIdPredicate)); err != nil {
			return err
		}
	}

	return nil
}

func (p Predicate) MarshalJSON() ([]byte, error) {
	return ffjson.Marshal([]interface{}{
		p.Type,
		p.Pre,
	})
}

func (p *Predicate) UnmarshalJSON(data []byte) error {
	raw := make([]json.RawMessage, 2)
	if err := ffjson.Unmarshal(data, &raw); err != nil {
		return err
	}

	if len(raw) != 2 {
		return ErrInvalidInputLength
	}

	if err := ffjson.Unmarshal(raw[0], &p.Type); err != nil {
		return err
	}

	switch p.Type {
	case PredicateAccountNameEqLit:
		p.Pre = &AccountNameEqLitPredicate{}
	case PredicateAssetSymbolEqLit:
		p.Pre = &AssetSymbolEqLitPredicate{}
	case PredicateBlockId:
		p.Pre = &BlockIdPredicate{}
	}

	if err := ffjson.Unmarshal(raw[1], p.Pre); err != nil {
		return err
	}

	return nil
}
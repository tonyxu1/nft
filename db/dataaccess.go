package db

import (
	"log"
	"nft/graph/model"
)

func GetTokenInfo() (*Token, error) {

	token := &Token{ID: 1}
	err := db.Model(token).WherePK().Select()
	if err != nil {
		return nil, err
	}
	return token, nil
}

// PutTokenInfo create records info in the database for the models, and update last read block in tokens table
func PushTransactionInfo(
	gqlTransacction []*model.Transaction,
	gqlTrxInternal []*model.TrxInternal,
	gqlTrxErc20 []*model.TrxErc20,
	gqlTrxErc721 []*model.TrxErc721,
	gqlTrxErc1155 []*model.TrxErc1155,
	currentBlock uint64,
) error {

	if len(gqlTransacction) > 0 {

		_, err := db.Model(&gqlTransacction).OnConflict("DO NOTHING").Insert()
		if err != nil {
			log.Println("Error:", err)
		}
	}

	if len(gqlTrxInternal) > 0 {

		_, err := db.Model(&gqlTrxInternal).OnConflict("DO NOTHING").Insert()
		if err != nil {
			log.Println("Error:", err)
		}
	}

	if len(gqlTrxErc20) > 0 {

		_, err := db.Model(&gqlTrxErc20).OnConflict("DO NOTHING").Insert()
		if err != nil {
			log.Println("Error:", err)
		}

	}

	if len(gqlTrxErc721) > 0 {

		_, err := db.Model(&gqlTrxErc721).OnConflict("DO NOTHING").Insert()
		if err != nil {
			log.Println("Error:", err)
		}
	}

	if len(gqlTrxErc1155) > 0 {

		_, err := db.Model(&gqlTrxErc1155).OnConflict("DO NOTHING").Insert()
		if err != nil {
			log.Println("Error:", err)
		}
	}

	token := &Token{ID: 1}
	token.LastReadBlock = currentBlock
	_, err := db.Model(token).Column("last_read_block").WherePK().Update()
	if err != nil {
		log.Println("Error:", err)
	}

	return nil
}

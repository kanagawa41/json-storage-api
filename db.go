package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// DBのレコードの構造体
type Stock struct {
	ID   int64
	UUID string
	JSON string
}

// func main() {
// 	// テーブル作成
// 	_, err = db.Exec(
// 		`CREATE TABLE IF NOT EXISTS "STOCKS" ("ID" INTEGER PRIMARY KEY, "UUID" VARCHAR(255), "JSON" BLOB)`,
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	las := lastID(db)
// 	cID := createStock(db, las+1)

// 	s := selectStock(db, cID)
// 	fmt.Println(s)

// 	updateStock(db, cID, "{\"test\":\"value\"}")

// 	s = selectStock(db, cID)
// 	fmt.Println(s)
// }

func connection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./stock.db")

	if err != nil {
		return db, err
	}

	// テーブル作成
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS "STOCKS" ("ID" INTEGER PRIMARY KEY, "UUID" VARCHAR(255), "JSON" BLOB)`,
	)
	if err != nil {
		return db, err
	}

	return db, nil
}

// func testID() {
//   // Output will be stored here.
//   var output string

//   // Prepare your query
//   query, err := DB.Prepare("SELECT COUNT(*) as count FROM  STOCKS")

//   if err != nil {
//       fmt.Printf("%s", err)
//   }

//   defer query.Close()

//   // Execute query using 'id' and place value into 'output'
//   err = query.QueryRow().Scan(&output)

//   // Catch errors
//   switch {
//   case err == sql.ErrNoRows:
//           fmt.Printf("No notebook with that ID.")
//   case err != nil:
//           fmt.Printf("%s", err)
//   default:
//           fmt.Printf("Counted %s notebooks\n", output)
//   }
// }

func lastID(db *sql.DB) (lastID int) {
	// 1件取得
	row := db.QueryRow(`SELECT COUNT(*) as count FROM STOCKS`)
	// row := DB.QueryRow(`SELECT COUNT(*)`)

	err := row.Scan(&lastID)

	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("Not found")
	case err != nil:
		panic(err)
	}

	return lastID
}

func selectStock(db *sql.DB, uuid string) (s Stock, err error) {
	// 1件取得
	row := db.QueryRow(
		`SELECT * FROM STOCKS WHERE UUID=?`,
		uuid,
	)

	err = row.Scan(&s.ID, &s.UUID, &s.JSON)

	switch {
	case err == sql.ErrNoRows:
		return s, err
	case err != nil:
		panic(err)
	}

	return s, nil
}
func selectStockWithId(db *sql.DB, id int64) (s Stock, err error) {
	// 1件取得
	row := db.QueryRow(
		`SELECT * FROM STOCKS WHERE ID=?`,
		id,
	)

	err = row.Scan(&s.ID, &s.UUID, &s.JSON)

	switch {
	case err == sql.ErrNoRows:
		return s, err
	case err != nil:
		panic(err)
	}

	return s, nil
}

func createStock(db *sql.DB, id int) (Stock, error) {
	// データの挿入
	res, err := db.Exec(
		`INSERT INTO STOCKS (ID, UUID, JSON) VALUES (?, lower(hex(randomblob(16))), ?)`,
		id,
		"{}",
	)
	if err != nil {
		panic(err)
	}

	// 挿入処理の結果からIDを取得
	cID, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return selectStockWithId(db, cID)
}

func updateStock(db *sql.DB, uuid string, json string) (int64, error) {
	res, err := db.Exec(
		`UPDATE STOCKS SET JSON=? WHERE UUID=?`,
		json,
		uuid,
	)
	if err != nil {
		return -1, err
	}

	// 更新されたレコード数
	affect, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}

	return affect, nil
}

func deleteStock(db *sql.DB, uuid string) (int64, error) {
	res, err := db.Exec(
		`DELETE STOCKS WHERE UUID=?`,
		uuid,
	)
	if err != nil {
		return -1, err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}

	return affect, nil
}

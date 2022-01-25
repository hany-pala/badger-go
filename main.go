package main

import (
	"github.com/dgraph-io/badger"
	"log"
)

func main() {
	db, e := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if e != nil {
		log.Fatal(e)
	}
	//fmt.Println(db)
	/** Read-only transactions */
	//err := db.View(func(txn *badger.Txn) error {
	//	return nil
	//})

	/** Read-write transactions */
	//err := db.Update(func(txn *badger.Txn) error {
	//	return nil
	//})

	// [update]
	//updates := make(map[string]string)
	//txn := db.NewTransaction(true)
	//for k, v := range updates {
	//	if err := txn.Set([]byte(k), []byte(v)); err == badger.ErrTxnTooBig {
	//		_ = txn.Commit()
	//		txn = db.NewTransaction(true)
	//		_ = txn.Set([]byte(k), []byte(v))
	//	}
	//}
	//_ = txn.Commit()

	// [write]
	//txn := db.NewTransaction(true)
	//defer txn.Discard()
	//
	//// Use the transaction...
	//e := txn.Set([]byte("answer"), []byte("4545"))
	//
	//if e != nil {
	//	fmt.Println(e)
	//}
	//
	//if err := txn.Commit(); err != nil {
	//	fmt.Println(err)
	//}

	//err := db.Update(func(txn *badger.Txn) error {
	//	err := txn.Set([]byte("answer"), []byte("42"))
	//	return err
	//})
	//fmt.Println(err)

	//err := db.Update(func(txn *badger.Txn) error {
	//	e := badger.NewEntry([]byte("answer"), []byte("42"))
	//	err := txn.SetEntry(e)
	//	return err
	//})
	//fmt.Println(err)

	//err := db.View(func(txn *badger.Txn) error {
	//	item, err := txn.Get([]byte("answer"))
	//	fmt.Println(err)
	//	fmt.Println(item) // key="answer", version=13, meta=40
	//
	//	var valNot, valCopy []byte
	//	/** value copy case 1 */
	//	//e := item.Value(func(val []byte) error {
	//	//	fmt.Printf("The answer is: %s\n", val) //42로 가지고 온다.
	//	//	valCopy = append([]byte{}, val...)
	//	//	fmt.Printf("value: %p\n, %v\n, %d\n", valCopy, valCopy, val) // 0xc00001c418
	//	//	valNot = val
	//	//	return nil
	//	//})
	//	fmt.Println(e)
	//	fmt.Printf("NEVER do this. %s\n", valNot)
	//	fmt.Printf("The answer is: %s\n", valCopy)
	//
	//
	//	/** value copy case 2 */
	//	valCopy, err = item.ValueCopy(nil)
	//	fmt.Printf("The answer is: %s\n", valCopy)
	//	return nil
	//})
	//
	//fmt.Println(err)

	/** Monotonically increasing integers */
	//seq, err := db.GetSequence([]byte("answer"), 1000)
	//defer seq.Release()
	//for {
	//	num, err := seq.Next()
	//	fmt.Printf("number is %d\n", num)
	//	fmt.Println("err: '", err)
	//}
	//fmt.Println(err)

	defer db.Close()
}

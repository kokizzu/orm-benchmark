package benchs

import (
	"fmt"
	"github.com/gocraft/dbr"
)

var con *dbr.Connection

func init() {
	st := NewSuite("dbr")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, DbrInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, DbrInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, DbrUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, DbrRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, DbrReadSlice)

		var err error
		con, err = dbr.Open("mysql", ORM_SOURCE, nil)
		if err != nil {
			panic(err)
		}
		con.SetMaxIdleConns(ORM_MAX_IDLE)
		con.SetMaxOpenConns(ORM_MAX_CONN)
	}
}

func DbrInsert(b *B) {
	sess := con.NewSession(nil)
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if _, err := sess.InsertInto(`model`).Columns(`name`,`title`,`fax`,`web`,`age`,`right`,`counter`).Record(m).Exec(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func dbrInsert(m *Model) error {
	sess := con.NewSession(nil)
	res, err := sess.InsertInto(`model`).Columns(`name`,`title`,`fax`,`web`,`age`,`right`,`counter`).Record(m).Exec()
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	m.Id = int(id)
	return err
}

func DbrInsertMulti(b *B) {
	panic(fmt.Errorf("Not support multi insert"))
}

func DbrUpdate(b *B) {
	sess := con.NewSession(nil)
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		dbrInsert(m)
	})

	for i := 0; i < b.N; i++ {
		_, err := sess.Update(`model`).Set(`name`,m.Name).Set(`title`,m.Title).Set(`fax`,m.Fax).Set(`web`,m.Web).Set(`age`,m.Age).Set(`right`,m.Right).Set(`counter`,m.Counter).Exec()
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func DbrRead(b *B) {
	sess := con.NewSession(nil)
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		dbrInsert(m)
	})

	for i := 0; i < b.N; i++ {
		_, err := sess.Select("*").From(`model`).Where(`id = ?`,m.Id).Limit(1).Load(m)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func DbrReadSlice(b *B) {
	sess := con.NewSession(nil)
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			dbrInsert(m)
		}
	})
	for i := 0; i < b.N; i++ {
		res := map[string]interface{}{}
		_, err := sess.Select("*").From(`model`).Where(`id > ?`, 0).Limit(100).Load(&res)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

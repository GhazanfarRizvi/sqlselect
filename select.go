package sqlselect

import(
	"database/sql"
	"fmt"
) 

func Select(db *sql.DB, sel string) (selectRows *sql.Rows) {
	selectRows, err := db.Query(sel)
	if err != nil {
		fmt.Println("Error running ODM_FWA_SUBS_TRAFFIC_4GAVG select query")
		fmt.Println(err)
	}
	return selectRows
}

func Odmfwasubsmaster1(db *sql.DB, sel string) (selectRows *sql.Rows) {
	selectRows, err := db.Query(sel)
	if err != nil {
		fmt.Println("Error running ODM_FWA_SUBS_TRAFFIC_4GAVG select query")
		fmt.Println(err)
	}
	return selectRows
}
package sqlselect

import(
	"database/sql"
	"fmt"
) 

//SQL SELECT command for all tables
func Select(db *sql.DB, sel string) (selectRows *sql.Rows) {
	selectRows, err := db.Query(sel)
	if err != nil {
		fmt.Println("Error running ODM_FWA_SUBS_TRAFFIC_4GAVG select query")
		fmt.Println(err)
	}
	return selectRows
}

package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	_ "github.com/microsoft/go-mssqldb"
)

var userid, password, server, database *string

func initFlags() {
	user := "sa"
	pass := "admin@1234"
	svr := "assignment_sqlserver" //"localhost" // 127.0.0.1
	dbname := "STUDENTDATA"

	userid = flag.String("U", user, "login_id")
	password = flag.String("P", pass, "password")
	server = flag.String("S", svr, "server_name[\\instance_name]")
	database = flag.String("d", dbname, "db_name")

}

func getDBConnection() (*sql.DB, string) {
	var retString string = ""

	// if flags are already initialized, no need to initialize them again
	if flag.Lookup("U") == nil {
		initFlags()
	}
	flag.Parse()

	dsn := "server=" + *server + ";user id=" + *userid + ";password=" + *password + ";database=" + *database
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		retString = "Cannot connect: " + err.Error() + "\n"
		return db, retString
	} else {
		retString = "DB Connected.\n"
	}

	err = db.Ping()
	if err != nil {
		retString = "Cannot connect: " + err.Error() + "\n"
		return db, retString
	} else {
		retString = retString + "Ping successful.\n"
	}

	return db, retString
}

func exec(db *sql.DB, cmd string) (string, error) {
	var retString string = ""

	rows, err := db.Query(cmd)
	if err != nil {
		return retString, err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return retString, err
	}
	if cols == nil {
		return retString, nil
	}
	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
		if i != 0 {
			fmt.Print("\t")
			retString = retString + "\t"
		}
		fmt.Print(cols[i])
		retString = retString + cols[i]
	}
	fmt.Println()
	retString = retString + "\n"
	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			fmt.Println(err)
			retString = retString + err.Error() + "\n"
			continue
		}
		for i := 0; i < len(vals); i++ {
			if i != 0 {
				fmt.Print("\t")
				retString = retString + "\t"
			}
			str := printValue(vals[i].(*interface{}))
			retString = retString + str
		}
		fmt.Println()
		retString = retString + "\n"
	}
	if rows.Err() != nil {
		return retString, rows.Err()
	}
	return retString, nil
}

func printValue(pval *interface{}) string {
	retString := ""

	switch v := (*pval).(type) {
	case nil:
		fmt.Print("NULL")
	case bool:
		if v {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	case []byte:
		fmt.Print(string(v))
		retString = string(v)
	case int:
		fmt.Print(v)
	//case time.Time:
	//	fmt.Print(v.Format("2006-01-02 15:04:05.999"))
	default:
		fmt.Print(v)
		retString = fmt.Sprint(v)
		//fmt.Print(retString)
	}
	return retString
}

func update(db *sql.DB, cmd string) (string, error) {
	retString := ""

	result, err := db.Exec(cmd, "Rkj", 3)
	if err == nil {
		rowcount, _ := result.RowsAffected()
		fmt.Println("No. of rows affected: ", rowcount)
		retString = "No. of rows affected: " + fmt.Sprint(rowcount) + "\n"
	}
	return retString, err
}

func checkdbConnection(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ENV_NAME=%s\n", getEnv())
	fmt.Fprintf(w, "checkdbConnection() started.\n")
	db, str := getDBConnection()
	fmt.Fprint(w, str)

	db.Close()

	fmt.Fprintf(w, "checkdbConnection() completed.")
}

func execSql(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ENV_NAME=%s\n", getEnv())
	fmt.Fprintf(w, "execSql() started.\n")

	// Connect to DB
	db, str := getDBConnection()
	fmt.Fprint(w, str)

	defer db.Close()

	sqlQuery := "select * from Student"
	val := r.URL.Query()
	if val.Has("query") {
		sqlQuery = string(val["query"][0])
	}

	retString, err := exec(db, sqlQuery)
	if err != nil {
		fmt.Println("Execution failed: ", err.Error())
		fmt.Fprintf(w, "Execution failed: %s\n", err.Error())
	} else {
		fmt.Fprintf(w, "%s", retString)
	}

	fmt.Fprintf(w, "execSql() completed.")
}

func updateSql(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ENV_NAME=%s\n", getEnv())

	fmt.Fprintf(w, "updateSql() started.\n")

	// Connect to DB
	db, str := getDBConnection()
	fmt.Fprint(w, str)

	defer db.Close()

	retString, err := update(db, "update Student set name=? where rollno=?")
	if err != nil {
		fmt.Println("Execution failed: ", err.Error())
		fmt.Fprintf(w, "Execution failed: %s", err.Error())
	} else {
		fmt.Fprintf(w, "%s", retString)
	}

	fmt.Fprintf(w, "updateSql() completed.")
}

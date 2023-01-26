package impl

import "database/sql"

type FlinkJobDaoImpl struct {
	db *sql.DB
}

func NewFlinkJobDaoImpl(db *sql.DB) *FlinkJobDaoImpl {
	return &FlinkJobDaoImpl{db: db}
}

func (f *FlinkJobDaoImpl) GetByTag(tag string) []string {
	rows, err := f.db.Query("select job_name from t_flink_app where instr(tags,?)>0", tag)
	defer rows.Close()
	var r []string
	if err != nil {
		return r
	}
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return r
		}
		r = append(r, name)
	}
	return r
}

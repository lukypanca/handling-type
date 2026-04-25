package config

import "database/sql"

type Databases struct {
	mufcms *sql.DB
	mufam  *sql.DB
}

func (d *Databases) MUFCMS() *sql.DB {
	return d.mufcms
}

func (d *Databases) MUFAM() *sql.DB {
	return d.mufam
}

func (d *Databases) Close() {
	d.mufcms.Close()
	d.mufam.Close()
}

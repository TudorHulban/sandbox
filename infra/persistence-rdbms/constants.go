package main

import "github.com/TudorHulban/GolangSandbox/config"

const _DDLRoot = "hera"

var paramsPG = config.DBInfo{
	DBName:     "sandbox",
	DBUser:     "postgres",
	DBPassword: "password",
	Host:       "localhost",
	Port:       5432,
}

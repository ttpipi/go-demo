package db

import (
	"database/sql"
	"fmt"
	mydb "go_filestore/db/mysql"
)

func OnFileUploadFinished(filehash, filename string, filesize int64, fileaddr string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_file(file_sha1, file_name, file_size, file_addr, status) values(?,?,?,?,1)")
	if err != nil {
		fmt.Println("Failed to Prepare statement, err:", err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if rf, err := ret.RowsAffected(); err == nil {
		if rf <= 0 {
			fmt.Printf("File with hash: %s has been upload before", filehash)
		}
		return true
	}
	return false
}

type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileAddr sql.NullString
	FileSize sql.NullInt64
}

func OnGetFileMeta(fileHash string) (*TableFile, error) {
	stmt, err := mydb.DBConn().Prepare(
		"select file_sha1, file_name, file_addr, file_size from tbl_file where file_sha1=? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return &TableFile{}, err
	}
	defer stmt.Close()

	file := TableFile{}
	err = stmt.QueryRow(fileHash).Scan(&file.FileHash, &file.FileName, &file.FileAddr, &file.FileSize)
	if err != nil {
		fmt.Println(err.Error())
		return &TableFile{}, err
	}
	return &file, nil
}

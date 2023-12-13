package meta

import mydb "demo/example/filestore/db"

type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

func GetFileMeta(fSha1 string) FileMeta {
	return fileMetas[fSha1]
}

func GetFileMetaDB(fSha1 string) (*FileMeta, error) {
	tfile, err := mydb.OnGetFileMeta(fSha1)
	if err != nil {
		return nil, err
	}

	fileMeta := FileMeta{
		FileSha1: tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		Location: tfile.FileAddr.String,
	}

	return &fileMeta, nil
}

func UpdateFileMeta(fMeta FileMeta) {
	fileMetas[fMeta.FileSha1] = fMeta
}

func UpdateFileMetaDB(meta FileMeta) bool {
	return mydb.OnFileUploadFinished(meta.FileSha1, meta.FileName, meta.FileSize, meta.Location)
}

func RemoveFileMeta(fSha1 string) {
	delete(fileMetas, fSha1)
}

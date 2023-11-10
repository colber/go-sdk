package models

type (
	File struct {
		Id       			string  	
		Type				string 		
		Tags				[]string 	
		CreatedAt			time.Time	
		UpdatedAt			time.Time	
		Status  			string		
	}

	Chunk struct {
		Id       			string  	`reindex:"id,,pk"`
		FileId       		string  	`reindex:"file_id"`
		Seq					int64
		Content				[]byte
	}
)